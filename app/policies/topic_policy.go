/**
 * @Author: fanjinghua
 * @Description:
 * @File: topic_policy
 * @Version: 1.0.0
 * @Date: 2022/1/16 10:20
 */

// Package policies 用户授权
package policies

import (
	"github.com/ZimoBoy/gohub/app/models/topic"
	"github.com/ZimoBoy/gohub/pkg/auth"
	"github.com/gin-gonic/gin"
)

func CanModifyTopic(c *gin.Context, _topic topic.Topic) bool {
	return auth.CurrentUID(c) == _topic.UserID
}
