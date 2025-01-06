#!/bin/bash

# Step 1: Generate a GPG key if not already present
echo "Generating GPG key..."
gpg --full-generate-key

# Step 2: List secret keys and extract the GPG Key ID
echo "Listing GPG secret keys..."
KEY_ID=$(gpg --list-secret-keys --keyid-format LONG | grep sec | awk '{print $2}' | cut -d'/' -f2)

# Step 3: Display the generated Key ID
echo "Generated GPG Key ID: $KEY_ID"

# Step 4: Encrypt the values-secrets-to-encrypt.yaml file using the generated GPG key
echo "Encrypting the values-secrets-to-encrypt.yaml file..."
sops -e --pgp "$KEY_ID" values-secrets-to-encrypt.yaml > values-secrets-encrypted.yaml

echo "Encryption complete. Encrypted file saved as values-secrets-encrypted.yaml"
