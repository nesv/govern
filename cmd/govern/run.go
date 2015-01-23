package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func runPlaybook(ctx *cli.Context) {
	inv, err := loadInventory(ctx.String("inventory"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
