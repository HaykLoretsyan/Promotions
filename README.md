**_Promotions_ is a repository that contains all the necessary code for /promotions/:id endpoint to function.**

### Requirements
You should have mongoDB installed on your machine and the following environment variables specified:

```
DATABASE_URI "URI for mongoDB service"
CSV_FILE_PATH "Path to csv file"   
```

### Running
After meeting requirements you can just do `go run .` in the root directory of the repo.

For retrieving an entry from the endpoint just run `curl <hostname>:8000/promotions/<id>`.

### Additional Requirements
1. ***The .csv file is very big (billions of entries)***. In this case multiple things can be done. We can read csv file from multiple goroutines
by chunks, we can delegate update part to worker services, and we can scale mongodb to be able to handle big chunks of data
faster.
2. ***How would this application perform in peak periods (millions of requests per
   minute)***. It depends on the machine that it is running on, but it is a good practice to have load balancer with multiple instances of this app 
running behind it in case of large amount of requests. 
3. ***How would you operate this app in production (e.g. deployment, scaling, monitoring)***. It depends on the amount of data/users that this app 
is going to handle. We can use docker-compose or kubernetes for deploying/scaling this app according to the requirements. For monitoring system state 
we can use well known monitoring services e.g. Grafana/Prometheus, Kibana/elasticsearch, etc.