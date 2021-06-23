package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dimfeld/httptreemux/v5"
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

	m := httptreemux.NewContextMux()
	m.Handle(http.MethodGet, "/test", nil)
	return nil
}