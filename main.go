package main

import "github.com/mtheusvalle/gopportunities/router"

func main() {
	router.Initialize()

	r := router.Teste{
		Name: "Teste",
	}

	print(r)
}
