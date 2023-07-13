// Package policies 用户授权
package policies

import (
	"github.com/gin-gonic/gin"
	"github.com/lihc520/gohub/app/models/topic"
	"github.com/lihc520/gohub/pkg/auth"
)

func CanModifyTopic(c *gin.Context, _topic topic.Topic) bool {
	return auth.CurrentUID(c) == _topic.UserID
}
