package cache

import (
	"fmt"
	"github.com/bluele/gcache"
)

/**
前提请先了解 LRU LFU的 使用
参考文件:
https://www.cnblogs.com/sddai/p/9739900.html
LRU，即：最近最少使用淘汰算法（Least Recently Used）。LRU是淘汰最长时间没有被使用的页面。
LFU，即：最不经常使用淘汰算法（Least Frequently Used）。LFU是淘汰一段时间内，使用次数最少的页面。
ARC，即：在LRU 和 ARC 中间不断平衡，取得最佳的结果。
*/

func init() {
	gc := gcache.New(20).
		LRU().
		Build()
	gc.Set("key", "ok")
	value, err := gc.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println("Get:", value)
}
