package app

import (
	"html/template"
)

var Templates *template.Template

const MainLayout string = "strn/templates/main.html"
const RedisAddr string = "127.0.0.1:6379"
