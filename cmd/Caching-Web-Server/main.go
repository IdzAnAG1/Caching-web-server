package main

import "CachingWebServer/internal/app"

func main() {
	A := app.NewApp()
	err := A.Launch()
	if err != nil {

	}
}
