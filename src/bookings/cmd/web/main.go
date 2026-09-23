package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/bny64/bookings/helpers"
	"github.com/bny64/bookings/internal/config"
	"github.com/bny64/bookings/internal/driver"
	"github.com/bny64/bookings/internal/handlers"
	"github.com/bny64/bookings/internal/models"
	"github.com/bny64/bookings/internal/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

// main is the application function
func main() {

	db, err := run()

	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	defer close(app.MailChan)

	fmt.Println("Starting mail listener...")
	listenForMail()

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() (*driver.DB, error) {

	// 세션(scs)은 내부적으로 encoding/gob을 사용하여 데이터를 직렬화/역직렬화합니다.
	// 기본(primitive) 타입 외에 models.Reservation 같은 사용자 정의 구조체를 세션에 저장하려면
	// gob 패키지에 미리 타입을 등록(Register)해야 런타임 오류 없이 저장 및 조회가 가능합니다.
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})
	gob.Register(models.RoomRestriction{})

	mailChan := make(chan models.MailData)
	app.MailChan = mailChan

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	//change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteStrictMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	//데이터베이스를 연결
	log.Println("Connecting to database...")
	db, err := driver.ConnectSQL("host=bny64.ddns.net port=30004 dbname=bookings user=bny64 password=Namyul64!")
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	fmt.Println("Successfully connected to database")

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("cannot create template cache", err)
		return nil, err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)
	helpers.NewHelpers(&app)
	return db, nil
}
