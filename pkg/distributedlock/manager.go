package distributedlock

import (
	"context"
	"distributedlock-go/pkg/redis"
	"github.com/bsm/redislock"
	"time"
)

var (
	managerInstance *LockManager
)

type LockManager struct {
	locker *redislock.Client
}

func GetLockManager() (instance *LockManager, err error) {
	if managerInstance == nil {
		managerInstance, err = NewLockManager()
		if err != nil {
			return nil, err
		}
	}
	return managerInstance, nil
}

func NewLockManager() (*LockManager, error) {
	locker, err := redis.GetLocker()
	if err != nil {
		return nil, err
	}
	return &LockManager{
		locker: locker,
	}, nil
}

func (l *LockManager) Lock(ctx context.Context, o Object) (lock *Lock, err error) {
	wrappedLock, err := l.obtain(ctx, o)
	if err != nil {
		return nil, err
	}
	return &Lock{
		wrappedLock: wrappedLock,
	}, nil
}

func (l *LockManager) obtain(ctx context.Context, o Object) (*redislock.Lock, error) {
	opt := &redislock.Options{
		RetryStrategy: redislock.LinearBackoff(time.Millisecond * 100),
	}
	lock, err := l.locker.Obtain(ctx, o.Key(), 10*time.Second, opt)
	if err != nil {
		return nil, err
	}
	return lock, nil
}

type Object interface {
	Key() string
}

//func (l *LockManager)tryLock(ctx context.Context,hello string) (lock *redislock.Lock,err error){
//	if ctx==nil{
//		ctx=context.Background()
//	}
//	for{
//		select {
//		case <-ctx.Done():
//			return nil,ctx.Err()
//		case <-time.After(time.Second*100):
//			return nil,ctx.Err()
//		default:
//			lock,err:=l.obtain(ctx,hello)
//			if err != nil {
//				if err == redislock.ErrNotObtained {
//					time.Sleep(time.Millisecond*100)
//					continue
//				}
//				return nil, err
//			}
//			return lock,nil
//		}
//	}
//}
