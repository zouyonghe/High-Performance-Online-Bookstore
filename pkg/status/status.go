package status

type BaseStatus struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

type ShopStatus struct {
	BaseStatus
	StatusInfo string `json:"status_info"`
}
