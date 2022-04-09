---
title: "Go Fast With Go: An Introduction to Go"
author: NeoConf, Jordan Brauer
date: 2022-04-09
---

# Go Fast with Go: An Introduction to Go

_~25 minutes_ \_\_φ(．．)

A brief intro to the Go programming language and some of its applications,
benefits, and niceties.

```go
package main

func main() {
    println("Hello, 世界")
}
```

---

# (・\_・ヾ What Exactly _is_ Go? 

> I've heard of Go, but never used it or looked into it

Created by [Robert Griesemer](https://en.wikipedia.org/wiki/Robert_Griesemer),
[Rob Pike](https://en.wikipedia.org/wiki/Rob_Pike), and [Ken Thompson](https://en.wikipedia.org/wiki/Ken_Thompson)
at Google, born out of frustration with existing languages and the trade-offs
they came with when solving modern problems.

- Open source and supported by Google;
- Imperative, statically typed, compiled language;
- Simple, light-weight, and easy to learn;
- Built-in concurrency with a robust standard library;
- Memory safe and garbage collected (no memory management required!);

---

# (￢\_￢) What Can Go Do?

> Yeah, yeah, yeah, but like ... what can I make with it?

Pretty much anything you can think of!

- Networking, cloud services, infrastructure & devops;
- CLI programs;
- Web servers;
- Games (2D & 3D) programming;
- Systems programming;


## Who Uses Go?

- Kubernetes, Dgraph, Turborepo, esbuild;
- YouTube, Netflix, Twitch;
- PayPal, AmericanExpress, CapitalOne;
- CloudFlare, DropBox, Bitly, Twitter, Uber, Microsoft, Salesfore, Riot Games;

---

# ( ˘▽˘)っ♨  Go's Best Features

> Okay, but how's that different than `<insert my favourite language here>`?

Yes, but does `<insert your favourite language here>` do **_this_**?!

- _Blazingly_ fast cross-platform compilation;
    - Produces a single statically linked binary;
- Built-in language tools for just about everything;
- Standard library does (almost) everything;
- Interfaces and composition in the form of _type embedding_;
- Built-in concurrency primitives;

---

# ヽ(~\_~(・\_・ )ゝ What Go _Isn't_

> What about ...?

No, but...

- Called Golang or GO – simply _Go_;
    - Stylized as _GO_ in the logo, and easily searched by _golang_;
- Object-oriented or functional (kind of);
- Only used by Google;

---

# (/\_＼) Demo Time

_woohoo_!

1. Functions, variables, & constants;
2. Interfaces, structs, & methods;
3. Go routines & channels;
4. Sample app;

---

# \_\_〆(￣ー￣ ) Resources

Some of the (in my opinion) best Go resources!

## Reference Documentation

Surprisingly, the official documentation is actually good.

- [The Go Website](https://go.dev/)
- [Language Spec](https://go.dev/ref/spec)
- [Official FAQ](https://go.dev/doc/faq)
- [Gopher Model Sheet](https://go.dev/doc/gopher/modelsheet.jpg)

## Tutorials & Guides

If you like interactive learning, these are for you.

- [How to Write Go Code](https://go.dev/doc/code)
- [A Tour of Go](https://go.dev/tour/welcome/1)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)

---

_thank you_!
