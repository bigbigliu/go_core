package pkgs

import (
	"encoding/binary"
	"github.com/satori/go.uuid"
	"strconv"
	"sync"
	"time"
)

var (
	lastTimestamp uint64
	mutex         sync.Mutex
)

// GenUniversalId 生成全局唯一 ID
func GenUniversalId() string {
	mutex.Lock()
	defer mutex.Unlock()

	// 获取当前时间戳
	currentTimestamp := uint64(time.Now().UnixNano() / 1000000) // 毫秒级时间戳

	// 如果当前时间戳与上一个相同，则递增序列号
	if currentTimestamp <= lastTimestamp {
		lastTimestamp++
	} else {
		lastTimestamp = currentTimestamp
	}

	// 构建 UUID，包括时间戳和序列号
	u4 := uuid.NewV4()
	uuidBytes := u4.Bytes()
	binary.BigEndian.PutUint64(uuidBytes[0:8], lastTimestamp)

	// 将 UUIDBytes 转换为 UUID 类型
	uuidFromBytes, _ := uuid.FromBytes(uuidBytes)

	// 生成 UUID 版本 5
	u5 := uuid.NewV5(uuidFromBytes, "51app")
	i64 := binary.BigEndian.Uint64(u5.Bytes())

	return strconv.FormatUint(i64, 10)
}
