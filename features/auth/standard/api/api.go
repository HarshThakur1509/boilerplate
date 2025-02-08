	// USER AUTHENTICATION

	router.HandleFunc("POST /login", controllers.CustomLogin)
	router.HandleFunc("POST /register", controllers.CustomRegister)

	router.HandleFunc("POST /reset", controllers.ResetPasswordHandler)
	router.HandleFunc("POST /forgot", controllers.ForgotPasswordHandler)

	router.HandleFunc("GET /auth", gothic.BeginAuthHandler)
	router.HandleFunc("GET /auth/callback", controllers.GoogleCallbackHandler)

	router.HandleFunc("GET /cookie", controllers.GetCookie)

	// Add code here

	authRouter := http.NewServeMux()
	authRouter.HandleFunc("GET /auth/logout", controllers.GothLogout)
	authRouter.HandleFunc("GET /api/validate", controllers.Validate)

	router.Handle("/", middleware.AuthMiddleware(authRouter))