package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"

	"github.com/labstack/gommon/log"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type (
	Task func() error

	// wrapTask struct {
	// 	Task
	// 	chErr chan error
	// }

	// workerPool struct {
	// 	maxError   int32
	// 	cntError   int32
	// 	limitError chan struct{}
	// }
)

// func (wp *workerPool) run(maxError int32, cntError int32, limitError chan struct{}) {

// }

// func worker(chTsk chan wrapTask, wg *sync.WaitGroup, chBreak chan struct{}) {
// 	defer func() {
// 		wg.Done()
// 	}()
// 	for wt := range chTsk {
// 		select {
// 		case <-chBreak:
// 			return
// 		case wt.chErr <- wt.Task():
// 		}
// 	}
// }

// type executor struct {
// 	chkErr     bool
// 	cntErr     int32
// 	maxErr     int32
// 	ttlWrk     int
// 	chTsk      chan wrapTask
// 	chErr      chan error
// 	chMaxError chan struct{}
// 	err        error
// }

// func newExecutor(n, m int) *executor {
// 	chkErr := false
// 	if m > 0 {
// 		chkErr = true
// 	}

// 	ex := &executor{
// 		maxErr:     int32(m),
// 		ttlWrk:     n,
// 		chkErr:     chkErr,
// 		chErr:      make(chan error),
// 		chTsk:      make(chan wrapTask),
// 		chMaxError: make(chan struct{}),
// 	}

// 	return ex
// }

// func (ex *executor) run(tasks []Task) error {
// 	wg := &sync.WaitGroup{}

// 	wg.Add(1)
// 	go ex.checker(wg)

// 	wg.Add(1)
// 	go ex.sender(tasks, wg)

// 	wgw := &sync.WaitGroup{}
// 	for i := 0; i < ex.ttlWrk; i++ {
// 		wgw.Add(1)
// 		go worker(ex.chTsk, wgw, ex.chMaxError)
// 	}
// 	wgw.Wait()
// 	close(ex.chErr)

// 	wg.Wait()
// 	return ex.err
// }

// // checker will receive tasks result up to empty channel, after it will stop
// // if it  need to count errors limit, will set stop signal after above limit.
// func (ex *executor) checker(wg *sync.WaitGroup) {
// 	defer func() {
// 		wg.Done()
// 	}()

// 	for e := range ex.chErr {
// 		if e != nil && ex.chkErr {
// 			atomic.AddInt32(&ex.cntErr, 1)
// 			if atomic.CompareAndSwapInt32(&ex.cntErr, ex.maxErr, ex.maxErr) {
// 				close(ex.chMaxError)
// 				ex.err = ErrErrorsLimitExceeded
// 			}
// 		}
// 	}
// }

// // sender will sending tasks up to last or will stop if get the  break signal.
// func (ex *executor) sender(tasks []Task, wg *sync.WaitGroup) {
// 	defer func() {
// 		close(ex.chTsk)
// 		wg.Done()
// 	}()
// 	for _, t := range tasks {
// 		select {
// 		case <-ex.chMaxError:
// 			return

// 		case ex.chTsk <- wrapTask{Task: t, chErr: ex.chErr}:
// 		}
// 	}
// }

// // Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
// func Run(tasks []Task, n, m int) error {
// 	// Place your code here.
// 	ex := newExecutor(n, m)
// 	return ex.run(tasks)
// }

type WP struct {
	limitError int32
	countErr   int32
	countWrk   int
	chTsk      chan Task
	chBreak    chan struct{}
}

func (wp *WP) run(wg *sync.WaitGroup) {
	for i := 0; i < wp.countWrk; i++ {
		wg.Add(1)
		go func(no int, chkErr bool) {
			log.Infof("worker-%v, start", no)
			defer func() {
				log.Infof("worker-%v, stop", no)
				wg.Done()
				wp.chBreak <- struct{}{}
			}()

			for t := range wp.chTsk {
				err := t()
				if err != nil && chkErr {
					atomic.AddInt32(&wp.countErr, 1)
					if atomic.CompareAndSwapInt32(&wp.countErr, wp.limitError, wp.limitError) {
						log.Infof("worker-%v, limit errors", no)
						return
					}
				}
			}
		}(i, wp.limitError > 0)
	}
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	// Place your code here.

	wp := &WP{
		limitError: int32(m),
		countWrk:   n,
		chTsk:      make(chan Task),
		chBreak:    make(chan struct{}, 1),
	}
	wg := &sync.WaitGroup{}
	wp.run(wg)
	defer func() {
		log.Info("waiting WaitGroup")
		wg.Wait()
	}()
	for _, t := range tasks {
		select {
		case <-wp.chBreak:
			log.Info("sender, limit errors, stop sending")
			close(wp.chTsk)
			return ErrErrorsLimitExceeded
		case wp.chTsk <- t:
		}
	}

	return nil
}
