# grpc-health

## Description
CLI tool to health check GRPC services. The service is expected to implement the **Health** method as follows:
```protobuf
import "google/protobuf/empty.proto";

package com.example;

service SomeService {
  rpc Health (google.protobuf.Empty)  returns (google.protobuf.Empty) {}

  // Remaining methods
}
// Remaining stuff
```

## Installation
```
$ make
```

## Usage
```
$ grpc-health [OPTIONAL FLAGS] --serverUrl=<server_url> --serviceName=<fully_qualified_service_name>
```
For a more detailed usage, see
```
$ grpc-health --help
```

## Example
```
$ grpc-health --verbose --timeout=2 --serverUrl=localhost:8081 --serviceName=com.example.SomeService
2017/08/24 20:09:06.024428 Starting health check for com.example.SomeService running at localhost:8081 (timeout=2s)
2017/08/24 20:09:06.024522 Attempting connection...
2017/08/24 20:09:06.027936 Connection successful
2017/08/24 20:09:06.027952 Invoking /com.example.SomeService/Health
2017/08/24 20:09:06.028788 Invoke successful
2017/08/24 20:09:06.028803 Healthy
```
