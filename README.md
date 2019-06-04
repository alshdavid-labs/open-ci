# Open CI

*This is still in development, this readme functions as a description of features to expect*

Simple CI solution you can self host which gives you the power of pipeline systems like BitBucket and GitLab pipelines.
The self hosted service includes an HTTP server which you must run as a daemon on your inux based server. 
You need docker running on your host.

Basic auth is used to control access. 
dashboard is available on `/admin` and will allow you to set config like environment variables 
to trigger a build dispatch a webhook to `/trigger/{git-url}`

### Basic example

```yaml
name: test-deployment

actions:
  step-one:
    image: node:latest
    persist:
       - ./dist 
    steps:
      - echo "Hi" >> dist/output
      - echo "World" >> dist/output

  step-two:
    image: debian:latest
    steps:
      - echo Hi!
      - ls dist

```

### Parallel builds

```yaml
name: test-deployment

groups:
  - one
  - two

actions:
  step-one:
    group: one
    image: node:latest
    steps:
      - echo "Hi"
      - echo "World"

  step-two:
    group: two
    image: node:latest
    steps:
      - echo "Yo"
      - echo "What's up"
```

### Functions

```yaml
name: test-deployment

functions:
  reuse_me:
    params:
      VAL_A
      VAL_B
    steps:
      - echo $VAL_A $VAL_B

actions:
  step-one:
    image: node:latest
    steps:
      - ::reuse_me(test, value)

  step-two:
    image: node:latest
    steps:
      - ::reuse_me(test, value)
```

### Imports

```yaml
name: test-deployment

imports:
  - ./functions.yml

actions:
  step-one:
    image: node:latest
    steps:
      - ::reuse_me(test, value)
```

