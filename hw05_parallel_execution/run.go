package hw05parallelexecution

import (
	"errors"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type (
	Task func() error
	WP   struct {
		limitError int32
		countErr   int32
		countWrk   int
		chTsk      chan Task
		chStop     chan struct{}
		once       sync.Once
	}
)

func (wp *WP) run(wg *sync.WaitGroup) {
	for i := 0; i < wp.countWrk; i++ {
		wg.Add(1)
		go func(chkErr bool) {
			defer func() {
				wg.Done()
			}()

			for {
				select {
				case <-wp.chStop:
					return
				case t, more := <-wp.chTsk:
					if !more {
						return
					}
					err := t()
					if err != nil && chkErr {
						atomic.AddInt32(&wp.countErr, 1)
						if atomic.CompareAndSwapInt32(&wp.countErr, wp.limitError, wp.limitError) {
							wp.once.Do(func() {
								close(wp.chStop)
							})
						}
					}
				}
			}
		}(wp.limitError > 0)
	}
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	// Place your code here.
	wp := &WP{
		limitError: int32(m),
		countWrk:   n,
		chTsk:      make(chan Task),
		chStop:     make(chan struct{}),
	}
	wg := &sync.WaitGroup{}
	wp.run(wg)
	defer func() {
		close(wp.chTsk)
		wg.Wait()
	}()
	for _, t := range tasks {
		select {
		case <-wp.chStop:
			return ErrErrorsLimitExceeded
		case wp.chTsk <- t:
		}
	}

	return nil
}

func getGorutineNo() (int32, error) {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		return 0, err
	}
	return int32(id), err
}
