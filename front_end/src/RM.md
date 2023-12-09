// GeneralInfo, SupermarketInfo, interact with the Frontend
type GeneralInfo struct {
Date time.Time `json:"date"`
Event string `json:"event"`
WarehouseProduct map[string]int `json:"warehouseProduct"`
}

type SupermarketInfo struct {
ID string `json:"id"`
ProductLeft map[string]int `json:"productLeft"` // The number of products left in the supermarket
ProductAdd map[string]int `json:"productAdd"` // The number of products added to the supermarket
DeliveryTime int `json:"deliveryTime"`
}
