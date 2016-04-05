package main
import (
	"fmt"
	"github.com/go-martini/martini"

)


func main() {
	fmt.Println("Starting Server .. ")
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.RunOnAddr(":9090")
}