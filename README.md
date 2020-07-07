# Setec Astronomy

![CI](https://github.com/anthonysterling/setec/workflows/CI/badge.svg)

_Setec Astronomy_ is an anagram of _too many secrets_ which I stole from the excellent movie [Sneakers](https://www.rottentomatoes.com/m/sneakers), which stars Robert Redford, Dan Aykroyd, Ben Kingsley, Mary McDonnell, River Phoenix, Sidney Poitier, and David Strathairn! 🤩

Go watch it.

## Overview

Setec (_pronounced see-tek_) is a utility tool that encrypts and decrypts secrets that are managed by Bitnami's [Sealed Secrets](https://github.com/bitnami-labs/sealed-secrets). Whilst we're technically not meant to be doing this, I had a use case and wanted to share this.

### Obtaining Sealed Secrets Certificate and Key

The tool requires the Sealed Secrets key to decrypt a value, and the Sealed Secrets certificate to encrypt a value. Where these are located in your Kubernetes cluster is most likely something you know already, I found mine with these commands; which may help.

```
kubectl get secrets --namespace kube-system --field-selector type=kubernetes.io/tls -o json | jq -r '.items[].data."tls.crt"' | base64 -D
```

```
kubectl get secrets --namespace kube-system --field-selector type=kubernetes.io/tls -o json | jq -r '.items[].data."tls.key"' | base64 -D
```

### Usage

Sealed Secrets are, optionally, scoped by Kubernetes namespace and name. If a Sealed Secret was scoped as cluster-wide you can omit the `--namespace` and `--name` flags.


```
cat plain-secret.txt | setec encrypt --public-key-path /tmp/backup.pub --namespace production --name rails
```

```
cat encrypted-secret.txt | setec decrypt --private-key-path /tmp/backup.key --namespace production --name rails
```

## Contributing

Contributions to this project are [released](https://help.github.com/articles/github-terms-of-service/#6-contributions-under-repository-license) to the public under the [MIT license](LICENSE.md).

### Submitting a pull request

0. Fork and clone the repository
0. Make sure the test and build succeed on your machine: `script/test` and `script/build`
0. Create a new branch: `git checkout -b my-branch-name`
0. Make your change, add tests, and make sure the tests still pass
0. Push to your fork and submit your pull request
