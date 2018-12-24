package providers

import (
	"azure-blockchain-connector/aad"
	"azure-blockchain-connector/proxy"
	"context"
	"golang.org/x/oauth2"
	"net/http"
)

type OAuthAuthCode struct {
	*oauth2.Config
	SvcAddr string
	ArgName string
	client  *http.Client
}

func (ac *OAuthAuthCode) RequestAccess() error {
	ctx := context.Background()

	tok, err := aad.AuthCodeGrantWebview(ctx, ac.Config, ac.ArgName)
	//tok, err := aad.AuthCodeGrantServer(ctx, ac.Config, ac.SvcAddr)
	if err != nil {
		return err
	}
	ac.client = ac.Config.Client(ctx, tok)

	return nil
}

func (ac *OAuthAuthCode) Client() *http.Client {
	return ac.client
}

func (ac *OAuthAuthCode) Modify(params *proxy.Params, req *http.Request) {
}
