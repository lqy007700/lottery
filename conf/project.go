package conf

import "time"

const SysTimeForm = "2006-01-02 15:04:05"
const SysTimeFormShort = "2006-01-02"

var SysTimeLocation, _ = time.LoadLocation("Local")

var SignSecret = []byte("0123qwe23")

var CookieSecret = "helloliu"
