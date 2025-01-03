#!/bin/bash

set -e  # Exit on error

# Deploy EKS_1
echo "Initializing and applying EKS_1..."
cd EKS_us_east_1
terraform init
terraform validate
terraform plan 
terraform apply -auto-approve

# Deploy EKS_2
echo "Initializing and applying EKS_2..."
cd ../EKS_us_east_2
terraform init
terraform validate
terraform plan 
terraform apply -auto-approve