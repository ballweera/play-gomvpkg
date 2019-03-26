package main

import (
	"github.com/ballweera/play-gomvpkg/common/consumer"
	"github.com/ballweera/play-gomvpkg/event"
)

func main() {
	consumer.Consume()
	event.Process()
}
