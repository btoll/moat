# moat

`moat` is a command-line tool that generates a password of configurable length using consisting of either a `Diceware` passphrase or a more traditional password of alphanumeric and special characters.

## What is Diceware?

From the [official Diceware site][diceware]:

> Diceware™ is a method for picking passphrases that uses dice to select words at random from a special list called the Diceware Word List.

This project uses the [new word list] from the [EFF], which contain many improvements over the original Diceware word list.

## Examples

```bash
$ moat diceware
regulatorvotinggauntlettreadingwidenuntried
```

```bash
$ moat diceware --delimiter " "
cozily troubling flask trapezoid waving throwback
```

```bash
$ moat diceware --delimiter "-"
ablaze-marmalade-axis-appetite-skyward-enigmatic
```

```bash
$ moat diceware --words 10 --delimiter "-"
that-giddily-repressed-setback-musty-predict-recede-bagging-pointer-uncurled
```

```bash
$ moat sillypass
WRdmdvJC^=6C
```

```bash
$ moat sillypass --length 30
6nPBoeqfh+.vOeNS?B2ssZ2tajlI<4
```

## Installation

```bash
go install github.com/btoll/moat@latest
```

## CLI Usage

```bash
moat is a passphrase generator

Usage:
  moat [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  diceware    Create a Diceware passphrase
  help        Help about any command
  sillypass   Create a sillypass password

Flags:
  -h, --help   help for moat

Use "moat [command] --help" for more information about a command.
```

## License

[GPLv3](COPYING)

## Author

Benjamin Toll

[diceware]: http://world.std.com/~reinhold/diceware.html
[new word list]: https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases
[EFF]: https://www.eff.org/

