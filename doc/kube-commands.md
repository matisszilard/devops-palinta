# Palinta K8S commands

```sh
# Create
kubectl create deployment influxdb --image=docker.io/influxdb:1.6.4

# Get information about the deployment

kubectl get deployments
kubectl describe deployment influxdb
```

```sh
# Configure kubernetes secrets on influxdb

kubectl create secret generic influxdb-creds \
  --from-literal=INFLUXDB_DATABASE=twittergraph \
  --from-literal=INFLUXDB_USERNAME=root \
  --from-literal=INFLUXDB_PASSWORD=root \
  --from-literal=INFLUXDB_HOST=influxdb

kubectl get secret influxdb-creds
kubectl describe secret influxdb-creds
```

```sh
kubectl edit deployment influxdb

spec:
  template:
    spec:
      containers:
      - image: docker.io/influxdb:1.6.4
        imagePullPolicy: IfNotPresent
        name: influxdb

spec:
  containers:
  - name: influxdb
    envFrom:
    - secretRef:
        name: influxdb-creds
```

```sh
kubectl describe deployment influxdb
```

## Create persistent volume for influxdb

```yml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  labels:
    app: influxdb
    project: twittergraph
  name: influxdb
spec:
  accessModes:
  - ReadWriteOnce
  resources:
    requests:
      storage: 2Gi
```

```sh
kubectl create -f pvc.yaml
kubectl get pvc

```

```sh
kubectl expose deployment influxdb --port=8086 --target-port=8086 --protocol=TCP --type=ClusterIP
```

## Create config map from terminal

```
kubectl create configmap grafana-config \
  --from-file=datasource.yml=$DEVOPS/devops-palinta/kube/configmap/data_source.yaml \
  --from-file=dashboard-provider.yml=$DEVOPS/devops-palinta/kube/configmap/grafana-dashboard-provider.yml \
  --from-file=dashboard.json=$DEVOPS/devops-palinta/kube/configmap/dashboard.json
```