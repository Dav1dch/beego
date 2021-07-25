/*
 * @Author: Junlang
 * @Date: 2021-07-17 16:50:16
 * @LastEditTime: 2021-07-17 23:41:58
 * @LastEditors: Junlang
 * @FilePath: /OpenCTScore/main.go
 */
package main

import (
	_ "OpenCTScore/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}
