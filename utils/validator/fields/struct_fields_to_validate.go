package fields

import (
	"fmt"
	"reflect"
	"strings"
)

// StructFieldsToValidate is a map of fields to validate from a struct
type StructFieldsToValidate struct {
	Fields                       map[string]string // Key is the field name and value is the name used in the validation error
	NestedStructFieldsToValidate map[string]*StructFieldsToValidate
}

// CreateGRPCStructFieldsToValidate creates the fields to validate from a gRPC struct
func CreateGRPCStructFieldsToValidate(exampleStruct interface{}) (*StructFieldsToValidate, error) {
	// Reflection of data
	valueReflection := reflect.ValueOf(exampleStruct)

	// If data is a pointer, dereference it
	if valueReflection.Kind() == reflect.Ptr {
		valueReflection = valueReflection.Elem()
	}

	// Initialize the map fields and the map of nested fields to validate
	rootStructFieldsToValidate := make(map[string]string)
	rootNestedStructFieldsToValidate := make(map[string]*StructFieldsToValidate)

	// Reflection of the type of data
	typeReflection := reflect.TypeOf(valueReflection)
	var protobufTag string
	var protobufName string
	for i := 0; i < typeReflection.NumField(); i++ {
		// Get the field type through reflection
		field := typeReflection.Field(i)
		fieldType := field.Type

		// Check if the field is a pointer to a struct
		if fieldType.Kind() != reflect.Ptr {
			// Get the Protobuf tag of the field
			protobufTag = field.Tag.Get("protobuf")
			if protobufTag == "" {
				return nil, fmt.Errorf(MissingProtobufTagError, field.Name)
			}
		} else {
			fieldType = fieldType.Elem()

			// Check if the element type is not a struct which would mean that it is an optional scalar type
			if fieldType.Kind() != reflect.Struct {
				continue
			}

			// Get the Protobuf tag of the field
			protobufTag = field.Tag.Get("protobuf")
			if protobufTag == "" {
				return nil, fmt.Errorf(MissingProtobufTagError, field.Name)
			}

			// Check the tag to determine if it contains 'oneof', which means it is an optional field
			if ok := strings.Contains(protobufTag, "oneof"); ok {
				continue
			}

			// Create a new StructFieldsToValidate for the nested field
			fieldNestedStructFieldsToValidate, err := CreateGRPCStructFieldsToValidate(reflect.New(fieldType).Interface())
			if err != nil {
				return nil, err
			}

			// Add the nested fields to the map
			rootNestedStructFieldsToValidate[field.Name] = fieldNestedStructFieldsToValidate
		}

		// Get the field name from the Protobuf tag
		tagParts := strings.Split(protobufTag, ",")
		protobufName = ""
		for _, part := range tagParts {
			if strings.HasPrefix(part, "name=") {
				protobufName = strings.TrimPrefix(part, "name=")
				break
			}
		}

		// Check if the field name is empty
		if protobufName == "" {
			return nil, fmt.Errorf(MissingProtobufNameError, field.Name)
		}

		// Add the field to the map
		rootStructFieldsToValidate[field.Name] = protobufName
	}

	return &StructFieldsToValidate{
		Fields:                       rootStructFieldsToValidate,
		NestedStructFieldsToValidate: rootNestedStructFieldsToValidate,
	}, nil
}
