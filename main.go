package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

func main()  {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	fmt.Println("Hello world")
	if 1 == 3 {
		return errors.New("random error")
	}

	return nil
}