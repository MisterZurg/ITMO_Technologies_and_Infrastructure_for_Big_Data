# BigData-Exam-Questions
## Basic concepts of BigData
### 1. What is Big Data? History of Big Data, V-concepts.
Big Data is data sets that are so voluminous and complex that can not be processed within single PC in reasonable time.

At the end of 90 s and begging of 2000 s the large data could be found
- Retail companies
- the company with queuing service internet operators, mobile operators, banks
- stock exchanges
- large corporations

The separate place was allocated to following scientific areas
- biomedicine
- hydrometeorology
- astronomy
- high energy physics

Main concept used before 2009 year **high load databases**

Volume  the size of the data does not allow to be processed within a reasonable time or be physically placed on the one computing node (or one server)
Variety diversity, and the lack of structuring (relational DBMS can not be apply directly).
Velocity growth rate and update data (more than 60% per year)
Veracity truthfulness and precision (introduced by IBM)
Value importance, price

### 2. Streaming solutions for Big Data.
Streaming data means a continuous flow of data from various sources – data streams can be processed, stored and analyzed using special technologies as it is generated in real-time.
The goal of this streaming is to ensure a constant flow of data to be processed without needing to download it from the source first.

**Google Cloud DataFlow**
- Autoscaling of resources and dynamic work rebalancing (Data inputs are partitioned automatically and constantly rebalanced to even out worker resource utilization and reduce the effect of “hot keys” on pipeline performance)
- Scheduling and pricing for batch processing
- Ready-to-use real-time AI patterns (real-time reactions with near-human intelligence to large torrents of events)

**Amazon Kinesis**
- Capturing, processing and storing video/data streams
- Streaming data analysis using SQL or Apache Flink
- Integration with Amazon AI/ML services

**Azure Stream Analytics**
- SQL syntax and extensible with JavaScript and C# custom code
- Enterprise-grade reliability with built-in recovery and built-in machine learning
- Security and compliance, built in (ISO, CSA/CMM, HIPAA, IRS 1075)

**Apache streaming projects** 
- Kafka (distributed replicated message broker)
- Flink (calculation for continios datastream, no storage so it has connectors to diff solutions)

**RabbitMQ**
- Message Queue FIFO
- Balancing Pull Model

### MapReduce concept and its main steps.
Distributed computing model for processing Big Data. 
We can implement an algorithm in map reduce paradigm, used wen classic algorithm struggles by memory
1) Map - data processing.
2) Reduce - data convolution.

Input files, are splitted so mappers (workers on stage Map) stores data on drive.
Going to Reduce stage, data procced by mappers is stored as a key->value. 
Data with equal keys comes on the same reducer (worker on stage Reduce)
Result stored in output files

Input data should be splittable (split = block in hdfs hadoop distributed), data in split should be independent.
One worker process one split.
Middleware data stores on local disk.
For each Reducer, Mapper creates Key->Value.
Data with the same Key goes into one Reducer.
Reducers start working when all mappers finished processing.
Data from Reducer "N" is written output file "N", which are stored in hdfs.

## Storages and resource managers - supplementary tools
### What is HDFS? HDFS Architecture. What is a chunk? How are chunks being processed into records? How to customize the way chunks are being processed?
Hadoop Distributed File System - open source analog of Google's GFS. 
It works on a cluster of server, and for user is viewed as one big drive.
It's application above standart filesystems. Ext3, Ext4, XFS.
Fault tolerant - data is not lost, if one or some disck/server doesn't lost. 

HDFS supports write-once-read-many semantics on files.

### Replica in HDFS. Safe mode. NameNode. DataNode.
Files and blocks are replicated wile writing:
- Same block is stored on several Datanodes
- Replicas by default (replication factor) is eq 3 (for example 1 replica would be stored on first Rack local machine second on the same Rack but different machine, and the third replica)
Daemon processes:
* Namenode - runs on 1-st **dedicated** machine; responsible for namespace, meta information, where data is stored on cluster.
* Datanode - runs on each cluster machine, stores data blocks, sends its state to Namenode
* Secondary Namenode - (optional daemon) updates fsimage, logs from datanode. It's not a backup of Namenode

