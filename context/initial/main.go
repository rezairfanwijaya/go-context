package main

import (
	"context"
	"log"
)

func main() {
	contextBackgroud := context.Background() // context kosongan
	contextTodo := context.TODO()            // context kosongan yang belum tau akan diapakan

	log.Println("context background : ", contextBackgroud)
	log.Println("context todo : ", contextTodo)
}
