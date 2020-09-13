package utils

import (
	"context"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	elastic "github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v7"
)

//ILogger Interface for Logger
type ILogger interface {
	GetCurLevel() logrus.Level
	SetLevel(level logrus.Level)
	Level() logrus.Level
}

var defaultLogger *Logger

const (
	defaultServiceName = "Foxconn"
	formatterJSON      = "json"
	formatterText      = "text"
)

func init() {
	localIP, err := GetLocalIP()
	if err != nil {
		fmt.Println(err)
	}

	defaultLogger = &Logger{
		ServiceName: getServiceName(),
		IP:          localIP,
		logger: &logrus.Logger{
			Out:       os.Stderr,
			Formatter: newFormatter(),
			Hooks:     make(logrus.LevelHooks),
			Level:     getLevel(),
		},
	}
}

func getServiceName() string {
	serviceName := os.Getenv("LOGGER_SERVICENAME")
	if serviceName == "" {
		serviceName = defaultServiceName
	}
	return serviceName
}

func getLevel() logrus.Level {
	levelstr := os.Getenv("LOGGER_LEVEL")
	if levelstr == "" {
		return logrus.InfoLevel
	}
	level, err := strconv.Atoi(levelstr)
	if err != nil {
		fmt.Printf("get log level from env error:%v,level string:%s", err, levelstr)
		return logrus.InfoLevel
	}
	return logrus.Level(level)
}

func newFormatter() logrus.Formatter {
	formatter := os.Getenv("LOGGER_FORMATTER")
	if formatter == "" {
		formatter = formatterJSON
	}
	switch formatter {
	case formatterText:
		return new(logrus.TextFormatter)
	default:
		return new(logrus.JSONFormatter)
	}
}

//Logger 环境变量名LOGGER_FORMATTER，只能设置两个值：json、text
//TODO back log :watch level change
//TODO back log :static
//TODO review code: log
//TODO load config from envirement
type Logger struct {
	ServiceName string
	IP          string
	logger      *logrus.Logger
}

func (l *Logger) newField(ctx context.Context) logrus.Fields {
	fields := logrus.Fields{
		"serviceName": l.ServiceName,
		"ip":          l.IP,
	}

	if ctx != nil {
		//TODO get traceid
	}

	pc, filename, line, ok := runtime.Caller(2)
	if ok {
		funcname := runtime.FuncForPC(pc).Name()
		funcname = filepath.Ext(funcname)
		funcname = strings.TrimPrefix(funcname, ".")
		fields["func"] = funcname
		fields["file"] = filepath.Base(filename)
		fields["line"] = line
	}
	return fields
}

//SetServiceName 优先于环境变量
// 环境变量名LOGGER_SERVICENAME
func SetServiceName(serviceName string) {
	defaultLogger.ServiceName = serviceName
}

//SetLevel 优先于环境变量
// 环境变量名LOGGER_LEVEL,5=Debug 4=Info,3=Warn,2=Error,1=Fatal,0=Panic
func SetLevel(level logrus.Level) {
	defaultLogger.logger.SetLevel(level)
}

//GetCurLevel (Alias Name Of Level() func)
// return result from Level()
func GetCurLevel() logrus.Level {
	return Level()
}

//IsDebugLevel Returen is debug mode or not
func IsDebugLevel() bool {
	return IsLogLevel(5)
}

//IsLogLevel Returen is <level> mode or not
func IsLogLevel(level int) bool {
	return Level() == logrus.Level(level)
}

//Level return current log level
func Level() logrus.Level {
	return defaultLogger.logger.Level
}

//UseJSONFormatter 优先于环境变量
func UseJSONFormatter() {
	defaultLogger.logger.Formatter = new(logrus.JSONFormatter)
}

//UseTextFormatter 优先于环境变量
func UseTextFormatter() {
	defaultLogger.logger.Formatter = new(logrus.TextFormatter)
}

//AddESHook AddESHook
func AddESHook(level logrus.Level, urls ...string) {
	client, err := elastic.NewClient(elastic.SetURL(urls...))
	if err != nil {
		defaultLogger.logger.Panic(nil, err)
	}

	hook, err := elogrus.NewAsyncElasticHook(client, defaultLogger.ServiceName, level, "huspluslog")
	if err != nil {
		defaultLogger.logger.Panic(nil, err)
	}
	defaultLogger.logger.Hooks.Add(hook)
}

//Debugf show log with format
func Debugf(ctx context.Context, format string, args ...interface{}) {
	defaultLogger.logger.WithFields(defaultLogger.newField(ctx)).Debugf(format, args...)
}