Safemode in Apache Hadoop is a maintenance state of **NameNode**, during which NameNode doesn’t allow any modifications to the file system.
In Safemode, HDFS cluster is in read-only and doesn’t replicate or delete Data Blocks.

> _When NameNode starts:_
> It loads the file system namespace from the last saved FsImage into its main memory and the edits log file.
> Merges edits log file on fsimage and results in new file system namespace.
> Then it receives block reports containing information about block location from all datanodes. **In SafeMode** NameNode performs collection of block reports from datanodes. NameNode enters safemode automatically during its start up.
> NameNode leaves Safemode after the DataNodes have reported that most blocks are available.

### CAP theorem. Communication Protocols. Data consistency.
CAP theorem states that any distributed data store can provide only two of the following three guarantees:
* **Consistency** - Every read receives the most recent write or an error.
* **Availability** - Every request receives a (non-error) response, without the guarantee that it contains the most recent write.
* **Partition tolerance** - The system continues to operate despite an arbitrary number of messages being dropped (or delayed) by the network between nodes.

When a network partition failure happens, it must be decided whether to do one of the following:
* cancel the operation and thus decrease the availability but ensure consistency
* proceed with the operation and thus provide availability but risk inconsistency.

**best-effort broadcast protocol**
guarantees that if the sender does not crash, the message is delivered to all non-faulty processes in a group. A simple way to implement it is to send the message to all processes in a group one-by-one over reliable links but if sender fails mid-way, some processes will never receive the message.

**reliable broadcast protocol**
guarantees that the message is eventually delivered to all non-faulty processes in the group, event if the sender crashes before the message has been fully delivered. One way to implement reliable broadcast is to have each processes re-transmit the message to the rest of the group the first time it is delivered. This approach is also known as eager reliable broadcast -ERB. Although it guarantees that all non-faulty processes eventually recieve the message, It’s costly as it requires sending the message N² times for a group of N processes. The number of messages can be reduced by re-transmitting a message on delivery to a random subset of processes.

**gossip protocol**
It is a probabilistic protocol, it does not guarantee that a message will be delivered to all processes. That said, it is possible to make that probability negligible by tuning the protocol’s parameters. Gossip protocols are particularly useful when broadcasting to a large number of processes, like thousands or more, where deterministic protocol would not scale.

Data consistency is a crucial aspect that ensures the accuracy and reliability of data. So if data is inconsistent, there is nothing right: 

Data consistency is the accuracy, completeness, and correctness of data stored in a database. 
The same data across all related systems, applications, and databases is when we say that data is consistent. 
Inconsistent data can lead to incorrect analysis, decision-making, and outcomes.
The key metrics such as accuracy, completeness, timeliness, and relevance are used to analyze or measure data consistency.

### Data Formats: txt, csv, binary, json, parquet. What's better?
For data lakes, in the Hadoop ecosystem, HDFS file system is used. 
However, most cloud providers have replaced it with their own deep storage system such as S3 or GCS. 
When using deep storage choosing the right file format is crucial, and you need to consider the format, compression and especially how you partition your data.

The most common formats are TXT, CSV, BINARY, JSON, PARQUET
- The structure of your data
- Performance
- Easy to read
- Compression
- Schema evolution
- Compatibility

**TXT:** Plain text file that contains unformatted text. Great choice for storing and sharing simple, unformatted information, such as notes, lists, or code snippets.

**CSV:** Good option for compatibility, spreadsheet processing and human readable data. The data must be flat. It is not efficient and cannot handle nested data. There may be issues with the separator which can lead to data quality issues. Use this format for exploratory analysis, POCs or small data sets.

