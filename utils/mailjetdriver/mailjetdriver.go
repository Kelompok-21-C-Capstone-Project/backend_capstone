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
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return errors.New("unable to get the current filename")
	}
	log.Print()
	filePath := filepath.Join(filepath.Dir(filename), "./media/billing.html")
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Print("Membuka file gagal")
		log.Print(err)
		return
	}

	// Convert []byte to string
	text := string(fileContent)

	text = strings.Replace(text, "<%name%>", name, 1)
	text = strings.Replace(text, "<%product_name%>", bill.Product, 1)
	text = strings.Replace(text, "<%payment_price%>", strconv.Itoa(int(bill.Billed)), 1)
	text = strings.Replace(text, "<%payment_details%>", bill.PaymentDetails, 1)
	text = strings.Replace(text, "<%deadline%>", bill.Deadline.Format("02-01-2006 15:04:05"), 1)
	if strings.Contains(bill.VaNumber, "http") {
		paymentImg := "<img src=\"" + bill.VaNumber + "\" alt=\"QR Code\" style=\"display: block; margin-left: auto; margin-right: auto;\">"
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
		log.Print("Email gagal dikirim")
		log.Print(err)
		return
	}
	log.Print("Berhasil kirim email")
	return
}

func (d *MailjetDriver) SendInvoice(name string, email string, bill dto.BillClient) (err error) {
	fileContent, err := ioutil.ReadFile("./media/invoice.html")
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string
	text := string(fileContent)

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
	res, err := d.mailjetClient.SendMailV31(&messages)
	log.Printf("Data: %+v\n", res)
	if err != nil {
		return
	}
	return
}
