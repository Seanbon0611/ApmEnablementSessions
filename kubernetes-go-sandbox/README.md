# Distributed Tracing

## 1. Install the dependecies
- cd into GinApp1 and run go get to install the dependencies:
```
cd /apm-enablement-sessions/kubernetes-go-sandbox/GinApp1
go get
```
- cd into GinApp2 and run the same command:
```
cd /apm-enablement-sessions/kubernetes-go-sandbox/GinApp2
go get
```

## 2. Run The applications
- within the GinApp1 directory run:
```
go run main.go
```
- Open up a new terminal, access the GinApp2 directory and run:
```
cd /apm-enablement-sessions/kubernetes-go-sandbox/GinApp2
go run main.go
```

## 3. Generate Traffic
curl into the endpoint:
```
curl localhost:8080/test
```