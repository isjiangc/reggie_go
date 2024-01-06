package server

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	apiV1 "reggie_go/api/v1"
	"reggie_go/docs"
	"reggie_go/internal/handler"
	"reggie_go/internal/middleware"
	"reggie_go/pkg/jwt"
	"reggie_go/pkg/log"
	"reggie_go/pkg/server/http"
)

func NewHTTPServer(
	logger *log.Logger,
	conf *viper.Viper,
	jwt *jwt.JWT,
	userHandler *handler.UserHandler,
	employeeHandler *handler.EmployeeHandler,
	categoryHandler *handler.CategoryHandler,
	dishHandler *handler.DishHandler,
	setmealHandler *handler.SetmealHandler,
	addressbookHandler *handler.AddressbookHandler,
	usersHandler *handler.UsersHandler,
) *http.Server {
	gin.SetMode(gin.DebugMode)
	s := http.NewServer(
		gin.Default(),
		logger,
		http.WithServerHost(conf.GetString("http.host")),
		http.WithServerPort(conf.GetInt("http.port")),
	)
	store := cookie.NewStore([]byte("reggie_go_cookie"))
	// swagger doc
	docs.SwaggerInfo.BasePath = "/"
	s.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		//ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", conf.GetInt("app.http.port"))),
		ginSwagger.DefaultModelsExpandDepth(-1),
	))

	s.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
		//middleware.SignMiddleware(log),
		// 使用session
		sessions.Sessions("reggie_go", store),
	)
	s.GET("/", func(ctx *gin.Context) {
		logger.WithContext(ctx).Info("hello")
		apiV1.HandleSuccess(ctx, map[string]interface{}{
			":)": "Thank you for using nunu!",
		})
	})
	// 员工
	s.POST("/employee", employeeHandler.Save)
	s.PUT("/employee", employeeHandler.UpdateEmployee)
	emp := s.Group("/employee")
	{
		emp.POST("/login", employeeHandler.Login)
		emp.POST("/logout", employeeHandler.Logout)
		emp.GET("/page", employeeHandler.GetEmployeeList)
		emp.GET("/:id", employeeHandler.GetEmployeeById)
	}
	// 分类
	s.POST("/category", categoryHandler.CreateCategory)
	s.DELETE("/category", categoryHandler.DeleteCategory)
	s.PUT("/category", categoryHandler.UpdateCategory)
	cate := s.Group("/category")
	{
		cate.GET("/page", categoryHandler.GetCategoryList)
	}

	s.POST("/dish", dishHandler.CreateDishWithFlavor)
	s.GET("/dish/page", dishHandler.GetDishList)
	s.GET("/dish/:id", dishHandler.GetDishById)

	s.DELETE("/setmeal", setmealHandler.DeleteSetmeal)
	setmeal := s.Group("/setmeal")
	{
		setmeal.GET("/page", setmealHandler.GetSetmealList)
		setmeal.POST("/:status/", setmealHandler.UpdateSetmealStatus)
	}

	s.POST("/addressBook", addressbookHandler.SaveAddressBook)
	addressBook := s.Group("/addressBook")
	{
		addressBook.GET("/list/:userid", addressbookHandler.GetAddressbookByUserId)
		addressBook.PUT("/default", addressbookHandler.UpdateAddressIsDefault)
		addressBook.GET("/:id", addressbookHandler.GetAddressBookById)
		addressBook.GET("/default/:userid", addressbookHandler.GetDefaultAddressBook)

	}
	s.POST("/users/login", usersHandler.UsersLogin)
	s.POST("/users/sendMsg", usersHandler.SendMsg)

	v1 := s.Group("/v1")
	{
		// No route group has permission
		noAuthRouter := v1.Group("/")
		{
			noAuthRouter.POST("/register", userHandler.Register)
			noAuthRouter.POST("/login", userHandler.Login)
		}
		// Non-strict permission routing group
		noStrictAuthRouter := v1.Group("/").Use(middleware.NoStrictAuth(jwt, logger))
		{
			noStrictAuthRouter.GET("/user", userHandler.GetProfile)
		}

		// Strict permission routing group
		strictAuthRouter := v1.Group("/").Use(middleware.StrictAuth(jwt, logger))
		{
			strictAuthRouter.PUT("/user", userHandler.UpdateProfile)
		}
	}

	return s
}
