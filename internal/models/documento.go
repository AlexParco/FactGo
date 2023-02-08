package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
)

// Modelo relacionado a las estructuras de las Facturas, Boletas, etc
//

type Documento struct {
	Cia                        sql.NullString `json:"cia,omitempty"`
	IdPlanta                   sql.NullString `json:"id_planta,omitempty"`
	IdPlantaDireccion          sql.NullString `json:"id_planta_direccion,omitempty"`
	IdTipoDoc                  sql.NullString `json:"id_tipo_doc,omitempty"`
	Serie                      sql.NullString `json:"serie,omitempty"`
	NroDocumento               sql.NullString `json:"nro_documento,omitempty"`
	SubTotal                   sql.NullString `json:"sub_total,omitempty"`
	Fecha                      sql.NullString `json:"fecha,omitempty"`
	Hora                       sql.NullString `json:"hora,omitempty"`
	IdMonedaDoc                sql.NullString `json:"id_moneda_doc,omitempty"`
	NroDi                      sql.NullString `json:"nro_di,omitempty"`
	DescripcionCliente         sql.NullString `json:"descripcion_cliente,omitempty"`
	IdDistrito                 sql.NullString `json:"id_distrito,omitempty"`
	DireccionCliente           sql.NullString `json:"direccion_cliente,omitempty"`
	DepartamentoCliente        sql.NullString `json:"departamento_cliente,omitempty"`
	ProvinciaCliente           sql.NullString `json:"provincia_cliente,omitempty"`
	DistritoCliente            sql.NullString `json:"distrito_cliente,omitempty"`
	RazonSocialCliente         sql.NullString `json:"razon_social_cliente,omitempty"`
	EmailCliente               sql.NullString `json:"email_cliente,omitempty"`
	FechaVence                 sql.NullString `json:"fecha_vence,omitempty"`
	Igv                        sql.NullString `json:"igv,omitempty"`
	PorcIgv                    sql.NullString `json:"porc_igv,omitempty"`
	Total                      sql.NullString `json:"total,omitempty"`
	MonedaDescripcion          sql.NullString `json:"moneda_descripcion,omitempty"`
	IdTipoDocRef               sql.NullString `json:"id_tipo_doc_ref,omitempty"`
	SerieRef                   sql.NullString `json:"serie_ref,omitempty"`
	NroDocumentoRef            sql.NullString `json:"nro_documento_ref,omitempty"`
	IdMotivoTipoSunat          sql.NullString `json:"id_motivo_tipo_sunat,omitempty"`
	DescripcionNota            sql.NullString `json:"descripcion_nota,omitempty"`
	MontoPer                   sql.NullString `json:"monto_per,omitempty"`
	TotalMontoPer              sql.NullString `json:"total_monto_per,omitempty"`
	RazonSocialTransporte      sql.NullString `json:"razon_social_transporte,omitempty"`
	RucEmpTransporte           sql.NullString `json:"ruc_emp_transporte,omitempty"`
	ChoferDescrip              sql.NullString `json:"chofer_descrip,omitempty"`
	IdChofer                   sql.NullString `json:"id_chofer,omitempty"`
	PlacaTractor               sql.NullString `json:"placa_tractor,omitempty"`
	NroAutorizacion            sql.NullString `json:"nro_autorizacion,omitempty"`
	DireccionOrigen            sql.NullString `json:"direccion_origen,omitempty"`
	DepartOrigen               sql.NullString `json:"depart_origen,omitempty"`
	ProvinOrigen               sql.NullString `json:"provin_origen,omitempty"`
	DistriOrigen               sql.NullString `json:"distri_origen,omitempty"`
	UbigeoOrigen               sql.NullString `json:"ubigeo_origen,omitempty"`
	PlantaOrigen               sql.NullString `json:"planta_origen,omitempty"`
	DireccionDestino           sql.NullString `json:"direccion_destino,omitempty"`
	DistriDestino              sql.NullString `json:"distri_destino,omitempty"`
	ProvinDestino              sql.NullString `json:"provin_destino,omitempty"`
	DepartDestino              sql.NullString `json:"depart_destino,omitempty"`
	UbigeoDestino              sql.NullString `json:"ubigeo_destino,omitempty"`
	CondicionPago              sql.NullString `json:"condicion_pago,omitempty"`
	NrpScop                    sql.NullString `json:"nrp_scop,omitempty"`
	PlacaCisterna              sql.NullString `json:"placa_cisterna,omitempty"`
	NroCubicacion              sql.NullString `json:"nro_cubicacion,omitempty"`
	MontoDescuento             sql.NullString `json:"monto_descuento,omitempty"`
	FechaSistema               sql.NullString `json:"fecha_sistema,omitempty"`
	PorcPercepcion             sql.NullString `json:"porc_percepcion,omitempty"`
	IdCliente                  sql.NullString `json:"id_cliente,omitempty"`
	IdTipoDocCredito           sql.NullString `json:"id_tipo_doc_credito,omitempty"`
	NroPago                    sql.NullString `json:"nro_pago,omitempty"`
	TotalPeso                  sql.NullString `json:"total_peso,omitempty"`
	FechaActual                sql.NullString `json:"fecha_actual,omitempty"`
	CantidadItem               sql.NullString `json:"cantidad_item,omitempty"`
	MotivoSunat                sql.NullString `json:"motivo_sunat,omitempty"`
	OrdenCompra                sql.NullString `json:"orden_compra,omitempty"`
	IdPlantaRef                sql.NullString `json:"id_planta_ref,omitempty"`
	IdCondicionPago            sql.NullString `json:"id_condicion_pago,omitempty"`
	PorcNosRetiene             sql.NullString `json:"porc_nos_retiene,omitempty"`
	FlagEsAgenteRetenedor      sql.NullString `json:"flag_es_agente_retenedor,omitempty"`
	MontoRetenedor             sql.NullString `json:"monto_retenedor,omitempty"`
	MontoPorcRetencion         sql.NullString `json:"monto_porc_retencion,omitempty"`
	IdClasificacionTipoNegocio sql.NullString `json:"id_clasificacion_tipo_negocio,omitempty"`
	FlagCondicionPagoCredito   sql.NullString `json:"flag_condicion_pago_credito,omitempty"`
	MontoAplicaDebitoRef       sql.NullString `json:"monto_aplica_debito_ref,omitempty"`
	PerCorrelativoCancelacion  sql.NullString `json:"per_correlativo_cancelacion,omitempty"`
}

func (b *Documento) PrintStruct() {
	docJson, err := json.MarshalIndent(b, "", " ")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Documento: ", string(docJson))
}
