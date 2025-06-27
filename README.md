### README:

# Kubernetes Capstone - Go Application with HPA, Volumes, and Health Checks

This project is a containerized Go application designed for deployment in a Kubernetes cluster using a **StatefulSet**, **Persistent Volume**, **Horizontal Pod Autoscaler**, and a **readiness probe**. It exposes multiple endpoints for testing application behavior, persistence, and scaling.

---

## 🔧 Features

- `/saveString`, `/getString` – Save and retrieve persistent data
- `/busywait` – Triggers CPU-intensive task to test autoscaling
- `/isAlive` – Readiness probe endpoint
- `/hello`, `/foo`, `/kill` – Basic route tests
- Environment-aware using ConfigMap, Secret, and Env Vars

---

## Docker Image

The application is already containerized and published.

```
docker pull apeterson30/assignment03:latest
```

## How to Deploy (From Scratch)

Make sure kubectl is pointed to a live Kubernetes cluster. Docker Desktop with Kubernetes enabled works fine for local testing.
1. Clone the Repo
```
git clone https://github.com/Ajosephp/Kubernetes-Capstone.git
cd Kubernetes-Capstone
```
2. Apply Core Kubernetes Resources
```
kubectl apply -f k8s-deployment/namespace.yaml
kubectl apply -f k8s-deployment/configmap.yaml
kubectl apply -f k8s-deployment/secret.yaml
kubectl apply -f k8s-deployment/service.yaml
kubectl apply -f k8s-deployment/statefulset.yaml
kubectl apply -f k8s-deployment/hpa.yaml
```
3. Verify the Deployment
Check that everything is running:
```
kubectl get pods -n apeterson30
kubectl get svc -n apeterson30
kubectl get pvc -n apeterson30
kubectl get hpa -n apeterson30
```
Test Application
Test Basic Route
```
curl http://localhost:8080/
```
Save & Retrieve Data (Persistent Volume Test)
```
curl -X POST -H "Content-Type: application/json" -d '{"data": "Hello from Andrew"}' http://localhost:8080/saveString
curl http://localhost:8080/getString
```
Force CPU Load (Trigger Autoscaling)
```
curl http://localhost:8080/busywait
```
Then monitor the HPA reacting:
```
kubectl get hpa -n apeterson30 -w
```
Readiness Probe
The pod only becomes “Ready” once this health check passes:
```
curl http://localhost:30010/isAlive
```
You should see:
```
true
```
Tear Down
```
kubectl delete namespace apeterson30
```
Tech Stack
```
    Go (net/http)
    Docker
    Kubernetes:
        StatefulSet
        PersistentVolumeClaim
        ConfigMap / Secret
        LoadBalancer Service
        Horizontal Pod Autoscaler
        Readiness Probes
```
