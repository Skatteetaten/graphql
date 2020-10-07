package graphql

import "strings"

// Extension of error to provide a graphql structure without losing backwards compatibility
// When returning an error of type Error, standard error handling will work, but the
// underlying graphql error structure will be available.
//
// Type assert to get the underlying errors:
//   err := client.Run(..)
//   if err != nil {
//     if gqlErrors, ok := err.(graphql.Errors); ok {
//       for _, e := range gqlErrors {
//         // Server returned an error
//       }
//     }
//     // Another error occurred
//   }

// Errors contains the errors that were returned by the GraphQL server.
type Errors []Error

// Error() method ensures that this can be returned as error with backward compatibility
func (errors Errors) Error() string {
	if len(errors) == 0 {
		return "no errors"
	}
	errs := make([]string, len(errors))
	for i, e := range errors {
		errs[i] = e.Message
	}
	return "graphql: " + strings.Join(errs, "; ")
}

// An Error contains error information returned by the GraphQL server.
type Error struct {
	// Message contains the error message.
	Message string
	// Locations contains the locations in the GraphQL document that caused the error
	Locations []Location
	// Path contains the key path of the response field which got the error.
	Path []interface{}
	// Extensions may contain additional fields set by the GraphQL service.
	Extensions map[string]interface{}
}

// A Location is a location in the GraphQL query that resulted in an error.
type Location struct {
	Line   int
	Column int
}

func (e Error) Error() string {
	return "graphql: " + e.Message
}
