package middleware

import (
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

type wrappedWrite struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWrite) writeHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func Logger(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapped := &wrappedWrite{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		next.ServeHTTP(wrapped, r)
		log.Println(wrapped.statusCode, r.Method, r.URL.Path, time.Since(start))
	}
}

func RecoveryMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if err := recover(); err != nil {
				msg := "Caught panic: %v, Stack treace: %s"
				log.Printf(msg, err, string(debug.Stack()))

				er := http.StatusInternalServerError
				http.Error(w, "Internal Server Error", er)
			}
		}()

		next.ServeHTTP(w, r)
	}
}

// func RequireAuth(next http.Handler) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// Retrieve the Authorization cookie
// 		cookie, err := r.Cookie("jwt")
// 		if err != nil {
// 			http.Error(w, "Unauthorized - missing token", http.StatusUnauthorized)
// 			return
// 		}
// 		tokenString := cookie.Value

// 		// Parse the JWT token
// 		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 			}
// 			return []byte(os.Getenv("SECRET")), nil
// 		})

// 		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

// 			if float64(time.Now().Unix()) > claims["exp"].(float64) {
// 				http.Error(w, "Unauthorized - token expired", http.StatusUnauthorized)
// 				return
// 			}
// 			var user models.User
// 			initializers.DB.First(&user, claims["sub"])

// 			if user.ID == 0 {
// 				http.Error(w, "Unauthorized - user not found", http.StatusUnauthorized)
// 				return
// 			}
// 			ctx := context.WithValue(r.Context(), "user", user)
// 			next.ServeHTTP(w, r.WithContext(ctx))
// 		} else {
// 			http.Error(w, "Unauthorized - token expired", http.StatusUnauthorized)
// 			return
// 		}

// 	}
// }

// func RequireAdmin(next http.Handler) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		// Access user information from context (set by RequireAuth)
// 		user, ok := r.Context().Value("user").(models.User)
// 		if !ok {
// 			http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 			return
// 		}

// 		// Check for admin status
// 		if !user.Admin {
// 			http.Error(w, "Forbidden - User is not an admin", http.StatusForbidden)
// 			return
// 		}

// 		// Continue processing the request if user is admin
// 		next.ServeHTTP(w, r)
// 	}
// }

type Middleware func(http.Handler) http.HandlerFunc

func MiddlewareChain(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.HandlerFunc {
		for _, middleware := range middlewares {
			next = middleware(next)
		}
		return next.ServeHTTP
	}
}
