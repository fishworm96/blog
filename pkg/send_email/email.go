package send_email

import (
	"blog/setting"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"

	"github.com/jordan-wright/email"
)

func SendCode(toUserEmail, code string) error {
	e := email.NewEmail()
	e.From = "Get <eb2ban_zxh@qq.com>"
	e.To = []string{toUserEmail}
	e.Subject = "验证码已发送，请查收"
	e.HTML = []byte("您的验证码: <b>" + code + "</b><br><b>验证码需要在5分钟内使用</b>")
	return e.Send("smtp.qq.com:25", smtp.PlainAuth("", "eb2ban_zxh@qq.com", setting.Conf.EmailSecretCode, "smtp.qq.com"))
	// return e.SendWithTLS("smtp.qq.com:25",
	// 	smtp.PlainAuth("", "eb2ban_zxh@qq.com", "", "smtp.qq.com"),
	// 	&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
}

func GetRand() string {
	rand.Seed(time.Now().UnixNano())
	s := ""
	for i := 0; i < 6; i++ {
		s += strconv.Itoa(rand.Intn(10))
	}
	return s
}
