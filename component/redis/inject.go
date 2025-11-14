package redis

import (
	"fmt"
	"go.uber.org/dig"
)

func Inject(container *dig.Container) error {
	if err := container.Provide(newClient); err != nil {
		return err
	}

	if err := container.Provide(newClusterClient); err != nil {
		return fmt.Errorf("cluster: %v", err)
	}

	return nil
}
