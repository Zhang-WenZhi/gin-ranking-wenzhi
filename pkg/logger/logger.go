package logger

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"time"
	"runtime/debug"
	"net/http"
)

/*
1、请求的日志(这个日志也可以有nginx来做)
2、当程序报错崩溃的日志，让正常返回code=500的错误码而不是什么也没返回就是报错
3、程序员自己想打印的日志,比如logger.Write()
*/

// 工具包都放在pkg目录下

func init() {
	// 设置日志为json格式
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// 设置日志级别为debug
	// logrus.SetLevel(logrus.DebugLevel)
	// 禁用调用函数信息
	logrus.SetReportCaller(false)
}

func Write(msg string, fileName string) {
	setOutPutFile(logrus.InfoLevel, fileName)
	logrus.Info(msg)
}

func Debug(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.DebugLevel, "debug")
	logrus.WithFields(fields).Debug(args...)
}

func Info(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.InfoLevel, "info")
	logrus.WithFields(fields).Info(args...)
}

func Warn(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.WarnLevel, "warn")
	logrus.WithFields(fields).Warn(args...)
}

func Fatal(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.FatalLevel, "fatal")
	logrus.WithFields(fields).Fatal(args...)
}

func Error(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.ErrorLevel, "error")
	logrus.WithFields(fields).Error(args...)
}

func Panic(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.PanicLevel, "panic")
	logrus.WithFields(fields).Panic(args...)
}

func Trace(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.TraceLevel, "trace")
	logrus.WithFields(fields).Trace(args...)
}

func setOutPutFile(level logrus.Level, logName string) {
	if _, err := os.Stat("./runtime/log"); os.IsNotExist(err) {
		err = os.MkdirAll("./runtime/log", os.ModePerm)
		if err != nil {
			panic(fmt.Errorf("create log dir '%s' error: %s", "./runtime/log", err))
		}
	}
	

	timeStr := time.Now().Format("2006-01-02")
	fileName := path.Join("./runtime/log", logName + "_" + timeStr + ".log")

	var err error
	os.Stderr, err = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open log file error:", err)
		return
	}
	logrus.SetOutput(os.Stderr)
	logrus.SetLevel(level)
	// return
}

func LoggerToFile() gin.LoggerConfig {
	if _, err := os.Stat("./runtime/log"); os.IsNotExist(err) {
		err = os.Mkdir("./runtime/log", 0777)
		if err != nil {
			panic(fmt.Errorf("create log dir '%s' error: %s", "./runtime/log", err))
		}
	}

	timeStr := time.Now().Format("2006-01-02")
	fileName := path.Join("./runtime/log", "success" + "_" + timeStr + ".log")
	os.Stderr, _ = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	var conf = gin.LoggerConfig{
		Formatter: func(param gin.LogFormatterParams) string {
			return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
				param.TimeStamp.Format("2006-01-02 15:04:05"),
				param.ClientIP,
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
			)	
		},
		Output: io.MultiWriter(os.Stdout, os.Stderr),
	}
	return conf
}

func Recover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			if _, errDir := os.Stat("./runtime/log"); os.IsNotExist(errDir) {
				errDir = os.MkdirAll("./runtime/log", 0777)
				if errDir != nil {
					panic(fmt.Errorf("create log dir '%s' error: %s", "./runtime/log", errDir))
				}
			}

			timeStr := time.Now().Format("2006-01-02")
			fileName := path.Join("./runtime/log", "error_" + timeStr + ".log")

			f, errFile := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if errFile != nil {
				fmt.Println("open log file error:", errFile)
			}
			timeFileStr := time.Now().Format("2006-01-02 15:04:05")
			f.WriteString("panic error time:" + timeFileStr + "\n")
			f.WriteString(fmt.Sprintf("%v\n", err))
			f.WriteString("stacktrace from panic:\n" + string(debug.Stack()) + "\n")
			f.Close()
			c.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg": fmt.Sprintf("%v", err),
			})
			// 终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}
	}()
	c.Next()
}

// fittencode 生成
func setOutPutFileOne1(level logrus.Level, fileName string) {
	// 获取当前时间
	now := time.Now()
	// 获取日志文件路径
	logPath := fmt.Sprintf("%s/%s/%s.log", os.Getenv("LOG_PATH"), now.Format("2006-01-02"), fileName)
	// 判断日志文件是否存在，不存在则创建
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		os.MkdirAll(path.Dir(logPath), os.ModePerm)
		f, err := os.Create(logPath)
		if err!= nil {
			fmt.Println("create log file error:", err)
			return
		}
		defer f.Close()
	}
	// 设置日志输出
	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open log file error:", err)
		return
	}
	// 设置日志输出
	logrus.SetOutput(io.MultiWriter(file, os.Stdout))
	// 设置日志级别
	logrus.SetLevel(level)
}

// fittencode 生成
func RecoveryOne1(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			// 获取当前时间
			now := time.Now()
			// 获取日志文件路径
			logPath := fmt.Sprintf("%s/%s/panic.log", os.Getenv("LOG_PATH"), now.Format("2006-01-02"))
			// 判断日志文件是否存在，不存在则创建
			if _, err := os.Stat(logPath); os.IsNotExist(err) {
				os.MkdirAll(path.Dir(logPath), os.ModePerm)
				f, err := os.Create(logPath)
				if err != nil {
					fmt.Println("create log file error:", err)
					return
				}
				defer f.Close()
			}
			// 设置日志输出
			file, err := os.OpenFile(logPath, os.O_APPEND|os.O_WRONLY, 0666)
			if err != nil {
				fmt.Println("open log file error:", err)
				return
			}
			// 设置日志输出
			logrus.SetOutput(io.MultiWriter(file, os.Stdout))
			// 设置日志级别
			logrus.SetLevel(logrus.PanicLevel)
			// 记录panic信息
			logrus.WithFields(logrus.Fields{
				"err": err,
				"stack": string(debug.Stack()),
			}).Panic("panic occurred")
			// 响应客户端
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}()
	c.Next()
}



