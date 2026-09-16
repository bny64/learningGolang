package dbrepo

import (
	"context"
	"time"

	"github.com/bny64/bookings/internal/models"
)

// AllUsers는 시스템의 모든 사용자 목록 조회를 수행합니다.
func (m *postgreDBRepo) AllUsers() bool {
	return true
}

// InsertReservation은 예약 정보를 데이터베이스에 저장합니다.
func (m *postgreDBRepo) InsertReservation(res models.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newID int
	query := `
		INSERT INTO reservations (
			first_name, last_name, email, phone,
			start_date, end_date, room_id,
			created_at, updated_at
		) VALUES (
			$1, $2, $3, $4,
			$5, $6, $7,
			$8, $9
		) RETURNING id`

	err := m.DB.QueryRowContext(
		ctx,
		query,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomID,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, err
}

// InsertRoomRestriction은 예약 정보를 데이터베이스에 저장합니다.
func (m *postgreDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		INSERT INTO room_restrictions (
			room_id, start_date, end_date,
			created_at, updated_at,
			reservation_id, restriction_id
		) VALUES (
			$1, $2, $3,
			$4, $5,
			$6, $7
		)`

	_, err := m.DB.ExecContext(
		ctx,
		query,
		r.RoomID,
		r.StartDate,
		r.EndDate,
		time.Now(),
		time.Now(),
		r.ReservationID,
		r.RestrictionID,
	)

	if err != nil {
		return err
	}

	return nil
}

// SearchAvailablilityByDates는 특정 기간의 예약 가능 여부를 확인합니다.
func (m *postgreDBRepo) SearchAvailablilityByDates(start, end time.Time, roomID int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		SELECT 
			count(id)
		FROM
			room_restrictions
		WHERE
			$1 < end_date and $2 > start_date and room_id = $3;
		`

	var numRows int

	row := m.DB.QueryRowContext(
		ctx,
		query,
		start,
		end,
		roomID,
	)
	err := row.Scan(&numRows)

	if err != nil {
		return false, err
	}

	if numRows == 0 {
		return true, nil
	}

	return false, nil
}
