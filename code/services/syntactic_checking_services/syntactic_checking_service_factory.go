package syntactic_checking_services

import (
	"logger/goinggo_services"
	"syntactic_checker/code/services/syntactic_checking_services/contract"
	"syntactic_checker/code/services/syntactic_checking_services/internal"
	"syntactic_checker/code/services/syntactic_checking_services/internal/configuration_getters"
	"syntactic_checker/code/services/syntactic_checking_services/internal/configuration_getters/cells_preparers"
	"syntactic_checker/code/services/syntactic_checking_services/internal/configuration_getters/object_model"
)

type SyntacticCheckingServiceFactory struct{}

func (
	factory SyntacticCheckingServiceFactory) Create(
	configuration_file_path string,
	logger *goinggo_services.Logger) contract.ISyntacticCheckingServices {

	run_configuration :=
		factory.
			get_current_run_configuration(
				configuration_file_path)

	syntactic_checking_service :=
		new(
			internal.
				SyntacticCheckingServices)

	syntactic_checking_service.
		Run_configuration =
		run_configuration

	factory.
		load_in_scope_cell_list(
			syntactic_checking_service)

	syntactic_checking_service.
		Logger =
		logger

	return syntactic_checking_service
}

func (
	factory SyntacticCheckingServiceFactory) get_current_run_configuration(
	configuration_file_path string) object_model.RunConfigurations {

	configuration_getter_factory :=
		new(
			configuration_getters.
				ConfigurationGetterFactories)

	configuration_getter :=
		configuration_getter_factory.
			Create()
	run_configuration :=
		*configuration_getter.
			Get_configuration(
				configuration_file_path)
	return run_configuration
}

func (
	SyntacticCheckingServiceFactory) load_in_scope_cell_list(
	syntactic_checking_service *internal.SyntacticCheckingServices) {

	check_configuration :=
		syntactic_checking_service.
			Run_configuration.
			Check_configuration

	identity_column_name :=
		check_configuration.
			Identity_column_name

	check_column_name :=
		check_configuration.
			Check_column_name

	csv_filename :=
		check_configuration.
			Input_csv_file_name

	cells_preparer_factory :=
		new(
			cells_preparers.CellsPreparerFactory)

	cells_preparer :=
		cells_preparer_factory.
			Create(
				csv_filename,
				check_column_name,
				identity_column_name)

	syntactic_checking_service.
		In_scope_cell_list =
		cells_preparer.
			Get_in_scope_identified_cells()

}
