// Application wide logging package. Singleton Log enables logging trough the application.
package logger

import (
	"github.com/sirupsen/logrus"
)

var Log = logrus.New()