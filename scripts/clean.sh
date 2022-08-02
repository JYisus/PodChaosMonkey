#!/usr/bin/bash

kubectl delete -f ./kubernetes/workloads
kubectl delete -f ./kubernetes/pod-chaos-monkey
