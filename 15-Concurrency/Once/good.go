package main

import (
	"fmt"
	"sync"
)

var Once sync.Once

func GetConfigWithOnce() *Config {

	Once.Do(func() {
		fmt.Println("Create new Config")
		config = &Config{}
	})
	return config
}
