package entity

type OrderRepository interface {
	Save(o *Order) error
	GetTotalTransactions() (int, error)
}
