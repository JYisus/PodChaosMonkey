# Chaos Monkey

Application for testing the resilience of workloads in kubernetes.

- [How it works?](#how-it-works)
  - [Configuration](#configuration)

## How it works?

Chaos Monkey deletes a random pod in a defined namespace (`workloads` by default) on a schedule
(each 5 seconds by default).

### Configuration

You can set the following configuration parameters as environment variables:

|     Configuration      |                           Description                           |   Default   |
|:----------------------:|:---------------------------------------------------------------:|:-----------:|
|     `CM_NAMESPACE`     |                Namespace of the victim workloads                |  workloads  |
|     `CM_SCHEDULE`      |                     Schedule to delete pods                     | '* * * * *' |
|  `CM_SCHEDULE_FORMAT`  | Format of the schedule*, could be 'cron' or 'cron-with-seconds' |   'cron'    |
| `CM_IS_INSIDE_CLUSTER` |     Indicates if Chaos Monkey is running inside the cluster     |    true     |

> \* Format are explained [here](#schedule-format).

### Schedule format

As Cron format doesn't support schedule in seconds interval, Chaos Monkey allow users to set a Cron with seconds custom 
format. This format is the same as Cron, but with an addition field for seconds at the beginning.
