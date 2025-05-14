package pkg

import (
	"github.com/haobaojiang/diexample/common"
	"github.com/samber/do/v2"
)

type Orders struct {
	payment common.Payment
}

func (Self *Orders) GetById() (common.Order, error) {
	return &Order{id: 1, payment: Self.payment}, nil
}

type Order struct {
	id      int64
	payment common.Payment
}

func (Self *Order) Id() int64 {
	return Self.id
}

func (Self *Order) Pay() error {
	return Self.payment.Pay(Self.id)
}

// ---------------------

func InitModule(injector do.Injector) {
	common.SafeDoProvide(injector, func(_injector do.Injector) (common.Orders, error) {
		payment := do.MustInvoke[common.Payment](injector)
		orders := &Orders{payment: payment}
		return orders, nil
	})
}
