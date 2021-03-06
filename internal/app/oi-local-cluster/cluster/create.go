/*
Copyright 2020 Rafael Fernández López <ereslibre@ereslibre.es>

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

package cluster

import (
	"fmt"

	localcluster "github.com/oneinfra/oneinfra/internal/pkg/local-cluster"
)

// Create creates a cluster with name clusterName, with private size
// privateClusterSize and public size publicClusterSize
func Create(clusterName, kubernetesVersion string, privateClusterSize, publicClusterSize int, remote bool) error {
	cluster := localcluster.NewHypervisorCluster(clusterName, kubernetesVersion, publicClusterSize, privateClusterSize, remote)
	if err := cluster.Create(); err != nil {
		return err
	}
	if err := cluster.Wait(); err != nil {
		return err
	}
	if remote {
		if err := cluster.StartRemoteCRIEndpoints(); err != nil {
			return err
		}
	}
	fmt.Print(cluster.Specs())
	return nil
}
