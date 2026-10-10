# uru-databases-2-go-service-common

Go module with helpers shared by the Pixel Plaza gRPC microservices. The current branch only holds a JWT claims validator.

**Note:** This repository is archived and read-only. The services (`api-gateway`, `auth-service`, `user-service`) import many more packages from tagged versions of this module (config, env and flag loading, Google Cloud credentials, MongoDB, gRPC interceptors, TLS, JWT issuing) that are not present in this branch.

## Contents

- **`jwt/token/validator/`** — `DefaultClaimsValidator`, which checks the token's refresh-token and id claims, enforces access versus refresh token use, and then checks validity through Redis when a Redis validator is configured.
- **Dependencies** — `uru-databases-2-protobuf-common`, `ralvarezdev/go-grpc`, `ralvarezdev/go-jwt` and `ralvarezdev/go-loader`.

## License

GPL-3.0. See [LICENSE](LICENSE).
