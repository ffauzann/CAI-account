DB_PG_SCHEMA=cai_user

protocgen:
	cd proto/ && \
	protoc --go_out=. --go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=false *.proto -I${GOPATH}/src -I. && \
	cd ..

# requires bufbuild/buf/buf
buf-gen:
	cd proto/ && \
	buf generate && \
	cd ..

mockgen:
	mockery --config mockery.yaml

migrate-up:
	migrate -path internal/migration -database "postgres://${DB_PG_HOST}:${DB_PG_PORT}/${DB_PG_SCHEMA}?sslmode=disable" up

migrate-down:
	migrate -path internal/migration -database "postgres://${DB_PG_HOST}:${DB_PG_PORT}/${DB_PG_SCHEMA}?sslmode=disable" down

docker-build:
	docker build -t user . && \
	docker tag user $(USERNAME)/user:$(VERSION) && \
	docker push $(USERNAME)/user:$(VERSION)

k8s-destroy: # Clear data async
	kubectl delete deployments user & \
	kubectl delete services user & \
	kubectl delete pvc postgres-persistent-storage-postgres-statefulset-0 & \
	kubectl delete pvc postgres-pvc & \
	kubectl delete secret postgres-secret &\
	kubectl delete statefulset postgres-statefulset &\
	kubectl delete service postgres &\
	kubectl delete deployments redis &\
	kubectl delete service redis &\
	kubectl delete job create-db-job

k8s-deploy:
	kubectl apply -f k8s/local

test:
	go test -cover -race ./...

.PHONY: gen test protocgen