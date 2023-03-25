package util

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func GenerateTransCode() string {
	currentTime := time.Now()
	saltCode := fmt.Sprint(currentTime.Format("010206-150405"))
	uniqueStr := uuid.NewString()[:8]
	return "INV" + "-" + saltCode + "-" + uniqueStr
}
