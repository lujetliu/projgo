package limiter

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit" // TODO: 熟悉限流器原理, ratelimit 包的使用
)

type LimiterIface interface {
	Key(c *gin.Context) string
	GetBucket(key string) (*ratelimit.Bucket, bool)
	AddBuckets(rules ...LimiterBucketRule) LimiterIface // 函子, 函数式编程

	/*
		限流器是存在多种实现的, 可能某一类接口需要限流器 A, 另外一类接口需要
		限流器 B, 所采用的策略不是完全一致的, 因此需要声明 LimiterIfac 这类
		通用接口, 保证其接口的设计, 初步的在 Iface 接口中, 一共声明了三个方法:
		Key: 获取对应的限流器的键值对名称
		GetBucket: 获取令牌桶
		AddBuckets: 新增多个令牌桶
	*/
}

// TODDO: 为何不直接用map, 确保包外不能访问此map?
type Limiter struct {
	// 存储令牌桶与键值对名称的映射关系
	limiterBuckets map[string]*ratelimit.Bucket
}

type LimiterBucketRule struct {
	Key          string        // 自定义键值对名称
	FillInterval time.Duration // 间隔多久时间放 Quantum 个令牌
	Capacity     int64         // 令牌桶的容量
	Quantum      int64         // 每次到达间隔时间后所放的具体令牌数量
}
