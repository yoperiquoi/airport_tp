# Go project : airport sensor

School project for the first year of master diploma of IMT Atlantique.
Developing an application with go. Working with broker et IoT-simulated devices.

## To setup redis database

1) To launch the docker Image :
```shell
docker run -p 6379:6379 -it --rm redislabs/redistimeseries
```

2) Go into redis-cli of docker and create the timeseries :
```shell
TS.CREATE sensor:temperature:NTE LABELS airport_id NTE sensor_type temperature unit C
TS.CREATE sensor:wind:NTE LABELS airport_id NTE sensor_type wind unit KMH
TS.CREATE sensor:pressure:NTE LABELS airport_id NTE sensor_type pressure unit percent
```

## To launch services

1) Using go run / build method : 
```shell
go run xxx.go
```
or
```shell
go build
```
In the directory of each services because the program will search for config file in the location where it's launched

2) Using the go install method :
```shell
go install ./...
```
on the root of the project, will build all the binaries and resolve dependencies. All the binaries files are located in 
the GOROOT/bin folder 

To launch the binary take care to have the config file in the same location (for captors and subscribers), if not the program will stop.

## To launch the front end application

```shell
npm start
```
