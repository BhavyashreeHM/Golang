package context

import (
	"context"
	"fmt"
)

// func c() {
// 	ctx := context.WithValue(context.Background(), "place", "Msore")
// 	processConex(ctx)

// }

// t //  {}  :=  t y
func processConex(ctx context.Context) {
	place, ok := ctx.Value("place").(string)
	if !ok {
		fmt.Println("Place is no founc in conex")
		return
	}
	fmt.Println("Processing daa for place", place)

}
