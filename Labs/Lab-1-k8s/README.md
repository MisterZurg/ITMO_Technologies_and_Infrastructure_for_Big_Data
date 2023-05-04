# Lab 1 — Kubernetes

> You could follow the guide to understand what you're doing
> 
> OR
> 
> Go to [Short what to do](#short-what-to-do)

## Connect to Cluster
<details>
  <summary>ℹ️ Collapse</summary>

1. Start and connect VPN
2. Check the connectivity
```shell
(base) misterzurg@MacBook-Pro-Denis ~ % ping 10.32.7.101
PING 10.32.7.101 (10.32.7.101): 56 data bytes
64 bytes from 10.32.7.101: icmp_seq=0 ttl=61 time=30.223 ms
```
3. Connect to Cluster
```shell
(base) misterzurg@MacBook-Pro-Denis ~ % ssh <your_login>@gateway.st
password >>> <your_password>
```

4. Get the external port
```shell
[<your_login>@gateway ~]$ kubectl get svc
NAME                TYPE       CLUSTER-IP      EXTERNAL-IP   PORT(S)                         AGE
jupyter-spark-svc   NodePort   10.129.154.22   <none>        8888:32074/TCP,4040:31385/TCP   5m25s
```

5. Type in browsers URL
```
node03.st:32074
```

6. Get token
```shell
[<your_login>@gateway ~]$ kubectl logs jupyter-spark-7c5b4455cc-fv7dr 
```
</details>

## 1. (1 point) Create Jupyter’s Deployment and Service.
Requirements:
<details>
  <summary>а. Guaranteed QoS. Resources amount - 1 CPU, 2 Gb RAM</summary>
  
  ```yaml
  # Jupyter’s Deployment
  # Guaranteed QoS - means that Requests and Limits are the same 
  # OR No Requests, only Limits
  resources:
    requests:
      memory: "2Gi"
      cpu: "1"
    limits:
      memory: "2Gi"
      cpu: "1"
  ```
</details>

- b. Replicas number - 1.
- c. Pod has to contain the following label - `jupyter: lab0`
- d. Deployment name: `lab0-jupyter`
- e. Service name: `lab0-jupyter-service`
- f. Service has to forward requests only to this deployment’s pod. Ports number for traffic forwarding - 1.
- g. Service type - `NodePort`

> Your Jupyter image is 
> ```
> node03.st:5000/pyspark-hdfs-jupyter:<your_login>
> ```

```shell
# Create file where you store the description of your Deployment
[<your_login>@gateway ~]$ vim lab0-jupyter-deployment.yaml
~                                                                               
~ 
:wq

# Create file where you store the description of your Service
[<your_login>@gateway ~]$ vim lab0-jupyter-service.yaml
~                                                                               
~ 
:wq
```
### Jupyter’s Service
```yaml
apiVersion: v1
kind: Service
metadata:
  name: lab0-jupyter-service
  
spec:
  type: NodePort  # Kubernetes control plane allocates a port from a range specified by --service-node-port-range flag (default: 30000-32767)
  selector:       # The set of Pods targeted by a Service
    jupyter: lab0
```

### Jupyter’s Deployment
```yaml
apiVersion: v1
kind: Deployment
metadata:
  name: lab0-jupyter

# Deployment settings
spec:
  replicas: 1
  selector: # defines how the created ReplicaSet finds which Pods to manage
    matchLabels:
      jupyter: lab0
  template:
    metadata:
      labels:
        jupyter: lab0
    # Pod settings that Deployment is going to create
    spec: 
      containers:
      - name: jupyter
        image: node03.st:5000/pyspark-hdfs-jupyter:<your_login> # !!! PUT YOUR LOGIN HERE !!!
        resources:
          requests:
            memory: "2Gi"
            cpu: "1"
          limits:
            memory: "2Gi"
            cpu: "1"
```

```shell
# Create a Deployment — that provides declarative updates for Pods and ReplicaSets.
kubectl create -f lab0-jupyter-deployment.yaml
# Create a Service — method for exposing a network application that is running as one or more Pods in your cluster.
kubectl create -f lab0-jupyter-service.yaml
```

## 2. (2 points) Modify previous Deployment and Service to face following conditions:
<details> 
  <summary>
  Jupyter has to listen on <code>8282</code> port (you can set it up using args).
  </summary>

  ```
  NotebookApp.port : Int
  Default: 8888
  The port the notebook server will listen on (env: JUPYTER_PORT).
  ```
</details>

- b. Jupyter’s token has to match your gateway.st password.
- c. Service has to receive requests on port `80` and forward them on `8282` port .
- d. Deployment and Service names have to stay the same!


### Jupyter’s Service
```yaml
apiVersion: v1
kind: Service
metadata:
  name: lab0-jupyter-service

spec:
  type: NodePort  # Kubernetes control plane allocates a port from a range specified by --service-node-port-range flag (default: 30000-32767)
  selector:
    jupyter: lab0
  ports:
    - protocol: TCP
      port: 80
      targetPort: 8282  
```


### Jupyter’s Deployment
```yaml
apiVersion: v1
kind: Deployment
metadata:
  name: lab0-jupyter

spec:
  replicas: 1
  selector:
    matchLabels:
      jupyter: lab0
  template:
    metadata:
      labels:
        jupyter: lab0
    spec:    
      ports:
         - containerPort: 8282
      containers:
      - name: jupyter
        image: node03.st:5000/pyspark-hdfs-jupyter:<your_login>
        args: ["--NotebookApp.token='<your_password>'", "--NotebookApp.port=8282"]
        resources:
          requests:
            memory: "2Gi"
            cpu: "1"
          limits:
            memory: "2Gi"
            cpu: "1"
```


## 3. (2 points) Create the following ConfigMap (CM):
- a. CM name: `lab0-jupyter-cm`
- b. CM has to contain a  file-like key. Key - `jupyter_notebook_config.py`
- c. Content for the `jupyter_notebook_config.py` key (2 rows): 
  - `с.NotebookApp.trust_xheaders = True`
  - `c.NotebookApp.quit_button = False`

### Jupyter’s ConfigMap
> `ConfigMap` — an object to store the data in key-value pairs.
```shell
# Create file where you store the description of your ConfigMap
[<your_login>@gateway ~]$ vim lab0-jupyter-cm.yaml
~                                                                               
~ 
:wq
```

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: lab0-jupyter-cm
  labels:
    jupyter: lab0
data:
  # file-like keys
  jupyter_notebook_config.py: |
    c.NotebookApp.trust_xheaders = True
    c.NotebookApp.quit_button = False
```

```shell
# Create a ConfigMap 
kubectl create -f lab0-jupyter-cm.yaml
```

## 4. (3 points) Mount jupyter_notebook_config.py (only this file) from your CM into the Jupyter pod:
- а. Mount has to be completed from the CM.
- b. Mount path - `/home/jovyan/.jupyter/jupyter_notebook_config.py`
- c. Mount has to be read only.

### Jupyter’s Deployment
```yaml
apiVersion: v1
kind: Deployment
metadata:
  name: lab0-jupyter

spec:
  replicas: 1
  selector:
    matchLabels:
      jupyter: lab0
  template:
    metadata:
      labels:
        jupyter: lab0
    spec:    
      containers:
      - name: jupyter
        image: node03.st:5000/pyspark-hdfs-jupyter:<your_login>
        ports:
        - containerPort: 8282
        command: ["start-notebook.sh"]  # Without CrashBackOff
        args: ["--NotebookApp.token='<your_password>'", "--NotebookApp.port=8282"]
        resources:
          requests:
            memory: "2Gi"
            cpu: "1"
          limits:
            memory: "2Gi"
            cpu: "1"
        # Mount the config map
        volumeMounts:
          - name: config-volume
            mountPath: /home/jovyan/.jupyter
            readOnly: true
      volumes:
        - name: config-volume
          configMap:
            # Provide the name of the ConfigMap containing the files you want
            # to add to the container
            name: lab0-jupyter-cm
            # An array of keys from the ConfigMap to create as files
            items:
              - key: "jupyter_notebook_config.py"
                path: "jupyter_notebook_config.py"
```

## 5. (1 point) Save the Jupyter’s pod log with timestamps into the file lab0-jupyter.log

> ```shell
> kubectl logs <POD_NAME> --timestamps >> lab0-jupyter.log
> ```


```shell
[<your_login>@gateway ~]$ kubectl get pods
NAME                             READY   STATUS    RESTARTS   AGE
jupyter-spark-7c5b4455cc-fv7dr   1/1     Running   0          5d20h
lab0-jupyter-79b957db4b-zsd4p    1/1     Running   0          6s

[<your_login>@gateway ~]$  kubectl logs lab0-jupyter-79b957db4b-zsd4p --timestamps >> lab0-jupyter.log
```

## 6. (1 point) Save in file lab0-ls-lah.sh the command to list files in the Jupyter container path /home/jovyan/.jupyter.


> ```shell
> # lab0-ls-lah.sh
> kubectl exec <POD_NAME> -- ls -a /home/jovyan/.jupyter
> ```

```shell
# lab0-ls-lah.sh
[<your_login>@gateway ~]$ kubectl exec lab0-jupyter-79b957db4b-zsd4p -- ls -a /home/jovyan/.jupyter
```

> **Attention all Fortnite gamers** 
> 
> pods, deployments, services, etc **HAVE TO BE PRESENTED** in your namespace.

So we have to add in each `*.yaml`
```yaml
apiVersion: v1
kind: <Deployment | Service | ConfigMap>
metadata:
  namespace: <your_login>
```

## Cheking everything is ok
1. Get your external port
```shell
[<your_login>@gateway ~]$ kubectl get service
NAME                   TYPE       CLUSTER-IP       EXTERNAL-IP   PORT(S)                         AGE
jupyter-spark-svc      NodePort   10.129.154.22    <none>        8888:32074/TCP,4040:31385/TCP   7d17h
lab0-jupyter-service   NodePort   10.129.187.168   <none>        80:30960/TCP                    42h
```
> `lab0-jupyter-service` ... 80:**30960**

2. Go to `http://node03.st:<external_port>`
> http://node03.st:**30960**

3. On login page as a token input `<your_password>` the one that you hard code!

4. There will be smth like
>```shell
> / 📁
>   | - 📁 nltk-data
>   | - 📁 work
> ```

## Short what to do
### Create 3 or 1 YAMLS's and 1 shell
```shell
[<your_login>@gateway ~]$ vim <filename>.yaml
```
### CTRL+C/V updated
> Update all
> `<put what you need here>` and `TODO HARDCODE` stuff

### Check that they appeared
```shell
[<your_login>@gateway ~]$ ls
<your_login>-jupyter.yaml  lab0-jupyter-cm.yaml  lab0-jupyter-deployment.yaml  lab0-jupyter-service.yaml  lab0-jupyter.log  lab0-ls-lah.sh
```
### Apply them
```shell
[<your_login>@gateway ~]$ kubectl apply -f <filename>.yaml
```

### Logs
```shell
[<your_login>@gateway ~]$ kubectl get pods
[<your_login>@gateway ~]$ kubectl logs <POD_NAME> --timestamps >> lab0-jupyter.log
```

> You're Genius 🗿


[kubectl Cheat Sheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/)