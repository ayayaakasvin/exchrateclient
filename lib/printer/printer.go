package printer

import (
	"fmt"
	"reflect"
)

// PrintStruct prints the field names and their values for a given struct.
// It takes an interface{} as input and uses reflection to extract field names and values.
func PrintStruct(obj interface{}) {
	value := reflect.ValueOf(obj) // Get the reflection value of the input object.

	// If the input is a pointer, get the value it points to.
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	// Check if the value is of kind Struct. If not, print an error message.
	if value.Kind() != reflect.Struct {
		fmt.Printf("Provided data is not struct: %s\n", value.Kind().String())
		return // Exit the function if the input is not a struct.
	}

	// Iterate over the fields of the struct.
	for i := 0; i < value.NumField(); i++ {
		fieldName := value.Type().Field(i)                 // Get the type information for the field.
		fieldValue := value.Field(i).Interface()           // Get the actual value of the field.

		// Print the field name and its corresponding value.
		fmt.Printf("%s : %v\n", fieldName.Name, fieldValue)	
	}
}
