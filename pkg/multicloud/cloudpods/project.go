// Copyright 2019 Yunion
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

package cloudpods

import (
	"yunion.io/x/onecloud/pkg/apis/compute"
	api "yunion.io/x/onecloud/pkg/apis/identity"
	"yunion.io/x/onecloud/pkg/cloudprovider"
	"yunion.io/x/onecloud/pkg/mcclient/modules"
	"yunion.io/x/onecloud/pkg/multicloud"
)

type SProject struct {
	multicloud.SResourceBase
	multicloud.CloudpodsTags

	cli *SCloudpodsClient

	api.ProjectDetails
}

func (self *SProject) GetName() string {
	return self.Name
}

func (self *SProject) GetId() string {
	return self.Id
}

func (self *SProject) GetGlobalId() string {
	return self.Id
}

func (self *SProject) GetStatus() string {
	return compute.EXTERNAL_PROJECT_STATUS_AVAILABLE
}

func (self *SCloudpodsClient) GetProjects() ([]SProject, error) {
	ret := []SProject{}
	return ret, self.list(&modules.Projects, nil, &ret)
}

func (self *SCloudpodsClient) GetIProjects() ([]cloudprovider.ICloudProject, error) {
	projects, err := self.GetProjects()
	if err != nil {
		return nil, err
	}
	ret := []cloudprovider.ICloudProject{}
	for i := range projects {
		projects[i].cli = self
		ret = append(ret, &projects[i])
	}
	return ret, nil
}
