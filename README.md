# spartan-cli

⚔️️ SPARTAN Framework CLI

> I need a predifined structure that allows me to easy understand where things
> are and how they work.

> Building applications is hard and complex. To help understand and navigate any
> project you need for it to have good structure that allows you to make sense
> of how things are organized and give you a good set of tooling ready to go
> (pun intended).

> Spartan CLI is a scafolding tool that allows you to quick generate base code
> for packages, services, handlers, entities, etc. Event though it uses a
> predefined set of known modules, it tries to minimize the dependency on
> them and they are easly replaceble by another modules if needed. Doing all
> this while trying to follow some good architecture standards, mainly
> "Clean Architecture Go" described by [Elton Minetto](https://twitter.com/eminetto)
> in this 2nd posts [Clean Architecture, 2 years later](https://eltonminetto.dev/en/post/2020-07-06-clean-architecture-2years-later/)
> and [Clean Architecture using Golang ](https://eltonminetto.dev/en/post/2018-03-05-clean-architecture-using-go/),
  plus a few things that make the adventure more easy.

> IMPORTANT: This is an ALPHA version, please use at your own risk.

## Instalation

Download the latest version from the `Releases` into your `/usr/local/bin`.

TIP: (use this short version to help calling the command)

```bash
echo "alias spa='/usr/local/bin/spartan'" >> ~/.bashrc
```

> All following examples will use the short version `spa`

## Usage

### Create new Project

```bash
spa new -n <projectname> -m github.com/<username>/<projectname>
```
