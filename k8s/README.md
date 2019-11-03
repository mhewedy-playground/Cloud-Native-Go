First Create docker image and push to docker hub   
```bash
cd ../ # go to the parent directory where Dockerfile lives
docker build -t cloud-native-go:1.0.0 .
docker tag cloud-native-go:1.0.0 mhewedy/cloud-native-go:1.0.0
docker push
```

1. go to kubesail.io and create an account   
then copy their connection settings into ~/.kube/config 
2. create deployment:  
`kubectl create -f deployment.yaml`
3. create service for the deployment:   
`kubectl create -f service.yaml`
4. create ingress for the service:   
`kubectl create -f ingress`
5. Access the service using:   
http://cloud-native-go.mhewedy.usw1.kubesail.io/api/books 
(or whatever url in HOSTS as result of `kubectl get ingress cloud-native-go`)