**BINARY:** Performance, small size. Not human readable. Advantages in terms of speed of access.

**JSON:** Heavily used in APIs. Nested format. It is widely adopted and human readable but it can be difficult to read if there are lots of nested fields. Great for small data sets, landing data or API integration. If possible convert to more efficient format before processing large amounts of data.

**Parquet:** Columnar storage. It has schema support. It works very well with Hive and Spark as a way to store columnar data in deep storage that is queried using SQL. Because it stores data in columns, query engines will only read files that have the selected columns and not the entire data set as opposed to Avro. Use it as a reporting layer.


## Apache Spark
### What is Apache Spark? Which components does its architecture consist of? 
Apache Spark™ is a multi-language engine for executing data engineering, data science, and machine learning on single-node machines or clusters.


### What is RDD? Which principal components does it consist of? What is a transformation? What is an action?

Composition of RDD
**partition**
When an RDD is generated by using an RDD operation, a piece of data is divided into n pieces. 
Each piece of data corresponds to a partition in the RDD, and the RDD stores an array of partitions. 
The number of partitions can be specified during RDD initialization. If the number is not specified, read the spark.default.parallelism configuration item to set the number.

**`compute()`**
The compute function is used to calculate each partition and is implemented by a subclass of RDD.
Each RDD has a compute function, which is responsible for receiving the partition transferred by the previous RDD.

**dependencies**
Indicate the dependency between an RDD and other RDDs. Dependencies are generated by RDD operations, for example, val lineLengths = lines.map(s => s.length). 
This indicates that RDD lineLengths depends on RDD lines. RDD lines are also called the parent RDD of RDD lineLengths.


RDD transformations – Transformations are lazy operations, instead of updating an RDD, these operations return another RDD.
RDD actions – operations that trigger computation and return RDD values.


### What is RDD lineage? How is it formed? How can different RDD lineages be combined?


### What is a partition in RDD? How does a partition in RDD relate to a partition in HDFS?
### What is broadcast and a broadcasting variable in Apache Spark? What goals does it serve? Give an example of its usage.
### What is caching and persistence in Apache Spark? What is the difference between caching and persistence in Apache Spark?  Which modes of persistence exist? What is spillover? 
### What is shuffling? Explain how it works. 

## Spark SQL 
### What is SparkSQL? What is SparkSession? What is DataFrame? What is Column? Name differences between RDD and DataFrame. What is Tungsten?
### UDF in Spark SQL. How to write an UDF (give a simple example)? What pitfalls should you be aware about when writing a custom udf?
### What is a schema of DataFrame? How to get a schema of DataFrame? What is groupby and how does it work?  
### Supported join operations in Spark SQL. Broadcast joins. Basic principles of functioning. Which types of broadcast-based joins exist and what are the differences in their functioning? Conditions of when one can apply these types of joins.
### Sort Merge Join. What is it for? Basic principles of functioning. Conditions of when one can apply these types of joins. Advantages and disadvantages in comparison with other joins. 
### Bucketing: what it is, its advantages and basic principles of functioning. 

## ClickHouse
### What is ClickHouse? Which components does its architecture consist of? 
### What is Materialized View and what is its purpose? How does Materialized Views work? Does it make sense to create a materialized view over a distributed table? If the answer is positive, why?
### Describe MergeTree engine work principles. How does ClickHouse store data in the filesystem? How does parts merging works from in terms of files?
### What is sharding, why is it useful and how is it implemented in ClickHouse? How can one distribute data into shards using ClickHouse?
### How does AggregatingMergeTree work? Are there any differences between aggregation functions in AggregatingMergeTree and simple queries?
### Why is it recommended inserting data into ClickHouse by large batches? What can one do if data arrives in short portions frequently? Name and explain existing mechanisms in ClickHouse to deal with the problem.
### Describe sparse index work principles. How is it implemented in ClickHouse? How can we set up indexes in ClickHouse?

