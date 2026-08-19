package main

import "log"

func order() {
	log.Println("START API Ordering....")
	sendEmail()
	sendSms()
	log.Println("SUCCESS API ordering")
}

func sendEmail() {
	defer func() {
		if r:= recover(); r != nil {
			log.Println("Fallback")
		}
	}()
	log.Println("START API sending email...")
	panic("Send email failed")
	log.Println("SUCCESS API sending email")
}

func sendSms() {
	log.Println("START API sending sms....")
	log.Println("SUCCESS API sending sms")
}

func main() {
	log.Println("Process API ordering....")
	order()
	log.Println("END API ORDER")
}
