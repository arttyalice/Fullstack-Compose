package repositorys

import (
	"deeploy-exam/app/database"
	"deeploy-exam/app/models"
	"log"
)

type StoreProduct struct{}

func (*StoreProduct) Getlist(p *models.SearchStoreProduct) (list []models.ListStoreProduct, count int64) {
	list = make([]models.ListStoreProduct, 0)
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

func (*StoreProduct) GetInfo(id uint) (info *models.StoreProductInfo) {
	// db := database.UseSQL()
	// info = new(models.StoreProductInfo)
	// product := new(models.ListProduct)
	// store := new(models.ListStore)

	// var err error
	// var wg sync.WaitGroup
	// wg.Add(2)
	// go func() {
	// 	err = db.Where("id = ?", id).First(info).Error
	// 	wg.Done()
	// }()
	// go func() {

	// }()
	// wg.Wait()
	// if err != nil {
	// 	log.Print(err)
	// 	return
	// }
	// info.Products = products
	return
}

func (*StoreProduct) Create(req []models.CreateStoreProduct) (err error) {
	db := database.UseSQL()
	err = db.Create(&req).Error
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (*StoreProduct) Update(req *models.UpdateStoreProduct) (err error) {
	db := database.UseSQL()
	err = db.Updates(req).Error
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func (*StoreProduct) Delete(reqID uint) (err error) {
	db := database.UseSQL()
	err = db.Where("id = ?", reqID).Delete(&models.StoreProduct{}).Error
	if err != nil {
		log.Print(err)
		return
	}
	return
}
