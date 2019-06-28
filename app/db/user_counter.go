package db

type UserCounter struct {
	UserID           int64  `db:"user_id"`
	ReservationCount uint64 `db:"reservation_count"`
}

func GetUserCounter(userID int64) (*UserCounter, error) {
	q := "SELECT * FROM `user_counter` WHERE `user_id` = ?"
	var counter UserCounter
	if err := db.Get(&counter, q, userID); err != nil {
		return nil, err
	}

	return &counter, nil
}

func IncrementUserReservationCount(userID int64) error {
	q := "INSERT INTO `user_counter` (`user_id`, `reservation_count`) VALUES (?, ?) " +
		" ON DUPLICATE KEY UPDATE `reservation_count` = `reservation_count` + 1"
	if _, err := db.Exec(q, userID, 1); err != nil {
		return err
	}
	return nil
}

func DecrementUserReservationCount(userID int64) error {
	q := "UPDATE `user_counter` SET `reservation_count` = `reservation_count` - 1 WHERE `user_id` = ?"
	if _, err := db.Exec(q, userID); err != nil {
		return err
	}
	return nil
}
