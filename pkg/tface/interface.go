package tface

import (
	tektoncd "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	versionedResource "github.com/tektoncd/pipeline/pkg/client/resource/clientset/versioned"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

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

func (p *TektonParam) tektonClient(config *rest.Config) (tektoncd.Interface, error) {
	cfg, err := tektoncd.NewForConfig(config)
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
