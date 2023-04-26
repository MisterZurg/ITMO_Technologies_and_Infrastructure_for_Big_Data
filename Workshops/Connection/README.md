# Workshop 1 — Сonnecting to Сluster

## Set Down
### 1. Download FortiClient
> Download The VPN-only version of FortiClient offers SSL VPN and IPSecVPN, but does not include any support. Download the best VPN software for multiple devices.

[FortiClient VPN](https://www.fortinet.com/support/product-downloads#vpn)

### 2. Add to /etc/hosts
#### Mac Gateway
```shell
10.32.7.101 gateway.st
10.32.7.103 node03.st
10.32.7.104 node04.st
10.32.7.105 node05.st
10.32.7.106 node06.st
10.32.7.107 node07.st
10.32.7.108 node08.st
10.32.7.109 node09.st
10.32.7.110 node10.st
10.32.7.111 node11.st
10.32.7.112 node12.st
10.32.7.113 node13.st
10.32.7.114 node14.st
10.32.7.115 node15.st
10.32.7.116 node16.st
10.32.7.117 node17.st
10.32.7.118 node18.st
10.32.7.119 node19.st
10.32.7.120 node20.st
10.32.7.121 node21.st
10.32.7.122 node22.st
10.32.7.123 node23.st
10.32.7.124 node24.st
10.32.7.125 node25.st
10.32.7.126 node26.st
10.32.7.127 node27.st
10.32.7.128 node28.st
10.32.7.129 node29.st
10.32.7.130 node30.st
10.32.7.130 node30.st
10.32.7.133 node33.st
10.32.7.134 node34.st
10.32.7.135 node35.st
10.32.7.136 node36.st
10.32.7.137 node37.st
10.32.7.138 node38.st
10.32.7.139 node39.st
10.32.7.140 node40.st
10.32.7.141 node41.st
10.32.7.142 node42.st
10.32.7.143 node43.st
10.32.7.144 node44.st
10.32.7.145 node45.st
10.32.7.146 node46.st
10.32.7.147 node47.st
10.32.7.148 node48.st
10.32.7.149 node49.st
10.32.7.150 node50.st
10.32.7.151 node51.st
10.32.7.152 node52.st
10.32.7.153 node53.st
10.32.7.154 node54.st
10.32.7.155 node55.st
10.32.7.156 node56.st
10.32.7.157 node57.st
10.32.7.158 node58.st
10.32.7.159 node59.st
10.32.7.160 node60.st
```
#### Windows Gateway
```shell
10.32.7.101 gateway.st
10.32.7.200 docker.st
10.32.7.101 gateway.st
10.32.7.103 node03.st
10.32.7.104 node04.st
10.32.7.105 node05.st
10.32.7.106 node06.st
10.32.7.107 node07.st
10.32.7.108 node08.st
10.32.7.109 node09.st
10.32.7.110 node10.st
10.32.7.111 node11.st
10.32.7.112 node12.st
10.32.7.113 node13.st
10.32.7.114 node14.st
10.32.7.115 node15.st
10.32.7.116 node16.st
10.32.7.117 node17.st
10.32.7.118 node18.st
10.32.7.119 node19.st
10.32.7.120 node20.st
10.32.7.121 node21.st
10.32.7.122 node22.st
10.32.7.123 node23.st
10.32.7.124 node24.st
10.32.7.125 node25.st
10.32.7.126 node26.st
10.32.7.127 node27.st
10.32.7.128 node28.st
10.32.7.129 node29.st
10.32.7.130 node30.st
10.32.7.133 node33.st
10.32.7.134 node34.st
10.32.7.135 node35.st
10.32.7.136 node36.st
10.32.7.137 node37.st
10.32.7.138 node38.st
10.32.7.139 node39.st
10.32.7.140 node40.st
10.32.7.141 node41.st
10.32.7.142 node42.st
10.32.7.143 node43.st
10.32.7.144 node44.st
10.32.7.145 node45.st
10.32.7.146 node46.st
10.32.7.147 node47.st
10.32.7.148 node48.st
10.32.7.149 node49.st
10.32.7.150 node50.st
10.32.7.151 node51.st
10.32.7.152 node52.st
10.32.7.153 node53.st
10.32.7.154 node54.st
10.32.7.155 node55.st
10.32.7.156 node56.st
10.32.7.157 node57.st
10.32.7.158 node58.st
10.32.7.159 node59.st
10.32.7.160 node60.st
```

### 3. FortiClient New VPN Connection
| Connection Name |             any-name             |
|:---------------:|:--------------------------------:|
| Remote Gateway  |  `fgvpn.onti.actcognitive.org`   |
|    Username     | Put anything, `AMOGUS` for example |

### 4. Forti Client Start VPN
#### 4.1 Get Creds
1. Go to ISU
2. Click open outlook.office365.com
3. Find letter with title *VPN and cluster credentials* from `itmo.bigdatacourse@gmail.com`
There will be smth like:
> ```shell
> Login:    <your_login>
> Password: <your_password>
> ```

#### 4.2
1. Fill the data
2. Check the connectivity
```shell
ping 10.32.7.101
PING 10.32.7.101 (10.32.7.101): 56 data bytes
64 bytes from 10.32.7.101: icmp_seq=0 ttl=61 time=28.850 ms
64 bytes from 10.32.7.101: icmp_seq=1 ttl=61 time=42.152 ms
64 bytes from 10.32.7.101: icmp_seq=2 ttl=61 time=29.246 ms
```

### Connecting to Cluster
```shell
ssh <your_login>@gateway.st
password >>> <your_password>
```

> Horray you're connected!

### Cluster Stuff
#### Checking ports
```shell
[<your_login>@gateway ~]$ kubectl get pods
No resources found in  <your_login> namespace.
```

See <your_login>.yaml
```
[<your_login>@gateway ~]$ ls
<your_login>-jupyter.yaml
[<your_login>@gateway ~]$ vim <your_login>-jupyter.yaml
```
### Inside `<your_login>-jupyter.yaml`
```yaml
---
apiVersion: v1
kind: Service
metadata:
  name: jupyter-spark-svc
  namespace: <your_login> # Your namespave stuff
spec:
  type: NodePort
  ports:
    - port: 8888
      protocol: TCP
      name: jupyter
    - port: 4040
      protocol: TCP
      name: spark
  selector:
    app: jupyter-spark
# ...
```

### Apply yaml to start services
```shell
[<your_login>@gateway ~]$ kubectl apply -f <your_login>-jupyter.yaml
deployment.apps/jupyter-spark created
service/jupyter-spark-svc created
```
Check started pods
```shell
[<your_login>@gateway ~]$ kubectl get pods
NAME                             READY   STATUS    RESTARTS   AGE
jupyter-spark-7c5b4455cc-fv7dr   1/1     Running   0          20s
```
Check started services
```shell
[<your_login>@gateway ~]$ kubectl get svc
NAME                TYPE       CLUSTER-IP      EXTERNAL-IP   PORT(S)                         AGE
jupyter-spark-svc   NodePort   10.129.154.22   <none>        8888:32074/TCP,4040:31385/TCP   63s
```

## Acsessing servises cheking the labels
```shell
[<your_login>@gateway ~]$ kubectl describe svc jupyter-spark-svc 
Name:                     jupyter-spark-svc
Namespace:                <your_login>
Labels:                   <none>
Annotations:              <none>
Selector:                 app=jupyter-spark
Type:                     NodePort
IP Families:              <none>
IP:                       10.129.154.22
IPs:                      10.129.154.22
Port:                     jupyter  8888/TCP
TargetPort:               8888/TCP
NodePort:                 jupyter  32074/TCP
Endpoints:                10.128.202.147:8888
Port:                     spark  4040/TCP
TargetPort:               4040/TCP
NodePort:                 spark  31385/TCP
Endpoints:                10.128.202.147:4040
Session Affinity:         None
External Traffic Policy:  Cluster
Events:                   <none>
```
## Connect to jupyter-server
### Get the external port
```shell
[<your_login>@gateway ~]$ kubectl get svc
NAME                TYPE       CLUSTER-IP      EXTERNAL-IP   PORT(S)                         AGE
jupyter-spark-svc   NodePort   10.129.154.22   <none>        8888:32074/TCP,4040:31385/TCP   5m25s
```
> 8888:**32074**
### Type in browsers URL
```
node03.st:32074
```
### Get token
```shell
[<your_login>@gateway ~]$ kubectl logs jupyter-spark-7c5b4455cc-fv7dr 
<outout ommited>    
    To access the notebook, open this file in a browser:
        file:///home/jovyan/.local/share/jupyter/runtime/nbserver-1-open.html
    Or copy and paste one of these URLs:
        http://jupyter-spark-7c5b4455cc-fv7dr:8888/?token=a3954971b43ab0e584aab51b99be344656ff011fbe911b18
     or http://127.0.0.1:8888/?token=a3954971b43ab0e584aab51b99be344656ff011fbe911b18
[I 17:41:00.537 NotebookApp] 302 GET / (10.128.103.192) 0.830000ms
[I 17:41:00.631 NotebookApp] 302 GET /tree? (10.128.103.192) 1.030000ms
```
> ?token=*COPY_THAT_PART*

Put that token into `node03.st:32074` input field

### When you're done
```shell
[<your_login>@gateway ~]$ kubectl delete pod jupyter-spark-7c5b4455cc-fv7dr
```
> When you delete pod and token is expired and stuff that you store note in
> ```shell
>   /
>   | - 📁 shared-data
>   | - 📁 nsf-home
> ```
> **will be lost!**


## Creating permanent token
> F security, lets hardcode the token
1. Open yaml
```shell
[<your_login>@gateway ~]$ vim <your_login>-jupyter.yaml
```
2. Uncomment `args`
```yaml
      containers:
      - name: jupyter
        ...
        command: [ "start-notebook.sh" ]
        # args: [ "--NotebookApp.token=''" ]  # Here you can specify your password for Jupyter Server
```
3. Put any token you like for example `P@ssw0rD`
```yaml
      containers:
      - name: jupyter
        ...
        command: [ "start-notebook.sh" ]
        args: [ "--NotebookApp.token='P@ssw0rD'" ]  # Here you can specify your password for Jupyter Server
```
4. Save yaml
> Now you can login by inputing the `P@ssw0rD`

## Usefull comands
> `scp` Copies files from local machine into remote srv
```shell
[<your_login>@gateway ~]$ scp ..... login@gateway.st:~
```