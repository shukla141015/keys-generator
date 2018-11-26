package main

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

type ethereumKey struct {
	private string
	public  string
}

var hardcodedEthereumLastPageKeys = []ethereumKey{
	{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", public: "0x3f17f1962B36e491b30A40b2405849e597Ba5FB5"},
	{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364142", public: "0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf"},
	{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364143", public: "0x2B5AD5c4795c026514f8317c7a215E218DcCD6cF"},
	{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364144", public: "0x6813Eb9362372EEF6200f3b1dbC3f819671cBA69"},
	{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364145", public: "0x1efF47bc3a10a45D4B230B5d10E37751FE6AA718"},
	{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364146", public: "0xe1AB8145F7E55DC933d51a18c793F901A3A0b276"},
	{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364147", public: "0xE57bFE9F44b819898F47BF37E5AF72a0783e1141"},
	{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364148", public: "0xd41c057fd1c78805AAC12B0A94a405c0461A6FBb"},
	{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364149", public: "0xF1F6619B38A98d6De0800F1DefC0a6399eB6d30C"},
	{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036414a", public: "0xF7Edc8FA1eCc32967F827C9043FcAe6ba73afA5c"},
	{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036414b", public: "0x4CCeBa2d7D2B4fdcE4304d3e09a1fea9fbEb1528"},
	{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036414c", public: "0x3DA8D322CB2435dA26E9C9fEE670f9fB7Fe74E49"},
	{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036414d", public: "0xDbc23AE43a150ff8884B02Cea117b22D1c3b9796"},
	{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036414e", public: "0x68E527780872cda0216Ba0d8fBD58b67a5D5e351"},
	{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036414f", public: "0x5A83529ff76Ac5723A87008c4D9B436AD4CA7d28"},
	{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364150", public: "0x8735015837bD10e05d9cf5EA43A2486Bf4Be156F"},
	{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364151", public: "0xfaE394561e33e242c551d15D4625309EA4c0B97f"},
	{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364152", public: "0x252Dae0A4b9d9b80F504F6418acd2d364C0c59cD"},
	{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364153", public: "0x79196B90D1E952C5A43d4847CAA08d50b967c34A"},
	{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364154", public: "0x4bd1280852Cadb002734647305AFC1db7ddD6Acb"},
	{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364155", public: "0x811da72aCA31e56F770Fc33DF0e45fD08720E157"},
	{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364156", public: "0x157bFBEcd023fD6384daD2Bded5DAD7e27Bf92E4"},
	{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364157", public: "0x37dA28C050E3c0A1c0aC3BE97913EC038783dA4C"},
	{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364158", public: "0x3Bc8287F1D872df4217283b7920D363F13Cf39D8"},
	{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364159", public: "0xf4e2B0fcbd0DC4b326d8A52B718A7bb43BdBd072"},
	{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036415a", public: "0x9a5279029e9A2D6E787c5A09CB068AB3D45e209d"},
	{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036415b", public: "0xc39677F5F47d5fE65ab24e66750e8FCa127c15BE"},
	{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036415c", public: "0x1dc728786E09F862E39Be1f39dD218EE37feB68D"},
	{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036415d", public: "0x636CC65783084b9F370789c90F733DBBeb88925D"},
	{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036415e", public: "0x4a7A7c2E09209dbE44A582cD92b0eDd7129E74be"},
	{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036415f", public: "0xA56160A359F2EAa66f5c9df5245542B07339A9a6"},
}

func generateEthereumKeys(pageNumber string, keysPerPage int) (keys []ethereumKey) {
	basePage := new(big.Int).Sub(makeBigInt(pageNumber), one)

	// convert the "int" to "string" because i dont know how to create a bigInt from an "int"
	stringInt := fmt.Sprintf("%d", keysPerPage)

	firstSeed := new(big.Int).Mul(basePage, makeBigInt(stringInt))

	ethereumKeys := make([]ethereumKey, 0, keysPerPage)

	for i := 0; i < keysPerPage; i++ {
		// convert the seed to hex and left-pad it with zeroes until its 64 chars long.
		privateKey := fmt.Sprintf("%064x", firstSeed)

		var publicKey string

		if privateKey == "0000000000000000000000000000000000000000000000000000000000000000" {
			publicKey = "0x3f17f1962B36e491b30A40b2405849e597Ba5FB5"
		} else {
			key, _ := crypto.HexToECDSA(privateKey)

			publicKey = crypto.PubkeyToAddress(key.PublicKey).Hex()
		}

		ethereumKeys = append(ethereumKeys, ethereumKey{
			public:  publicKey,
			private: privateKey,
		})

		// this is the last seed that the ethereum crypto package can generate a public key for,
		// a seed higher than this will crash. There are more valid addresses after this seed,
		// they are hardcoded in this file.
		if privateKey == "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364140" {
			return append(ethereumKeys, hardcodedEthereumLastPageKeys...)
		}

		firstSeed.Add(firstSeed, one)
	}

	return ethereumKeys
}

func findEthPrivateKeyPage(privateKey string, keysPerPage int) string {
	hex := strings.TrimLeft(privateKey, "0")

	if hex == "" {
		return "1"
	}

	baseBigInt, _ := new(big.Int).SetString(hex, 16)

	kppBigInt, _ := new(big.Int).SetString(fmt.Sprintf("%d", keysPerPage), 10)

	divided, _ := new(big.Int).DivMod(baseBigInt, kppBigInt, kppBigInt)

	finalBigInt := new(big.Int).Add(divided, one)

	return fmt.Sprintf("%d", finalBigInt)
}
