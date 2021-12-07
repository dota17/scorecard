// Copyright 2021 Security Scorecard Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package giteerepo

import (
	"context"
	"gitee.com/openeuler/go-gitee/gitee"
	"github.com/ossf/scorecard/v3/clients"
)

type statusesHandler struct {
	client *gitee.APIClient
	ctx    context.Context
	owner  string
	repo   string
}

func (handler *statusesHandler) init(ctx context.Context, owner, repo string) {
	handler.ctx = ctx
	handler.owner = owner
	handler.repo = repo
}

func (handler *statusesHandler) listStatuses(ref string) ([]clients.Status, error) {
	//statuses, _, err := handler.client.Repositories.ListStatuses(handler.ctx, handler.owner, handler.repo, ref,
	//	&gitee.ListOptions{})
	//if err != nil {
	//	return nil, sce.WithMessage(sce.ErrScorecardInternal, fmt.Sprintf("ListStatuses: %v", err))
	//}
	return statusesFrom(), nil
}

func statusesFrom() []clients.Status {
	var statuses []clients.Status
	//for _, status := range data {
	//	statuses = append(statuses, clients.Status{
	//		State:   status.GetState(),
	//		Context: status.GetContext(),
	//		URL:     status.GetURL(),
	//	})
	//}
	return statuses
}
