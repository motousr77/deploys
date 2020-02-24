### Create 'Hello'
~~~bash
kubectl apply \
  -f pg-cfg.yaml \
  -f pg-deployment.yaml \
  -f pg-svc 

kubectl get nodes,pods

# get podname and save to variable pod $
export pod=$(kubectl get pods -l app=postgres --no-headers -o custom-columns=":metadata.name")

# go to inside pod $
kubectl exec -it $pod sh
~~~
~~~SQL
# inside pod create db and table
createdb -U postgres hello
psql -U postgres
\c hello 
CREATE TABLE hits (total int);
INSERT INTO hits (total) VALUES (0);
\q
exit
~~~

~~~bash
kubectl describe pods -l app=postgres
kubectl apply -f pg-svc.yaml
kubectl get all

build the docker image
docker build -t hello .

if we hame error: unknown flag from
do this:
Remove --chown from the COPY line in the Dockerfile and use it under RUN flag

kubectl get pods -n kube-system

kubectl apply -f deployment.yaml
kubelet get pods
kubectl apply -f service.yaml
kubectl get services

read -p "Enter the cluster IP address for hello-svc: " clusterip

curl $clusterip:30000

kubectl scale deployment _deployment_name_ --replicas=_replica_count

# copy to here privat_key from vagrant directories:
mkdir ~/machine-ssh/worker-1/private_key
cat > ~/machine-ssh/worker-1/private_key << EOF
EOF
chomd 600 ~/machine-ssh/worker-1/private_key

mkdir ~/machine-ssh/worker-2/private_key
cat > ~/machine-ssh/worker-1/private_key << EOF
EOF
chmod 600 ~/machine-ssh/worker-2/private_key

scp -i '~/machine-ssh/worker-1/private_key':/home/vagrant/readme_copy /home/vagrant/readme

# You will need to save the Docker image as a tar file:

docker save -o <path for generated tar file> <image name>
# Then copy your image to a new system with regular file transfer tools such as cp, scp or rsync(preferred for big files). After that you will have to load the image into Docker:

docker load -i <path to image tar file>
# PS: You may need to sudo all commands.

# EDIT: You should add filename (not just directory) with -o, for example:

docker save -o c:/myfile.tar centos:16

docker tag local-image:tagname new-repo:tagname
docker push new-repo:tagname
docker tag hello:latest devcodemy/hellops:latest
docker push hello:latest
