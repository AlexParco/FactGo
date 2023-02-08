package models

import (
	"encoding/xml"
	"log"
	"os"
)

type VoidedDocuments struct {
	Filename string `xml:"-"`

	// attrs
	XmlnsSac string `xml:"xmlns:sac,attr"`
	XmlnsCac string `xml:"xmlns:cac,attr"`
	XmlnsExt string `xml:"xmlns:ext,attr"`
	Xmlns    string `xml:"xmlns,attr"`
	XmlnsCbc string `xml:"xmlns:cbc,attr"`

	// childs
	UblVersionID            string                  `xml:"cbc:UBLVersionID"`
	CustomizationID         string                  `xml:"cbc:CustomizationID"`
	Id                      string                  `xml:"cbc:ID"`
	ReferenceDate           string                  `xml:"cbc:ReferenceDate"`
	IssueDate               string                  `xml:"cbc:IssueDate"`
	Signature               Signature               `xml:"cac:Signature"`
	AccountingSupplierParty AccountingSupplierParty `xml:"cac:AccountingSupplierParty"`
	VoidedDocumentsLine     VoidedDocumentsLine     `xml:"sac:VoidedDocumentsLine"`
}

type Signature struct {
	Id                         string                     `xml:"cbc:ID"`
	SignatoryParty             SignatoryParty             `xml:"cac:SignatoryParty"`
	DigitalSignatureAttachment DigitalSignatureAttachment `xml:"cac:DigitalSignatureAttachment"`
}

type SignatoryParty struct {
	PartyIdentification struct {
		Id string `xml:"cbc:ID"`
	} `xml:"cac:PartyIdentification"`

	PartyName struct {
		Name string `xml:"cbc:Name"`
	} `xml:"cac:PartyName"`
}

type DigitalSignatureAttachment struct {
	ExternalReference struct {
		URI string `xml:"cbc:URI"`
	} `xml:"cac:ExternalReference"`
}

type AccountingSupplierParty struct {
	CustomerAssignedAccountID string `xml:"cbc:CustomerAssignedAccountID"`
	AdditionalAccountID       string `xml:"cbc:AdditionalAccountID"`
	Party                     Party  `xml:"cac:Party"`
}

type Party struct {
	PartyName struct {
		Name string `xml:"cbc:Name"`
	} `xml:"cac:PartyName"`
	PostalAddress    PostalAddress `xml:"cac:PostalAddress"`
	PartyLegalEntity struct {
		RegistrationName string `xml:"cbc:RegistrationName"`
	} `xml:"cac:PartyLegalEntity"`
}

type PostalAddress struct {
	Id               string `xml:"cbc:ID"`
	StreetName       string `xml:"cbc:StreetName"`
	CityName         string `xml:"cbc:CityName"`
	CountrySubentity string `xml:"cbc:CountrySubentity"`
	District         string `xml:"cbc:District"`
	Country          struct {
		IdentificationCode string `xml:"cbc:IdentificationCode"`
	} `xml:"cac:Country"`
}

type VoidedDocumentsLine struct {
	LineID                string `xml:"cbc:LineID"`
	DocumentTypeCode      string `xml:"cbc:DocumentTypeCode"`
	DocumentSerialID      string `xml:"sac:DocumentSerialID"`
	DocumentNumberID      string `xml:"sac:DocumentNumberID"`
	VoidReasonDescription string `xml:"sac:VoidReasonDescription"`
}

func (baja *VoidedDocuments) BuildXml() {
	file, err := xml.MarshalIndent(baja, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	file = []byte(xml.Header + string(file))

	err = os.WriteFile(baja.Filename+".xml", file, 0644)
	if err != nil {
		log.Fatal(err.Error())
	}
}
