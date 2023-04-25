package lumberjack

import (
	"go.uber.org/dig"
)

func Inject(container *dig.Container) error {
	if err := container.Provide(newClient); err != nil {
		return err
	}
	return nil
}
