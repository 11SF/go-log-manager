start container
docker run --rm --name log-family-pay --network=family-pay -p 8080:8080 -d log-family-pay:1.0.0