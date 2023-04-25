package mysql

import (
	"go.uber.org/dig"
	"gorm.io/gorm"
)

func Inject(container *dig.Container) error {
	if err := container.Provide(newClient); err != nil {
		return err
	}
	return container.Provide(func(client *Client) *gorm.DB {
		return client.DB
	})
}

func newClient(configs MultiConfig) (*Client, error) {
	client := new(Client)
	client.configs = configs
	if err := client.connect(); err != nil {
		return nil, err
	}
	return client, nil
}
