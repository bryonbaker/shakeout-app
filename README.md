# shakeout-app

## Overview

`shakeout-app` is a lightweight, containerized Go application designed for validating routing and persistent storage on OpenShift clusters, including single-node setups. It acts as a simple HTTP echo server with persistent logging: every access to `/hello` is logged (with timestamp and pod name) to a file on a mounted volume, and the log is displayed back in each response. Its goal is to help users confirm network, storage, and orchestration functionality in OpenShift environments.

---

## Features

- **Route Validation:** Easily verify that OpenShift ingress and routing are working.
- **Persistent Logging:** Writes request logs to a persistent volume so you can confirm storage attachment and data persistence across pod restarts.
- **On-demand Log Viewer:** The `/hello` endpoint returns the full content of the access log so you can observe activity and test cluster scaling.
- **Minimal Container Image:** Uses Red Hat UBI minimal, keeping runtime small, fast, and secure.
- **Easy Deployment:** Includes manifests and a Makefile for rapid setup and teardown.

---

## Repository Structure

- `shakeout.go` &mdash; Main Go application source code.
- `Containerfile` &mdash; Dockerfile/Containerfile for building the container image.
- `Makefile` &mdash; Automates building, pushing, and cleaning up the container image.
- `manifest/aws/shakeout-app-dep.yaml` &mdash; Deployment manifest for AWS/gp2 persistent volumes.
- `manifest/lvm/shakeout-app-dep.yaml` &mdash; Deployment manifest for clusters using LVM provisioner.
- `README.md` &mdash; Project documentation (this file).

---

## Prerequisites

- **OpenShift**: Single-node or multi-node cluster.
- **Podman** or **Docker**: For building and pushing container images.
- **OpenShift CLI (`oc`)**
- **Go**: For local development (not required for running the container).
- **Access to a container registry** (e.g., quay.io, Docker Hub).

---

## Building the Container Image

1. **Clone the repository:**
    ```
    git clone <repo-url>
    cd shakeout-app
    ```

2. **Update the Makefile** (if needed):
   Customize `IMAGE_NAME` and `TAG` to match your container registry and preferred tag.
   ```
   IMAGE_NAME := <your-registry>/<your-namespace>/shakeout-app
   TAG := latest
   ```

3. **Build the image:**
    ```
    make clean
    make build
    ```

4. **Push the image to your registry:**
    ```
    make push
    ```

---

## Deploying on OpenShift

1. **Edit the deployment manifest** (`manifest/aws/shakeout-app-dep.yaml` or `manifest/lvm/shakeout-app-dep.yaml`)
   - Update the `image:` field under `spec.containers` to your registry and tag.
   - Adjust the `storageClassName` in the PersistentVolumeClaim to suit your environment (`gp2`, `lvm`, etc.).

2. **Apply the manifest:**
    ```
    oc apply -f manifest/aws/shakeout-app-dep.yaml
    # or
    oc apply -f manifest/lvm/shakeout-app-dep.yaml
    ```

---

## Accessing the Application

1. **Get the exposed route:**
    ```
    oc get routes
    ```

2. **Visit the `/hello` endpoint in your browser:**
    ```
    http://<route-host>/hello
    ```
   - Each request writes a log entry and displays the full access log.

---

## Customization

- **Change Container Image Name:**
  Update both the `Makefile` and deployment manifest to reference your preferred registry and repo.
- **Change Storage Class:**
  Modify the `storageClassName` field in the manifest to match your OpenShift cluster's provisioner.
- **Log Path or Format:**
  Change the `logFilePath` or output format in `shakeout.go`.

---

## File Breakdown

- **`shakeout.go`:**
    - HTTP server on port 9000
    - Handles `/hello` route
    - Logs request timestamp and pod name to `/data/access-log.log`
    - Responds with HTML including all previous log entries

- **`Containerfile`:**
    - Based on Red Hat UBI minimal image
    - Installs Go toolchain, builds the binary, exposes port 9000

- **`manifest/aws/shakeout-app-dep.yaml` & `manifest/lvm/shakeout-app-dep.yaml`:**
    - OpenShift deployments with appropriate PVC and route setup for different storage classes

- **`Makefile`:**
    - Automates container build, push, and cleanup via Podman

---

## Maintenance

- This image and manifest are for demonstration and testing only.
- Periodically update the base image and dependencies for security.
- Does not implement authentication or input validation; not suitable for production or public exposure.
- Logs can grow without limit; add log rotation or persistent log retention if necessary for long-running use.

---

## License

MIT License
