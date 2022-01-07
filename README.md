# RunescapePathfinding
A grpc pathfinding service for oldschool runescape. There is also a http / json api exposed.

The runescape map is a sparse 5000x15000x4 grid with over 10 million nodes. This service uses contraction hierarchies to solve paths in ~200 ms. There is a client library for epicbot which can be found on the forums. With that being said, the service is designed to be used with any client. 

To get started, make a post request to /v1/get_path. There is a test server set up at http://ec2-54-244-194-50.us-west-2.compute.amazonaws.com:8081/. This is just a free ec2 instance so it will definately fail under extreme stress and has no guarantee of stability or uptime. With that being said, it has handled hundreds of requests a minute in the past without difficulty. 

```
curl -X POST -d '{"x_start":2965, "y_start":3380, "x_end":2965, "y_end": 3490}' http://ec2-54-244-194-50.us-west-2.compute.amazonaws.com:8081/v1/get_path
```

If you are a client owner who is interested in hosting their own server or implementing their own client library to use this service just send me a dm. 
