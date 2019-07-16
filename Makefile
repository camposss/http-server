docker: 
	docker build -t http-server:${TAG} -f Dockerfile .


push: docker
	echo "${DOCKER_PASSWORD}" | docker login --username "${DOCKER_LOGIN}" --password-stdin \
	docker push camposss/http-server:${TAG}