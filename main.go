package main

import (
	"log"
	"os"

	"astuart.co/edgeos-rest/pkg/edgeos"
)

func main() {
	c, err := edgeos.NewClient(os.Getenv("ERLITE_ADDR"), os.Getenv("ERLITE_USER"), os.Getenv("ERLITE_PASS"))
	if err != nil {
		log.Fatal(err)
	}

	if err := c.Login(); err != nil {
		log.Fatal(err)
	}

	feat, err := c.Feature(edgeos.PortForwarding)
	if err != nil {
		log.Fatal(err)
	}

	d := feat["data"].(map[string]interface{})["rules-config"].([]interface{})

	d = d[:len(d)-1]

	feat["data"].(map[string]interface{})["rules-config"] = d

	log.Println(feat["data"])
	log.Println()
	log.Println()

	// log.Println(c.SetFeature(edgeos.PortForwarding, feat["data"]))
}
