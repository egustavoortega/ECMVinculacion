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
	"github.com/e-capture/ECMVinculacion/example"
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
	const (
		fontPtSize = 8.0
	)

	pdf := gofpdf.New("P", "mm", "A4", "")
	titleStr := "Ficha de Identificacion del Cliente Persona Fisica "
		pdf.SetHeaderFunc(func() {
			pdf.Image(example.ImageFile("ecapture.png"), 10, 6, 30, 0, false, "", 0, "")
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
			pdf.Image(example.ImageFile("itc.png"), 170, 6, 30, 0, false, "", 0, "")
			pdf.SetY(10)
			// Line break
			pdf.Ln(5)
		})
	pdf.SetFooterFunc(func() {
		pdf.SetY(-15)
		pdf.SetFont("Arial", "I", 8)
		pdf.CellFormat(0, 10, fmt.Sprintf("Page %d", pdf.PageNo()),
			"", 0, "C", false, 0, "")
	})
	pdf.AddPage()
	pdf.Ln(20)
	pdf.SetFont("Arial", "", fontPtSize)
	pdf.SetFillColor(38, 78, 114)
	pdf.SetTextColor(255, 255, 255)
	pdf.CellFormat(0, 8, "Datos Personales - Persona Fisica", "1", 1, "C", true, 0, "")
	pdf.Ln(1)
	pdf.SetTextColor(0, 0, 0)
	pdf.WriteAligned(0,35, "Apellido(s): "+model.Lastname, "")
	pdf.WriteAligned(0,35, "Nombre(s): "+model.Name, "R")
	pdf.Ln(10)
	pdf.WriteAligned(0,35,"Documento de Identificacion:"+model.Type_Doc, "")
	pdf.WriteAligned(0,35, "Numero: "+model.Num_Doc, "R")
	pdf.Ln(10)
	pdf.WriteAligned(0,35,"Nacionalidad:"+model.Nationality, "")
	pdf.WriteAligned(0,35, "Sexo: "+model.Sex, "R")
	pdf.Ln(10)
	pdf.WriteAligned(0,35,"Lugar de nacimiento:"+model.Place_Birth, "")
	pdf.WriteAligned(0,35, "Fecha de Nacimiento: "+model.Day_Birth, "R")
	pdf.Ln(10)
	pdf.WriteAligned(0,35,"Estado Civil:"+model.Status_Civil, "")
	pdf.WriteAligned(0,35, "Telefono de Residencia Numero: "+model.Tel_Resi, "R")
	pdf.Ln(10)
	pdf.WriteAligned(0,35,"Telefono Celular Num.:"+model.Tel_Cel, "")
	pdf.WriteAligned(0,35, "Telefono Oficina Num.: "+model.Tel_ofi, "R")
	pdf.Ln(10)
	pdf.WriteAligned(0,35,"Direccion de Residencia.:"+model.Dir_Resi, "")
	pdf.Ln(10)
	pdf.WriteAligned(0,35,"Ciudad/Pais:"+model.City, "")
	pdf.WriteAligned(0,35, "Codigo Postal RD: "+model.Cod_Postal, "R")
	pdf.Ln(10)
	pdf.WriteAligned(0,35,"Nombre de la Empresa donde labora"+model.Nom_Business, "")
	pdf.WriteAligned(0,35, "Tipo de Negocio: "+model.Type_Business, "R")
	pdf.Ln(10)
	pdf.WriteAligned(0,35,"Cargo:"+model.Cargo, "")
	pdf.WriteAligned(0,35, "Anos en la empresa: "+model.Anos_Business, "R")
	pdf.Ln(10)
	pdf.WriteAligned(0,35,"Direccion de Oficina"+model.Dir_Business, "")
	pdf.WriteAligned(0,35, "Ciudad/Pais: "+model.City_Business, "R")
	pdf.Ln(10)
	pdf.WriteAligned(0,35,"Ciudadano o Residente de EE.UU. :"+model.Resident, "")
	pdf.Ln(10)
	pdf.WriteAligned(0,35, "Si la respuesta es 'Si', indicar el numero de Seguridad Social : "+model.Num_Security, "")
	pdf.Ln(20)
	pdf.SetFont("Arial", "", fontPtSize)
	pdf.SetFillColor(38, 78, 114)
	pdf.SetTextColor(255, 255, 255)
	pdf.CellFormat(0, 8, "Comunicaciones y Correspondencia", "1", 1, "C", true, 0, "")
	pdf.SetTextColor(0, 0, 0)
	pdf.WriteAligned(0,35, "Fax Numero: "+model.Number_Fax, "")
	pdf.WriteAligned(0,35, "Correo Electronico : "+model.Email, "R")
	pdf.Ln(10)
	pdf.WriteAligned(0,35, "Lugar para entrega de correspondencia y estados de cuenta: "+model.Correspondence, "")
	pdf.Ln(10)
	pdf.WriteAligned(0,35, "Indique si acepta recibir confirmaciones y enviar instrucciones: "+model.Confirmations, "")
	pdf.Ln(10)
	pdf.WriteAligned(0,35, "Numero de Cuenta Bancaria : ", "")
	pdf.WriteAligned(0,35, "Tipo de Cuenta: ", "C")
	pdf.WriteAligned(0,35, "Institucion: ", "R")
	pdf.Ln(10)
	pdf.WriteAligned(0,35, ""+model.Account_Num1, "")
	pdf.WriteAligned(0,35, ""+model.Account_Type1, "C")
	pdf.WriteAligned(0,35, ""+model.Institution1, "R")
	pdf.Ln(5)
	pdf.WriteAligned(0,35, ""+model.Account_Num2, "")
	pdf.WriteAligned(0,35, ""+model.Account_Type2, "C")
	pdf.WriteAligned(0,35, ""+model.Institution2, "R")
	pdf.Ln(5)
	pdf.WriteAligned(0,35, ""+model.Account_Num3, "")
	pdf.WriteAligned(0,35, ""+model.Account_Type3, "C")
	pdf.WriteAligned(0,35, ""+model.Institution3, "R")
	pdf.Ln(5)
	pdf.WriteAligned(0,35, "Numero de Cuenta de Custodia de Valores: "+model.Custodian_Account, "")
	pdf.WriteAligned(0,35, "Institucion: "+model.Institution_Other, "R")
	pdf.Ln(20)
	pdf.SetFont("Arial", "", fontPtSize)
	pdf.SetFillColor(38, 78, 114)
	pdf.SetTextColor(255, 255, 255)
	pdf.CellFormat(0, 8, "Formulario Perfil del Inversionista", "1", 1, "C", true, 0, "")
	pdf.Ln(1)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Arial", "", fontPtSize)
	lineHt4 := pdf.PointConvert(fontPtSize)
	htmlStr4 :=`El proposito del siguiente cuestionario es informar a AFI Reservas, S.A., los objetivos de inversion del Cliente.  En ningún caso AFI Reservas, S.A.,  prestara al Cliente asesoria respecto de sus inversiones.`
	html4 := pdf.HTMLBasicNew()
	html4.Write(lineHt4, htmlStr4)
	pdf.Cell(1,9,"")
	pdf.Ln(1)
	pdf.WriteAligned(0,35, "Experiencia y Conocimiento como Inversionista: "+model.Knowledge, "")
	pdf.Ln(10)
	pdf.WriteAligned(0,35, "Obje. de Inversion en un Ciclo de Mercado: "+model.Market_Cycle, "")
	pdf.Ln(10)
	pdf.WriteAligned(0,35, "Nivel Academico: "+model.Academic_Level, "")
	pdf.Ln(20)
	pdf.WriteAligned(0,35, "Ingresos Anuales: "+model.Annual_Income, "")
	pdf.WriteAligned(0,35, "Total de Activos Liquidos (dinero e inversiones): "+model.Total_Equity, "C")
	pdf.WriteAligned(0,35, "Patrimonio Total: "+model.Liquid_Assets, "C")
	pdf.Ln(20)
	pdf.WriteAligned(0,35, "Objetivos de Inversion : "+model.Investment_Objectives, "")
	pdf.Ln(20)
	pdf.SetFont("Arial", "B", fontPtSize)
	pdf.SetFillColor(38, 78, 114)
	pdf.SetTextColor(255, 255, 255)
	pdf.CellFormat(0, 8, "Tolerancia al Riesgo del Cliente", "1", 1, "C", true, 0, "")
	pdf.Ln(1)
	pdf.SetTextColor(0, 0, 0)
	pdf.WriteAligned(0,35, " "+model.Tolerance_Risk, "")
	pdf.Ln(20)
	pdf.SetFont("Arial", "", fontPtSize)
	pdf.SetFillColor(38, 78, 114)
	pdf.SetTextColor(255, 255, 255)
	pdf.CellFormat(0, 8, "Analisis del Riesgo Patrimonial (Si existe algun cambio patrimonial importante entre el año anterior y el presente)", "1", 1, "C", true, 0, "")
	pdf.SetTextColor(0, 0, 0)
	pdf.WriteAligned(0,35, "Total Patrimonio", "")
	pdf.WriteAligned(0,35, "Total Activos Liquidos", "R")
	pdf.Ln(5)
	pdf.WriteAligned(0,35, ""+model.Total_Patrimony1, "")
	pdf.WriteAligned(0,35, ""+model.Total_Assets1, "R")
	pdf.Ln(5)
	pdf.WriteAligned(0,35, ""+model.Total_Patrimony2, "")
	pdf.WriteAligned(0,35, ""+model.Total_Assets2, "R")
	pdf.Ln(20)
	pdf.SetFont("Arial", "", fontPtSize)
	pdf.SetFillColor(38, 78, 114)
	pdf.SetTextColor(255, 255, 255)
	pdf.CellFormat(0, 8, "Declaracion Jurada de Origen y Destino Licito de Fondos ", "1", 1, "C", true, 0, "")
	pdf.Ln(5)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Arial", "", fontPtSize)
	lineHt := pdf.PointConvert(fontPtSize)
	htmlStr :=`Por medio del presente documento declaro (declaramos) bajo fe de juramento que el dinero, capitales, haberes, valores o titulos utilizados en los negocios realizados con AFI Reservas S.A., no tienen relacion con actividades producto del delito de legitimacion de capitales, y por lo tanto no guardan vinculacion ninguna con actividades ilicitas, asi como tampoco emanan de ningun delito previsto en cualquier ley penal vigente. Mediante la presente declaracion ademas se autoriza a AFI Reservas, S.A., para que verifique los datos suministrados y contenidos en esta planilla.`
	html := pdf.HTMLBasicNew()
	html.Write(lineHt, htmlStr)
	pdf.Ln(20)
	pdf.SetFont("Arial", "", fontPtSize)
	pdf.SetFillColor(38, 78, 114)
	pdf.SetTextColor(255, 255, 255)
	pdf.CellFormat(0, 8, "Origen de los Fondos", "1", 1, "C", true, 0, "")
	pdf.Ln(1)
	pdf.SetTextColor(0, 0, 0)
	pdf.WriteAligned(0,10, ""+model.Source_Funds, "")
	pdf.Ln(20)
	pdf.Cell(0,9,"")
	pdf.Ln(20)
	pdf.SetFont("Arial", "B", fontPtSize)
	pdf.SetFillColor(38, 78, 114)
	pdf.SetTextColor(255, 255, 255)
	pdf.CellFormat(0, 8, "Declaracion Jurada de Personas Expuestas Politicamente (PEP'S)", "1", 1, "C", true, 0, "")
	pdf.Ln(1)
	pdf.SetTextColor(0, 0, 0)
	ht := pdf.PointConvert(fontPtSize)
	write := func(str string) {
		pdf.CellFormat(0, ht, str, "", 1, "", false, 0, "")
		pdf.Ln(ht)
	}
	pdf.Ln(2 * ht)
	write("Por medio del presente documento declaro (declaramos) lo siguiente:")
	write("1: Ocupo un cargo publico por eleccion popular. SI____   NO____.")
	write("2: Mantengo una relacion de empleo con organos o entidades del Estado bajo las disposiciones del Codigo de Trabajo, ocupando una")
	write("   posicion de gerencia o direccion. SI___   NO___.")
	write("3: Ocupo un cargo de direccion o gerencia dentro de la Administracion Publica Central  o Descentralizada  o de un Ayuntamiento por ")
	write("   designacion de la autoridad competente. SI___   NO___.")
	write("4: Que en virtud de lo anterior ostento u ostente el cargo de ______________________________________________________")
	write("5: Formo parte de las Fuerzas Armadas o de la Policia Nacional, con un rango de Capitan, Mayor, Coronel, General, Contralmirante, Mayor General,")
	write("   Vicealmirante, Almirante, Contralmirante o Teniente General . SI___   NO___. ")
	write("6: Formo parte del organo de direccion de un partido politico. SI___   NO___.")
	write("7: En los ultimos ocho (8) anos he ocupado algunas de las funciones o cargos enumerados anteriormente. SI___   NO___.")
	write("8: Mi conyuge cumple con algunas de las caracteristicas enumeradas anteriormente. SI___   NO___.")
	write("9: Me encuentro vinculado por parentesco por consanguinidad o afinidad hasta el 1er grado___    2do grado___ 3er grado___ con una")
	write("   persona que cumpla las caracteristicas enumeradas anteriormente. SI___   NO___.")
	write("10: A mi juicio cumplo con los requisitos para ser considerado como Persona Politicamente Expuesta. SI___   NO___.")

	pdf.Ln(10)
	pdf.SetFont("Arial", "", fontPtSize)
	pdf.SetFillColor(38, 78, 114)
	pdf.SetTextColor(255, 255, 255)
	pdf.CellFormat(0, 8, "Declaracion Jurada sobre Informaciones Suministradas", "1", 1, "C", true, 0, "")
	pdf.Ln(1)
	pdf.SetTextColor(0, 0, 0)
	lineHt2:= pdf.PointConvert(fontPtSize)
	htmlStr2 :=`EL CLIENTE reconoce y declara que: (i) la informacion suministrada a AFI RESERVAS, S.A., sobre su ciudadania, estatus de residencia y domicilio son veraces, por tanto, en caso de ser ciudadano o residente en un pais extranjero, asi como, que a juicio de cualquier oficial de AFI RESERVAS, S.A., muestre cualquier indicio de que tal fuera el caso, EL CLIENTE se compromete a facilitar el llenado de los formularios correspondientes en caso de que aplique; (ii) informara por escrito a AFI RESERVAS S.A., de cualquier cambio que ocurra en su estatus migratorio, ciudadania o residencia ante cualquier pais extranjero, y facilitara el llenado de los formularios correspondientes en los casos que aplique, en un plazo no mayor a treinta (30) dias calendarios contados a partir de la ocurrencia de este; (iii) autoriza a AFI RESERVAS, S.A., a compartir la informacion suministrada a requerimiento de entidades regulatorias y tributarias competentes en cumplimiento de las leyes vigentes en la Republica Dominicana y en los casos que aplique con el Internal Revenue Services (por sus siglas en ingles IRS) y/o cualquier otra autoridad norteamericana competente, en virtud de la ley Foreign Account Tax Compliance  Act (por sus siglas en ingles FATCA); (iv) descarga a AFI RESERVAS, S.A., de cualquier responsabilidad por la entrega y manejo de la informacion suministrada; y (v) autoriza a AFI RESERVAS, S.A., a realizar las indagatorias correspondientes, para verificar y confirmar las informaciones por el suministradas, incluyendo pero no limitado a indagatorias en los Buros de Credito y cualquier base de datos publicas o privadas.`
	html2 := pdf.HTMLBasicNew()
	html2.Write(lineHt2, htmlStr2)

	pdf.Ln(10)
	pdf.SetFont("Arial", "", fontPtSize)
	pdf.SetFillColor(38, 78, 114)
	pdf.SetTextColor(255, 255, 255)
	pdf.CellFormat(0, 8, "Tarjeta de Firma", "1", 1, "C", true, 0, "")
	pdf.Ln(5)
	pdf.SetTextColor(0, 0, 0)
	lineHt3:= pdf.PointConvert(fontPtSize)
	htmlStr3 :=`El(los) firmante(s) (el "Cliente"") declara haber leido, comprendido y expresamente acepta los Terminos y Condiciones del Contrato de Participacion o Suscripcion de Cuotas de AFI Reservas, S.A. ("El Contrato").  El Cliente ademas ratifica su intencion de someter los contratos y operaciones celebradas con AFI Reservas, S.A. a arbitraje institucional de conformidad con lo dispuesto en el Contrato.`
	html3 := pdf.HTMLBasicNew()
	html3.Write(lineHt3, htmlStr3)

	pdf.Ln(10)
	pdf.SetFont("Arial", "", fontPtSize)
	pdf.SetFillColor(38, 78, 114)
	pdf.SetTextColor(255, 255, 255)
	pdf.CellFormat(0, 8, "Recaudos Anexos", "1", 1, "C", true, 0, "")
	pdf.Ln(1)
	pdf.SetTextColor(0, 0, 0)
	pdf.Cell(0,6,"El Cliente anexa y suministra copia fiel a la AFI Reservas, S.A., de los  recaudos siguientes:")
	pdf.WriteAligned(0,10, ""+model.Attachments, "")
	pdf.Ln(20)
	pdf.SetFont("Arial", "", fontPtSize)
	pdf.Ln(20)
	pdf.Cell(0,9,"")
	pdf.Ln(20)
	pdf.SetFont("Arial", "B", fontPtSize)
	pdf.SetFillColor(38, 78, 114)
	pdf.SetTextColor(255, 255, 255)
	pdf.CellFormat(0, 8, "Referencias Bancarias", "1", 1, "C", true, 0, "")
	pdf.Ln(1)
	pdf.SetTextColor(0, 0, 0)
	pdf.WriteAligned(0,10, "Institucion", "")
	pdf.WriteAligned(0,10, "Direccion", "C")
	pdf.WriteAligned(0,10, "Telefono y Contacto", "R")
	pdf.Ln(5)
	pdf.WriteAligned(0,10, ""+model.Bank_References_Institution1, "")
	pdf.WriteAligned(0,10, ""+model.Bank_References_Address1, "C")
	pdf.WriteAligned(0,10, ""+model.Bank_References_Phone1, "R")
	pdf.Ln(5)
	pdf.WriteAligned(0,10, ""+model.Bank_References_Institution2, "")
	pdf.WriteAligned(0,10, ""+model.Bank_References_Address2, "C")
	pdf.WriteAligned(0,10, ""+model.Bank_References_Phone2, "R")
	pdf.Ln(5)
	pdf.WriteAligned(0,10, ""+model.Bank_References_Institution3, "")
	pdf.WriteAligned(0,10, ""+model.Bank_References_Address3, "C")
	pdf.WriteAligned(0,10, ""+model.Bank_References_Phone3, "R")
	pdf.Ln(20)
	pdf.SetFont("Arial", "", fontPtSize)
	pdf.SetFillColor(38, 78, 114)
	pdf.SetTextColor(255, 255, 255)
	pdf.CellFormat(0, 8, "Otros Recaudados Presentados", "1", 1, "C", true, 0, "")
	pdf.Ln(1)
	pdf.SetTextColor(0, 0, 0)
	pdf.WriteAligned(0,10, ""+model.Other_Collected_Proceeds1, "")
	pdf.Ln(10)
	pdf.WriteAligned(0,10, ""+model.Other_Collected_Proceeds2,"")
	pdf.Ln(20)
	pdf.SetFont("Arial", "", fontPtSize)
	pdf.SetFillColor(38, 78, 114)
	pdf.SetTextColor(255, 255, 255)
	pdf.CellFormat(0, 8, "Por AFI Reservas, S.A.", "1", 1, "C", true, 0, "")
	pdf.Ln(1)
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(0,8,"Original: Gerencia Administrativa y de Operaciones, Copia: Gerencia de Negocios, Oficial de Cumplimiento.", "0", 0, "R", false, 0, "")
	pdf.Ln(10)
	pdf.SetFont("Arial", "", 5)
	pdf.Cell(0,8,"AFI-006                     01-2016                        RNC 1-31-37781-5  ")

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
	err = pdf.OutputFileAndClose("Vinculacion.pdf")
	if err != nil {
		txt := fmt.Sprintf("Ha ocurrido un erro en la generacion del PDF: %s",err.Error())
		log.Write(txt)
		return err
	}

	response := response.Response{}
	response.Error = "false"
	response.Data = strconv.FormatInt(result, 10)
	response.Message = "Registro insertado exitosamente!"


	return c.JSON(http.StatusOK, &response)
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
