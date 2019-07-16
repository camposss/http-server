docker: 
	docker build -t http-server:${TAG} -f Dockerfile .


push: docker
	${DOCKER_PASSWORD} | docker login -u ${DOCKER_LOGIN} --password-stdin \
	docker push camposss/http-server:${TAG}