package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type UserTable_20181202_133212 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &UserTable_20181202_133212{}
	m.Created = "20181202_133212"
	_ = migration.Register("UserTable_20181202_133212", m)
}

// Run the migrations
func (m *UserTable_20181202_133212) Up() {
	m.SQL("CREATE TABLE public.user(id BIGSERIAL PRIMARY KEY, login VARCHAR(255) UNIQUE NOT NULL, full_name VARCHAR(255), password VARCHAR(511), salt VARCHAR(511) )")

}

// Reverse the migrations
func (m *UserTable_20181202_133212) Down() {
	m.SQL("DROP TABLE user")

}
