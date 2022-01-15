/**
 * @Author: fanjinghua
 * @Description:
 * @File: user_factory.go
 * @Version: 1.0.0
 * @Date: 2022/1/15 13:11
 */

// Package factories 存放工厂方法
package factories

import (
	"github.com/ZimoBoy/gohub/app/models/user"
	"github.com/ZimoBoy/gohub/pkg/helpers"
	"github.com/bxcodec/faker/v3"
)

func MakeUsers(times int) []user.User {
	var objs []user.User

	// 设置唯一值
	faker.SetGenerateUniqueValues(true)

	for i := 0; i < times; i++ {
		model := user.User{
			Name:     faker.Username(),
			Email:    faker.Email(),
			Phone:    helpers.RandomNumber(11),
			Password: "$2a$14$oPzVkIdwJ8KqY0erYAYQxOuAAlbI/sFIsH0C0R4MPc.3JbWWSuaUe",
		}
		objs = append(objs, model)
	}
	return objs
}
