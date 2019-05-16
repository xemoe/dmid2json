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
DOCKER_GO_IMAGE="my-golang-serial"

#
# Functions
#
define docker_build
	@echo
	docker build -t ${DOCKER_GO_IMAGE} .
endef

define docker_run
	@echo
	@docker run --rm \
		-v ${ROOT_DIR}:${DOCKER_GO_PATH} \
		-w ${DOCKER_GO_PATH} \
		--device /dev/mem:/dev/mem \
		--cap-add SYS_RAWIO \
		${DOCKER_GO_IMAGE} \
		$1
endef

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
	$(call docker_build)
	@echo "[x]  Go build image has been created!"

#
# target: test
#
.PHONY: go-get
go-get:
	$(call docker_run,go get)
	@echo "[x]  Done go get!"

.PHONY: test
test: go-get
	$(call docker_run,go test ./dmid ./parser)
	@echo "[x]  Done go test!"

.PHONY: test-verbose
test-verbose: go-get
	$(call docker_run,go test -v ./dmid ./parser)
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
	docker rm $(DOCKER_PS_EXITED)
endif

