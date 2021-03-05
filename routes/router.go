package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Prevent abnormal shutdown while panic
func recoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				log.Print(string(debug.Stack()))
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// Put params in context for sharing them between handlers
func wrapHandler(next http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), "params", ps)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

type TokenRes struct {
	Exp int `json:"exp"`
}

//RouterConfig function
func RouterConfig() http.Handler {
	router := httprouter.New()
	router.PanicHandler = panicHandler
	//indexHandlers := alice.New(recoverHandler)

	setPingRoutes(router)
	GetAboutPrivacyPolicy(router)

	router.Handler("GET", "/swagger", httpSwagger.WrapHandler)
	router.Handler("GET", "/swagger/:one", httpSwagger.WrapHandler)
	router.Handler("GET", "/swagger/:one/:two", httpSwagger.WrapHandler)
	router.Handler("GET", "/swagger/:one/:two/:three", httpSwagger.WrapHandler)
	router.Handler("GET", "/swagger/:one/:two/:three/:four", httpSwagger.WrapHandler)
	router.Handler("GET", "/swagger/:one/:two/:three/:four/:five", httpSwagger.WrapHandler)
	handler := cors.AllowAll().Handler(router)
	return handler
}

func panicHandler(w http.ResponseWriter, r *http.Request, c interface{}) {
	fmt.Println("(alcochange-dtx)Recovering from panic-Reason: %+v", c.(error))
	debug.PrintStack()
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(c.(error).Error()))
}
