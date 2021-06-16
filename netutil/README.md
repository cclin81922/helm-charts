# Docker
docker run --name server --hostname server --rm -d -e PROTOCOL=tcp -e PORT=5000 jimlintw/netutil:v20210616 /opt/server.sh
docker run --name client --hostname client --rm -d -e PROTOCOL=tcp -e PORT=5000 -e IP=172.17.0.3 jimlintw/netutil:v20210616 /opt/client.sh
docker logs -f server
docker logs -f client

# K8s
cat <<YAML | kubectl apply -f -
apiVersion: apps/v1
kind: Deployment
metadata:
  name: netutil-server
spec:
  replicas: 1
  selector:
    matchLabels:
      app: netutil
      component: server
  template:
    metadata:
      labels:
        app: netutil
        component: server
    spec:
      containers:
      - name: netutil
        image: jimlintw/netutil:v20210616
        env:
        - name: PROTOCOL
          value: tcp
        - name: PORT
          value: "30300"
        command: ["/opt/server.sh"]
---
apiVersion: v1
kind: Service
metadata:
  name: netutil-server
spec:
  type: NodePort
  selector:
    app: netutil
    component: server
  ports:
    - protocol: TCP
      port: 30300
      targetPort: 30300
      nodePort: 30300
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: netutil-client
spec:
  replicas: 1
  selector:
    matchLabels:
      app: netutil
      component: client
  template:
    metadata:
      labels:
        app: netutil
        component: client
    spec:
      containers:
      - name: netutil
        image: jimlintw/netutil:v20210616
        env:
        - name: PROTOCOL
          value: tcp
        - name: PORT
          value: "30300"
        - name: IP
          value: "172.17.0.3"
        command: ["/opt/client.sh"]
---

YAML
