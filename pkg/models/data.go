package models

var Db []Pizza

type Pizza struct {
	ID    int     `json:"id"`
	Size  int     `json:"size"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

func init() {
	pizza1 := Pizza{
		ID:    1,
		Size:  30,
		Title: "Margarita",
		Price: 300,
	}
	pizza2 := Pizza{
		ID:    2,
		Size:  26,
		Title: "Pepperoni",
		Price: 200,
	}
	pizza3 := Pizza{
		ID:    3,
		Size:  35,
		Title: "BBQ",
		Price: 450,
	}

	Db = append(Db, pizza1, pizza2, pizza3)
}

func FindPizzaByID(id int) (Pizza, bool) {
	var pizza Pizza
	var found bool
	for _, p := range Db {
		if p.ID == id {
			pizza = p
			found = true
			break
		}
	}
	return pizza, found
}
