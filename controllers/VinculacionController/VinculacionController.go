package VinculacionController

import (
	"github.com/labstack/echo"
	"fmt"
	"github.com/e-capture/ECMVinculacion/structurs/response"
	"github.com/e-capture/MAJOSystem/db"
	"github.com/e-capture/ECMVinculacion/models/vinculacion"
	"github.com/e-capture/MAJOSystem/log"
	"net/http"
	"strconv"
	"github.com/e-capture/ECMVinculacion/structurs/request"
	"encoding/json"
	"github.com/jung-kurt/gofpdf"
	"golang.org/ejemplo/example"
	"io/ioutil"
)
func Index(c echo.Context) error {

	db, err := db.Open()
	defer db.Close()


	result, err := vinculacion.GetAll(db)
	if err != nil {
		text := fmt.Sprintf("Error DB, error en la consulta a la base de datos en la tabla usr_bonding, %s: ", err)
		log.Write(text)
		return err
	}

	response := response.ResponseVinculacionCreate{}
	response.Error = "false"
	response.Data = result
	response.Message = "La consulta se realizo exitosamente!"

	return c.JSON(http.StatusOK, &response)
}
func Create(c echo.Context) error {

	data := request.RequestCreateVinculacion{}

	defer c.Request().Body.Close()
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(b))
	err = json.Unmarshal(b,&data)
	if err != nil{
		fmt.Println(err)
	}

	db, err := db.Open()
	if err != nil {
		text := fmt.Sprintf("Error DB: ha ocurrido un error al concetar a la base de datos: %s", err)
		log.Write(text)
		return nil
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		text := fmt.Sprintf("Error DB: ha ocurrido un error al crear la transaccion en la base de datos: %s", err)
		log.Write(text)
		return nil
	}

	model := vinculacion.Usr_Bonding{}

	model.Name 							= data.Name
	model.Lastname 						= data.Lastname
	model.Type_Doc 						= data.Type_Doc
	model.Num_Doc 						= data.Num_Doc
	model.Nationality 					= data.Nationality
	model.Sex 							= data.Sex
	model.Place_Birth 					= data.Place_Birth
	model.Day_Birth 					= data.Day_Birth
	model.Status_Civil 					= data.Status_Civil
	model.Tel_Resi 						= data.Tel_Resi
	model.Tel_Cel 						= data.Tel_Cel
	model.Tel_ofi 						= data.Tel_ofi
	model.Dir_Resi 						= data.Dir_Resi
	model.City 							= data.City
	model.Cod_Postal 					= data.Cod_Postal
	model.Nom_Business 					= data.Nom_Business
	model.Type_Business 				= data.Type_Business
	model.Cargo 						= data.Cargo
	model.Anos_Business 				= data.Anos_Business
	model.Dir_Business 				    = data.Dir_Business
	model.City_Business 				= data.City_Business
	model.Resident   					= data.Resident
	model.Num_Security 					= data.Num_Security
	model.Number_Fax 					= data.Number_Fax
	model.Email 						= data.Email
	model.Correspondence				= data.Correspondence
	model.Confirmations					= data.Confirmations
	model.Account_Num1					= data.Account_Num1
	model.Account_Num2					= data.Account_Num2
	model.Account_Num3					= data.Account_Num3
	model.Account_Type1					= data.Account_Type1
	model.Account_Type2					= data.Account_Type2
	model.Account_Type3					= data.Account_Type3
	model.Institution1					= data.Institution1
	model.Institution2					= data.Institution2
	model.Institution3					= data.Institution3
	model.Custodian_Account				= data.Custodian_Account
	model.Institution_Other				= data.Institution_Other
	model.Knowledge						= data.Knowledge
	model.Market_Cycle					= data.Market_Cycle
	model.Academic_Level				= data.Academic_Level
	model.Annual_Income					= data.Annual_Income
	model.Total_Equity					= data.Total_Equity
	model.Liquid_Assets					= data.Liquid_Assets
	model.Investment_Objectives			= data.Investment_Objectives
	model.Tolerance_Risk				= data.Tolerance_Risk
	model.Total_Patrimony1				= data.Total_Patrimony1
	model.Total_Patrimony2				= data.Total_Patrimony2
	model.Total_Assets1					= data.Total_Assets1
	model.Total_Assets2					= data.Total_Assets2
	model.Source_Funds					= data.Source_Funds
	model.Attachments					= data.Attachments
	model.Bank_References_Institution1	= data.Bank_References_Institution1
	model.Bank_References_Institution2	= data.Bank_References_Institution2
	model.Bank_References_Institution3	= data.Bank_References_Institution3
	model.Bank_References_Address1		= data.Bank_References_Address1
	model.Bank_References_Address2		= data.Bank_References_Address2
	model.Bank_References_Address3		= data.Bank_References_Address3
	model.Bank_References_Phone1		= data.Bank_References_Phone1
	model.Bank_References_Phone2		= data.Bank_References_Phone2
	model.Bank_References_Phone3		= data.Bank_References_Phone3
	model.Other_Collected_Proceeds1		= data.Other_Collected_Proceeds1
	model.Other_Collected_Proceeds2		= data.Other_Collected_Proceeds2

	result, err := vinculacion.Insert(&model, tx)
	if err != nil {
		tx.Rollback()
		text := fmt.Sprintf("Error DB: ha ocurrido un error en el insert de la tabla usr_bonding en el metodo Create: %s", err)
		log.Write(text)
		return err
	}

	tx.Commit()
	plantilla1 := fmt.Sprintf(`Apellido(s): %s					Nombre(s) %s
	                                  Documento de Identificacion: %s	Numero: %s
									  Nacionalidad: %s					Sexo: %s
	                                  Lugar de Nacimiento: %s			Fecha de Nacimiento: %s
	                                  Estado Civil: %s					Telefono de Recidencia Número: %s
	                                  Telefono Celular Número: %s		Telefono Oficina Número: %s
	                                  Direccion de Residencia %s
			                          Ciudad/País: %s					Código Postal RD: %s
	                                  Nombre de la Empresa labora: %s	Tipo de Negocio: %s
	                                  Cargo: %s							Años en la Empresa: %s
	                                  Direccion de Oficina: %s
	                                  Ciudad/pais: %s					Ciudadano o Residente de EE.UU: %s
	                                  Si la Respuesta es "Si", Indicar el Numero de Seguridad Social: %s
	                                  `,model.Lastname,model.Name,model.Type_Doc,model.Num_Doc,model.Nationality,model.Sex,model.Place_Birth,
										model.Day_Birth,model.Status_Civil,model.Tel_Resi,model.Tel_Cel,model.Tel_ofi,model.Dir_Resi,model.City,model.Cod_Postal,
										model.Nom_Business,model.Type_Business,model.Cargo,model.Anos_Business,model.Dir_Business,model.City_Business,model.Resident,
										model.Num_Security)

	plantilla2 := fmt.Sprintf(`Fax Numero: %s
	                                  Correo Electronico: %s
									  Lugar para entrega de correspondencia y estados de cuenta: %s
	                                  Indique si acepta recibir confirmaciones y enviar instrucciones: %s
	                                  Número de Cuenta Bancaria: %s
	                                  Tipo de Cuenta: %s
			                          Institución: %s
	                                  `,model.Number_Fax,model.Email,model.Correspondence,model.Confirmations,model.Account_Num1,
										model.Account_Type1,model.Institution1)
	GeneratedPdf(plantilla1,plantilla2)
	response := response.Response{}
	response.Error = "false"
	response.Data = strconv.FormatInt(result, 10)
	response.Message = "Registro insertado exitosamente!"


	return c.JSON(http.StatusOK, &response)
}


