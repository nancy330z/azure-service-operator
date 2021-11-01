// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

type ThroughputSettingsGetResults_StatusARM struct {
	//Id: The unique resource identifier of the ARM resource.
	Id *string `json:"id,omitempty"`

	//Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	//Name: The name of the ARM resource.
	Name *string `json:"name,omitempty"`

	//Properties: The properties of an Azure Cosmos DB resource throughput
	Properties *ThroughputSettingsGetProperties_StatusARM `json:"properties,omitempty"`
	Tags       map[string]string                          `json:"tags,omitempty"`

	//Type: The type of Azure resource.
	Type *string `json:"type,omitempty"`
}

type ThroughputSettingsGetProperties_StatusARM struct {
	Resource *ThroughputSettingsGetProperties_Status_ResourceARM `json:"resource,omitempty"`
}

type ThroughputSettingsGetProperties_Status_ResourceARM struct {
	//AutoscaleSettings: Cosmos DB resource for autoscale settings. Either throughput
	//is required or autoscaleSettings is required, but not both.
	AutoscaleSettings *AutoscaleSettingsResource_StatusARM `json:"autoscaleSettings,omitempty"`

	//Etag: A system generated property representing the resource etag required for
	//optimistic concurrency control.
	Etag *string `json:"_etag,omitempty"`

	//MinimumThroughput: The minimum throughput of the resource
	MinimumThroughput *string `json:"minimumThroughput,omitempty"`

	//OfferReplacePending: The throughput replace is pending
	OfferReplacePending *string `json:"offerReplacePending,omitempty"`

	//Rid: A system generated property. A unique identifier.
	Rid *string `json:"_rid,omitempty"`

	//Throughput: Value of the Cosmos DB resource throughput. Either throughput is
	//required or autoscaleSettings is required, but not both.
	Throughput *int `json:"throughput,omitempty"`

	//Ts: A system generated property that denotes the last updated timestamp of the
	//resource.
	Ts *float64 `json:"_ts,omitempty"`
}

type AutoscaleSettingsResource_StatusARM struct {
	//AutoUpgradePolicy: Cosmos DB resource auto-upgrade policy
	AutoUpgradePolicy *AutoUpgradePolicyResource_StatusARM `json:"autoUpgradePolicy,omitempty"`

	//MaxThroughput: Represents maximum throughput container can scale up to.
	MaxThroughput int `json:"maxThroughput"`

	//TargetMaxThroughput: Represents target maximum throughput container can scale up
	//to once offer is no longer in pending state.
	TargetMaxThroughput *int `json:"targetMaxThroughput,omitempty"`
}

type AutoUpgradePolicyResource_StatusARM struct {
	//ThroughputPolicy: Represents throughput policy which service must adhere to for
	//auto-upgrade
	ThroughputPolicy *ThroughputPolicyResource_StatusARM `json:"throughputPolicy,omitempty"`
}

type ThroughputPolicyResource_StatusARM struct {
	//IncrementPercent: Represents the percentage by which throughput can increase
	//every time throughput policy kicks in.
	IncrementPercent *int `json:"incrementPercent,omitempty"`

	//IsEnabled: Determines whether the ThroughputPolicy is active or not
	IsEnabled *bool `json:"isEnabled,omitempty"`
}
