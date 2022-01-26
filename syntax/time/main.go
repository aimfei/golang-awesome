package time_example

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	duration, _ := time.ParseDuration("1h")
	add := now.Add(-24 * duration)
	fmt.Println(add.Unix())
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(add.Format("2006-01-02 15:04:05"))
	sub := add.Sub(now)
	fmt.Println(sub.String())

	tick := time.Tick(1 * time.Second)
	for {
		select {
		case <-tick:
			fmt.Println("tick runing ")
		default:
			fmt.Println("running finish")
			time.Sleep(999 * time.Millisecond)
		}
	}

}
