# Troubleshooting APM in Kubernetes Environments

## Prerequisites
---

In order to run this sandbox, you will need Minikube.

### If you do not have Minikube:
1. Download and install Docker to use as the VM Driver for Minikube
Docker: Install Docker Desktop on Mac

2. Install the kubectl client to interact with your Kubernetes cluster

```
brew install kubernetes-cli
```
3. Install Minikube (minikube start)
```
brew install minikube
```
4. Now that Minikube has been installed, you can start it with the default parameters
```
minikube start
```
<br/> 

### Deploy the Agent on your Minikube environment via Helm:

1. Install Helm by running the following command:

```
brew install helm
```

2. For the first install, you will have to add the Datadog repository and the Helm stable (for KSM Legacy) repository with the following commands:

```
helm repo add datadog https://helm.datadoghq.com
```
You can then fetch the latest charts from this repos by using the command:

```
helm repo update
```

Set your API key: 
```
kubectl create secret generic datadog-agent --from-literal='api-key=<API_KEY>'
```

Then install and upgrade everything with the command:
```
helm install datadog -f values.yaml datadog/datadog
helm upgrade datadog -f values.yaml datadog/datadog
```

<br/> 
<br/> 
<br/> 

## Getting Started
---
1. Ensure theat we are working in the kubernetes_exercise directory and spin up the application using the command:
```
kubectl apply -f application_deployment.yaml
```

Lets verify our pod is running with kubectl get pods

copy the pod name and lets make a few requests for this app:
```
kubectl exec -it <pod_name> -- bash
```
We can see that our requests are going through, so lets check stdout for our application
```
kubectl logs <pod_name>
```

Where could we go from here?

