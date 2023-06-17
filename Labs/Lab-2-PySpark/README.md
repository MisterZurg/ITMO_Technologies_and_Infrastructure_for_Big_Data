# Lab 2 — PySpark

> <picture>
>   <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/light-theme/warning.svg">
>   <img alt="Warning" src="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/dark-theme/warning.svg">
> </picture><br>
> 
> **PLAGIARISM IS PROHIBITED!**
> <details>
>  <summary>ByPassing</summary>
>
>  `name()` is an alias for `alias()`
> 
>  `orderBy()` is an alias for `sort()`
>  
>  `F.col("col_name").asc()` === `F.asc("col_name")`
>
>  `F.col("col_name.propery").getItem(0)` === `F.col('col_name')['propery'][0]`
> 
>  `.where(F.col("col_name").isNotNull())` ===   `.where("col_name <> ''")` === `.where("col_name is not null")`
> 
> `groupBy(["col_1", "col_2"])` == `groupBy(F.col("col_1"), F.col("col_2"))`
> </details>

## Connect to Cluster
> <picture>
>   <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/light-theme/info.svg">
>   <img alt="Info" src="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/dark-theme/info.svg">
> </picture><br>
> <details>
>  <summary>Collapse</summary>
>
> 1. Start and connect VPN
> 2. Check the connectivity
> ```shell
>(base) misterzurg@MacBook-Pro-Denis ~ % ping 10.32.7.101
> PING 10.32.7.101 (10.32.7.101): 56 data bytes
> 64 bytes from 10.32.7.101: icmp_seq=0 ttl=61 time=30.223 ms
> ```
> 3. Connect to Cluster
> ```shell
>(base) misterzurg@MacBook-Pro-Denis ~ % ssh <your_login>@gateway.st
> password >>> <your_password>
> ```
> 
> 4. Get the external port
> ```shell
> [<your_login>@gateway ~]$ kubectl get svc
> NAME                TYPE       CLUSTER-IP      EXTERNAL-IP   PORT(S)                         AGE
> jupyter-spark-svc   NodePort   10.129.154.22   <none>        8888:32074/TCP,4040:31385/TCP   5m25s
> ```
> 
> 5. Type in browsers URL
> ```
> node03.st:32074
> ```
> 
> 6. Get token
> ```shell
> [<your_login>@gateway ~]$ kubectl logs jupyter-spark-7c5b4455cc-fv7dr 
> ```
> </details>

## General info
> <picture>
>   <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/light-theme/info.svg">
>   <img alt="Info" src="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/dark-theme/info.svg">
> </picture><br>
> 
> To submit the lab, please put the completed notebook in your **NFS home directory** in the following path:
> 
> (path in **gateway.st**) `/nfs/home/<your-login>/spark-lab.ipynb`
> 
> (path in **jupyter containers**) `/home/jovyan/nfs-home/spark-lab.ipynb`
> 
> This notebook is available in your Jupyter servers by following path /home/jovyan/shared-data/notebooks/BDML/SparkLab-Template.ipynb.

## How to Get template and Save
> <picture>
>   <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/light-theme/info.svg">
>   <img alt="Info" src="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/dark-theme/info.svg">
> </picture><br>
> <details>
>  <summary>Collapse</summary>
>
> 1. Start and connect VPN
> 2. Open jupyter workspace `http://node03.st:<your_port>`
> 3. Go to `shared-data/notebooks/BDML/SparkLab-Template.ipynb.`
> 4. Download `SparkLab-Template.ipynb` to your Device
> 5. Go back to home of jupyter workspace
> 6. Click on `nfs-home` you'll see 
>```shell
> 📁 / nfs-home
>   | - 📁 ...
>   empty
> ```
> 7. Click `Upload` button at the top right corner, to upload `SparkLab-Template.ipynb`
> 8. Rename it to `spark-lab.ipynb` and you're ready to work!
> > If everything is ok the `nfs-home` should look smth like that
> > ```shell
> >  📁 / nfs-home
> >   | - 📁 ...
> >   | - 📁 conf
> >   | - 📁 logs
> >   | - 📁 tmp
> >   | - 📓 spark-lab.ipynb
> > ```
> > And `Running` menu:
> > ```shell
> >  Notebooks
> >   📓 shared-data/notebooks/BDML/SparkLab-Template.ipynb
> >   📓 nfs-home/spark-lab.ipynb
> > ```
> </details>


