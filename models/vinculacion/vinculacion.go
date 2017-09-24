package vinculacion
import (
	"database/sql"
	"fmt"
)
type Sys_Domain struct {

	Id				string		`json:"id"`
	Name			string		`json:"name"`
	Description		string		`json:"description"`
	Id_Dad			string		`json:"id_dad"`
	Created_at		string		`json:"created_at"`
	Updated_at		string		`json:"updated_at"`
	Status  		string		`json:"status"`
	Eforms  		string		`json:"eforms"`
	Modules  		string		`json:"modules"`

}
type Usr_Bonding struct {

	Name							string	`json:"name"`
	Lastname						string	`json:"lastname"`
	Type_Doc						string	`json:"type_doc"`
	Num_Doc							string	`json:"num_doc"`
	Nationality						string	`json:"nationality"`
	Sex								string	`json:"sex"`
	Place_Birth						string	`json:"place_birth"`
	Day_Birth						string	`json:"day_birth"`
	Status_Civil					string	`json:"status_civil"`
	Tel_Resi						string	`json:"tel_resi"`
	Tel_Cel							string	`json:"tel_cel"`
	Tel_ofi							string	`json:"tel_ofi"`
	Dir_Resi						string	`json:"dir_resi"`
	City							string	`json:"city"`
	Cod_Postal						string	`json:"cod_postal"`
	Nom_Business					string	`json:"nom_business"`
	Type_Business					string	`json:"type_business"`
	Cargo							string	`json:"cargo"`
	Anos_Business					string	`json:"anos_business"`
	Dir_Business					string	`json:"anos_business"`
	City_Business					string	`json:"city_business"`
	Resident						string	`json:"resident"`
	Num_Security					string	`json:"num_security"`
	Number_Fax						string	`json:"number_fax"`
	Email							string	`json:"email"`
	Correspondence					string	`json:"correspondence"`
	Confirmations					string	`json:"confirmations"`
	Account_Num1					string	`json:"account_num_1"`
	Account_Num2					string	`json:"account_num_2"`
	Account_Num3					string	`json:"account_num_3"`
	Account_Type1					string	`json:"account_type_1"`
	Account_Type2					string	`json:"account_type_2"`
	Account_Type3					string	`json:"account_type_3"`
	Institution1					string	`json:"institution_1"`
	Institution2					string	`json:"institution_2"`
	Institution3					string	`json:"institution_3"`
	Custodian_Account				string  `json:"custodian_account"`
	Institution_Other				string  `json:"institution_other"`
	Knowledge						string	`json:"knowledge"`
	Market_Cycle					string	`json:"market_cycle"`
	Academic_Level					string	`json:"academic_level"`
	Annual_Income					string	`json:"annual_income"`
	Total_Equity					string	`json:"total_equity"`
	Liquid_Assets					string	`json:"liquid_assets"`
	Investment_Objectives			string	`json:"investment_objectives"`
	Tolerance_Risk					string	`json:"tolerance_risk"`
	Total_Patrimony1				string	`json:"total_patrimony_1"`
	Total_Patrimony2				string	`json:"total_patrimony_2"`
	Total_Assets1					string	`json:"total_assets_1"`
	Total_Assets2					string	`json:"total_assets_2"`
	Source_Funds					string 	`json:"source_funds"`
	Attachments						string	`json:"attachments"`
	Bank_References_Institution1	string	`json:"bank_references_institution_1"`
	Bank_References_Institution2	string	`json:"bank_references_institution_2"`
	Bank_References_Institution3	string	`json:"bank_references_institution_3"`
	Bank_References_Address1		string	`json:"bank_references_address_1"`
	Bank_References_Address2		string	`json:"bank_references_address_2"`
	Bank_References_Address3		string	`json:"bank_references_address_3"`
	Bank_References_Phone1			string	`json:"bank_references_phone_1"`
	Bank_References_Phone2			string	`json:"bank_references_phone_2"`
	Bank_References_Phone3			string	`json:"bank_references_phone_3"`
	Other_Collected_Proceeds1		string	`json:"other_collected_proceeds_1"`
	Other_Collected_Proceeds2		string	`json:"other_collected_proceeds_2"`
	Created_at						string	`json:"created_at"`
	Updated_at						string	`json:"updated_at"`
}
type Sys_Domains []Sys_Domain
type Users_Bonding []Usr_Bonding

