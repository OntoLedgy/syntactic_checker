package results_processor

import (
	"storage/csv"
	"syntactic_checker/code/services/syntactic_check_service/configuration_handler"
)

func Report_syntactic_check_outputs(
	syntactic_check_results map[string][][]string,
	run_configuration *configuration_handler.Configurations) {

	issues_file_name := run_configuration.Output_configuration.Output_issues_file_absolute_path
	issue_parameters_file_name := run_configuration.Output_configuration.Output_issue_parameters_file_absolute_path
	fixes_file_name := run_configuration.Output_configuration.Output_fixes_file_absolute_path

	syntactic_check_issues_set :=
		syntactic_check_results["syntactic_check_issues_set"]

	Report_syntactic_check_issues(
		syntactic_check_issues_set,
		issues_file_name)

	Report_syntactic_check_issue_parameters(
		syntactic_check_results["syntactic_check_issue_parameters_set"],
		issue_parameters_file_name)

	Report_syntactic_check_fixes(
		syntactic_check_results["syntactic_check_fix_transactions_set"],
		fixes_file_name)
}

func Report_syntactic_check_issues(
	issue_transactions [][]string,
	output_csv_filename string) {

	output_header := []string{
		"check_uuids",
		"check_type_uuids",
		"identifiers"}

	storage.Write_slice_with_header_to_csv(
		issue_transactions,
		output_header,
		output_csv_filename)

}

func Report_syntactic_check_issue_parameters(
	issue_parameter_transactions [][]string,
	output_csv_filename string) {

	output_header := []string{
		"check_uuids",
		"parameter_sequence_number",
		"parameter_values"}

	storage.Write_slice_with_header_to_csv(
		issue_parameter_transactions,
		output_header,
		output_csv_filename)

}

func Report_syntactic_check_fixes(
	fix_transactions [][]string,
	output_csv_filename string) {

	output_header := []string{
		"cell_values_original",
		"cell_values_marked",
		"cell_values_fixed",
		"fix_uuids",
		"identifiers"}

	storage.Write_slice_with_header_to_csv(
		fix_transactions,
		output_header,
		output_csv_filename)

}