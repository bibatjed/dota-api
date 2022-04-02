package models

type Hero struct {
	HeroName      string `json:"hero_name" db:"hero_name"`
	LocalizedName string `json:"localized_name" db:"localized_name"`
	ClassName string `json:"class_name" db:"class_name"`
}
