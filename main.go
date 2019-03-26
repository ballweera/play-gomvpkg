package main

import (
	"github.com/ballweera/play-gomvpkg/event"
	"github.com/ballweera/play-gomvpkg/kafka"
)

func main() {
	kafka.Consume()
	event.Process()
}
