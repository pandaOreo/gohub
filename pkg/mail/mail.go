/**
 * @Author: fanjinghua
 * @Description:
 * @File: mail
 * @Version: 1.0.0
 * @Date: 2022/1/11 21:50
 */

// Package mail 发送短信
package mail

import (
	"github.com/ZimoBoy/gohub/pkg/config"
	"sync"
)

type From struct {
	Address string
	Name    string
}

type Email struct {
	From    From
	To      []string
	Bcc     []string
	Cc      []string
	Subject string
	Text    []byte
	HTML    []byte
}

type Mailer struct {
	Driver Driver
}

var once sync.Once
var interfaceMailer *Mailer

// NewMailer 单例模式
func NewMailer() *Mailer {
	once.Do(func() {
		interfaceMailer = &Mailer{
			Driver: &SMTP{},
		}
	})
	return interfaceMailer
}

func (m *Mailer) Send(email Email) bool {
	return m.Driver.Send(email, config.GetStringMapString("mail.smtp"))
}
