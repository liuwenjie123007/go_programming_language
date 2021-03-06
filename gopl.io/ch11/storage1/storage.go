package storage

import (
	"fmt"
	"log"
	"net/smtp"
)

func bytesInuse(user string) int64 {
	return 0 /* ... */
}

// 邮件发送者配置
// 注意：永远不要把密码放到源代码中
const sender = "notification@example.com"
const password = "correcthoresebatterystaple"
const hostname = "smt.example.com"

const template = `Warning: you are using %d bytes of storage, %d%% of your quota.`

func CheckQuota(username string) {
	used := bytesInuse(username)
	const quota = 1000000000 // 1GB
	percent := 100 * used    // quota
	if percent < 90 {
		return // OK
	}
	msg := fmt.Sprintf(template, used, percent)
	auth := smtp.PlainAuth("", sender, password, hostname)
	err := smtp.SendMail(hostname+":587", auth, sender, []string{username}, []byte(msg))
	if err != nil {
		log.Printf("smtp.SendMail(%s) failed: %s", username, err)
	}
}