## Kubernetes
### What is Kubernetes? Its architecture, main components.
K8s is an open-source system for automating deployment, scaling, and management of containerized applications.
> **Control Plane** - The container orchestration layer that exposes the API and interfaces to define, deploy, and manage the lifecycle of containers.

**ApiServer**
The API server is a component of the Kubernetes control plane that exposes the Kubernetes API. The API server is the front end for the Kubernetes control plane.
kube-apiserver is designed to scale horizontally—that is, it scales by deploying more instances. You can run several instances of kube-apiserver and balance traffic between those instances.


**etcd**
Consistent and highly-available key value store used as Kubernetes' backing store for all cluster data.

**Scheduler**
Control plane component that watches for newly created Pods with no assigned node, and selects a node for them to run on.

Factors taken into account for scheduling decisions include: individual and collective resource requirements, hardware/software/policy constraints, affinity and anti-affinity specifications, data locality, inter-workload interference, and deadlines.

**Kube-Controller** 
Control plane component that runs controller processes. 
> a control loop is a non-terminating loop that regulates the state of a system on clusters and moves the current state into disired.



**Kubelet** 
An agent that runs on each node in the cluster. It makes sure that containers are running in a Pod.
 
The kubelet takes a set of PodSpecs that are provided through various mechanisms and ensures that the containers described in those PodSpecs are running and healthy. The kubelet doesn't manage containers which were not created by Kubernetes.

**Kube-Proxy** 
It is a network proxy that runs on each node in your cluster, implementing part of the Kubernetes Service concept.
kube-proxy maintains network rules on nodes. These network rules allow network communication to your Pods from network sessions inside or outside of your cluster.
kube-proxy uses the operating system packet filtering layer if there is one and it's available. Otherwise, kube-proxy forwards the traffic itself.


### What is Pod? Difference between pod and container.
Pods are the smallest deployable units of computing that you can create and manage in Kubernetes.

A Pod (as in a pod of whales or pea pod) is a group of one or more containers, with shared storage and network resources, and a specification for how to run the containers.

Container is a lightweight and portable executable image that contains software and all of its dependencies.

Each node in a Kubernetes cluster runs the containers that form the Pods assigned to that node. Containers in a Pod are co-located and co-scheduled to run on the same node.

### What is Service? How does it work? Service types: nodeport, clusterip, headless service.
In Kubernetes, a Service is a method for exposing a network application that is running as one or more Pods in your cluster.

A key aim of Services in Kubernetes is that you don't need to modify your existing application to use an unfamiliar service discovery mechanism. You can run code in Pods, whether this is a code designed for a cloud-native world, or an older app you've containerized. You use a Service to make that set of Pods available on the network so that clients can interact with it.

The Service API, part of Kubernetes, is an abstraction to help you expose groups of Pods over a network. Each Service object defines a logical set of endpoints (usually these endpoints are Pods) along with a policy about how to make those pods accessible.

**NodePort**
Exposes the Service on each Node's IP at a static port (the NodePort). To make the node port available, Kubernetes sets up a cluster IP address, the same as if you had requested a Service of type: ClusterIP.

**ClusterIP**
Exposes the Service on a cluster-internal IP. Choosing this value makes the Service only reachable from within the cluster. This is the default that is used if you don't explicitly specify a type for a Service. You can expose the Service to the public internet using an Ingress or a Gateway.

**Headless Services**
Sometimes you don't need load-balancing and a single Service IP. In this case, you can create what are termed headless Services, by explicitly specifying "None" for the cluster IP address (.spec.clusterIP).

For headless Services, a cluster IP is not allocated, kube-proxy does not handle these Services, and there is no load balancing or proxying done by the platform for them.

### Container resources: how is it implemented? CPU, RAM, Storage. QoS classes.
**Limits**
- Number of resources that POD can use
- Upper bound
**Requests**
- Number of resources that is reserved for POD on Node
- Not shared between PODs on node

