The best source of truth for running tests (custom) and oci-conformance tests is
to look at the latest [GitHub workflow](https://github.com/vsoch/django-oci/blob/master/.github/workflows/main.yml#L33)
that runs them. We will also provide a brief example here.

First, clone django-oci

```bash
$ git clone https://github.com/vsoch/django-oci
$ cd django-oci
```

Typically you'd be installing Django-OCI into your Django project. For the purposes
of testing, we are going to use the example project. Note that you'll need to either
provide your own conformance binary, or use the one provided in the repository.

## Python Environment

Create a Python environment to install dependencies to.

```bash
$ python -m venv env
$ source env/bin/activate
$ pip install -r tests/requirements.txt
```

At this point in the actions we would be downloading the opencontainers/distributon-spec.
It's assumed you have installed Go.

```bash
$ git clone https://github.com/opencontainers/distribution-spec dist-spec
$ cd dist-spec/conformance
$ go mod vendor
$ CGO_ENABLED=0 go test -c -o ../../conformance.test
cd ../../
```

And then you can run two sets of tests. The first are Django based:

```bash
python manage.py makemigrations django_oci
python manage.py makemigrations
python manage.py migrate
python manage.py migrate django_oci
python manage.py test tests.test_api
rm db-test.sqlite3
```

It's important to remove the database at the end so we can start fresh. Next run
the conformance tests.

```bash
python manage.py makemigrations django_oci
python manage.py makemigrations
python manage.py migrate
python manage.py migrate django_oci
DISABLE_AUTHENTICATION=yes python manage.py test tests.test_conformance
```
Note that there is one failing test because a 404 is erroneously returned on an attempt to get mounted blob.
This was some change in the conformance testing that I haven't been able to fix (right before the release it worked okay!)

```bash

Summarizing 1 Failure:

[Fail] OCI Distribution Conformance Tests Push Cross-Repository Blob Mount [It] GET request to test digest within cross-mount namespace should return 200 
/tmp/django-oci/dist-spec/conformance/02_push_test.go:242

Ran 60 of 65 Specs in 2.332 seconds
FAIL! -- 59 Passed | 1 Failed | 0 Pending | 5 Skipped
```

It shouldn't be an issue if you disable this functionality. This was a side project
and I cannot work on it full time so [please open an issue if you can help](https://github.com/vsoch/django-oci/issues)!
