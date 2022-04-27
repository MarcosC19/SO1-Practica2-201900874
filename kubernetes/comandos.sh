# CREACION DE KLUSTER K8S
gcloud container clusters create practica2 --num-nodes=1 --tags=allin,allout --machine-type=n1-standard-2 --no-enable-network-policy

# INSTALACION RABBITMQ OPERATOR
kubectl apply -f https://github.com/rabbitmq/cluster-operator/releases/latest/download/cluster-operator.yml

# CREACION NAMESPACE
kubectl create namespace practica2-201900874 --dry-run -o yaml > config201900874.yaml

# CREACION DEPLOYMENT
kubectl create deploy grpc --namespace practica2-201900874 --image=curtex19/client_grpc_201900874 --image=curtex19/server_grpc_201900874 --replica=1 --dry-run -o yaml >> config201900874.yaml

# CREACION RABBIT SUBSCRIBER
kubectl run subscriber --image=curtex19/rabbit_subscriber_201900874 --restart=Never --namespace=practica2-201900874 --dry-run -o yaml >> config201900874.yaml

# CREACION LOADBALANCER
kubectl expose deploy/grpc --type=LoadBalancer --port=8080 --namespace=practica2-201900874 --dry-run=client -o yaml >> config201900874.yaml
