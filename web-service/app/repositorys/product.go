package repositorys

import (
	"deeploy-exam/app/database"
	"deeploy-exam/app/models"
	"log"
	"sync"
)

type Product struct{}

func (*Product) Getlist(p *models.SearchProduct) (list []models.ListProduct, count int64) {
	list = make([]models.ListProduct, 0)
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

func (*Product) GetInfo(id uint) (info *models.ProductInfo) {
	db := database.UseSQL()
	info = new(models.ProductInfo)
	stores := make([]models.ListStore, 0)
	var category models.ProductCategory

	var err error
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		err = db.Where("id = ?", id).First(info).Error
		err = db.Where("id = ?", info.CategoryID).First(&category).Error
		info.Category = category
		wg.Done()
	}()
	go func() {
		err = db.
			Table("stores s").
			Joins("LEFT JOIN store_products sp ON s.id = sp.store_id").
			Where("sp.product_id = ?", id).
			Select("DISTINCT ON(s.id) s.*").
			Scan(&stores).Error
		wg.Done()
	}()
	wg.Wait()
	if err != nil {
		log.Print(err)
		return
	}
	info.Stores = stores
	return
}

func (*Product) Create(product *models.CreateProduct) (err error) {
	db := database.UseSQL()
	err = db.Create(product).Error
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (*Product) Update(product *models.UpdateProduct) (err error) {
	db := database.UseSQL()
	log.Printf("%+v", product)
	err = db.Updates(product).Error
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (*Product) Delete(productID uint) (err error) {
	db := database.UseSQL()
	err = db.Where("id = ?", productID).Delete(&models.Product{}).Error
	if err != nil {
		log.Print(err)
		return
	}
	return
}
