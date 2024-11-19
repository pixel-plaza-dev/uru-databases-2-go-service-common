package gcloud

import (
	"context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/idtoken"
	"google.golang.org/api/option"
	"google.golang.org/grpc/credentials/oauth"
)

// LoadServiceAccountCredentials loads the service account credentials
func LoadServiceAccountCredentials(
	ctx context.Context, url string,
) (*google.Credentials, *oauth.TokenSource, error) {
	// Construct the GoogleCredentials object with the default configuration from the working environment
	credentials, err := google.FindDefaultCredentials(ctx)
	if err != nil {
		return nil, nil, FailedToGenerateDefaultGoogleCredentialsError
	}

	// Use the default service account credentials
	ts, err := idtoken.NewTokenSource(ctx, url, option.WithCredentials(credentials))
	if err != nil {
		return nil, nil, FailedToCreateTokenSourceError
	}

	return credentials, &oauth.TokenSource{TokenSource: ts}, nil
}
