package main

import (
	"booking-cinema-backend/api/handler"
	"booking-cinema-backend/api/middlewares"
	"booking-cinema-backend/api/routes"
	"booking-cinema-backend/repository"
	roleservice "booking-cinema-backend/services/role"
	"booking-cinema-backend/services/user"
	"booking-cinema-backend/services/user/seeds"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type EnvConfig struct {
	DBSource string `mapstructure:"DB_SOURCE"`
	Port     string `mapstructure:"PORT"`
}

var USER_SEEDS_AMOUNT = 100 
var SEEDS_GOROUTINE = 10 //dont add too much goroutine

func main() {
	config, err := loadEnv(".")
	if err != nil {
		log.Fatal("cannot load config")
	}
	repository.ConnectDB(config.DBSource)
	db := repository.GetDB()
	runDBMigration(db)
	addSeedData(db)
	r := SetupRouter(db)
	_ = r.Run(":8080")
}

func loadEnv(path string) (config EnvConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func runDBMigration(db *gorm.DB) {
	err := db.AutoMigrate(&userservice.User{}, &roleservice.Role{})
	if err != nil {
		log.Fatal("fail to run migrate up: ", err)
	}
	err = db.Migrator().CreateIndex(&userservice.User{}, "Username")
	if err != nil {
		log.Fatal("fail to create username index: ", err)
	}
	err = db.Migrator().CreateIndex(&roleservice.Role{}, "RoleName")
	if err != nil {
		log.Fatal("fail to create rolename index: ", err)
	}

	log.Println("db migrate successfully")
}

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Cors())
	r.Use(middlewares.JSONAppErrorReporter())

	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	// roleRepo := repository.NewRoleRepo(db)
	// roleService := services.NewRoleService(roleRepo)
	// rolehandler := handler.NewRoleHandler(roleService)
	// roleRoutes := routes.NewRoleRoutes(rolehandler)
	// roleRoutes.Setup(r)

	userRepo := repository.NewUserRepo(db)
	userService := userservice.NewUserService(userRepo)
	userhandler := handler.NewUserHandler(userService)
	userRoutes := routes.NewUserRoutes(userhandler)
	userRoutes.Setup(r)

	return r
}

func addSeedData(db *gorm.DB) {
	addSeedUser(db)
}

func addSeedUser(db *gorm.DB) {
	var wg sync.WaitGroup
	userRepo := repository.NewUserRepo(db)
	userService := userservice.NewUserService(userRepo)
	err := userseeds.CreateAdminUser(userService)
	if err != nil {
		log.Println("error when trying to add super admin user: ", err)
	}
	log.Println("add super admin user successfull")
	wg.Add(SEEDS_GOROUTINE)
	for i := 0; i < SEEDS_GOROUTINE; i++ {
		go func() {
			for i:= 0; i < USER_SEEDS_AMOUNT / SEEDS_GOROUTINE; i++ {
				err := userseeds.CreateUser(userService)
				if err != nil {
					log.Println("error when trying to add user seed data: ", err)
				}
			}
			defer wg.Done()
		}()
	}
	wg.Wait()
	log.Printf("add %v seed users successfull \n", USER_SEEDS_AMOUNT)
	

}
