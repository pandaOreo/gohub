/**
 * @Author: fanjinghua
 * @Description:
 * @File: sms
 * @Version: 1.0.0
 * @Date: 2022/1/11 17:52
 */

// Package sms 发送短信
package sms

import (
	"github.com/ZimoBoy/gohub/pkg/config"
	"sync"
)

// Message 短信的结构体
type Message struct {
	Template string
	Data     map[string]string
	Context  string
}

// SMS 发送短信的操作类
type SMS struct {
	Driver Driver
}

// once 单例模式
var once sync.Once

// internalSMS SMS对象
var internalSMS *SMS

// NewSMS 单例模式获取
func NewSMS() *SMS {
	once.Do(func() {
		internalSMS = &SMS{
			Driver: &Aliyun{},
		}
	})
	return internalSMS
}

// Send 发送短信
func (sms *SMS) Send(phone string, message Message) bool {
	return sms.Driver.Send(phone, message, config.GetStringMapString("sms.aliyun"))
}
