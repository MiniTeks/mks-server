# Minimal Tekton Server

It is an API server which exposes a few APIs that users can use to create the resource which will be watched by the controller to further create the Tekton resources.

---
## Tekton Resources Can Be Created

- Task
- TaskRun
- PipelineRun

---

## How To Run ?

- Clone the repository using:
```bash
git clone git@github.com:albertplaya/tinder-clone.git
```
- Run the below commands:
```bash
go mod tidy

go mod vendor
```
- Build the project using:
```bash
go build -o mks-server
```
- Make sure minikube cluster is running and tekton pipeline installed
- Apply Custom Resource Definition present in config/ according to custom resource you want create
```bash
kubectl apply -f config/<crd-file-name>.yaml
```
- Run the executable using:
```bash
./mks-server -kubeconfig=$HOME/.kube/config
```
- For checking create custom resource by applying custom-resource file
```bash
kubectl apply -f config/<cr-example>.yaml
```


