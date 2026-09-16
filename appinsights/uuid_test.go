package appinsights

import (
	"io"
	"sync"
	"testing"
)

func TestNewUUID(t *testing.T) {
	var start sync.WaitGroup
	var finish sync.WaitGroup

	start.Add(1)

	goroutines := 250
	uuidsPerRoutine := 10
	results := make(chan string, 100)

	// Start normal set of UUID generation:
	for range goroutines {
		finish.Go(func() {
			start.Wait()
			for range uuidsPerRoutine {
				results <- newUUID().String()
			}
		})
	}

	// Start broken set of UUID generation
	brokenGen := newUuidGenerator(&brokenReader{})
	for range goroutines {
		finish.Go(func() {
			start.Wait()
			for range uuidsPerRoutine {
				results <- brokenGen.newUUID().String()
			}
		})
	}

	// Close the channel when all the goroutines have exited
	go func() {
		finish.Wait()
		close(results)
	}()

	used := make(map[string]bool)
	start.Done()
	for id := range results {
		if _, ok := used[id]; ok {
			t.Errorf("UUID was generated twice: %s", id)
		}

		used[id] = true
	}
}

type brokenReader struct{}

func (reader *brokenReader) Read(b []byte) (int, error) {
	return 0, io.EOF
}
