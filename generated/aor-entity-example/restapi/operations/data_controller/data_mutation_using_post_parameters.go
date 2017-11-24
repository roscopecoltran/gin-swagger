package data_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/roscopecoltran/aor-gin-swagger/api"

	strfmt "github.com/go-openapi/strfmt"
	// DefaultImports
	// Imports
)

// BusinessLogicDataMutationUsingPOST executes the core logic of the related
// route endpoint.
func BusinessLogicDataMutationUsingPOST(f func(ctx *gin.Context, params *DataMutationUsingPOSTParams) *api.Response) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// generate params from request
		params := &DataMutationUsingPOSTParams{}
		err := params.bindRequest(ctx)
		if err != nil {
			errObj := err.(*errors.CompositeError)
			problem := api.Problem{
				Title:  "Unprocessable Entity.",
				Status: int(errObj.Code()),
				Detail: errObj.Error(),
			}
			ctx.Writer.Header().Set("Content-Type", "application/problem+json")
			ctx.JSON(problem.Status, problem)
			return
		}

		resp := f(ctx, params)
		switch resp.Code {
		case http.StatusNoContent:
			ctx.AbortWithStatus(resp.Code)
		default:
			ctx.JSON(resp.Code, resp.Body)
		}
	}
}

// DataMutationUsingPOSTParams contains all the bound params for the data mutation using p o s t operation
// typically these are obtained from a http.Request
//
// swagger:parameters dataMutationUsingPOST
type DataMutationUsingPOSTParams struct {

	/*allRequestParams
	  In: body
	*/
	AllRequestParams interface{}
	/*entity
	  Required: true
	  In: path
	*/
	Entity string
}

// DataMutationUsingPOSTParamsFromCtx gets the params struct from the gin context.
func DataMutationUsingPOSTParamsFromCtx(ctx *gin.Context) *DataMutationUsingPOSTParams {
	params, _ := ctx.Get("params")
	return params.(*DataMutationUsingPOSTParams)
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *DataMutationUsingPOSTParams) bindRequest(ctx *gin.Context) error {
	var res []error
	formats := strfmt.NewFormats()

	if runtime.HasBody(ctx.Request) {
		var body interface{}
		if err := ctx.BindJSON(&body); err != nil {
			res = append(res, errors.NewParseError("allRequestParams", "body", "", err))
		} else {

			if len(res) == 0 {
				o.AllRequestParams = body
			}
		}

	}

	rEntity := []string{ctx.Param("entity")}
	if err := o.bindEntity(rEntity, true, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DataMutationUsingPOSTParams) bindEntity(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Entity = raw

	return nil
}

// vim: ft=go