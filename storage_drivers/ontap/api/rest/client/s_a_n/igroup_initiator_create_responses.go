// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// IgroupInitiatorCreateReader is a Reader for the IgroupInitiatorCreate structure.
type IgroupInitiatorCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IgroupInitiatorCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIgroupInitiatorCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIgroupInitiatorCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIgroupInitiatorCreateCreated creates a IgroupInitiatorCreateCreated with default headers values
func NewIgroupInitiatorCreateCreated() *IgroupInitiatorCreateCreated {
	return &IgroupInitiatorCreateCreated{}
}

/* IgroupInitiatorCreateCreated describes a response with status code 201, with default header values.

Created
*/
type IgroupInitiatorCreateCreated struct {
	Payload *models.IgroupInitiatorResponse
}

func (o *IgroupInitiatorCreateCreated) Error() string {
	return fmt.Sprintf("[POST /protocols/san/igroups/{igroup.uuid}/initiators][%d] igroupInitiatorCreateCreated  %+v", 201, o.Payload)
}
func (o *IgroupInitiatorCreateCreated) GetPayload() *models.IgroupInitiatorResponse {
	return o.Payload
}

func (o *IgroupInitiatorCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IgroupInitiatorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIgroupInitiatorCreateDefault creates a IgroupInitiatorCreateDefault with default headers values
func NewIgroupInitiatorCreateDefault(code int) *IgroupInitiatorCreateDefault {
	return &IgroupInitiatorCreateDefault{
		_statusCode: code,
	}
}

/* IgroupInitiatorCreateDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 1254193 | Adding an initiator would cause the initiator to be mapped to the same LUN more than once. |
| 1254324 | Adding an initiator would cause the initiator to have the same logical unit identifier for multiple LUN maps. |
| 5373969 | A supplied initiator name looks like an iSCSI IQN initiator, but the portions after the prefix are missing. |
| 5373971 | A supplied initiator name looks like an iSCSI IQN initiator, but the date portion is invalid. |
| 5373972 | A supplied initiator name looks like an iSCSI IQN initiator, but the naming authority portion is invalid. |
| 5373977 | A supplied initiator name looks like an iSCSI EUI initiator, but the length is invalid. |
| 5373978 | A supplied initiator name looks like an iSCSI EUI initiator, but the format is invalid. |
| 5373992 | A supplied initiator name was too long to be valid. |
| 5373993 | A supplied initiator name did not match any valid format. |
| 5374033 | Initiators must be supplied. |
| 5374035 | A supplied initiator is already in the initiator group. |
| 5374038 | An invalid Fibre Channel WWPN was supplied. |
| 5374039 | An invalid iSCSI initiator name was supplied. |
| 5374734 | An initiator is already in another initiator group with a conflicting operating system type. |
| 5374852 | The initiator group specified in the URI does not exist. |
| 5374917 | Multiple matching initiators have been supplied with conflicting comments. |

*/
type IgroupInitiatorCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the igroup initiator create default response
func (o *IgroupInitiatorCreateDefault) Code() int {
	return o._statusCode
}

func (o *IgroupInitiatorCreateDefault) Error() string {
	return fmt.Sprintf("[POST /protocols/san/igroups/{igroup.uuid}/initiators][%d] igroup_initiator_create default  %+v", o._statusCode, o.Payload)
}
func (o *IgroupInitiatorCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IgroupInitiatorCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
