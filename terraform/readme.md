# Terraform

All resources should be added do one of the region submodules under `_modules`.

### Install Terraform

```sh
brew install tfenv # install terraform version manager

# install a terraform version (check your `versions.tf` file for the right version!)
tfenv install
tfenv use
terraform -v
```

Or go to terraform [install doc page](https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli)

## Local Testing

While it can be difficult to validate terraform changes without deploying them out, the `terraform validate` command can be used to catch most common syntax errors.

## Terraform Validate

```sh
# downloads all needed libraries into a cache in a .terraform folder
terraform init -backend=false
terraform validate # runs the validate command
```

## Local Formatting

```sh
# Automatically formats all .tf files in all subdirectories
terraform fmt -recursive .
```

## Deploying

Create a bucket s3 manualy in your account to do the version control and update the `terraform.tf` file with the name and region.

### Terminal Environment Variables

```
export AWS_ACCESS_KEY_ID=your_aws_access_key
export AWS_SECRET_ACCESS_KEY=your_aws_secret_access_key
```

### Initialize

```
terraform init
```

### Create the plan

```
terraform plan -out dev.tfplan -var-file='dev.tfvars'
```

### Deploy it

```
terraform apply -var-file='dev.tfvars'
```

### Destroy it

```
terraform destroy -var-file='dev.tfvars'
```
