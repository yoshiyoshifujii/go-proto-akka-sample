package main

import (
	"fmt"
	"github.com/asynkron/protoactor-go/actor"
	"time"
)

func main() {

	system := actor.NewActorSystem()

	props := actor.PropsFromProducer(func() actor.Actor {
		return &HelloWorldActor{}
	})

	pid := system.Root.Spawn(props)

	system.Root.Send(pid, &Hello{Who: "Roger"})
	time.Sleep(1 * time.Second)

}

type Hello struct {
	Who string
}

type HelloWorldActor struct{}

func (state *HelloWorldActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *Hello:
		fmt.Printf("Hello %v\n", msg.Who)
	}
}
