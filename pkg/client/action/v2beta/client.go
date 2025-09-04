package v2beta

import (
	"context"

	"github.com/atefhaloui/zitadel-go/v3/pkg/client/zitadel"
	action "github.com/atefhaloui/zitadel-go/v3/pkg/client/zitadel/action/v2beta"
)

type Client struct {
	Connection *zitadel.Connection
	action.ActionServiceClient
}

func NewClient(ctx context.Context, issuer, api string, scopes []string, options ...zitadel.Option) (*Client, error) {
	conn, err := zitadel.NewConnection(ctx, issuer, api, scopes, options...)
	if err != nil {
		return nil, err
	}

	return &Client{
		Connection:          conn,
		ActionServiceClient: action.NewActionServiceClient(conn.ClientConn),
	}, nil
}
