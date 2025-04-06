package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/wneessen/go-mail"
)

func main() {
	to := flag.String("to", "", "Recipient email address")
	subject := flag.String("subject", "", "Email subject")
	body := flag.String("body", "", "Email body")

	flag.Parse()

	if !strings.Contains(*to, "@") {
		contacts := parseContacts(os.Getenv("CONTACTS"))
		*to = contacts[*to]
	}

	paths := flag.Args()

	username := os.Getenv("GMAIL_USERNAME")
	password := os.Getenv("GMAIL_PASSWORD")

	from := username
	smtpHost := "smtp.gmail.com"

	message := mail.NewMsg()

	if err := message.From(from); err != nil {
		panic(err)
	}

	if err := message.To(*to); err != nil {
		panic(err)
	}

	message.Subject(*subject)
	message.SetBodyString(mail.TypeTextPlain, *body)

	for _, path := range paths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			fmt.Printf("File %s does not exist\n", path)
			continue
		}

		message.AttachFile(path)
	}

	client, err := mail.NewClient(smtpHost, mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(username), mail.WithPassword(password))

	if err != nil {
		panic(err)
	}

	if err := client.DialAndSend(message); err != nil {
		panic(err)
	}

	fmt.Println("Email sent successfully!")
}

func parseContacts(contacts string) map[string]string {
	result := make(map[string]string)
	contactStrings := strings.SplitSeq(contacts, ";")

	for contactString := range contactStrings {
		contact := strings.Split(strings.Trim(contactString, " "), " ")
		if len(contact) != 2 {
			continue
		}
		name := contact[0]
		email := contact[1]
		result[name] = email
	}
	return result
}
