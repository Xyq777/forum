package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func register(i Info) error {

	if strings.ContainsAny(i.Password, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") && len(i.Password) < 8 {
		return errors.New("密码格式错误")
	} else {
		err := setUser(i.Username, i.Password)
		if err != nil {
			return err
		}
		return nil

	}
}

func logIn(i Info) error {

	user, err := getUser(i.Username)
	if err != nil {
		return err
	}
	if user.Password != i.Password {
		return errors.New("密码错误")

	}
	return nil

}
func currentTime() string {
	currentTime1 := time.Now()

	// 获取年、月、日、小时和分钟
	year := currentTime1.Year()
	month := currentTime1.Month()
	day := currentTime1.Day()
	hour := currentTime1.Hour()
	minute := currentTime1.Minute()

	// 将它们转化为字符串形式
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d", year, month, day, hour, minute)

}
