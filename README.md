[![Build Status](https://travis-ci.org/trigun117/ProxyGrabber.svg?branch=master)](https://travis-ci.org/trigun117/ProxyGrabber) [![codecov](https://codecov.io/gh/trigun117/ProxyGrabber/branch/master/graph/badge.svg)](https://codecov.io/gh/trigun117/ProxyGrabber) [![Go Report Card](https://goreportcard.com/badge/github.com/trigun117/ProxyGrabber)](https://goreportcard.com/report/github.com/trigun117/ProxyGrabber)
# ProxyGrabber

ProxyGrabber

![site screenshot](https://github.com/trigun117/ProxyGrabber/blob/master/image.JPG)

On http://tproxyt.tk you can find live example

# Get Started

For start, build docker image from Dockerfile and run with this command
```
docker run -d \
-p 80:80 \
-e EF=email_from \
-e EFL=email_from_login \
-e EFP=email_from_password \
-e APIPAS=api_password \
-e CORSS=authorized_domain \
-e TARGET=target_site \
-e REG=regular_expression \
-e ET=email_to \
--restart always \
image_name
```
and then open http://localhost or http://your_server_ip

# License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
