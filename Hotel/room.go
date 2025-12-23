package Hotel

// Room представляет модель данных номера отеля
type Room struct {
	RoomNumber    string
	Type          string
	PricePerNight float64
	IsOccupied    bool
}