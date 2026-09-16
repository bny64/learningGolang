package dbrepo

import (
	"context"
	"time"

	"github.com/bny64/bookings/internal/models"
)

func (m *postgreDBRepo) AllUsers() bool {
	return true
}

// InsertReservation은 예약 정보를 데이터베이스에 저장합니다.
func (m *postgreDBRepo) InsertReservation(res models.Reservation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into reservations (first_name, last_name, email, phone, start_date, end_date, room_id, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id`

	_, err := m.DB.ExecContext(ctx, stmt, res.FirstName, res.LastName, res.Email, res.Phone, res.StartDate, res.EndDate, res.RoomID, time.Now(), time.Now())

	return err
}
