package context

import (
	"github.com/xxiiaaon/wechat/v2/credential"
	"github.com/xxiiaaon/wechat/v2/miniprogram/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
