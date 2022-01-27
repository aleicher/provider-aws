/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
)

// +kubebuilder:skipversion
type APIKey_SDK struct {
	CreatedDate *metav1.Time `json:"createdDate,omitempty"`

	CustomerID *string `json:"customerID,omitempty"`

	Description *string `json:"description,omitempty"`

	Enabled *bool `json:"enabled,omitempty"`

	ID *string `json:"id,omitempty"`

	LastUpdatedDate *metav1.Time `json:"lastUpdatedDate,omitempty"`

	Name *string `json:"name,omitempty"`

	StageKeys []*string `json:"stageKeys,omitempty"`

	Tags map[string]*string `json:"tags,omitempty"`

	Value *string `json:"value,omitempty"`
}

// +kubebuilder:skipversion
type APIStage struct {
	APIID *string `json:"apiID,omitempty"`

	Stage *string `json:"stage,omitempty"`

	Throttle map[string]*ThrottleSettings `json:"throttle,omitempty"`
}

// +kubebuilder:skipversion
type AccessLogSettings struct {
	DestinationARN *string `json:"destinationARN,omitempty"`

	Format *string `json:"format,omitempty"`
}

// +kubebuilder:skipversion
type Authorizer_SDK struct {
	AuthType *string `json:"authType,omitempty"`

	AuthorizerCredentials *string `json:"authorizerCredentials,omitempty"`

	AuthorizerResultTtlInSeconds *int64 `json:"authorizerResultTtlInSeconds,omitempty"`

	AuthorizerURI *string `json:"authorizerURI,omitempty"`

	ID *string `json:"id,omitempty"`

	IdentitySource *string `json:"identitySource,omitempty"`

	IdentityValidationExpression *string `json:"identityValidationExpression,omitempty"`

	Name *string `json:"name,omitempty"`

	ProviderARNs []*string `json:"providerARNs,omitempty"`
	// The authorizer type. Valid values are TOKEN for a Lambda function using a
	// single authorization token submitted in a custom header, REQUEST for a Lambda
	// function using incoming request parameters, and COGNITO_USER_POOLS for using
	// an Amazon Cognito user pool.
	Type *string `json:"type_,omitempty"`
}

// +kubebuilder:skipversion
type BasePathMapping_SDK struct {
	BasePath *string `json:"basePath,omitempty"`

	RestAPIID *string `json:"restAPIID,omitempty"`

	Stage *string `json:"stage,omitempty"`
}

// +kubebuilder:skipversion
type CanarySettings struct {
	DeploymentID *string `json:"deploymentID,omitempty"`

	PercentTraffic *float64 `json:"percentTraffic,omitempty"`

	StageVariableOverrides map[string]*string `json:"stageVariableOverrides,omitempty"`

	UseStageCache *bool `json:"useStageCache,omitempty"`
}

// +kubebuilder:skipversion
type DeploymentCanarySettings struct {
	PercentTraffic *float64 `json:"percentTraffic,omitempty"`

	StageVariableOverrides map[string]*string `json:"stageVariableOverrides,omitempty"`

	UseStageCache *bool `json:"useStageCache,omitempty"`
}

// +kubebuilder:skipversion
type Deployment_SDK struct {
	APISummary map[string]map[string]*MethodSnapshot `json:"apiSummary,omitempty"`

	CreatedDate *metav1.Time `json:"createdDate,omitempty"`

	Description *string `json:"description,omitempty"`

	ID *string `json:"id,omitempty"`
}

// +kubebuilder:skipversion
type DocumentationPartLocation struct {
	Method *string `json:"method,omitempty"`

	Name *string `json:"name,omitempty"`

	Path *string `json:"path,omitempty"`

	StatusCode *string `json:"statusCode,omitempty"`

	Type *string `json:"type_,omitempty"`
}

// +kubebuilder:skipversion
type DocumentationPart_SDK struct {
	ID *string `json:"id,omitempty"`
	// Specifies the target API entity to which the documentation applies.
	Location *DocumentationPartLocation `json:"location,omitempty"`

	Properties *string `json:"properties,omitempty"`
}

// +kubebuilder:skipversion
type DocumentationVersion_SDK struct {
	CreatedDate *metav1.Time `json:"createdDate,omitempty"`

	Description *string `json:"description,omitempty"`

	Version *string `json:"version,omitempty"`
}

