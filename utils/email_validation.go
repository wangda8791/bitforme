package utils

import (
	"github.com/badoux/checkmail"
)

func Email_Validation(email string) error {

	// Format
	err_f := checkmail.ValidateFormat(email)
	if err_f != nil {		
		return err_f
	}

	// Domain
	// err_d := checkmail.ValidateHost(email)
	// if err_d != nil {
	// 	fmt.Println(err_d)
		
	// 	return err_d
	// }

	// User
	// err := checkmail.ValidateHost("unknown-user-129083726@gmail.com")
	// if smtpErr, ok := err.(checkmail.SmtpError); ok && err != nil {
	// 	fmt.Printf("Code: %s, Msg: %s", smtpErr.Code(), smtpErr)
	// }

	return nil
}