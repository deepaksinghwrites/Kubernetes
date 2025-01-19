kubectl create secret tls certs-secret \
  --cert=/home/user/certs/tls.crt \
  --key=/home/user/certs/tls.key

kubectl create secret generic auth-secret \
  --from-file=/home/ubuntu/path/to/your/file

kubectl create secret docker-registry nginx-secret \
  --docker-server=my-registry:5000 \
  --docker-username=myuser \
  --docker-password=mypasswd \

