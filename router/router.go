package router

import (
	_ "High-Performance-Online-Bookstore/docs" // docs is generated by Swag CLI, you have to import it.
	"High-Performance-Online-Bookstore/handler/book"
	"High-Performance-Online-Bookstore/handler/cart"
	"High-Performance-Online-Bookstore/handler/order"
	"High-Performance-Online-Bookstore/handler/state"
	"High-Performance-Online-Bookstore/handler/user/admin"
	"High-Performance-Online-Bookstore/handler/user/common"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/router/middleware"
	"errors"
	"github.com/gin-contrib/pprof"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swag "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go.uber.org/zap"
	"net/http"
	"time"
)

// InitRouter creates a gin router,
// load middlewares and start listening.
func InitRouter() {
	gin.SetMode(viper.GetString("level"))
	g := gin.New()
	Load(
		// Cores
		g,

		// External Middlewares
		ginzap.Ginzap(zap.L(), time.RFC3339, false),
		ginzap.RecoveryWithZap(zap.L(), true),
		middleware.RequestId(),
	)

	go testPing()

	startListen(g)
}

// testPing tests if the server is listening
// normally by ping the server address.
func testPing() {
	if err := pingServer(); err != nil {
		log.ErrNoResponse(err)
	}
	log.RouterDeployed()
}

// startListen starts listening the
// requests by http or https protocol.
func startListen(g *gin.Engine) {
	cert := viper.GetString("tls.cert")
	key := viper.GetString("tls.key")
	if cert != "" && key != "" {
		go func() {
			httpsAddr := viper.GetString("tls.addr")
			log.StartListenHTTPS(httpsAddr)
			if err := http.ListenAndServeTLS(httpsAddr, cert, key, g); err != nil {
				log.ErrListenHTTPS(err)
			}
		}()
	}
	httpAddr := viper.GetString("addr")
	log.StartListenHTTP(httpAddr)
	if err := http.ListenAndServe(httpAddr, g); err != nil {
		log.ErrListenHTTP(err)
	}
}

// pingServer pings the http server
// to make sure the server is working.
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/state/health`.
		resp, err := http.Get(viper.GetString("url") + "/state/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.WaitForRouter()
		time.Sleep(time.Second)
	}
	return errors.New("connect to the router failed")
}

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Internal Middlewares.
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(middleware.HasPermission)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "The api does not exist.",
		})
	})

	// swagger api docs
	g.GET("/swagger/*any", swag.WrapHandler(swaggerFiles.Handler))

	// pprof router
	pprof.Register(g)

	// The health check handlers
	st := g.Group("/state")
	{
		st.GET("/health", state.HealthCheck)
		st.GET("/disk", state.DiskCheck)
		st.GET("/cpu", state.CPUCheck)
		st.GET("/ram", state.RAMCheck)
	}

	// api version group
	v1 := g.Group("/v1")

	u := v1.Group("/user")
	{
		// user login router
		u.POST("/login", common.Login)
		// user register router
		u.POST("/register", common.Register)
	}

	// common user router group
	c := u.Group("/common")
	{
		// common user router
		c.GET("", common.SelfCheck)
		c.PUT("", common.SelfUpd)
		c.DELETE("", common.SelfDel)
	}
	// admin user router group
	a := u.Group("/admin")
	{
		// admin user router
		a.DELETE("/:id", admin.Delete)
		a.PUT("/:id", admin.Update)
		a.GET("", admin.List)
		a.GET("/:id", admin.Get)
		a.POST("/register", admin.Register)
	}
	// book manager router
	b := v1.Group("/book")
	{
		// book manager router
		b.POST("", book.Add)
		b.GET("", book.List)
		b.DELETE("/:id", book.Delete)
		b.PUT("/:id", book.Update)
		b.GET("/:id", book.Get)
	}
	ca := v1.Group("/cart")
	{
		// cart router
		ca.GET("", cart.Show)
		ca.PUT("", cart.Add)
		ca.DELETE("", cart.Delete)
		ca.DELETE("/all", cart.Clear)
	}
	o := v1.Group("/order")
	{
		// order router
		o.POST("", order.Create)
		o.PUT("", order.Deal)
		/*		o.GET("", order.List)
				o.GET("/:id", order.Get)
				o.PUT("/:id", order.Update)
				o.DELETE("/:id", order.Delete)*/
	}

	return g
}
