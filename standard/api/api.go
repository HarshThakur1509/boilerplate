package api

import (
	"fmt"
	"net/http"

	"github.com/HarshThakur1509/boilerplate/standard/controllers"
	"github.com/HarshThakur1509/boilerplate/standard/middleware"
	"github.com/markbates/goth/gothic"
	"github.com/rs/cors"
)

type ApiServer struct {
	addr string
}

func NewApiServer(addr string) *ApiServer {
	return &ApiServer{addr: addr}
}

func (s *ApiServer) Run() error {
	router := http.NewServeMux()

	router.HandleFunc("GET /health", controllers.Health)

	// USER AUTHENTICATION

	router.HandleFunc("POST /login", controllers.CustomLogin)
	router.HandleFunc("POST /register", controllers.CustomRegister)

	router.HandleFunc("GET /auth", gothic.BeginAuthHandler)
	router.HandleFunc("GET /auth/callback", controllers.GoogleCallbackHandler)

	router.HandleFunc("GET /cookie", controllers.GetCookie)

	// Add code here

	authRouter := http.NewServeMux()
	authRouter.HandleFunc("GET /auth/logout", controllers.GothLogout)
	authRouter.HandleFunc("GET /api/user", controllers.GetUser)

	router.Handle("/", middleware.AuthMiddleware(authRouter))

	stack := middleware.MiddlewareChain(middleware.Logger, middleware.RecoveryMiddleware)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Specify your React frontend origin
		AllowCredentials: true,                              // Allow cookies and credentials
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{
			"Content-Type",
			"Authorization",
			"Accept",
			"Origin",
			"X-Requested-With"},
	}).Handler(stack(router))

	server := http.Server{
		Addr:    s.addr,
		Handler: corsHandler,
	}
	fmt.Println("Server has started", s.addr)
	return server.ListenAndServe()
}
