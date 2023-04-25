package mongo

import (
	"go.uber.org/dig"
)

func Inject(container *dig.Container) error {
	return container.Provide(newClient)
}

func newClient(conf *Config) (*Client, error) {
	c := new(Client)
	c.conf = conf

	if err := c.connect(); err != nil {
		return nil, err
	}
	return c, nil
}
