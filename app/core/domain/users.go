package domain

type Users struct {
	Id       string     `json:"user_id" bson:"_id" gorm:"primaryKey;column:user_id;type:string"`
	Username string     `json:"username" bson:"username" gorm:"column:username;unique"`
	Password string     `json:"password" bson:"password" gorm:"column:password"`
	Role     string     `json:"role" bson:"role" gorm:"column:role;default:user"`
	Bookings []Bookings `json:"bookings" bson:"bookings" gorm:"foreignKey:UserId;references:Id"`
}
