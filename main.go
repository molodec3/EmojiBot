package main

import (
	"log"
	"strings"

	"gopkg.in/telegram-bot-api.v4"
)

type User struct {
	ID       int
	Username string
	ChatID   int
}

var Users map[int]*User
var Data map[string]string

func main() {
	var emoji = map[string]string{"а": "\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n",
		"б": "\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n",
		"в": "\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n",
		"г": "\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n",
		"д": "\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\n\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n",
		"е": "\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n",
		"ё": "\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n",
		"ж": "\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31A\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31A\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n",
		"з": "\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n",
		"и": "\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31A\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n",
		"й": "\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31A\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31A\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n",
		"к": "\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n",
		"л": "\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n",
		"м": "\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\n\U0001F31D\U0001F31A\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31D\n\U0001F31D\U0001F31A\U0001F31D\U0001F31A\U0001F31D\U0001F31A\U0001F31D\U0001F31A\U0001F31D\n\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\n\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\n\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\n\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n",
		"н": "\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n",
		"о": "\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n",
		"п": "\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n",
		"р": "\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n",
		"с": "\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n",
		"т": "\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n",
		"у": "\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n",
		"ф": "\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\n\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\n\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\n\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n",
		"х": "\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n",
		"ц": "\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n",
		"ч": "\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n",
		"ш": "\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\n\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\n\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\n\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\n\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\n\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\n\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n",
		"щ": "\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\n\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\n\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\n\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\n\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\n\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\n\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31A\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n",
		"ь": "\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n",
		"ъ": "\U0001F31D\U0001F31A\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n",
		"ы": "\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\n\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\n\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\n\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\n\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31A\U0001F31D\n\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31A\U0001F31D\n\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n",
		"э": "\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n",
		"ю": "\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31A\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\n\U0001F31D\U0001F31A\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\n\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\n\U0001F31D\U0001F31A\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\n\U0001F31D\U0001F31A\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\n\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n",
		"я": "\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31A\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\U0001F31A\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n",
		" ": "\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\U0001F31D\n"}

	Users = make(map[int]*User)

	bot, err := tgbotapi.NewBotAPI(/*"bot token"*/)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message == nil {
			continue
		}
		user, ok := Users[update.Message.From.ID]

		if !ok {
			log.Println("New user!!!")
			user = &User{ID: update.Message.From.ID, ChatID: int(update.Message.Chat.ID), Username: update.Message.From.UserName}
			Users[user.ID] = user
		}
		if prepare(update.Message.Text) == "/start" {
			text := `
			Show me some Russian words`
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
			bot.Send(msg)
			continue
		} else {
			var answer, message string
			answer = ""
			message = prepare(update.Message.Text)
			for i := 0; i < len([]rune(message)); i++ {
				val, ok := emoji[string([]rune(message)[i])]
				if ok {
					answer += val
				} else {
					answer = "Try something Russain (w/o punctuation)"
					break
				}
			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, answer)
			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
			}
		}
	}
}

func prepare(msg string) string {
	return strings.ToLower(msg)
}