package domain

type RoomStatus string

const (
	RoomStatusAvailable   RoomStatus = "available"
	RoomStatusOccupied    RoomStatus = "occupied"
	RoomStatusMaintenance RoomStatus = "maintenance"
)

func (rs RoomStatus) IsValid() bool {
	switch rs {
	case RoomStatusAvailable, RoomStatusOccupied, RoomStatusMaintenance:
		return true
	default:
		return false
	}
}