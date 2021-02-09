package entities

// WeekOrder stands for all orders within a week
type WeekOrder struct {
	Orders []Order `json:"week_orders"`
}
