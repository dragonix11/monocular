package charts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kubernetes-helm/monocular/src/api/swagger/models"
)

/*GetChartVersionOK charts response

swagger:response getChartVersionOK
*/
type GetChartVersionOK struct {

	// In: body
	Payload *models.ResourceArrayData `json:"body,omitempty"`
}

// NewGetChartVersionOK creates GetChartVersionOK with default headers values
func NewGetChartVersionOK() *GetChartVersionOK {
	return &GetChartVersionOK{}
}

// WithPayload adds the payload to the get chart version o k response
func (o *GetChartVersionOK) WithPayload(payload *models.ResourceArrayData) *GetChartVersionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get chart version o k response
func (o *GetChartVersionOK) SetPayload(payload *models.ResourceArrayData) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChartVersionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetChartVersionDefault unexpected error

swagger:response getChartVersionDefault
*/
type GetChartVersionDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetChartVersionDefault creates GetChartVersionDefault with default headers values
func NewGetChartVersionDefault(code int) *GetChartVersionDefault {
	if code <= 0 {
		code = 500
	}

	return &GetChartVersionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get chart version default response
func (o *GetChartVersionDefault) WithStatusCode(code int) *GetChartVersionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get chart version default response
func (o *GetChartVersionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get chart version default response
func (o *GetChartVersionDefault) WithPayload(payload *models.Error) *GetChartVersionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get chart version default response
func (o *GetChartVersionDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChartVersionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
