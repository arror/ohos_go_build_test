package log

/*
#include <stdlib.h>

typedef enum {
    LOG_APP = 0,
} LogType;

typedef enum {
    LOG_DEBUG = 3,
    LOG_INFO = 4,
    LOG_WARN = 5,
    LOG_ERROR = 6,
    LOG_FATAL = 7,
} LogLevel;

static void OHOS_LOG(LogType type, LogLevel level, unsigned int domain, const char *tag, const char *message) {
    extern int OH_LOG_Print(LogType type, LogLevel level, unsigned int domain, const char *tag, const char *fmt, ...);
    OH_LOG_Print(type, level, domain, tag, "%{public}s", message);
}
*/
import "C"
import "unsafe"

var osLog = func(level C.LogLevel, message string) {
	temp := C.CString(message)
	tag := C.CString("GOLANG")
	defer C.free(unsafe.Pointer(temp))
	defer C.free(unsafe.Pointer(tag))
	C.OHOS_LOG(C.LOG_APP, level, 0x0F, tag, temp)
}

func Debug(message string) {
	osLog(C.LOG_DEBUG, message)
}

func Info(message string) {
	osLog(C.LOG_INFO, message)
}

func Wran(message string) {
	osLog(C.LOG_WARN, message)
}

func Error(message string) {
	osLog(C.LOG_ERROR, message)
}

func Fatal(message string) {
	osLog(C.LOG_FATAL, message)
}
