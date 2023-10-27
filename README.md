# RunescapePathfinding
A grpc pathfinding service for oldschool runescape. There is also a http / json api exposed.

The runescape map is a sparse 5000x15000x4 grid with over 10 million nodes. This service uses contraction hierarchies to solve paths in ~200 ms. There is a client library for epicbot and OSB which can be found on the forums. With that being said, the service is designed to be used with any client. 

To get started, make a post request to /v1/get_path. 

```
curl -X POST -d '{"x_start":2965, "y_start":3380, "x_end":2965, "y_end": 3490}' localhost
```

There is an integration with explv's map here. You may have issues with ssl as github sites default to using ssl and the back end does not have ssl keys. 

http://phobichd.github.io/explv/

If you are a client owner who is interested in hosting their own server or implementing their own client library to use this service just send me a dm. 

https://www.youtube.com/watch?v=2oq3lZxJRc8
