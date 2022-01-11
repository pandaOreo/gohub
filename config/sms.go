/**
 * @Author: fanjinghua
 * @Description:
 * @File: sms
 * @Version: 1.0.0
 * @Date: 2022/1/11 17:59
 */

// Package config 站点配置信息
package config

import "github.com/ZimoBoy/gohub/pkg/config"

func init() {
	config.Add("sms", func() map[string]interface{} {
		return map[string]interface{}{
			// 默认是阿里云的测试 sign_name 和 template_code
			"aliyun": map[string]interface{}{
				"access_key_id":     config.Env("SMS_ALIYUN_ACCESS_ID"),
				"access_key_secret": config.Env("SMS_ALIYUN_ACCESS_SECRET"),
				"sign_name":         config.Env("SMS_ALIYUN_ACCESS_NAME", "阿里云短信测试"),
				"template_code":     config.Env("SMS_ALIYUN_TEMPLATE_CODE", "SMS_154950909"),
			},
		}
	})
}
