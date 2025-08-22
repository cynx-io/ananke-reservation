package response

import "github.com/cynx-io/cynx-core/src/response"

const (
	// Expected Error
	CodeSuccess            response.Code = "00"
	codeValidationError    response.Code = "VE"
	codeUnauthorized       response.Code = "UA"
	codeNotAllowed         response.Code = "NA"
	codeNotFound           response.Code = "NF"
	codeInvalidCredentials response.Code = "IC"

	// Internal
	codeInternalError response.Code = "I-IE"

	// Micro
	codeMicroHermesError response.Code = "M-HE"
	codeMicroPlutusError response.Code = "M-PL"

	// External Errors

	// Database Errors
	codeDatabasePreorderError     response.Code = "D-PR"
	codeDatabasePreorderTypeError response.Code = "D-PT"
)

var responseCodeNames = map[response.Code]string{
	// Expected Error
	CodeSuccess:            "Success",
	codeValidationError:    "Validation Error",
	codeUnauthorized:       "Not Authorized",
	codeNotAllowed:         "Not Allowed",
	codeNotFound:           "Not Found",
	codeInvalidCredentials: "Invalid Credentials",

	// Internal
	codeInternalError: "Internal Error",

	// Micro
	codeMicroHermesError: "Hermes Error",
	codeMicroPlutusError: "Plutus Error",

	// External Errors

	// Database Errors
	codeDatabasePreorderError:     "Database Preorder Error",
	codeDatabasePreorderTypeError: "Database Preorder Type Error",
}
