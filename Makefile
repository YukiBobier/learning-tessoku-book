.PHONY: init_module commit_solution commit_refactoring

init_module:
ifndef MODULE
	$(error MODULE is not set)
endif
	mkdir $(MODULE)
	cp ./_template/main.go $(MODULE)/main.go
	cp ./_template/main_test.go $(MODULE)/main_test.go
	cd $(MODULE) && go mod init $(MODULE)
	go work use $(MODULE)
	git add $(MODULE) go.work
	git commit -m "build: init $(MODULE)"

commit_solution:
ifndef MODULE
	$(error MODULE is not set)
endif
	git add $(MODULE)
	git commit -m "feat: solve $(MODULE)"

commit_refactoring:
ifndef MODULE
	$(error MODULE is not set)
endif
	git add $(MODULE)
	git commit -m "refactor: refactor $(MODULE)"
