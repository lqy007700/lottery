package common

import (
	"crypto/md5"
	"fmt"
	"lottery/conf"
	"lottery/models"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func ClientIp(req *http.Request) string {
	ip, _, _ := net.SplitHostPort(req.RemoteAddr)
	return ip
}

func Redirect(writer http.ResponseWriter, url string) {
	writer.Header().Add(conf.CookieKey, url)
	writer.WriteHeader(http.StatusFound)
}

func GetLoginUser(req *http.Request) *models.LoginUser {
	cookie, err := req.Cookie("lottery_user")
	if err != nil {
		return nil
	}

	query, err := url.ParseQuery(cookie.Value)
	if err != nil {
		return nil
	}

	uid, err := strconv.Atoi(query.Get("uid"))
	if err != nil {
		return nil
	}

	atoi, err := strconv.Atoi(query.Get("now"))
	if err != nil || NowUnix()-atoi > 86400*30 {
		return nil
	}

	user := &models.LoginUser{
		Uid:      int64(uid),
		Username: query.Get("username"),
		Now:      time.Now().Unix(),
		Ip:       ClientIp(req),
		Sign:     query.Get("sign"),
	}

	if user.Sign != createSign(user) {
		return nil
	}

	return user
}

func SetLoginUser(writer http.ResponseWriter, user *models.LoginUser) {
	if user == nil || user.Uid < 1 {
		c := &http.Cookie{
			Name:   conf.CookieKey,
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		}
		http.SetCookie(writer, c)
		return
	}

	if user.Sign == "" {
		user.Sign = createSign(user)
	}

	params := url.Values{}
	params.Add("uid", strconv.Itoa(int(user.Uid)))
	params.Add("username", user.Username)
	params.Add("now", strconv.Itoa(NowUnix()))
	params.Add("ip", user.Ip)
	params.Add("sign", user.Sign)
	c := &http.Cookie{
		Name:  conf.CookieKey,
		Value: params.Encode(),
		Path:  "/",
	}
	http.SetCookie(writer, c)
}

func createSign(user *models.LoginUser) string {
	str := fmt.Sprintf("uid=%d&username=%s&secret=%s&now=%d",
		user.Uid, user.Username, user.Sign, user.Now)
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
