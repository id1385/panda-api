package services

import (
	"time"
	"net/url"
	"github.com/wilsontamarozzi/panda-api/services/models"
	"github.com/wilsontamarozzi/panda-api/helpers"
	"github.com/wilsontamarozzi/panda-api/logger"
)

func GetPeople(pagination helpers.Pagination, q url.Values) models.People {

	var people models.People

	db := Con

	if q.Get("filter") != "" {
		db = db.Where("name iLIKE ?", "%" + q.Get("filter") + "%").
			Or("company_name iLIKE ?", "%" + q.Get("filter") + "%")
	}

	if q.Get("code") != "" {
		db = db.Where("code = ?", q.Get("code"))
	}

	if q.Get("name") != "" {
		db = db.Where("name iLIKE ?", "%" + q.Get("name") + "%")	
	}

	if q.Get("company_name") != "" {
		db = db.Where("company_name iLIKE ?", "%" + q.Get("company_name") + "%")	
	}

	if q.Get("gender") != "" {
		db = db.Where("gender = ?", q.Get("gender"))
	}

	if q.Get("type") != "" {
		db = db.Where("type = ?", q.Get("type"))
	}

	if q.Get("only_users") != "" {
		db = db.Where("is_user = ?", q.Get("only_users"))
	}
	
	db.Limit(pagination.ItemPerPage).
		Offset(pagination.StartIndex).
		Order("registered_at desc").
		Find(&people)

    return people
}

func GetPerson(personId string) models.Person {

	var person models.Person

	Con.Where("uuid = ?", personId).
		First(&person)

	return person
}

func DeletePerson(personId string) error {
	err := Con.Where("uuid = ?", personId).Delete(&models.Person{}).Error

	if err != nil {
		logger.Fatal(err)
	}

	return err;
}

func CreatePerson(person models.Person) (models.Person, error) {
	
	record := models.Person{
		Type 				: person.Type,
		Name 				: person.Name,
		CityName 			: person.CityName,
		CompanyName 		: person.CompanyName,
		Address 			: person.Address,
		Number 				: person.Number,
		Complement 			: person.Complement,
		District 			: person.District,
		Zip 				: person.Zip,
		BirthDate 			: person.BirthDate,
		Cpf 				: person.Cpf,
		Rg 					: person.Rg,
		Gender 				: person.Gender,
		BusinessPhone 		: person.BusinessPhone,
		HomePhone 			: person.HomePhone,
		MobilePhone 		: person.MobilePhone,
		Cnpj 				: person.Cnpj,
		StateInscription 	: person.StateInscription,
		Phone 				: person.Phone,
		Fax 				: person.Fax,
		Email 				: person.Email,
		Website 			: person.Website,
		Observations 		: person.Observations,
		RegisteredAt 		: time.Now(),
		RegisteredByUUID	: person.RegisteredByUUID,
	}

	err := Con.Set("gorm:save_associations", false).
		Create(&record).Error

	if err != nil {
		logger.Fatal(err)
	}

	return record, err
}

func UpdatePerson(person models.Person) (models.Person, error) {
	
	record := models.Person{
		Name 				: person.Name,
		CityName 			: person.CityName,
		CompanyName 		: person.CompanyName,
		Address 			: person.Address,
		Number 				: person.Number,
		Complement 			: person.Complement,
		District 			: person.District,
		Zip 				: person.Zip,
		BirthDate 			: person.BirthDate,
		Cpf 				: person.Cpf,
		Rg 					: person.Rg,
		Gender 				: person.Gender,
		BusinessPhone 		: person.BusinessPhone,
		HomePhone 			: person.HomePhone,
		MobilePhone 		: person.MobilePhone,
		Cnpj 				: person.Cnpj,
		StateInscription 	: person.StateInscription,
		Phone 				: person.Phone,
		Fax 				: person.Fax,
		Email 				: person.Email,
		Website 			: person.Website,
		Observations 		: person.Observations,
	}

	err := Con.Set("gorm:save_associations", false).
		Model(&models.Person{}).
		Where("uuid = ?", person.UUID).
		Updates(&record).Error

	if err != nil {
		logger.Fatal(err)
	}

	return record, err
}

func CountRowsPerson() int {
	var count int
	Con.Model(&models.Person{}).Count(&count)

	return count
}