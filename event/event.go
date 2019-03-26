package event

import "github.com/ballweera/play-gomvpkg/kafka"

func Process() {
	kafka.Consume()
}
