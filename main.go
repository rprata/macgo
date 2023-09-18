package main

import (
	"fmt"
	"os"

	"github.com/rprata/macgo/internal/app"
)

func main() {
	app := app.Generate()
	if error := app.Run(os.Args); error != nil {
		fmt.Println(error)
	}
}
