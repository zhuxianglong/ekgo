package service

import (
	"ekgo/app/model"
	"ekgo/app/repository"
	"github.com/small-ek/antgo/crypto/hash"
	"github.com/small-ek/antgo/encoding/base64"
	"github.com/small-ek/antgo/jwt"
	"github.com/small-ek/antgo/os/config"
	"github.com/small-ek/antgo/request"
	"github.com/small-ek/antgo/response"
	"time"
)

type Admin struct {
	PageParam request.PageParam
	Model     model.Admin
}

//Index 分页
func (s *Admin) Index() (*response.Page, error) {
	var list, total, err = repository.Admin.Default(s.Model).GetPage(s.PageParam)
	return &response.Page{List: list, Total: total}, err
}

//Show 查询
func (s *Admin) Show() (interface{}, error) {
	var result, err = repository.Admin.Default(s.Model).FindByID()
	return result, err
}

//Store 添加
func (s *Admin) Store() (interface{}, error) {
	var result, err = repository.Admin.Default(s.Model).Create()
	result.Password = ""
	result.Salt = ""
	return result, err
}

//Update 修改
func (s *Admin) Update() error {
	var err = repository.Admin.Default(s.Model).Update()
	return err
}

//Delete 删除
func (s *Admin) Delete() error {
	var err = repository.Admin.Default(s.Model).Delete()
	return err
}

//Login 登录
func (s *Admin) Login() (map[string]interface{}, error) {
	var result, err = repository.Admin.Default(s.Model).FindByUserName()
	if result.Id > 0 {
		var ClearTextPassword, _ = base64.Decode(s.Model.Password)
		var password = hash.Sha256(result.Salt + ClearTextPassword)

		if password == result.Password {
			var private_key = config.Decode().Get("jwt").Get("private_key").Bytes()
			var public_key = config.Decode().Get("jwt").Get("public_key").Bytes()

			var authorization, err = jwt.Default(public_key, private_key).Encrypt(map[string]interface{}{
				"id":       result.Id,
				"username": result.Username,
				"exp":      time.Now().Add(time.Hour * 168).Unix(),
			})
			return map[string]interface{}{
				"authorization":   authorization,
				"id":              result.Id,
				"username":        result.Username,
				"name":            result.Name,
				"super":           result.Super,
				"role_id":         result.RoleId,
				"profile_picture": result.ProfilePicture,
			}, err
		}
	}
	return nil, err
}
