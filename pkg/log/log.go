/*
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     https://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package log

import (
	"fmt"
	"io"
	"log"
	"os"
)

type LogLevel int

const (
	LogPrefix = "[go-adc] "
	ErrorPrefix = "[error] "
	WarningPrefix = "[warn] "
	InfoPrefix = "[info] "
	DebugPrefix = "[debug] "
	DefaultLogLevel = DebugLevel
)

const (
	ErrorLevel LogLevel = iota
	WarningLevel
	InfoLevel
	DebugLevel
)

type Logger struct {
	level LogLevel
	*log.Logger
}

var	logger = &Logger{
	level:  DefaultLogLevel,
	Logger: log.New(os.Stderr, LogPrefix, log.LstdFlags),
}

func Init(out io.Writer) {
	logger.SetOutput(out)
}

func Error(format string, v ...interface{}) {
	if logger.level >= ErrorLevel {
		logger.Println(fmt.Sprintf(ErrorPrefix + format, v...))
	}
}

func Warning(format string, v ...interface{}) {
	if logger.level >= WarningLevel {
		logger.Println(fmt.Sprintf(WarningPrefix + format, v...))
	}
}

func Info(format string, v ...interface{}) {
	if logger.level >= InfoLevel {
		logger.Println(fmt.Sprintf(InfoPrefix + format, v...))
	}
}

func Debug(format string, v ...interface{}) {
	if logger.level >= DebugLevel {
		logger.Println(fmt.Sprintf(DebugPrefix + format, v...))
	}
}
