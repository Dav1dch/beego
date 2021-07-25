/*
 * @Author: Junlang
 * @Date: 2021-07-17 16:50:16
 * @LastEditTime: 2021-07-17 21:57:25
 * @LastEditors: Junlang
 * @FilePath: /OpenCTScore/routers/router.go
 */
package routers

import (
	"OpenCTScore/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/?:id", &controllers.UserController{})
	beego.Router("/user", &controllers.UserController{})
}
