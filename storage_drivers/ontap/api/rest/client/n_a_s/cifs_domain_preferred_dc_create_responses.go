// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// CifsDomainPreferredDcCreateReader is a Reader for the CifsDomainPreferredDcCreate structure.
type CifsDomainPreferredDcCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CifsDomainPreferredDcCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCifsDomainPreferredDcCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCifsDomainPreferredDcCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCifsDomainPreferredDcCreateCreated creates a CifsDomainPreferredDcCreateCreated with default headers values
func NewCifsDomainPreferredDcCreateCreated() *CifsDomainPreferredDcCreateCreated {
	return &CifsDomainPreferredDcCreateCreated{}
}

/* CifsDomainPreferredDcCreateCreated describes a response with status code 201, with default header values.

Created
*/
type CifsDomainPreferredDcCreateCreated struct {
	Payload *models.CifsDomainPreferredDc
}

func (o *CifsDomainPreferredDcCreateCreated) Error() string {
	return fmt.Sprintf("[POST /protocols/cifs/domains/{svm.uuid}/preferred-domain-controllers][%d] cifsDomainPreferredDcCreateCreated  %+v", 201, o.Payload)
}
func (o *CifsDomainPreferredDcCreateCreated) GetPayload() *models.CifsDomainPreferredDc {
	return o.Payload
}

func (o *CifsDomainPreferredDcCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CifsDomainPreferredDc)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCifsDomainPreferredDcCreateDefault creates a CifsDomainPreferredDcCreateDefault with default headers values
func NewCifsDomainPreferredDcCreateDefault(code int) *CifsDomainPreferredDcCreateDefault {
	return &CifsDomainPreferredDcCreateDefault{
		_statusCode: code,
	}
}

/* CifsDomainPreferredDcCreateDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 2621516    | Only data SVMs allowed. |
| 655918     | The fully qualified domain name cannot be longer than 254 bytes. |
| 656408     | RPC failure occured during the CIFS preferred-dc configuration validation. |
| 656407     | Failed to validate CIFS preferred-dc for domain. Reason: Configuration not found at SecD. Contact technical support for assistance. |
| 655366     | Invalid domain controller. |
| 655506     | Failed to add preferred-dc. |

*/
type CifsDomainPreferredDcCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the cifs domain preferred dc create default response
func (o *CifsDomainPreferredDcCreateDefault) Code() int {
	return o._statusCode
}

func (o *CifsDomainPreferredDcCreateDefault) Error() string {
	return fmt.Sprintf("[POST /protocols/cifs/domains/{svm.uuid}/preferred-domain-controllers][%d] cifs_domain_preferred_dc_create default  %+v", o._statusCode, o.Payload)
}
func (o *CifsDomainPreferredDcCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CifsDomainPreferredDcCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
