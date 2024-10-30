### README:

# Assignment 03 - Go Application with Docker

Ensure you have a running Kubernetes cluster and `kubectl` configured.
Metrics Server Configuration

If using Docker Desktop, apply the included components.yaml in k8s-deployment to configure the metrics-server with --kubelet-insecure-tls. This is required for the Horizontal Pod Autoscaler (HPA) to function:

### Apply Kubernetes Configuration:

1. Apply the namespace, ConfigMap, Secret, Service, StatefulSet and HPA from the k8's-deployment folder.
    ```bash
    kubectl apply -f k8s-deployment/namespace.yaml
    kubectl apply -f k8s-deployment/configmap.yaml
    kubectl apply -f k8s-deployment/secret.yaml
    kubectl apply -f k8s-deployment/service.yaml
    kubectl apply -f k8s-deployment/statefulset.yaml
    kubectl apply -f k8s-deployment/hpa.yaml
    ```

2. Check that all resources are created and running:
    ```bash
    Pods:
    kubectl get pods -n apeterson30

    Services:
    kubectl get svc -n apeterson30

    Persistent Volume Claims:
    kubectl get pvc -n apeterson30

    Horizontal Pod Autoscaler:
    kubectl get hpa -n apeterson30
    ```
    B. Readiness Probe Resources

    Apply the Probe Namespace, Deployment, and Service:
    ```bash
    kubectl apply -f readiness-probe/probe-namespace.yaml
    kubectl apply -f readiness-probe/probe-deployment.yaml
    kubectl apply -f readiness-probe/probe-service.yaml
    ```
    Verify Probe Resources:
    # Pods in Probe Namespace
    ```kubectl get pods -n probe```

    # Services in Probe Namespace
    ```kubectl get svc -n probe```

3. Key Features and Endpoints
    Data Persistence:
        Save data with POST /saveString
        Retrieve saved data with GET /getString

    CPU Load Simulation:
        Trigger CPU-intensive work with GET /busywait to test HPA scaling.

    Health Checks:
        /isAlive readiness probe ensures only healthy pods serve requests.

4. Testing Steps

    Test Data Persistence:
        Save a string to the volume:
        ```curl -X POST -H "Content-Type: application/json" -d '{"data": "testPersistence"}' localhost:8080/saveString```

    Verify retrieval:
        ```curl localhost:8080/getString```

    Restart the pod and confirm persistence:
    ```kubectl delete pod docker-gs-ping-statefulset-0 -n apeterson30```
    ```curl localhost:8080/getString  # Run after the pod is Ready```

    Verify Autoscaling:
    Simulate high CPU usage:
    ```curl localhost:8080/busywait```

    Monitor scaling:
    ```kubectl get hpa -n apeterson30 -w```

    Check Pod Readiness and Troubleshoot:
    After applying manifests, check pod readiness:
    ```kubectl describe pod docker-gs-ping-statefulset-0 -n apeterson30```

5. Accessing the /isAlive Endpoint

Once the Kubernetes Service is running, you can access the /isAlive endpoint to check the health of the docker-gs-ping application. This endpoint is exposed on port 30010 through the Service.
```
kubectl get service docker-gs-ping-service -n apeterson30
```
Note the EXTERNAL-IP of the docker-gs-ping-service.
```curl http://<EXTERNAL-IP>:30010/isAlive```
