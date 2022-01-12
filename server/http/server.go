package http

import (
	"github.com/gin-gonic/gin"
	"github.com/penk110/distribute_cache/cache/impl"
	"github.com/penk110/distribute_cache/cluster"
	"net"
)

type Server struct {
	impl.Cache
	cluster.Cluster
}

var (
	_engine *gin.Engine
)

func NewServer(cache impl.Cache, cluster cluster.Cluster) *Server {
	var s *Server

	s = &Server{
		Cache:   cache,
		Cluster: cluster,
	}
	return s
}

func (s *Server) Listen(listener net.Listener) {
	var err error
	_engine = gin.Default()

	v1 := _engine.Group("/api/v1")

	cacheGp := v1.Group("/cache")
	regCacheRouter(cacheGp)

	//statusGp := v1.Group("/status")
	//regStatusRouter(statusGp)
	//
	//clusterGp := v1.Group("/cluster")
	//regStatusRouter(clusterGp)
	//
	//rebalancedGp := v1.Group("/rebalanced")
	//regRebalancedRouter(rebalancedGp)

	if err = _engine.RunListener(listener); err != nil {
		panic("try start engine failed, err: " + err.Error())
	}
}

func regCacheRouter(rg *gin.RouterGroup) {
	rg.GET("/:key", GetCacheRouter().Retrieve)
	rg.POST("", GetCacheRouter().Create)
	rg.PUT("/:key", GetCacheRouter().Create)
	rg.DELETE("/:key", GetCacheRouter().Create)
}

func regStatusRouter(rg *gin.RouterGroup) {
	rg.POST("", GetCacheRouter().Create)
	rg.GET("", GetCacheRouter().Create)
	rg.PUT("/:key", GetCacheRouter().Create)
	rg.DELETE("/:key", GetCacheRouter().Create)
}

func regClusterRouter(rg *gin.RouterGroup) {
	rg.POST("", GetCacheRouter().Create)
	rg.GET("", GetCacheRouter().Create)
	rg.PUT("/:key", GetCacheRouter().Create)
	rg.DELETE("/:key", GetCacheRouter().Create)
}

func regRebalancedRouter(rg *gin.RouterGroup) {
	rg.POST("", GetCacheRouter().Create)
	rg.GET("", GetCacheRouter().Create)
	rg.PUT("/:key", GetCacheRouter().Create)
	rg.DELETE("/:key", GetCacheRouter().Create)
}
