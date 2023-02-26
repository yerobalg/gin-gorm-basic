package entity

var (
	HealthCategory     = "Health"
	TechnologyCategory = "Technology"
	BusinessCategory   = "Business"
	TravelCategory     = "Travel"
)

type Category struct {
	ID   uint
	Name string `gorm:"type:VARCHAR(30)"`
}