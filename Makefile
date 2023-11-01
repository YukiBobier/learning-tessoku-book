.PHONY: init_module

init_module:
ifndef MODULE
	$(error MODULE is not set)
endif
	mkdir $(MODULE)
	cd $(MODULE) && go mod init $(MODULE)
	go work use $(MODULE)
	git add $(MODULE) go.work
	git commit -m "build: init $(MODULE)"