func GetAll(db *sql.DB) (*Users_Bonding, error){

	query := fmt.Sprintf(`SELECT name, lastname,type_doc,num_doc,nationality,sex,place_birth,day_birth,status_civil,tel_resi,
								  tel_cel,tel_ofi,dir_resi,city,cod_postal,nom_business,type_business,cargo,anos_business, dir_business,
	                              city_business,resident,num_security,number_fax,email,correspondence,confirmations,account_num1,account_num2,account_num3,
	                              account_type1,account_type2,account_type3,institution1,institution2,institution3,custodian_account,institution_other,knowledge,market_Cycle,academic_level,annual_income,
	                              total_equity,liquid_assets,investment_objectives,tolerance_risk,total_patrimony1,total_patrimony2, total_assets1, total_assets2,source_funds,attachments,
	                              bank_references_institution1,bank_references_institution2,bank_references_institution3,bank_references_address1,
	                              bank_references_address2,bank_references_address3,bank_references_phone1,bank_references_phone2, bank_references_phone3,other_collected_proceeds1,other_collected_proceeds2,created_at, updated_at
								  FROM usr_bonding`)

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	data := Users_Bonding{}
	for rows.Next(){
		regist := Usr_Bonding{}
		err := rows.Scan(
			&regist.Name,
			&regist.Lastname,
			&regist.Type_Doc,
			&regist.Num_Doc,
			&regist.Nationality,
			&regist.Sex,
			&regist.Place_Birth,
			&regist.Day_Birth,
			&regist.Status_Civil,
			&regist.Tel_Resi,
			&regist.Tel_Cel,
			&regist.Tel_ofi,
			&regist.Dir_Resi,
			&regist.City,
			&regist.Cod_Postal,
			&regist.Nom_Business,
			&regist.Type_Business,
			&regist.Cargo,
			&regist.Anos_Business,
			&regist.Dir_Business,
			&regist.City_Business,
			&regist.Resident,
			&regist.Num_Security,
			&regist.Number_Fax,
			&regist.Email,
			&regist.Correspondence,
			&regist.Confirmations,
			&regist.Account_Num1,
			&regist.Account_Num2,
			&regist.Account_Num3,
			&regist.Account_Type1,
			&regist.Account_Type2,
			&regist.Account_Type3,
			&regist.Institution1,
			&regist.Institution2,
			&regist.Institution3,
			&regist.Custodian_Account,
			&regist.Institution_Other,
			&regist.Knowledge,
			&regist.Market_Cycle,
			&regist.Academic_Level,
			&regist.Annual_Income,
			&regist.Total_Equity,
			&regist.Liquid_Assets,
			&regist.Investment_Objectives,
			&regist.Tolerance_Risk,
			&regist.Total_Patrimony1,
			&regist.Total_Patrimony2,
			&regist.Total_Assets1,
			&regist.Total_Assets2,
			&regist.Source_Funds,
			&regist.Attachments,
			&regist.Bank_References_Institution1,
			&regist.Bank_References_Institution2,
			&regist.Bank_References_Institution3,
			&regist.Bank_References_Address1,
			&regist.Bank_References_Address2,
			&regist.Bank_References_Address3,
			&regist.Bank_References_Phone1,
			&regist.Bank_References_Phone2,
			&regist.Bank_References_Phone3,
			&regist.Other_Collected_Proceeds1,
			&regist.Other_Collected_Proceeds2,
			&regist.Created_at,
			&regist.Updated_at,
		)
		if err != nil {
			return nil, err
		}

		data = append(data, regist)
	}
	return &data, nil
}
func Insert(model *Usr_Bonding, tx *sql.Tx) (int64, error) {

	query := fmt.Sprintf(` INSERT INTO usr_bonding
								 (name, lastname,type_doc,num_doc,nationality,sex,place_birth,day_birth,status_civil,tel_resi,
								  tel_cel,tel_ofi,dir_resi,city,cod_postal,nom_business,type_business,cargo,anos_business, dir_business,
	                              city_business,resident,num_security,number_fax,email,correspondence,confirmations,account_num1,account_num2,account_num3,
	                              account_type1,account_type2,account_type3,institution1,institution2,institution3,custodian_account,institution_other,knowledge,market_Cycle,academic_level,annual_income,
	                              total_equity,liquid_assets,investment_objectives,tolerance_risk,total_patrimony1,total_patrimony2, total_assets1, total_assets2,
								  source_funds,attachments,bank_references_institution1,bank_references_institution2,bank_references_institution3,bank_references_address1,
	                              bank_references_address2,bank_references_address3,bank_references_phone1,bank_references_phone2, bank_references_phone3,other_collected_proceeds1,other_collected_proceeds2,created_at, updated_at)
						 		 values
								 ('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s',
	                              '%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s',
	                              '%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s',
	                              '%s','%s','%s',GETDATE(), GETDATE())`, model.Name, model.Lastname,model.Type_Doc,model.Num_Doc, model.Nationality,model.Sex,model.Place_Birth,model.Day_Birth,model.Status_Civil,
		                          model.Tel_Resi,model.Tel_Cel, model.Tel_ofi,model.Dir_Resi,model.City,model.Cod_Postal,model.Nom_Business,model.Type_Business,model.Cargo,model.Anos_Business,model.Dir_Business,
		                          model.City_Business,model.Resident,model.Num_Security,model.Number_Fax,model.Email,model.Correspondence,model.Confirmations,model.Account_Num1,model.Account_Num2,model.Account_Num3,
		                          model.Account_Type1,model.Account_Type2,model.Account_Type3,model.Institution1,model.Institution2,model.Institution3,model.Custodian_Account,model.Institution_Other,model.Knowledge,model.Market_Cycle,model.Academic_Level,
		                          model.Annual_Income,model.Total_Equity,model.Liquid_Assets,model.Investment_Objectives,
		                          model.Tolerance_Risk,model.Total_Patrimony1,model.Total_Patrimony2,model.Total_Assets1,model.Total_Assets2,model.Source_Funds,
		                          model.Attachments,model.Bank_References_Institution1,model.Bank_References_Institution2,
		                          model.Bank_References_Institution3,model.Bank_References_Address1,model.Bank_References_Address2,model.Bank_References_Address3,model.Bank_References_Phone1,model.Bank_References_Phone2,
							      model.Bank_References_Phone3,model.Other_Collected_Proceeds1,model.Other_Collected_Proceeds2)

	stmt, err := tx.Prepare(query)
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec()
	if err != nil {
		return 0, err
	}

	validate, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	if validate != 1 {
		return 0, err
	}

	id,err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
func ShowTypeDoc(db *sql.DB) (*Sys_Domains, error){
	query := fmt.Sprintf(`select id, description FROM sys_domain where name='TIPODOC' and id_dad='18' and status='1'`)
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	data := Sys_Domains{}
	for rows.Next(){
		regist := Sys_Domain{}
		err := rows.Scan(
			&regist.Id,
			&regist.Description,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, regist)
	}
	return &data, nil
}
func ShowSex(db *sql.DB) (*Sys_Domains, error){
	query := fmt.Sprintf(`select id, description FROM sys_domain where name='SEXO' and id_dad='8' and status='1'`)
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	data := Sys_Domains{}
	for rows.Next(){
		regist := Sys_Domain{}
		err := rows.Scan(
			&regist.Id,
			&regist.Description,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, regist)
	}
	return &data, nil
}
func ShowEstadoCivil(db *sql.DB) (*Sys_Domains, error){
	query := fmt.Sprintf(`select id, description FROM sys_domain where name='ESTADOCIVIL' and id_dad='11' and status='1'`)
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	data := Sys_Domains{}
	for rows.Next(){
		regist := Sys_Domain{}
		err := rows.Scan(
			&regist.Id,
			&regist.Description,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, regist)
	}
	return &data, nil
}