// +kubebuilder:skipversion
type DomainName_SDK struct {
	CertificateARN *string `json:"certificateARN,omitempty"`

	CertificateName *string `json:"certificateName,omitempty"`

	CertificateUploadDate *metav1.Time `json:"certificateUploadDate,omitempty"`

	DistributionDomainName *string `json:"distributionDomainName,omitempty"`

	DistributionHostedZoneID *string `json:"distributionHostedZoneID,omitempty"`

	DomainName *string `json:"domainName,omitempty"`

	DomainNameStatus *string `json:"domainNameStatus,omitempty"`

	DomainNameStatusMessage *string `json:"domainNameStatusMessage,omitempty"`
	// The endpoint configuration to indicate the types of endpoints an API (RestApi)
	// or its custom domain name (DomainName) has.
	EndpointConfiguration *EndpointConfiguration `json:"endpointConfiguration,omitempty"`
	// If specified, API Gateway performs two-way authentication between the client
	// and the server. Clients must present a trusted certificate to access your
	// custom domain name.
	MutualTLSAuthentication *MutualTLSAuthentication `json:"mutualTLSAuthentication,omitempty"`

	RegionalCertificateARN *string `json:"regionalCertificateARN,omitempty"`

	RegionalCertificateName *string `json:"regionalCertificateName,omitempty"`

	RegionalDomainName *string `json:"regionalDomainName,omitempty"`

	RegionalHostedZoneID *string `json:"regionalHostedZoneID,omitempty"`

	SecurityPolicy *string `json:"securityPolicy,omitempty"`

	Tags map[string]*string `json:"tags,omitempty"`
}

// +kubebuilder:skipversion
type EndpointConfiguration struct {
	Types []*string `json:"types,omitempty"`

	VPCEndpointIDs []*string `json:"vpcEndpointIDs,omitempty"`
}

// +kubebuilder:skipversion
type Integration struct {
	CacheKeyParameters []*string `json:"cacheKeyParameters,omitempty"`

	CacheNamespace *string `json:"cacheNamespace,omitempty"`

	ConnectionID *string `json:"connectionID,omitempty"`

	ConnectionType *string `json:"connectionType,omitempty"`

	ContentHandling *string `json:"contentHandling,omitempty"`

	Credentials *string `json:"credentials,omitempty"`

	HTTPMethod *string `json:"httpMethod,omitempty"`

	IntegrationResponses map[string]*IntegrationResponse `json:"integrationResponses,omitempty"`

	PassthroughBehavior *string `json:"passthroughBehavior,omitempty"`

	RequestParameters map[string]*string `json:"requestParameters,omitempty"`

	RequestTemplates map[string]*string `json:"requestTemplates,omitempty"`

	TimeoutInMillis *int64 `json:"timeoutInMillis,omitempty"`

	TLSConfig *TLSConfig `json:"tlsConfig,omitempty"`
	// The integration type. The valid value is HTTP for integrating an API method
	// with an HTTP backend; AWS with any AWS service endpoints; MOCK for testing
	// without actually invoking the backend; HTTP_PROXY for integrating with the
	// HTTP proxy integration; AWS_PROXY for integrating with the Lambda proxy integration.
	Type *string `json:"type_,omitempty"`

	URI *string `json:"uri,omitempty"`
}

// +kubebuilder:skipversion
type IntegrationResponse struct {
	ContentHandling *string `json:"contentHandling,omitempty"`

	ResponseParameters map[string]*string `json:"responseParameters,omitempty"`

	ResponseTemplates map[string]*string `json:"responseTemplates,omitempty"`

	SelectionPattern *string `json:"selectionPattern,omitempty"`
	// The status code.
	StatusCode *string `json:"statusCode,omitempty"`
}

