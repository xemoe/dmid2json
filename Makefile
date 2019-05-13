#
# Default target
#
.PHONY: all
all: deps test clean

#
# Variables
#
ROOT_DIR=${PWD}
DOCKER_GO_PATH=/usr/src/myapp
DOCKER_GO_IMAGE="my-golang-app"

#
# Dependencies
#
.PHONY: deps
deps: go-build-image
	@echo "[x]  Prepare dependencies"

#
# Go build with docker
#
.PHONY: go-build-image
go-build-image:
	@echo "[ ]  Creating go build image..."
	@echo
	@docker build -t ${DOCKER_GO_IMAGE} .
	@echo
	@echo "[x]  Go build image has been created!"

#
# target: test
#
.PHONY: go-get
go-get:
	@echo "[ ]  Starting go get..."
	@docker run -v ${ROOT_DIR}:${DOCKER_GO_PATH} \
		--rm \
		-w ${DOCKER_GO_PATH} \
		${DOCKER_GO_IMAGE} \
		go get
	@echo "[x]  Done go get!"

.PHONY: test
test: go-get
	@echo "[ ]  Starting go test..."
	@echo
	@docker run -v ${ROOT_DIR}:${DOCKER_GO_PATH} \
		--rm \
		--device /dev/mem:/dev/mem \
		--cap-add SYS_RAWIO \
		-w ${DOCKER_GO_PATH} \
		${DOCKER_GO_IMAGE} \
		go test ./dmid ./parser
	@echo
	@echo "[x]  Done go test!"

.PHONY: test-verbose
test-verbose: go-get
	@echo "[ ]  Starting go test..."
	@echo
	@docker run -v ${ROOT_DIR}:${DOCKER_GO_PATH} \
		--rm \
		--device /dev/mem:/dev/mem \
		--cap-add SYS_RAWIO \
		-w ${DOCKER_GO_PATH} \
		${DOCKER_GO_IMAGE} \
		go test -v ./dmid ./parser
	@echo
	@echo "[x]  Done go test!"

#
# target: clean
#
.PHONY: clean
clean: clean-docker-exited

#
# clean exited docker process
#
DOCKER_PS_EXITED= \
	$(shell docker ps -a -f status=exited -f ancestor=${DOCKER_GO_IMAGE} -q)

.PHONY: clean-docker-exited
clean-docker-exited:
ifneq "$(DOCKER_PS_EXITED)" ""
	@echo "Clean exited docker build process"
	@docker rm $(DOCKER_PS_EXITED)
endif

