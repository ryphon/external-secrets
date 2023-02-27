/*
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

package v1beta1

import smmeta "github.com/external-secrets/external-secrets/apis/meta/v1"

// KeeperSecurityProvider Configures a store to sync secrets using Keeper Security.
type KeeperSecurityProvider struct {
	Auth     smmeta.SecretKeySelector `json:"authRef"`
	Hostname string                   `json:"hostname"`
	FolderID string                   `json:"folderID"`
}
