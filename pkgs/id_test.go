package pkgs

import (
	"sync"
	"testing"
)

func TestGenUniversalIdConcurrent(t *testing.T) {
	// 设置并发测试的协程数
	numGoroutines := 10000

	// 使用 WaitGroup 等待所有协程完成
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// 用于存储生成的 UUID
	ids := make(map[string]struct{})
	var mutex sync.Mutex // 用于保护 ids 的互斥锁

	// 启动并发协程生成 UUID
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			uuid := GenUniversalId()

			// 通过互斥锁保护对 ids 的并发访问
			mutex.Lock()
			defer mutex.Unlock()

			// 判断是否有重复的 UUID
			if _, exists := ids[uuid]; exists {
				t.Errorf("发现重复的 UUID: %s", uuid)
			}

			// 打印生成的 UUID
			//t.Log("uuid: ", uuid)

			// 将 UUID 加入到 ids 中
			ids[uuid] = struct{}{}
		}()
	}

	// 等待所有协程完成
	wg.Wait()
}
