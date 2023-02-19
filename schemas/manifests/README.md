# Manifests

## Structure of this `manifests` directory
- `argocd` folder contains the argocd application
- `k8s` folder contains the kubernetes yaml files of our `gocom` project

## Sequence of commands to apply the manifests
- from `manifests` dir
- `kubectl apply -f k8s/postgresql-db.yaml`
    - apply the DB manifests and wait until DB becomes ready
    - confirm by `kubectl get pods`
- `kubectl apply -f k8s/auth.yaml`
- `kubectl apply -f k8s/ecom.yaml`

## Steps to trigger the GitOps's CD part by ArgoCD

### For the First time (this is needed only once initially)
- create `argocd` namespace: `kubectl create namespace argocd`
- install `argocd`: `kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml`:
    - make sure pods are running: `kubectl get pods -n argocd`
    - check the services: `kubectl get svc -n argocd`
- port forward of the `argocd` server (to access the `argocd` UI): `kubectl port-forward -n argocd svc/argocd-server 8080:443`
- login to argocd UI(`localhost:8080`):
    - username: `admin`
    - password: get by `kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo`
- connect the git repository from argocd UI:
    - go to repositories -> add github
    - you may need to provide git personal access token for the respective git repo
- apply argocd application: `kubectl apply -f argocd/application.yaml` 


### From next time
- open a PR for changing the manifests(basically update the `gocom` project's services's docker image in respective manifests from `k8s` directory)
- when the PR get merged argocd will trigger the changes in the respective cluster
> Note: argocd syncs in every 3minutes
