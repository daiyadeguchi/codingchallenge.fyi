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
	var isWord bool
	var isCharacter bool
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
			&cli.BoolFlag{
				Name:        "word, w",
				Usage:       "Count the words in the file",
				Destination: &isWord,
			},
			&cli.BoolFlag{
				Name:        "m",
				Usage:       "Count the characters in the file",
				Destination: &isCharacter,
			},
		},
		Action: func(ctx *cli.Context) error {
			filepath := ctx.Args().Get(0)
			file, err := os.Open(filepath)
			if err != nil {
				return err
			}
			dat, err := os.ReadFile(filepath)
			if err != nil {
				return err
			}

			if isCount {
				stat, err := file.Stat()
				if err != nil {
					return err
				}
				fmt.Println(stat.Size(), file.Name())
			}

			if isLine {
				fmt.Println(getLineCount(dat), file.Name())
			}

			if isWord {
				fmt.Println(getWordCount(dat), file.Name())
			}

			if isCharacter {
				fmt.Println(getCharcterCount(dat), file.Name())
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
