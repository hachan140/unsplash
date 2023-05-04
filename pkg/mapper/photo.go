package mapper

import (
	"gin_unsplash/pkg/dto"
	"gin_unsplash/pkg/model"
)

func PhotosToDTOs(ms []*model.Photo) []*dto.Photo {
	ds := make([]*dto.Photo, 0, len(ms))
	for _, m := range ms {
		d := PhotoToDTO(m)
		ds = append(ds, d)
	}
	return ds
}

func PhotoToDTO(m *model.Photo) *dto.Photo {
	if m == nil {
		return nil
	}
	return &dto.Photo{
		ID:             m.ID,
		CreatedAt:      m.CreatedAt,
		UpdatedAt:      m.UpdatedAt,
		Width:          m.Width,
		Height:         m.Height,
		Url:            m.Url,
		Description:    m.Description,
		AltDescription: m.AltDescription,
		Likes:          m.Likes,
	}
}

func PhotoToModel(d *dto.Photo) *model.Photo {
	if d == nil {
		return nil
	}
	return &model.Photo{
		ID:             d.ID,
		CreatedAt:      d.CreatedAt,
		UpdatedAt:      d.UpdatedAt,
		Width:          d.Width,
		Height:         d.Height,
		Url:            d.Url,
		Description:    d.Description,
		AltDescription: d.AltDescription,
		Likes:          d.Likes,
	}
}
