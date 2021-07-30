package main

import "SyncImageBE/presentation"

func main() {
	s := presentation.CreateService()
	r := presentation.CreateRouter(s)
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
