# ttyname

Prints the file name of the terminal connected to standard input.

## Usage

```go
tty, err := ttyname.TTY()
if err != nil {
	panic(err)
}

fmt.Println(tty)
```