func GeneratedPdf (data string, data2 string){
	pdf := gofpdf.New("P", "mm", "A4", "")
	titleStr := "Ficha de Identificacion del Cliente Persona Fisica "
	pdf.SetTitle(titleStr, false)
	pdf.SetAuthor("Jules Verne", false)

	pdf.SetHeaderFunc(func() {
		pdf.Image(example.ImageFile("logo.png"), 10, 6, 30, 0, false, "", 0, "")
		pdf.SetY(5)
		// Arial bold 12
		pdf.SetFont("Arial", "B", 12)
		// Calculate width of title and position
		wd := pdf.GetStringWidth(titleStr) + 6
		pdf.SetX((210 - wd) / 2)
		// Colors of frame, background and text
		pdf.SetDrawColor(255, 250, 250)
		pdf.SetFillColor(250, 250, 255)
		pdf.SetTextColor(0, 0, 0)
		// Thickness of frame (1 mm)
		pdf.SetLineWidth(1)
		// Title
		pdf.CellFormat(wd, 9, titleStr, "1", 1, "J", true, 0, "")
		// Line break
		pdf.Ln(20)
	})
	pdf.SetFooterFunc(func() {
		// Position at 1.5 cm from bottom
		pdf.SetY(-15)
		// Arial italic 8
		pdf.SetFont("Arial", "I", 8)
		// Text color in gray
		pdf.SetTextColor(38, 78, 114)
		// Page number
		pdf.CellFormat(0, 10, fmt.Sprintf("Page %d", pdf.PageNo()),
			"", 0, "C", false, 0, "")
	})
	chapterTitle := func(chapNum int, titleStr string) {
		// 	// Arial 9
		pdf.SetFont("Arial", "", 9)
		// Background color
		pdf.SetFillColor(200, 220, 255)
		// Title
		pdf.CellFormat(0, 6, fmt.Sprintf(" %d  %s", chapNum, titleStr),
			"1", 1, "L", true, 0, "")
		// Line break
		pdf.Ln(4)
	}
	chapterBody := func(fileStr string) {
		// Read text file
		txtStr, err := ioutil.ReadFile(fileStr)
		if err != nil {
			pdf.SetError(err)
		}
		// Times 12
		pdf.SetFont("Times", "", 10)
		// Output justified text
		pdf.MultiCell(0, 5, string(txtStr), "", "", false)
		// Line break
		pdf.Ln(-1)
		// Mention in italics
		pdf.SetFont("", "I", 0)
	}
	printChapter := func(chapNum int, titleStr, fileStr string) {
		pdf.AddPage()
		chapterTitle(chapNum, titleStr)
		chapterBody(fileStr)
	}
	printChapter(1, "Datos Personales - Persona Fisica",data)
	printChapter(2, "Datos Personales - Persona Fisica",data2)

	fileStr := example.Filename("Vinculacion")
	err := pdf.OutputFileAndClose(fileStr)
	example.Summary(err, fileStr)

}

