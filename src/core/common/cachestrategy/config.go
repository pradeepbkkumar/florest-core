package cachestrategy

import (
	"github.com/jabong/florest-core/src/components/cache"
	"github.com/jabong/florest-core/src/core/common/threadpool"
)

type Config struct {
	Strategy      string
	DBAdapterType string
	Cache         cache.Config
	ThreadPool    threadpool.Config
}
