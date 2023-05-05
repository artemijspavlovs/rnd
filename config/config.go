package cfg

import (
	"fmt"
	"go.uber.org/config"
)

type Configuration struct {
	Tasks []string
}

func Load(cp string) (*Configuration, error) {
	var c Configuration

	provider, err := config.NewYAML(config.File(cp))
	if err != nil {
		fmt.Println("failed to load file")
		return nil, err
	}

	if err := provider.Get("").Populate(&c); err != nil {
		fmt.Println("failed to populate config")
		return nil, err
	}

	return &c, nil
}
