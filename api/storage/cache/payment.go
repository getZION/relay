package cache

import "github.com/getzion/relay/api"

func (c *cacheStorage) GetPayments() ([]api.Payment, error) {
	// var payments []v1.PaymentORM
	// result := s.connection.DB.Find(&payments)
	// if result.Error != nil {
	// 	return nil, result.Error
	// }
	// return payments, nil
	return nil, nil
}

func (c *cacheStorage) InsertPayment(*api.Payment) error {
	// payment.Zid = uuid.NewString()

	// err := validator.Struct(&payment)
	// if err != nil {
	// 	return nil, err
	// }

	// result := s.connection.DB.Create(&payment)
	// if result.Error != nil {
	// 	return nil, result.Error
	// }

	// return &payment, nil
	return nil
}
