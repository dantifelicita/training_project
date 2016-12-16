package talk_training

import "time"

type Talks struct {
	ID         int64     `json:"id,string"`
	ProductID  int64     `json:"product_id"`
	Message    string    `json:"message"`
	CreateTime time.Time `json:"create_time"`
}

type Messages struct {
	ProductID int64  `json:"product_id"`
	ShopID    int64  `json:"shop_id"`
	UserID    int64  `json:"user_id"`
	Message   string `json:"message"`
}
