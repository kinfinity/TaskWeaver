# Apache Mesos Architecture Study

Mesos consists of a master daemon that manages agent daemons running on each cluster node, and Mesos frameworks that run tasks on these agents.

The master decides how many resources to offer to each framework according to a given organizational policy, such as fair sharing or strict priority.  
To support a diverse set of policies, the master employs a modular architecture that makes it easy to add new allocation modules via a plugin mechanism.

[https://mesos.apache.org/documentation/latest/architecture/](https://mesos.apache.org/documentation/latest/architecture/)
