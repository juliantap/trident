// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SecurityConfigModifyReader is a Reader for the SecurityConfigModify structure.
type SecurityConfigModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityConfigModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewSecurityConfigModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityConfigModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityConfigModifyAccepted creates a SecurityConfigModifyAccepted with default headers values
func NewSecurityConfigModifyAccepted() *SecurityConfigModifyAccepted {
	return &SecurityConfigModifyAccepted{}
}

/* SecurityConfigModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SecurityConfigModifyAccepted struct {
	Payload *models.JobLinkResponse
}

func (o *SecurityConfigModifyAccepted) Error() string {
	return fmt.Sprintf("[PATCH /security][%d] securityConfigModifyAccepted  %+v", 202, o.Payload)
}
func (o *SecurityConfigModifyAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SecurityConfigModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityConfigModifyDefault creates a SecurityConfigModifyDefault with default headers values
func NewSecurityConfigModifyDefault(code int) *SecurityConfigModifyDefault {
	return &SecurityConfigModifyDefault{
		_statusCode: code,
	}
}

/* SecurityConfigModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 5636142 | This operation is not supported in a mixed-release cluster. |
| 52428817 | SSLv3 is not supported when FIPS is enabled. |
| 52428824 | TLSv1 is not supported when FIPS is enabled. |
| 52428830 | Cannot enable FIPS-compliant mode because the configured minimum security strength for certificates is not compatible. |
| 52559974 | Cannot enable FIPS-compliant mode because a certificate that is not FIPS-compliant is in use. |
| 196608081 | Cannot start software encryption conversion while there are data volumes in the cluster. |
| 196608082 | The operation is not valid when the MetroCluster is in switchover mode. |

*/
type SecurityConfigModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the security config modify default response
func (o *SecurityConfigModifyDefault) Code() int {
	return o._statusCode
}

func (o *SecurityConfigModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /security][%d] security_config_modify default  %+v", o._statusCode, o.Payload)
}
func (o *SecurityConfigModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityConfigModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
