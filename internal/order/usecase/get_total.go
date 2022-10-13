package usecase

import "github.com/vyctor/pfa-go/internal/order/entity"

type GetTotalOutputDto struct {
	Total int
}

type GetTotalUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetTotalUseCase(orderRepository entity.OrderRepositoryInterface) *GetTotalUseCase {
	return &GetTotalUseCase{OrderRepository: orderRepository}
}

func (uc *GetTotalUseCase) Execute() (*GetTotalOutputDto, error) {
	total, err := uc.OrderRepository.GetTotal()
	if err != nil {
		return nil, err
	}
	return &GetTotalOutputDto{Total: total}, nil
}
