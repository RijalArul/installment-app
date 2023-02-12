package web

type GoodRequestDTO struct {
	Name         string `json:"name" form:"name"`
	Slug         string `json:"slug" form:"slug"`
	Price        int    `json:"price" form:"price"`
	Rate         int    `json:"rate" form:"rate"`
	GoodsOwnerID uint
}

type GoodResponseBodyDTO struct {
	Name         string `json:"name" form:"name"`
	Slug         string `json:"slug" form:"slug"`
	Price        int    `json:"price" form:"price"`
	Rate         int    `json:"rate" form:"rate"`
	GoodsOwnerID uint
}
