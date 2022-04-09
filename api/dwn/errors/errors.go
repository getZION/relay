package errors

import "github.com/sirupsen/logrus"

const (
	RequestLevelProcessingErrorMessage string = "The request could not be processed correctly"
	TargetDIDNotFoundErrorMessage      string = "Target DID not found within the Identity Hub instance"

	ImproperlyConstructedErrorMessage   string = "The message was malformed or improperly constructed"
	InterfaceNotImplementedErrorMessage string = "The interface method is not implemented"
	AuthorizationFailedErrorMessage     string = "The message failed authorization requirements"

	MessageSuccessfullyMessage string = "The message was successfully processed"
)

type MessageLevelError struct {
	Message string
	Code    int
	Error   error
}

func NewMessageLevelError(code int, message string, err error) *MessageLevelError {

	if code >= 500 {
		logrus.Errorf("%s: %v", message, err)
	} else {
		logrus.Infof("%s: %v", message, err)
	}

	return &MessageLevelError{
		Code:    code,
		Message: message,
		Error:   err,
	}
}
