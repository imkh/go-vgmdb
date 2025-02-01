package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"

	"github.com/imkh/go-vgmdb"
)

func main() {
	godotenv.Load()

	cookie := os.Getenv("AUTH_COOKIE")
	s, err := vgmdb.NewScraper(vgmdb.WithCookie(cookie))
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to create scraper: %v\n", err)
		os.Exit(1)
	}

	app := &cli.App{
		Name:  "vgmdb",
		Usage: "a VGMdb web scraper",
		Commands: []*cli.Command{
			{
				Name:  "role",
				Usage: "get a role by its ID",
				Action: func(cCtx *cli.Context) error {
					id, err := strconv.Atoi(cCtx.Args().First())
					if err != nil {
						return fmt.Errorf("invalid role ID: %w", err)
					}
					role, err := s.Roles.GetRole(id)
					if err != nil {
						return fmt.Errorf("unable to get role: %w", err)
					}

					output, err := json.MarshalIndent(role, "", "  ")
					if err != nil {
						return fmt.Errorf("unable to marshal role: %w", err)
					}
					fmt.Println(string(output))

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
