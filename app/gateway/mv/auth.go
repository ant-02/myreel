package mv

import (
	"context"
	"myreel/app/gateway/pack"
	"myreel/pkg/constants"
	"myreel/pkg/util"

	"github.com/cloudwego/hertz/pkg/app"
)

func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		token := string(c.GetHeader(constants.AuthHeader))
		_, uid, err := util.CheckToken(token)
		if err != nil {
			pack.RespError(c, err)
			c.Abort()
			return
		}

		c.Set(constants.CtxUserIdKey, uid)
		c.Next(ctx)
	}
}
