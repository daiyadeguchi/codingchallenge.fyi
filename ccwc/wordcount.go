package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	var isCount bool
	var isLine bool
	app := &cli.App{
		Name:  "ccwc",
		Usage: "wc for coding challenge",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "count, c",
				Usage:       "Count the size of the file",
				Destination: &isCount,
			},
			&cli.BoolFlag{
				Name:        "line, l",
				Usage:       "Count the number of lines in the file",
				Destination: &isLine,
			},
		},
		Action: func(ctx *cli.Context) error {
			filepath := ctx.Args().Get(0)
			file, err := os.Open(filepath)
			if err != nil {
				log.Fatal(err)
			}

			if isCount {
				stat, _ := file.Stat()
				fmt.Println(stat.Size(), file.Name())
			}

			if isLine {
				dat, _ := os.ReadFile(filepath)
				content := string(dat)
				lineCount := 0
				for _, line := range content {
					if line == '\n' {
						lineCount++
					}
				}
				fmt.Println(lineCount, file.Name())
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
