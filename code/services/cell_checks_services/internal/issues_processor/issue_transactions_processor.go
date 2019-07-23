package issues_processor

import (
	"fmt"
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/services/cell_check_services/contract"
	"syntactic_checker/code/services/cell_checks_services/internal/check_result_processors"
)

func (
	issues_processor *issuesProcessors) process_issue_transactions(
	cell_check_service contract.ICellCheckServices) *issues.Issues {

	issues_found := cell_check_service.Get_check_result() != nil && cell_check_service.Get_check_result().Check_result_string_edit_ranges != nil

	if issues_found {

		fmt.Printf(
			"\nprocessing issues...\n")

		cell_check_issue :=
			issues_processor.
				generate_issue_transaction(
					cell_check_service)

		return &cell_check_issue

	}
	return nil
}

func (
	issues_processor *issuesProcessors) generate_issue_transaction(
	cell_check_service contract.ICellCheckServices) issues.Issues {

	cell_check_issue :=
		new(
			issues.
				Issues)

	cell_check_issue.
		Object_uuid =
		cell_check_issue.
			Objects.
			Set_object_uuid()

	check_result_processor :=
		new(
			check_result_processors.
				CheckResultProcessors)

	check_result_processor.
		Check_results =
		cell_check_service.
			Get_check_result()

	check_result_processor.
		In_scope_issue_type =
		cell_check_issue.
			Issue_type

	check_result_processor.
		In_scope_cell =
		issues_processor.
			in_scope_cell

	check_result_processor.
		Process_regex_result()

	cell_check_issue.
		Cell_edit_history =
		*check_result_processor.
			Cell_edit_history

	return *cell_check_issue
}