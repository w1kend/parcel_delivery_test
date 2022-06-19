//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Orders = newOrdersTable("public", "orders", "")

type ordersTable struct {
	postgres.Table

	//Columns
	ID                postgres.ColumnString
	FromAddr          postgres.ColumnString
	ToAddr            postgres.ColumnString
	Status            postgres.ColumnString
	Price             postgres.ColumnInteger
	SenderName        postgres.ColumnString
	SenderPassportNum postgres.ColumnString
	RecipientName     postgres.ColumnString
	Weight            postgres.ColumnInteger
	CreatedAt         postgres.ColumnTimestampz
	CreatedBy         postgres.ColumnString
	CourierID         postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type OrdersTable struct {
	ordersTable

	EXCLUDED ordersTable
}

// AS creates new OrdersTable with assigned alias
func (a OrdersTable) AS(alias string) *OrdersTable {
	return newOrdersTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new OrdersTable with assigned schema name
func (a OrdersTable) FromSchema(schemaName string) *OrdersTable {
	return newOrdersTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new OrdersTable with assigned table prefix
func (a OrdersTable) WithPrefix(prefix string) *OrdersTable {
	return newOrdersTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new OrdersTable with assigned table suffix
func (a OrdersTable) WithSuffix(suffix string) *OrdersTable {
	return newOrdersTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newOrdersTable(schemaName, tableName, alias string) *OrdersTable {
	return &OrdersTable{
		ordersTable: newOrdersTableImpl(schemaName, tableName, alias),
		EXCLUDED:    newOrdersTableImpl("", "excluded", ""),
	}
}

func newOrdersTableImpl(schemaName, tableName, alias string) ordersTable {
	var (
		IDColumn                = postgres.StringColumn("id")
		FromAddrColumn          = postgres.StringColumn("from_addr")
		ToAddrColumn            = postgres.StringColumn("to_addr")
		StatusColumn            = postgres.StringColumn("status")
		PriceColumn             = postgres.IntegerColumn("price")
		SenderNameColumn        = postgres.StringColumn("sender_name")
		SenderPassportNumColumn = postgres.StringColumn("sender_passport_num")
		RecipientNameColumn     = postgres.StringColumn("recipient_name")
		WeightColumn            = postgres.IntegerColumn("weight")
		CreatedAtColumn         = postgres.TimestampzColumn("created_at")
		CreatedByColumn         = postgres.StringColumn("created_by")
		CourierIDColumn         = postgres.StringColumn("courier_id")
		allColumns              = postgres.ColumnList{IDColumn, FromAddrColumn, ToAddrColumn, StatusColumn, PriceColumn, SenderNameColumn, SenderPassportNumColumn, RecipientNameColumn, WeightColumn, CreatedAtColumn, CreatedByColumn, CourierIDColumn}
		mutableColumns          = postgres.ColumnList{FromAddrColumn, ToAddrColumn, StatusColumn, PriceColumn, SenderNameColumn, SenderPassportNumColumn, RecipientNameColumn, WeightColumn, CreatedAtColumn, CreatedByColumn, CourierIDColumn}
	)

	return ordersTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                IDColumn,
		FromAddr:          FromAddrColumn,
		ToAddr:            ToAddrColumn,
		Status:            StatusColumn,
		Price:             PriceColumn,
		SenderName:        SenderNameColumn,
		SenderPassportNum: SenderPassportNumColumn,
		RecipientName:     RecipientNameColumn,
		Weight:            WeightColumn,
		CreatedAt:         CreatedAtColumn,
		CreatedBy:         CreatedByColumn,
		CourierID:         CourierIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
