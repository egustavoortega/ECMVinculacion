package response

import (
	"github.com/e-capture/ECMVinculacion/models/vinculacion"
)

type Response struct {
	Error		string	`json:"error"`
	Data		string	`json:"data"`
	Message		string	`json:"message"`
}

type ModelVinculacion struct {
	Error		string						`json:"error"`
	Data		*vinculacion.Sys_Domain	`json:"data"`
	Message		string						`json:"message"`
}
type ResponseVinculacion struct {
	Error		string						`json:"error"`
	Data		*vinculacion.Sys_Domains	`json:"data"`
	Message		string						`json:"message"`
}
type ModelVinculacionCreate struct {
	Error		string						`json:"error"`
	Data		*vinculacion.Usr_Bonding	`json:"data"`
	Message		string						`json:"message"`
}
type ResponseVinculacionCreate struct {
	Error		string						`json:"error"`
	Data		*vinculacion.Users_Bonding	`json:"data"`
	Message		string						`json:"message"`
}