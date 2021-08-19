# spartan-cli

⚔️️ SPARTAN Framework CLI

> This is an ALPHA version, please use at your own risk.

## Instalation

Download the latest version from the `Releases` into your `/usr/local/bin`.

TIP: (use this short version to help calling the command)

```bash
echo "alias spa='/usr/local/bin/spartan'" >> ~/.bashrc
```

> All following examples will use the short version `ign`

## Usage

### Create new Project

```bash
spa new <projectname>
```
or
```bash
spartan new <projectname>
```

extra flags:

* `-p /Users/DarthVader/Documents` - Specificy a path to instalation
* `-m github.com/DeathVader/DeathStar` - Specifiy a module name to be added to `go.mod`
