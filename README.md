# rest_microservices_comms_golang

## Values to change in the helm chart + quick tutorial on deploying a microservice via helm.

This application (look at commit history to see what changes had to be made) creates a very simple rest api in golang.  
Here are the steps in order:  
[Full Tutorial](https://www.techtarget.com/searchitoperations/tutorial/Build-and-deploy-a-microservice-with-Kubernetes)

1. Create a project directory and run `go mod init example.com/microsvc`
2. Create the `main.go` script which takes care of http requests.
3. Add a `.DOCKERFILE` in which you create the script's artifact. You will also need to set the following:
    ```dockerfile
    ENV PORT 4317
    EXPOSE 4317
    ```
   Set the port to whatever you would like. Then build and push the docker image.
4. Navigate to the root project directory and run the following command:
    ```shell
    helm create microsvc
    ```
   This will create a helm chart for easy deployment/maintenance on Kubernetes.
5. Within the helm chart, navigate to following files, and make the following changes:
    * `values.yaml`
        * Change the image's repository name to the name and tag to the one you pushed to.
        * Update the `fullnameOverride` to whatever you wish.
        * Update the `service.port` to the port you will use.
        * Set `ingress.enabled` to true.
    * `templates/deployment.yaml`
        * Update `containers.ports.containerPort` to the same port used throughout.
        * Comment out the `containers.livenessProbe` and `containers.readinessProbe`.
        * Add a `containers.env` port, right after resources:
            ```yaml
          resources:
            {{- toYaml .Values.resources | nindent 12 }}
          env:
            - name: PORT
              value: "4317"
            ```
    * `templates/ingress.yaml`
        * Update `metadata.labels`:
          ```yaml
          labels:
            tier: backend
          ```
    * `templates/service.yaml`
        * Update `spec.externalName` to be:
          ```yaml
          spec:
            type: {{ .Values.service.type }}
            externalName: {{ include "microsvc.fullname" . }}.default.svc.cluster.local
          ```
        * Update `targetPort` to match the port used throughout.
6. Your microservice is now ready to deploy. Use helm install to deploy (command mentioned above).
7. Test your microservice on kubernetes by shelling into another cluster. Then issue the following command:

```shell
curl microsvc.default.svc.cluster.local:4317/employee
```

## Testing The External IP
This service is available as an external IP. You can test it from your local machine by just doing:
```shell
curl http://ac185762-4cc1-41f6-8fac-09b8f1dd97d3.k8s.civo.com/service-2-server
```
**NOTE:** this address (the first `.com` part) is different for every cluster you want to run this service on. It needs to
configured manually. Or at least until I find a way to do it automatically.
