/**
 * @Author: fanjinghua
 * @Description:
 * @File: verifycode
 * @Version: 1.0.0
 * @Date: 2022/1/11 20:37
 */

package config

import "github.com/ZimoBoy/gohub/pkg/config"

func init() {
	config.Add("verifycode", func() map[string]interface{} {
		return map[string]interface{}{
			// 验证码的长度
			"code_length": config.Env("VERIFY_CODE_LENGTH", 6),
			// 过期时间
			"expire_time": config.Env("VERIFY_EXPIRE_TIME", 15),
			// debug 模式下的过期时间, 方便本地开发调试
			"debug_expire_time": 10080,
			// 本地开发环境验证码使用 debug_code
			"debug_code": 123456,
			// 方便本地和 API 自动测试
			"debug_phone_prefix": "000",
			"debug_email_suffix": "@testing.com",
		}
	})
}
