# Onion-Gate: Secure Chat Room Prototype

Onion-Gate is designed as a demonstrative implementation of a secure chat room. Constructed with `HTMX` and `Go`, it serves as an educational reference. It is imperative to note that Onion-Gate is not intended for production deployment; its primary purpose is for illustrative and learning contexts.

## Key Features
- 🔒 **Enhanced Security**: Incorporates a secure login mechanism.
- 🚪 **Restricted Access**: Ensures chat room accessibility only post-authentication.
- ⚡ **Real-time Chat**: Seamlessly engage in instantaneous conversations with live chat updates.

## Deployment Options
You have a few deployment options of which I will walk through.

1. Local (testing etc...)
2. Tor (.onion) Deployment
3. Regular Deployment

### Local Deployment
To deploy and test locally you can use the docker file, or simply run the `GO` application.

#### Docker
You can run these commands to build your docker image locally (assuming your machine has docker installed.)

```bash
# Build your docker image.
docker build -t onion-gate:latest .

# Run your docker iamge.
docker run -p 8080:8080 onion-gate:latest

# Verify the container is running (Optional)
docker container ps
```

Now you can navigate to [http://localhost:8080/login](http://localhost:8080/login), and verify.

#### Go Run
You can also just run the applcaiton using the `GO` cli, by using the following comand.

```go run cmd/onion_gate/main.go
# Assuming you are in the root of the project directory.
go run cmd/onion_gate/main.go
```
Now you can navigate to [http://localhost:8080/login](http://localhost:8080/login), and verify.

### Tor (.onion) Deployment
This was my first time every playing around with this. You can follow these steps to deploy (I used an ec2, but any box would work. For example a digital ocean droplet, or even your personal box). It's all wrappe in docker, so you just need to run the below.

```bash
# Start the containers
docker-compose up -d --build

# Access your onion address
docker-compose exec tor cat /var/lib/tor/hidden_service/hostname
xuc3soorvlcwjbocrjqagab3t3cbfb7yn6dfaplf5gxtf5a2kbdy5vqd.onion

# Kill if needed
docker-compose down
```

You can now access the domain found above via your Tor browser.

NOTE: You can find formal Tor deployment docs [here](https://community.torproject.org/onion-services/setup/)

### Regular deploy
You can simpoly use the base Docker file `./deploy/Dockerfile` on any box such as a ec2 or DigitalOcean droplet, as shown in the local host section. If you want to run on ECS or something else, it would require more steps.