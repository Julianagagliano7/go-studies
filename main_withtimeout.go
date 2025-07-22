package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ctx, cancel := context.WithTimeout(
// 		context.Background(),
// 		5*time.Second,
// 	)

// 	defer cancel() //não faz diferença, é só para não dar erro no cancel

// 	printUntilCancel(ctx)
// }

// func printUntilCancel(ctx context.Context) {
// 	count := 0

// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("timeout reached, exiting")
// 			return
// 		default:
// 			time.Sleep(1 * time.Second)
// 			fmt.Printf("printing until cancel, number = %d \n", count)
// 			count += 1
// 		}
// 	}
// }
