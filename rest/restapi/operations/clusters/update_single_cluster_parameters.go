package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/radanalyticsio/oshinko-cli/rest/models"
)

// NewUpdateSingleClusterParams creates a new UpdateSingleClusterParams object
// with the default values initialized.
func NewUpdateSingleClusterParams() UpdateSingleClusterParams {
	var ()
	return UpdateSingleClusterParams{}
}

// UpdateSingleClusterParams contains all the bound params for the update single cluster operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateSingleCluster
type UpdateSingleClusterParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Requested cluster update
	  Required: true
	  In: body
	*/
	Cluster *models.NewCluster
	/*Name of the cluster
	  Required: true
	  In: path
	*/
	Name string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *UpdateSingleClusterParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.NewCluster
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("cluster", "body"))
			} else {
				res = append(res, errors.NewParseError("cluster", "body", "", err))
			}

		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Cluster = &body
			}
		}

	} else {
		res = append(res, errors.Required("cluster", "body"))
	}

	rName, rhkName, _ := route.Params.GetOK("name")
	if err := o.bindName(rName, rhkName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateSingleClusterParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Name = raw

	return nil
}
