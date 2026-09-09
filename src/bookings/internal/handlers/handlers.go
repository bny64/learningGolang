package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bny64/bookings/internal/config"
	"github.com/bny64/bookings/internal/forms"
	"github.com/bny64/bookings/internal/models"
	"github.com/bny64/bookings/internal/render"
)

// Repo는 핸들러들이 사용하는 리포지토리 인스턴스입니다.
var Repo *Repository

// Repository는 리포지토리 구조체 타입입니다.
type Repository struct {
	App *config.AppConfig
}

// NewRepo는 새로운 리포지토리를 생성합니다.
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers는 핸들러에 사용할 리포지토리를 설정합니다.
func NewHandlers(r *Repository) {
	Repo = r
}

// Home은 홈 페이지 핸들러입니다.
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	log.Println("Requested URL:", r.URL.Path)

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About은 소개 페이지 핸들러입니다.
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	log.Println("Requested URL:", r.URL.Path)
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}

// Reservation은 예약 페이지 핸들러입니다.
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	emptyReservation := models.Reservation{}
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// PostReservation은 예약 처리(POST) 핸들러입니다.
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}

	form := forms.New(r.PostForm)

	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3, r)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	// main.go에서 gob.Register(models.Reservation{})로 타입을 등록해 두었기 때문에
	// 사용자 정의 구조체인 reservation을 세션에 직렬화하여 저장할 수 있습니다.
	m.App.Session.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

// Generals는 장군(Generals) 객실 페이지 핸들러입니다.
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.tmpl", &models.TemplateData{})
}

// Majors는 소령(Majors) 객실 페이지 핸들러입니다.
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.page.tmpl", &models.TemplateData{})
}

// Availability는 예약 가능 여부 조회 페이지 핸들러입니다.
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// PostAvailability는 예약 가능 여부 조회 처리(POST) 핸들러입니다.
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte("Posted to search availability " + start + " to " + end))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON은 예약 가능 여부를 확인하고 JSON 응답을 전송하는 핸들러입니다.
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "     ")

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// Contact는 문의 페이지 핸들러입니다.
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

// ReservationSummary는 예약 요약 페이지 핸들러입니다.
func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		log.Println("cannot get reservation from session")
		m.App.Session.Put(r.Context(), "error", "Can't get reservation from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(r.Context(), "reservation")
	data := make(map[string]interface{})
	data["reservation"] = reservation

	render.RenderTemplate(w, r, "reservation-summary.page.tmpl", &models.TemplateData{
		Data: data,
	})
}
