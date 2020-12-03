# MonitoringMuxRestfulAPI

Implemented GET, POST, DELETE and PUT services on a books database using Mux router and Go. Tested it using Postman and CURL. Monitored it using Prometheus and Grafana.

## API

1. GET request using CURL - `curl http://localhost:8000/api/books -X GET | jq`  
Response -  
```
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   248  100   248    0     0   121k      0 --:--:-- --:--:-- --:--:--  242k
[
  {
    "id": 1,
    "isbn": "978-0812036381",
    "title": "Hamlet",
    "author": {
      "firstname": "William",
      "lastname": "Shakespeare"
    }
  },
  {
    "id": 2,
    "isbn": "978-0671027032",
    "title": "How to Win Friends & Influence People",
    "author": {
      "firstname": "Dale",
      "lastname": "Carnegie"
    }
  }
]
```

2. GET request using CURL - `curl http://localhost:8000/api/books/1 -X GET | jq`  
Response -  
```
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   108  100   108    0     0   105k      0 --:--:-- --:--:-- --:--:--  105k
{
  "id": 1,
  "isbn": "978-0812036381",
  "title": "Hamlet",
  "author": {
    "firstname": "William",
    "lastname": "Shakespeare"
  }
}
```

3. POST request using CURL - `curl http://localhost:8000/api/books -g -X POST -H "Content-Type: application/json" -d '{"isbn":"978-0060555665","title":"The Intelligent Investor","author":{"firstname":"Benjamin",  "lastname":"Graham"}}'`  

4. DELETE request using CURL - `curl http://localhost:8000/api/books/1 -X "DELETE"`  

5. PUT request using CURL - `curl http://localhost:8000/api/books/2 -X PUT -H "Content-Type: application/json" -d '{"isbn":"978-1503212831","title":"A Christmas Carol","author":{"firstname":"Charles",  "lastname":"Dickens"}}'`

## Docs Using Swagger And Hosting Using Redoc

1. `make swagger` to generate swagger.yaml file.
2. serving the swagger.yaml file to Redoc middlewares.
3. Getting the docs ready.

## Monitoring

1. Edit the api to get the matrices at `http://localhost:8000/metrics`
2. cd in the prometheus folder and change the `prometheus.yml` file accordingly.
3. cd to the prometheus download folder and `./prometheus --config.file=prometheus.yml` 
4. Check whether prometheus server is running and whether the targets are UP by going to - `http://localhost:9090/new/targets`
5. Start grafana 
```
sudo systemctl daemon-reload
sudo systemctl start grafana-server
sudo systemctl status grafana-server
```
6. Open `http://localhost:3000/`
7. Import prometheus as datasource in grafana
8. Import an already made dashboard.

![alt text](https://github.com/jack17529/MonitoringMuxRestfulAPI/blob/master/grafana/graph1.png)  
![alt text](https://github.com/jack17529/MonitoringMuxRestfulAPI/blob/master/grafana/graph2.png)

## References

1. https://www.keycdn.com/support/popular-curl-examples#:~:text=13.-,GET%20method,request%20GET%20or%20%2DX%20GET%20.
2. https://github.com/rusart/muxprom
3. https://www.youtube.com/watch?v=SonwZ6MF5BE
4. https://dzone.com/articles/go-microservices-part-15-monitoring-with-prometheu
5. https://banzaicloud.com/blog/monitoring-gin-with-prometheus/
6. https://prometheus.io/docs/prometheus/latest/getting_started/
7. https://medium.com/htc-research-engineering-blog/build-a-monitoring-dashboard-by-prometheus-grafana-741a7d949ec2
8. https://goswagger.io/generate/spec.html