## 1. Find top posts:
### 1a. by likes count (1 point)
```python
def task_1a(df: "pyspark.sql.dataframe.DataFrame", 
            F: "pyspark.sql.functions") -> "pyspark.sql.dataframe.DataFrame":
    # Place your code to transform DaraFrame here
    # Filed id indicates post id.
    modified_df = df.select(
            F.col("id").alias("post_id"), 
            F.col("likes.count").alias("likes_count")
        ).orderBy(
            F.col("likes_count").desc(),
            F.col("post_id").asc()
    )
    
    return modified_df
```
### 1b. by comments count (1 point)
```python
def task_1b(df: "pyspark.sql.dataframe.DataFrame", 
            F: "pyspark.sql.functions") -> "pyspark.sql.dataframe.DataFrame":
    # Place your code to transform DaraFrame here
    modified_df = df.select(
            F.col("id").alias("post_id"), 
            F.col("comments.count").alias("comments_count")
        ).orderBy(
            F.col("comments_count").desc(),
            F.col("post_id").asc()
    )
    
    return modified_df
```
### 1c. by reposts count (1 point)
```python
def task_1c(df: "pyspark.sql.dataframe.DataFrame", 
            F: "pyspark.sql.functions") -> "pyspark.sql.dataframe.DataFrame":
    # Place your code to transform DaraFrame here
    modified_df = df.select(
            F.col("id").alias("post_id"), 
            F.col("reposts.count").alias("reposts_count")
        ).orderBy(
            F.col("reposts_count").desc(),
            F.col("post_id").asc()
    )
    
    
    return modified_df
```
## 2. Find top users:
### 2a. by likes are received (1 point)
```python
def task_2a(df: "pyspark.sql.dataframe.DataFrame", 
            F: "pyspark.sql.functions") -> "pyspark.sql.dataframe.DataFrame":
    # Place your code to transform DaraFrame here
    # Look through df
    modified_df = df.groupby("ownerId")\
        .agg(F.count("ownerId").name("count"))\
        .orderBy(
            F.col("count").desc(),
            F.col("ownerId").asc()
    )
    
    return modified_df
```
### 2b. by reposts are made (1 point)
```python
def task_2b(df: "pyspark.sql.dataframe.DataFrame", 
            F: "pyspark.sql.functions") -> "pyspark.sql.dataframe.DataFrame":
    # Place your code to transform DaraFrame here 
    modified_df = df.where(F.col("copy_history").isNotNull())\
                    .select(
                        F.col("owner_id"),
                        F.col("copy_history.id").name("src_post_id"),
                    ).groupBy("owner_id").agg(
                        F.count("src_post_id").name("count"))\
                        .orderBy(
                            F.col("count").desc(),
                            F.col("owner_id").asc()
                    )
    
    return modified_df
```
## 3. Extract made users' reposts from the ITMO group (3 points)
```python
def task_3(df: "pyspark.sql.dataframe.DataFrame", 
           F: "pyspark.sql.functions") -> "pyspark.sql.dataframe.DataFrame":
    # Place your code to transform DaraFrame here

    modified_df = df.where(F.col("copy_history").isNotNull())\
                    .select(
                        F.col("id").alias("user_post_id"),
                        F.col("copy_history.id").getItem(0).alias("group_post_id"))\
                    .where(F.col("copy_history.owner_id").getItem(0) == -94)\
                    .groupby("group_post_id")\
                    .agg(F.sort_array(F.collect_list('user_post_id')).alias('user_post_ids')) \
                    .withColumn("reposts_count", F.size("user_post_ids"))\
                    .orderBy(
                        F.col("reposts_count").desc(),
                        F.col("group_post_id").asc()
                    ) # .show(10)
    
    
    return modified_df
```
## 4. Extrace emojis in user posts and find top positive emojis, top neutral emojis and top negative emojis. (3 points)
```python
def task_4(df: "pyspark.sql.dataframe.DataFrame",
           F: "pyspark.sql.functions",
           T: "pyspark.sql.types",
           emojis_data: dict,
           broadcast_func: "spark.sparkContext.broadcast") -> 'Tuple["pyspark.sql.dataframe.DataFrame"]':
    # You are able to modify any code inside this function.
    # Only 'emoji' package import is allowed for this task.
    
    import emoji
    reg_exp = emoji.get_emoji_regexp()

    # 0) Helpers
    @F.udf(returnType=T.ArrayType(T.StringType()))
    def getEmojiSlice(text):
        emojis = []
        for reg_match in reg_exp.finditer(text):
            emojis.append(reg_match.group())
        return emojis

    
    broarcasted_dict = broadcast_func(emojis_data)
    @F.udf(returnType=T.StringType())
    def getSentiment(col):
        return broarcasted_dict.value.get(col, None)

    
    # 1) Gather emoji's from df
    raw_text_emoji_df = df.where(F.col("text").isNotNull())\
                            .select(
                                F.col("id"), 
                                F.col("text")
                            )\
                            .withColumn("emojis", getEmojiSlice("text"))\
                            .where(F.size("emojis") > 0)\
                            .select(F.col("id").name("postID"), F.col("emojis"))

    # 2) Get all emojis counted
    emoji_df = raw_text_emoji_df.select(F.explode("emojis").name("emoji"))\
                               .groupBy("emoji")\
                               .agg(F.count("emoji").name("count"))\
                                .withColumn("sentiment", getSentiment(F.col("emoji")))

    positive_emojis = emoji_df.where(F.col("sentiment") == "positive")\
                                .select(
                                    F.col("emoji"),
                                    F.col("count"))\
                                .orderBy(F.col("count").desc())
    
    neutral_emojis = emoji_df.where(F.col("sentiment") == "neutral")\
                                .select(
                                    F.col("emoji"),
                                    F.col("count"))\
                                .orderBy(F.col("count").desc())
    
    negative_emojis = emoji_df.where(F.col("sentiment") == "negative")\
                                .select(
                                    F.col("emoji"),
                                    F.col("count"))\
                                .orderBy(F.col("count").desc())
    
    return (positive_emojis, neutral_emojis, negative_emojis)
```


