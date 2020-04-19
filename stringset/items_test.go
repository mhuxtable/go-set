package stringset

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewSetItem() string {
	return fmt.Sprintf("%d", rand.Int())
}
