package models

import "fmt"

// OptionInvalidError represents an error related to invalid options.
type OptionInvalidError struct {
	Code    int
	Message string
}

// Error returns a formatted error string for OptionInvalidError.
func (e *OptionInvalidError) Error() string {
	return fmt.Sprintf("Code %d: %s", e.Code, e.Message)
}

// NewOptionInvalidError creates a new OptionInvalidError with the given code and message.
func NewOptionInvalidError(code int, msg string) *OptionInvalidError {
	return &OptionInvalidError{
		Code:    code,
		Message: msg,
	}
}

// IpApiError represents an error from the IP API.
type IpApiError struct {
	Code    int
	Message string
}

// Error returns a formatted error string for IpApiError.
func (e *IpApiError) Error() string {
	return fmt.Sprintf("		Code %d: %s", e.Code, e.Message)
}

// NewErrorIpApiError creates a new IpApiError with the given code and message.
func NewErrorIpApiError(code int, msg string) *IpApiError {
	return &IpApiError{
		Code:    code,
		Message: msg,
	}
}

// CountryApiError represents an error from the Country API.
type CountryApiError struct {
	Code    int
	Message string
}

// Error returns a formatted error string for CountryApiError.
func (e *CountryApiError) Error() string {
	return fmt.Sprintf("		Code %d: %s", e.Code, e.Message)
}

// NewCountryApiError creates a new CountryApiError with the given code and message.
func NewCountryApiError(code int, msg string) *CountryApiError {
	return &CountryApiError{
		Code:    code,
		Message: msg,
	}
}

// CurrencyApiError represents an error from the Currency API.
type CurrencyApiError struct {
	Code    int
	Message string
}

// Error returns a formatted error string for CurrencyApiError.
func (e *CurrencyApiError) Error() string {
	return fmt.Sprintf("		Code %d: %s", e.Code, e.Message)
}

// NewCurrencyApiError creates a new CurrencyApiError with the given code and message.
func NewCurrencyApiError(code int, msg string) *CurrencyApiError {
	return &CurrencyApiError{
		Code:    code,
		Message: msg,
	}
}
