package identified_string_list_preparers

import (
	"fmt"
	"storage/csv"
	storage_slices "storage/slices"
	"syntactic_checker/code/object_model/identified_strings"
)

type identifiedStringListPreparers struct {
	identity_column_name string
	check_column_name    string
	csv_filename         string //should be path?
}

func (
	identified_string_preparer identifiedStringListPreparers) Get_in_scope_identified_identified_string() identified_strings.IdentifiedStringLists {

	csv_filename :=
		identified_string_preparer.
			csv_filename

	identity_colunmn_name :=
		identified_string_preparer.
			identity_column_name

	identified_string_value_column_name :=
		identified_string_preparer.
			check_column_name

	identified_string_list :=
		prepare_identified_string_data(
			csv_filename,
			identity_colunmn_name,
			identified_string_value_column_name)

	return identified_string_list
}

//TODO - Stage 3 - data reading needs to be improved

func prepare_identified_string_data(
	csv_filename string,
	identity_colunmn_name string,
	string_value_column_name string) identified_strings.IdentifiedStringLists {

	fmt.Printf(
		"\nReading CSV Data..")

	var identified_string_list identified_strings.IdentifiedStringLists

	identified_string_list_raw :=
		storage.Read_csv_data(
			csv_filename, "")

	fmt.Printf(
		"Preparing extracted data for checks (converting to interface)")

	identified_string_list_interface :=
		storage_slices.Convert_2d_string_to_interface(
			identified_string_list_raw)

	identified_string_list_with_headers :=
		storage.Get_csv_with_headers(
			identified_string_list_interface)

	identified_string_list.Identified_string_list =
		make(
			[]identified_strings.IdentifiedStrings,
			len(identified_string_list_interface))

	for index, value := range identified_string_list_with_headers {

		identified_string_list.Identified_string_list[index].String_identifier =
			value[identity_colunmn_name].(string)

		identified_string_list.Identified_string_list[index].String_value =
			value[string_value_column_name].(string)

	}

	return identified_string_list
}