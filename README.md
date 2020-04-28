# osi

An OSI stack built with Golang. Mostly as a learning tool.

## Organization

The OSI stack has 7 layers. Starting at the top,

- Application
- Presentation
- Session
- Transport
- Network
- Data link
- Physical

Each layer in this organization has a link to the layer above it and one to the layer below. Each layer uses these links to wrap the data that it gets from the layer above it and pass it to the layer under it.

This model is great to think about what's happening while computers are talking to each other. The implementation is split into a few different protocols. So this package is going to be walking through a simplified implementation of HTTP over TCP and how the data that you provide, for example, a curl command gets to the other computer.

We will also use a simplified version of IP addresses (ones without subnets) because the idea here is to demonstrate how the data flows through the OSI stack while using something like HTTP.

## Resources

- [Cloudflare's OSI page](https://www.cloudflare.com/learning/ddos/glossary/open-systems-interconnection-model-osi/)
- [Wikipedia OSI model entry](https://en.wikipedia.org/wiki/OSI_model)
- [Wikipedia protocol data unit entry](https://en.wikipedia.org/wiki/Protocol_data_unit)
- [Network programming in Go](https://ipfs.io/ipfs/QmfYeDhGH9bZzihBUDEQbCbTc5k5FZKURMUoUvfmc27BwL/index.html)
