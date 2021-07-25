/*
 * @Author: Junlang
 * @Date: 2021-07-17 16:50:16
 * @LastEditTime: 2021-07-17 23:47:03
 * @LastEditors: Junlang
 * @FilePath: /OpenCTScore/controllers/default.go
 */
package controllers

import (
	"OpenCTScore/models"
	"encoding/json"
	"fmt"
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

type UserController struct {
	beego.Controller
}

/**
 * @description: RestfulAPI Get method * @param {void}
 * @return {json}
 */
func (c *MainController) Get() {
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	// c.TplName = "index.tpl"
	m := make(map[string]string)
	m["name"] = "hello world"
	c.Data["json"] = m
	c.ServeJSON()
}

/**
 * @description: User RestfulAPI Get method
 * @param {void}
 * @return {user json}
 */
func (userController *UserController) Get() {
	id := userController.GetString(":id")
	id64, _ := strconv.ParseInt(id, 10, 64)
	u := new(models.User)
	u.GetOneUser(id64)
	userController.Data["json"] = u
	userController.ServeJSON()
	userController.StopRun()
}

/**
 * @description: User RestfulAPI Post method for create user
 * @param {user}
 * @return {*}
 */
func (userController *UserController) Post() {
	u := new(models.User)
	var err error
	if err = json.Unmarshal(userController.Ctx.Input.RequestBody, u); err == nil {
		fmt.Printf("%v", u)
		err1 := u.Create()
		if err1 != nil {
			fmt.Printf("create fail")
		}
	} else {
		fmt.Println("unmarshal json fail")
		userController.Data["json"] = err.Error()
	}
	userController.ServeJSON()

}

/**
 * @description: User RestfulAPI Put method for update user
 * @param {id, user}
 * @return {*}
 */
func (userController *UserController) Put() {
	u := new(models.User)
	var err error
	if err = json.Unmarshal(userController.Ctx.Input.RequestBody, u); err == nil {
		fmt.Printf("%v", u)
		err1 := u.UpdateUser(u.Id, u)
		if err1 != nil {
			fmt.Printf("update fail")
		}
	} else {
		userController.Data["json"] = err.Error()
		fmt.Println("unmarshal json fail")
	}
	userController.ServeJSON()

}

/**
 * @description: User RestfulAPI Delete method for delete user
 * @param {*}
 * @return {*}
 */
func (userController *UserController) Delete() {
	u := new(models.User)
	var err error
	if err = json.Unmarshal(userController.Ctx.Input.RequestBody, u); err == nil {
		fmt.Printf("%v", u)
		err1 := u.DeleteUser(u.Id)
		if err1 != nil {
			fmt.Printf("delete fail")
		}
	} else {
		userController.Data["json"] = err.Error()
		fmt.Println("unmarshal json fail")
	}
	userController.ServeJSON()
}
