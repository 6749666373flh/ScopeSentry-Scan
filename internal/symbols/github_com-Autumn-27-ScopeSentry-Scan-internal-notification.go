// Code generated by 'yaegi extract github.com/Autumn-27/ScopeSentry-Scan/internal/notification'. DO NOT EDIT.

package symbols

import (
	"github.com/Autumn-27/ScopeSentry-Scan/internal/notification"
	"reflect"
)

func init() {
	Symbols["github.com/Autumn-27/ScopeSentry-Scan/internal/notification/notification"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"FlushBuffer":            reflect.ValueOf(notification.FlushBuffer),
		"InitializeNotification": reflect.ValueOf(notification.InitializeNotification),
		"NotificationQueues":     reflect.ValueOf(&notification.NotificationQueues).Elem(),

		// type definitions
		"NotificationQueue": reflect.ValueOf((*notification.NotificationQueue)(nil)),
	}
}
