package waitgroup

import (
	"errors"
	"sync"
	"time"
)

// WaitGroup wraps sync.WaitGroup and adds a method WaitTimeout to abort waiting
// for long-running, blocked or leaked goroutines blocking Wait() from the
// underlying WaitGroup. A caller might use this functionality to terminate a
// program in a bounded time.
type WaitGroup struct {
	sync.WaitGroup
}

// ErrTimeout is returned when the timeout in WaitTimeout is exceeded
var ErrTimeout = errors.New("timed out")

// WaitTimeout blocks until the WaitGroup counter is zero or when timeout is
// exceeded. It spawns an internal goroutine. In case of timeout exceeded the
// error ErrTimeout is returned and the internally spawned goroutine might leak
// if the underlying WaitGroup never returns.
//
// It is safe to call WaitTimeout concurrently but doing so might leak
// additional goroutines as described above.
func (wg *WaitGroup) WaitTimeout(timeout time.Duration) error {
	doneCh := make(chan struct{})
	timer := time.NewTimer(timeout)
	defer timer.Stop()

	go func() {
		wg.Wait()
		close(doneCh)
	}()

	select {
	case <-timer.C:
		return ErrTimeout
	case <-doneCh:
		return nil
	}
}

// Waiter is the interface blocking on Wait(). sync.WaitGroup implements this
// interface.
type Waiter interface {
	Wait()
}

// WaitErrorer is the interface blocking on Wait() and returning any error that
// occurred from Wait(). errgroup.Group implements this interface.
type WaitErrorer interface {
	Wait() error
}

// Await is a convenience function that can be used instead of WaitGroup
// provided by this package. Await blocks until Waiter returns or the specified
// timeout is exceeded. In case of timeout exceeded the error ErrTimeout is
// returned and the internally spawned goroutine might leak if Waiter never returns.
//
// It is safe to call Await concurrently and multiple times but doing so might leak
// additional goroutines as described above.
func Await(wf Waiter, timeout time.Duration) error {
	doneCh := make(chan struct{})
	timer := time.NewTimer(timeout)
	defer timer.Stop()

	go func() {
		wf.Wait()
		close(doneCh)
	}()

	select {
	case <-timer.C:
		return ErrTimeout
	case <-doneCh:
		return nil
	}
}

// AwaitWithError is a convenience function that can be used instead of
// WaitGroup provided by this package. AwaitWithError blocks until WaitErrorer returns
// or the specified timeout is exceeded. Any error from WaitErrorer will be returned
// unless the timeout has been exceeded before. In case of timeout exceeded the
// error ErrTimeout is returned and the internally spawned goroutine might leak
// if WaitErrorer never returns.
//
// It is safe to call AwaitWithError concurrently and multiple times but doing
// so might leak additional goroutines as described above.
func AwaitWithError(we WaitErrorer, timeout time.Duration) error {
	doneCh := make(chan struct{})
	errCh := make(chan error)

	timer := time.NewTimer(timeout)
	defer timer.Stop()

	go func() {
		err := we.Wait()
		errCh <- err
		close(doneCh)
	}()

	select {
	case <-timer.C:
		return ErrTimeout
	case <-doneCh:
		return nil
	case err := <-errCh:
		return err
	}
}
