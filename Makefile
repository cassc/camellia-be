.PHONY: dev build deploy
dev:
	go run server/main.go
build:
	go build server/main.go
deploy: build
	rsync -av main ubgf:/opt/apps/goemt/goemt
	rsync -av templates ubgf:/opt/apps/goemt
	rsync -av news ubgf:/opt/apps/goemt
	ansible-playbook goemt-ansible.yml
