.PHONY: init_module

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
