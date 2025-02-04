package usecase

type userUsecaseItf interface {}

type userUsecase struct {}

func NewuserUsecase() userUsecaseItf {
    return &userUsecase{}
}
