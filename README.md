<!-- Shelds -->
[![Go Report Card][go-reportcard-sheild]][go-reportcard-url]
![Go][go-status-url]
[![go.dev reference][godoc-shield]][godoc-url]
[![MIT License][license-shield]][license-url]

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]

<br/>
<p align="center">
  <a href="https://github.com/blushft/strana"></a>
  <h3 align="center">Strana</h3>
  <p align="center">
  Smart, Simple, Scalable analytics for your websites and applications
    <br/>

  </p>
</p>

## Contents
- [Contents](#contents)
- [About](#about)
- [Motivation](#motivation)
- [Roadmap](#roadmap)
- [Architecture](#architecture)
  - [Modules](#modules)
    - [Source Modules](#source-modules)
    - [Processor Modules](#processor-modules)
    - [Sink Modules](#sink-modules)
  - [Platform](#platform)
    - [Server](#server)
    - [Bus](#bus)
    - [Store](#store)
    - [Logger](#logger)


## About

Strana is a modular analytics platform that follows the philosphy "batteries included but not required". This means that as your application grows, Strana can scale with it.

The name Strana is a portmanteau of ***str***_eaming_ ***ana***_lytics_.

## Motivation

Analytics is a crowded space. With so many options, why create another? Strana was born because even with so many products, few, if any, could be described as smart, simple, and scalable.

**Smart**
We don't want to use different analytics for every project. Strana strives to accomidate _your_ usecase instead of inheriting opinions from other verticals. We aren't writing e-commerce apps...

We also trust ourselves to guard the privacy of our users without depriving ourselves of the ability to improve products based on user-based tracking. Strana allows you to anonymize your analytics or not. The choice is yours.

**Simple**
We want to get started fast and learn as we go. It shouldn't take an IT team, marketing department, and data scientists to do basic analytics. Strana is designed for simplicity. Simple to install, simple to configure, simple to extend.

**Scalable**
In the cloud native age, we've come to demand cross-platform, planet-scale applications at the end of every Helm chart. Analytics should be no different. Strana deploys where you want to deploy without deep infrastructure opinions. Run on your laptop, VM, Docker, or Kubernetes, no problem. Run one process or a hundred, same experience.    

## Roadmap

The following items are prioritized for development with the goal of reaching the minimum viable product stage.

- [x] ~~Platform Internals (server, bus, store)~~
- [x] ~~Event and context definiition~~
- [x] ~~Module Interface (Source, Sink, Processor)~~
- [ ] API and Pixel tracking collector module
- [ ] Webhook collector module
- [x] ~~Enhancement processor module~~
- [x] ~~Sink processor module~~
- [x] ~~Raw Event Loader module~~
- [x] ~~Go Tracker~~
- [ ] Javascript Tracker
- [ ] Reporting Module
- [ ] Reporting Frontend

## Architecture

Strana is a modular application, consisting of three module types; Sources, Processors, and Sinks.

### Modules

#### Source Modules

Source modules originate events and publish them to other modules on the platform. Most source modules convert a payload to a raw event. For example, the Tracker module accepts a payload over HTTP, adds server-side contexts, and publishes the event under the topic `collected_raw_events`.

**Available Source Modules**

- Tracker
- Webhook
- Raw Event Loader

#### Processor Modules

Processor modules listen for raw events published by a source or other processor module and, as the name implies, process those events. Typically, processor modules both consume and publish events over through a bus.

It is possible to split or join an event stream through processors as well as produce several events from one source event.

It is important not to confuse Processor Modules with Event Processors. Any module type can employ Event Processors. The point of a Processor Module is to provide enhancement, filtering, or other work in a standalone part of the event stream. 

**Available Processor Modules**

- Enhancer
- Fan-In/Out
- Sink Passthrough

#### Sink Modules

Sink modules represent the end of the line for an event stream. These modules accept events over the bus but do not pass them any further. Sink modules are therefore useful for transforming events into new formats, writing into data warehouses, or storing raw results to be used by other applications.

**Available Sink Modules**

- Raw Event Loader
- Reporter

### Platform

The Strana platform is the set of resources injected into modules to allow events to be captured, published, consumed, and stored.

Central to the platform components is the configuration provider and app controller. The controller creates resources based on configuration as well as creating any configured modules. When modules are brought online, these resources are mounted to the module.

#### Server

The Strana HTTP server. Modules can mount routes on the server accessible to outside clients.

#### Bus

The Strana event bus. Users can configure the bus to control multiple internal or external brokers suchs as NATS, NSQ, Kafka, or various cloud message bus providers. Modules can mount publishers, subscribers, and message routes on the bus, which handles transport.

#### Store

The Strana storage provider. Modules mount services that can read or write from the storage provider.

#### Logger

The Strana structured logger. Modules can consume and extend the logger to provide rich, structured logging.



[go-reportcard-sheild]: https://goreportcard.com/badge/github.com/blushft/strana
[go-reportcard-url]: https://goreportcard.com/report/github.com/blushft/strana
[go-status-url]: https://github.com/blushft/strana/workflows/Go/badge.svg
[godoc-shield]: https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square
[godoc-url]: https://pkg.go.dev/github.com/blushft/redtape
[license-shield]: https://img.shields.io/github/license/blushft/strana.svg?style=flat-square
[license-url]: https://github.com/blushft/strana/blob/master/LICENSE
[contributors-shield]: https://img.shields.io/github/contributors/blushft/strana.svg?style=flat-square
[contributors-url]: https://github.com/blushft/strana/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/blushft/strana.svg?style=flat-square
[forks-url]: https://github.com/blushft/strana/network/members