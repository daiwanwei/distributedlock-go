package distributedlock

import (
	"context"
	"github.com/bsm/redislock"
)

type Lock struct {
	wrappedLock *redislock.Lock
}

func (l *Lock) Unlock(ctx context.Context) (err error) {
	err = l.wrappedLock.Release(ctx)
	if err != nil && err != redislock.ErrLockNotHeld {
		return err
	}
	return nil
}
