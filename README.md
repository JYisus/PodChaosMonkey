# PodChaosMonkey

Application for testing the resilience of workloads in kubernetes.

- [How it works?](#how-it-works)
  - [Configuration](#configuration)

## How it works?

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

You can set this configuration parameters in the following [configmap](./kubernetes/chaos-monkey/chaos-monkey.configmap.yml).

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

You can set the `blacklist.yml` on the following [configmap](./kubernetes/chaos-monkey/chaos-monkey.configmap.yml).

#### Example
The following `blacklist.yml` excludes pods with the label `app=my-app` and pods with status `Pending`:

```yaml
labels:
  app: my-app
fieldSelectors:
  status.phase: Pending
```