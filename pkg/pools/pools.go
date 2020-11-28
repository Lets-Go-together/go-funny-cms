package pools

import (
	"github.com/panjf2000/ants/v2"
	"gocms/pkg/config"
	"gocms/pkg/logger"
	"strconv"
	"sync"
)

func init() {
	var options = ants.Options{}
	maxTalks := config.GetInt("POOL_MAX_TASKS", 10)
	preAlloc := config.GetBool("POOL_PRE_All_OC", true)
	options.PreAlloc = preAlloc
	options.MaxBlockingTasks = maxTalks

	// ---
	pool, e := ants.NewPool(maxTalks, ants.WithOptions(options))
	logger.PanicError(e, "init goroutine pool", true)

	config.Pool = pool
}

func Initialize() {}

func PoolsExample(i int, wg *sync.WaitGroup) {
	_ = ants.Submit(func() {
		logger.Info("pool test", strconv.Itoa(i))
		wg.Done()
	})
}
