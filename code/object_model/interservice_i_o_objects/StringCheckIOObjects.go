package interservice_i_o_objects

import (
	"github.com/OntoLedgy/syntactic_checker/code/object_model/interservice_i_o_objects/service_inputs"
	"github.com/OntoLedgy/syntactic_checker/code/object_model/interservice_i_o_objects/service_results"
)

type StringChecksIOObjects struct {
	String_checks_input  *service_inputs.StringChecksInputs
	String_checks_result *service_results.StringChecksResults
}
