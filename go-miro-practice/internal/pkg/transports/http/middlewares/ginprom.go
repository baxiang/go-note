package middlewares

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"sync"
	"time"
	"github.com/prometheus/client_golang/prometheus"
)
const (
	metricsPath = "/metrics"
	faviconPath = "/favicon.ico"
)
var (
	httpHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "http_server",
		Name:      "requests_seconds",
		Help:      "Histogram of response latency (seconds) of http handlers.",
	}, []string{"method", "code", "uri"})
)
func init() {
	prometheus.MustRegister(httpHistogram)
}
type handlerPath struct {
	sync.Map
}

func(hp *handlerPath)get(handler string)string{
	if v,ok:=hp.Load(handler);ok{
		return v.(string)
	}
	return ""
}
func(hp *handlerPath)set(ginInfo gin.RouteInfo){
	hp.Store(ginInfo.Handler,ginInfo.Path)
}
type GinPrometheus struct {
	engine *gin.Engine
	ignored map[string]bool
	pathMap *handlerPath
	updated bool
}
type Option func(*GinPrometheus)

func Ignore(path ...string)Option{
	return func(gp *GinPrometheus) {
		for _,p :=range path{
			gp.ignored[p] = true
		}
	}
}

func NewGinPrometheus(e *gin.Engine,options ...Option)*GinPrometheus{
	if e == nil {
		return nil
	}
	gp := &GinPrometheus{
		engine: e,
		ignored: map[string]bool{
			metricsPath: true,
			faviconPath: true,
		},
		pathMap: &handlerPath{},
	}
	for _, o := range options {
		o(gp)
	}
	return gp
}
func (gp *GinPrometheus) updatePath() {
	gp.updated = true
	for _, ri := range gp.engine.Routes() {
		gp.pathMap.set(ri)
	}
}

// Middleware 返回中间件
func (gp *GinPrometheus) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !gp.updated {
			gp.updatePath()
		}
		// 把不需要的过滤掉
		if gp.ignored[c.Request.URL.String()] == true {
			c.Next()
			return
		}
		start := time.Now()

		c.Next()

		httpHistogram.WithLabelValues(
			c.Request.Method,
			strconv.Itoa(c.Writer.Status()),
			gp.pathMap.get(c.HandlerName()),
		).Observe(time.Since(start).Seconds())
	}
}

