# Types of Load Balancers: Layer 4, Application, SSL, and Network

This document explains the key differences between various types of load balancers—including Layer 4 Load Balancer, Application Load Balancer, SSL Load Balancer, and Network Load Balancer—in a cloud or data center environment.

---

## Table of Contents

- [Introduction](#introduction)
- [Comparison Table](#comparison-table)
- [Detailed Descriptions](#detailed-descriptions)
  - [Layer 4 Load Balancer](#layer-4-load-balancer)
  - [Application Load Balancer (ALB)](#application-load-balancer-alb)
  - [SSL Load Balancer](#ssl-load-balancer)
  - [Network Load Balancer (NLB)](#network-load-balancer-nlb)
- [Choosing the Right Load Balancer](#choosing-the-right-load-balancer)

---

## Introduction

Load balancers distribute network or application traffic across multiple servers to improve reliability, performance, and availability. Different types of load balancers operate at different layers of the OSI model and provide varying features, flexibility, and use cases.

---

## Comparison Table

| Load Balancer Type     | OSI Layer | Primary Protocols         | Features                                         | Use Cases                       |
|-----------------------|-----------|---------------------------|--------------------------------------------------|----------------------------------|
| Layer 4 Load Balancer | 4 (Transport) | TCP, UDP                | Fast, low latency, forwards packets by IP/port   | Any TCP/UDP service, simple routing |
| Application Load Balancer (ALB) | 7 (Application) | HTTP, HTTPS               | Routing by URL, Host, Headers, Cookies; Websockets, SSL termination | Web apps, microservices, APIs   |
| SSL Load Balancer     | 7 or 4*     | HTTPS / SSL/TLS           | Offloads SSL encryption/decryption (SSL termination), can combine with other LB types | Securing web traffic, compliance |
| Network Load Balancer (NLB) | 4 (Transport) | TCP, UDP, TLS passthrough | Handles millions of requests/sec, ultra-low latency, supports static IP | High-scale microservices, streaming, gaming |

**\*** An SSL Load Balancer terminates SSL at either layer 4 or layer 7 depending on configuration and product.

---

## Detailed Descriptions

### Layer 4 Load Balancer

- Operates at the transport layer (TCP/UDP).
- Routes traffic based on IP address and port.
- Does not inspect payload (application data).
- Fast and efficient, best for simple traffic distribution or non-HTTP protocols.

### Application Load Balancer (ALB)

- Works at the application layer (HTTP, HTTPS).
- Can make routing decisions based on content like host, path, header, or method.
- Supports advanced features including WebSockets, sticky sessions, and SSL termination.
- Ideal for microservices and modern web architectures.

### SSL Load Balancer

- Terminates or offloads SSL/TLS traffic so backend servers handle unencrypted traffic.
- Increases performance of applications by moving compute-heavy cryptography off servers.
- Often integrated into other load balancer types (ALB/NLB) as an option.
- Essential for HTTPS-enabled applications, compliance, and security.

### Network Load Balancer (NLB)

- Operates at layer 4 but is optimized for ultra-high performance and low latency.
- Can handle volatile or massive traffic spikes efficiently.
- Supports static IP addresses, preserves the client IP address, and is suitable for real-time applications.
- Often used for scaling TCP/UDP workloads, IoT, gaming, and edge computing.

---

## Choosing the Right Load Balancer

- **Layer 4/NLB**: Choose if you need high performance, low latency, or must support non-HTTP protocols.
- **ALB**: Choose for HTTP/HTTPS apps, web services, content-based routing, SSL termination, or microservices.
- **SSL LB**: Use when you require dedicated SSL offload for secure connections.
- Many cloud providers offer hybrid load balancers capable of multiple features.

---

*Always match your load balancer type to the needs of your application for optimal performance and security.*


