package main

import (
	"fmt"
	"github.com/smartwalle/mailx"
)

func main() {
	var client = mailx.NewClient(
		"smartwalle@gmail.com",
		"your password",
		"smtp.gmail.com",
		"587",
		mailx.WithMaxIdle(1),
		mailx.WithMaxActive(1),
	)

	var m = mailx.NewHTMLMessage("Title", "<a href='http://www.google.com'>Hello Google</a>")
	m.From = "Yang<smartwalle@gmail.com>"
	m.To = []string{"917996695@qq.com"}

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i, client.Send(m))
		}(i)
	}

	select {}
}
