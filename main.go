package main

import (
	"errors"
	"fmt"

	"github.com/michimani/gotwi"
	"github.com/xh3b4sd/tracer"
	"github.com/xh3b4sd/twurge/pkg/cache"
	"github.com/xh3b4sd/twurge/pkg/client"
	"github.com/xh3b4sd/twurge/pkg/timeline"
	"github.com/xh3b4sd/twurge/pkg/tweet"
	"github.com/xh3b4sd/twurge/pkg/user"
)

func main() {
	err := mainE()
	if isRateLimit(err) {
		fmt.Printf("rate limit hit\n")
	} else if err != nil {
		tracer.Panic(err)
	}
}

func mainE() error {
	var err error

	var cli *gotwi.Client
	{
		cli = client.New()
	}

	var use *cache.Cache
	{
		use = cache.New(cache.Config{
			Cac: "user",
			Das: false,
			Get: func() ([]string, error) {
				var use *user.User
				{
					use = user.New(cli)
				}

				var uid string
				{
					uid, err = use.Search()
					if err != nil {
						return nil, tracer.Mask(err)
					}
				}

				return []string{uid}, nil
			},
		})
	}

	var tim *cache.Cache
	{
		tim = cache.New(cache.Config{
			Cac: "timeline",
			Das: true,
			Get: func() ([]string, error) {
				var uid string
				{
					uid, err = use.Search()
					if err != nil {
						return nil, tracer.Mask(err)
					}
				}

				var tim *timeline.Timeline
				{
					tim = timeline.New(cli)
				}

				var pag string
				var lis []string
				{
					_, lis, err = tim.Search(uid, pag)
					if err != nil {
						return nil, tracer.Mask(err)
					}
				}

				return lis, nil
			},
		})
	}

	var tid string
	{
		tid, err = tim.Search()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var twe *tweet.Tweet
	{
		twe = tweet.New(cli)
	}

	{
		err = twe.Delete(tid)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	fmt.Printf("deleted tweet %s\n", tid)

	return nil
}

func isRateLimit(err error) bool {
	if err == nil {
		return false
	}

	var gte *gotwi.GotwiError
	{
		gte = &gotwi.GotwiError{}
	}

	if errors.As(err, &gte) {
		return gte.OnAPI && gte.RateLimitInfo != nil
	}

	return false
}
