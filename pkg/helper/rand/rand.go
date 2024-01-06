package rand

import (
	"fmt"
	"math/rand"
	"time"
)

// CreateCaptcha 生产6位随机数
func CreateCaptcha() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}
