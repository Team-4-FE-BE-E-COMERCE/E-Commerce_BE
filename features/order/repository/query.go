package data

import (
	"errors"
	"project/e-commerce/features/order"

	"gorm.io/gorm"
)

type transactionData struct {
	db *gorm.DB
}

func New(db *gorm.DB) order.DataInterface {
	return &transactionData{
		db: db,
	}
}

func (repo *transactionData) InsertData(token int, dataReq order.AddressCore) (int, error) {

	var inputTransaction []Results
	tx := repo.db.Model(&Product{}).Select("carts.id, carts.quantity, products.name, products.images, products.price, carts.user_id, carts.product_id").Joins("inner join carts on carts.product_id = products.id").Where("carts.user_id = ?", token).Scan(&inputTransaction)
	if tx.Error != nil {
		return -1, tx.Error
	}

	dataRes := insertJoin(inputTransaction)
	data := insert(dataRes)

	for i, v := range data {
		data[i].TotalPrice = v.TotalPrice * v.Quantity
		txCreate := repo.db.Create(&data[i])
		if txCreate.Error != nil {
			return -1, txCreate.Error
		}
	}

	var id []int
	str := "waiting"
	txId := repo.db.Model(&Cart{}).Select("transactions.id").Joins("inner join transactions on transactions.cart_id = carts.id").Where("carts.user_id = ? AND transactions.order_status = ?", token, str).Scan(&id)
	if txId.Error != nil {
		return -1, txId.Error
	}

	for _, v := range id {

		dataReq.TransactionID = uint(v)
		dataCreate := toDb(dataReq)
		txCreate := repo.db.Create(&dataCreate)
		if txCreate.Error != nil {
			return -1, txCreate.Error
		}

	}

	return 1, nil

}

func (repo *transactionData) UpdateStatus(token int, status string) (int, error) {

	var id []int
	str := "waiting"
	tx := repo.db.Model(&Cart{}).Select("transactions.id").Joins("inner join transactions on transactions.cart_id = carts.id").Where("carts.user_id = ? AND transactions.order_status = ?", token, str).Scan(&id)
	if tx.Error != nil {
		return -1, tx.Error
	}

	for _, valueId := range id {
		if valueId == 0 {
			return -1, errors.New("not found")
		}
		txUpdate := repo.db.Model(&Transaction{}).Where("id = ?", valueId).Update("order_status", status)
		if txUpdate.Error != nil {
			return -1, txUpdate.Error
		}
	}

	txDel := repo.db.Where("user_id = ?", token).Delete(&Cart{})
	if txDel.Error != nil {
		return -1, txDel.Error
	}

	return 1, nil

}

func (repo *transactionData) CancelOrder(token int, status string) (int, error) {
	var id []int
	str := "cancelled"
	tx := repo.db.Model(&Cart{}).Select("transactions.id").Joins("inner join transactions on transactions.cart_id = carts.id").Where("carts.user_id = ? AND transactions.order_status = ?", token, str).Order("transactions.id Desc").Scan(&id)
	if tx.Error != nil {
		return -1, tx.Error
	}

	for _, valueId := range id {
		if valueId == 0 {
			return -1, errors.New("no data")
		}
		txUpdate := repo.db.Model(&Transaction{}).Where("id = ?", valueId).Update("order_status", status)
		if txUpdate.Error != nil {
			return -1, txUpdate.Error
		}
	}

	return 1, nil
}

func (repo *transactionData) SelectOrder(token int) ([]order.HistoryOrder, error) {

	var data []order.HistoryOrder
	// status := "confirmed"
	tx := repo.db.Unscoped().Model(&Product{}).Select("products.images, products.name, products.price, carts.quantity, transactions.order_status").Joins("inner join carts on carts.product_id = products.id").Joins("inner join transactions on transactions.cart_id = carts.id").Where("carts.user_id = ? AND carts.deleted_at IS NOT NULL", token).Scan(&data)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return data, nil
}

//transactions.order_status = ? AND
