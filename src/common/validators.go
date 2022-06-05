package common

import (
	"regexp"

	"github.com/google/uuid"
)

func IsValidUUID(id string) bool {
	if len(id) == 0 {
		return true
	}

	_, err := uuid.Parse(id)
	return err == nil
}

func CheckRegexp(reg, str string) bool {
	return regexp.MustCompile(reg).Match([]byte(str))
}
