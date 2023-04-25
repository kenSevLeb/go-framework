package es

import (
	"github.com/olivere/elastic/v7"
	"go.uber.org/dig"
)

func Inject(container *dig.Container) error {
	return container.Provide(newClient)
}

func newClient(conf *Config) (*Client, error) {
	instance, err := elastic.NewClient(elastic.SetURL(conf.Host), elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}

	return &Client{
		conf:     conf,
		instance: instance,
	}, nil
}
