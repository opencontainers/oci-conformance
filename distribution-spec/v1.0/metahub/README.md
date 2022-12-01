# MetaHub Registry

![MetaHub Logo](https://gitlab.com/qnib-metahub/community-edition/-/raw/beta/img/metahub-logo.png)

MetaHub is a registry that takes the context of a user into account to deliver the best suited container image.

## Getting Started

To tip your toe into it let's download a generic image for `qnib/featuretest`:

```bash
$ echo hub |docker login metahub-registry.org --password-stdin  -u meta
Login Succeeded
$ docker run -ti --rm --pull=always metahub-registry.org/qnib/featuretest:latest
latest: Pulling from qnib/featuretest
88ecf269dec3: Pull complete
a2f2473b6a7e: Pull complete
Digest: sha256:fa4073c84baeda1fa38e12f7d0675476b435ebb5e70048ff7bc728aa3ab2dd9b
Status: Downloaded newer image for metahub-registry.org/qnib/featuretest:latest
>> This container is optimized for: amd64
```

Next we use the profile `zen3` to pull down a new image optimized for the particular hardware architecutre.

```bash
$ echo hub |docker login metahub-registry.org --password-stdin  -u meta
Login Succeeded
$ docker run -ti --rm --pull=always metahub-registry.org/qnib/featuretest:latest
latest: Pulling from qnib/featuretest
59bf1c3509f3: Pull complete
c1ac79912c87: Pull complete
Digest: sha256:14e67ca906c711fb5283b7241765a59827d55a8259af851cdeb3c822140fc10e
Status: Downloaded newer image for metahub-registry.org/qnib/featuretest:latest
>> This container is optimized for: arch:zen3
```