//Infof show log with format
func Infof(ctx context.Context, format string, args ...interface{}) {
	defaultLogger.logger.WithFields(defaultLogger.newField(ctx)).Infof(format, args...)
}

//Printf show log with format
func Printf(ctx context.Context, format string, args ...interface{}) {
	defaultLogger.logger.WithFields(defaultLogger.newField(ctx)).Printf(format, args...)
}

//Warnf show log with format
func Warnf(ctx context.Context, format string, args ...interface{}) {
	defaultLogger.logger.WithFields(defaultLogger.newField(ctx)).Warnf(format, args...)
}

//Warningf show log with format
func Warningf(ctx context.Context, format string, args ...interface{}) {
	defaultLogger.logger.WithFields(defaultLogger.newField(ctx)).Warningf(format, args...)
}

//Errorf show log with format
func Errorf(ctx context.Context, format string, args ...interface{}) {
	defaultLogger.logger.WithFields(defaultLogger.newField(ctx)).Errorf(format, args...)
}

//Fatalf show log with format
func Fatalf(ctx context.Context, format string, args ...interface{}) {
	defaultLogger.logger.WithFields(defaultLogger.newField(ctx)).Fatalf(format, args...)
}

//Panicf show log with format
func Panicf(ctx context.Context, format string, args ...interface{}) {
	defaultLogger.logger.WithFields(defaultLogger.newField(ctx)).Panicf(format, args...)
}

//Debug show log without format
func Debug(ctx context.Context, args ...interface{}) {
	defaultLogger.logger.WithFields(defaultLogger.newField(ctx)).Debug(args...)
}

//Info show log without format
func Info(ctx context.Context, args ...interface{}) {
	defaultLogger.logger.WithFields(defaultLogger.newField(ctx)).Info(args...)
}

//Print show log without format
func Print(ctx context.Context, args ...interface{}) {
	defaultLogger.logger.WithFields(defaultLogger.newField(ctx)).Print(args...)
}

//Warn show log without format
func Warn(ctx context.Context, args ...interface{}) {
	defaultLogger.logger.WithFields(defaultLogger.newField(ctx)).Warn(args...)
}

//Warning show log without format
func Warning(ctx context.Context, args ...interface{}) {
	defaultLogger.logger.WithFields(defaultLogger.newField(ctx)).Warning(args...)
}

//Error show log without format
func Error(ctx context.Context, args ...interface{}) {
	defaultLogger.logger.WithFields(defaultLogger.newField(ctx)).Error(args...)
}

//Fatal show log without format
func Fatal(ctx context.Context, args ...interface{}) {
	defaultLogger.logger.WithFields(defaultLogger.newField(ctx)).Fatal(args...)
}

//Panic show log without format
func Panic(ctx context.Context, args ...interface{}) {
	defaultLogger.logger.WithFields(defaultLogger.newField(ctx)).Panic(args...)
}

//Debugln show log without format with line(CLCF)
func Debugln(ctx context.Context, args ...interface{}) {
	defaultLogger.logger.WithFields(defaultLogger.newField(ctx)).Debugln(args...)
}

//Infoln show log without format with line(CLCF)
func Infoln(ctx context.Context, args ...interface{}) {
	defaultLogger.logger.WithFields(defaultLogger.newField(ctx)).Infoln(args...)
}

//Println show log without format with line(CLCF)
func Println(ctx context.Context, args ...interface{}) {
	defaultLogger.logger.WithFields(defaultLogger.newField(ctx)).Println(args...)
}

//Warnln show log without format with line(CLCF)
func Warnln(ctx context.Context, args ...interface{}) {
	defaultLogger.logger.WithFields(defaultLogger.newField(ctx)).Warnln(args...)
}

//Warningln show log without format with line(CLCF)
func Warningln(ctx context.Context, args ...interface{}) {
	defaultLogger.logger.WithFields(defaultLogger.newField(ctx)).Warningln(args...)
}

//Errorln show log without format with line(CLCF)
func Errorln(ctx context.Context, args ...interface{}) {
	defaultLogger.logger.WithFields(defaultLogger.newField(ctx)).Errorln(args...)
}

//Fatalln show log without format with line(CLCF)
func Fatalln(ctx context.Context, args ...interface{}) {
	defaultLogger.logger.WithFields(defaultLogger.newField(ctx)).Fatalln(args...)
}

//Panicln show log without format with line(CLCF)
func Panicln(ctx context.Context, args ...interface{}) {
	defaultLogger.logger.WithFields(defaultLogger.newField(ctx)).Panicln(args...)
}

//GetLocalIP get local private ip
func GetLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return "", err
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			//TODO IP6?
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", nil
}
