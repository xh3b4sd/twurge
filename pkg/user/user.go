package user

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/user/userlookup"
	"github.com/michimani/gotwi/user/userlookup/types"
	"github.com/xh3b4sd/tracer"
)

type User struct {
	cli *gotwi.Client
}

func New(cli *gotwi.Client) *User {
	return &User{cli: cli}
}

func (s *User) Search() (string, error) {
	var err error

	var inp *types.GetMeInput
	{
		inp = &types.GetMeInput{}
	}

	var out *types.GetMeOutput
	{
		out, err = userlookup.GetMe(context.Background(), s.cli, inp)
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	return gotwi.StringValue(out.Data.ID), nil
}
