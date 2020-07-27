apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: palinta-device
  name: palinta-device
spec:
  replicas: 1
  selector:
    matchLabels:
      app: palinta-device
  template:
    metadata:
      labels:
        app: palinta-device
    spec:
      containers:
        - image: mszg/palinta-device:v0.1.0
          name: palinta-device
          ports:
            - containerPort: 8080
              name: palinta-device
---
apiVersion: v1
kind: Service
metadata:
  name: palinta-device
  labels:
    app: palinta-device
    project: palinta
spec:
  type: NodePort
  selector:
      app: palinta-device
  ports:
    - port: 8080
      targetPort: 8080
      protocol: TCP
---
