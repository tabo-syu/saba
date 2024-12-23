package main

import (
	"context"
	"log"
	"os"

	"github.com/tabo-syu/saba/cmd"
)

const (
	success int = 0
	failure int = 1
)

func main() {
	if err := run(); err != nil {
		os.Exit(failure)
	}

	os.Exit(success)
}

func run() error {
	if err := cmd.Saba.Run(context.Background(), os.Args); err != nil {
		log.Println(err)

		return err
	}

	return nil
}
