package windowsiot

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/windowsiot/mgmt/2019-06-01/windowsiot"

// ServiceNameUnavailabilityReason enumerates the values for service name unavailability reason.
type ServiceNameUnavailabilityReason string

const (
	// AlreadyExists ...
	AlreadyExists ServiceNameUnavailabilityReason = "AlreadyExists"
	// Invalid ...
	Invalid ServiceNameUnavailabilityReason = "Invalid"
)

// PossibleServiceNameUnavailabilityReasonValues returns an array of possible values for the ServiceNameUnavailabilityReason const type.
func PossibleServiceNameUnavailabilityReasonValues() []ServiceNameUnavailabilityReason {
	return []ServiceNameUnavailabilityReason{AlreadyExists, Invalid}
}

// DeviceService the description of the Windows IoT Device Service.
type DeviceService struct {
	autorest.Response `json:"-"`
	// Etag - The Etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal ETag convention.
	Etag *string `json:"etag,omitempty"`
	// DeviceServiceProperties - The properties of a Windows IoT Device Service.
	*DeviceServiceProperties `json:"properties,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The Azure Region where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource Id for the resource
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for DeviceService.
func (ds DeviceService) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ds.Etag != nil {
		objectMap["etag"] = ds.Etag
	}
	if ds.DeviceServiceProperties != nil {
		objectMap["properties"] = ds.DeviceServiceProperties
	}
	if ds.Tags != nil {
		objectMap["tags"] = ds.Tags
	}
	if ds.Location != nil {
		objectMap["location"] = ds.Location
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for DeviceService struct.
func (ds *DeviceService) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "etag":
			if v != nil {
				var etag string
				err = json.Unmarshal(*v, &etag)
				if err != nil {
					return err
				}
				ds.Etag = &etag
			}
		case "properties":
			if v != nil {
				var deviceServiceProperties DeviceServiceProperties
				err = json.Unmarshal(*v, &deviceServiceProperties)
				if err != nil {
					return err
				}
				ds.DeviceServiceProperties = &deviceServiceProperties
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				ds.Tags = tags
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				ds.Location = &location
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				ds.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				ds.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				ds.Type = &typeVar
			}
		}
	}

	return nil
}

// DeviceServiceCheckNameAvailabilityParameters input values.
type DeviceServiceCheckNameAvailabilityParameters struct {
	// Name - The name of the Windows IoT Device Service to check.
	Name *string `json:"name,omitempty"`
}

// DeviceServiceDescriptionListResult the JSON-serialized array of DeviceService objects with a next link.
type DeviceServiceDescriptionListResult struct {
	autorest.Response `json:"-"`
	// Value - The array of DeviceService objects.
	Value *[]DeviceService `json:"value,omitempty"`
	// NextLink - READ-ONLY; The next link.
	NextLink *string `json:"nextLink,omitempty"`
}

// DeviceServiceDescriptionListResultIterator provides access to a complete listing of DeviceService
// values.
type DeviceServiceDescriptionListResultIterator struct {
	i    int
	page DeviceServiceDescriptionListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *DeviceServiceDescriptionListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeviceServiceDescriptionListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *DeviceServiceDescriptionListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter DeviceServiceDescriptionListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter DeviceServiceDescriptionListResultIterator) Response() DeviceServiceDescriptionListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter DeviceServiceDescriptionListResultIterator) Value() DeviceService {
	if !iter.page.NotDone() {
		return DeviceService{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the DeviceServiceDescriptionListResultIterator type.
func NewDeviceServiceDescriptionListResultIterator(page DeviceServiceDescriptionListResultPage) DeviceServiceDescriptionListResultIterator {
	return DeviceServiceDescriptionListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (dsdlr DeviceServiceDescriptionListResult) IsEmpty() bool {
	return dsdlr.Value == nil || len(*dsdlr.Value) == 0
}

// deviceServiceDescriptionListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (dsdlr DeviceServiceDescriptionListResult) deviceServiceDescriptionListResultPreparer(ctx context.Context) (*http.Request, error) {
	if dsdlr.NextLink == nil || len(to.String(dsdlr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(dsdlr.NextLink)))
}

// DeviceServiceDescriptionListResultPage contains a page of DeviceService values.
type DeviceServiceDescriptionListResultPage struct {
	fn    func(context.Context, DeviceServiceDescriptionListResult) (DeviceServiceDescriptionListResult, error)
	dsdlr DeviceServiceDescriptionListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *DeviceServiceDescriptionListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeviceServiceDescriptionListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.dsdlr)
	if err != nil {
		return err
	}
	page.dsdlr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *DeviceServiceDescriptionListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page DeviceServiceDescriptionListResultPage) NotDone() bool {
	return !page.dsdlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page DeviceServiceDescriptionListResultPage) Response() DeviceServiceDescriptionListResult {
	return page.dsdlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page DeviceServiceDescriptionListResultPage) Values() []DeviceService {
	if page.dsdlr.IsEmpty() {
		return nil
	}
	return *page.dsdlr.Value
}

// Creates a new instance of the DeviceServiceDescriptionListResultPage type.
func NewDeviceServiceDescriptionListResultPage(getNextPage func(context.Context, DeviceServiceDescriptionListResult) (DeviceServiceDescriptionListResult, error)) DeviceServiceDescriptionListResultPage {
	return DeviceServiceDescriptionListResultPage{fn: getNextPage}
}

// DeviceServiceNameAvailabilityInfo the properties indicating whether a given Windows IoT Device Service
// name is available.
type DeviceServiceNameAvailabilityInfo struct {
	autorest.Response `json:"-"`
	// NameAvailable - READ-ONLY; The value which indicates whether the provided name is available.
	NameAvailable *bool `json:"nameAvailable,omitempty"`
	// Reason - READ-ONLY; The reason for unavailability. Possible values include: 'Invalid', 'AlreadyExists'
	Reason ServiceNameUnavailabilityReason `json:"reason,omitempty"`
	// Message - The detailed reason message.
	Message *string `json:"message,omitempty"`
}

// DeviceServiceProperties the properties of a Windows IoT Device Service.
type DeviceServiceProperties struct {
	// Notes - Windows IoT Device Service notes.
	Notes *string `json:"notes,omitempty"`
	// StartDate - READ-ONLY; Windows IoT Device Service start date,
	StartDate *date.Time `json:"startDate,omitempty"`
	// Quantity - Windows IoT Device Service device allocation,
	Quantity *int64 `json:"quantity,omitempty"`
	// BillingDomainName - Windows IoT Device Service ODM AAD domain
	BillingDomainName *string `json:"billingDomainName,omitempty"`
	// AdminDomainName - Windows IoT Device Service OEM AAD domain
	AdminDomainName *string `json:"adminDomainName,omitempty"`
}

// ErrorDetails the details of the error.
type ErrorDetails struct {
	// Code - One of a server-defined set of error codes.
	Code *string `json:"code,omitempty"`
	// Message - A human-readable representation of the error.
	Message *string `json:"message,omitempty"`
	// Target - The target of the particular error.
	Target *string `json:"target,omitempty"`
	// Details - A human-readable representation of the error's details.
	Details *string `json:"details,omitempty"`
}

// OperationDisplayInfo the operation supported by Azure Data Catalog Service.
type OperationDisplayInfo struct {
	// Description - The description of the operation.
	Description *string `json:"description,omitempty"`
	// Operation - The action that users can perform, based on their permission level.
	Operation *string `json:"operation,omitempty"`
	// Provider - Service provider: Azure Data Catalog Service.
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed.
	Resource *string `json:"resource,omitempty"`
}

// OperationEntity the operation supported by Azure Data Catalog Service.
type OperationEntity struct {
	// Name - Operation name: {provider}/{resource}/{operation}.
	Name *string `json:"name,omitempty"`
	// Display - The operation supported by Azure Data Catalog Service.
	Display *OperationDisplayInfo `json:"display,omitempty"`
}

// OperationListResult result of the request to list Windows IoT Device Service operations. It contains a
// list of operations and a URL link to get the next set of results.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; List of Windows IoT Device Service operations supported by the Microsoft.WindowsIoT resource provider.
	Value *[]OperationEntity `json:"value,omitempty"`
	// NextLink - READ-ONLY; URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`
}

