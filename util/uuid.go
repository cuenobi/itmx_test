package util

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

func GenerateUuid() string {
	uuid := uuid.New().String()
	now := time.Now().Format("2006-01-02 15:04:05")
	now = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(now, "-", ""), ":", ""), " ", "")
	result := now + "-" + uuid
	return result
}
