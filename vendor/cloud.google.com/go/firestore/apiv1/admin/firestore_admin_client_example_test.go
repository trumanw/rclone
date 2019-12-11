// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by gapic-generator. DO NOT EDIT.

package apiv1_test

import (
	"context"

	apiv1 "cloud.google.com/go/firestore/apiv1/admin"
	"google.golang.org/api/iterator"
	adminpb "google.golang.org/genproto/googleapis/firestore/admin/v1"
)

func ExampleNewFirestoreAdminClient() {
	ctx := context.Background()
	c, err := apiv1.NewFirestoreAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleFirestoreAdminClient_CreateIndex() {
	ctx := context.Background()
	c, err := apiv1.NewFirestoreAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &adminpb.CreateIndexRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateIndex(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleFirestoreAdminClient_ListIndexes() {
	ctx := context.Background()
	c, err := apiv1.NewFirestoreAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &adminpb.ListIndexesRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListIndexes(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleFirestoreAdminClient_GetIndex() {
	ctx := context.Background()
	c, err := apiv1.NewFirestoreAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &adminpb.GetIndexRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetIndex(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleFirestoreAdminClient_DeleteIndex() {
	ctx := context.Background()
	c, err := apiv1.NewFirestoreAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &adminpb.DeleteIndexRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteIndex(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleFirestoreAdminClient_ImportDocuments() {
	ctx := context.Background()
	c, err := apiv1.NewFirestoreAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &adminpb.ImportDocumentsRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.ImportDocuments(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleFirestoreAdminClient_ExportDocuments() {
	ctx := context.Background()
	c, err := apiv1.NewFirestoreAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &adminpb.ExportDocumentsRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.ExportDocuments(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleFirestoreAdminClient_GetField() {
	ctx := context.Background()
	c, err := apiv1.NewFirestoreAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &adminpb.GetFieldRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetField(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleFirestoreAdminClient_ListFields() {
	ctx := context.Background()
	c, err := apiv1.NewFirestoreAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &adminpb.ListFieldsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListFields(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleFirestoreAdminClient_UpdateField() {
	ctx := context.Background()
	c, err := apiv1.NewFirestoreAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &adminpb.UpdateFieldRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateField(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
