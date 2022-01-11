/**
 * @Author: fanjinghua
 * @Description:
 * @File: driver_interface
 * @Version: 1.0.0
 * @Date: 2022/1/11 21:39
 */

package mail

type Driver interface {
	// Send 检查验证码
	Send(email Email, config map[string]string) bool
}
