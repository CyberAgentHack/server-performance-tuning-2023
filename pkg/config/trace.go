package config

import (
	"fmt"

	"github.com/aws/aws-xray-sdk-go/xray"
	"go.uber.org/zap"
)

func ConfigureTraceProvider(logger *zap.Logger) error {
	err := xray.Configure(xray.Config{})
	if err != nil {
		return fmt.Errorf("failed to configure xlay: %w", err)
	}
	return nil
}
