# PodChaosMonkey

Application for testing the resilience of workloads in kubernetes.

- [How it works](#how-it-works)
  - [Configuration](#configuration)
  - [Schedule format](#schedule-format)
  - [Blacklist](#blacklist)
- [Usage](#usage)
  - [Build image](#build-image)
  - [Deploy](#deploy)
  - [Clean deployed resources](#clean-deployed-resources)
- [Development](#development)
  - [Run PodChaosMonkey locally](#run-podchaosmonkey-locally)
  - [Generate mocks](#generate-mocks)
  - [Run tests](#run-tests)

## How it works

Chaos Monkey deletes a random pod in a defined namespace (`workloads` by default) on a schedule
(each 5 seconds by default).

### Configuration

You can set the following configuration parameters as environment variables:

|    Configuration    |                           Description                           |   Default   |
|:-------------------:|:---------------------------------------------------------------:|:-----------:|
|     `NAMESPACE`     |                   Namespace of the workloads                    |  workloads  |
|     `SCHEDULE`      |                     Schedule to delete pods                     | '* * * * *' |
|  `SCHEDULE_FORMAT`  | Format of the schedule*, could be 'cron' or 'cron-with-seconds' |   'cron'    |
| `IS_INSIDE_CLUSTER` |     Indicates if Chaos Monkey is running inside the cluster     |    true     |

> \* Format are explained [here](#schedule-format).

You can set this configuration parameters in the following [configmap](./kubernetes/pod-chaos-monkey/pod-chaos-monkey.configmap.yml).

### Schedule format

As Cron format doesn't support schedule in seconds interval, Chaos Monkey allow users to set a Cron with seconds custom 
format. This format is the same as Cron, but with an addition field for seconds at the beginning.

### Blacklist

In some scenarios, pods should not be deleted by PodChaosMonkey. An example is those pods that are running a long and
heavy data processing task which has to be restarted if it is interrupted. To avoid deleting these pods, you can add them
to the blacklist.

The blacklist is a YAML file with the following structure:

```yaml
labels: []
fieldSelectors: []
```

* `labels` are a key-value that represent the labels of the pods to be blacklisted.
* `fieldSelectors` are a key-value that represent the field selectors of the pods to be blacklisted. For more information,
  see [Kubernetes API](https://kubernetes.io/docs/concepts/overview/working-with-objects/field-selectors/).

You can set the `blacklist.yml` on the following [configmap](./kubernetes/pod-chaos-monkey/blacklist.configmap.yml).

#### Example
The following `blacklist.yml` excludes pods with the label `app=my-app` and pods with status `Pending`:

```yaml
labels:
  app: my-app
fieldSelectors:
  status.phase: Pending
```

## Usage

### Build Image

An image for PodChaosMonkey is available on [DockerHub](https://hub.docker.com/repository/docker/yisusisback/pod-chaos-monkey).
If you want to build your own image, run:
```bash
docker build -t [image-name] .
```

Also, you have to update de Kubernetes [deployment](./kubernetes/pod-chaos-monkey/pod-chaos-monkey.deployment.yml) for
PodChaosMonkey to use your image.

### Deploy
To deploy PodChaosMonkey, run the following command from the project's root directory:

```bash
./scripts/deploy.sh
```

### Clean deployed resources
To clean previously deployed resources, run the following command from the project's root directory:

```bash
./scripts/clean.sh
```

## Development

### Run PodChaosMonkey locally
You can run PodChaosMonkey locally by running the following command:

```bash
IS_INSIDE_CLUSTER=false go run main.go
```

If its running locally, IS_INSIDE_CLUSTER should be set to false, and the Kubernetes clientset is generated with your
kube config file.

### Generate mocks
```bash
go generate ./...
```

### Run tests
```bash
go test ./...
```
