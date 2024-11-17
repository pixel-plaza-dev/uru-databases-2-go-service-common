package gcloud

import (
	"context"
	"google.golang.org/api/idtoken"
	"google.golang.org/grpc/credentials/oauth"
)

// LoadServiceAccountCredentials loads the service account credentials
func LoadServiceAccountCredentials(ctx context.Context, url string) (*oauth.TokenSource, error) {
	// Use the default service account credentials
	ts, err := idtoken.NewTokenSource(ctx, url)
	if err != nil {
		return nil, FailedToCreateTokenSourceError
	}

	return &oauth.TokenSource{TokenSource: ts}, nil
}
