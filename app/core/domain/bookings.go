package domain

type Bookings struct {
	Id         string  `json:"booking_id" bson:"_id" gorm:"primaryKey;column:booking_id;type:string"`
	UserId     string  `json:"user_id" bson:"user_id" gorm:"column:user_id;type:string"`
	RoundId    string  `json:"round_id" bson:"round_id" gorm:"column:round_id;type:string"`
	SeatNumber int     `json:"seat_number" bson:"seat_number" gorm:"column:seat_number"`
	Price      float64 `json:"price" bson:"price" gorm:"column:price"`
	QrCode     string  `json:"qr_code" bson:"qr_code" gorm:"column:qr_code"`
}
