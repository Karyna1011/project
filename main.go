package main

import (
	"os"

	_ "gitlab.com/tokend/subgroup/project/docs"
	"gitlab.com/tokend/subgroup/project/internal/cli"
)

// @title           REST API PROJECT
// @version         1.0
// @description     Lab#2

// @host      localhost:9000
// @BasePath  /project
func main() {
	//e := echo.New()

	//	e.GET("/swagger/*", echoSwagger.WrapHandler)

	//	e.Logger.Fatal(e.Start(":1323"))

	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
