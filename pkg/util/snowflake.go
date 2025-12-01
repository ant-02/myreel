package util

import (
	"fmt"
	"sync"
	"time"
)

// 定义常量
const (
	workerIDBits     = 5  // 机器ID位数
	datacenterIDBits = 5  // 数据中心ID位数
	sequenceBits     = 12 // 序列号位数

	maxWorkerID     = -1 ^ (-1 << workerIDBits)     // 最大机器ID: 31
	maxDatacenterID = -1 ^ (-1 << datacenterIDBits) // 最大数据中心ID: 31
	maxSequence     = -1 ^ (-1 << sequenceBits)     // 最大序列号: 4095

	workerIDShift     = sequenceBits                                   // 机器ID左移位数: 12
	datacenterIDShift = sequenceBits + workerIDBits                    // 数据中心ID左移位数: 17
	timestampShift    = sequenceBits + workerIDBits + datacenterIDBits // 时间戳左移位数: 22
)

// Snowflake 雪花算法结构体
type Snowflake struct {
	mu            sync.Mutex
	startTime     int64 // 起始时间戳 (毫秒)
	workerID      int64 // 机器ID
	datacenterID  int64 // 数据中心ID
	sequence      int64 // 序列号
	lastTimestamp int64 // 上次生成ID的时间戳
}

// ParsedID 解析后的ID信息
type ParsedID struct {
	Timestamp    int64 // 时间戳
	DatacenterID int64 // 数据中心ID
	WorkerID     int64 // 机器ID
	Sequence     int64 // 序列号
}

// NewSnowflake 创建雪花算法实例
func NewSnowflake(workerID, datacenterID int64) (*Snowflake, error) {
	// 参数检查
	if workerID < 0 || workerID > maxWorkerID {
		return nil, fmt.Errorf("workerid must be between 0 and %d", maxWorkerID-1)
	}
	if datacenterID < 0 || datacenterID > maxDatacenterID {
		return nil, fmt.Errorf("datacenterid must be between 0 and %d", maxDatacenterID-1)
	}

	// 使用 2020-01-01 00:00:00 UTC 作为起始时间
	startTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli()

	return &Snowflake{
		startTime:    startTime,
		workerID:     workerID,
		datacenterID: datacenterID,
	}, nil
}

// Generate 生成唯一ID
func (s *Snowflake) Generate() (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	timestamp := time.Now().UnixMilli()

	// 检查时钟回拨
	if timestamp < s.lastTimestamp {
		return 0, fmt.Errorf("clock moved backwards, refusing to generate ID. Last timestamp: %d, current timestamp: %d",
			s.lastTimestamp, timestamp)
	}

	// 同一毫秒内生成
	if timestamp == s.lastTimestamp {
		s.sequence = (s.sequence + 1) & maxSequence
		// 序列号用完，等待下一毫秒
		if s.sequence == 0 {
			timestamp = s.waitNextMillis(timestamp)
		}
	} else {
		// 新的毫秒，重置序列号
		s.sequence = 0
	}

	s.lastTimestamp = timestamp

	// 组合生成ID
	id := ((timestamp - s.startTime) << timestampShift) |
		(s.datacenterID << datacenterIDShift) |
		(s.workerID << workerIDShift) |
		s.sequence

	return id, nil
}

// waitNextMillis 等待下一毫秒
func (s *Snowflake) waitNextMillis(lastTimestamp int64) int64 {
	timestamp := time.Now().UnixMilli()
	for timestamp <= lastTimestamp {
		time.Sleep(100 * time.Microsecond) // 休眠100微秒
		timestamp = time.Now().UnixMilli()
	}
	return timestamp
}

// ParseID 解析雪花ID
func (s *Snowflake) ParseID(id int64) *ParsedID {
	timestamp := (id >> timestampShift) + s.startTime
	datacenterID := (id >> datacenterIDShift) & maxDatacenterID
	workerID := (id >> workerIDShift) & maxWorkerID
	sequence := id & maxSequence

	return &ParsedID{
		Timestamp:    timestamp,
		DatacenterID: datacenterID,
		WorkerID:     workerID,
		Sequence:     sequence,
	}
}
