# profile for server start up

set up boost configuration for micro runtime which means using different components of go-micro

- local

  - registry mdns
  - config file
  - broker http
  - store memory

- k8s

  - registry kubernetes
  - config configmap
  - router static :dns

- dev
  - registry mdns
  - config environment
  - broker http
  - store memory
