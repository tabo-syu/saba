package cmd

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/urfave/cli/v3"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

var Saba = &cli.Command{
	Name:  "saba",
	Usage: "",
	Action: func(ctx context.Context, c *cli.Command) error {
		e := echo.New()
		e.Any("/*", func(c echo.Context) error {
			req := c.Request()

			header := req.Header
			body := req.Body

			log.Println("header:", header)
			log.Println("body", body)

			return nil
		})

		ngrok.ListenAndServeHTTP(ctx, config.HTTPEndpoint(config.WithURL()))

		tunnel, err := ngrok.Listen(ctx, config.HTTPEndpoint(config.()), ngrok.WithAuthtokenFromEnv())
		if err != nil {
			return err
		}

		if err := e.Start(":9000"); err != http.ErrServerClosed {
			return err
		}

		return nil
	},
}
