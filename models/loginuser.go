package models

type LoginUser struct {
	Uid      int64
	Username string
	Now      int64
	Ip       string
	Sign     string
}
