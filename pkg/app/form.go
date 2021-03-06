/*
 *    Copyright 2020 opensourceai
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"strconv"

	"github.com/opensourceai/go-api-service/pkg/e"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (httpCode, errCode int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, e.ERROR
	}
	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	return http.StatusOK, e.SUCCESS
}

// 使用 validates 数据校验
func Valid(form interface{}) (int, int) {
	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, e.ERROR
	}
	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, e.INVALID_PARAMS
	}
	return http.StatusOK, e.SUCCESS
}

// 认证信息
type Auth struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
}

// 获取gin上下文中的用户信息
func GetUserInfo(content *gin.Context) *Auth {
	var userId interface{}
	var username interface{}
	var exists bool
	if userId, exists = content.Get("userId"); !exists {
		panic("认证失败")
	}

	if username, exists = content.Get("username"); !exists {
		panic("认证失败")
	}

	return &Auth{
		UserId:   userId.(int),
		Username: com.ToStr(username),
	}

}

// 获取int类型查询参数
func QueryWithInt(ctx *gin.Context, key string) (int, error) {
	param := ctx.Query(key)
	if param == "" {
		return 0, nil
	}
	atoi, err := strconv.Atoi(param)
	if err != nil {
		return 0, err
	}
	return atoi, nil

}
