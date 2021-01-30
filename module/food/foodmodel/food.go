package foodmodel

type Food struct {
	Id        string
	Name      string
	Description    bool
	Price    float64
	RestaurantId string
	ShortDescription string
}

type GetFood struct {
	Id string
	Name string
	Price float64
	ShortDescription string
}