# Get Ansible Vault password from Hashicorp Vault

#### To compile *AnsibleVaultPassFromHashiVault.go* file just execute the following command (It will prepare _AnsibleVaultPassFromHashiVault_ file):
```bash
# go get -u github.com/pkg/errors
# go get -u github.com/hashicorp/vault/api
# go build AnsibleVaultPassFromHashiVault.go
```

#### By default *AnsibleVaultPassFromHashiVault* binary file tries find *CONFIG_FILE* (Path for the configuration file) environment file or *.auth.yaml* file in the current directory. We can configure one of them with two variable values for *VAULT_TOKEN* and *VAULT_URL*. 
```bash
# export CONFIG_FILE=./config.yaml
# cat .auth.yaml
token: ae3bc4cf-2f12-4377-0bd1-d61e2aa3ebf5
vault_addr: http://localhost:8200
```

#### To execute *AnsibleVaultPassFromHashiVault* file we can use the following command (After successful execution of the command it will write Ansible Vault password to the *vault_pass.txt* file):
```bash
# ./AnsibleVaultPassFromHashiVault -k secret/ansible
```
