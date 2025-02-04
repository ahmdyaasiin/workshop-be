package usecase

type itemUsecaseItf interface {}

type itemUsecase struct {}

func NewitemUsecase() itemUsecaseItf {
    return &itemUsecase{}
}
