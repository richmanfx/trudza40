package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type SettingsTable_20190310_223038 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &SettingsTable_20190310_223038{}
	m.Created = "20190310_223038"
	_ = migration.Register("SettingsTable_20190310_223038", m)
}

// Run the migrations
func (m *SettingsTable_20190310_223038) Up() {
	m.SQL("CREATE TABLE public.settings(" +
		"user_id bigserial PRIMARY KEY, " +
		"settings_name varchar(255), " +
		"browser_width smallint, " +
		"browser_height smallint, " +
		"host_page_url varchar(255), " +
		"flash_quantity smallint, " +
		"flash_period smallint, " +
		"flash_allowed boolean, " +
		"debug_level varchar(25), " +
		"min_area smallint, " +
		"max_area smallint, " +
		"min_rental_period smallint, " +
		"property_type varchar(100), " +
		"contract_type varchar(100), " +
		"country varchar(100), " +
		"property_location varchar(100), " +
		"sort_field_name varchar(100), " +
		"average_rental money, " +
		"profit_months smallint, " +
		"prior_repair money, " +
		"contract_registration money, " +
		"running_cost money, " +
		"yearly_insurance money, " +
		"monthly_heating money, " +
		"housing_office_maintenance money, " +
		"accounting_service money, " +
		"required_profit_margin smallint" +
		")")

}

// Reverse the migrations
func (m *SettingsTable_20190310_223038) Down() {
	m.SQL("DROP TABLE settings")

}
