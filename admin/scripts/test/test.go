package main

import (
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	ringpop "github.com/uber/ringpop-go/admin"
	"github.com/uber/tchannel-go"
)

func main() {
	ch, err := tchannel.NewChannel("ringpop-admin-go", nil)
	if err != nil {
		log.Fatalln(err)
	}

	admin := ringpop.NewAdminClient("10.0.1.2:3000", ch, time.Second)

	stats, err := admin.Stats()
	if err != nil {
		log.Fatalln(err)
	}

	spew.Dump(stats)
}
