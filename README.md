# Memcached Operator but with custom instrumentation

## The code here pretty much follows [Golang Based Operator Quickstart](https://sdk.operatorframework.io/docs/golang/quickstart/) with 2 small changes:
1. The operator name is customized - some simple tweaks are needed in the code to get this done. We will not go into the details of this here. 
2. The main purpose it show how to make the operator generate custom metrics. The inspiration for that is from [this](https://github.com/brancz/prometheus-example-app) example

## How to to generate custom metrics from the Operator
Though [Expose custom Metrics](https://sdk.operatorframework.io/docs/golang/monitoring/prometheus/#expose-custom-metrics) gives us some directives, this is not the recommended method anymore. The main issues is around exploiting Global Registry for adding custom metrics. This is not recommended anymore. To read more about this -  [Kubernetes advises not to use Global Registry for metrics](https://github.com/kubernetes/enhancements/blob/master/keps/sig-instrumentation/20181106-kubernetes-metrics-overhaul.md).

Therefore, in this example, under `pkg/custommetrics` - you will see that we:
1. create a new Registry and 
2. register the metrics to it
3. create a http handler to serve these metrics
4. And finally, we initialize this from `cmd/manager/main.go`


## How to visualize the metrics
The basic operator exposes metrics in 2 ports as you will see in `cmd/manager/main.go`
```
var (
	metricsHost               = "0.0.0.0"
	metricsPort         int32 = 8383
	operatorMetricsPort int32 = 8686
)
```
The basic operator creates:
1. kubernetes service and 
2. ServiceMonitor Custom Resource in the namespace in which the operator POD runs. 

However this exposes the `standard operator metrics`. But for the custom metrics, we need a few extra steps
- we need to modify the service created above to exposes the custom metric port explicitly. 
- we also need to modify the service monitor created above to look for the custom metric port.

Once ServiceMonitor CRs are created, there are known ways to visualize these metris in Prometheus/Grafana.

## How to just deploy and play
1. The operator image is available here: quay.io/bjoydeep/memcached-operator-inst:002 and the deploy/operator.yaml is pointing to it.
2. kubectl create -f deploy/crd/cache.example.com_memcachedinsts_crd.yaml
3. kubectl create -f deploy/service_account.yaml
4. kubectl create -f deploy/role.yaml
5. kubectl create -f deploy/role_binding.yaml
6. kubectl create -f deploy/operator.yaml
7. Modify the service and ServiceMonitor as explained above - and you will be set.