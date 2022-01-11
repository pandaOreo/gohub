/**
 * @Author: fanjinghua
 * @Description:
 * @File: custom_validators
 * @Version: 1.0.0
 * @Date: 2022/1/11 22:09
 */

// Package validators 存放自定义规则和验证器
package validators

import "github.com/ZimoBoy/gohub/pkg/captcha"

func ValidateCaptcha(captchaID, captchaAnswer string, errs map[string][]string) map[string][]string {
	if ok := captcha.NewCaptcha().VerifyCaptcha(captchaID, captchaAnswer); !ok {
		errs["captcha_answer"] = append(errs["captcha_answer"], "图片验证码错误")
	}
	return errs
}
