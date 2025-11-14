package es

import (
	"github.com/olivere/elastic/v7"
	"go.uber.org/dig"
)

func Inject(container *dig.Container) error {
	return container.Provide(newClient)
}

func newClient(conf *Config) (*Client, error) {
	clientOpt := []elastic.ClientOptionFunc{
		elastic.SetURL(conf.Host),
		elastic.SetSniff(false),
	}
	if conf.Username != "" {
		clientOpt = append(clientOpt, elastic.SetBasicAuth(conf.Username, conf.Pass))
	}
	instance, err := elastic.NewClient(clientOpt...)
	if err != nil {
		return nil, err
	}
	return &Client{
		conf:     conf,
		instance: instance,
	}, nil
}
