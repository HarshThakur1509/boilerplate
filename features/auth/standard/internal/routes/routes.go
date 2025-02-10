	// USER AUTHENTICATION

	router.HandleFunc("POST /login", handlers.CustomLogin)
	router.HandleFunc("POST /register", handlers.CustomRegister)

	router.HandleFunc("POST /reset", handlers.ResetPasswordHandler)
	router.HandleFunc("POST /forgot", handlers.ForgotPasswordHandler)

	router.HandleFunc("GET /auth", gothic.BeginAuthHandler)
	router.HandleFunc("GET /auth/callback", handlers.GoogleCallbackHandler)

	router.HandleFunc("GET /cookie", handlers.GetCookie)

	// Add code here

	authRouter := http.NewServeMux()
	authRouter.HandleFunc("GET /auth/logout", handlers.GothLogout)
	authRouter.HandleFunc("GET /api/validate", handlers.Validate)

	router.Handle("/", middleware.AuthMiddleware(authRouter))