func ShowTypeDoc (c echo.Context) error {
	db, err := db.Open()
	if err != nil {
		text := fmt.Sprintf("Error DB, error en la consulta a la base de datos en la tabla usr_bonding, %s: ", err)
		log.Write(text)
		return err
	}

	defer db.Close()

	result, err := vinculacion.ShowTypeDoc(db)

	response := response.ResponseVinculacion{}
	response.Error = "false"
	response.Data = result
	response.Message = "La consulta se realizo exitosamente!"

	return c.JSON(http.StatusOK, &response)
}
func ShowSex (c echo.Context) error {
	db, err := db.Open()
	if err != nil {
		text := fmt.Sprintf("Error DB, error en la consulta a la base de datos en la tabla usr_bonding, %s: ", err)
		log.Write(text)
		return err
	}

	defer db.Close()

	result, err := vinculacion.ShowSex(db)

	response := response.ResponseVinculacion{}
	response.Error = "false"
	response.Data = result
	response.Message = "La consulta se realizo exitosamente!"

	return c.JSON(http.StatusOK, &response)
}
func ShowEstadoCivil (c echo.Context) error {
	db, err := db.Open()
	if err != nil {
		text := fmt.Sprintf("Error DB, error en la consulta a la base de datos en la tabla usr_bonding, %s: ", err)
		log.Write(text)
		return err
	}

	defer db.Close()

	result, err := vinculacion.ShowEstadoCivil(db)

	response := response.ResponseVinculacion{}
	response.Error = "false"
	response.Data = result
	response.Message = "La consulta se realizo exitosamente!"

	return c.JSON(http.StatusOK, &response)
}