> Prevent negative influence due to unexpected behaviour
> Efficient resources utilization
> Scheduling is better

Quality of Service (QoS) class determines the pod’s scheduling and eviction priority. QoS class is used by the Kubernetes scheduler to make decisions about scheduling pods onto nodes.

QoS classes:
Guaranteed
- Requests and Limits are the same OR
- No Requests, only Limits
Burstable
- Requests and limits are specified and they are different. OR
- There are no limits specified.
BestEffort
- Requests and limits are not specified
“I do not care if my application receives enough resources”
Bad choice for CPU intensive application

- configMap - provides a way to inject configuration data into pods. The data stored in a ConfigMap can be referenced in a volume of type configMap and then consumed by containerized applications running in a pod.
- emptyDir - An emptyDir volume is first created when a Pod is assigned to a node, and exists as long as that Pod is running on that node. As the name says, the emptyDir volume is initially empty. All containers in the Pod can read and write the same files in the emptyDir volume, though that volume can be mounted at the same or different paths in each container. When a Pod is removed from a node for any reason, the data in the emptyDir is deleted permanently.
- hostPath - A hostPath volume mounts a file or directory from the host node's filesystem into your Pod.
  
### PV & PVC: what is it and how does it work? PV types. Volume provisioner.
When doing some cloud stuff like migraiting from Azure to Google. 

PersistentVolumeClaim describes requrements of volume. Disk size, file system, access type.
PersistentVolume stores params and volume status
Storage Class stores connection params.

There are a number types of PV:
- awsElasticBlockStore
- azureDisk
- cephfs
- csi
- gcePersistentDisk
- glusterfs
- local
- nfs
...

Provisioning

**Static**
A cluster administrator creates a number of PVs. They carry the details of the real storage, which is available for use by cluster users. They exist in the Kubernetes API and are available for consumption.

**Dynamic**
When none of the static PVs the administrator created match a user's PersistentVolumeClaim, the cluster may try to dynamically provision a volume specially for the PVC. This provisioning is based on StorageClasses: the PVC must request a storage class and the administrator must have created and configured that class for dynamic provisioning to occur. Claims that request the class "" effectively disable dynamic provisioning for themselves.

### StatefulSet. Difference from Deployment/Pod.
Manages the deployment and scaling of a set of Pods, and provides guarantees about the ordering and uniqueness of these Pods.

Unlike a Deployment, a StatefulSet maintains a sticky identity for each of its Pods. These pods are created from the same spec, but are not interchangeable: each has a persistent identifier that it maintains across any rescheduling.

StatefulSets are valuable for applications that require one or more of the following.
* Stable, unique network identifiers.
* Stable, persistent storage.
* Ordered, graceful deployment and scaling.
* Ordered, automated rolling updates.


### Label selector. What is it? How is it used? 
> Labels are key/value pairs that are attached to objects such as Pods. 
> Labels are intended to be used to specify identifying attributes of objects that are meaningful and relevant to users, but do not directly imply semantics to the core system. 
> Labels can be used to organize and to select subsets of objects. 
> Labels can be attached to objects at creation time and subsequently added and modified at any time. Each object can have a set of key/value labels defined.

Labels enable users to map their own organizational structures onto system objects in a loosely coupled fashion, without requiring clients to store these mappings.

Label selectors
Unlike names and UIDs, labels do not provide uniqueness. In general, we expect many objects to carry the same label(s).

Via a label selector, the client/user can identify a set of objects.

The API currently supports two types of selectors: equality-based and set-based.
- Equality- or inequality-based requirements allow filtering by label keys and values. Matching objects must satisfy all of the specified label constraints, though they may have additional labels as well. Three kinds of operators are admitted =,==,!=.
- Set-based label requirements allow filtering keys according to a set of values. Three kinds of operators are supported: in,notin and exists (only the key identifier).
