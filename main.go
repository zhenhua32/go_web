package main

import "tzh.com/web/cmd"

// @title Apiserver API
// @version 1.0
// @description This is a simple api server.

// @contact.name coolcat
// @contact.url http://coolcat.io/support
// @contact.email help@coolcat.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8081
// @BasePath /v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cmd.Execute()
}
