# ConfigMap & Secret replication for Kubernetes

**This repository is a fork of https://github.com/mittwald/kubernetes-replicator**

This repository contains a custom Kubernetes controller that can be used to make
secrets and config maps available in multiple namespaces.

## Deployment

```shellsession
$ # Create roles and service accounts
$ kubectl apply -f  https://raw.githubusercontent.com/zerospam/kubernetes-replicator/master/deploy/00-rbac.yaml
$ # Create actual deployment
$ kubectl apply -f https://raw.githubusercontent.com/zerospam/kubernetes-replicator/master/deploy/01-deployment.yaml
```

## Usage

- If a secret or configMap needs to be replicated to other namespaces, annotations should be added in that object permitting replication.
  - Add `replicator.v1.mittwald.de/replication-allowed-namespaces` annotation. Value of this annotation should contain a comma separated list or permitted namespaces or regular expressions. for e.g. `namespace-1,my-ns-2,app-ns-[0-9]*`, in this case replication will be performed only names `namespace-1`, `my-ns-2` and any namespace that matches the regular expression `app-ns-[0-9]*`.

    ```yaml
    apiVersion: v1
    kind: Secret
    metadata:
      annotations:
        replicator.v1.mittwald.de/replicate-allowed-namespaces: "my-ns-1,namespace-[0-9]*"
    data:
      key1: <value>
    ```

- Add the annotation `replicator.v1.mittwald.de/replicate-from` to any Kubernetes secret or config map object. The value of that annotation should contain the the name of another secret or config map (using `<namespace>/<name>` notation).

  ```yaml
  apiVersion: v1
  kind: Secret
  metadata:
    annotations:
      replicator.v1.mittwald.de/replicate-from: default/some-secret
  data: {}
  ```

  The replicator will then copy the `data` attribute of the referenced object into the annotated object and keep them in sync.   
