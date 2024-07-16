# Run the commands serially

init: ## Initialize the server with the basic setup
	@ansible-playbook playbooks/init.yml -i playbooks/hosts.ini -l server0 -u root -k

upgrade: ## Upgrade the server to the latest LTS
	@ansible-playbook playbooks/upgrade.yml -i playbooks/hosts.ini -l server0 -u ubuntu

add_ssh: ## Add the current user's SSH key to the server
	@ansible-playbook playbooks/add_ssh.yml -i playbooks/hosts.ini -l server0 -u ubuntu

add_docker: ## Setup Docker on the server
	@ansible-playbook playbooks/add_docker.yml -i playbooks/hosts.ini -l server0 -u ubuntu

add_caddy: ## Setup a simple 'hello world' Caddy server
	@ansible-playbook playbooks/add_caddy.yml -i playbooks/hosts.ini -l server0 -u ubuntu
