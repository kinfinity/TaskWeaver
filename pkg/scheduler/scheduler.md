# Scheduler
schedule tasks unto the Workers

Feasability:
    This phase assesses whether it’s even possible to schedule a
task onto a worker. check resources and capacity of the workers to handle the task

Scoring: This phase takes the workers identified by the feasability phase
and gives each one a score.

Picking: From the list of scores, the scheduler
picks the best one. This will be either the highest or lowest score.


