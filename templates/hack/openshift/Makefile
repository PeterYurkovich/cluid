# Deploy operators and resources for observability components.

all: operators resources

operators:
	kubectl apply -k config/operators
	../wait.sh subscription openshift-logging cluster-logging
	../wait.sh subscription openshift-operators-redhat loki-operator
	../wait.sh subscription openshift-netobserv-operator netobserv-operator
	../wait.sh subscription openshift-opentelemetry-operator opentelemetry-operator
	../wait.sh subscription openshift-tempo-operator tempo-operator
	kubectl patch consoles.operator.openshift.io/cluster --type=merge --patch '{ "spec": { "plugins": ["monitoring-plugin", "logging-view-plugin", "netobserv-plugin"]}}'

STORAGE_CLASS?=$(shell kubectl get storageclass -o=jsonpath='{.items[?(@.metadata.annotations.storageclass\.kubernetes\.io/is-default-class=="true")].metadata.name}')

MINIO_ROLLOUT=../wait.sh rollout minio deployment/minio

LOGGING_ROLLOUT=../wait.sh rollout openshift-logging \
deployment.apps/cluster-logging-operator \
deployment.apps/logging-loki-distributor \
deployment.apps/logging-loki-gateway \
deployment.apps/logging-loki-querier \
deployment.apps/logging-loki-query-frontend \
deployment.apps/logging-view-plugin

OPENTELEMETRY_ROLLOUT=../wait.sh rollout openshift-opentelemetry-operator \
deployment.apps/opentelemetry-operator-controller-manager

TEMPO_ROLLOUT=../wait.sh rollout openshift-tempo-operator \
deployment.apps/tempo-operator-controller \
deployment.apps/tempo-simplest-query-frontend \
deployment.apps/tempo-simplest-distributor \
deployment.apps/tempo-simplest-querier \
deployment.apps/tempo-simplest-compactor

NETOBSERV_ROLLOUT=../wait.sh rollout netobserv \
deployment.apps/loki-distributor \
deployment.apps/loki-gateway \
deployment.apps/loki-querier \
deployment.apps/loki-query-frontend \
deployment.apps/netobserv-plugin

resources:
ifeq ($(strip $(STORAGE_CLASS)),)
		$(error cannot determine storage class of cluster. exitting)
endif
	echo "STORAGE_CLASS=$(STORAGE_CLASS)" > config/resources/storage.env
	kubectl apply -k config/resources
	$(MINIO_ROLLOUT)
	$(LOGGING_ROLLOUT)
	$(OPENTELEMETRY_ROLLOUT)
	$(TEMPO_ROLLOUT)
	$(NETOBSERV_ROLLOUT)

clean-operators:
	kubectl delete --ignore-not-found -k config/operators
clean-resources:
	kubectl delete --ignore-not-found -k config/resources
clean-all: clean-resources clean-operators
