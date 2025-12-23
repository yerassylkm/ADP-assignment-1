package Hotel

import "errors"

var (
	ErrRoomNotFound = errors.New("room does not exist")
	ErrRoomOccupied = errors.New("room is already occupied")
)