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

package tconfig

import (
	tektonclient "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	versionedResource "github.com/tektoncd/pipeline/pkg/client/resource/clientset/versioned"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// TektonParam struct contains Tekton client and kubernetes configuration. 
type TektonParam struct {
	clients        *Client
	kubeConfigPath string
	kubeContext    string
	namespace      string
}

var _ Params = (*TektonParam)(nil)

func (p *TektonParam) SetKubeConfigPath(path string) {
	p.kubeConfigPath = path
}

func (p *TektonParam) SetKubeContext(context string) {
	p.kubeContext = context
}

func (p *TektonParam) tektonClient(config *rest.Config) (tektonclient.Interface, error) {
	cfg, err := tektonclient.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func (p *TektonParam) resourceClient(config *rest.Config) (versionedResource.Interface, error) {
	cfg, err := versionedResource.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func (p *TektonParam) kubeClient(config *rest.Config) (kubernetes.Interface, error) {
	cfg, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func (p *TektonParam) dynamicClient(config *rest.Config) (dynamic.Interface, error) {
	cfg, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func (p *TektonParam) Client() (*Client, error) {
	if p.clients != nil {
		return p.clients, nil
	}

	config, err := p.config()
	if err != nil {
		return nil, err
	}

	tekton, err := p.tektonClient(config)
	if err != nil {
		return nil, err
	}

	resource, err := p.resourceClient(config)
	if err != nil {
		return nil, err
	}

	kube, err := p.kubeClient(config)
	if err != nil {
		return nil, err
	}

	dynamic, err := p.dynamicClient(config)
	if err != nil {
		return nil, err
	}

	p.clients = &Client{
		Tekton:   tekton,
		Kube:     kube,
		Resource: resource,
		Dynamic:  dynamic,
	}

	return p.clients, nil

}

func (p *TektonParam) config() (*rest.Config, error) {
	lr := clientcmd.NewDefaultClientConfigLoadingRules()
	if p.kubeConfigPath != "" {
		lr.ExplicitPath = p.kubeConfigPath
	}

	coverride := &clientcmd.ConfigOverrides{}

	if p.kubeContext != "" {
		coverride.CurrentContext = p.kubeContext
	}

	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(lr, coverride)

	if p.namespace == "" {
		namespace, _, err := kubeconfig.Namespace()
		if err != nil {
			return nil, err
		}
		p.namespace = namespace
	}

	cfg, err := kubeconfig.ClientConfig()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func (p *TektonParam) SetNamespace(ns string) {
	p.namespace = ns
}

func (p *TektonParam) Namespace() string {
	return p.namespace
}
