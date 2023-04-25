package kafka

import "go.uber.org/dig"

func Inject(container *dig.Container) error {
	return container.Provide(NewClient)
}
