# Running in AWS

## Setup the environment

Follow the [instructions](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html) to install the `aws` CLI.

Then everything else:
```sh
python3 -m venv venv
source venv/bin/activate
pip install --upgrade pip
pip install aws-parallelcluster Flask==2.2.5

# Install Node.js and npm (on macOS used "port install nodejs18 npm9")
npm install aws-cdk
export PATH=$PWD/node_modules/.bin:$PATH
```

Follow the [instructions](https://docs.aws.amazon.com/parallelcluster/latest/ug/install-v3-configuring.html) to configure and run a cluster:
```sh
aws configure
aws ec2 create-key-pair \
    --key-name mobius-cluster-keypair \
    --key-type rsa \
    --key-format pem \
    --query "KeyMaterial" \
    --output text > mobius-cluster-keypair.pem
chmod 400 mobius-cluster-keypair.pem

pcluster configure --config cluster-config.yaml
# Edit cluster-config.yaml, set MinCount equal to MaxCount to have all worker nodes immediately available, add OnNodeConfigured scripts (see the `cluster-config.yaml` example)
pcluster create-cluster --cluster-configuration cluster-config.yaml --cluster-name mobius-cluster --region us-west-1
pcluster list-clusters
```

Wait until the last command shows `CREATE_COMPLETE`, then login:
```sh
pcluster ssh --cluster-name mobius-cluster -i ./mobius-cluster-keypair.pem
```

## Run mobius

Back to the head node, as the local user:
```sh
git clone https://github.com/danielpickens/mobius.git
cd mobius

# Download the mobius-kubelet binary (adjust the version in the URL)
wget https://github.com/danielpickens/mobius/releases/download/v0.1.0/mobius-kubelet_v0.1.0_linux_amd64.tar.gz
tar -zxvf mobius-kubelet_v0.1.0_linux_amd64.tar.gz
mkdir -p bin
mv mobius-kubelet bin/
```

Run each of the following in a separate window:
```sh
make run-kubemaster
make run-kubelet
```

And you are all set:
```sh
export KUBE_PATH=~/.k8sfs/kubernetes/
export KUBECONFIG=${KUBE_PATH}/admin.conf
kubectl get nodes
```

## Clean up the environment

To clean up the cluster:
```sh
pcluster delete-cluster --region us-west-1 --cluster-name mobius-cluster
```

To clean up the networks (VPCs), list them and then delete any ones that show up:
```sh
aws --region us-west-1 cloudformation list-stacks --stack-status-filter "CREATE_COMPLETE" --query "StackSummaries[].StackName" | grep -e "parallelclusternetworking-"
aws --region us-west-1 cloudformation delete-stack --stack-name <stack_name>
```
