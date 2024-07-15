package main

import (
	"fmt"
	"midincihuy/webgo/bootstrap"
	"midincihuy/webgo/pkg/env"
	"log"
)

func main() {
	app := bootstrap.NewApplication()
	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", env.GetEnv("APP_HOST", "localhost"), env.GetEnv("APP_PORT", "4000"))))
}