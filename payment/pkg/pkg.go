package pkg

import (
	"github.com/haobaojiang/diexample/common"
	"github.com/samber/do/v2"
)

type Payment struct {
}

func (p *Payment) Pay(id int64) error {
	return nil
}

func InitModule(injector do.Injector) {
	common.SafeDoProvide(injector, func(_injector do.Injector) (common.Payment, error) {
		return &Payment{}, nil
	})
}
