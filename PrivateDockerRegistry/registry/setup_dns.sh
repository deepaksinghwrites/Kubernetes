#!/bin/bash

# Set the registry name
REGISTRY_NAME="my-registry"

# Replace 'docker-registry' with the actual name of your service and namespace if applicable
SERVICE_NAME="private-docker-registry"
NAMESPACE="default"  # Replace with the actual namespace if it's different

# Print the registry name
echo "Registry Name: $REGISTRY_NAME"

# Get the IP address of the service
REGISTRY_IP=$(kubectl get svc $SERVICE_NAME -n $NAMESPACE -o jsonpath='{.spec.clusterIP}')

# Check if the IP was found
if [ -z "$REGISTRY_IP" ]; then
  echo "Error: Could not find the IP for service '$SERVICE_NAME'."
  exit 1
fi

# Print the IP address
echo "Docker Registry IP is: $REGISTRY_IP"

# Export the variables for use in other commands or scripts
export REGISTRY_NAME
export REGISTRY_IP

echo "$REGISTRY_IP $REGISTRY_NAME" | sudo tee -a /etc/hosts > /dev/null



