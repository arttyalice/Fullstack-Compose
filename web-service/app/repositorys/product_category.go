package repositorys

import (
	"deeploy-exam/app/database"
	"deeploy-exam/app/models"
	"log"
)

type Category struct{}

func (*Category) Getlist(p *models.SearchProductCategory) (list []models.ListProductCategory, count int64) {
	list = make([]models.ListProductCategory, 0)
	db := database.UseSQL()
	query := db.
		Model(&list)

	if p.Page != nil && p.Size != nil {
		offset := *p.Size * (*p.Page - 1)
		query = query.Limit(*p.Size).Offset(offset)
	}
	err := query.Count(&count).Error
	if err != nil {
		log.Print(err)
		return
	}
	err = query.Find(&list).Error
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (*Category) GetInfo(id uint) (info *models.ListProductCategory) {
	db := database.UseSQL()
	info = new(models.ListProductCategory)
	err := db.Where("id = ?", id).First(&info).Error
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (*Category) Create(category *models.CreateProductCategory) (err error) {
	db := database.UseSQL()
	err = db.Create(category).Error
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (*Category) Update(category *models.UpdateProductCategory) (err error) {
	db := database.UseSQL()
	err = db.Updates(category).Error
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (*Category) Delete(categoryID uint) (err error) {
	db := database.UseSQL()
	err = db.Where("id = ?", categoryID).Delete(&models.ProductCategory{}).Error
	if err != nil {
		log.Print(err)
		return
	}
	return
}
