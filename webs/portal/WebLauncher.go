package portal

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	controllers "iamcc.cn/doubanbookapi/webs/portal/controllers/v1"
	"iamcc.cn/doubanbookapi/webs/services/impl"
	"log"
	"strings"
)

/**
 * 变量定义
 */
const (
	DEFAULT_HOST = "0.0.0.0"
	DEFAULT_PORT = 6666
)

type WebLauncher struct {
	Host      string
	Port      int
	IsRunning bool

	Router *gin.Engine
}

/**
 * 构造函数
 */
func NewWebLauncherWithHostAndPort(host string, port int) (*WebLauncher, error) {
	// 定义
	var (
		result *WebLauncher
		err    error
	)

	// 参数
	if "" == strings.TrimSpace(host) {
		err := errors.New("Invalid argument value host.")
		return result, err
	}
	if port <= 0 || 65535 <= port {
		err := errors.New("Invalid argument value port.")
		return result, err
	}

	// 启动前检查
	err = impl.NewDefaultService().VerifyMongoDB()
	if nil != err {
		return nil, err
	}

	// 初始化
	result = &WebLauncher{
		Host: host,
		Port: port,
		// 注册一个默认的路由器
		Router: gin.Default(),
	}

	return result, err
}

func NewWebLauncher() (*WebLauncher, error) {
	return NewWebLauncherWithHostAndPort(DEFAULT_HOST, DEFAULT_PORT)
}

/**
 * 函数
 */
func (this *WebLauncher) Run() {

	log.Printf("Http listening on %s:%d", this.Host, this.Port)

	// Functions
	//this.Router.SetFuncMap(template.FuncMap{
	//	"FormatDatetime": functions.FormatDatetime,
	//	"SubString":      functions.SubString,
	//	"SubStr":         functions.SubStr,
	//})

	// Load Resources
	//this.Router.LoadHTMLGlob(filepath.Join(os.Getenv("GOPATH"), "src/iamcc.cn/smartreader/resources/templates/**/*"))

	// ---
	//this.Router.GET("/v1", func(c *gin.Context) {
	//})
	v1Group := this.Router.Group("/v1")
	{
		// ---
		v1Group.Any("/", func(c *gin.Context) {
			controllers.DefaultController_Version(c)
		})
		// ---
		buyRecordGroup := v1Group.Group("/buyrecord")
		{
			buyRecordGroup.GET("/", func(c *gin.Context) {
				controllers.BuyRecordController_Query(c)
			})
			buyRecordGroup.POST("/", func(c *gin.Context) {
				controllers.BuyRecordController_Add(c)
			})
		}
		// ---
		doubanGroup := v1Group.Group("/douban")
		{
			doubanGroup.GET("/*action", func(c *gin.Context) {
				controllers.BookApiController_Query(c)
			})
		}
		// ---
		bookGroup := v1Group.Group("/book")
		{
			bookGroup.GET("/*action", func(c *gin.Context) {
				controllers.BookController_ActionDispatcher(c)
			})
			bookGroup.POST("/", func(c *gin.Context) {
				log.Println("POST")
			})
		}
		// ---
	}
	// ---
	this.Router.Any("/favicon.ico", func(c *gin.Context) {
		controllers.DefaultController_Favicon(c)
	})
	this.Router.Any("/", func(c *gin.Context) {
		controllers.DefaultController_404Error(c)
	})
	// 50X Pages
	this.Router.Any("/50X", func(c *gin.Context) {
		controllers.DefaultController_50XError(c)
	})
	// 404 Page
	this.Router.Any("/404", func(c *gin.Context) {
		controllers.DefaultController_404Error(c)
	})
	this.Router.NoRoute(func(c *gin.Context) {
		controllers.DefaultController_404Error(c)
	})
	// ---
	this.Router.Run(fmt.Sprintf("%s:%d", this.Host, this.Port))

	log.Printf("Server exited ...")
}
