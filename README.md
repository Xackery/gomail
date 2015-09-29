# gomail
A super simple SMTP server for golang, catered for a gmail or google apps email account.

```
package main

import (
	"fmt"
	"github.com/xackery/gomail"
)

func main() {

	err := gomail.Send("yourname@gmail.com", "yourpassword", []string{"destination@somewhere.com"}, "subject", "content")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Sent!")
}
```