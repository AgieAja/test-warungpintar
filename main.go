package main

import (
	configDb "github.com/AgieAja/go-config-db/database/v2/mysql"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"tes-warungpintar/docs"
	"tes-warungpintar/routes/messageHttp"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"

	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog"
)

// !Important : Comments below are formatted as it is to be-
// read by Swagger tools (Swaggo)
// @title Service Tes Warung Pintar
// @version 1.0
// @description Service for tes warung pintar
// @BasePath /api/v1
// @contact.name Tiar Agisti
// @contact.email tiar.agisti@gmail.com
func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)

	err := godotenv.Load("config/.env")
	if err != nil {
		log.Error().Msg("Failed read configuration database")
		return
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Error().Msg("port cant empty")
		return
	}

	setMode := os.Getenv("SET_MODE")
	if setMode == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.New()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"POST", "DELETE", "GET", "OPTIONS", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "userid"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           720 * time.Hour,
	}))

	// Add a logger middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	r.Use(logger.SetLogger())

	// Custom logger
	subLog := zerolog.New(os.Stdout).With().Logger()

	var rxURL = regexp.MustCompile(`^/regexp\d*`)
	r.Use(logger.SetLogger(logger.Config{
		Logger:         &subLog,
		UTC:            true,
		SkipPath:       []string{"/skip"},
		SkipPathRegexp: rxURL,
	}))

	r.Use(gin.Recovery())

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	maxIdle := os.Getenv("MAX_IDLE")
	maxConn := os.Getenv("MAX_CONN")

	myMaxIdle, errMaxIdle := strconv.Atoi(maxIdle)
	if errMaxIdle != nil {
		log.Error().Msg("Failed convert errMaxIdle = " + errMaxIdle.Error())
		return
	}

	myMaxConn, errMaxConn := strconv.Atoi(maxConn)
	if errMaxConn != nil {
		log.Error().Msg("Failed convert errMaxConn = " + errMaxConn.Error())
		return
	}

	conn, errConn := getDBConnection(dbHost, dbPort, dbUser, dbPass, dbName, myMaxIdle, myMaxConn)
	if errConn != nil {
		log.Error().Msg("Failed connect database errConn : " + errConn.Error())
		return
	}

	defer func() {
		errConnClose := conn.Close()
		if errConnClose != nil {
			log.Error().Msg("errConnClose : " + errConnClose.Error())
		}
	}()

	log.Info().Msg("Last Update : " + time.Now().Format("2006-01-02 15:04:05"))
	log.Info().Msg("Service Running version 0.0.1 at port : " + port)

	//module message
	messageHttp.MessageRoute(r, conn)

	// use ginSwagger middleware to
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if errHTTP := http.ListenAndServe(":"+port, r); errHTTP != nil {
		log.Error().Msg(errHTTP.Error())
	}
}

func getDBConnection(dbHost, dbPort, dbUser, dbPass, dbName string, maxIdle, maxConn int) (*gorm.DB, error) {
	conn, err := configDb.ConnMySQLORM(dbHost, dbPort, dbUser, dbPass, dbName, maxIdle, maxConn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
