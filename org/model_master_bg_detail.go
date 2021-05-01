/*
 * Organization API
 *
 * Description of the Organization API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package org

import (
	"encoding/json"
)

// MasterBGDetail struct for MasterBGDetail
type MasterBGDetail struct {
	// An explanation about the purpose of this instance.
	ClientId *string `json:"clientId,omitempty"`
	// An explanation about the purpose of this instance.
	CreatedAt *string `json:"createdAt,omitempty"`
	// An explanation about the purpose of this instance.
	Domain *string `json:"domain,omitempty"`
	Entitlements *Entitlements `json:"entitlements,omitempty"`
	// An explanation about the purpose of this instance.
	Environments *[]Environment `json:"environments,omitempty"`
	// An explanation about the purpose of this instance.
	Id *string `json:"id,omitempty"`
	// An explanation about the purpose of this instance.
	IdproviderId *string `json:"idprovider_id,omitempty"`
	// An explanation about the purpose of this instance.
	IsAutomaticAdminPromotionExempt *bool `json:"isAutomaticAdminPromotionExempt,omitempty"`
	// An explanation about the purpose of this instance.
	IsFederated *bool `json:"isFederated,omitempty"`
	// An explanation about the purpose of this instance.
	IsMaster *bool `json:"isMaster,omitempty"`
	// An explanation about the purpose of this instance.
	MfaRequired *string `json:"mfaRequired,omitempty"`
	// An explanation about the purpose of this instance.
	Name *string `json:"name,omitempty"`
	// An explanation about the purpose of this instance.
	OwnerId *string `json:"ownerId,omitempty"`
	// An explanation about the purpose of this instance.
	ParentOrganizationIds *[]string `json:"parentOrganizationIds,omitempty"`
	// An explanation about the purpose of this instance.
	Properties *map[string]interface{} `json:"properties,omitempty"`
	// An explanation about the purpose of this instance.
	SubOrganizationIds *[]string `json:"subOrganizationIds,omitempty"`
	// An explanation about the purpose of this instance.
	TenantOrganizationIds *[]string `json:"tenantOrganizationIds,omitempty"`
	// An explanation about the purpose of this instance.
	UpdatedAt *string `json:"updatedAt,omitempty"`
	Owner *Owner `json:"owner,omitempty"`
	// An explanation about the purpose of this instance.
	SessionTimeout *int32 `json:"sessionTimeout,omitempty"`
	Subscription *Subscription `json:"subscription,omitempty"`
}

// NewMasterBGDetail instantiates a new MasterBGDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMasterBGDetail() *MasterBGDetail {
	this := MasterBGDetail{}
	var clientId string = ""
	this.ClientId = &clientId
	var createdAt string = ""
	this.CreatedAt = &createdAt
	var domain string = ""
	this.Domain = &domain
	var entitlements Entitlements = {}
	this.Entitlements = &entitlements
	var id string = ""
	this.Id = &id
	var idproviderId string = ""
	this.IdproviderId = &idproviderId
	var isAutomaticAdminPromotionExempt bool = false
	this.IsAutomaticAdminPromotionExempt = &isAutomaticAdminPromotionExempt
	var isFederated bool = false
	this.IsFederated = &isFederated
	var isMaster bool = false
	this.IsMaster = &isMaster
	var mfaRequired string = ""
	this.MfaRequired = &mfaRequired
	var name string = ""
	this.Name = &name
	var ownerId string = ""
	this.OwnerId = &ownerId
	var updatedAt string = ""
	this.UpdatedAt = &updatedAt
	var sessionTimeout int32 = 0
	this.SessionTimeout = &sessionTimeout
	return &this
}

// NewMasterBGDetailWithDefaults instantiates a new MasterBGDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMasterBGDetailWithDefaults() *MasterBGDetail {
	this := MasterBGDetail{}
	var clientId string = ""
	this.ClientId = &clientId
	var createdAt string = ""
	this.CreatedAt = &createdAt
	var domain string = ""
	this.Domain = &domain
	var entitlements Entitlements = {}
	this.Entitlements = &entitlements
	var id string = ""
	this.Id = &id
	var idproviderId string = ""
	this.IdproviderId = &idproviderId
	var isAutomaticAdminPromotionExempt bool = false
	this.IsAutomaticAdminPromotionExempt = &isAutomaticAdminPromotionExempt
	var isFederated bool = false
	this.IsFederated = &isFederated
	var isMaster bool = false
	this.IsMaster = &isMaster
	var mfaRequired string = ""
	this.MfaRequired = &mfaRequired
	var name string = ""
	this.Name = &name
	var ownerId string = ""
	this.OwnerId = &ownerId
	var updatedAt string = ""
	this.UpdatedAt = &updatedAt
	var sessionTimeout int32 = 0
	this.SessionTimeout = &sessionTimeout
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *MasterBGDetail) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *MasterBGDetail) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *MasterBGDetail) SetClientId(v string) {
	o.ClientId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *MasterBGDetail) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *MasterBGDetail) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *MasterBGDetail) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *MasterBGDetail) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *MasterBGDetail) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *MasterBGDetail) SetDomain(v string) {
	o.Domain = &v
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *MasterBGDetail) GetEntitlements() Entitlements {
	if o == nil || o.Entitlements == nil {
		var ret Entitlements
		return ret
	}
	return *o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetEntitlementsOk() (*Entitlements, bool) {
	if o == nil || o.Entitlements == nil {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *MasterBGDetail) HasEntitlements() bool {
	if o != nil && o.Entitlements != nil {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given Entitlements and assigns it to the Entitlements field.
func (o *MasterBGDetail) SetEntitlements(v Entitlements) {
	o.Entitlements = &v
}

// GetEnvironments returns the Environments field value if set, zero value otherwise.
func (o *MasterBGDetail) GetEnvironments() []Environment {
	if o == nil || o.Environments == nil {
		var ret []Environment
		return ret
	}
	return *o.Environments
}

// GetEnvironmentsOk returns a tuple with the Environments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetEnvironmentsOk() (*[]Environment, bool) {
	if o == nil || o.Environments == nil {
		return nil, false
	}
	return o.Environments, true
}

// HasEnvironments returns a boolean if a field has been set.
func (o *MasterBGDetail) HasEnvironments() bool {
	if o != nil && o.Environments != nil {
		return true
	}

	return false
}

// SetEnvironments gets a reference to the given []Environment and assigns it to the Environments field.
func (o *MasterBGDetail) SetEnvironments(v []Environment) {
	o.Environments = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MasterBGDetail) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MasterBGDetail) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MasterBGDetail) SetId(v string) {
	o.Id = &v
}

// GetIdproviderId returns the IdproviderId field value if set, zero value otherwise.
func (o *MasterBGDetail) GetIdproviderId() string {
	if o == nil || o.IdproviderId == nil {
		var ret string
		return ret
	}
	return *o.IdproviderId
}

// GetIdproviderIdOk returns a tuple with the IdproviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetIdproviderIdOk() (*string, bool) {
	if o == nil || o.IdproviderId == nil {
		return nil, false
	}
	return o.IdproviderId, true
}

// HasIdproviderId returns a boolean if a field has been set.
func (o *MasterBGDetail) HasIdproviderId() bool {
	if o != nil && o.IdproviderId != nil {
		return true
	}

	return false
}

// SetIdproviderId gets a reference to the given string and assigns it to the IdproviderId field.
func (o *MasterBGDetail) SetIdproviderId(v string) {
	o.IdproviderId = &v
}

// GetIsAutomaticAdminPromotionExempt returns the IsAutomaticAdminPromotionExempt field value if set, zero value otherwise.
func (o *MasterBGDetail) GetIsAutomaticAdminPromotionExempt() bool {
	if o == nil || o.IsAutomaticAdminPromotionExempt == nil {
		var ret bool
		return ret
	}
	return *o.IsAutomaticAdminPromotionExempt
}

// GetIsAutomaticAdminPromotionExemptOk returns a tuple with the IsAutomaticAdminPromotionExempt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetIsAutomaticAdminPromotionExemptOk() (*bool, bool) {
	if o == nil || o.IsAutomaticAdminPromotionExempt == nil {
		return nil, false
	}
	return o.IsAutomaticAdminPromotionExempt, true
}

// HasIsAutomaticAdminPromotionExempt returns a boolean if a field has been set.
func (o *MasterBGDetail) HasIsAutomaticAdminPromotionExempt() bool {
	if o != nil && o.IsAutomaticAdminPromotionExempt != nil {
		return true
	}

	return false
}

// SetIsAutomaticAdminPromotionExempt gets a reference to the given bool and assigns it to the IsAutomaticAdminPromotionExempt field.
func (o *MasterBGDetail) SetIsAutomaticAdminPromotionExempt(v bool) {
	o.IsAutomaticAdminPromotionExempt = &v
}

// GetIsFederated returns the IsFederated field value if set, zero value otherwise.
func (o *MasterBGDetail) GetIsFederated() bool {
	if o == nil || o.IsFederated == nil {
		var ret bool
		return ret
	}
	return *o.IsFederated
}

// GetIsFederatedOk returns a tuple with the IsFederated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetIsFederatedOk() (*bool, bool) {
	if o == nil || o.IsFederated == nil {
		return nil, false
	}
	return o.IsFederated, true
}

// HasIsFederated returns a boolean if a field has been set.
func (o *MasterBGDetail) HasIsFederated() bool {
	if o != nil && o.IsFederated != nil {
		return true
	}

	return false
}

// SetIsFederated gets a reference to the given bool and assigns it to the IsFederated field.
func (o *MasterBGDetail) SetIsFederated(v bool) {
	o.IsFederated = &v
}

// GetIsMaster returns the IsMaster field value if set, zero value otherwise.
func (o *MasterBGDetail) GetIsMaster() bool {
	if o == nil || o.IsMaster == nil {
		var ret bool
		return ret
	}
	return *o.IsMaster
}

// GetIsMasterOk returns a tuple with the IsMaster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetIsMasterOk() (*bool, bool) {
	if o == nil || o.IsMaster == nil {
		return nil, false
	}
	return o.IsMaster, true
}

// HasIsMaster returns a boolean if a field has been set.
func (o *MasterBGDetail) HasIsMaster() bool {
	if o != nil && o.IsMaster != nil {
		return true
	}

	return false
}

// SetIsMaster gets a reference to the given bool and assigns it to the IsMaster field.
func (o *MasterBGDetail) SetIsMaster(v bool) {
	o.IsMaster = &v
}

// GetMfaRequired returns the MfaRequired field value if set, zero value otherwise.
func (o *MasterBGDetail) GetMfaRequired() string {
	if o == nil || o.MfaRequired == nil {
		var ret string
		return ret
	}
	return *o.MfaRequired
}

// GetMfaRequiredOk returns a tuple with the MfaRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetMfaRequiredOk() (*string, bool) {
	if o == nil || o.MfaRequired == nil {
		return nil, false
	}
	return o.MfaRequired, true
}

// HasMfaRequired returns a boolean if a field has been set.
func (o *MasterBGDetail) HasMfaRequired() bool {
	if o != nil && o.MfaRequired != nil {
		return true
	}

	return false
}

// SetMfaRequired gets a reference to the given string and assigns it to the MfaRequired field.
func (o *MasterBGDetail) SetMfaRequired(v string) {
	o.MfaRequired = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MasterBGDetail) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MasterBGDetail) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MasterBGDetail) SetName(v string) {
	o.Name = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *MasterBGDetail) GetOwnerId() string {
	if o == nil || o.OwnerId == nil {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetOwnerIdOk() (*string, bool) {
	if o == nil || o.OwnerId == nil {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *MasterBGDetail) HasOwnerId() bool {
	if o != nil && o.OwnerId != nil {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *MasterBGDetail) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetParentOrganizationIds returns the ParentOrganizationIds field value if set, zero value otherwise.
func (o *MasterBGDetail) GetParentOrganizationIds() []string {
	if o == nil || o.ParentOrganizationIds == nil {
		var ret []string
		return ret
	}
	return *o.ParentOrganizationIds
}

// GetParentOrganizationIdsOk returns a tuple with the ParentOrganizationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetParentOrganizationIdsOk() (*[]string, bool) {
	if o == nil || o.ParentOrganizationIds == nil {
		return nil, false
	}
	return o.ParentOrganizationIds, true
}

// HasParentOrganizationIds returns a boolean if a field has been set.
func (o *MasterBGDetail) HasParentOrganizationIds() bool {
	if o != nil && o.ParentOrganizationIds != nil {
		return true
	}

	return false
}

// SetParentOrganizationIds gets a reference to the given []string and assigns it to the ParentOrganizationIds field.
func (o *MasterBGDetail) SetParentOrganizationIds(v []string) {
	o.ParentOrganizationIds = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *MasterBGDetail) GetProperties() map[string]interface{} {
	if o == nil || o.Properties == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetPropertiesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *MasterBGDetail) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]interface{} and assigns it to the Properties field.
func (o *MasterBGDetail) SetProperties(v map[string]interface{}) {
	o.Properties = &v
}

// GetSubOrganizationIds returns the SubOrganizationIds field value if set, zero value otherwise.
func (o *MasterBGDetail) GetSubOrganizationIds() []string {
	if o == nil || o.SubOrganizationIds == nil {
		var ret []string
		return ret
	}
	return *o.SubOrganizationIds
}

// GetSubOrganizationIdsOk returns a tuple with the SubOrganizationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetSubOrganizationIdsOk() (*[]string, bool) {
	if o == nil || o.SubOrganizationIds == nil {
		return nil, false
	}
	return o.SubOrganizationIds, true
}

// HasSubOrganizationIds returns a boolean if a field has been set.
func (o *MasterBGDetail) HasSubOrganizationIds() bool {
	if o != nil && o.SubOrganizationIds != nil {
		return true
	}

	return false
}

// SetSubOrganizationIds gets a reference to the given []string and assigns it to the SubOrganizationIds field.
func (o *MasterBGDetail) SetSubOrganizationIds(v []string) {
	o.SubOrganizationIds = &v
}

// GetTenantOrganizationIds returns the TenantOrganizationIds field value if set, zero value otherwise.
func (o *MasterBGDetail) GetTenantOrganizationIds() []string {
	if o == nil || o.TenantOrganizationIds == nil {
		var ret []string
		return ret
	}
	return *o.TenantOrganizationIds
}

// GetTenantOrganizationIdsOk returns a tuple with the TenantOrganizationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetTenantOrganizationIdsOk() (*[]string, bool) {
	if o == nil || o.TenantOrganizationIds == nil {
		return nil, false
	}
	return o.TenantOrganizationIds, true
}

// HasTenantOrganizationIds returns a boolean if a field has been set.
func (o *MasterBGDetail) HasTenantOrganizationIds() bool {
	if o != nil && o.TenantOrganizationIds != nil {
		return true
	}

	return false
}

// SetTenantOrganizationIds gets a reference to the given []string and assigns it to the TenantOrganizationIds field.
func (o *MasterBGDetail) SetTenantOrganizationIds(v []string) {
	o.TenantOrganizationIds = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *MasterBGDetail) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *MasterBGDetail) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *MasterBGDetail) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *MasterBGDetail) GetOwner() Owner {
	if o == nil || o.Owner == nil {
		var ret Owner
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetOwnerOk() (*Owner, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *MasterBGDetail) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given Owner and assigns it to the Owner field.
func (o *MasterBGDetail) SetOwner(v Owner) {
	o.Owner = &v
}

// GetSessionTimeout returns the SessionTimeout field value if set, zero value otherwise.
func (o *MasterBGDetail) GetSessionTimeout() int32 {
	if o == nil || o.SessionTimeout == nil {
		var ret int32
		return ret
	}
	return *o.SessionTimeout
}

// GetSessionTimeoutOk returns a tuple with the SessionTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetSessionTimeoutOk() (*int32, bool) {
	if o == nil || o.SessionTimeout == nil {
		return nil, false
	}
	return o.SessionTimeout, true
}

// HasSessionTimeout returns a boolean if a field has been set.
func (o *MasterBGDetail) HasSessionTimeout() bool {
	if o != nil && o.SessionTimeout != nil {
		return true
	}

	return false
}

// SetSessionTimeout gets a reference to the given int32 and assigns it to the SessionTimeout field.
func (o *MasterBGDetail) SetSessionTimeout(v int32) {
	o.SessionTimeout = &v
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *MasterBGDetail) GetSubscription() Subscription {
	if o == nil || o.Subscription == nil {
		var ret Subscription
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterBGDetail) GetSubscriptionOk() (*Subscription, bool) {
	if o == nil || o.Subscription == nil {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *MasterBGDetail) HasSubscription() bool {
	if o != nil && o.Subscription != nil {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given Subscription and assigns it to the Subscription field.
func (o *MasterBGDetail) SetSubscription(v Subscription) {
	o.Subscription = &v
}

func (o MasterBGDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.Entitlements != nil {
		toSerialize["entitlements"] = o.Entitlements
	}
	if o.Environments != nil {
		toSerialize["environments"] = o.Environments
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IdproviderId != nil {
		toSerialize["idprovider_id"] = o.IdproviderId
	}
	if o.IsAutomaticAdminPromotionExempt != nil {
		toSerialize["isAutomaticAdminPromotionExempt"] = o.IsAutomaticAdminPromotionExempt
	}
	if o.IsFederated != nil {
		toSerialize["isFederated"] = o.IsFederated
	}
	if o.IsMaster != nil {
		toSerialize["isMaster"] = o.IsMaster
	}
	if o.MfaRequired != nil {
		toSerialize["mfaRequired"] = o.MfaRequired
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OwnerId != nil {
		toSerialize["ownerId"] = o.OwnerId
	}
	if o.ParentOrganizationIds != nil {
		toSerialize["parentOrganizationIds"] = o.ParentOrganizationIds
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.SubOrganizationIds != nil {
		toSerialize["subOrganizationIds"] = o.SubOrganizationIds
	}
	if o.TenantOrganizationIds != nil {
		toSerialize["tenantOrganizationIds"] = o.TenantOrganizationIds
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.SessionTimeout != nil {
		toSerialize["sessionTimeout"] = o.SessionTimeout
	}
	if o.Subscription != nil {
		toSerialize["subscription"] = o.Subscription
	}
	return json.Marshal(toSerialize)
}

type NullableMasterBGDetail struct {
	value *MasterBGDetail
	isSet bool
}

func (v NullableMasterBGDetail) Get() *MasterBGDetail {
	return v.value
}

func (v *NullableMasterBGDetail) Set(val *MasterBGDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableMasterBGDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableMasterBGDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMasterBGDetail(val *MasterBGDetail) *NullableMasterBGDetail {
	return &NullableMasterBGDetail{value: val, isSet: true}
}

func (v NullableMasterBGDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMasterBGDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


