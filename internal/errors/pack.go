package errors

import "fmt"

type ErrPackNotFound struct {
	BaseErr error
	packId  string
}

func (e ErrPackNotFound) Error() string {
	return fmt.Sprintf("pack with ID %s not found: %s", e.packId, e.BaseErr.Error())
}

func NewErrPackNotFound(baseErr error, packId string) ErrPackNotFound {
	return ErrPackNotFound{
		BaseErr: baseErr,
		packId:  packId,
	}
}
