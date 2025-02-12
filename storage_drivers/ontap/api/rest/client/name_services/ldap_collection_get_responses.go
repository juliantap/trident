// Code generated by go-swagger; DO NOT EDIT.

package name_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// LdapCollectionGetReader is a Reader for the LdapCollectionGet structure.
type LdapCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LdapCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLdapCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLdapCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLdapCollectionGetOK creates a LdapCollectionGetOK with default headers values
func NewLdapCollectionGetOK() *LdapCollectionGetOK {
	return &LdapCollectionGetOK{}
}

/* LdapCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type LdapCollectionGetOK struct {
	Payload *models.LdapServiceResponse
}

func (o *LdapCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /name-services/ldap][%d] ldapCollectionGetOK  %+v", 200, o.Payload)
}
func (o *LdapCollectionGetOK) GetPayload() *models.LdapServiceResponse {
	return o.Payload
}

func (o *LdapCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LdapServiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLdapCollectionGetDefault creates a LdapCollectionGetDefault with default headers values
func NewLdapCollectionGetDefault(code int) *LdapCollectionGetDefault {
	return &LdapCollectionGetDefault{
		_statusCode: code,
	}
}

/* LdapCollectionGetDefault describes a response with status code -1, with default header values.

 Following error codes can be thrown as part of LDAP status information, if LDAP status is needed to be retrieved.
| Error Code | Description |
| ---------- | ----------- |
| 4915229    | DNS resolution failed due to an internal error. Contact technical support if this issue persists |
| 4915231    | DNS resolution failed for one or more of the specified LDAP servers. Verify that a valid DNS server is configured |
| 23724132   | DNS resolution failed for all the specified LDAP servers. Verify that a valid DNS server is configured |
| 4915258    | The LDAP configuration is invalid. Verify that the Active Directory domain or servers are reachable and that the network configuration is correct |
| 4915263    | Failed to check the current status of LDAP server. Reason:<Reason for the failure> |
| 4915234    | The specified LDAP server or preferred Active Directory server is not supported because it is one of the following: multicast, loopback, 0.0.0.0, or broadcast |
| 4915265    | The specified bind password or bind DN is invalid |
| 4915264    | Certificate verification failed. Verify that a valid certificate is installed |

*/
type LdapCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the ldap collection get default response
func (o *LdapCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *LdapCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /name-services/ldap][%d] ldap_collection_get default  %+v", o._statusCode, o.Payload)
}
func (o *LdapCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LdapCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
