/**
 * @Author: fanjinghua
 * @Description:
 * @File: login_controller
 * @Version: 1.0.0
 * @Date: 2022/1/12 13:59
 */

package auth

import (
	v1 "github.com/ZimoBoy/gohub/app/http/controllers/api/v1"
	"github.com/ZimoBoy/gohub/app/requests"
	"github.com/ZimoBoy/gohub/pkg/auth"
	"github.com/ZimoBoy/gohub/pkg/jwt"
	"github.com/ZimoBoy/gohub/pkg/response"
	"github.com/gin-gonic/gin"
)

// LoginController 用户登录控制器
type LoginController struct {
	v1.BaseApiController
}

// LoginByPhone 手机登录
func (lc *LoginController) LoginByPhone(c *gin.Context) {
	// 1. 验证表单
	request := requests.LoginByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPhone); !ok {
		return
	}

	// 2. 尝试登录
	user, err := auth.LoginByPhone(request.Phone)
	if err != nil {
		// 失败,显示错误提示
		response.Error(c, err, "账号不存在或密码错误")
	} else {
		// 登录成功
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)

		response.JSON(c, gin.H{
			"token": token,
		})
	}
}
