package router

import (
	"Jinshuzhai-Bookstore/handler/user"
	"errors"
	"github.com/spf13/viper"
	swag "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"net/http"
	"time"

	_ "Jinshuzhai-Bookstore/docs" // docs is generated by Swag CLI, you have to import it.
	"Jinshuzhai-Bookstore/handler/state"
	"Jinshuzhai-Bookstore/router/middleware"

	"github.com/gin-contrib/pprof"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter(logger *zap.Logger) {
	gin.SetMode(viper.GetString("mode"))
	g := gin.New()
	Load(
		// Cores
		g,
		/*		middleware.GinLogger(logger),
				middleware.GinRecovery(logger, true),*/
		// Middlewares
		ginzap.Ginzap(logger, time.RFC3339, true),
		ginzap.RecoveryWithZap(logger, true),
		middleware.RequestId(),
	)

	go testPing(logger)

	startListen(g, logger)
}

func testPing(logger *zap.Logger) {
	if err := pingServer(logger); err != nil {
		logger.Fatal("The router has no response, or it might took too long to start up.", zap.Error(err))
	}
	logger.Info("The router has been deployed successfully.")
}

func startListen(g *gin.Engine, logger *zap.Logger) {
	cert := viper.GetString("tls.cert")
	key := viper.GetString("tls.key")
	if cert != "" && key != "" {
		go func() {
			logger.Info("Start to listening the incoming requests on https address", zap.String("tls.addr", viper.GetString("tls.addr")))
			logger.Info(http.ListenAndServeTLS(viper.GetString("tls.addr"), cert, key, g).Error())
		}()
	}
	logger.Info("Start to listening the incoming requests on http address", zap.String("addr", viper.GetString("addr")))
	logger.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

// pingServer pings the http server to make sure the router is working.
func pingServer(logger *zap.Logger) error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/state/health`.
		resp, err := http.Get(viper.GetString("url") + "/state/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		logger.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("connect to the router failed")
}

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The requested resource is not found.")
	})

	// swagger api docs
	g.GET("/swagger/*any", swag.WrapHandler(swaggerFiles.Handler))

	// pprof router
	pprof.Register(g)

	// api for authentication functionalities
	g.POST("/login", user.Login)

	// The user handlers, requiring authentication
	u := g.Group("/v1/user")
	// use authentication middleware
	//u.Use(middleware.AuthMiddleware())
	{
		u.POST("", user.Create)
		u.DELETE("/:id", user.Delete)
		u.PUT("/:id", user.Update)
		u.GET("", user.List)
		u.GET("/:username", user.Get)
	}

	// The health check handlers
	st := g.Group("/state")
	{
		st.GET("/health", state.HealthCheck)
		st.GET("/disk", state.DiskCheck)
		st.GET("/cpu", state.CPUCheck)
		st.GET("/ram", state.RAMCheck)
	}

	return g
}