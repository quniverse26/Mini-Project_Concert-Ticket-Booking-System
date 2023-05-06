**Hasil Running Docker**

PS C:\go-work\Mini-Project_Concert-Ticket-Booking-System> docker login --username=quniverse26
Password:
Login Succeeded

Logging in with your password grants your terminal complete access to your account.
For better security, log in with a limited-privilege personal access token. Learn more at https://docs.docker.com/go/access-tokens/
PS C:\go-work\Mini-Project_Concert-Ticket-Booking-System> docker build -t github.com/quniverse26/miniproject .
[+] Building 301.8s (16/16) FINISHED
 => [internal] load build definition from Dockerfile                                                                                                 1.1s 
 => => transferring dockerfile: 624B                                                                                                                 0.4s 
 => [internal] load .dockerignore                                                                                                                    0.8s 
 => => transferring context: 2B                                                                                                                      0.2s 
 => [auth] docker/dockerfile:pull token for registry-1.docker.io                                                                                     0.1s 
 => CACHED docker-image://docker.io/docker/dockerfile:experimental@sha256:600e5c62eedff338b3f7a0850beb7c05866e0ef27b2d2e8c02aa468e78496ff5           0.0s 
 => [internal] load metadata for docker.io/library/golang:1.17-alpine                                                                                3.5s 
 => [auth] library/golang:pull token for registry-1.docker.io                                                                                        0.0s 
 => [internal] load build context                                                                                                                   50.6s 
 => => transferring context: 16.85MB                                                                                                                50.3s 
 => CACHED [2/7] WORKDIR /app                                                                                                                        0.0s 
 => [3/7] COPY go.mod ./                                                                                                                             1.0s 
 => [4/7] COPY go.sum ./                                                                                                                             0.4s 
 => [5/7] RUN go mod download                                                                                                                      123.6s 
 => [6/7] COPY . ./                                                                                                                                  6.6s 
 => [7/7] RUN go build -o /github.com/quniverse26/miniproject                                                                                       63.2s 
 => exporting to image                                                                                                                              36.8s 
 => => exporting layers                                                                                                                             33.7s 
 => => writing image sha256:c478c7da8c7edcb9a7642ba0e6bdfd562ced3884b88390d302a7cf3e3846aa19                                                         0.7s 
 => => naming to github.com/quniverse26/miniproject                                                                                                  0.1s 
PS C:\go-work\Mini-Project_Concert-Ticket-Booking-System> docker images
REPOSITORY                           TAG       IMAGE ID       CREATED              SIZE
quniverse26/clean-code               latest    6ff1c0ef3c15   3 weeks ago          448MB
docker/disk-usage-extension          0.2.5     5f8f804dbfa2   9 months ago         2.79MB
PS C:\go-work\Mini-Project_Concert-Ticket-Booking-System> docker image push github.com/quniverse26/miniproject
Using default tag: latest
The push refers to repository [github.com/quniverse26/miniproject]
006c6717dfda: Preparing
4d1fe581c2ab: Preparing
74cf1978a4dd: Preparing
434028243aef: Preparing
fed9565c145c: Waiting
60af18e613ad: Waiting
925e9633fd86: Waiting
7b35f2def65d: Waiting
error parsing HTTP 403 response body: invalid character 'C' looking for beginning of value: "Cookies must be enabled to use GitHub."
PS C:\go-work\Mini-Project_Concert-Ticket-Booking-System> docker images
REPOSITORY                           TAG       IMAGE ID       CREATED         SIZE
github.com/quniverse26/miniproject   latest    c478c7da8c7e   5 minutes ago   471MB
quniverse26/clean-code               latest    6ff1c0ef3c15   3 weeks ago     448MB
docker/disk-usage-extension          0.2.5     5f8f804dbfa2   9 months ago    2.79MB
PS C:\go-work\Mini-Project_Concert-Ticket-Booking-System> docker run c478c7da8c7e 
