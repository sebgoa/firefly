// Copyright © 2021 Kaleido, Inc.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sqlcommon

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kaleido-io/firefly/internal/log"
	"github.com/kaleido-io/firefly/pkg/database"
	"github.com/kaleido-io/firefly/pkg/fftypes"
	"github.com/stretchr/testify/assert"
)

func TestUpsertGroupE2EWithDB(t *testing.T) {
	log.SetLevel("debug")

	s := newQLTestProvider(t)
	defer s.Close()
	ctx := context.Background()

	// Create a new group
	identity1 := "identity1"
	identity2 := "identity2"
	group := &fftypes.Group{
		ID:        fftypes.NewUUID(),
		Namespace: "ns1",
		Recipients: fftypes.Recipients{
			{Identity: identity1},
			{Identity: identity2},
		},
		Created: fftypes.Now(),
	}
	err := s.UpsertGroup(ctx, group, true)
	assert.NoError(t, err)

	// Check we get the exact same group back
	groupRead, err := s.GetGroupByID(ctx, group.ID)
	assert.NoError(t, err)
	groupJson, _ := json.Marshal(&group)
	groupReadJson, _ := json.Marshal(&groupRead)
	assert.Equal(t, string(groupJson), string(groupReadJson))

	// Update the group (this is testing what's possible at the database layer,
	// and does not account for the verification that happens at the higher level)
	identity3 := "identity3"
	groupUpdated := &fftypes.Group{
		ID:          group.ID,
		Namespace:   "ns1",
		Description: "my group",
		Recipients: fftypes.Recipients{
			{Identity: identity3},
			{Identity: identity2},
		},
		Created: fftypes.Now(),
		Message: fftypes.NewUUID(),
	}

	err = s.UpsertGroup(context.Background(), groupUpdated, true)
	assert.NoError(t, err)

	// Check we get the exact same group back - note the removal of one of the data elements
	groupRead, err = s.GetGroupByID(ctx, group.ID)
	assert.NoError(t, err)
	groupJson, _ = json.Marshal(&groupUpdated)
	groupReadJson, _ = json.Marshal(&groupRead)
	assert.Equal(t, string(groupJson), string(groupReadJson))

	// Query back the group
	fb := database.GroupQueryFactory.NewFilter(ctx)
	filter := fb.And(
		fb.Eq("id", groupUpdated.ID),
		fb.Eq("namespace", groupUpdated.Namespace),
		fb.Eq("description", groupUpdated.Description),
		fb.Eq("message", groupUpdated.Message),
		fb.Gt("created", "0"),
	)
	groups, err := s.GetGroups(ctx, filter)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(groups))
	groupReadJson, _ = json.Marshal(groups[0])
	assert.Equal(t, string(groupJson), string(groupReadJson))

	// Negative test on filter
	filter = fb.And(
		fb.Eq("id", groupUpdated.ID.String()),
		fb.Eq("created", "0"),
	)
	groups, err = s.GetGroups(ctx, filter)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(groups))

	// Update
	up := database.GroupQueryFactory.NewUpdate(ctx).
		Set("description", "new description")
	err = s.UpdateGroup(ctx, group.ID, up)
	assert.NoError(t, err)

	// Test find updated value
	filter = fb.And(
		fb.Eq("id", groupUpdated.ID.String()),
		fb.Eq("description", "new description"),
	)
	groups, err = s.GetGroups(ctx, filter)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(groups))

	s.callbacks.AssertExpectations(t)
}

