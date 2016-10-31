package managers

import (
	"time"
)

//GetDateTimeNow returns the datetime now
func GetDateTimeNow() string {
	return time.Now().Format(time.UnixDate)
}

//SayHello returns a greeting
func SayHello(name string) string {
	return "Hello " + name
}
