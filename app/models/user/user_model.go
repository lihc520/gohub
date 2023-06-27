// Package user 存放用户 Model 相关逻辑
package user

import "github.com/lihc520/gohub/app/models"

// User 用户类型
type User struct {
	models.BaseModel

	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}
