package mapper

import (
	"gin_unsplash/pkg/dto"
	"gin_unsplash/pkg/model"
)

func UsersToDTOs(ms []*model.User) []*dto.User {
	ds := make([]*dto.User, len(ms))
	for _, m := range ms {
		d := UserToDTO(m)
		ds = append(ds, d)

	}
	return ds
}
func UserToDTO(u *model.User) *dto.User {
	if u == nil {
		return nil
	}
	return &dto.User{
		Id:          u.Id,
		Username:    u.Username,
		Password:    u.Password,
		FullName:    u.FullName,
		PhoneNumber: u.PhoneNumber,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}
func UserToModel(d *dto.User) *model.User {
	if d == nil {
		return nil
	}
	return &model.User{
		Id:          d.Id,
		Username:    d.Username,
		Password:    d.Password,
		FullName:    d.FullName,
		PhoneNumber: d.PhoneNumber,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}
