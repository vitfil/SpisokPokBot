app_name = $(notdir $(shell pwd))
build_dir = _build

.PHONY: run
run:
	@if pgrep $app_name; then pkill $app_name; fi
	@rm -rf $(build_dir) && mkdir $(build_dir) && clear
	@go build -ldflags "-s -w" -o $(build_dir)/$(app_name) && $(build_dir)/$(app_name) -debug
	@rm -f $(build_dir)/$(app_name)

.PHONY: deploy
deploy:
	@rm -rf $(build_dir) && mkdir $(build_dir)
	@GOOS=linux GOACRCH=amd64 go build -ldflags "-s -w" -o $(build_dir)/$(app_name) && upx $(build_dir)/$(app_name)
	@ssh -i ~/.ssh/vitfil_github vitfil@firefly "sudo systemctl stop $(app_name).service; rm -f /opt/$(app_name)/$(app_name)"
	@scp -i ~/.ssh/vitfil_github $(build_dir)/$(app_name) vitfil@firefly:/opt/$(app_name)
	@ssh -i ~/.ssh/vitfil_github vitfil@firefly "sudo systemctl start $(app_name).service"
