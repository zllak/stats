package logstats

import (
	"log"

	"github.com/segmentio/stats"
)

func NewBackend(logger *log.Logger) stats.Backend {
	return stats.BackendFunc(func(e stats.Event) {
		logger.Printf("%s %s [%v] %g/%g\n", e.Type, e.Name, e.Tags, e.Value, e.Sample)
	})
}