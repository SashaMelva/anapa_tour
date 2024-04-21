package marcketmodel

type Production struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Count       int    `json:"count"`
}

type ByProduction struct {
	IdProd     int `json:"production_id"`
	IdAccount  int `json:"account_id"`
	Price      int `json:"price"`
	ContScklad int `json:"count_scklad"`
}
