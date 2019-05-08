#
# Default target
#
.PHONY: all
all: test

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

#
# Go build with docker
#
.PHONY: go-build-image
go-build-image:
	@docker build -t ${DOCKER_GO_IMAGE} .

#
# target: test
#
.PHONY: test
test:
	@echo "[ ]  Starting go get"
	@docker run -v ${ROOT_DIR}:${DOCKER_GO_PATH} \
		--rm \
		-v ${ROOT_DIR}/.cache:/go \
		-w ${DOCKER_GO_PATH} \
		${DOCKER_GO_IMAGE} \
		go get
	@echo "[x]  Done go get!"
	@echo "[ ]  Starting go test"
	@echo
	@docker run -v ${ROOT_DIR}:${DOCKER_GO_PATH} \
		--rm \
		--device /dev/mem:/dev/mem \
		--cap-add SYS_RAWIO \
		-v ${ROOT_DIR}/.cache:/go \
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

