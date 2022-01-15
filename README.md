# airport_tp

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

3) To launch api :
```shell
go build
./api
```
4) 
