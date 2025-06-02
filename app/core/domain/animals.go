package domain

type Animals struct {
	Id           string `json:"animal_id" bson:"_id" gorm:"primaryKey;column:animal_id;type:string"`
	Name         string `json:"name" bson:"name" gorm:"column:name"`
	Species      string `json:"species" bson:"species" gorm:"column:species"`
	Type         string `json:"type" bson:"type" gorm:"column:type"`
	ShowDuration int    `json:"show_duration" bson:"show_duration" gorm:"column:show_duration"`
}
