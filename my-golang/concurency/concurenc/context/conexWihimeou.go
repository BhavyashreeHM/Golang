package context

import (
	"context"
	"fmt"
	"time"
)

// {

// }

// receiver := <-ch

func chcckEvenorOdd(cx context.Context, n int) string {
	select {
	case <-cx.Done():
		return  "Operaion Cancelled"
	default:
		if n%2 == 0 {
			return fmt.Sprintf("the number is Even %d", n)
		} else {
			return fmt.Sprintf("the number is  Odd %d", n)
		}

	}
}

func ma() {
	cx := context.TODO()
	resul:=chcckEvenorOdd(cx, 5)
	fmt.Println("Resul from conex backgrounf",resul)

	cx=context.Background()
	
	cx, cancel:= context.WithTimeout(cx, 1*time.Second)
	defer cancel()
	resul=chcckEvenorOdd(cx, 22)
	fmt.Println("Resul from conex backgrounf",resul)

	time.Sleep(2*time.Second)
	resul=chcckEvenorOdd(cx, 252)
	fmt.Println("Resul from conex backgrounf",resul)
	





}
