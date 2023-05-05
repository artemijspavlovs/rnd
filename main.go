package main

import (
	"fmt"
	cfg "github.com/artemijspavlovs/rnd/config"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	c, err := cfg.Load("config.yml")
	if err != nil {
		fmt.Printf("failed to load config: %v\n", err)
		return
	}

	fmt.Printf("do '%s' or '%s'\n", c.Tasks[rand.Intn(len(c.Tasks))], c.Tasks[rand.Intn(len(c.Tasks))])
}
