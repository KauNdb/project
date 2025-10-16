package auth

import "project/db"

type AuthRepository struct {
	DataBase *db.Db
}

func NewAuthRepository(db *db.Db) *AuthRepository {
	return &AuthRepository{
		DataBase: db,
	}
}

func (authRepo *AuthRepository) CreatePhone(phone *Phone) (*Phone, error) {
	result := authRepo.DataBase.DB.Create(phone)
	if result.Error != nil {
		return nil, result.Error
	}

	return phone, nil
}

func (authRepo *AuthRepository) UpatePhone(phone *Phone) (*Phone, error) {
	result := authRepo.DataBase.DB.Where("phone = ?", phone.Phone).Updates(phone)
	if result.Error != nil {
		return nil, result.Error
	}

	return phone, nil
}

func (authRepo *AuthRepository) GetPhone(phoneNum string) error {
	var phone Phone
	result := authRepo.DataBase.DB.First(&phone, "phone = ?", phoneNum)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (authRepo *AuthRepository) GetPhoneByCode(sessionId string, code int) (*Phone, error) {
	var phone Phone
	result := authRepo.DataBase.DB.First(&phone, "session_id = ? AND  code = ?", sessionId, code)
	if result.Error != nil {
		return nil, result.Error
	}

	return &phone, nil
}
