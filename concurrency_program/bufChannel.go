package main
import (
	"fmt"
)

func main(){
	numbers := make(chan int, 5)
	counter := 10
	for i := 0; i < counter; i++ {
		select {
			// 往通道送东西
		case numbers <- i:
		default:
			fmt.Println("Not enough space for ", i)
			
		}
	}
	for i := 0; i < counter+5; i++ {
		select {
			// 从通道拿东西
		case num := <- numbers:
			fmt.Println(num)
		default:
			fmt.Println("Finish！")
			break
			
		}
	}
	
}
