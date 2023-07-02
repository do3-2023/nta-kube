# nta-kube

## Web App

### Docker

```bash
docker build -t nta-kube-webapp:latest .
docker run -p 3000:3000 nta-kube-webapp:latest
```

### Kubernetes

```bash
kubectl apply -f webapp/infra
kubectl port-forward deployment/nta-kube-webapp 8080:3000
```
