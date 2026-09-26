package dbrepo

import (
	"context"
	"errors"
	"time"

	"github.com/bny64/bookings/internal/models"
	"golang.org/x/crypto/bcrypt"
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

// SearchAvailablilityByDatesByRoomID는 특정 기간의 예약 가능 여부를 확인합니다.
func (m *postgreDBRepo) SearchAvailablilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
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

// SearchAvailabilityForAllRooms는 전체 기간의 예약 가능 여부를 확인합니다.
func (m *postgreDBRepo) SearchAvailabilityForAllRooms(
	start, end time.Time,
) ([]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var rooms []models.Room

	query := `
		SELECT 
			r.id, r.room_name
		FROM
			rooms r
		WHERE
			r.id not in (
				SELECT room_id FROM room_restrictions WHERE
				$1 < end_date and $2 > start_date
			)
		`

	rows, err := m.DB.QueryContext(
		ctx,
		query,
		start,
		end,
	)

	if err != nil {
		return rooms, err
	}

	for rows.Next() {
		var room models.Room
		err := rows.Scan(&room.ID, &room.RoomName)
		if err != nil {
			return rooms, err
		}
		rooms = append(rooms, room)
	}

	if err = rows.Err(); err != nil {
		return rooms, err
	}

	return rooms, nil
}

// GetRoomByID gets a room by id
func (m *postgreDBRepo) GetRoomByID(id int) (models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var room models.Room

	query := `
		select id, room_name, created_at, updated_at from rooms where id = $1
	`

	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&room.ID,
		&room.RoomName,
		&room.CreatedAt,
		&room.UpdatedAt,
	)

	if err != nil {
		return room, err
	}

	return room, nil
}

// GetUserByID gets a user by id
func (m *postgreDBRepo) GetUserByID(id int) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, first_name, last_name, email, password, access_level, created_at, updated_at
					from users where id = $1
	`
	row := m.DB.QueryRowContext(ctx, query, id)

	var u models.User

	err := row.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Password,
		&u.AccessLevel,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		return u, err
	}

	return u, nil
}

// UpdateUser updates a user in the database
func (m *postgreDBRepo) UpdateUser(u models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		update users set first_name = $1, last_name = $2, email = $3, access_level = $4, updated_at = $5 where id = $6
	`

	_, err := m.DB.ExecContext(ctx, query,
		u.FirstName,
		u.LastName,
		u.Email,
		u.AccessLevel,
		u.UpdatedAt,
		u.ID)

	if err != nil {
		return err
	}

	return nil
}

// Authenticate authenticates a user by email and password
func (m *postgreDBRepo) Authenticate(email string, testPassword string) (int, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id int
	var hashedPassword string

	query := `
		select id, password
		from users where email = $1
	`

	row := m.DB.QueryRowContext(ctx, query, email)
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		return id, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(testPassword))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, "", errors.New("incorrect password")
	} else if err != nil {
		return 0, "", err
	}

	return id, hashedPassword, nil
}
