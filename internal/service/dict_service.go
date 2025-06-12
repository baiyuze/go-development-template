package service

import (
	"app/internal/common/log"
	"app/internal/dto"
	"app/internal/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

type DictService interface {
	Create(c *gin.Context, body *model.Dict) error
	Delete(c *gin.Context, body dto.DeleteIds) error
	List(context *gin.Context, query dto.ListQuery, name string) (dto.Result[dto.List[model.Dict]], error)
	Update(c *gin.Context, id int, body *model.Dict) error
	GetOptionsByDictCode(c *gin.Context, code string) ([]*model.DictItem, error)
}

type dictService struct {
	db  *gorm.DB
	log *log.LoggerWithContext
}

func NewDictService(
	db *gorm.DB,
	log *log.LoggerWithContext) DictService {

	return &dictService{db: db, log: log}
}

func ProvideDictService(container *dig.Container) {
	if err := container.Provide(NewDictService); err != nil {
		panic(err)
	}
}

// Create 创建
func (s *dictService) Create(c *gin.Context, body *model.Dict) error {
	if err := s.db.Transaction(func(tx *gorm.DB) error {
		// 添加数据到dict
		if err := tx.Create(&model.Dict{
			Name:        body.Name,
			Code:        body.Code,
			En:          body.En,
			Description: body.Description,
		}).Error; err != nil {
			return err
		}
		// 添加数据到dict_items
		for _, item := range body.Items {
			if err := tx.Create(&model.DictItem{
				DictCode: body.Code,
				Value:    item.Value,
				LabelZh:  item.LabelZh,
				LabelEn:  item.LabelEn,
				Status:   item.Status,
			}).Error; err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return err
	}
	return nil
}

// List 根据不同的code获取不同的options
func (s *dictService) List(c *gin.Context, query dto.ListQuery, name string) (dto.Result[dto.List[model.Dict]], error) {
	logger := s.log.WithContext(c)
	var dicts []model.Dict
	limit := query.PageSize
	offset := query.PageNum*query.PageSize - query.PageSize

	if result := s.db.
		Where("name LIKE ?", "%"+name+"%").
		Preload("Items").
		Limit(limit).
		Offset(offset).
		Order("create_time asc").
		Find(&dicts); result.Error != nil {
		logger.Error(result.Error.Error())
		return dto.ServiceFail[dto.List[model.Dict]](result.Error), result.Error
	}
	var count int64
	if result := s.db.Model(&model.Dict{}).Count(&count); result.Error != nil {
		logger.Error(result.Error.Error())
		return dto.ServiceFail[dto.List[model.Dict]](result.Error), result.Error
	}
	data := dto.ServiceSuccess(dto.List[model.Dict]{
		Items:    dicts,
		PageSize: query.PageSize,
		PageNum:  query.PageNum,
		Total:    count,
	})
	return data, nil
}

// GetOptionsByDictCode 根据不同的code获取不同的options
func (s *dictService) GetOptionsByDictCode(c *gin.Context, code string) ([]*model.DictItem, error) {
	var dictItems []*model.DictItem
	if err := s.db.Where("dict_code = ?", code).Find(&dictItems).Error; err != nil {
		return nil, err
	}

	return dictItems, nil
}

// Delete 删除
func (s *dictService) Delete(c *gin.Context, body dto.DeleteIds) error {
	var dicts []model.Dict
	if err := s.db.Preload("Items").Find(&dicts, body.Ids).Error; err != nil {
		return err
	}
	// 如果想删除主表记录同步删除从表记录，需要设置constraint:OnDelete:CASCADE
	// 当前采用先删除从表，再删除主表记录
	for _, dict := range dicts {
		for _, item := range dict.Items {
			if err := s.db.Delete(&item).Error; err != nil {
				return err
			}
		}
		if err := s.db.Delete(&dict).Error; err != nil {
			return err
		}
	}

	return nil
}

// Update 更新字典数据,需要使用事务
func (s *dictService) Update(c *gin.Context, id int, body *model.Dict) error {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		var dict model.Dict
		if err := tx.First(&dict, id).Error; err != nil {
			return err
		}
		//更新关系，先清空再更新
		// 不允许clear是因为外键不能为null导致的
		//if err := tx.Model(&dict).Association("Items").Clear(); err != nil {
		//	return err
		//}
		//	先更新id数据，再更新关联关系
		if err :=
			tx.Model(&dict).Updates(model.Dict{
				Name:        body.Name,
				En:          body.En,
				Code:        body.Code,
				Description: body.Description,
			}).Error; err != nil {
			return nil
		}

		if body.Items != nil {
			for _, item := range body.Items {
				dictItems := model.DictItem{
					DictCode: body.Code,
					Value:    item.Value,
					LabelZh:  item.LabelZh,
					LabelEn:  item.LabelEn,
					Status:   item.Status,
				}
				if item.ID != 0 {
					if err := tx.Model(&model.DictItem{
						ID: item.ID,
					}).Updates(&dictItems).Error; err != nil {
						return err
					}
				} else {
					if err := tx.Create(&dictItems).Error; err != nil {
						return err
					}
				}
			}
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
