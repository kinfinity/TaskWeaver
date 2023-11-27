The task  
The job  
The scheduler  
The manager  
The worker  
The cluster  
The CLI

processes commands which can be ran against tasks @runtime besides \[stop|\]  
command interface which can be implemented by each task

The first is long-running services that should “never” go  
down, and handle short-lived latency-sensitive requests (a  
few ms to a few hundred ms)

Task:

Granularity: A task is typically smaller in scope and represents a specific action or activity that contributes to achieving a larger goal or project.  
Duration: Tasks are generally shorter in duration and can be completed within hours or days.  
Dependencies: Tasks can have dependencies on other tasks and often form part of a sequential workflow.  
Assignee: Tasks are often assigned to individuals or team members responsible for their completion.

Job:

Granularity: A job can refer to a broader unit of work, which may encompass multiple tasks or activities.  
Duration: Jobs are usually longer in duration, potentially taking weeks, months, or even years to complete.  
Dependencies: Jobs may involve complex dependencies but can also encompass multiple tasks or sub-jobs.  
Assignee: Jobs may be assigned to teams, departments, or individuals responsible for overseeing and managing the entire job.

JOB

- name
- owner
- number of tasks + replicas
- priority
- resources required for the job
- time limit per task
- command line template
- environment variables
- input files/directories
- output directory
- log file location
- list of dependencies on other jobs or tasks. If any  
  dependency fails then this job will not run at all.
- status: pending | running | completed | failed | killed
- start time
- end time
- exit code if it has finished successfully
- reason why it stopped

Job descriptions are written in a declarative language

Jobs can have constraints to force  
its tasks to run on machines with particular attributes such as  
processor architecture, OS version, or an external IP address

Each task maps to a set of Linux processes running in  
a container on a machine

no VMs or Virtualization

A task has properties too, such as its resource requirements and the task’s index within the job. Most task properties are the same across all tasks in a job, but can be overridden – e.g., to provide task-specific command-line flags.  
Each resource dimension (CPU cores, RAM, disk space,  
disk access rate, TCP ports,2 etc.) is specified independently  
at fine granularity; we don’t impose fixed-sized buckets or  
slots (x5.4).

A user can change the properties of some or all of the  
tasks in a running job by pushing a new job configuration  
to Borg, and then instructing Borg to update the tasks to  
the new specification

STATES  
submit > accept | reject

- Pending (update) (failed, kill, lost) -> schedule
- Running (update, evict: pending) -> finish, failed, kill, lost
- Dead (submit: pending)  
  Completed

look into handling different workload types

- container runtimes
- stateful workloads
- batch workloads
- interactive workloads
- streaming workloads
- data processing workloads
- distributed training workloads
- serverless workloads

Task Definition

- name of container image for task
- amount of resources for task
- number of replicas
- restart policies

The scheduler should perform these functions:

1.  Determine a set of candidate machines on which a task could run.
2.  Score the candidate machines from best to worst.
3.  Pick the machine with the best score

EPVM scheduler (used as part of Google’s Borg scheduler)

The manager is the brain of an orchestrator and the main entry point for  
users.The manager also periodically collects metrics from each of its workers,  
which are used in the scheduling

The manager should do the following:

1.  Accept requests from users to start and stop tasks.
2.  Schedule tasks onto worker machines.
3.  Keep track of tasks, their states, and the machine on which they run.
4.  Persist jobs in the system in the datastore

The worker provides the muscles of an orchestrator.

(Tainting Workers)  
The worker is responsible for the following:

1.  Running tasks as Docker containers.
2.  Accepting tasks to run from a manager.
3.  Providing relevant statistics to the manager for the purpose of  
    scheduling tasks.
4.  Keeping track of its tasks and their state.

Finally, our CLI, the main user interface, should allow a user to:

1.  Start and stop tasks
2.  Get the status of tasks
3.  See the state of machines (i.e. the workers)
4.  Start the manager
5.  Start the worker

servers - leader & followers

- replication across servers  
  workers - clients
- use rpc to communicate with servers

Operators should define how Tasks are executed => BashOperator, PythonOperator, GoOperator, DockerOperator

DataSources define data/ params for the task -> move data between Tasks

\=> Components of Master Node - Control Plane

- In memory persistent storage - etcd
- scheduler
- api-server
- controller manager
- Plugin Controller - connect to trigger external workloads

\=> Components of Worker Nodes

- proxy - for network routing
- Task Operators
- Task Agent