// +kubebuilder:skipversion
type Method struct {
	APIKeyRequired *bool `json:"apiKeyRequired,omitempty"`

	AuthorizationScopes []*string `json:"authorizationScopes,omitempty"`

	AuthorizationType *string `json:"authorizationType,omitempty"`

	AuthorizerID *string `json:"authorizerID,omitempty"`

	HTTPMethod *string `json:"httpMethod,omitempty"`
	// Represents an HTTP, HTTP_PROXY, AWS, AWS_PROXY, or Mock integration.
	//
	// In the API Gateway console, the built-in Lambda integration is an AWS integration.
	//
	// Creating an API (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-create-api.html)
	MethodIntegration *Integration `json:"methodIntegration,omitempty"`

	MethodResponses map[string]*MethodResponse `json:"methodResponses,omitempty"`

	OperationName *string `json:"operationName,omitempty"`

	RequestModels map[string]*string `json:"requestModels,omitempty"`

	RequestParameters map[string]*bool `json:"requestParameters,omitempty"`

	RequestValidatorID *string `json:"requestValidatorID,omitempty"`
}

// +kubebuilder:skipversion
type MethodResponse struct {
	ResponseModels map[string]*string `json:"responseModels,omitempty"`

	ResponseParameters map[string]*bool `json:"responseParameters,omitempty"`
	// The status code.
	StatusCode *string `json:"statusCode,omitempty"`
}

// +kubebuilder:skipversion
type MethodSetting struct {
	CacheDataEncrypted *bool `json:"cacheDataEncrypted,omitempty"`

	CacheTtlInSeconds *int64 `json:"cacheTtlInSeconds,omitempty"`

	CachingEnabled *bool `json:"cachingEnabled,omitempty"`

	DataTraceEnabled *bool `json:"dataTraceEnabled,omitempty"`

	LoggingLevel *string `json:"loggingLevel,omitempty"`

	MetricsEnabled *bool `json:"metricsEnabled,omitempty"`

	RequireAuthorizationForCacheControl *bool `json:"requireAuthorizationForCacheControl,omitempty"`

	ThrottlingBurstLimit *int64 `json:"throttlingBurstLimit,omitempty"`

	ThrottlingRateLimit *float64 `json:"throttlingRateLimit,omitempty"`

	UnauthorizedCacheControlHeaderStrategy *string `json:"unauthorizedCacheControlHeaderStrategy,omitempty"`
}

// +kubebuilder:skipversion
type MethodSnapshot struct {
	APIKeyRequired *bool `json:"apiKeyRequired,omitempty"`

	AuthorizationType *string `json:"authorizationType,omitempty"`
}

// +kubebuilder:skipversion
type Model_SDK struct {
	ContentType *string `json:"contentType,omitempty"`

	Description *string `json:"description,omitempty"`

	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Schema *string `json:"schema,omitempty"`
}

// +kubebuilder:skipversion
type MutualTLSAuthentication struct {
	TruststoreURI *string `json:"truststoreURI,omitempty"`

	TruststoreVersion *string `json:"truststoreVersion,omitempty"`

	TruststoreWarnings []*string `json:"truststoreWarnings,omitempty"`
}

// +kubebuilder:skipversion
type MutualTLSAuthenticationInput struct {
	TruststoreURI *string `json:"truststoreURI,omitempty"`

	TruststoreVersion *string `json:"truststoreVersion,omitempty"`
}

// +kubebuilder:skipversion
type PatchOperation struct {
	From *string `json:"from,omitempty"`

	Op *string `json:"op,omitempty"`

	Path *string `json:"path,omitempty"`

	Value *string `json:"value,omitempty"`
}

// +kubebuilder:skipversion
type QuotaSettings struct {
	Limit *int64 `json:"limit,omitempty"`

	Offset *int64 `json:"offset,omitempty"`

	Period *string `json:"period,omitempty"`
}

// +kubebuilder:skipversion
type Resource_SDK struct {
	ID *string `json:"id,omitempty"`

	ParentID *string `json:"parentID,omitempty"`

	Path *string `json:"path,omitempty"`

	PathPart *string `json:"pathPart,omitempty"`

	ResourceMethods map[string]*Method `json:"resourceMethods,omitempty"`
}

// +kubebuilder:skipversion
type RestAPI_SDK struct {
	APIKeySource *string `json:"apiKeySource,omitempty"`

	BinaryMediaTypes []*string `json:"binaryMediaTypes,omitempty"`

	CreatedDate *metav1.Time `json:"createdDate,omitempty"`

	Description *string `json:"description,omitempty"`

	DisableExecuteAPIEndpoint *bool `json:"disableExecuteAPIEndpoint,omitempty"`
	// The endpoint configuration to indicate the types of endpoints an API (RestApi)
	// or its custom domain name (DomainName) has.
	EndpointConfiguration *EndpointConfiguration `json:"endpointConfiguration,omitempty"`

	ID *string `json:"id,omitempty"`

	MinimumCompressionSize *int64 `json:"minimumCompressionSize,omitempty"`

	Name *string `json:"name,omitempty"`

	Policy *string `json:"policy,omitempty"`

	Tags map[string]*string `json:"tags,omitempty"`

	Version *string `json:"version,omitempty"`

	Warnings []*string `json:"warnings,omitempty"`
}

