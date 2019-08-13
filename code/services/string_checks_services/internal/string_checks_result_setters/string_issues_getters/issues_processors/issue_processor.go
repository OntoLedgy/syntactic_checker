package issues_processors

import (
	"syntactic_checker/code/object_model/issues"
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/services/string_check_services"
)

//TODO - add separate type for this

func (
	issues_processor *issuesProcessors) get_string_check_issue(
	in_scope_issue_type issues.IssueTypes) *issues.Issues {

	string_check_parameter := new(
		service_parameters.
			StringCheckParameters)

	string_check_parameter.
		In_scope_issue_type =
		in_scope_issue_type

	string_check_parameter.
		String_value =
		issues_processor.
			string_checks_parameter.
			String_value

	String_check_service_factory :=
		new(
			string_check_services.
				StringCheckServiceFactory)

	string_check_service :=
		String_check_service_factory.
			Create(
				*string_check_parameter)

	string_check_service.
		Set_string_check_result()

	string_check_issue :=
		issues_processor.
			process_issue_transactions(
				string_check_service)

	return string_check_issue

}

func (
	issues_processor *issuesProcessors) process_string_check_issue(
	string_check_issue *issues.Issues,
	in_scope_issue_type issues.IssueTypes) {

	there_is_an_issue :=
		string_check_issue != nil

	if there_is_an_issue {

		issues_processor.
			append_string_check_issue(
				string_check_issue,
				in_scope_issue_type)

	}
}

func (
	issues_processor *issuesProcessors) append_string_check_issue(
	string_check_issue *issues.Issues,
	in_scope_issue_type issues.IssueTypes) {

	string_check_issue.
		Issue_type =
		in_scope_issue_type

	issues_processor.
		string_checks_issues =
		append(
			issues_processor.string_checks_issues,
			*string_check_issue)

}