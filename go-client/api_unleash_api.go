/*
 * GitLab API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type UnleashApiApiService service

/*
UnleashApiApiService
Get a list of features
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param projectId The ID of a project
 * @param optional nil or *UnleashApiApiGetApiV4FeatureFlagsUnleashProjectIdClientFeaturesOpts - Optional Parameters:
     * @param "InstanceId" (optional.String) -  The instance ID of Unleash Client
     * @param "AppName" (optional.String) -  The application name of Unleash Client


*/

type UnleashApiApiGetApiV4FeatureFlagsUnleashProjectIdClientFeaturesOpts struct {
	InstanceId optional.String
	AppName    optional.String
}

func (a *UnleashApiApiService) GetApiV4FeatureFlagsUnleashProjectIdClientFeatures(ctx context.Context, projectId string, localVarOptionals *UnleashApiApiGetApiV4FeatureFlagsUnleashProjectIdClientFeaturesOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v4/feature_flags/unleash/{project_id}/client/features"
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", fmt.Sprintf("%v", projectId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.InstanceId.IsSet() {
		localVarQueryParams.Add("instance_id", parameterToString(localVarOptionals.InstanceId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AppName.IsSet() {
		localVarQueryParams.Add("app_name", parameterToString(localVarOptionals.AppName.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
UnleashApiApiService
Get a list of features (deprecated, v2 client support)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param projectId The ID of a project
 * @param optional nil or *UnleashApiApiGetApiV4FeatureFlagsUnleashProjectIdFeaturesOpts - Optional Parameters:
     * @param "InstanceId" (optional.String) -  The instance ID of Unleash Client
     * @param "AppName" (optional.String) -  The application name of Unleash Client


*/

type UnleashApiApiGetApiV4FeatureFlagsUnleashProjectIdFeaturesOpts struct {
	InstanceId optional.String
	AppName    optional.String
}

func (a *UnleashApiApiService) GetApiV4FeatureFlagsUnleashProjectIdFeatures(ctx context.Context, projectId string, localVarOptionals *UnleashApiApiGetApiV4FeatureFlagsUnleashProjectIdFeaturesOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v4/feature_flags/unleash/{project_id}/features"
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", fmt.Sprintf("%v", projectId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.InstanceId.IsSet() {
		localVarQueryParams.Add("instance_id", parameterToString(localVarOptionals.InstanceId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AppName.IsSet() {
		localVarQueryParams.Add("app_name", parameterToString(localVarOptionals.AppName.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
