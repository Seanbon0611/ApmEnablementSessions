#check if user has minkube based on the output of the minikube command
#if you have minikube installed but aren't using it, comment out the below if statement.
minkube_output=$(minikube)
SUB='command not found'
if [ "$minikube_output" != *"$SUB"* ]
then
  echo "Minikube detected, will create docker images for minikube"
#   eval $(minikube -p minikube docker-env) #change test1 to minikube
fi

#Deploy GinApp2
kubectl apply -f gin_app2_deployment.yaml

#Deploy GinApp1
kubectl apply -f gin_app1_deployment.yaml

echo "Sleeping for 1 minute before pulling service IP"
sleep 60


#Retrieve IP from GinApp2
gin_app2_ip=$(kubectl get pods -l app=ginapp2 -o=jsonpath="{range .items[*]}{.status.podIP}{end}")

echo "IP address for GinApp2:"
echo $gin_app2_ip

echo "Export IP as env var: KUBERNETES_GINAPP_IP"
export KUBERNETES_GINAPP_IP=$gin_app2_ip
echo "Exporting done!"

printenv | grep KUBERNETES_GINAPP_IP

#Retreive pod name
pod_name=$(kubectl get pod -l app=ginapp1 -o jsonpath="{.items[0].metadata.name}")
# pod_name=$(kubectl get pods -l app=ginapp1 -o custom-columns=:metadata.name)
echo "pod name is:"
echo $pod_name

#curl endpoints
echo "hitting endpoints of GinApp1"
echo "kubectl exec -it $pod_name -- curl localhost:8080/test"

kubectl exec -it $pod_name -- curl localhost:8080/test

