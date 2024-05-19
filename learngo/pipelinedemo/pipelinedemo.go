package main

import (
	"context"
	"log"

	"github.com/pkg/errors"
)

func lineListSource(ctx context.Context, lines ...string) (
	<-chan string, <-chan error, error) {
	if len(lines) == 0 {
		// Handle an error that occurs before the goroutine begins.
		return nil, nil, errors.Errorf("no lines provided")
	}
	out := make(chan string)
	errc := make(chan error, 1)
	go func() {
		defer close(out)
		defer close(errc)
		for lineIndex, line := range lines {
			if line == "" {
				// Handle an error that occurs during the goroutine.
				errc <- errors.Errorf("line %v is empty", lineIndex+1)
				return
			}
			// Send the data to the output channel but return early
			// if the context has been cancelled.
			select {
			case out <- line:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out, errc, nil
}

func main() {

	out, _, _ := lineListSource(context.TODO(), "12312", "234", "error")
	for line := range out {
		log.Println(line)
	}
	//for line := range errc {
	//	log.Println(line)
	//}
}
