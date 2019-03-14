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
		"browser_width int, " +
		"browser_height int, " +
		"host_page_url varchar(255), " +
		"flash_quantity int, " +
		"flash_period int, " +
		"flash_allowed boolean, " +
		"debug_level varchar(25), " +
		"min_area int, " +
		"max_area int, " +
		"min_rental_period int, " +
		"property_type varchar(100), " +
		"contract_type varchar(100), " +
		"country varchar(100), " +
		"property_location varchar(100), " +
		"sort_field_name varchar(100), " +
		"average_rental int, " +
		"profit_months int, " +
		"prior_repair int, " +
		"contract_registration int, " +
		"running_cost int, " +
		"yearly_insurance int, " +
		"monthly_heating int, " +
		"housing_office_maintenance int, " +
		"accounting_service int, " +
		"required_profit_margin int" +
		")")

}

// Reverse the migrations
func (m *SettingsTable_20190310_223038) Down() {
	m.SQL("DROP TABLE settings")

}
