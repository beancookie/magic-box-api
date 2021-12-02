package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/rs/zerolog/log"
)

type PaginationRequest struct {
	Page int `form:"page" json:"page"`
	Size int `form:"size" json:"size"`
}

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		log.Info().Msgf("%s %s", err.Key, err.Message)
	}

	return
}
