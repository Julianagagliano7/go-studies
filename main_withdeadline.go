package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ctx, _ := context.WithDeadline(
// 		context.Background(),
// 		time.Now().Add(5*time.Second),
// 	)

// 	printUntilCancel(ctx)
// }

// func printUntilCancel(ctx context.Context) {
// 	count := 0

// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("Time to stop operation has come, exiting")
// 			return
// 		default:
// 			time.Sleep(1 * time.Second)
// 			fmt.Printf("printing until cancel, number = %d \n", count)
// 			count += 1
// 		}
// 	}
// }
