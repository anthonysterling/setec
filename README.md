# Setec Astronomy

_Setec Astronomy_ is an anagram of _too many secrets_ which I stole from the excellent movie [Sneakers](https://www.rottentomatoes.com/m/sneakers), starring Robert Redford, Dan Aykroyd, Ben Kingsley, Mary McDonnell, River Phoenix, Sidney Poitier, and David Strathairn! 🤩

Go watch it.

## Overview

Setec (_pronounced see-tek_) is a utility tool that encrypts and decrypts secrets that are managed by Bitnami's [Sealed Secrets](https://github.com/bitnami-labs/sealed-secrets).

### Obtaining Sealed Secrets Certificate and Key

```
kubectl get secrets --namespace kube-system --field-selector type=kubernetes.io/tls -o json | jq -r '.items[].data."tls.crt"' | base64 -D
```

```
kubectl get secrets --namespace kube-system --field-selector type=kubernetes.io/tls -o json | jq -r '.items[].data."tls.key"' | base64 -D
```

### Usage

Sealed Secrets are, optionally, scoped by Kubernetes namespace and name. If a Sealed Secret was scoped as cluster-wide you can omit the `--namespace` and `--name` flags.


```
cat plain-secret.txt | setec encrypt --public-key-path string /tmp/backup.pub --namespace production --name rails
```

```
cat encrypted-secret.txt | setec decrypt --private-key-path string /tmp/backup.key --namespace production --name rails
```
