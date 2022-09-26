package test_log

import (
	"context"
	"errors"
	"testing"
	"time"

	clog "github.com/pip-services3-gox/pip-services3-components-gox/log"
	"github.com/stretchr/testify/assert"
)

type LoggerFixture struct {
	logger *clog.CachedLogger
}

func NewLoggerFixture(logger *clog.CachedLogger) *LoggerFixture {
	lf := LoggerFixture{}
	lf.logger = logger
	return &lf
}

func (c *LoggerFixture) TestLogLevel(t *testing.T) {
	assert.True(t, c.logger.Level() >= clog.LevelNone)

	assert.True(t, c.logger.Level() <= clog.LevelTrace)
}

func (c *LoggerFixture) TestSimpleLogging(t *testing.T) {
	ctx := context.Background()
	c.logger.SetLevel(clog.LevelTrace)

	c.logger.Fatal(ctx, "", nil, "Fatal error message")
	c.logger.Error(ctx, "", nil, "Error message")
	c.logger.Warn(ctx, "", "Warning message")
	c.logger.Info(ctx, "", "Information message")
	c.logger.Debug(ctx, "", "Debug message")
	c.logger.Trace(ctx, "", "Trace message")
	c.logger.Dump(ctx)

	select {
	case <-time.After(time.Duration(1000) * time.Millisecond):
	}
}

func (c *LoggerFixture) TestErrorLogging(t *testing.T) {
	ctx := context.Background()

	var ex error = errors.New("Testing error throw")

	c.logger.Fatal(ctx, "123", ex, "Fatal error")
	c.logger.Error(ctx, "123", ex, "Recoverable error")
	assert.NotNil(t, ex)
	c.logger.Dump(ctx)
	select {
	case <-time.After(time.Duration(1000) * time.Millisecond):
	}
}
