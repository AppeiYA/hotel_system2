package domain

type RoomType string 

const (
	RoomTypeSingle RoomType = "single"
	RoomTypeDouble RoomType = "double"
	RoomTypeSuite  RoomType = "suite"
	RoomTypePresidential RoomType = "presidential"
)

func (rt RoomType) IsValid() bool {
	switch rt {
	case RoomTypeSingle, RoomTypeDouble, RoomTypeSuite, RoomTypePresidential:
		return true
	default:
		return false
	}
}