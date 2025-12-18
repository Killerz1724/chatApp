package usecase

import (
	"chatApp/entity"
	"chatApp/repository"
	"context"
)

type ConvUsecaseItf interface {
	UsecaseGetListConvs(context.Context, entity.ReqListConvBody) (*[]entity.ResListConvBody, error)
}

type ConvUsecaseImpl struct {
	cr repository.ConvRepositoryItf
}

func NewConvUsecase(cr repository.ConvRepositoryItf) ConvUsecaseImpl {
	return ConvUsecaseImpl{cr: cr}
}

func (cu ConvUsecaseImpl) UsecaseGetListConvs(c context.Context, req entity.ReqListConvBody) (*[]entity.ResListConvBody, error) {

	res, err := cu.cr.RepoGetListConvs(c, req)

	if err != nil {
		return nil, err
	}

	return res, nil 
}