package domain

type PerformanceStage struct {
	Id           string  `json:"stage_id" bson:"_id" gorm:"primaryKey;column:stage_id;type:string"`
	RoomNumber   string  `json:"room_number" bson:"room_number" gorm:"column:room_number"`
	SeatCapacity int     `json:"seat_capacity" bson:"seat_capacity" gorm:"column:seat_capacity"`
	PricePerSeat float64 `json:"price_per_seat" bson:"price_per_seat" gorm:"column:price_per_seat"`
}
