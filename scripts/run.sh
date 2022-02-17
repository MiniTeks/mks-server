#!/bin/sh
echo "Creating a new project"
oc new-project mks
sleep 2

echo
echo "Applying mks crd files on the cluster"
kubectl apply -f ./config/mksCRDs
sleep 1

echo "Add Openshift Pipeline Operator from OC UI"
# sleep 10

# echo
# echo "Installing tekton pipelines operator in the cluster"
# oc apply --filename https://storage.googleapis.com/tekton-releases/pipeline/latest/release.notags.yaml
# oc apply --filename https://storage.googleapis.com/tekton-releases/dashboard/latest/tekton-dashboard-release.yaml
# sleep 2

echo
echo "Installing user defined tasks"
kubectl apply -f ./tekton/tasks
sleep 1

echo
echo "Installing required tasks from hub.tekton.dev"
kubectl apply -f https://raw.githubusercontent.com/tektoncd/catalog/main/task/buildah/0.3/buildah.yaml
kubectl apply -f https://raw.githubusercontent.com/tektoncd/catalog/main/task/git-clone/0.5/git-clone.yaml
kubectl apply -f https://raw.githubusercontent.com/tektoncd/catalog/main/task/golang-test/0.2/golang-test.yaml
# kubectl apply -f https://raw.githubusercontent.com/tektoncd/catalog/main/task/kaniko/0.5/kaniko.yaml
sleep 1

echo
echo "Installing user defined pipelines "
kubectl apply -f ./tekton/pipelines
sleep 2

echo
echo "Installing user defined pvcs "
kubectl apply -f ./tekton/pvcs
sleep 2
# echo
# echo "to install the latest release of Tekton Triggers and its dependencies"
# kubectl apply --filename https://storage.googleapis.com/tekton-releases/triggers/latest/release.yaml
# kubectl apply --filename https://storage.googleapis.com/tekton-releases/triggers/latest/interceptors.yaml
# sleep 2

echo
echo  "Get db server up and running."
oc apply -f ./k8s/db-config
sleep 2

echo
echo "Get mks-server up and running."
oc apply -f ./k8s/mks-config
sleep 2

echo
echo "Apply the files to create the mks resource examples"
oc apply -f ./config/mksResourceExample
sleep 2
