MAIN_FILE = ./cmd/northrend/main.go
OUT_FILE = northrend

run:
	go run ${MAIN_FILE}

build:
	go build -o ${OUT_FILE} ${MAIN_FILE}

mongo:
	-docker run --name northrend-mongo \
		-p 27017:27017 \
		-e MONGO_INITDB_ROOT_USERNAME=root \
      	-e MONGO_INITDB_ROOT_PASSWORD=root \
		-v mongo-volume:/data/db \
		--network northrend-net \
		-d --rm mongo
	-docker run --name northrend-mongo-exp \
		-p 8081:8081 \
		-e ME_CONFIG_MONGODB_ADMINUSERNAME=root \
      	-e ME_CONFIG_MONGODB_ADMINPASSWORD=root \
    	-e ME_CONFIG_MONGODB_URL=mongodb://root:root@northrend-mongo:27017/ \
		--network northrend-net \
		-d --rm mongo-express

mongostop:
	-docker kill northrend-mongo
	-docker kill northrend-mongo-exp

docker:
	docker build -t yosaa5782/northrend .
	docker run --rm --name northrend-web -p 5000:5000 --network northrend-net -d yosaa5782/northrend

net:
	docker network create northrend-net

deps:
	go mod tidy
	go get -u github.com/gin-gonic/gin
	go get -u github.com/bytedance/sonic
	go get -u github.com/rs/zerolog
	go get -u go.mongodb.org/mongo-driver/mongo
	go get -u github.com/redis/go-redis/v9
	go get -u github.com/gorilla/sessions
