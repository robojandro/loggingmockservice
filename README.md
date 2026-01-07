# Overview

*LoggingMockService* outputs a stream of log messages.

It can be executed by doing:
```
go run main.go
```

Example of the output:

![Simulated loggin as a service](example_output.png "example of output")

The ratios of how many lines logged at specific levels can be customized using the following *optional* arguments:

```
 -fatal_ratio  integer
 -error_ratio    "
 -warn_ratio     "
 -info_ratio     "
 -debug_ratio    "
 -trace_ratio    "
```

where those integers represent the ratio out of 100 percent that you want of a given level.

By default, *NO* Fatal or Trace level messages are logged and the distribution should approximate:

```
Error 10%
Warn  20%
Info  50%
Debug 20%
```

The outputs will not match that exactly as the output involves randomization, so there will be some deviation.

Additional arguments to tailor the output are:

```
-output integer
    default amount of log lines to output before exiting (default 1000)

-delay int
      delay between log messages in milliseconds (default 10)

-verbose
    whether to output execution statistics to STDOUT after log statements have been output
```

Setting the delay to 0 results in a near instant finish, setting it above 200 produces a more natural human readable rate of output.

### Other ways of executing

The LoggingMockService can be built as a docker container using the included `Dockerfile` and spun up as Kubernetes pod using the included `pod.yaml`.

The `Dockerfile` uses the following arguments for customize the output:
```
--delay 200 --output 20000 --verbose
```

Load the image so that kubectl can find it:
```
minikube image load logging-mock-service:latest
```

Apply the pod yaml through kubectl to get the logging service running:
```
kubectl apply -f pod.yaml
```

### See also:

Early proof-of-concept for tool that can stream and parse 
- https://github.com/robojandro/whaletail

Library used to generate the actual log output
- https://github.com/robojandro/loggenerator 


### Related:
- Kubernetes
- Minikube
- Docker
- Stern

