package common

import "github.com/getzion/relay/api"

func (c *Connection) GetPayments() ([]api.Payment, error) {
	var payments []api.Payment
	result := c.db.Find(&payments)
	if result.Error != nil {
		return nil, result.Error
	}
	return payments, nil
}

func (c *Connection) InsertPayment(*api.Payment) error {
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
