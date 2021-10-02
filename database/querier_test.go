package database

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

type Sample struct {
	id     int
	field1 string
	field2 string
}

func TestSampleQuery(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	mockedRows := sqlmock.NewRows([]string{"id", "field1", "field2"}).AddRow(1, "sample11", "sample12").AddRow(2, "sample21", "sample22")
	mock.ExpectQuery("^SELECT (.+) FROM bills$").WillReturnRows(mockedRows)

	actualRows, err := db.Query("SELECT * FROM bills")
	if err != nil {
		t.Fatal(err)
	}

	var data []Sample

	for actualRows.Next() {
		var id int
		var field1 string
		var field2 string

		_ = actualRows.Scan(&id, &field1, &field2)
		datum := Sample{
			id:     id,
			field1: field1,
			field2: field2,
		}

		data = append(data, datum)
	}

	assert.Equal(t, len(data), 2)
	assert.Equal(t, data[0].field2, "sample12")
	assert.Equal(t, data[1].field1, "sample21")
}

func TestGetAccountIdFromName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	mockAxisRow := sqlmock.NewRows([]string{"id"}).AddRow(19)
	mock.ExpectQuery("Axis").WillReturnRows(mockAxisRow)
	axisId := GetAccountIdFromName(db, "Axis Bank", 1)

	assert.Equal(t, axisId, 19)
}

func TestAbstractClass(t *testing.T) {
	// db, err := getConnection()
	// if err != nil {
	// 	return nil, err
	// }

	// ExecuteQuery("select id from accounts where name = \"Axis Bank\" and ")

	// assert.Equal(t, hdfcAccount.Name, "HDFC Bank")
	// assert.Equal(t, hdfcAccount.TransactionDetails.Type, DEBIT)
}
