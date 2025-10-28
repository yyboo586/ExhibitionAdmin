/*
* @desc:token功能
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/9/27 17:01
 */

package token

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/tiger1103/gfast/v3/internal/app/common/consts"
	commonModel "github.com/tiger1103/gfast/v3/internal/app/common/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/liberr"
	"github.com/yyboo586/common/authUtils/tokenUtils"
)

type sToken struct {
	*tokenUtils.Token
}

func New() service.IGfToken {
	var (
		ctx = gctx.New()
		opt *commonModel.TokenOptions
		err = g.Cfg().MustGet(ctx, "gfToken").Struct(&opt)
		fun tokenUtils.OptionFunc
	)
	liberr.ErrIsNil(ctx, err)

	switch opt.CacheModel {
	case consts.CacheModelRedis:
		fun = tokenUtils.WithGRedis() //redis缓存
	case consts.CacheModelMem:
		fun = tokenUtils.WithGCache() // 内存缓存
	default:
		panic("invalid cache model, only support redis and memory")
	}
	return &sToken{
		Token: tokenUtils.NewToken(
			tokenUtils.WithCacheKey(opt.CacheKey),
			tokenUtils.WithTimeout(opt.Timeout),
			tokenUtils.WithMaxRefresh(opt.MaxRefresh),
			tokenUtils.WithMultiLogin(opt.MultiLogin),
			tokenUtils.WithExcludePaths(opt.ExcludePaths),
			fun,
		),
	}
}

func init() {
	service.RegisterGToken(New())
}
