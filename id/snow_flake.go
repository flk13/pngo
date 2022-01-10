package id

import (
	"fmt"
	"sync"
	"time"
)

const (
	MAX_WORK         = 10                    //节点数
	MAX_NUMBER       = 12                    //1ms内可生成的id序号的二进制位数
	MAX_WORK_VALUE   = 0x3ff                 //节点id的最大值，防止溢出
	MAX_NUMBER_VALUE = 0xfff                 //id序号的最大值，防止溢出
	OFFSET_TIME      = MAX_NUMBER + MAX_WORK //时间戳向左的偏移量
	OFFSET_WORK      = MAX_NUMBER            //节点id向左的偏移量
)

var (
	oldNow int64 = 1641799227
)

type SnowFlake struct {
	//添加一个互斥锁保证并发安全
	sync.Mutex
	//记录时间戳
	timeStamp int64
	//该节点的id
	workNode int64
	//当前毫秒已经生成的id序列号，从0开始累加，1ms内最多最多4096个id
	number int64
}

//实例化对象
func NewSnowFlake(workNode int64) (*SnowFlake, error) {
	//先检测当前节点id是否在上面定义的范围内
	if workNode > MAX_WORK_VALUE || workNode < 0 {
		return nil, fmt.Errorf("input number is tto big or <0, right number is [0, 1023]")
	}
	return &SnowFlake{workNode: workNode}, nil
}

//获取一个新id
func (s *SnowFlake) GetID() int64 {
	//获取id的关键一步，上锁，最后记得解锁
	s.Lock()
	defer s.Unlock()
	//获取现在时间的时间戳
	now := time.Now().UnixNano() / 1e6 //纳秒转换成毫秒

	//判断当前工作节点是否在1ms内已经生成了MAX_WORK_VALUE个id
	if now == s.timeStamp {
		s.number++
		//如果溢出超过上限，则等待1ms
		if s.number > MAX_NUMBER_VALUE {
			for now <= s.timeStamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		//如果当前时间与工作节点上一次生成ID的时间不一致 则需要重置工作节点生成ID的序号
		s.number = 0
		s.timeStamp = now // 将机器上一次生成ID的时间更新为当前时间
	}
	return (s.timeStamp-oldNow)<<OFFSET_TIME | s.workNode<<OFFSET_WORK | s.number
}
