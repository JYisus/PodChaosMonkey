#!/usr/bin/bash

function createNamespaces() {
    echo "Creating namespace \"workloads\""
    kubectl apply -f ./kubernetes/workloads/workloads.namespace.yml
     echo "Creating namespace \"chaos\""
    kubectl apply -f ./kubernetes/pod-chaos-monkey/chaos.namespace.yml
}

function createResources() {
  echo "Creating \"workloads\" resources"
  kubectl apply -f ./kubernetes/workloads
  echo "Creating \"pod-chaos-monkey\" resources"
  kubectl apply -f ./kubernetes/pod-chaos-monkey
}

createNamespaces
createResources
