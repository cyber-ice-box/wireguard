sqlc:
	docker-compose -f docker-compose.dev.yml run --rm sqlc

proto:
	docker-compose -f docker-compose.dev.yml run --rm protobufCompiler

push:
	cd .. && docker image build -t cybericebox/wireguard . -f ./wireguard/Dockerfile
	docker push cybericebox/wireguard
