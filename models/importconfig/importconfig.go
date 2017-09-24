package importconfig

import (
	"fmt"
	"github.com/e-capture/MAJOSystem/db"
	"github.com/e-capture/MAJOSystem/log"
)



type Dataimportconfig struct {
	Id			int64 	`json:"id"`
	Name 			string	`json:"name"`
	Sweptdirectory		string 	`json:"sweptdirectory"`
	Backupdirectory		string	`json:"backupdirectory"`
	Separator		string	`json:"separator"`
	Extimport		string	`json:"extimport"`
	Frequency		int	`json:"frequency"`
	Updated_at		string	`json:"updated_at"`
}

type Dataimportconfigs []Dataimportconfig



// Metodo para obtener todos los procesos de Importacion configurados en el sistema
func ImportAll() Dataimportconfigs {
	db, err := db.Open()
	if err != nil {
		text := fmt.Sprintf("Error DB: ha ocurrido un error al concetar a la base de datos: %s", err)
		log.Write(text)
		return nil
	}
	defer db.Close()
	sqlImport := `SELECT id,name,sweptdirectory,backupdirectory,separator,extimport,frequency,updated_at
  		    FROM set_imports`
	importconfig := Dataimportconfigs{}
	stmt, err := db.Prepare(sqlImport)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		imp := Dataimportconfig{}
		err = rows.Scan(
			&imp.Id,
			&imp.Name,
			&imp.Sweptdirectory,
			&imp.Backupdirectory,
			&imp.Separator,
			&imp.Extimport,
			&imp.Frequency,
			&imp.Updated_at,
		)
		if err != nil {
			fmt.Println(err)
		}
		importconfig = append(importconfig, imp)
	}
	return importconfig
}

type Dataimportconf struct {
	Id			int64 	`json:"id"`
	Name 			string	`json:"name"`
	Sequences		int 	`json:"sequences"`
	Doctype_id		int64	`json:"doctype_id"`
	Keyword_id		int64	`json:"keyword_id"`
	Fieldtypes_id		int64  `json:"fieldtypes_id"`
}

type Dataimportconfs []Dataimportconf

func ImportConfig(proceso, doctype string) Dataimportconfs {
	fmt.Println("ingreso moledo import config: ",proceso)
	db, err := db.Open()
	if err != nil {
		text := fmt.Sprintf("Error DB: ha ocurrido un error al concetar a la base de datos: %s", err)
		log.Write(text)
		return nil
	}
	defer db.Close()
	sqlImport := fmt.Sprintf(`SELECT si.id, si.name,si.sequences,si.doctype_id,keyword_id ,k.fieldtypes_id
				FROM [dbo].[set_settingsimports] si join [dbo].[keywords] k ON si.keyword_id = k.id
				WHERE si.settingsimport_id = %s and si.doctype_id = %s
				ORDER BY si.doctype_id, si.sequences`,proceso,doctype)

	fmt.Println(sqlImport)
	importconfig := Dataimportconfs{}
	stmt, err := db.Prepare(sqlImport)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		imp := Dataimportconf{}
		err = rows.Scan(
			&imp.Id,
			&imp.Name,
			&imp.Sequences,
			&imp.Doctype_id,
			&imp.Keyword_id,
			&imp.Fieldtypes_id,
		)
		if err != nil {
			fmt.Println(err)
		}
		importconfig = append(importconfig, imp)
	}
	return importconfig
}