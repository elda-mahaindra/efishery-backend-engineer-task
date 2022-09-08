package api

import (
	"fmt"

	"auth-app/store"
	"auth-app/token"
	"auth-app/util"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

const (
	LOGIN_ROUTE        = "/login"
	PING_ROUTE         = "/ping"
	REGISTER_ROUTE     = "/register"
	VERIFY_TOKEN_ROUTE = "/verify-token"
)

type ApiServer struct {
	config       util.Config
	userStore    store.UserStore
	Router       *gin.Engine
	TokenManager token.Manager
}

func NewApiServer(config util.Config, userStore store.UserStore, tokenManager token.Manager) *ApiServer {
	apiServer := &ApiServer{
		config:       config,
		userStore:    userStore,
		TokenManager: tokenManager,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("role", validRole)
	}

	apiServer.setupRouter()

	return apiServer
}

func (server *ApiServer) authorize(serviceName string, role string) (bool, error) {
	super := util.SUPER
	admin := util.ADMIN
	user := util.USER

	services := map[string][]string{
		// ================= auth service =================
		"VerifyToken": {super, admin, user},
	}

	authorizedRoles, ok := services[serviceName]
	if !ok {
		return false, fmt.Errorf("service not registered")
	}

	isAuthorized := false

	for _, authorizedRole := range authorizedRoles {
		isAuthorized = isAuthorized || (authorizedRole == role)
	}

	return isAuthorized, nil
}

func (server *ApiServer) setupRouter() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	additionalAllowedHeaders := []string{"X-Amz-Date", "Authorization", "X-Api-Key", "X-Amz-Security-Token", "locale"}

	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, additionalAllowedHeaders...)

	router.Use(cors.New(config))

	router.GET(PING_ROUTE, server.ping)
	router.POST(REGISTER_ROUTE, server.register)
	router.POST(LOGIN_ROUTE, server.login)
	router.Use(authMiddleware(server.TokenManager)).GET(VERIFY_TOKEN_ROUTE, server.verifyToken)

	server.Router = router
}

func (server *ApiServer) Start(config util.Config) error {
	address := fmt.Sprintf(":%s", config.Port)

	return server.Router.Run(address)
}

func errorResponse(err error, message string) gin.H {
	return gin.H{
		"description": message,
		"error":       err.Error(),
	}
}
