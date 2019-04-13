package config

import (
	"os"
	"time"
)

var Token = os.Getenv("TOKEN")
var Prefix = "g!"
var StartTime = time.Now()
