package defer_panic

import "fmt"

func main() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("Something bad happened!")
}
