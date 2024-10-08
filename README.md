### README:

# Assignment 02 - Go Application with Docker

Ensure you have a running Kubernetes cluster and `kubectl` configured.

### Apply Kubernetes Configuration:

1. Apply the namespace, ConfigMap, Secret, Deployment, and Service:
    ```bash
    kubectl apply -f ubernetes-deployment.yaml
    ```

    *Note: The namespace will be created automatically as defined in `kubernetes-deployment.yaml`.*

2. Verify the deployment:
    ```bash
    kubectl get all -n apeterson30
    ```

    *This command will list all resources in the `apeterson30` namespace.*

3. Access the application at [http://localhost:30000](http://localhost:30000).

## Testing the Endpoints

1. **GET /foo**
    ```bash
    curl http://localhost:30000/foo
    ```
    **Response:**
    ```
    bar
    ```

2. **POST /hello**
    ```bash
    curl -H "Accept: application/json" -H "Content-Type: application/json" -X POST --data '{"name": "YourName"}' http://localhost:30000/hello
    ```
    **Response:**
    ```
    Hello YourName!
    ```

3. **GET /kill**
    ```bash
    curl http://localhost:30000/kill
    ```
    **Response:**
    ```
    Shutting down...
    ```
    *(The server will shut down after this request.)*

4. **GET /configValue**
    ```bash
    curl http://localhost:30000/configValue
    ```
    **Response:**
    ```
    snake
    ```

5. **GET /secretValue**
    ```bash
    curl http://localhost:30000/secretValue
    ```
    **Response:**
    ```
    secretSnake
    ```

6. **GET /envValue**
    ```bash
    curl http://localhost:30000/envValue
    ```
    **Response:**
    ```
    environmentSnake
    ```