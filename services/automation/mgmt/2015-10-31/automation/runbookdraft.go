package automation

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"io"
	"net/http"
)

// RunbookDraftClient is the automation Client
type RunbookDraftClient struct {
	BaseClient
}

// NewRunbookDraftClient creates an instance of the RunbookDraftClient client.
func NewRunbookDraftClient(subscriptionID string, resourceGroupName string) RunbookDraftClient {
	return NewRunbookDraftClientWithBaseURI(DefaultBaseURI, subscriptionID, resourceGroupName)
}

// NewRunbookDraftClientWithBaseURI creates an instance of the RunbookDraftClient client.
func NewRunbookDraftClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string) RunbookDraftClient {
	return RunbookDraftClient{NewWithBaseURI(baseURI, subscriptionID, resourceGroupName)}
}

// CreateOrUpdate updates the runbook draft with runbookStream as its content.
//
// automationAccountName is the automation account name. runbookName is the runbook name. runbookContent is the runbook
// draft content. runbookContent will be closed upon successful return. Callers should ensure closure when receiving an
// error.
func (client RunbookDraftClient) CreateOrUpdate(ctx context.Context, automationAccountName string, runbookName string, runbookContent io.ReadCloser) (result RunbookDraftCreateOrUpdateFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.ResourceGroupName,
			Constraints: []validation.Constraint{{Target: "client.ResourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "automation.RunbookDraftClient", "CreateOrUpdate")
	}

	req, err := client.CreateOrUpdatePreparer(ctx, automationAccountName, runbookName, runbookContent)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client RunbookDraftClient) CreateOrUpdatePreparer(ctx context.Context, automationAccountName string, runbookName string, runbookContent io.ReadCloser) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"resourceGroupName":     autorest.Encode("path", client.ResourceGroupName),
		"runbookName":           autorest.Encode("path", runbookName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsOctetStream(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/content", pathParameters),
		autorest.WithFile(runbookContent),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client RunbookDraftClient) CreateOrUpdateSender(req *http.Request) (future RunbookDraftCreateOrUpdateFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client RunbookDraftClient) CreateOrUpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieve the runbook draft identified by runbook name.
//
// automationAccountName is the automation account name. runbookName is the runbook name.
func (client RunbookDraftClient) Get(ctx context.Context, automationAccountName string, runbookName string) (result RunbookDraft, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.ResourceGroupName,
			Constraints: []validation.Constraint{{Target: "client.ResourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "automation.RunbookDraftClient", "Get")
	}

	req, err := client.GetPreparer(ctx, automationAccountName, runbookName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RunbookDraftClient) GetPreparer(ctx context.Context, automationAccountName string, runbookName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"resourceGroupName":     autorest.Encode("path", client.ResourceGroupName),
		"runbookName":           autorest.Encode("path", runbookName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RunbookDraftClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RunbookDraftClient) GetResponder(resp *http.Response) (result RunbookDraft, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetContent retrieve the content of runbook draft identified by runbook name.
//
// automationAccountName is the automation account name. runbookName is the runbook name.
func (client RunbookDraftClient) GetContent(ctx context.Context, automationAccountName string, runbookName string) (result ReadCloser, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.ResourceGroupName,
			Constraints: []validation.Constraint{{Target: "client.ResourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "automation.RunbookDraftClient", "GetContent")
	}

	req, err := client.GetContentPreparer(ctx, automationAccountName, runbookName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "GetContent", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetContentSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "GetContent", resp, "Failure sending request")
		return
	}

	result, err = client.GetContentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "GetContent", resp, "Failure responding to request")
	}

	return
}

// GetContentPreparer prepares the GetContent request.
func (client RunbookDraftClient) GetContentPreparer(ctx context.Context, automationAccountName string, runbookName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"resourceGroupName":     autorest.Encode("path", client.ResourceGroupName),
		"runbookName":           autorest.Encode("path", runbookName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/content", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetContentSender sends the GetContent request. The method will close the
// http.Response Body if it receives an error.
func (client RunbookDraftClient) GetContentSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetContentResponder handles the response to the GetContent request. The method always
// closes the http.Response Body.
func (client RunbookDraftClient) GetContentResponder(resp *http.Response) (result ReadCloser, err error) {
	result.Value = &resp.Body
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK))
	result.Response = autorest.Response{Response: resp}
	return
}

// Publish publish runbook draft.
//
// automationAccountName is the automation account name. runbookName is the parameters supplied to the publish runbook
// operation.
func (client RunbookDraftClient) Publish(ctx context.Context, automationAccountName string, runbookName string) (result RunbookDraftPublishFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.ResourceGroupName,
			Constraints: []validation.Constraint{{Target: "client.ResourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "automation.RunbookDraftClient", "Publish")
	}

	req, err := client.PublishPreparer(ctx, automationAccountName, runbookName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "Publish", nil, "Failure preparing request")
		return
	}

	result, err = client.PublishSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "Publish", result.Response(), "Failure sending request")
		return
	}

	return
}

// PublishPreparer prepares the Publish request.
func (client RunbookDraftClient) PublishPreparer(ctx context.Context, automationAccountName string, runbookName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"resourceGroupName":     autorest.Encode("path", client.ResourceGroupName),
		"runbookName":           autorest.Encode("path", runbookName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/publish", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PublishSender sends the Publish request. The method will close the
// http.Response Body if it receives an error.
func (client RunbookDraftClient) PublishSender(req *http.Request) (future RunbookDraftPublishFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	return
}

// PublishResponder handles the response to the Publish request. The method always
// closes the http.Response Body.
func (client RunbookDraftClient) PublishResponder(resp *http.Response) (result Runbook, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UndoEdit retrieve the runbook identified by runbook name.
//
// automationAccountName is the automation account name. runbookName is the runbook name.
func (client RunbookDraftClient) UndoEdit(ctx context.Context, automationAccountName string, runbookName string) (result RunbookDraftUndoEditResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.ResourceGroupName,
			Constraints: []validation.Constraint{{Target: "client.ResourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "automation.RunbookDraftClient", "UndoEdit")
	}

	req, err := client.UndoEditPreparer(ctx, automationAccountName, runbookName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "UndoEdit", nil, "Failure preparing request")
		return
	}

	resp, err := client.UndoEditSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "UndoEdit", resp, "Failure sending request")
		return
	}

	result, err = client.UndoEditResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.RunbookDraftClient", "UndoEdit", resp, "Failure responding to request")
	}

	return
}

// UndoEditPreparer prepares the UndoEdit request.
func (client RunbookDraftClient) UndoEditPreparer(ctx context.Context, automationAccountName string, runbookName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"resourceGroupName":     autorest.Encode("path", client.ResourceGroupName),
		"runbookName":           autorest.Encode("path", runbookName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-10-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/undoEdit", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UndoEditSender sends the UndoEdit request. The method will close the
// http.Response Body if it receives an error.
func (client RunbookDraftClient) UndoEditSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// UndoEditResponder handles the response to the UndoEdit request. The method always
// closes the http.Response Body.
func (client RunbookDraftClient) UndoEditResponder(resp *http.Response) (result RunbookDraftUndoEditResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
