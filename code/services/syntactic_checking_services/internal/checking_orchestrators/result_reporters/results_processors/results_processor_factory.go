package results_processors

import (
	"syntactic_checker/code/object_model/service_results"
	"syntactic_checker/code/services/syntactic_checking_services/internal/configuration_getters/object_model"
)

type ResultsProcessorFactory struct{}

func (
	ResultsProcessorFactory) Create(
	identified_string_list_checks_result service_results.IdentifiedStringListChecksResults,
	output_configuration object_model.OutputConfigurations,
) *resultsProcessors {

	results_processor :=
		new(
			resultsProcessors)

	results_processor.
		output_configuration =
		output_configuration

	results_processor.
		syntactic_checking_results =
		identified_string_list_checks_result

	return results_processor

}