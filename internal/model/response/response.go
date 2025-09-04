package response

import (
	"github.com/cynx-io/cynx-core/src/response"
)

func setResponse[Resp response.Generic](resp Resp, code response.Code) {
	resp.GetBase().Code = code.String()
	resp.GetBase().Desc = responseCodeNames[code]
}

func Success[Resp response.Generic](resp Resp) {
	setResponse(resp, CodeSuccess)
}

func ErrorValidation[Resp response.Generic](resp Resp) {
	setResponse(resp, codeValidationError)
}

func ErrorUnauthorized[Resp response.Generic](resp Resp) {
	setResponse(resp, codeUnauthorized)
}

func ErrorNotAllowed[Resp response.Generic](resp Resp) {
	setResponse(resp, codeNotAllowed)
}

func ErrorNotFound[Resp response.Generic](resp Resp) {
	setResponse(resp, codeNotFound)
}

func ErrorInvalidCredentials[Resp response.Generic](resp Resp) {
	setResponse(resp, codeInvalidCredentials)
}

func ErrorInternal[Resp response.Generic](resp Resp) {
	setResponse(resp, codeInternalError)
}

func ErrorDatabasePreorder[Resp response.Generic](resp Resp) {
	setResponse(resp, codeDatabasePreorderError)
}

func ErrorDatabasePreorderType[Resp response.Generic](resp Resp) {
	setResponse(resp, codeDatabasePreorderTypeError)
}

func ErrorDatabaseWaitlist[Resp response.Generic](resp Resp) {
	setResponse(resp, codeDatabaseWaitlistError)
}

func ErrorDatabaseWaitlistType[Resp response.Generic](resp Resp) {
	setResponse(resp, codeDatabaseWaitlistTypeError)
}

func ErrorMicroHermes[Resp response.Generic](resp Resp) {
	setResponse(resp, codeMicroHermesError)
}

func ErrorMicroPlutus[Resp response.Generic](resp Resp) {
	setResponse(resp, codeMicroPlutusError)
}

func ErrorExternalRecaptcha[Resp response.Generic](resp Resp) {
	setResponse(resp, codeExternalRecaptcha)
}

func ErrorExternalBrevo[Resp response.Generic](resp Resp) {
	setResponse(resp, codeExternalBrevo)
}
