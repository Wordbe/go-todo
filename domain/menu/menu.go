package menu

type Menu struct {
	Id    int64   `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

var Menus = []Menu{
	{Id: 1, Title: "Blue Cake", Price: 56.99},
	{Id: 2, Title: "Fish", Price: 17.99},
	{Id: 3, Title: "Sarah Spaghetti and Clifford Potatoes", Price: 39.99},
}
