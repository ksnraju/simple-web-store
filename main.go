package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/ksnraju/simple-web-store/entity"
	"github.com/ksnraju/simple-web-store/handler"
	"github.com/ksnraju/simple-web-store/handler/cart"
	"github.com/ksnraju/simple-web-store/handler/item"
	"github.com/ksnraju/simple-web-store/handler/order"
	"github.com/ksnraju/simple-web-store/handler/user"
	"github.com/ksnraju/simple-web-store/middleware"
	"github.com/ksnraju/simple-web-store/validators"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DB() (db *gorm.DB) {
	db, err := gorm.Open(postgres.Open("dbname=simple-web-store user=admin password=admin@123"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return
}

func main() {
	db := DB()
	db.AutoMigrate(
		&entity.User{},
		&entity.Cart{},
		&entity.Item{},
		&entity.Order{},
		&entity.CartItem{},
		&entity.OrderItem{},
	)
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"*"},
	}))

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("alphaspace", validators.AlphaSpace)
		v.RegisterValidation("alphanumspace", validators.AlphaNumSpace)
		v.RegisterValidation("username", validators.Username)
		v.RegisterValidation("password", validators.Password)
	}

	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/", handler.Index())
	r.Use(middleware.ExtractToken)
	r.GET("/user/list", user.ListUsers(db))
	r.GET("/item/list", item.ListItems(db))
	r.GET("/cart/list", cart.ListCarts(db))
	r.GET("/order/list", order.ListOrders(db))
	r.POST("/user/create", user.CreateUser(db))
	r.POST("/user/login", user.UserLogin(db))
	r.POST("/item/create", item.CreateItem(db))
	r.POST("/cart/add", cart.AddToCart(db))
	r.POST("/cart/:CartID/complete", order.CreateOrder(db))

	r.NoMethod(handler.AppError(400, "No Method Found"))
	r.NoRoute(handler.AppError(404, "No Route Found"))
	r.Run()
}
