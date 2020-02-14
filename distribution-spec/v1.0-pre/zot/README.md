zot
===

# Install
```
go get -u github.com/anuvu/zot/cmd/zot
```

# Setup 

## Conformance Tests

```
git clone https://github.com/opencontainers/${spec_name}.git
cd ${spec_name}/conformance
go test -c
```

## Without BASIC auth

```
cat << EOF > no-auth.json
{
  "version":"0.1.0-dev",
  "storage":{
    "rootDirectory":"/tmp/zot"
  },
  "http": {
    "address":"127.0.0.1",
    "port":"8080"
  },
  "log":{
    "level":"debug"
  }
}
EOF
```

```
zot serve no-auth.json
```

```
OCI_ROOT_URL="http://localhost:8080" OCI_NAMESPACE="test/repo" ./conformance.test
```

## With BASIC auth

```
cat << EOF > auth.json
{
  "version":"0.1.0-dev",
  "storage":{
    "rootDirectory":"/tmp/zot"
  },
  "http": {
    "address":"127.0.0.1",
    "port":"8080",
    "realm":"zot",
    "auth": {
      "htpasswd": {
        "path": "htpasswd"
      }
    }
  },
  "log":{
    "level":"debug"
  }
}
EOF
```

```
htpasswd -bBn test test123 > htpasswd 
```
```
zot serve auth.json
```

```
OCI_ROOT_URL="http://localhost:8080" OCI_NAMESPACE="test/repo" OCI_USERNAME="test" OCI_PASSWORD="test123" ./conformance.test
```

# References

[zot](https://github.com/anuvu/zot)
