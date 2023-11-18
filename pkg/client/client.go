package client

import (
	"fmt"
	"os"

	"github.com/michimani/gotwi"
	"github.com/xh3b4sd/tracer"
)

const (
	OAuthTokenEnvKeyName       = "GOTWI_ACCESS_TOKEN"
	OAuthTokenSecretEnvKeyName = "GOTWI_ACCESS_TOKEN_SECRET"
)

var (
	cfg = &gotwi.NewClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		OAuthToken:           os.Getenv(OAuthTokenEnvKeyName),
		OAuthTokenSecret:     os.Getenv(OAuthTokenSecretEnvKeyName),
	}
)

func New() *gotwi.Client {
	var err error

	var cli *gotwi.Client
	{
		cli, err = gotwi.NewClient(cfg)
		if err != nil {
			tracer.Panic(err)
		}
	}

	if !cli.IsReady() {
		tracer.Panic(fmt.Errorf("client misconfigured"))
	}

	return cli
}
