package middleware

/*
	应用的调用链如果是应用 A =》应用 B =》应用 C, 如果应用 C 出现了问题,
	在没有任何约束的情况下持续调用, 就会导致应用 A、B、C 均出现问题,
	也就是很常见的上下游应用的互相影响, 导致连环反应, 最终使得整个集群应
	用出现一定规模的不可用;

	在进行多应用/服务的调用时, 把父级的上下文信息(ctx)不断地传递下去,
	那么在统计超时控制的中间件中所设置的超时时间, 其实是针对整条链路的,
	而不是针对单单每一条, 如果你需要针对额外的链路进行超时时间的调整,
	那么只需要调用像 context.WithTimeout 等方法对父级 ctx 进行设置,
	然后取得子级 ctx, 再进行新的上下文传递就可以

*/

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

func ContextTimeout(t time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)

		/*
			调用了 context.WithTimeout 方法设置当前 context 的超时时间,
			并重新赋予给了 gin.Context, 在当前请求运行到指定的时间后,
			在使用了该 context 的运行流程就会针对 context 所提供的超
			时时间进行处理, 并在指定的时间进行取消行为
		*/
		c.Next()
	}
}
