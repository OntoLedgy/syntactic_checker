package contract

import (
	"syntactic_checker/code/object_model/service_parameters"
	"syntactic_checker/code/object_model/service_results"
)

type IStringCheckServices interface {
	Set_string_check_result()
	Get_check_parameter() *service_parameters.StringCheckParameters
	Get_string_check_result() *service_results.StringCheckResults
	Set_string_check_result_value(*service_results.StringCheckResults)
}