// +kubebuilder:skipversion
type SDKConfigurationProperty struct {
	DefaultValue *string `json:"defaultValue,omitempty"`

	Description *string `json:"description,omitempty"`

	FriendlyName *string `json:"friendlyName,omitempty"`

	Name *string `json:"name,omitempty"`

	Required *bool `json:"required,omitempty"`
}

// +kubebuilder:skipversion
type StageKey struct {
	RestAPIID *string `json:"restAPIID,omitempty"`

	StageName *string `json:"stageName,omitempty"`
}

// +kubebuilder:skipversion
type Stage_SDK struct {
	// Access log settings, including the access log format and access log destination
	// ARN.
	AccessLogSettings *AccessLogSettings `json:"accessLogSettings,omitempty"`

	CacheClusterEnabled *bool `json:"cacheClusterEnabled,omitempty"`
	// Returns the size of the CacheCluster.
	CacheClusterSize *string `json:"cacheClusterSize,omitempty"`
	// Returns the status of the CacheCluster.
	CacheClusterStatus *string `json:"cacheClusterStatus,omitempty"`
	// Configuration settings of a canary deployment.
	CanarySettings *CanarySettings `json:"canarySettings,omitempty"`

	ClientCertificateID *string `json:"clientCertificateID,omitempty"`

	CreatedDate *metav1.Time `json:"createdDate,omitempty"`

	DeploymentID *string `json:"deploymentID,omitempty"`

	Description *string `json:"description,omitempty"`

	DocumentationVersion *string `json:"documentationVersion,omitempty"`

	LastUpdatedDate *metav1.Time `json:"lastUpdatedDate,omitempty"`

	MethodSettings map[string]*MethodSetting `json:"methodSettings,omitempty"`

	StageName *string `json:"stageName,omitempty"`

	Tags map[string]*string `json:"tags,omitempty"`

	TracingEnabled *bool `json:"tracingEnabled,omitempty"`

	Variables map[string]*string `json:"variables,omitempty"`

	WebACLARN *string `json:"webACLARN,omitempty"`
}

// +kubebuilder:skipversion
type TLSConfig struct {
	InsecureSkipVerification *bool `json:"insecureSkipVerification,omitempty"`
}

// +kubebuilder:skipversion
type ThrottleSettings struct {
	BurstLimit *int64 `json:"burstLimit,omitempty"`

	RateLimit *float64 `json:"rateLimit,omitempty"`
}

// +kubebuilder:skipversion
type UpdateRequestValidatorOutput struct {
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	ValidateRequestBody *bool `json:"validateRequestBody,omitempty"`

	ValidateRequestParameters *bool `json:"validateRequestParameters,omitempty"`
}

// +kubebuilder:skipversion
type UpdateVPCLinkOutput struct {
	Description *string `json:"description,omitempty"`

	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Status *string `json:"status,omitempty"`

	StatusMessage *string `json:"statusMessage,omitempty"`

	Tags map[string]*string `json:"tags,omitempty"`

	TargetARNs []*string `json:"targetARNs,omitempty"`
}

// +kubebuilder:skipversion
type UsagePlanKey_SDK struct {
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Type *string `json:"type_,omitempty"`

	Value *string `json:"value,omitempty"`
}

// +kubebuilder:skipversion
type UsagePlan_SDK struct {
	APIStages []*APIStage `json:"apiStages,omitempty"`

	Description *string `json:"description,omitempty"`

	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	ProductCode *string `json:"productCode,omitempty"`
	// Quotas configured for a usage plan.
	Quota *QuotaSettings `json:"quota,omitempty"`

	Tags map[string]*string `json:"tags,omitempty"`
	// The API request rate limits.
	Throttle *ThrottleSettings `json:"throttle,omitempty"`
}
