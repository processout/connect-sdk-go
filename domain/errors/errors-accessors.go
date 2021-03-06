// Copyright 2017 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen-accessors; DO NOT EDIT.

package errors

// GetCategory returns the Category field if it's non-nil, zero value otherwise.
func (a *APIError) GetCategory() string {
	if a == nil || a.Category == nil {
		return ""
	}
	return *a.Category
}

// GetCode returns the Code field if it's non-nil, zero value otherwise.
func (a *APIError) GetCode() string {
	if a == nil || a.Code == nil {
		return ""
	}
	return *a.Code
}

// GetHTTPStatusCode returns the HTTPStatusCode field if it's non-nil, zero value otherwise.
func (a *APIError) GetHTTPStatusCode() int32 {
	if a == nil || a.HTTPStatusCode == nil {
		return 0
	}
	return *a.HTTPStatusCode
}

// GetID returns the ID field if it's non-nil, zero value otherwise.
func (a *APIError) GetID() string {
	if a == nil || a.ID == nil {
		return ""
	}
	return *a.ID
}

// GetMessage returns the Message field if it's non-nil, zero value otherwise.
func (a *APIError) GetMessage() string {
	if a == nil || a.Message == nil {
		return ""
	}
	return *a.Message
}

// GetPropertyName returns the PropertyName field if it's non-nil, zero value otherwise.
func (a *APIError) GetPropertyName() string {
	if a == nil || a.PropertyName == nil {
		return ""
	}
	return *a.PropertyName
}

// GetRequestID returns the RequestID field if it's non-nil, zero value otherwise.
func (a *APIError) GetRequestID() string {
	if a == nil || a.RequestID == nil {
		return ""
	}
	return *a.RequestID
}

// GetErrorID returns the ErrorID field if it's non-nil, zero value otherwise.
func (e *ErrorResponse) GetErrorID() string {
	if e == nil || e.ErrorID == nil {
		return ""
	}
	return *e.ErrorID
}

// GetErrors returns the Errors field if it's non-nil, zero value otherwise.
func (e *ErrorResponse) GetErrors() []APIError {
	if e == nil || e.Errors == nil {
		return nil
	}
	return *e.Errors
}
