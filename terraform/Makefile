plan_file=.terraform/plan.terraform

.PHONY: init refresh plan apply destroy list output fmt lint

init:
	@rm -rf ./.terraform
	@terraform init

refresh:
	@terraform refresh

plan:
	@terraform plan -refresh=true -out=${plan_file}

apply:
	@terraform apply ${plan_file}

destroy:
	@terraform destroy

list:
	@terraform state list

output:
	@terraform output

fmt:
	@terraform fmt -recursive

lint:
	@tflint
