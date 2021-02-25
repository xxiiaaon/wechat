package context

import (
	"github.com/xxiiaaon/wechat/v2/credential"
	"github.com/xxiiaaon/wechat/v2/officialaccount/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
