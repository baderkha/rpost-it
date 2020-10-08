package main

import (
	"comment-me/src/dependency"
	"comment-me/src/repository"
	"encoding/json"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// Models our json input
type ServerEnv struct {
	DSN          string `json:"dsn"`
	HashStrength uint   `json:"hashStrength"`
	JWTTokenTime int64  `json:"jwtTokenTimeMinutes"`
	JWTIssuer    string `json:"jwtIssuer"`
	JWTSecret    string `json:"jwtSecret"`
}

// reads the enviroment file if it exists , if not then there's a problem
func readEnviromentFile() *ServerEnv {
	var serverEnv ServerEnv
	jsonFile, err := os.Open("env.json")
	if err != nil {
		panic(":0 PANIK , enviroment json not found")
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &serverEnv)
	if err != nil {
		panic(":0 PANIK , enviroment json could not decode")
	}
	return &serverEnv
}

// fetch the db connection here
func getDB(serverEnv *ServerEnv) *gorm.DB {
	db, err := gorm.Open(mysql.Open(serverEnv.DSN), &gorm.Config{})
	if err != nil {
		panic(":0 PANIk , DB CONN COULD NOT BE MADE")
	}
	db.Logger = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			LogLevel: logger.Info, // Log level
			Colorful: true,        // Disable color
		},
	)
	return db
}

func migrate(db *gorm.DB) {
	var acc repository.Account
	_ = db.AutoMigrate(&acc)

	var user repository.User
	_ = db.AutoMigrate(&user)
}

func makeCors(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://shrter.xyz", "http://127.0.0.1:5500", "http://localhost:8080"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "DELETE", "GET"},
		AllowHeaders:     []string{"Origin", "content-type", "authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}

// Define the gin routes in here using the router
func makeRestRoutes(router *gin.Engine, controller *dependency.Dependency) {
	api := router.Group("api")
	{
		account := api.Group("accounts")
		{
			// create account
			account.POST("", controller.POSTAccount)
			// generate a jwt for account if valid
			account.POST("jwt", controller.POSTAccountJWT)
		}
	}
}

func makeRouter(router *gin.Engine, controller *dependency.Dependency) {
	makeCors(router)
	makeRestRoutes(router, controller)
}

// no buisness logic lives here
func main() {
	env := readEnviromentFile()
	db := getDB(env)
	migrate(db)
	controller := dependency.MakeDependencies(db, &dependency.ServerEnvDependency{
		HashStrength: env.HashStrength,
		JWTIssuer:    env.JWTIssuer,
		JWTSecret:    env.JWTSecret,
		JWTTokenTime: env.JWTTokenTime,
	})
	router := gin.Default()
	makeRouter(router, controller)
	_ = router.Run()
}
