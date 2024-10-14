# Overview

This is a simple shakeout app to test routes and use of persistent volumes on Single Node OpenShift.

# Running on OpenSHift

Deploy the app with `oc apply -f shakeout-app-dep.yaml`

Then find the route and open  the route in a browser with the path `hello`.

E.g. 
```
$ oc get routes
NAME                 HOST/PORT                                                 PATH   SERVICES               PORT   TERMINATION   WILDCARD
shakeout-app-route   shakeout-app-route-shakeout.apps.kmarthub.bakerapps.net          shakeout-app-service   9000                 None
```

Then navigate to `http://shakeout-app-route-shakeout.apps.kmarthub.bakerapps.net/hello`


# Maintaining the base image patch level

I don't promise to keep this patched on the latest RHEL UBI. So to make this your own you need to maintain two specific files.

**Makefile:**  

Update this file to generate the image using the correct name and tag.
```
IMAGE_NAME := quay.io/bryonbaker/shakeout-app
TAG := latest
```

**shakeout-app-dep.yaml**

Update this file to use your image from your registry.
```
    spec:
      containers:
      - name: shakeout-app
        image: quay.io/bryonbaker/shakeout-app:latest
```

# Building the Image
To build this yourself run:  
```make clean``` to clean up any local image.  
```make build``` to build the new image.   
```make push``` to push the image to the target registry.   
