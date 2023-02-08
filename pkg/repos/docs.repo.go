package repos

import (
	"log"
	"reflect"

	"github.com/alexparco/FactGo/internal/connection"
	cnst "github.com/alexparco/FactGo/internal/constants"
	"github.com/alexparco/FactGo/internal/models"
)

type DocRepo interface {
	FindAllDocs() ([]*models.Documento, error)
	FindAllBajas() ([]*models.Baja, error)
}

type docRepo struct {
	db *connection.Server
}

func NewDocRepo(db *connection.Server) DocRepo {
	return &docRepo{db}
}

// findAllDocs implements DocRepo
func (d *docRepo) FindAllDocs() ([]*models.Documento, error) {
	rows, err := d.db.Query(cnst.SelectQueryDocs)
	if err != nil {
		return nil, err
	}

	var documentos []*models.Documento

	for rows.Next() {
		var doc models.Documento

		values := reflect.ValueOf(&doc).Elem()

		nCols := values.NumField()

		cols := make([]interface{}, nCols)

		for i := 0; i < nCols; i++ {
			field := values.Field(i)
			cols[i] = field.Addr().Interface()
		}

		if err := rows.Scan(cols...); err != nil {
			log.Fatal(err)
		}

		doc.PrintStruct()

		documentos = append(documentos, &doc)
	}
	return documentos, nil
}

// findAllBajas implements DocRepo
func (d *docRepo) FindAllBajas() ([]*models.Baja, error) {
	var bajas []*models.Baja

	rows, err := d.db.Query(cnst.SelectQueryBajas)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var baja models.Baja

		values := reflect.ValueOf(&baja).Elem()

		nCols := values.NumField()

		cols := make([]interface{}, nCols)

		for i := 0; i < nCols; i++ {
			field := values.Field(i)
			cols[i] = field.Addr().Interface()
		}

		if err := rows.Scan(cols...); err != nil {
			log.Fatal(err)
			return nil, err
		}
		baja.PrintStruct()

		bajas = append(bajas, &baja)
	}
	return bajas, nil
}