// OperationListResultIterator provides access to a complete listing of OperationEntity values.
type OperationListResultIterator struct {
	i    int
	page OperationListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *OperationListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationListResultIterator) Response() OperationListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationListResultIterator) Value() OperationEntity {
	if !iter.page.NotDone() {
		return OperationEntity{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the OperationListResultIterator type.
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return OperationListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (olr OperationListResult) IsEmpty() bool {
	return olr.Value == nil || len(*olr.Value) == 0
}

// operationListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (olr OperationListResult) operationListResultPreparer(ctx context.Context) (*http.Request, error) {
	if olr.NextLink == nil || len(to.String(olr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(olr.NextLink)))
}

// OperationListResultPage contains a page of OperationEntity values.
type OperationListResultPage struct {
	fn  func(context.Context, OperationListResult) (OperationListResult, error)
	olr OperationListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.olr)
	if err != nil {
		return err
	}
	page.olr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *OperationListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationListResultPage) NotDone() bool {
	return !page.olr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationListResultPage) Response() OperationListResult {
	return page.olr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationListResultPage) Values() []OperationEntity {
	if page.olr.IsEmpty() {
		return nil
	}
	return *page.olr.Value
}

// Creates a new instance of the OperationListResultPage type.
func NewOperationListResultPage(getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return OperationListResultPage{fn: getNextPage}
}

// ProxyResource the resource model definition for a ARM proxy resource. It will have everything other than
// required location and tags
type ProxyResource struct {
	// ID - READ-ONLY; Fully qualified resource Id for the resource
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty"`
}

// Resource the core properties of ARM resources
type Resource struct {
	// ID - READ-ONLY; Fully qualified resource Id for the resource
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty"`
}

// TrackedResource the resource model definition for a ARM tracked top level resource
type TrackedResource struct {
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The Azure Region where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource Id for the resource
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for TrackedResource.
func (tr TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if tr.Tags != nil {
		objectMap["tags"] = tr.Tags
	}
	if tr.Location != nil {
		objectMap["location"] = tr.Location
	}
	return json.Marshal(objectMap)
}
