---
layout: post
title: "Twirp: from RPC to a tiny Go service"
date: 2026-10-11
categories: go rpc
description: "A beginner's tour of RPC, gRPC, and Twirp, with a runnable Go example."
---

I met Twirp through a new project and wanted the shortest useful path from “what is RPC?” to a request I could actually send. Here is that path.

**Prefer pictures?** [Start the six-page visual walkthrough]({{ '/prototypes/twirp-in-action/diagrams/01-rpc.html' | relative_url }}). Its links take you from the first RPC to a real Twirp HTTP request; each page has a click-to-open glossary.

## RPC first

Imagine a website backend needs to greet Ada. The website backend is the **client**: it asks for work. A separate Greeter service is the **server**: it does the work and replies. “Client” names the caller's role in this request; here, that caller is a Go backend.

A **remote procedure call** (RPC) lets the website backend invoke the Greeter service's named `Greet` operation. A local call such as `Greet("Ada")` hides no network. An RPC has to encode the request into bytes, send them, find the server method, decode the request, and send back a response or an error. That network boundary is why timeouts and failures matter even when the client API looks like an ordinary function call.

**Encoding** turns a structured value such as `Name: "Ada"` into bytes for the trip; **decoding** rebuilds the value at the other end. JSON makes those bytes readable as `{"name":"Ada"}`. Protobuf can encode the same fields more compactly in binary.

In a **protobuf** service definition, you describe the method and its messages once:

```proto
service Greeter {
  rpc Greet(GreetRequest) returns (GreetResponse);
}
```

Code generators turn that contract into client and server types. Protobuf is the schema and message format; it is not, by itself, the transport that moves a call over the network.

## Where gRPC fits

gRPC is an RPC framework that can generate those client and server APIs from a service definition. It commonly uses protobuf messages and runs over HTTP/2. It supports a single request and response (**unary RPC**) as well as client, server, and bidirectional **streaming** calls. [The gRPC core concepts guide](https://grpc.io/docs/what-is-grpc/core-concepts/) lays out all four call types.

## Why Twirp appeared

Twitch wanted the typed, generated service model but found gRPC's HTTP/2 requirement and broader transport features more than its mainly request/response services needed. Its engineers [introduced Twirp in 2018](https://blog.twitch.tv/en/2018/01/16/twirp-a-sweet-new-rpc-framework-for-go-5f2febbf35f/) to keep the service contract and generated code while making the wire protocol simpler. That is Twitch's stated motivation, not a claim that gRPC is generally too complex.

Twirp maps a method to an ordinary HTTP `POST` URL, accepts protobuf or JSON request bodies, and works over HTTP/1.1 as well as HTTP/2. It uses Go's standard `net/http` handler. The trade-off is scope: Twirp handles unary calls, while gRPC also offers streaming. The [Twirp wire protocol](https://twitchtv.github.io/twirp/docs/spec_v7.html) spells out the route and encoding rules.

## Twirp in action

The [Go prototype](https://github.com/koffeeandkode-gs/koffee-and-kode/tree/main/prototypes/twirp-in-action) defines one `Greeter/Greet` RPC. Its generated Twirp handler receives this request:

```http
POST /twirp/demo.greeter.Greeter/Greet
Content-Type: application/json

{"name":"Ada"}
```

The server replies with `{"message":"Hello, Ada!"}`. You can send the request with `curl`, without a generated client. The Go client in the prototype uses Twirp's generated protobuf client for the same RPC. If the name is empty, the service returns a Twirp `invalid_argument` error as JSON with HTTP 400. [Twirp's routing guide](https://twitchtv.github.io/twirp/docs/routing.html) documents these paths and error responses.

That is the useful mental model: **protobuf defines the API, generated code handles the HTTP details, and your Go method supplies the behavior.**
