package main

import (
	"os"

	"github.com/IaC/go-kcloutie/pkg/cmd/gokcloutie"
	"github.com/IaC/go-kcloutie/pkg/params"
)

// @title       go-kcloutie example go app
// @version     1.0
// @description go-kcloutie Golang CLI and Rest API application

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host     localhost:8080

// @securityDefinitions.basic BasicAuth
func main() {
	cliParams := params.New()
	cli := gokcloutie.Root(cliParams)

	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
