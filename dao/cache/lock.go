package cache

type Locker struct {
	cacheTime int
	key       string
}

type AllLocker struct {
	SignLocker *Locker
	ChooseLock *Locker
}

const (
	SignLock   = "sign_lock"
	ChooseLock = "choose_lock"
)

func newAllLock() *AllLocker {
	return &AllLocker{
		SignLocker: newSignLock(),
		ChooseLock: newChooseLock(),
	}
}

func newSignLock() *Locker {
	return &Locker{
		cacheTime: 10,
		key:       SignLock,
	}
}

func newChooseLock() *Locker {
	return &Locker{
		cacheTime: 10,
		key:       ChooseLock,
	}
}

func (l *Locker) Lock(key int64) bool {
	rds := CachePool.Get()
	defer rds.Close()
	reply, _ := rds.Do("EVAL", `if redis.call("setnx",KEYS[1],ARGV[1])==1 then return redis.call("expire",KEYS[1],ARGV[2]) end return 0`, 1, l.key, key, l.cacheTime)
	if reply == 1 {
		return true
	}
	return false
}

func (l *Locker) UnLock(key int64) bool {
	rds := CachePool.Get()
	defer rds.Close()
	reply, _ := rds.Do("EVAL", `if redis.call("get",KEYS[1])==ARGV[1] then return redis.call("del",KEYS[1]) end return 0`, 1, l.key, key)
	if reply == 1 {
		return true
	}
	return false
}
