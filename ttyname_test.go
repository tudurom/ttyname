package ttyname

import "fmt"

func ExampleTTY() {
	tty, err := TTY()
	if err != nil {
		panic(err)
	}

	fmt.Println(tty)
}
