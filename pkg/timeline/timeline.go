package timeline

import (
	"context"
	"time"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/tweet/timeline"
	"github.com/michimani/gotwi/tweet/timeline/types"
	"github.com/xh3b4sd/tracer"
)

const (
	ninetydays time.Duration = time.Hour * 24 * 90
)

type Timeline struct {
	cli *gotwi.Client
}

func New(cli *gotwi.Client) *Timeline {
	return &Timeline{cli: cli}
}

func (s *Timeline) Search(use string, pag string) (string, []string, error) {
	var err error

	var inp *types.ListTweetsInput
	{
		inp = &types.ListTweetsInput{
			ID:              use,
			MaxResults:      types.ListMaxResults(100),
			PaginationToken: pag,
			EndTime:         staTim(ninetydays),
			TweetFields: []fields.TweetField{
				fields.TweetFieldCreatedAt,
			},
		}
	}

	var out *types.ListTweetsOutput
	{
		out, err = timeline.ListTweets(context.Background(), s.cli, inp)
		if err != nil {
			return "", nil, tracer.Mask(err)
		}
	}

	var lis []string
	for _, x := range out.Data {
		lis = append(lis, gotwi.StringValue(x.ID))
	}

	{
		pag = gotwi.StringValue(out.Meta.NextToken)
	}

	return pag, lis, nil
}

func staTim(dur time.Duration) *time.Time {
	t := time.Now().UTC().Add(-dur)
	return &t
}
