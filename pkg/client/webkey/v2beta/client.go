package v2

import (
	"context"

	"github.com/atefhaloui/zitadel-go/v3/pkg/client/zitadel"
	webkey "github.com/atefhaloui/zitadel-go/v3/pkg/client/zitadel/webkey/v2beta"
)

type Client struct {
	Connection *zitadel.Connection
	webkey.WebKeyServiceClient
}

func NewClient(ctx context.Context, issuer, api string, scopes []string, options ...zitadel.Option) (*Client, error) {
	conn, err := zitadel.NewConnection(ctx, issuer, api, scopes, options...)
	if err != nil {
		return nil, err
	}

	return &Client{
		Connection:          conn,
		WebKeyServiceClient: webkey.NewWebKeyServiceClient(conn.ClientConn),
	}, nil
}
