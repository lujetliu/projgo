package service

import (
	"blog/internal/model"
	"blog/pkg/app"
)

/*
	github.com/go-playground/validator/v10 包校验规则:
	required	必填
 	gt	大于
 	gte	大于等于
 	lt	小于
 	lte	小于等于
 	min	最小值
 	max	最大值
 	oneof	参数集内的其中之一
 	len	长度要求与 len 给定的一致
*/

type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,defalut=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,defalut=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name      string `form:"name" binding:"min=3, max=100"`
	CreatedBy string `form:"created_by" binding:"required, min=3, max=100"`
	State     uint8  `form:"state, default=1" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	ID         uint32 `form:"id" binding:"required, gte=1"`
	Name       string `form:"name" binding:"minx=3, max=100"`
	State      uint8  `form:"state" binding:"required, oneof=0 1"`
	ModifiedBy string `json:"modified_by" binding:"required, min=3, max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required, gte=1"`
}

func (svc *Service) CountTag(param *CountTagRequest) (int, error) {
	return svc.dao.CountTag(param.Name, param.State)
}

func (svc *Service) GetTagList(param *TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return svc.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) CreateTag(param *CreateTagRequest) error {
	return svc.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

func (svc *Service) UpdateTag(param *UpdateTagRequest) error {
	return svc.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
}

func (svc *Service) DeleteTag(param *DeleteTagRequest) error {
	return svc.dao.DeleteTag(param.ID)
}
