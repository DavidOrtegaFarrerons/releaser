package release

import (
	"fmt"
	"time"
)

func GenerateTag(env string) string {
	t := time.Now()
	return fmt.Sprintf("%s-%d-%02d-%02d", env, t.Year(), int(t.Month()), t.Day())
}
