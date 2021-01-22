package chat

import "fmt"

type Chat struct {

}

func (c* Chat) Run(){
	for{
		fmt.Println("Chat Running")
	}
}

func Start(){
	c := Chat{}
	c.Run()
}