package models

import (
	"encoding/json"
	"fmt"
	"log"
)

type Baja struct {
	Cia          string `json:"cia,omitempty"`
	IDPlanta     string `json:"id_planta,omitempty"`
	IdTipoDoc    string `json:"id_tipo_doc,omitempty"`
	Serie        string `json:"serie,omitempty"`
	NroDocumento string `json:"nro_documento,omitempty"`
	Fecha        string `json:"fecha,omitempty"`
	Fechaactual  string `json:"fechaactual,omitempty"`
}

func (b *Baja) PrintStruct() {
	docJson, err := json.MarshalIndent(b, "", " ")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Baja: ", string(docJson))
}
