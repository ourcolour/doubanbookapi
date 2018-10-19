package portal

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"iamcc.cn/doubanbookapi/webs/portal/controllers"
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
	this.Router.GET("/version", func(c *gin.Context) {
		controllers.DefaultController_Version(c)
	})
	doubanGroup := this.Router.Group("/douban")
	{
		doubanGroup.GET("/isbn/:isbn", func(c *gin.Context) {
			controllers.BookApiController_QueryByIsbn(c)
		})
		doubanGroup.GET("/id/:id", func(c *gin.Context) {
			controllers.BookApiController_QueryById(c)
		})
	}
	// ---
	bookGroup := this.Router.Group("/book")
	{
		bookGroup.GET("/id/:id", func(c *gin.Context) {
			controllers.BookController_QueryById(c)
		})
		bookGroup.GET("/:isbn", func(c *gin.Context) {
			controllers.BookController_QueryByIsbn(c)
		})
		bookGroup.POST("/", func(c *gin.Context) {
			log.Println("POST")
		})
	}
	// ---
	this.Router.Run(fmt.Sprintf("%s:%d", this.Host, this.Port))

	log.Printf("Server exited ...")
}
