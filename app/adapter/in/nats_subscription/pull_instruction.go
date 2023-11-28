package nats_subscription

import (
	"fmt"
	"instruction-processor/app/shared/archetype/container"
	einar "instruction-processor/app/shared/archetype/nats"

	"github.com/nats-io/nats.go"
)

func init() {
	go container.InjectInboundAdapter(func() error {
		_, err := einar.Conn.Subscribe("public-server", func(msg *nats.Msg) {
			fmt.Printf("Recibido un mensaje en [%s]: %s\n", msg.Subject, string(msg.Data))
		})
		return err
	})
}
