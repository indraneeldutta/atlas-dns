package apis

import (
	dnsservice "github.com/atlas-dns/services/dns_service"
	mongoservice "github.com/atlas-dns/services/mongo_service"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitialiseApplication() {
	if viper.GetString("ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	securityConfig := secure.DefaultConfig()
	securityConfig.SSLRedirect = false
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(secure.New(securityConfig))

	router.Use(gin.Recovery())

	mongoClient := mongoservice.NewMongoConnection()

	dnsService := dnsservice.NewDnsService(mongoClient)

	v1 := router.Group("/v1")

	NewDnsController(v1, &dnsService)

	router.Run(viper.GetString("SERVER_PORT"))
}

// SetupEnvironment sets up the configs and environment for the application to start
func SetupEnvironment() error {
	err := viper.BindEnv("gopath")
	if err != nil {
		panic(err)
	}

	viper.SetConfigName("config")
	viper.AddConfigPath("./config/")
	viper.SetConfigType("yml")

	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return err
}
