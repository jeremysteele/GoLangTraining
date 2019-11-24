## How's It Work?

The API server will insert a person into a `people` list in redis. The worker checks the list every second, processes the data, and spits it back
into the `results` list. 

Type validation is done by the API (for free by Unmarshal), but no extra validation is done in this example.

## Running

`docker-compose up -d --build`

## Testing

Run the following curl command to send an example person to the API

```
curl -X POST \
  http://localhost:8080/add \
  -H 'Content-Type: application/json' \
  -H 'Host: localhost:8080' \
  -H 'cache-control: no-cache' \
  -d '{"first_name":"Jeremy", "last_name": "Steele", "age":123}'
```

You should get back a 200 with an `added person to queue` message.

Open redis commander at http://localhost:8081/ and you should see the
resulting string of the combined/processed data in the `results` list.  