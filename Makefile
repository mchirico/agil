
docker-build:
	docker build --build-arg buildtime_variable=${GITHUB_TOKEN} --build-arg webhook=${GITHUB_WEBHOOK_SECRET} --no-cache -t gcr.io/mchirico/agil:test -f Dockerfile .

push:
	docker push gcr.io/mchirico/agil:test

build:
	go build -v .

run:
	docker run --name agil --rm -it -p 3000:3000  gcr.io/mchirico/agil:test


deploy:
	docker build --build-arg buildtime_variable=${GITHUB_TOKEN} --build-arg webhook=${GITHUB_WEBHOOK_SECRET}  --no-cache -t gcr.io/mchirico/agil:test -f Dockerfile .
	docker push gcr.io/mchirico/agil:test
	gcloud run deploy agil  --image gcr.io/mchirico/agil:test --platform managed \
            --allow-unauthenticated --project mchirico \
            --region us-east1 --port 3000 --max-instances 1  --memory 256Mi



