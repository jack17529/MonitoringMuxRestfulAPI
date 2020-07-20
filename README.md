# MonitoringMuxRestfulAPI

Implemented GET, POST, DELETE and PUT services on a books database using Mux router and Go. Tested it using Postman and CURL. Monitored it using Prometheus and Grafana.

## API

1. GET request using CURL - `curl -X GET http://localhost:8000/api/books`  
Response - `[{"id":"1","isbn":"448743","title":"Book One","author":{"firstname":"John","lastname":"Doe"}},{"id":"2","isbn":"448744","title":"Book Two","author":{"firstname":"Steve","lastname":"Smith"}}]`

2. GET request using CURL - `curl -X GET http://localhost:8000/api/books/1` 
Response - `{"id":"1","isbn":"448743","title":"Book One","author":{"firstname":"John","lastname":"Doe"}}`

3. POST request using CURL - `curl -g -X POST -H "Content-Type: application/json" -d '{"isbn":"4545454","title":"Book Three","author":{"firstname":"Harry",  "lastname":"White"}}' http://localhost:8000/api/books`  
Response - `{"id":"","isbn":"","title":"","author":null}`

4. DELETE request using CURL - `curl -X "DELETE" http://localhost:8000/api/books/1`  
Response - `[{"id":"2","isbn":"448744","title":"Book Two","author":{"firstname":"Steve","lastname":"Smith"}},{"id":"7131847","isbn":"4545454","title":"Book Three","author":{"firstname":"Harry","lastname":"White"}}]`

5. PUT request using CURL - `curl -X PUT -H "Content-Type: application/json" -d '{"isbn":"2121212","title":"Updated Title","author":{"firstname":"Charles",  "lastname":"Dickens"}}' http://localhost:8000/api/books/2`  
Response - `{"id":"2","isbn":"2121212","title":"Updated Title","author":{"firstname":"Charles","lastname":"Dickens"}}`

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

![alt text]()
![alt text]()

## References

1. https://www.keycdn.com/support/popular-curl-examples#:~:text=13.-,GET%20method,request%20GET%20or%20%2DX%20GET%20.
2. https://github.com/rusart/muxprom
3. https://www.youtube.com/watch?v=SonwZ6MF5BE
4. https://dzone.com/articles/go-microservices-part-15-monitoring-with-prometheu
5. https://banzaicloud.com/blog/monitoring-gin-with-prometheus/
6. https://prometheus.io/docs/prometheus/latest/getting_started/
7. https://medium.com/htc-research-engineering-blog/build-a-monitoring-dashboard-by-prometheus-grafana-741a7d949ec2
