/*
 * @Author: Junlang
 * @Date: 2021-07-17 21:12:19
 * @LastEditTime: 2021-07-24 15:50:57
 * @LastEditors: Junlang
 * @FilePath: /openscore-beego/models/user.go
 */
package models

import (
	"errors"
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/builder"
	"xorm.io/xorm"
)

var x *xorm.Engine

type User struct {
	Id   int64  `xorm:"id pk autoincr"`
	Name string `xorm:"varchar(50)"`
}

func init() {
	var err error
	user, _ := beego.AppConfig.String("mysql::MYSQL_USER")
	password, _ := beego.AppConfig.String("mysql::MYSQL_PASSWORD")
	host, _ := beego.AppConfig.String("mysql::MYSQL_HOST")
	port, _ := beego.AppConfig.String("mysql::MYSQL_PORT")
	database, _ := beego.AppConfig.String("mysql::MYSQL_DATABASE")
	mysqlDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, password, host, port, database)
	fmt.Println(mysqlDSN)
	x, err = xorm.NewEngine("mysql", mysqlDSN)
	if err != nil {
		fmt.Println("Fail to create xorm engine")
		panic(err)
	}
	x.Sync2(new(User))
	x.ShowSQL(true)
}

func (u *User) GetOneUser(id int64) error {
	has, err := x.Where(builder.Eq{"id": id}).Get(u)
	if !has || err != nil {
		fmt.Println("could not find user")
	}
	return err
}

func (u *User) Create() error {
	_, err := x.Insert(u)
	return err
}

func (u *User) UpdateUser(id int64, v *User) error {
	affec, err := x.Where(builder.Eq{"id": id}).Update(v)
	if err == nil && affec == 0 {
		err = errors.New("update user error")
	}
	return err
}

func (u *User) DeleteUser(id int64) error {
	affec, err := x.Where(builder.Eq{"id": id}).Delete(u)
	if affec == 0 && err == nil {
		err = errors.New("delete user error")
	}
	return err
}
