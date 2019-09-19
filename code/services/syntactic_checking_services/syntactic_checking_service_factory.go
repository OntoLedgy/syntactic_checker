package syntactic_checking_services

import (
	"syntactic_checker/code/object_model/configurations"
	"syntactic_checker/code/services/syntactic_checking_services/contract"
	"syntactic_checker/code/services/syntactic_checking_services/internal"
)

type SyntacticCheckingServiceFactory struct{}

func (
	factory SyntacticCheckingServiceFactory) Create(
	service_run_data *configurations.SyntacticCheckingData) contract.ISyntacticCheckingServices {

	syntactic_checking_service :=
		new(
			internal.
				SyntacticCheckingServices)

	syntactic_checking_service.
		Syntactic_checking_service_data =
		service_run_data

	return syntactic_checking_service
}
