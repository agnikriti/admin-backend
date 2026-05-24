package utils

import (
	"fmt"
	"strconv"

	"agnikriti_admin_backend/config"

	"gopkg.in/gomail.v2"
)

func SendProposalEmail(
	title string,
	description string,
	email string,
	mobile string,
) error {

	port, err := strconv.Atoi(config.AppConfig.SMTPPort)

	if err != nil {
		return err
	}

	mailer := gomail.NewDialer(
		config.AppConfig.SMTPHost,
		port,
		config.AppConfig.SMTPEmail,
		config.AppConfig.SMTPPassword,
	)

	message := gomail.NewMessage()

	message.SetHeader(
		"From",
		config.AppConfig.SMTPEmail,
	)

	message.SetHeader(
		"To",
		config.AppConfig.SMTPEmail,
	)

	message.SetHeader(
		"Subject",
		title,
	)

	body := fmt.Sprintf(
		`
		New Proposal Received

		Description:
		%s

		Sender Email:
		%s

		Sender Mobile:
		%s
		`,
		description,
		email,
		mobile,
	)

	message.SetBody("text/plain", body)

	err = mailer.DialAndSend(message)

	return err
}
