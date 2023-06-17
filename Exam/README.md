# BigData-Exam-Questions
## Basic concepts of BigData
### What is Big Data? History of Big Data, V-concepts.
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

### Streaming solutions for Big Data.

### MapReduce concept and its main steps.

## Storages and resource managers - supplementary tools
What is HDFS?  HDFS Architecture. What is a chunk? How are chunks being processed into records? How to customize the way chunks are being processed?
Replica in HDFS. Safe mode. NameNode. DataNode.
CAP theorem. Communication Protocols. Data consistency.
Data Formats: txt,csv,binary, json,parquet. What's better?

## Apache Spark
What is Apache Spark? Which components does its architecture consist of? 
What is RDD? Which principal components does it consist of? What is a transformation? What is an action?
What is RDD lineage? How is it formed? How can different RDD lineages be combined?
What is a partition in RDD? How does a partition in RDD relate to a partition in HDFS?
What is broadcast and a broadcasting variable in Apache Spark? What goals does it serve? Give an example of its usage.
What is caching and persistence in Apache Spark? What is the difference between caching and persistence in Apache Spark?  Which modes of persistence exist? What is spillover? 
What is shuffling? Explain how it works. 

## Spark SQL 
What is SparkSQL? What is SparkSession? What is DataFrame? What is Column? Name differences between RDD and DataFrame. What is Tungsten?
UDF in Spark SQL. How to write an UDF (give a simple example)? What pitfalls should you be aware about when writing a custom udf?
What is a schema of DataFrame? How to get a schema of DataFrame? What is groupby and how does it work?  
Supported join operations in Spark SQL. Broadcast joins. Basic principles of functioning. Which types of broadcast-based joins exist and what are the differences in their functioning? Conditions of when one can apply these types of joins.
Sort Merge Join. What is it for? Basic principles of functioning. Conditions of when one can apply these types of joins. Advantages and disadvantages in comparison with other joins. 
Bucketing: what it is, its advantages and basic principles of functioning. 

## ClickHouse
What is ClickHouse? Which components does its architecture consist of? 
What is Materialized View and what is its purpose? How does Materialized Views work? Does it make sense to create a materialized view over a distributed table? If the answer is positive, why?
Describe MergeTree engine work principles. How does ClickHouse store data in the filesystem? How does parts merging works from in terms of files?
What is sharding, why is it useful and how is it implemented in ClickHouse? How can one distribute data into shards using ClickHouse?
How does AggregatingMergeTree work? Are there any differences between aggregation functions in AggregatingMergeTree and simple queries?
Why is it recommended inserting data into ClickHouse by large batches? What can one do if data arrives in short portions frequently? Name and explain existing mechanisms in ClickHouse to deal with the problem.
Describe sparse index work principles. How is it implemented in ClickHouse? How can we set up indexes in ClickHouse?

## Kubernetes
What is Kubernetes? Its architecture, main components.
What is Pod? Difference between pod and container.
What is Service? How does it work? Service types: nodeport, clusterip, headless service.
Container resources: how is it implemented? CPU, RAM, Storage. QoS classes.
PV & PVC: what is it and how does it work? PV types. Volume provisioner.
StatefulSet. Difference from Deployment/Pod.
Label selector. What is it? How is it used? 
