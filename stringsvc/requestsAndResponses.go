package main

type (
    uppercaseRequest struct {
		// S : string
		S string `json:"s"`
	}

	uppercaseResponse struct {
		// V : value
		V string `json:"v"`
		// Err : error. Errors don't JSON Marshal so we use a string
		Err string `json:"err,omitempty"`
	}

	countRequest struct {
		// S : string
		S string `json:"s"`
	}

	countResponse struct {
		// V : value
		V int `json:"v"`
	}
)