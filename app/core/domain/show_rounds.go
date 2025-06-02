package domain

type ShowRounds struct {
	Id       string     `json:"round_id" bson:"_id" gorm:"primaryKey;column:round_id;type:string"`
	AnimalId string     `json:"animal_id" bson:"animal_id" gorm:"column:animal_id;type:string"`
	StageId  string     `json:"stage_id" bson:"stage_id" gorm:"column:stage_id;type:string"`
	ShowTime string     `json:"show_time" bson:"show_time" gorm:"column:show_time;type:timestamp"`
	Bookings []Bookings `json:"bookings" bson:"bookings" gorm:"foreignKey:RoundId;references:Id"`
}
