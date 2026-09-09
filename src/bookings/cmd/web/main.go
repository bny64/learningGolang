package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/bny64/bookings/internal/config"
	"github.com/bny64/bookings/internal/handlers"
	"github.com/bny64/bookings/internal/models"
	"github.com/bny64/bookings/internal/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the application function
func main() {

	// 세션(scs)은 내부적으로 encoding/gob을 사용하여 데이터를 직렬화/역직렬화합니다.
	// 기본(primitive) 타입 외에 models.Reservation 같은 사용자 정의 구조체를 세션에 저장하려면
	// gob 패키지에 미리 타입을 등록(Register)해야 런타임 오류 없이 저장 및 조회가 가능합니다.
	gob.Register(models.Reservation{})

	//change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteStrictMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	// _ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
