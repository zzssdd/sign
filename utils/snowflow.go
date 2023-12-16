package utils

import (
	"sign/conf"
	"sync"
	"time"
)

type SnowFlow struct {
	mutex         sync.Mutex
	workId        int64
	sequence      int64
	timeStamp     int64
	beginStamp    int64
	timeStampBits int
	workIdBits    int
	sequenceBits  int
}

func NewSnowFlow(conf *conf.SnowFlow) *SnowFlow {
	return &SnowFlow{
		mutex:         sync.Mutex{},
		workId:        conf.WorkId,
		beginStamp:    conf.BeginStamp,
		timeStampBits: conf.TimeStampBits,
		workIdBits:    conf.WorkIdBits,
		sequenceBits:  conf.SequenceBits,
	}
}

func (s *SnowFlow) GenID() int64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	now := time.Now().UnixNano() / 1000000
	if now == s.timeStamp {
		s.sequenceBits++
	} else {
		s.timeStamp = now
		s.sequenceBits = 0
	}
	t := now - s.beginStamp
	return ((t & (1<<s.timeStampBits - 1)) << int64(s.workIdBits+s.sequenceBits)) + ((s.workId & (1<<s.workIdBits - 1)) << int64(s.sequenceBits)) + (s.sequence & (1<<s.sequenceBits - 1))
}
