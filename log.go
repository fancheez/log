// Copyright 2014 Unknwon
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package log

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var (
	Prefix     = "[Log]"
	TimeFormat = "06-01-02 15:04:05"

	NonColor    bool
	ShowDepth   bool
	CallerDepth = 2

	LevelFlags = []string{"DEBUG", " INFO", " WARN", "ERROR", "FATAL"}
)

func init() {
	if runtime.GOOS == "windows" {
		NonColor = true
	}
}

type Level int

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func Print(level Level, format string, args ...interface{}) {
	var depthInfo string
	if ShowDepth {
		pc, file, line, ok := runtime.Caller(CallerDepth)
		if ok {
			// Get caller function name.
			fn := runtime.FuncForPC(pc)
			var fnName string
			if fn == nil {
				fnName = "?()"
			} else {
				fnName = strings.TrimLeft(filepath.Ext(fn.Name()), ".") + "()"
			}
			depthInfo = fmt.Sprintf("[%s:%d %s] ", filepath.Base(file), line, fnName)
		}
	}
	if NonColor {
		fmt.Printf("%s %s [%s] %s%s\n",
			Prefix, time.Now().Format(TimeFormat), LevelFlags[level], depthInfo,
			fmt.Sprintf(format, args...))
		if level == FATAL {
			os.Exit(1)
		}
		return
	}

	switch level {
	case DEBUG:
		fmt.Printf("%s \033[36m%s\033[0m [\033[34m%s\033[0m] %s%s\n",
			Prefix, time.Now().Format(TimeFormat), LevelFlags[level], depthInfo,
			fmt.Sprintf(format, args...))
	case INFO:
		fmt.Printf("%s \033[36m%s\033[0m [\033[32m%s\033[0m] %s%s\n",
			Prefix, time.Now().Format(TimeFormat), LevelFlags[level], depthInfo,
			fmt.Sprintf(format, args...))
	case WARNING:
		fmt.Printf("%s \033[36m%s\033[0m [\033[33m%s\033[0m] %s%s\n",
			Prefix, time.Now().Format(TimeFormat), LevelFlags[level], depthInfo,
			fmt.Sprintf(format, args...))
	case ERROR:
		fmt.Printf("%s \033[36m%s\033[0m [\033[31m%s\033[0m] %s%s\n",
			Prefix, time.Now().Format(TimeFormat), LevelFlags[level], depthInfo,
			fmt.Sprintf(format, args...))
	case FATAL:
		fmt.Printf("%s \033[36m%s\033[0m [\033[35m%s\033[0m] %s%s\n",
			Prefix, time.Now().Format(TimeFormat), LevelFlags[level], depthInfo,
			fmt.Sprintf(format, args...))
		os.Exit(1)
	default:
		fmt.Printf("%s %s [%s] %s%s\n",
			Prefix, time.Now().Format(TimeFormat), LevelFlags[level], depthInfo,
			fmt.Sprintf(format, args...))
	}
}

func Debug(format string, args ...interface{}) {
	Print(DEBUG, format, args...)
}

func Warn(format string, args ...interface{}) {
	Print(WARNING, format, args...)
}

func Info(format string, args ...interface{}) {
	Print(INFO, format, args...)
}

func Error(format string, args ...interface{}) {
	Print(ERROR, format, args...)
}

func Fatal(format string, args ...interface{}) {
	Print(FATAL, format, args...)
}
