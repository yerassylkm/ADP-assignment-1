package Hotel

type HotelManager struct {
	rooms map[string]Room
}

func NewHotel() *HotelManager {
	return &HotelManager{
		rooms: make(map[string]Room),
	}
}

func (h *HotelManager) AddRoom(r Room) {
	h.rooms[r.RoomNumber] = r
}

func (h *HotelManager) CheckIn(roomNumber string) error {
	room, exists := h.rooms[roomNumber]
	if !exists {
		return ErrRoomNotFound
	}
	if room.IsOccupied {
		return ErrRoomOccupied
	}

	room.IsOccupied = true
	h.rooms[roomNumber] = room
	return nil
}

func (h *HotelManager) CheckOut(roomNumber string) error {
	room, exists := h.rooms[roomNumber]
	if !exists {
		return ErrRoomNotFound
	}

	room.IsOccupied = false
	h.rooms[roomNumber] = room
	return nil
}

func (h *HotelManager) GetVacantRooms() []Room {
	var vacant []Room
	for _, room := range h.rooms {
		if !room.IsOccupied {
			vacant = append(vacant, room)
		}
	}
	return vacant
}