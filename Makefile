app_name = $(notdir $(shell pwd))
build_dir = _deploy

.PHONY: run
run:
	@if pgrep $app_name; then pkill $app_name; fi
	@rm -f $(build_dir)/$(app_name) && clear
	@go build -ldflags "-s -w" -o $(build_dir)/$(app_name) && $(build_dir)/$(app_name) -debug
	@rm -f $(build_dir)/$(app_name)

#.PHONY: deploy
#deploy:
#	@rm -rf $(dist) && mkdir $(dist)
#	@GOOS=linux GOACRCH=amd64 go build -ldflags "-s -w" -o $(dist)/$(app_name) && upx $(dist)/$(app_name)
#	@ssh -i ~/.ssh/est_rsa root@est-grup.com "systemctl stop est-api-server.service; rm -f /opt/est-server/$(app_name)"
#	@scp -i ~/.ssh/est_rsa $(dist)/$(app_name) root@est-grup.com:/opt/est-server
#	@ssh -i ~/.ssh/est_rsa root@est-grup.com "systemctl start est-api-server.service"
