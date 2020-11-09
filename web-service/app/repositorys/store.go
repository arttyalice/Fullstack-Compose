package repositorys

import (
	"deeploy-exam/app/database"
	"deeploy-exam/app/models"
	"log"
	"sync"
)

type Store struct{}

func (*Store) Getlist(p *models.SearchStore) (list []models.ListStore, count int64) {
	list = make([]models.ListStore, 0)
	db := database.UseSQL()
	query := db.Table("stores").Where("deleted_at IS NULL")

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

func (*Store) GetInfo(id uint) (info *models.StoreInfo) {
	db := database.UseSQL()
	info = new(models.StoreInfo)
	products := make([]models.ListProduct, 0)

	var err error
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		err = db.Where("id = ?", id).First(info).Error
		wg.Done()
	}()
	go func() {
		err = db.
			Table("products p").
			Joins("LEFT JOIN store_products sp ON p.id = sp.product_id").
			Where("sp.store_id = ?", id).
			Select("DISTINCT ON(p.id) p.*, sp.product_quantity as quantity").
			Scan(&products).Error
		wg.Done()
	}()
	wg.Wait()
	if err != nil {
		log.Print(err)
		return
	}
	info.Products = products
	return
}

func (*Store) Create(Store *models.CreateStore) (err error) {
	db := database.UseSQL()
	err = db.Create(Store).Error
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (*Store) Update(Store *models.UpdateStore) (err error) {
	db := database.UseSQL()
	log.Printf("%+v", Store)
	err = db.Updates(Store).Error
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (*Store) Delete(StoreID uint) (err error) {
	db := database.UseSQL()
	err = db.Where("id = ?", StoreID).Delete(&models.Store{}).Error
	if err != nil {
		log.Print(err)
		return
	}
	return
}
