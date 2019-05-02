package zaman

import (
	"fmt"
	"time"
)

func BugünNe() {
	fmt.Println("Henuz bilmiyorum")
	p := fmt.Println
	now := time.Now()
	p(now)
}
