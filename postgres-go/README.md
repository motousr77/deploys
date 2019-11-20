kubectl apply -f pg-cfg.yaml
kubectl apply -f pg-deployment.yaml

kubectl get nodes
kubectl get pods

<!-- get podname and save to variable pod $ -->
export pod=$(kubectl get pods -l app=postgres --no-headers -o custom-columns=":metadata.name")
<!-- go to inside pod $ -->
kubectl exec -it $pod sh
<!-- inside pod create db and table -->
createdb -U postgres hello
psql -U postgres
\c hello 
CREATE TABLE hits (total int);
INSERT INTO hits (total) VALUES (0);
\q
exit
<!-- # -->
kubectl describe pods -l app=postgres
kubectl apply -f pg-svc.yaml
kubectl get all
<!-- build the docker image -->
docker build -t hello .
<!-- if we hame error: unknown flag from
do this:
Remove --chown from the COPY line in the Dockerfile and use it under RUN flag
 -->
kubectl get pods -n kube-system
<!-- # -->
kubectl apply -f deployment.yaml
kubelet get pods
kubectl apply -f service.yaml
kubectl get services
<!--  -->
read -p "Enter the cluster IP address for hello-svc: " clusterip
<!--  -->
curl $clusterip:30000
<!--  -->
kubectl scale deployment _deployment_name_ --replicas=_replica_count
<!-- end. -->