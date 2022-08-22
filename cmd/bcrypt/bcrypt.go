package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "hash",
				Aliases: []string{"h"},
				Usage:   "Returns a hash for a given string",
				Action: func(ctx *cli.Context) error {
					hash(ctx.Args().First())
					return nil
				},
			},
			{
				Name:    "compare",
				Aliases: []string{"c"},
				Usage:   "Compares a secret password to a hash",
				Action: func(ctx *cli.Context) error {
					password := ctx.Args().First()
					hash := ctx.Args().Get(ctx.Args().Len() - 1)

					compare(password, hash)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func hash(password string) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("error hashing: %v\n", err)
		return
	}

	hash := string(hashedBytes)
	fmt.Printf("hash: %q\n", hash)
}

func compare(password, hash string) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Printf("password and hash do not match: %v\n", err)
		return
	}

	fmt.Printf("Password is correct")
}
