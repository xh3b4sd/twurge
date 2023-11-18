package tweet

import (
	"context"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/managetweet"
	"github.com/michimani/gotwi/tweet/managetweet/types"
	"github.com/xh3b4sd/tracer"
)

type Tweet struct {
	cli *gotwi.Client
}

func New(cli *gotwi.Client) *Tweet {
	return &Tweet{cli: cli}
}

func (s *Tweet) Delete(tid string) error {
	var err error

	var inp *types.DeleteInput
	{
		inp = &types.DeleteInput{
			ID: tid,
		}
	}

	{
		_, err = managetweet.Delete(context.Background(), s.cli, inp)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
