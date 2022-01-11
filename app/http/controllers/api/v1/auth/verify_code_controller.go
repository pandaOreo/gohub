/**
 * @Author: fanjinghua
 * @Description:
 * @File: verify_code_controller
 * @Version: 1.0.0
 * @Date: 2022/1/11 08:39
 */

package auth

import (
	v1 "github.com/ZimoBoy/gohub/app/http/controllers/api/v1"
	"github.com/ZimoBoy/gohub/pkg/captcha"
	"github.com/ZimoBoy/gohub/pkg/logger"
	"github.com/ZimoBoy/gohub/pkg/response"
	"github.com/gin-gonic/gin"
)

// VerifyCodeController 用户控制器
type VerifyCodeController struct {
	v1.BaseApiController
}

// ShowCaptcha 显示图片验证码
func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	// 生成验证码
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	// 记录错误日志,因为验证码是用户的入口,出错时应该记 error 等级的日志
	logger.LogIf(err)
	// 返回给用户
	response.JSON(c, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}
