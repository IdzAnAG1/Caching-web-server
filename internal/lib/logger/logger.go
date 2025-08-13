package logger

import (
	"CachingWebServer/internal/config"
	"errors"
	"log/slog"
	"os"
)

const (
	DEBUG = "debug"
	LOCAL = "local"
	PROD  = "prod"
)

func IdentifyLogger(cfg config.Config) (*slog.Logger, error) {
	switch cfg.Logger.Level {
	case DEBUG:
		return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource:   false,
			Level:       slog.LevelDebug,
			ReplaceAttr: nil,
		})), nil
	case LOCAL:
		return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource:   true,
			Level:       slog.LevelDebug,
			ReplaceAttr: nil,
		})), nil
	case PROD:
		return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource:   false,
			Level:       slog.LevelInfo,
			ReplaceAttr: nil,
		})), nil
	default:
		return nil, errors.New("Logger level is not specified ")
	}
}

// #TODO Implement
//type ZapHandler struct {
//	zap         *zap.Logger
//	level       *slog.Level
//	addSource   bool
//	replaceAttr func(groups []string, a slog.Attr) slog.Attr
//}
//
//func (z *ZapHandler) Enabled(ctx context.Context, level slog.Level) bool {
//	Level := slog.LevelInfo
//	if z.level != nil {
//		Level = z.level.Level()
//	}
//	return level >= Level
//}
//
//func (z *ZapHandler) Handle(_ context.Context, r slog.Record) error {
//}
//func attrToZapField(a slog.Attr) zap.Field {
//}
//
//func (z *ZapHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
//
//}
//func (z *ZapHandler) WithGroup(name string) slog.Handler {
//
//}
