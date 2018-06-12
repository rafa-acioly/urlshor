## URLShor is a simple url shortner made with Go, Redis and Postgres

Why shorten the URL?

- Easy URL memorization
- Less characters to use in social media

## Advantages to use this project

- Decoupled front-end (build your own front-end)
- Decoupled back-end (open API)
- Easy to customize
- Dockerized

## Main configurations

Each URL sent to be shortened [will be cached for 30 minutes.](https://github.com/rafa-acioly/urlshor/blob/master/redis/redis.go#L29)
as well saved in database. Any request that was erased from redis but still in database [will be bounded again to redis](https://github.com/rafa-acioly/urlshor/blob/master/urlshor.go#L99).

The back-end API only respond with a simple `json` that contains the `code` to be used after your domain, example:
```
curl -X POST http://localhost:5000/short \
  -d '{
	"url": "http://an-long-url.com.br/page/"
}'


{"url":"F9WL"}
```

### Http port

By default the [port `5000` is used](https://github.com/rafa-acioly/urlshor/blob/master/urlshor.go#L25), feel free to changed it, but don't forget to [change on docker compose file](https://github.com/rafa-acioly/urlshor/blob/master/docker-compose.yml#L11) as well.

### Routes
- GET `/` - show the main html page with a form

- GET `/{id}` - redirect the user for the given encoded key given if exist on redis or in database

- POST `/short` - return a `json` with the given key for a url encoded
    - example: `{"url": "9ZYOP"}`