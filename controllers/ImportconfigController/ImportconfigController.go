package ImportconfigController

import (
	"github.com/e-capture/ECMConfiguration/models/importconfig"
	"github.com/labstack/echo"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

// Estructuras para retornar JSON
func Import(c echo.Context) error {
	dimport := importconfig.ImportAll()
	return c.JSON(http.StatusOK, dimport) //String(http.StatusOK, "name:" + name + ", email:" + email)
}


type Dataconfigimport struct {
	Process	  string `json:"process"`
	DocType_Id string `json:"doctype_id"`
}
func ImportConfig(c echo.Context) error {
	data := Dataconfigimport{}
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
	fmt.Println("Paramentro Process: ", data)
	// w.Header().Set("Content-Type", "application/json")
	dimport := importconfig.ImportConfig(data.Process,data.DocType_Id)
	// user.Users,_ = json.Marshal(&users)
	fmt.Println("Import")
	return c.JSON(http.StatusOK, dimport) //String(http.StatusOK, "name:" + name + ", email:" + email)
}
