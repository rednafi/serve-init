# Make all the targets phony
.PHONY: $(MAKECMDGOALS)


##########################################
# App
##########################################
app-lint: ## Lint the demo app
	@cd app && go fmt ./... && go vet ./...

app-test: ## Test the demo app
	@cd app && go test ./...

app-deploy: ## Deploy the demo app
	@kamal setup

##########################################
# Infrastructure
##########################################

# Provision (this is a script because it needs to load the env vars)
tf-init: ## Create the server
	./bin/tf tf-init

tf-plan: ## Plan the server
	./bin/tf tf-plan

tf-apply: ## Apply the server
	./bin/tf tf-apply

tf-refresh: ## Refresh
	./bin/tf tf-refresh

tf-destroy: ## Destroy the server
	./bin/tf tf-destroy

# Configuration
an-init: ## Initialize the server with the basic setup
	@ansible-galaxy install -r ansible/requirements.yml
	@ansible-playbook ansible/init.yml -i ansible/hosts.ini -l server0 --ask-pass

an-upgrade: ## Upgrade the server to the latest LTS
	@ansible-playbook ansible/upgrade.yml -i ansible/hosts.ini -l server1
