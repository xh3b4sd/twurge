package cache

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/xh3b4sd/tracer"
)

type Config struct {
	// Cac is the cache name, basically used to differentiate cache locations.
	Cac string
	// Das, delete after search, determines whether to delete the cached item from
	// cache, that got returned by search. Items can be kept in cache or be
	// removed after consumption.
	Das bool
	// Get is a data fecthing function that provides all the data written to the
	// cache, once it was found to be empty.
	Get func() ([]string, error)
}

type Cache struct {
	cac string
	das bool
	get func() ([]string, error)
}

func New(c Config) *Cache {
	if c.Cac == "" {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Cac must not be empty", c)))
	}
	if c.Get == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Get must not be empty", c)))
	}

	return &Cache{
		cac: c.Cac,
		das: c.Das,
		get: c.Get,
	}
}

func (c *Cache) Search() (string, error) {
	var err error

	// Try to read from cache and prepare a string slice from the row based file
	// content, if possible.
	var lis []string
	{
		lis, err = c.cacRea()
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	// Execute the configured getter function if there is no data cached right
	// now. This provides the same string slice format as if from cache.
	if len(lis) == 0 {
		lis, err = c.get()
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	// Something went awfully wrong here, should basically never happen.
	if len(lis) == 0 {
		return "", tracer.Mask(fmt.Errorf("not found"))
	}

	// Select the first item from the string slice and write the rest back to the
	// cache.
	var str string
	{
		str, err = c.cacWri(lis)
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	return str, nil
}

func (c *Cache) cacPat() string {
	return fmt.Sprintf("./.%s.cache.twurge.txt", c.cac)
}

func (c *Cache) cacRea() ([]string, error) {
	var err error

	var byt []byte
	{
		byt, err = os.ReadFile(c.cacPat())
		if os.IsNotExist(err) {
			return nil, nil
		} else if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	if len(byt) == 0 {
		return nil, nil
	}

	var lis []string
	{
		for _, x := range bytes.Split(byt, []byte("\n")) {
			lis = append(lis, string(x))
		}
	}

	return lis, err
}

func (c *Cache) cacWri(lis []string) (string, error) {
	var err error

	var str string
	{
		str = lis[0]
	}

	if c.das {
		lis = lis[1:]
	}

	{
		err = os.WriteFile(c.cacPat(), []byte(strings.Join(lis, "\n")), 0600)
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	return str, nil
}
