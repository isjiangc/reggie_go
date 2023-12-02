package v1

type CreateDishRequest struct {
	Name        string    `json:"name"`        //菜品名称
	Price       float64   `json:"price"`       //菜品价格
	Code        string    `json:"code"`        //商品码
	Image       string    `json:"image"`       //图片
	Description string    `json:"description"` //描述信息
	Status      int       `json:"status"`      //0 停售 1 起售
	CategoryId  string    `json:"categoryId"`  //菜品分类id
	Flavors     []Flavors `json:"flavors"`     // 口味
}
type Flavors struct {
	Name       string `json:"name"`  // 口味名称
	Value      string `json:"value"` // 口味数据list
	ShowOption bool   `json:"showOption"`
}
