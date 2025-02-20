package transport

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"market/src/service"
	"net/http"
)

type Router struct {
	log    *logrus.Entry
	router *gin.Engine
	app    service.MarketService
}

func NewRouter(log *logrus.Logger, app service.MarketService) *Router {
	r := &Router{
		log:    log.WithField("component", "router"),
		router: gin.Default(),
		app:    app,
	}
	r.router.GET("/metrics", prometheusHandler())
	g := r.router.Group("/api/v1")
	g.GET("/wallet/:id", r.getCollection)
	g.POST("/wallet", r.addCollection)
	return r
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func (r *Router) Run(_ context.Context, addr string) error {
	return r.router.Run(addr)
}

func (r *Router) addCollection(c *gin.Context) {
	var input *CreateCollectionReq
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	id, err := r.app.CreateCollection(c, input)
	if err != nil {
		r.log.Errorf("failed to store date: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (r *Router) getCollection(c *gin.Context) {
	val := c.Param("id")
	id, err := uuid.Parse(val)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	w, err := r.app.GetCollection(c, id)
	switch {
	case err == nil:
	case errors.Is(err, pgx.ErrNoRows):
		c.JSON(http.StatusNotFound, err)
		return
	default:
		r.log.Errorf("failed to get collection: %v", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, w)
}
