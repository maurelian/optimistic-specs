COMPOSEFLAGS=-d
ITESTS_L2_HOST=http://localhost:9545

build: submodules opnode contracts integration-tests
.PHONY: build

build-go: submodules opnode l2os

build-ts: submodules contracts integration-tests
.PHONY: build-ts

submodules:
	# CI will checkout submodules on its own (and fails on these commands)
	if [ -z "$$GITHUB_ENV" ]; then \
		git submodule init; \
		git submodule update; \
	fi
.PHONY: submodules

opnode:
	go build -o ./bin/op ./opnode/cmd
.PHONY: opnode

contracts:
	cd ./packages/contracts && yarn install && yarn build
.PHONY: contracts

integration-tests:
	cd ./packages/integration-tests && yarn install && yarn build:contracts
.PHONY: integration-tests

clean:
	rm -rf ./bin
.PHONY: clean

devnet-up:
	@bash ./ops/devnet-up.sh
.PHONY: devnet-up

devnet-down:
	@(cd ./ops && GENESIS_TIMESTAMP=$(shell date +%s) docker-compose down -v)
.PHONY: devnet-stop

devnet-clean: devnet-down
	rm -rf ./packages/contracts/deployments/devnetL1
	rm -rf ./.devnet
	cd ./ops && docker-compose rm
.PHONY: devnet-clean

test-unit:
	cd ./opnode && make test
	cd ./packages/contracts && yarn test
.PHONY: test

test-integration:
	bash ./ops/test-integration.sh \
		./packages/contracts/deployments/devnetL1
.PHONY: test-integration

devnet-genesis:
	bash ./ops/devnet-genesis.sh
.PHONY: devnet-genesis

l2os:
	go build -o ./bin/l2os ./l2os/cmd/l2os
.PHONY: l2os