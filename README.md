# ocp-s2i

A simple go appliation to play with a number of concepts, including:

- multiple go modules in a single git repo, using go.work file: https://github.com/golang/tools/blob/master/gopls/doc/workspace.md
- OCP s2i to build a simple go app in OCP, and deploy it.
  - Dockerfile based approach, so the same app can be built identically locally with podman, and by s2i in OCP.
- OCP Service env vars - utilise the OCP namespace scoped service variables to call from one app to another

## Tasks

- [x] create and run simple hello http server in go
- [x] Build it from Dockerfile
- [x] Build and run image locally with podman
- [x] Build and run from Dockerfile in OCP using docker context with s2i
- [x] build and deploy second identical app to same OCP namespace
- [ ] confirm service variables exist in each pod's container, and can be used to call other app, even after restarts...
- [ ] [optional] update app2 to actually call app1 using the above env var techniques


## Commands

### Run and test app1:
```
cd hello-app1
go run hello.go
```
In another terminal
```
curl http://localhost:8080/hello
```

### To build and run app1
```
cd hello-app1
go build hello.go
./hello
```
In another terminal
```
curl http://localhost:8080/hello
```

### Build from Containerfile

```
cd hello-app1
podman build -t hello-app1 .
```

### Run image

```
podman run -p 8080:8080 --name hello-app1 -d hello-app1
podman ps
curl localhost:8080/hello
podman stop hello-app1
```

### Build app in OCP from Dockerfile
```
oc new-app --context-dir=hello-app1 --strategy=docker --name=hello-app1 https://github.com/thisisdavidbell/ocp-s2i
oc status
oc get pods
```

**Notes**: 
  - By default, all resources take the git repo name, if `--name=...` omitted
  - If the Dockerfile has no port `EXPOSE`d, then no service is created. In this case delete everything and re-run `oc new-app ...`

### Delete all deployed objects

To delete all, including the image stream and build config:

```
oc get all
oc delete all --selector app=hello-app1
oc get all
```

**Note**: to rebuild and redeploy, you now run the full `oc new-app ...` again...

### To rebuild and redeploy in place, without deletion

```
# make changes and push to github

oc start-build hello-app1
oc get pods -w
```

**Note**: you can see it run the build, and if successful start a new pod for the newly build hello-app1, and terminate the original hello-app1 pod.

### Expose and call app
```
oc get all
oc get service
oc get route
oc expose service hello-app1
oc get route
curl $(oc get route -o json hello-app1 | jq -r .spec.host)/hello
```

### deploy a second identical app, with different name
```
oc new-app --context-dir=hello-app1 --strategy=docker https://github.com/thisisdavidbell/ocp-s2i --name=hello-app2
oc get service
oc expose service hello-app2
curl $(oc get route -o json hello-app2 | jq -r .spec.host)/hello
```