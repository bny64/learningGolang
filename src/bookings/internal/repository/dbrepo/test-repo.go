package dbrepo

import (
	"errors"
	"time"

	"github.com/bny64/bookings/internal/models"
)

// AllUsers는 시스템의 모든 사용자 목록 조회를 수행합니다.
func (m *testDBRepo) AllUsers() bool {
	return true
}

// InsertReservation은 예약 정보를 데이터베이스에 저장합니다.
func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {
	return 1, nil
}

// InsertRoomRestriction은 예약 정보를 데이터베이스에 저장합니다.
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	return nil
}

// SearchAvailablilityByDatesByRoomID는 특정 기간의 예약 가능 여부를 확인합니다.
func (m *testDBRepo) SearchAvailablilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
	return false, nil
}

// SearchAvailabilityForAllRooms는 전체 기간의 예약 가능 여부를 확인합니다.
func (m *testDBRepo) SearchAvailabilityForAllRooms(
	start, end time.Time,
) ([]models.Room, error) {
	var rooms []models.Room
	return rooms, nil
}

// GetRoomByID gets a room by id
func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {

	var room models.Room
	if id > 2 {
		return room, errors.New("Some error")
	}

	return room, nil
}
