package status

var (
	Ok   = BaseStatus{Code: 0, Status: "ok"}
	Fail = BaseStatus{Code: 1, Status: "fail"}

	Open  = ShopStatus{Ok, "open"}
	Close = ShopStatus{Fail, "close"}
)
