package main

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)


func main() {
	logrus.Println("Hello ,go module mode\n")
	logrus.Println(uuid.NewString())
}