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
$ git clone -b v1.0.1 https://github.com/opencontainers/distribution-spec dist-spec
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

This will generate the report and junit.xml files in the present working directory.
The files present in this repository were run with the distribution-spec, version 1.0.1.
