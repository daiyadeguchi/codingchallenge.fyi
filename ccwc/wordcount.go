package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:  "ccwc",
		Usage: "wc for coding challenge",
		Action: func(ctx *cli.Context) error {
			filepath := ctx.Args().Get(0)
			file, err := os.Open(filepath)
			if err != nil {
				log.Fatal(err)
			}

			stat, err := file.Stat()
			fmt.Printf("size of file is %d", stat.Size())

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
