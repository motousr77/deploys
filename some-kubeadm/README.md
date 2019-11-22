# Lets try
<!-- install kubeadm -->
sudo apt-get update && sudo apt-get install -y apt-transport-https curl
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -

cat <<EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list
deb https://apt.kubernetes.io/ kubernetes-xenial main
EOF

sudo apt-get update && sudo apt-get install -y kubelet kubeadm kubectl && sudo apt-mark hold kubelet kubeadm kubectl
<!--  -->
When using Docker, kubeadm will automatically detect the cgroup driver for the kubelet and set it in the /var/lib/kubelet/kubeadm-flags.env file during runtime.

If you are using a different CRI, you have to modify the file /etc/default/kubelet (/etc/sysconfig/kubelet for CentOS, RHEL, Fedora) with your cgroup-driver value, like so:

KUBELET_EXTRA_ARGS=--cgroup-driver=<value>
<!-- downgrade docker -->
sudo apt-get autoremove -y docker-engine && sudo apt-get purge docker-engine -y && sudo rm -rf /etc/docker/ && sudo rm -f /etc/systemd/system/multi-user.target.wants/docker.service && sudo rm -rf /var/lib/docker &&  sudo systemctl daemon-reload
<!--  -->
sudo apt install -y docker-ce=$(apt-cache madison docker-ce | grep 18.06 | head -1 | awk '{print $3}')

sudo kubeadm init --apiserver-advertise-address $(hostname -i)

sudo kubeadm init --pod-network-cidr=10.244.0.0/16
10.244.0.0/16
sudo kubeadm init --apiserver-advertise-address 192.168.6.11 --pod-network-cidr=10.244.0.0/16

To start using your cluster, you need to run the following as a regular user:

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  https://kubernetes.io/docs/concepts/cluster-administration/addons/
kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/k8s-manifests/kube-flannel-rbac.yml
kubectl apply -f https://docs.projectcalico.org/v3.10/manifests/calico.yaml

Then you can join any number of worker nodes by running the following on each as root:

sudo kubeadm join 192.168.6.11:6443 --token c04la0.s5vgijnart3d18ut --discovery-token-ca-cert-hash sha256:6a165fc8a4671241a15641b05098298ada62c92156cdf8e2efa0839efdd4cd11

<!--  -->
mkdir -p $HOME/.kube && sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config && sudo chown $(id -u):$(id -g) $HOME/.kube/config
<!--  -->
systemctl daemon-reload && systemctl restart docker
<!--  -->
kubeadm join --token SOMETOKEN SOMEIPADDRESS --discovery-token-ca-cert-hash SOMESHAHASH
<!--  -->
kubectl apply -n kube-system -f "https://cloud.weave.works/k8s/net?k8s-version=$(kubectl version | base64 |tr -d '\n')"
<!--  -->

ping -c 2 192.168.6.11 && ping -c 2 192.168.6.21 && ping -c 2 192.168.6.22

kubectl run nginx --image=nginx:1.17.6 --replicas=4

Flannel: "Network": "10.244.0.0/16