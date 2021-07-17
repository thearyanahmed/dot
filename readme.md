## DOT 

DOT accepts conventional DNS requests and proxy it to DNS servers with [DNS over TLS (DoT)](https://en.wikipedia.org/wiki/DNS_over_TLS). DoT provides privacy and security improvements taking advantage of encrypted DNS traffic. For more information follow [RFC7858](https://datatracker.ietf.org/doc/html/rfc7858) and [RFC7626](https://datatracker.ietf.org/doc/html/rfc7626) [https://developers.cloudflare.com/1.1.1.1/dns-over-tls](https://developers.cloudflare.com/1.1.1.1/dns-over-tls)
.

### Running with docker 

```bash
docker run -itp 531:53 -p 531:53/UDP thearyanahmed/dot:1.0
```

Set your nameserver to `127.0.0.1` in `/etc/resolv.conf`.

using dig to make the request 
```bash
dig cloudflare.com @127.0.0.1 -p 531
```

### Running locally 
```
git clone git@github.com:thearyanahmed/dot.git
```

cd into the project. 

If you wish to change the  env files make a copy of `.env.example` to `.env`

By default, the values are 
```txt
UPSTREAM_TIMEOUT=2000ms
UPSTREAM_SERVER=1.1.1.1
UPSTREAM_PORT=853
ENABLED_TCP=true
ENABLED_UDP=true
```
To run the application, 

```
go run main.go
```


**Note** Make sure your ports are available. It might not send of an error but will not print anything after priting `setting up dns handler`. I used docker to use a different port.

### Testing
**Note** -p $port must match
```
dig cloudflare.com @127.0.0.1 -p 531
```

Thanks to [Shajal Ahamed](https://github.com/shajalahamedcse) for the idea. 

Useful Links

- [RFC7858](https://datatracker.ietf.org/doc/html/rfc7858)
- [RFC7626](https://datatracker.ietf.org/doc/html/rfc7626)
- [Docker image](https://hub.docker.com/repository/docker/thearyanahmed/dot)
- [github.com/MatthewVance/stubby-docker](https://github.com/MatthewVance/stubby-docker)
- [github.com/miekg/dns](https://github.com/miekg/dns)
- [github.com/jonathanbeber/burrow](https://github.com/jonathanbeber/burrow)

You should use it for learning purpose only.