package persistence

import "strings"

func isUniqueConstraintError(err error, field string) bool {
	if err != nil && strings.Contains(err.Error(), "unique") && strings.Contains(err.Error(), field) {
		return true
	}
	return false
}
