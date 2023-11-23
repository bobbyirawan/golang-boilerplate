package logger

import (
	"go-chat/pkg/utils"
	"testing"
)

func TestFields(t *testing.T) {

	logger := utils.NewLogger(utils.NewTimeService())

	logger.Info()

}
