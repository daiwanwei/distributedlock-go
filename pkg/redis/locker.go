package redis

import (
	"github.com/bsm/redislock"
)

var (
	lockerInstance *redislock.Client
)

func GetLocker() (instance *redislock.Client, err error) {
	if lockerInstance == nil {
		instance, err = NewLocker()
		if err != nil {
			return nil, err
		}
		lockerInstance = instance
	}
	return lockerInstance, nil
}

func NewLocker() (instance *redislock.Client, err error) {
	r, err := GetRedis()
	if err != nil {
		return nil, err
	}
	return redislock.New(r), nil
}
