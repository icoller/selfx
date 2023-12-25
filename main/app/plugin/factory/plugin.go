/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-25 14:17:01
 * @Desc:
 */
package factory

import (
	"errors"
	"fmt"
	"path/filepath"
	"selfx/app/plugin/entity"
	"selfx/config"
	"selfx/constant"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewPlugin(entry entity.PluginEntry) (*entity.Plugin, error) {
	info := entry.Info()
	if info == nil {
		return nil, errors.New("plugin info undefined")
	}
	if info.ID == "" {
		return nil, constant.ErrIdRequired
	}
	if !info.RunEnable {
		info.CronEnable = false // 如果未启用run，则也不启用cron
	}
	return &entity.Plugin{Entry: entry, Info: info, Log: newPluginLog(info.ID)}, nil
}

func GetPluginLogFilePath(id string) string {
	return filepath.Join(constant.LogDir, "plugin", fmt.Sprintf("%s.log", id))
}

func newPluginLog(id string) *zap.Logger {
	return zap.New(
		zapcore.NewCore(zapcore.NewJSONEncoder(zapcore.EncoderConfig{
			MessageKey:     "msg",
			LevelKey:       "level",
			TimeKey:        "time",
			NameKey:        "logger",
			CallerKey:      "file",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.EpochMillisTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder, // 短路径编码器
			EncodeName:     zapcore.FullNameEncoder,
		}),
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(&lumberjack.Logger{
				Filename:   GetPluginLogFilePath(id),
				MaxSize:    config.Config.Log.Plugin.MaxSize,
				MaxAge:     config.Config.Log.Plugin.MaxAge,
				MaxBackups: config.Config.Log.Plugin.MaxBackups,
				Compress:   config.Config.Log.Plugin.Compress,
			})),
			zap.NewAtomicLevelAt(zapcore.DebugLevel),
		),
		zap.AddCallerSkip(0), zap.AddCaller(),
	)
}
