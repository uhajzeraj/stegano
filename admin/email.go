package main

import (
	"log"
	"net/smtp"
)

func send(body string, to string) {
	from := "steganoteam@gmail.com"
	pass := "Stegano123"
	// to := "foobarbazz@mailinator.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("sent, visit http://foobarbazz.mailinator.com")
}

// $ip = $_SERVER["REMOTE_ADDR"];
// mysqli_query($connection, "INSERT INTO `ip` (`address` ,`timestamp`)VALUES ('$ip',CURRENT_TIMESTAMP)");
// $result = mysqli_query($connection, "SELECT COUNT(*) FROM `ip` WHERE `address` LIKE '$ip' AND `timestamp` > (now() - interval 10 minute)");
// $count = mysqli_fetch_array($result, MYSQLI_NUM);

// if($count[0] > 3){
//   echo "Your are allowed 3 attempts in 10 minutes";
// }
