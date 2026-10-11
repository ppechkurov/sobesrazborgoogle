package channels_test

import (
	"testing"
	"time"

	"github.com/ppechkurov/sobesrazborgoogle/internal/channels"
)

func TestMerge(t *testing.T) {
	t.Parallel()
	t.Run("reads all", func(t *testing.T) {
		t.Parallel()

		chans := make([]<-chan int, 0, 10)
		for i := range 10 {
			ch := make(chan int)
			chans = append(chans, ch)

			go func() {
				ch <- i
				close(ch)
			}()
		}

		timer := time.NewTimer(time.Millisecond)
		defer timer.Stop()

		out := channels.Merge(chans...)
		got, want := 0, 0+1+2+3+4+5+6+7+8+9

	lp:
		for {
			select {
			case r, more := <-out:
				if !more {
					break lp
				}
				got += r

			case <-timer.C:
				t.Fatal("reading out channel takes too long")
			}
		}

		if want != got {
			t.Fatalf("want %v, got %v", want, got)
		}
	})
}
