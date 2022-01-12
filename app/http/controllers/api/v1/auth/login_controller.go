/**
 * @Author       : Jinghua Fan
 * @Date         : 2022-01-12 13:59:45
 * @LastEditors  : Jinghua Fan
 * @LastEditTime : 2022-01-12 20:02:29
 * @Description  : 佛祖保佑,永无BUG
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

func (lc *LoginController) LoginByPassword(c *gin.Context) {
	// 1. 验证表单
	request := requests.LoginByPasswordRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPassword); !ok {
		return
	}
	// 2. 尝试登录
	user, err := auth.Attempt(request.LoginID, request.Password)
	if err != nil {
		// 失败,显示错误提示
		response.Unauthorized(c, "登录失败")
	} else {
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)
		response.JSON(c, gin.H{
			"token": token,
		})
	}
}

// RefreshToken 刷新 Access Token
func (lc *LoginController) RefreshToken(c *gin.Context) {
	token, err := jwt.NewJWT().RefreshToken(c)

	if err != nil {
		response.Error(c, err, "令牌刷新失败")
	} else {
		response.JSON(c, gin.H{
			"token": token,
		})
	}
}
