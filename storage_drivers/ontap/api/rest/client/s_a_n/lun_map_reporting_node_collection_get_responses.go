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

// LunMapReportingNodeCollectionGetReader is a Reader for the LunMapReportingNodeCollectionGet structure.
type LunMapReportingNodeCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LunMapReportingNodeCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLunMapReportingNodeCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLunMapReportingNodeCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLunMapReportingNodeCollectionGetOK creates a LunMapReportingNodeCollectionGetOK with default headers values
func NewLunMapReportingNodeCollectionGetOK() *LunMapReportingNodeCollectionGetOK {
	return &LunMapReportingNodeCollectionGetOK{}
}

/* LunMapReportingNodeCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type LunMapReportingNodeCollectionGetOK struct {
	Payload *models.LunMapReportingNodeResponse
}

func (o *LunMapReportingNodeCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/san/lun-maps/{lun.uuid}/{igroup.uuid}/reporting-nodes][%d] lunMapReportingNodeCollectionGetOK  %+v", 200, o.Payload)
}
func (o *LunMapReportingNodeCollectionGetOK) GetPayload() *models.LunMapReportingNodeResponse {
	return o.Payload
}

func (o *LunMapReportingNodeCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LunMapReportingNodeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLunMapReportingNodeCollectionGetDefault creates a LunMapReportingNodeCollectionGetDefault with default headers values
func NewLunMapReportingNodeCollectionGetDefault(code int) *LunMapReportingNodeCollectionGetDefault {
	return &LunMapReportingNodeCollectionGetDefault{
		_statusCode: code,
	}
}

/* LunMapReportingNodeCollectionGetDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 5374875 | The specified LUN does not exist or is not accessible to the caller. |
| 5374878 | The specified initiator group does not exist, is not accessible to the caller, or is not in the same SVM as the specified LUN. |
| 5374922 | The specified LUN map does not exist. |

*/
type LunMapReportingNodeCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the lun map reporting node collection get default response
func (o *LunMapReportingNodeCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *LunMapReportingNodeCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/san/lun-maps/{lun.uuid}/{igroup.uuid}/reporting-nodes][%d] lun_map_reporting_node_collection_get default  %+v", o._statusCode, o.Payload)
}
func (o *LunMapReportingNodeCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LunMapReportingNodeCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