> <picture>
>   <source media="(prefers-color-scheme: d)" srcset="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/light-theme/tip.svg">
>   <img alt="Tip" src="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/dark-theme/tip.svg">
> </picture><br>
>
> <details>
>  <summary>Alternative code 1</summary>
> 
> ```python
> def task_4(df: "pyspark.sql.dataframe.DataFrame",
>            F: "pyspark.sql.functions",
>            T: "pyspark.sql.types",
>            emojis_data: dict,
>            broadcast_func: "spark.sparkContext.broadcast") -> 'Tuple["pyspark.sql.dataframe.DataFrame"]':
>     
>     # You are able to modify any code inside this function.
>     # Only 'emoji' package import is allowed for this task.
>     
>     import emoji
>     
>     # 0) Helpers
>     schema = T.ArrayType(T.StringType(), containsNull=False)
>     extractEmoji = lambda s: (c for c in s if c in emoji.UNICODE_EMOJI["en"])
> 
>     
>     # 1) Gather emoji's from df
>    raw_text_emoji_df = df.where(F.col("text").isNotNull())\
>                 .select(
>                     F.col("id").alias("post_id"), 
>                     F.col("text").alias("raw_text")
>                 )\
>                 .withColumn(
>                     "emojis", 
>                     F.udf(lambda s: list(extractEmoji(s)),schema)("raw_text")
>                 )
>     
>     # 2) Get all emojis counted
>     emoji_df = raw_text_emoji_df.where(F.size("emojis") > 0)\
>                 .select(
>                     F.col("post_id"),
>                     F.explode("emojis").alias("emoji"),
>                 ).groupBy("emoji").agg(
>                         F.count("emoji").alias("count"))\
>                 .orderBy(
>                             F.col("count").desc()
>                 )
> 
>     # Helper  I hope
>     broarcasted_variable = broadcast_func(emojis_data)
>     @F.udf(returnType=T.StringType())
>     def udf(col):
>         return broarcasted_variable.value.get(col, None)
>     
>     
>     emoji_df = emoji_df.select(
>             F.col("emoji"),
>             F.col("count")
>         ).withColumn('sentiment', udf(F.col("emoji")))
>     
>     
>     positive_emojis = emoji_df.where(F.col("sentiment") == "positive")\
>                     .select(
>                         F.col("emoji"),
>                         F.col("count")
>                     )
>     neutral_emojis = emoji_df.where(F.col("sentiment") == "neutral")\
>                     .select(
>                         F.col("emoji"),
>                         F.col("count")
>                     )
>     negative_emojis = emoji_df.where(F.col("sentiment") == "negative")\
>                     .select(
>                         F.col("emoji"),
>                        F.col("count")
>                     )
>     
>    
>     return (positive_emojis, neutral_emojis, negative_emojis)
> ```
> </details>

