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
	"fmt"
	"gitee.com/openeuler/go-gitee/gitee"
	"sync"


	"github.com/ossf/scorecard/v3/clients"
)

type contributorsHandler struct {
	ghClient     *gitee.APIClient
	once         *sync.Once
	ctx          context.Context
	errSetup     error
	owner        string
	repo         string
	contributors []clients.Contributor
}

func (handler *contributorsHandler) init(ctx context.Context, owner, repo string) {
	handler.ctx = ctx
	handler.owner = owner
	handler.repo = repo
	handler.errSetup = nil
	handler.once = new(sync.Once)
}

func (handler *contributorsHandler) setup() error {
	handler.once.Do(func() {
		contribs, _, err := handler.ghClient.RepositoriesApi.GetV5ReposOwnerRepoContributors(
			handler.ctx, handler.owner, handler.repo, &gitee.GetV5ReposOwnerRepoContributorsOpts{})
		if err != nil {
			handler.errSetup = fmt.Errorf("error during ListContributors: %w", err)
		}
		if contribs.Name != "" {

		}
		//for _, contrib := range contribs {
		//	if contrib.Login() == "" {
		//		continue
		//	}
		//	contributor := clients.Contributor{
		//		NumContributions: contrib.GetContributions(),
		//		User: clients.User{
		//			Login: contrib.GetLogin(),
		//		},
		//	}
		//	orgs, _, err := handler.ghClient.OrganizationsApi.GetV5UsersUsernameOrgs(handler.ctx, contrib.Login(), nil)
		//	// This call can fail due to token scopes. So ignore error.
		//	if err == nil {
		//		for _, org := range orgs {
		//			contributor.Organizations = append(contributor.Organizations, clients.User{
		//				Login: org.Login,
		//			})
		//		}
		//	}
		//	user, _, err := handler.ghClient.UsersApi.GetV5UsersUsername(handler.ctx, contrib.Login(),nil)
		//	if err != nil {
		//		handler.errSetup = fmt.Errorf("error during Users.Get: %w", err)
		//	}
		//	contributor.Company = user.Company()
		//	handler.contributors = append(handler.contributors, contributor)
		//}
		handler.errSetup = nil
	})
	return handler.errSetup
}

func (handler *contributorsHandler) getContributors() ([]clients.Contributor, error) {
	if err := handler.setup(); err != nil {
		return nil, fmt.Errorf("error during contributorsHandler.setup: %w", err)
	}
	return handler.contributors, nil
}
