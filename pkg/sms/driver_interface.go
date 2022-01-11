/**
 * @Author: fanjinghua
 * @Description:
 * @File: driver_interface
 * @Version: 1.0.0
 * @Date: 2022/1/11 17:32
 */

package sms

type Driver interface {
	// 发送短信
	Send(phone string, message Message, config map[string]string) bool
}
