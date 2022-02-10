// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2022 Satyam Bhardwaj <sabhardw@redhat.com>
// SPDX-FileCopyrightText: 2022 Utkarsh Chaurasia <uchauras@redhat.com>
// SPDX-FileCopyrightText: 2022 Avinal Kumar <avinkuma@redhat.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//    http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
Minimal Tekton Server is a small application written in golang which listens for
'mks' custom resource being created and send requests to Tekton APIs to create
corresponding resource on a Kubernetes/OpenShift Cluster. It only takes few
required fields into consideration and provides options to customise them.

It runs as a server and currently three types of resources are supported.

	- MksTask (Equivalent to Tekton 'Task')
	- MksTaskRun (Equivalent to Tekton 'TaskRun')
	- MksPipelineRun (Equivalent to Tekton 'PipelineRun')

Usage:

	./mks-server [flags]

The flags are:

	-kubeconfig
		Kubernetes configuration file. Automatically selects configuration from
		the cluster and user's home($HOME/.kube/config).
	-master
		The address of the kubernetes API server. Overrides any value in the
		kubeconfig.
	-addr
		The address of the Redis server. Defaults to 127.0.0.1:6379.
	-password
		The password of the Redis database.

This minimal server also serves the redis database. Database is used to store
the count the mks resources being created, deleted, is active, completed or
failed. The choice of the database is Redis due to its simplicity and easy
integration (see https://github.com/go-redis/redis).

This server can receive request via exposed API or by using Minimal Tekton
Server CLI application aka 'mks-cli' (see https://github.com/MiniTeks/mks-cli).
You can create, delete, list or update a particular resource and the changes are
transferred to the Tekton API and applied. See more about mks-cli in its
respective repository.

The live statistics of the resources can be seen using Minimal Tekton Server
Dashboard aka 'mks-ui' (see https://github.com/MiniTeks/mks-ui). The dashboard
is very minimal and shows all the stas in a tabular form. See more about mks-ui
in its respective repository.

*/

package main
