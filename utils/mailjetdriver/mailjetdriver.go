package mailjetdriver

import (
	"backend_capstone/configs"
	"backend_capstone/services/transaction"
	"backend_capstone/services/transaction/dto"
	"errors"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/mailjet/mailjet-apiv3-go/v4"
)

type MailjetDriver struct {
	CompanyName   string
	CompanyEmail  string
	PrivateKey    string
	PublicKey     string
	mailjetClient *mailjet.Client
}

func NewTransactionMailjetService(configs *configs.AppConfig) transaction.Mailjet {
	api := &MailjetDriver{
		CompanyName:   configs.App.NAME,
		CompanyEmail:  configs.API_Mailjet.EMAIL,
		PrivateKey:    configs.API_Mailjet.PRIVATE_KEY,
		PublicKey:     configs.API_Mailjet.PUBLIC_KEY,
		mailjetClient: mailjet.NewMailjetClient(configs.API_Mailjet.PUBLIC_KEY, configs.API_Mailjet.PRIVATE_KEY),
	}
	return api
}

func (d *MailjetDriver) SendBill(name string, email string, bill dto.BillClient) (err error) {
	var text string

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return errors.New("unable to get the current filename")
	}
	log.Print()
	filePath := filepath.Join(filepath.Dir(filename), "./media/billing.html")
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Print("Membuka file gagal")
		text = HTML_BILLING
	} else {
		// Convert []byte to string
		text = string(fileContent)
	}

	text = strings.Replace(text, "<%name%>", name, 1)
	text = strings.Replace(text, "<%product_name%>", bill.Product, 1)
	text = strings.Replace(text, "<%payment_price%>", strconv.Itoa(int(bill.Billed)), 1)
	text = strings.Replace(text, "<%payment_details%>", bill.PaymentDetails, 1)
	text = strings.Replace(text, "<%deadline%>", bill.Deadline.Format("02-01-2006 15:04:05"), 1)
	if strings.Contains(bill.VaNumber, "http") {
		paymentImg := `<img src="` + bill.VaNumber + `" alt=\"QR Code\" style=\"display: block; margin-left: auto; margin-right: auto;\">`
		text = strings.Replace(text, "<%payment_code%>", paymentImg, 1)
	} else {
		text = strings.Replace(text, "<%payment_code%>", bill.VaNumber, 1)
	}

	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Name:  d.CompanyName,
				Email: d.CompanyEmail,
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: email,
					Name:  name,
				},
			},
			Subject:  "Transaction Is on Due!",
			TextPart: "Dear " + name + ", your transaction is waiting for you!",
			HTMLPart: text,
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	_, err = d.mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Print("Email billing gagal dikirim")
		log.Print(err)
		return
	}
	log.Print("Berhasil kirim email billing")
	return
}

func (d *MailjetDriver) SendInvoice(name string, email string, bill dto.BillClient) (err error) {
	var text string

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return errors.New("unable to get the current filename")
	}
	log.Print()
	filePath := filepath.Join(filepath.Dir(filename), "./media/invoice.html")
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Print("Membuka file gagal invoice")
		text = HTML_INVOICE
	} else {
		// Convert []byte to string
		text = string(fileContent)
	}

	text = strings.Replace(text, "<%name%>", name, 1)
	text = strings.Replace(text, "<%product_name%>", bill.Product, 1)
	text = strings.Replace(text, "<%payment_price%>", strconv.Itoa(int(bill.Billed)), 1)
	text = strings.Replace(text, "<%payment_details%>", bill.PaymentDetails, 1)

	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: d.CompanyEmail,
				Name:  d.CompanyName,
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: email,
					Name:  name,
				},
			},
			Subject:  "Transaction Compleate!",
			TextPart: "Dear " + name + ", your transaction has been compleated!",
			HTMLPart: text,
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	_, err = d.mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Print("Email invoice gagal dikirim")
		log.Print(err)
		return
	}
	log.Print("Berhasil kirim email invoice")
	return
}

var (
	HTML_BILLING = `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Billing</title>
	</head>
	<body style="margin: 0; padding: 0;">
		<table align="center" border="0" cellpadding="0" cellspacing="0" width="600" style="border: 1px solid #cccccc;">
			<tr>
				<td align="center" bgcolor="#FFD700" style="padding: 40px 0 30px 0;">
					<img src="https://raw.githubusercontent.com/ixtza/backend_capstone/feat/system/utils/mailjetdriver/media/payzone.svg" alt="Payzone" width="300" height="100" style="display: block;" />
				 </td>
			</tr>
			<tr>
				<td bgcolor="#ffffff" style="padding: 40px 30px 40px 30px; font-family: Arial, sans-serif; font-size: 14px;">
					<table border="0" cellpadding="0" cellspacing="0" width="100%">
						<tr>
						 <td style="padding: 40px 30px 40px 30px; font-family: Arial, sans-serif; font-size: 30px; font-weight: bold;">
							Your transaction is waiting for you!
						 </td>
						</tr>
						<tr>
						 <td style="padding: 40px 30px 40px 30px;">
							Dear <%name%>, your purchase for <%product_name%> is on due. Please compleate this purchase by paying Rp<%payment_price%>,00 via <strong><%payment_details%></strong> before <%deadline%>.
						 </td>
						</tr>
						<tr>            
						 <td style="padding: 40px 30px 40px 30px; font-size: 20px;  font-weight: bold; text-align: center;">
							<%payment_code%>
						 </td>
						</tr>
					 </table>
				 </td>
			</tr>
			<tr>
					<td bgcolor="#FFD700" style="padding: 30px 30px 30px 30px; font-family: Arial, sans-serif; font-size: 14px;">
					Payzone 2022 &#169
				 </td>
			</tr>
		 </table>
	 </body>
	</html>`
	HTML_INVOICE = `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Invoice</title>
	</head>
	<body style="margin: 0; padding: 0;">
		<table align="center" border="0" cellpadding="0" cellspacing="0" width="600" style="border: 1px solid #cccccc;">
			<tr>
				<td align="center" bgcolor="#FFD700" style="padding: 40px 0 30px 0;">
					<img src="https://raw.githubusercontent.com/ixtza/backend_capstone/feat/system/utils/mailjetdriver/media/payzone.svg" alt="Payzone" width="300" height="100" style="display: block;" />
				 </td>
			</tr>
			<tr>
				<td bgcolor="#ffffff" style="padding: 40px 30px 40px 30px; font-family: Arial, sans-serif; font-size: 14px;">
					<table border="0" cellpadding="0" cellspacing="0" width="100%">
						<tr>
						 <td style="padding: 40px 30px 40px 30px; font-family: Arial, sans-serif; font-size: 30px; font-weight: bold;">
							Your transaction has been Compleated!
						 </td>
						</tr>
						<tr>
						 <td style="padding: 40px 30px 40px 30px;">
							Dear <%name%>, your purchase for <%product_name%> has been compleated.
							<table style="padding: 40px 30px 40px 30px;">
								<tr>
									<td>
										Item : <%product_name%>
									</td>
								</tr>
								<tr>
									<td>
										Price :Rp<%payment_price%>,00
									</td>
								</tr>
								<tr>
									<td>
										Payment Method: <%payment_details%>
									</td>
								</tr>
							</table>
						 </td>
						</tr>
						<tr>    
						</tr>
					 </table>
				 </td>
			</tr>
			<tr>
					<td bgcolor="#FFD700" style="padding: 30px 30px 30px 30px; font-family: Arial, sans-serif; font-size: 14px;">
					Payzone 2022 &#169
				 </td>
			</tr>
		 </table>
	 </body>
	</html>`
)
