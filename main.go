package main

import (
	"log"
	"os"

	"astuart.co/edgeos-rest/pkg/edgeos"
)

func main() {
	c, _ := edgeos.NewClient(os.Getenv("ERLITE_ADDR"), os.Getevn("ERLITE_USER"), os.Getenv("ERLITE_PASS"))

	if err := c.Login(); err != nil {
		log.Fatal(err)
	}

	log.Println(c.Get())
}
