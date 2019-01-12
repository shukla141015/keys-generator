# Keys.lol generator
This repository contains the key generator for [Keys.lol](https://keys.lol)

The code for the website is in a [separate repository](https://github.com/SjorsO/keys)

## Building and installing
1. cd to `~/go/src/github.com/sjorso/keys-generator`
2. install required packages with `go get`
3. build the executable with `go build`
4. include the executable in `$PATH`: `sudo cp keys-generator /usr/local/bin`

## Usage
For generating keys, run:

```bash
keys-generator btc <page number>
keys-generator eth <page number>
```

For searching by private key, run:
```bash
keys-generator btc-search <btc private key>
keys-generator eth-search <eth private key>
```

## License

This project is open-sourced software licensed under the [MIT license](http://opensource.org/licenses/MIT)
