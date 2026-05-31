package utils

import (
	"fmt"

	"agnikriti_admin_backend/config"

	"github.com/resend/resend-go/v3"
)

func SendProposalEmail(
	title string,
	description string,
	email string,
	mobile string,
	quote int32,
) error {

	client := resend.NewClient(config.AppConfig.ResendAPIKey)

	quoteRow := ""
	if quote > 0 {
		quoteRow = fmt.Sprintf(`
			<tr>
				<td style="padding:10px 16px;border-bottom:1px solid #e5e7eb;color:#6b7280;font-weight:600;width:140px;white-space:nowrap;">Budget Quote</td>
				<td style="padding:10px 16px;border-bottom:1px solid #e5e7eb;color:#111827;">&#8377;%d</td>
			</tr>`, quote)
	}

	body := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head><meta charset="UTF-8"><meta name="viewport" content="width=device-width,initial-scale=1.0"></head>
<body style="margin:0;padding:0;background-color:#f3f4f6;font-family:Arial,sans-serif;">
  <table width="100%%" cellpadding="0" cellspacing="0" style="background-color:#f3f4f6;padding:40px 0;">
    <tr><td align="center">
      <table width="600" cellpadding="0" cellspacing="0" style="background-color:#ffffff;border-radius:8px;overflow:hidden;box-shadow:0 1px 3px rgba(0,0,0,0.1);">

        <tr>
          <td style="background-color:#1d4ed8;padding:28px 32px;">
            <p style="margin:0;font-size:13px;color:#bfdbfe;letter-spacing:1px;text-transform:uppercase;">Agnikriti Solutions</p>
            <h1 style="margin:6px 0 0;font-size:22px;color:#ffffff;">New Proposal Received</h1>
          </td>
        </tr>

        <tr>
          <td style="padding:28px 32px;">
            <h2 style="margin:0 0 6px;font-size:18px;color:#111827;">%s</h2>
            <p style="margin:0 0 24px;font-size:14px;color:#6b7280;">A new service proposal has been submitted. Details are listed below.</p>

            <table width="100%%" cellpadding="0" cellspacing="0" style="border:1px solid #e5e7eb;border-radius:6px;font-size:14px;border-collapse:collapse;">
              <tr>
                <td style="padding:10px 16px;border-bottom:1px solid #e5e7eb;background-color:#f9fafb;font-weight:700;color:#374151;" colspan="2">Proposal Details</td>
              </tr>
              <tr>
                <td style="padding:10px 16px;border-bottom:1px solid #e5e7eb;color:#6b7280;font-weight:600;width:140px;white-space:nowrap;">Description</td>
                <td style="padding:10px 16px;border-bottom:1px solid #e5e7eb;color:#111827;">%s</td>
              </tr>
              <tr>
                <td style="padding:10px 16px;border-bottom:1px solid #e5e7eb;color:#6b7280;font-weight:600;">Email</td>
                <td style="padding:10px 16px;border-bottom:1px solid #e5e7eb;color:#111827;">%s</td>
              </tr>
              <tr>
                <td style="padding:10px 16px;border-bottom:1px solid #e5e7eb;color:#6b7280;font-weight:600;">Mobile</td>
                <td style="padding:10px 16px;border-bottom:1px solid #e5e7eb;color:#111827;">%s</td>
              </tr>
              %s
            </table>
          </td>
        </tr>

        <tr>
          <td style="padding:16px 32px;background-color:#f9fafb;border-top:1px solid #e5e7eb;text-align:center;">
            <p style="margin:0;font-size:12px;color:#9ca3af;">This is an automated notification from Agnikriti Solutions.</p>
          </td>
        </tr>

      </table>
    </td></tr>
  </table>
</body>
</html>`,
		title, description, email, mobile, quoteRow,
	)

	params := &resend.SendEmailRequest{
		From:    fmt.Sprintf("Agnikriti Solutions <%s>", config.AppConfig.SMTPEmail),
		To:      []string{config.AppConfig.SMTPEmail},
		Subject: fmt.Sprintf("New Proposal: %s", title),
		Html:    body,
	}

	_, err := client.Emails.Send(params)
	return err
}
