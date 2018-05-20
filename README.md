# ConfigMap & Secret replication for Kubernetes

**This repository is a fork of https://github.com/mittwald/kubernetes-replicator**

This repository contains a custom Kubernetes controller that can be used to make
secrets and config maps available in multiple namespaces.

## Deployment

```shellsession
$ # Create roles and service accounts
$ kubectl apply -f  https://github.intuit.com/raw/dev-build/kubernetes-replicator/master/deploy/rbac.yaml
$ # Create actual deployment
$ kubectl apply -f https://github.intuit.com/raw/dev-build/kubernetes-replicator/master/deploy/deployment.yaml
```

## Usage

### 1. Grant permission for replicator

Replicator must be granted role in both the source and destination namespace(s) to manage `ConfigMap` / `Secret` objects. If a secret needs to be replicated to more than one namespace, role should be granted in the source namespace and in all destination namespaces.

```shellsession
$ # Create role binding
$ kubectl apply -f  https://github.intuit.com/raw/dev-build/kubernetes-replicator/master/deploy/role-binding.yaml --namespace <namespace>
```

### 2. Create empty secret

Add the annotation `replicator.v1.mittwald.de/replicate-from` to any Kubernetes
secret or config map object. The value of that annotation should contain the
the name of another secret or config map (using `<namespace>/<name>` notation).

```yaml
apiVersion: v1
kind: Secret
metadata:
  annotations:
    replicator.v1.mittwald.de/replicate-from: default/some-secret
data: {}
```

The replicator will then copy the `data` attribute of the referenced object into
the annotated object and keep them in sync.   
