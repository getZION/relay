package handler

import (
	"github.com/getzion/relay/api/dwn"
	"github.com/getzion/relay/api/models"
)

type RequestContext struct {
	ModelManager *models.ModelManager
	Message      *dwn.Message
}
