install_helm:
	helm install service-2-server charts/service-2-server/ --values charts/service-2-server/values.yaml
	helm install service-1-client charts/service-1-client/ --values charts/service-1-client/values.yaml

uninstall_helm:
	helm uninstall service-1-client
	helm uninstall service-2-server

reinstall: uninstall_helm install_helm