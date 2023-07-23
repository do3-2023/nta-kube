# ☸️ Orchestration Exercise

This repository contains the code for an assignment as part of the course Orchestration at Polytech Montpellier.

## Table of contents

1. [Instructions](#instructions)
2. [The application](#the-application)
3. [Architecture](#architecture)
4. [Install guide (Taskfile)](#install-guide-taskfile)
5. [Install guide (manual)](#install-guide-manual)
6. [Development](#development)
7. [Known issues](#known-issues)
8. [License](#license)

## Instructions

<details>
  <summary>Click to see the exercise instructions</summary>
  <img src="https://github.com/do3-2023/nta-kube/assets/78071629/6b6e390e-a904-448e-a42b-65c8d692e9ca"/>
</details>

## The application

This application simply consists of a list of drinks and a button to add a new random drink to the list. The list of drinks is stored in a PostgreSQL database and the API is used to retrieve and add drinks to the database.

<details>
  <summary>Click to see a demo</summary>
  <img src="https://github.com/do3-2023/nta-kube/assets/78071629/51da1226-2f90-4b84-b20a-d231ad8b4b59"/>
</details>

## Architecture

### Applications

- The web application has been made using [SvelteKit](https://kit.svelte.dev/) as it allows users to create SSR only routes which is required for this assignment.
    - Latest image: `ghcr.io/do3-2023/nta-kube/webapp:main`
    - Source code: [webapp](webapp)
- The API has been built using [Go](https://golang.org/) and the API library named [go-chi](https://go-chi.io).
    - Latest image: `ghcr.io/do3-2023/nta-kube/api:main`
    - Source code: [api](api)
- The database is a simple [PostgreSQL](https://www.postgresql.org/) database.
- Each application is built and pushed as an image to [Github Container Registry (ghcr.io)](https://github.com/features/packages) using [Github Actions](https://github.com/features/actions).
    - Source code: [workflows](.github/workflows)

### Infrastructure

- Each application has its own namespace as required by the assignment.
- Since each application is deployed in its own namespace, they can't share the same Secrets or ConfigMaps. Consequently, each namespace has its own Secret and ConfigMap with identical values.
- The applications communicate with each other using the cluster's internal DNS. This means that they use the following URLs: `<app-name>.<namespace>.svc.cluster.local`.
- The database uses a PersistentVolumeClaim (PVC) to store its data. This PVC is backed by a PersistentVolume (PV) which resides on a local volume on the host machine. In this case, the volume is stored in the /tmp/nta-kube directory.
- The web application doesn't have an Ingress and needs to be port forwarded to be accessed.
- The API has a replica count of 3 to demonstrate that it can be scaled.

![k8s infrastructure](https://github.com/do3-2023/nta-kube/assets/78071629/cf7ddb33-4986-4b06-abb8-e494a8f5dad0)

## Install guide (Taskfile)

This project features a [Taskfile](https://taskfile.dev/#/) which automates the creation of the cluster and the deployment of the applications. You can find out about all the available tasks by running `task --list`.

### Requirements

- [kubectl](https://kubernetes.io/docs/tasks/tools/)
- [k3d](https://k3d.io/#installation)
- [Taskfile](https://taskfile.dev/#/installation)

### Steps

1. Create the k3d cluster with everything deployed and ready to go.
```bash
task deploy
```

2. Port forward the web application to your local machine. This should make the web application available [here](http://localhost:8080). You may need to wait a few dozen seconds for the web application to be ready.
```bash
task port-forward
```

3. Delete the cluster (when you're done).
```bash
task delete
```

## Install guide (manual)

### Requirements

- [kubectl](https://kubernetes.io/docs/tasks/tools/)
- [k3d](https://k3d.io/#installation)

### Steps

1. Create a k3d cluster named `nta-kube`.
```bash
k3d cluster create nta-kube
```

2. Create the namespaces for the applications. This has to be done before deploying the applications.
```bash
kubectl apply -f infra/namespaces
```

3. Deploy the PostgreSQL database.
```bash
kubectl apply -f infra/db
```

4. Deploy the API.
```bash
kubectl apply -f infra/api
```

5. Deploy the web application.
```bash
kubectl apply -f infra/webapp
```

6. Port forward the web application to your local machine. This should make the web application available [here](http://localhost:8080). You may need to wait for the web application to be ready. It can take up to a minute.
```bash
kubectl port-forward -n webapp svc/webapp 8080:3000
```

## Development

A docker-compose file is available to run the applications locally. This docker-compose file doesn't include the web application which needs to be run separately using `npm run dev`.

- Using the Taskfile:
```bash
task dev
```

- Without the task command:
```bash
docker-compose up -d
```

## Known issues

- The web application sometimes has trouble loading and fetching data on the first load. This issue seems to happen only on Firefox. A simple refresh fixes the issue.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.