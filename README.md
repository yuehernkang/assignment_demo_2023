# assignment_demo_2023

![Tests](https://github.com/TikTokTechImmersion/assignment_demo_2023/actions/workflows/test.yml/badge.svg)

## Architecture
![Architecture](./tiktok.drawio.png)

### NGINX
NGINX is an open-source web server and reverse proxy solution. It is used as a load balancer to the 
2 HTTP servers. It uses a round robin algorithm to determine which server to use.

### Cassandra Database
Cassandra is a NoSQL database that is used to store the data. It is a distributed database that is scalable. 
I think the Cassandra would fit the requirements of the project, especially for a chat application.

## Installation

Requirement:

- golang 1.18+
- docker

To install dependency tools:

```bash
make pre
```

## Run

```bash
docker-compose up -d
```

Check if it's running:

```bash
curl localhost:8080/ping
```

## Areas of improvement
- Cassandra Database:
  - Using Cassandra database, we can easily scale horizontally by adding more nodes to the cluster. 
  - One way we can improve the Cassandra Database implementation is to increase the amount of nodes.
