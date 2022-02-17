# client-go-demo
Simple go client to demonstrate how to connect to k8s cluster (local) and list down some of the k8s resources like Pods and deployments

env GOOS=linux GOARCH=amd64 go build

docker build -t harishkadamudi/client-go:v0.0.2 -f Dockerfile .
docker tag harishkadamudi/client-go:v0.0.2 harishkadamudi/client-go:v0.0.2
docker push  harishkadamudi/client-go:v0.0.2
kubectl create role podepllist --resource pods,deployments --verb list
kubectl create rolebinding podepllist --role podepllist --serviceaccount default:default

If you're using kind then - https://kind.sigs.k8s.io/docs/user/quick-start/#loading-an-image-into-your-cluster

enjoy reading - https://dev.to/tidalmigrations/how-to-cross-compile-go-app-for-apple-silicon-m1-27l6
