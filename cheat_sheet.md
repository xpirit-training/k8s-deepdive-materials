# kubectl installed?
Test with `kubectl version`

 [Linux-install](https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/)

 [macOS-install](https://kubernetes.io/docs/tasks/tools/install-kubectl-macos/)

 [Windows](https://kubernetes.io/docs/tasks/tools/install-kubectl-windows/) please consider above options :)

# Kubeconfig

kubernetes takes config file in `~/.kube/config` OR file configured in $KUBECONFIG
One config file can contain multiple contexts.


Request your kubeconfig from your trainer 

`cp trainee-cluster-<your-name>.kubeconfig ~/.kube/config`

and own it:
`chmod 600 ~/.kube/config`

Test with `kubectl config current-context` and `kubectl get nodes`

# Shortcuts

Who has time to type `kubectl` ? `alias k=kubectl`


Do this! `source <(kubectl completion bash)`

Krew is a plugin manager for kubectl that can help us with exploring the cluster
```bash
# Install Krew
(
  set -x; cd "$(mktemp -d)" &&
  curl -fsSLO "https://github.com/kubernetes-sigs/krew/releases/latest/download/krew.tar.gz" &&
  tar zxvf krew.tar.gz &&
  KREW=./krew-"$(uname | tr '[:upper:]' '[:lower:]')_$(uname -m | sed -e 's/x86_64/amd64/' -e 's/\(arm\)\(64\)\?.*/\1\2/' -e 's/aarch64$/arm64/')" &&
  "$KREW" install krew
)
```

Once installed you can install plugins with `kubectl krew install <plugin-name<`

e.g. `kubectl krew install ctx` and `kubectl krew install ns` and `kubectl krew install pod-inspect`

explore krew plugins with `kubectl krew search`


# Helpful cluster commands

`kubectl get nodes`

`kubectl api-resources `

# Helpful commands


## Creating yamls via kubectl
`kubectl run busybox-pod --image=busybox --dry-run=client -o yaml > busybox.yaml`

`kubectl create secret generic example-secret --from-file=./secret.file --dry-run=client -o yaml > secret.yaml`


## Other helpful and basic commands
`kubectl get <resource-name>`  some have the `-o wide` option for more info

`kubectl -n <namespace-name> get <resource-name>`

`kubectl edit <resource-name>`

`kubectl port-forward <pod-name> <local-port>:<pod-port>`

`kubectl scale <scalable-resource> --replicas=4`


## Kube config 
`kubectl config current-context`

`kubectl config get contexts`

`kubectl config use-context <context-name>`

`kubectl config view`

# Creating and deleting

`kubectl create -f <file.yaml>`

`kubectl apply -f <file.yaml>`

`kubectl delete <resource-name> <pod-name>`

`kubectl delete -f <file.yaml>`


# Debugging stuff

`kubectl logs <pod-name>` (optional with `-c <container-name>`)

`kubectl get all`

`kubectl describe <ressource-name>`

`kubectl exec -it <pod-name> -- /bin/sh`

# Other

`kubectl explain <resource-name>`

`kubectl cp <pod-name:<pod-path> <local-path>`

`kubectl run -it --image=nicolaka/netshoot -- bash`