func TestUpsertGroupFailBegin(t *testing.T) {
	s, mock := newMockProvider().init()
	mock.ExpectBegin().WillReturnError(fmt.Errorf("pop"))
	err := s.UpsertGroup(context.Background(), &fftypes.Group{}, true)
	assert.Regexp(t, "FF10114", err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpsertGroupFailSelect(t *testing.T) {
	s, mock := newMockProvider().init()
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT .*").WillReturnError(fmt.Errorf("pop"))
	mock.ExpectRollback()
	groupID := fftypes.NewUUID()
	err := s.UpsertGroup(context.Background(), &fftypes.Group{ID: groupID}, true)
	assert.Regexp(t, "FF10115", err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpsertGroupFailInsert(t *testing.T) {
	s, mock := newMockProvider().init()
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT .*").WillReturnRows(sqlmock.NewRows([]string{}))
	mock.ExpectExec("INSERT .*").WillReturnError(fmt.Errorf("pop"))
	mock.ExpectRollback()
	groupID := fftypes.NewUUID()
	err := s.UpsertGroup(context.Background(), &fftypes.Group{ID: groupID}, true)
	assert.Regexp(t, "FF10116", err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpsertGroupFailUpdate(t *testing.T) {
	s, mock := newMockProvider().init()
	groupID := fftypes.NewUUID()
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT .*").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(groupID.String()))
	mock.ExpectExec("UPDATE .*").WillReturnError(fmt.Errorf("pop"))
	mock.ExpectRollback()
	err := s.UpsertGroup(context.Background(), &fftypes.Group{ID: groupID}, true)
	assert.Regexp(t, "FF10117", err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpsertGroupFailLoadRecipients(t *testing.T) {
	s, mock := newMockProvider().init()
	groupID := fftypes.NewUUID()
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT .*").WillReturnRows(sqlmock.NewRows([]string{"id"}))
	mock.ExpectExec("INSERT .*").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectQuery("SELECT .*").WillReturnError(fmt.Errorf("pop"))
	mock.ExpectRollback()
	err := s.UpsertGroup(context.Background(), &fftypes.Group{ID: groupID}, true)
	assert.Regexp(t, "FF10115", err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpsertGroupFailCommit(t *testing.T) {
	s, mock := newMockProvider().init()
	groupID := fftypes.NewUUID()
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT .*").WillReturnRows(sqlmock.NewRows([]string{"id"}))
	mock.ExpectExec("INSERT .*").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectQuery("SELECT .*").WillReturnRows(sqlmock.NewRows([]string{"identity"}))
	mock.ExpectCommit().WillReturnError(fmt.Errorf("pop"))
	err := s.UpsertGroup(context.Background(), &fftypes.Group{ID: groupID}, true)
	assert.Regexp(t, "FF10119", err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetRecipientsScanFail(t *testing.T) {
	s, mock := newMockProvider().init()
	groupID := fftypes.NewUUID()
	mock.ExpectBegin()
	tx, _ := s.db.Begin()
	mock.ExpectQuery("SELECT .*").WillReturnRows(sqlmock.NewRows([]string{"identity"}).AddRow("not the uuid you are looking for"))
	_, err := s.getRecipients(context.Background(), &txWrapper{sqlTX: tx}, groupID)
	assert.Regexp(t, "FF10121", err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateRecipientsBlankIdentity(t *testing.T) {
	s, mock := newMockProvider().init()
	groupID := fftypes.NewUUID()
	mock.ExpectBegin()
	tx, _ := s.db.Begin()
	mock.ExpectQuery("SELECT .*").WillReturnRows(sqlmock.NewRows([]string{"identity", "idx"}).AddRow("id1", 1))
	err := s.updateRecipients(context.Background(), &txWrapper{sqlTX: tx}, &fftypes.Group{
		ID:         groupID,
		Recipients: fftypes.Recipients{{Identity: ""}},
	})
	assert.Regexp(t, "FF10220", err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateGroupDataDeleteFail(t *testing.T) {
	s, mock := newMockProvider().init()
	groupID := fftypes.NewUUID()
	mock.ExpectBegin()
	tx, _ := s.db.Begin()
	mock.ExpectQuery("SELECT .*").WillReturnRows(sqlmock.NewRows([]string{"identity", "idx"}).AddRow("id1", 1))
	mock.ExpectExec("DELETE .*").WillReturnError(fmt.Errorf("pop"))
	err := s.updateRecipients(context.Background(), &txWrapper{sqlTX: tx}, &fftypes.Group{
		ID: groupID,
	})
	assert.Regexp(t, "FF10118", err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateGroupDataAddFail(t *testing.T) {
	s, mock := newMockProvider().init()
	groupID := fftypes.NewUUID()
	mock.ExpectBegin()
	tx, _ := s.db.Begin()
	mock.ExpectQuery("SELECT .*").WillReturnRows(sqlmock.NewRows([]string{"identity"}))
	mock.ExpectExec("INSERT .*").WillReturnError(fmt.Errorf("pop"))
	err := s.updateRecipients(context.Background(), &txWrapper{sqlTX: tx}, &fftypes.Group{
		ID:         groupID,
		Recipients: fftypes.Recipients{{Identity: "id1"}},
	})
	assert.Regexp(t, "FF10116", err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateGroupDataSwitchIDxFail(t *testing.T) {
	s, mock := newMockProvider().init()
	groupID := fftypes.NewUUID()
	mock.ExpectBegin()
	tx, _ := s.db.Begin()
	mock.ExpectQuery("SELECT .*").WillReturnRows(sqlmock.NewRows([]string{"identity", "idx"})).
		WillReturnRows(sqlmock.NewRows(
			[]string{"identity", "idx"},
		).AddRow(
			"id1", 0,
		))
	mock.ExpectExec("INSERT .*").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("UPDATE .*").WillReturnError(fmt.Errorf("pop"))
	err := s.updateRecipients(context.Background(), &txWrapper{sqlTX: tx}, &fftypes.Group{
		ID:         groupID,
		Recipients: fftypes.Recipients{{Identity: "id2"}, {Identity: "id1"}},
	})
	assert.Regexp(t, "FF10117", err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestLoadRecipientsQueryFail(t *testing.T) {
	s, mock := newMockProvider().init()
	groupID := fftypes.NewUUID()
	mock.ExpectQuery("SELECT .*").WillReturnError(fmt.Errorf("pop"))
	err := s.loadRecipients(context.Background(), []*fftypes.Group{{ID: groupID}})
	assert.Regexp(t, "FF10115", err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestLoadRecipientsScanFail(t *testing.T) {
	s, mock := newMockProvider().init()
	groupID := fftypes.NewUUID()
	mock.ExpectQuery("SELECT .*").WillReturnRows(sqlmock.NewRows([]string{"identity"}).AddRow("only one"))
	err := s.loadRecipients(context.Background(), []*fftypes.Group{{ID: groupID}})
	assert.Regexp(t, "FF10121", err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestLoadRecipientsEmpty(t *testing.T) {
	s, mock := newMockProvider().init()
	groupID := fftypes.NewUUID()
	group := &fftypes.Group{ID: groupID}
	mock.ExpectQuery("SELECT .*").WillReturnRows(sqlmock.NewRows([]string{"identity"}))
	err := s.loadRecipients(context.Background(), []*fftypes.Group{group})
	assert.NoError(t, err)
	assert.Equal(t, fftypes.Recipients{}, group.Recipients)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetGroupByIDSelectFail(t *testing.T) {
	s, mock := newMockProvider().init()
	groupID := fftypes.NewUUID()
	mock.ExpectQuery("SELECT .*").WillReturnError(fmt.Errorf("pop"))
	_, err := s.GetGroupByID(context.Background(), groupID)
	assert.Regexp(t, "FF10115", err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetGroupByIDNotFound(t *testing.T) {
	s, mock := newMockProvider().init()
	groupID := fftypes.NewUUID()
	mock.ExpectQuery("SELECT .*").WillReturnRows(sqlmock.NewRows([]string{"id"}))
	group, err := s.GetGroupByID(context.Background(), groupID)
	assert.NoError(t, err)
	assert.Nil(t, group)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetGroupByIDScanFail(t *testing.T) {
	s, mock := newMockProvider().init()
	groupID := fftypes.NewUUID()
	mock.ExpectQuery("SELECT .*").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("only one"))
	_, err := s.GetGroupByID(context.Background(), groupID)
	assert.Regexp(t, "FF10121", err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetGroupByIDLoadRecipientsFail(t *testing.T) {
	s, mock := newMockProvider().init()
	groupID := fftypes.NewUUID()
	mock.ExpectQuery("SELECT .*").WillReturnRows(sqlmock.NewRows(groupColumns).
		AddRow(groupID.String(), nil, "ns1", "", fftypes.Now()))
	mock.ExpectQuery("SELECT .*").WillReturnError(fmt.Errorf("pop"))
	_, err := s.GetGroupByID(context.Background(), groupID)
	assert.Regexp(t, "FF10115", err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetGroupsBuildQueryFail(t *testing.T) {
	s, _ := newMockProvider().init()
	f := database.GroupQueryFactory.NewFilter(context.Background()).Eq("id", map[bool]bool{true: false})
	_, err := s.GetGroups(context.Background(), f)
	assert.Regexp(t, "FF10149.*id", err)
}

func TestGetGroupsQueryFail(t *testing.T) {
	s, mock := newMockProvider().init()
	mock.ExpectQuery("SELECT .*").WillReturnError(fmt.Errorf("pop"))
	f := database.GroupQueryFactory.NewFilter(context.Background()).Eq("id", "")
	_, err := s.GetGroups(context.Background(), f)
	assert.Regexp(t, "FF10115", err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetGroupsReadGroupFail(t *testing.T) {
	s, mock := newMockProvider().init()
	mock.ExpectQuery("SELECT .*").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("only one"))
	f := database.GroupQueryFactory.NewFilter(context.Background()).Eq("id", "")
	_, err := s.GetGroups(context.Background(), f)
	assert.Regexp(t, "FF10121", err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetGroupsLoadRecipientsFail(t *testing.T) {
	s, mock := newMockProvider().init()
	groupID := fftypes.NewUUID()
	mock.ExpectQuery("SELECT .*").WillReturnRows(sqlmock.NewRows(groupColumns).
		AddRow(groupID.String(), nil, "ns1", "", fftypes.Now()))
	mock.ExpectQuery("SELECT .*").WillReturnError(fmt.Errorf("pop"))
	f := database.GroupQueryFactory.NewFilter(context.Background()).Gt("created", "0")
	_, err := s.GetGroups(context.Background(), f)
	assert.Regexp(t, "FF10115", err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGroupUpdateBeginFail(t *testing.T) {
	s, mock := newMockProvider().init()
	mock.ExpectBegin().WillReturnError(fmt.Errorf("pop"))
	u := database.GroupQueryFactory.NewUpdate(context.Background()).Set("id", "anything")
	err := s.UpdateGroup(context.Background(), fftypes.NewUUID(), u)
	assert.Regexp(t, "FF10114", err)
}

func TestGroupUpdateBuildQueryFail(t *testing.T) {
	s, mock := newMockProvider().init()
	mock.ExpectBegin()
	u := database.GroupQueryFactory.NewUpdate(context.Background()).Set("id", map[bool]bool{true: false})
	err := s.UpdateGroup(context.Background(), fftypes.NewUUID(), u)
	assert.Regexp(t, "FF10149.*id", err)
}

func TestGroupsUpdateBuildFilterFail(t *testing.T) {
	s, mock := newMockProvider().init()
	mock.ExpectBegin()
	f := database.GroupQueryFactory.NewFilter(context.Background()).Eq("id", map[bool]bool{true: false})
	u := database.GroupQueryFactory.NewUpdate(context.Background()).Set("description", "my desc")
	err := s.UpdateGroups(context.Background(), f, u)
	assert.Regexp(t, "FF10149.*id", err)
}

func TestGroupUpdateFail(t *testing.T) {
	s, mock := newMockProvider().init()
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE .*").WillReturnError(fmt.Errorf("pop"))
	mock.ExpectRollback()
	u := database.GroupQueryFactory.NewUpdate(context.Background()).Set("description", fftypes.NewUUID())
	err := s.UpdateGroup(context.Background(), fftypes.NewUUID(), u)
	assert.Regexp(t, "FF10117", err)
}
