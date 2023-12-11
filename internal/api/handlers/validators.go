package handlers

import "github.com/google/uuid"

func IsValidUUID(input string) bool {
	if _, err := uuid.Parse(input); err != nil {
		return false
	}

	return true
}