> <picture>
>   <source media="(prefers-color-scheme: d)" srcset="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/light-theme/tip.svg">
>   <img alt="Tip" src="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/dark-theme/tip.svg">
> </picture><br>
>
> <details>
>  <summary>Alternative code 2</summary>
> 
> ```python
> def task_4(df: "pyspark.sql.dataframe.DataFrame",
>            F: "pyspark.sql.functions",
>            T: "pyspark.sql.types",
>            emojis_data: dict,
>            broadcast_func: "spark.sparkContext.broadcast") -> 'Tuple["pyspark.sql.dataframe.DataFrame"]':
>     
>     # You are able to modify any code inside this function.
>     # Only 'emoji' package import is allowed for this task.
>     
>     import emoji
>     
>     broarcasted_variable = broadcast_func(emojis_data)
>     
>     @F.udf(returnType=T.StringType())
>     def udf(col):
>         if type(col) == str:
>             chars = ''.join([k['emoji'] for k in emoji.emoji_lis(col)])
>             if chars != '':
>                 return chars
>             return None
>         return None
>     
> 
>     @F.udf(returnType=T.ArrayType(elementType=T.StringType()))
>     def udf2(col):
>         items = [k for k in ''.join(col)]
>         return items
>     
> 
>     @F.udf(returnType=T.StringType())
>     def udf3(col):
>         return broarcasted_variable.value.get(col, None)
>     
> 
>     df = df.select('text')\
>            .withColumn('text', udf(F.col('text')))\
>            .dropna()\
>            .withColumn("text", F.collect_list("text"))\
>            .withColumn('text', udf2(F.col('text')))\
>            .withColumn('text', F.explode("text"))\
>            .groupby('text')\
>            .count()\
>            .orderBy(["count",'text'], ascending=[0, 1])\
>            .withColumn('sentiment', udf3(F.col("text")))\
>     
>     positive_emojis = df.where(F.col('sentiment') == 'positive')\
>                         .select(F.col('text').alias('emoji'), F.col('count'))
>     neutral_emojis = df.where(F.col('sentiment') == 'neutral')\
>                         .select(F.col('text').alias('emoji'), F.col('count'))
>     negative_emojis = df.where(F.col('sentiment') == 'negative')\
>                         .select(F.col('text').alias('emoji'), F.col('count'))
> 
>     return (positive_emojis, neutral_emojis, negative_emojis)
> ```
> </details>

