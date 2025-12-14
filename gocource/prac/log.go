package log

import (
	"fmt"
	"log"
)

func main() {
	log.Println(" i dnt know what is log")

	//log flags
	log.SetFlags(log.Ldate) //print only date
	log.Println("guess wt it print")

	log.SetFlags(log.Ltime) // print only time
	log.Println("now guess what is will print")

	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	fmt.Println("one more chance")

}
