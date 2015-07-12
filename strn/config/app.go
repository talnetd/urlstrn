package app

import (
	"html/template"
	"os"
)

var Templates *template.Template

const MainLayout string = "strn/templates/main.html"

// var RedisAddress string = "127.0.0.1"
// var RedisPort string = "6379"

var RedisAddress string = os.Getenv("REDIS_PORT_6379_TCP_ADDR")
var RedisPort string = os.Getenv("REDIS_PORT_6379_TCP_PORT")
