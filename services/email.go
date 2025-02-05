package services

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/bn_funds/utils"
)

func SendResetTokenEmail(to string, token string) error {
	from := os.Getenv("EMAIL")
	pass := os.Getenv("EMAILPASSWORD")

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Reset Token\n\n" +
		"You requested to reset your password. \n" +
		"If you did not request a change of password, please ignore this e-mail.\n\n" +
		"To reset your password follow this link: \n" + utils.GetEnv("BASE_URL", "") + "user/changepassword/do?email=" + to + "&t=" + token

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return err
	}

	log.Print(fmt.Sprintf("Reset token sent to %s : token %s", to, token))
	return err
}
