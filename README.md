# workshop-go

## Useful cammand lines

```bash
# Delete all docker containers and images
docker rm $(docker ps -a -q) -f ; docker rmi $(docker images -q) -f

# Create a local namespace
kubectl create namespace local

# Setup kubernetes config
kubectl config set-context docker-desktop --namespace=local
kubectl config use-context docker-desktop --namespace=local

# Get kubernetes pods
kubectl get pods
```
