package main

import (
	"fmt"

	"github.com/fronomenal/go_jwt/httpd/inits"
)

func init() {
	inits.SetupEnv()
}

func main() {
	fmt.Println("project init")
}