## 5. Probable fans (4 points)
```python
def task_5(df: "pyspark.sql.dataframe.DataFrame",
           F: "pyspark.sql.functions",
           W: "pyspark.sql.window.Window",
           top_n_likers: int) -> "pyspark.sql.dataframe.DataFrame":
    # Place your code to transform DaraFrame here
    
    
    modified_df = df.groupBy(["ownerId", "likerId"])\
                    .agg(F.count("ownerId").name("count"))
    
    # removing self likes
    modified_df = modified_df.where(modified_df.ownerId != modified_df.likerId)
    
    # using window function
    windowSpec  = W.partitionBy("ownerId")\
                    .orderBy(
                        F.col("ownerId").asc(), 
                        F.col("count").desc(), 
                        F.col("likerId").asc()
    )
    
    modified_df = modified_df.withColumn("row_number", F.row_number().over(windowSpec))
    
    modified_df = modified_df.where(modified_df.row_number <= top_n_likers)
    
    modified_df = modified_df.drop("row_number")
    
    modified_df = modified_df.orderBy(
                        F.col("ownerId").asc(),
                        F.col("count").desc(),
                        F.col("likerId").asc()
    )
    
    return modified_df
```
## 6. Probable friends (5 points)
```python
def task_6(df: "pyspark.sql.dataframe.DataFrame",
           F: "pyspark.sql.functions",
           W: "pyspark.sql.window.Window") -> "pyspark.sql.dataframe.DataFrame":
    # Place your code to transform DaraFrame here    
    userA = df.groupBy(
                    F.col("likerId"),
                    F.col("ownerId"))\
                .agg(F.count("ownerId").name("count"))

    # Removing self likes
    userA = userA.where(userA.ownerId != userA.likerId)

    # SEARCH LIKES
    window = W.partitionBy("likerId").orderBy(
            # F.col("ownerId").asc(),
            F.col("count").desc(),
            F.col("likerId").asc()
    )


    userA = userA.withColumn("max_likes", F.max("count").over(window))\
                 .where(F.col("count") == F.col("max_likes"))\
                 .drop("max_likes") 
    # Left only top likers (with "count" == topG count within "ownerId")
    
    userB = userA.alias("userB")
    
    modified_df = userA.alias("userA").join(
                                            userB, \
                                            (F.col("userA.ownerId") == F.col("userB.likerId")) &
                                            (F.col("userA.likerId") == F.col("userB.ownerId")) &
                                            (F.col("userA.ownerId") < F.col("userB.ownerId")),
                                            "inner")\
                                        .select(
                                            F.col("userA.ownerId").name("user_a"),
                                            F.col("userB.ownerId").name("user_b"),
                                            F.col("userB.count").name("likes_from_a"),
                                            F.col("userA.count").name("likes_from_b")
                                        )

    modified_df = modified_df.withColumn("mutual_likes", modified_df.likes_from_b + modified_df.likes_from_a)

    # Removing identical (user_a = 123, user_b = 456 and user_a = 456, user_b = 123) entries
    modified_df = modified_df.withColumn(
        "sorted_ids",
        F.concat(
            F.least(F.col("user_a"), F.col("user_b")),
            F.lit("_"),
            F.greatest(F.col("user_a"), F.col("user_b"))
        )
    )

    modified_df = modified_df.dropDuplicates(["sorted_ids"])

    modified_df = modified_df.select(
                                F.col("user_a"), 
                                F.col("user_b"), 
                                F.col("likes_from_a"), 
                                F.col("likes_from_b"), 
                                F.col("mutual_likes")
    )

    modified_df = modified_df.sort(F.col("mutual_likes").desc(), F.col("user_a").asc(), F.col("user_b").asc())
    
    return modified_df
```

> <picture>
>   <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/light-theme/success.svg">
>   <img alt="Success" src="https://raw.githubusercontent.com/Mqxx/GitHub-Markdown/main/blockquotes/badge/dark-theme/success.svg">
> </picture><br>
>
> You're Genius 🗿