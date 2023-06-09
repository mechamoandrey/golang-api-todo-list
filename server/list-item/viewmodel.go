package server

type CreateListItemRequest struct {
	Name        string
	Description string
}

type CreateListItemResponse struct {
	ListItemUUID string `json:"list_item_uuid"`
	Name         string `json:"name"`
	Description  string `json:"description"`
}
