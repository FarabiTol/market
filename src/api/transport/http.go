package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Router struct {
	log    *logrus.Entry
	router *gin.Engine
	secret []byte
}
