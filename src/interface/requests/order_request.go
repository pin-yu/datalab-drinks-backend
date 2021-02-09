package requests

// OrderRequestBody models order request body
type OrderRequestBody struct {
	OrderBy string `json:"order_by"`
	Size    string `json:"size"`
	ItemID  uint   `json:"item_id"`
	SugarID uint   `json:"sugar_id"`
	IceID   uint   `json:"ice_id"`
}

// IsSchemaValid returs false if there is wrong schema
func (o *OrderRequestBody) IsSchemaValid() bool {
	if o.OrderBy == "" {
		return false
	}

	if o.Size != "medium" && o.Size != "large" {
		return false
	}

	if o.ItemID == 0 {
		return false
	}

	if o.SugarID == 0 {
		return false
	}

	if o.IceID == 0 {
		return false
	}

	return true
}
