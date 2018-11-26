package main

import (
	"reflect"
	"testing"
)

func Test_generateEthereumKeys(t *testing.T) {
	type args struct {
		pageNumber  string
		keysPerPage int
	}
	tests := []struct {
		name     string
		args     args
		wantKeys []ethereumKey
	}{
		{
			"It can generate keys starting from the first page",
			args{"1", 18},
			[]ethereumKey{
				{private: "0000000000000000000000000000000000000000000000000000000000000000", public: "0x3f17f1962B36e491b30A40b2405849e597Ba5FB5"},
				{private: "0000000000000000000000000000000000000000000000000000000000000001", public: "0x7E5F4552091A69125d5DfCb7b8C2659029395Bdf"},
				{private: "0000000000000000000000000000000000000000000000000000000000000002", public: "0x2B5AD5c4795c026514f8317c7a215E218DcCD6cF"},
				{private: "0000000000000000000000000000000000000000000000000000000000000003", public: "0x6813Eb9362372EEF6200f3b1dbC3f819671cBA69"},
				{private: "0000000000000000000000000000000000000000000000000000000000000004", public: "0x1efF47bc3a10a45D4B230B5d10E37751FE6AA718"},
				{private: "0000000000000000000000000000000000000000000000000000000000000005", public: "0xe1AB8145F7E55DC933d51a18c793F901A3A0b276"},
				{private: "0000000000000000000000000000000000000000000000000000000000000006", public: "0xE57bFE9F44b819898F47BF37E5AF72a0783e1141"},
				{private: "0000000000000000000000000000000000000000000000000000000000000007", public: "0xd41c057fd1c78805AAC12B0A94a405c0461A6FBb"},
				{private: "0000000000000000000000000000000000000000000000000000000000000008", public: "0xF1F6619B38A98d6De0800F1DefC0a6399eB6d30C"},
				{private: "0000000000000000000000000000000000000000000000000000000000000009", public: "0xF7Edc8FA1eCc32967F827C9043FcAe6ba73afA5c"},
				{private: "000000000000000000000000000000000000000000000000000000000000000a", public: "0x4CCeBa2d7D2B4fdcE4304d3e09a1fea9fbEb1528"},
				{private: "000000000000000000000000000000000000000000000000000000000000000b", public: "0x3DA8D322CB2435dA26E9C9fEE670f9fB7Fe74E49"},
				{private: "000000000000000000000000000000000000000000000000000000000000000c", public: "0xDbc23AE43a150ff8884B02Cea117b22D1c3b9796"},
				{private: "000000000000000000000000000000000000000000000000000000000000000d", public: "0x68E527780872cda0216Ba0d8fBD58b67a5D5e351"},
				{private: "000000000000000000000000000000000000000000000000000000000000000e", public: "0x5A83529ff76Ac5723A87008c4D9B436AD4CA7d28"},
				{private: "000000000000000000000000000000000000000000000000000000000000000f", public: "0x8735015837bD10e05d9cf5EA43A2486Bf4Be156F"},
				{private: "0000000000000000000000000000000000000000000000000000000000000010", public: "0xfaE394561e33e242c551d15D4625309EA4c0B97f"},
				{private: "0000000000000000000000000000000000000000000000000000000000000011", public: "0x252Dae0A4b9d9b80F504F6418acd2d364C0c59cD"},
			},
		},
		{
			"It can generate keys for the last page",
			args{"904625697166532776746648320380374280100293470930272690489102837043110636675", 128},
			[]ethereumKey{
				// 96 keys on the last page
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364100", public: "0xbbF3316f2Fa21d9e0A8a07F5047F37A467f01a5B"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364101", public: "0xFB7DC16619EdD43a08eFD9cE20De94c94682D13a"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364102", public: "0xd79dc898B43e1404Beb1471D6d566F9F981f118F"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364103", public: "0x3c62B34927d02A4e55379293488Be63cE69b05F8"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364104", public: "0xd259E3B2470098EE45B70673630a1922f638761e"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364105", public: "0x956bcC8D3a53E1A80f729802431D501c23Aa9276"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364106", public: "0xE9A97ebdE002Cdbd4ef86d61cd479f7B36DBb1e7"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364107", public: "0xA036362f74E84039D2702ec4e2c2750aa1F43fEf"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364108", public: "0xf32746816286f894981122e02E2640569f824fC0"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364109", public: "0x279661BD2Cbf7675A51b42Ab08801EB718D99Cc1"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036410a", public: "0x246D6FbCBf9601f55A7e2DaC06eE6BDe84CF2120"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036410b", public: "0x7eC80c82DA721Fed253bB16A0ACBbDb269409E5E"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036410c", public: "0x1F5Ce48feCEEA16759D22c4f52c90204974411FF"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036410d", public: "0x19026841bb9B57587e120CE13FaE9Dd7C20B7F32"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036410e", public: "0xF82867F877acf02EAb87886Ac1F4ffcF48d03962"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036410f", public: "0x76618F5A6eE6C138eBfC2AdF8e92367A9A21d1A6"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364110", public: "0xdb108Df98704cCF44Fe13e25F08B0A0AA230B9A6"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364111", public: "0x9db2dE6864185CdECA7d6709406F3E1acCfFD5dB"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364112", public: "0x7D8D458D014aC223de08d2F80D25438901f15c82"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364113", public: "0x603F312db28F24FEAC8f539dCA8cAb442407D356"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364114", public: "0xFC8705Eff1d89Cc66Cd0B2CaC3FA8c986Ab96EE3"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364115", public: "0x980545ad727dC273B51E0a5352586fA5Cd548683"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364116", public: "0x48442b572C9339923a9BCcBd09612B160CD15849"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364117", public: "0x32D5f8FCD62ffA771b1DB65E7C2211e9DEfD348F"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364118", public: "0x0255b88a30dE3Db1d5b6D63d5343114c6Ce140c4"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364119", public: "0x9D990c3d16241DACb92f36e8E3eAC450eca4935E"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036411a", public: "0xEf2E1F33EbD377B6AcB5470F82A120aC23061E31"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036411b", public: "0xBA5935b3BC656E62158A1077246135d6E1A10Df8"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036411c", public: "0x3fb21F5f512D614328CBe1196177A1Dd80da1e90"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036411d", public: "0x1a7a11C766A414B66F9C4D59a36D7e730E4Bca1D"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036411e", public: "0x233987e78A38D754C44816643e96Ca1e5815dAeA"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036411f", public: "0x6a716064358CDAb0009010E05DC6aF539ab53d8A"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364120", public: "0x95B1fD7b3879CD52ffd36F948AF67166D08cDF11"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364121", public: "0x6687C40EE5F12F7916Db9E2368534Cb0040CF3e4"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364122", public: "0xb419888537465EB564662e4CB5bf2E7400c9ECc7"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364123", public: "0x142110Ba8A897a0212efEea44BF4acB8Ea80462e"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364124", public: "0x7c1e26881DA999Ac729695a47F909DC1BaD2cec0"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364125", public: "0xde0073Ce497e7eAEe5ea97798D823B2D2C723f71"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364126", public: "0x0E2511Dd112A63Cf18c3513B23316e011Afc3afE"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364127", public: "0x0de7755E7475097F42DA221bFb153Eafba2E9F5D"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364128", public: "0xcfF97B2D79Ded7D1dB9502cBF0706935B2a78656"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364129", public: "0x84289d222E4765fFF2Be4e406800Ed4465D4845B"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036412a", public: "0x9bD05480754b3D5816984CAc5E88e60497657199"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036412b", public: "0xB0211d6477BeA0c686Bd6E407eab5cE37aCcA893"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036412c", public: "0x7c51c9A72Cb650e159215A54d6d9D69a69547b5A"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036412d", public: "0x4AB5e175Cdd5B31AA1044D7a7Bba0B90CB9208Cb"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036412e", public: "0xe20307eF6c7b1E5428aC7ca9873dfD1850A147d2"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036412f", public: "0x1475e0534C40F7AAE5DaefB0D2C9Ab58FB01eb8F"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364130", public: "0xD5ea7A94F67d24171b40987f99D26C5DD762596C"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364131", public: "0x44E9F52C16F2b5f232543EDBFC8e9837931D33B3"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364132", public: "0xE8DE258b404F7D5116DB7bFaA1F7F4C8208C2BcD"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364133", public: "0xF66B31d0638d8558c04d75F3F857095e5048F166"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364134", public: "0xAD98c8a3FA5bB03C8C249a0B3e727E8503333Fd2"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364135", public: "0xb67983fE9CCE1EF4fa6E5B339c5FF5B2A9b27395"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364136", public: "0x5c6bD1597b1411cce0e79A0841FD11073120493B"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364137", public: "0xdF8e88eB567f6C901491fDE5636b4bD7611Bd873"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364138", public: "0xBDc7a9D74e7194E279bCde320496dDB314Ac4303"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364139", public: "0x6023eB78B679DAF4f8e14E096e97774f75c5140E"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036413a", public: "0xcA193534a86C4e536722676E3F92E03804A436d0"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036413b", public: "0xdc6999513539883ee37f4f1a0a2Ad573812B6A68"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036413c", public: "0x941171032778e26a70A00Da92b841a6C7fB5b676"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036413d", public: "0xb69f25896e3CFac20C89eC1Ce8866F4eB2828c36"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036413e", public: "0x2Ef1f47E3244806c0FAf4Bd42D96cD1e05AefFeC"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036413f", public: "0x92D48Ff5523c9B04Aa426191b4bD21e6080F074A"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364140", public: "0x80C0dbf239224071c59dD8970ab9d542E3414aB2"},
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
			},
		},
		{
			"It can generate keys for the second to last page",
			args{"904625697166532776746648320380374280100293470930272690489102837043110636674", 128},
			[]ethereumKey{
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364080", public: "0x24c28b494275D758d9eCc27aD4A54EB7a5E19f88"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364081", public: "0xc0Fc35d6B3D9Bc95ab36B55C0d1e429491A4e4D8"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364082", public: "0x5E8dc623b42c4bBA6f804D104d9B41D323bA5f1C"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364083", public: "0x8fb4b8eB836C2b9D72F17B7fB5B3Dc2417a1481d"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364084", public: "0x464bb6ee22d547f9183ed0DfDD06d30c10A54714"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364085", public: "0xaf4Ff8938215795bA24F4533CFaEbCB15FA72d84"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364086", public: "0x8A73654C5b16E2645178420Bf298Bdf5E23535F9"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364087", public: "0x4fC7105Cf594A471077475D28b03174DE27d4fC7"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364088", public: "0x7CCA5C103879204458CAEeed10A09a2629C908f2"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364089", public: "0xD13E9B95F69b5E9553554d79cf1AF5Ca35dda38A"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036408a", public: "0x45eC01EdeC168643DE5999bEEccF394A37DEd4BF"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036408b", public: "0xC961Cd443Eb275f9F6856f3e5E1e1a1DC12fcC89"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036408c", public: "0x3dFfcC0bdd956A419405B675C33E65f944529EE0"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036408d", public: "0x54cE04Df9e8127bbE04c6C4f6d819B1060F5C14e"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036408e", public: "0xfCe9F8ddDdfAB9C6538ff2cD05F76ded16518905"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036408f", public: "0x1b817fA3735C57f59d0c9D3e6e4e1A4e4a6F020a"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364090", public: "0xd7a0Bbd722a302b2B4079abc962d839c6Ad3A7E2"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364091", public: "0x8597d471D161C1606a3C5A8d7f9dbb7DcBa29f64"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364092", public: "0xF6F2a0752CD1162f836ed6C12D7716d4BDe2D45c"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364093", public: "0x5A0A69269F0764cC01539997511418d73DF6Adb9"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364094", public: "0xa6BAD5AC97fdA351a784c38D557014c267596BA3"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364095", public: "0xE703F0102927c5F0Fb828C429e43f91C7549c54D"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364096", public: "0x22F274009d12E0ECc0D9866ce6424910E46DbA39"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364097", public: "0x036f5bE2A77B5a40De057564faC86ccAc1Ecc105"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364098", public: "0xC4Ca0853326D03C00C221780fEA5Bd7b057EAe73"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364099", public: "0xb39e749993E1f4952812Bd2eeF05331D8ea7e04A"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036409a", public: "0x636D9D40b1F0eec51F0af83e0ae8962fE87b0E55"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036409b", public: "0xAb34dF310aefCF51aff94574D19E88B191555be0"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036409c", public: "0xa27229415a22854d4c045eb864210C42c160D4Af"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036409d", public: "0xfCE982a5Ae7918EfC275817E41044e02783DF9b1"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036409e", public: "0x0248479ee063Ed471E6fA743416975096CeDd1Af"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036409f", public: "0x2D729b4897967f05fBC9ABEBE5f3990e79aAbd8e"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640a0", public: "0xf8961751715D7824B636d09d1aeD94db1358070b"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640a1", public: "0x802e217Dae478aCB08F6083Ac9F6978804cc96Fa"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640a2", public: "0x557659b033587f5F88A798543c751522da9A77c3"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640a3", public: "0x2B4701117AC9A9D83F20091A82c193776938b33a"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640a4", public: "0x808632B8C1b6B71474Bf802AC5330338d991E262"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640a5", public: "0x2AEDF2E5E44ff1E921e06258168C20C0297BE0b0"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640a6", public: "0x4c5D822457F031D0D837Af6FD016487Ac808501B"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640a7", public: "0xae92713086df0115dE138A1803c1830ebA24c04c"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640a8", public: "0xf2dDf8403B63E793eefE5776F775218D644A6Ed4"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640a9", public: "0x9Fe41B5675101d215CbD65b952829255F5Fd190B"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640aa", public: "0x8C43dc73bbda7c745acfc4c5fd5DA0Fb8b0789DD"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640ab", public: "0x6dDbC88dD571a3D62354F4697EE94FA2F674F60b"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640ac", public: "0x49Ad143cd3e1A50044A603e52Caf9aAb45F9039e"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640ad", public: "0xf0dF3E487a4DfE1E0c1A8D79Ae9a6D2246F4c132"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640ae", public: "0x6E13371f15a952AA59ce3081832350f29824DD3D"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640af", public: "0xF5e2e4B23f5E402E7bf8cB6B030132843627dabc"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640b0", public: "0x547e8B70462cd11fedF450DbABE27606335208C7"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640b1", public: "0x9a5Cf8caf63325B601eDD16FF3aFE35E19C33D4B"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640b2", public: "0x05e7F54281E28358F9F80d94224119C366921baE"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640b3", public: "0x3A42E3DA23548C1e869D2DC296905D830be2aeDe"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640b4", public: "0x1E4E79D2e4EDE040ed7F4ECfEAf0dCcA009c0Cf1"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640b5", public: "0x6afa631b162c62A63E871cafBDEf97aa8c4F17d6"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640b6", public: "0x894883b7CccB4380886d6F266bd4B31e5116dBB6"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640b7", public: "0xd41a68cD6d5468885Ca4106E9c091053E4B98712"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640b8", public: "0x082144C36517bE5567a211458A2f9CE8dd3C9631"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640b9", public: "0xB3421D04D4CEf251748e52554f7bB6898162457D"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640ba", public: "0x77dc4a011db73b747b6C35831F9A3d708C4Bc06b"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640bb", public: "0xE45eC6A57aF401025780082066fb8b0e852948Ea"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640bc", public: "0x6694e6f4aeb487396d6380e16aE7045A6959a08F"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640bd", public: "0x36826A179bc0F443F3BE335b950E5984fe66A0E9"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640be", public: "0x7C2dAd3C003D0D7f2754240dCde0AE85D8966891"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640bf", public: "0xe3Dd2A563D497B1bd13F0990801F872126EF9619"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640c0", public: "0x3c14170974F780573a8EAA4a898b5f7295af388C"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640c1", public: "0x0436F67BE14Fe01581Ec95cfdd258292B612b1e0"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640c2", public: "0xA8A5D0d6B212D0E3400b2CF5797F9069E9D242f8"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640c3", public: "0xFcDa3687cB8EDDA171bB4E0e77dD8e5c4f0bEB36"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640c4", public: "0x7B70043eA73A741Fbb0308b0f144b9d558F1ee6d"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640c5", public: "0xbf3920FC367E6DeD48205922E6BcAA5094C11984"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640c6", public: "0x63D61B4b88394813a6Eb57F12F7eb814B2C55559"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640c7", public: "0xa8b0137B4993550507b82bDe299769E6FCB06bc0"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640c8", public: "0x800111796Cf8CF1fe19d1A88A8b14e61ABb2F64c"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640c9", public: "0x489db12C6494Cb82f4B4595AF8e1cc03c7aCe7BC"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640ca", public: "0x5dA264B9c694025e493b88a74d24926B6D820046"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640cb", public: "0x6a404a4dDE355Bd5A70DBe2496E41f86BF7195B7"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640cc", public: "0xdACc0CD0A39F5790ebDD89338fD89D997109079B"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640cd", public: "0x49CdA6da93eDA358DdD5bBF9fc9CEdE1bff7CF37"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640ce", public: "0x74997224c265153206A3583389853d4875eAF2f0"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640cf", public: "0x16F961B881c4aa185437dC788411966267743599"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640d0", public: "0xc56bAF7ae68d346E1C3e1E4FADACA482922C689D"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640d1", public: "0xc046AcF78Fb69A459509C88c22bb589C3dA47f1E"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640d2", public: "0x93dD9E9C613A612294F66F86d6B4f745Fb9613e3"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640d3", public: "0xcB1FA10AC9BD63ED02f6c4c208A6ed4C38DDd2F7"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640d4", public: "0x132481670A563Ebc85F1780aD0ce3aD7C8Ec8550"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640d5", public: "0x6deDf2f67d8D2416bA27529465Ac3799A685cD94"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640d6", public: "0x147ecbA866a76C06e5Df312464820B3F0ac22336"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640d7", public: "0xf7fE46bb9DCb35e84f862B3f0f29ad88f98729D3"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640d8", public: "0x01C45FF9b24b1A30c7a67a47AF35f20Ee6E861Df"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640d9", public: "0x08b677c1702f893d7af551CcF0aB73Cc76b08725"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640da", public: "0x9c9B8ad417948F4992917604D5e474A0985Da72B"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640db", public: "0xDe4115CEd57a730a98Aeb51262272F70176D517c"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640dc", public: "0x2376A74C7e27734513232fCcB9BF4987FFd32828"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640dd", public: "0x85C42EBbc874B4d7Bd41B422ba52cba74005a708"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640de", public: "0x64499795Be8f25f7835355FAd147ef4bE67A63eA"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640df", public: "0x5e217cedC12065090504E6fefFfD4F34b5392C96"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640e0", public: "0x0046d140192095d0d5a496A1c6aeFe505cB97b25"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640e1", public: "0x803E207333090038357df8850B6505082C22202c"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640e2", public: "0x1D008738A1a2edDDCCc647414a012952C7aaa703"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640e3", public: "0xc24e4Cb2cD7aF8652014A5895d718DE3989Ec908"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640e4", public: "0x3534CE8B77D3a6763F621eE5E695AC1F069939C1"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640e5", public: "0xbFc243E6A1AD1C0849619836803fD4ec69aA3627"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640e6", public: "0x1EB41699Bb4Cf60562016b2feA19F892b747bf7e"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640e7", public: "0x7F9b7b09b2FD7dFfA4cE34eEf655fEED77C262Cf"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640e8", public: "0x17C2Be1182AdC7856927a874483d34872941e26A"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640e9", public: "0xCacd403c3A66ee51D70D988C3a205a8B55632c08"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640ea", public: "0xE4ce5B9B8eb22168490456c04fEc9BF8B9351226"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640eb", public: "0xdC283e9cE10796590409904014A6ca5A652A3fe7"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640ec", public: "0x6c670CBdD33A15AfCA8A123ddafD2E5bd076b320"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640ed", public: "0xc632AE04355AF36c642b2Db29e37fA2E47290B46"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640ee", public: "0xE847f260CA14716206Bf35A855Eec2A231Ce4829"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640ef", public: "0x0BF649ce3f9c67c700f35019495EC129B2adC25f"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640f0", public: "0x53E2BBeF6240F9c4f7DDA12B8ca548d51d9faa5D"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640f1", public: "0xeE3744837d7e2bC55061a90cd7eE9f2249CA83EC"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640f2", public: "0xb8bc3965F623dc485DA38f3CA5dBcFF905911C14"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640f3", public: "0x4fBA73b1cA808602e2472A91eA70CAa3C1d599e0"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640f4", public: "0x9cc1bA4282517CE299a50C024e3D3b2e6062c299"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640f5", public: "0x4430BaeD72189129eAa9c6D521B601Fe4eb21E86"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640f6", public: "0xbF6f1e7389dcBf98F36C30643591B3556e4A862E"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640f7", public: "0xa96020783A4366Fa30E62322Ae31113C5ceEE49D"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640f8", public: "0x9694dc44cF34bFE044f13C1E553fC51b9c4e9F17"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640f9", public: "0xB9ea59981494040e4E01c79F2aE94c0A3F306C7C"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640fa", public: "0x31A22784aB2Eaae663390BD3876cA744f8993189"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640fb", public: "0x47236b8BEbBFe5d8555eF87b10463509C0284272"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640fc", public: "0xDdE3c33431448B7e1053Fe20dCD1C2fc61e9cf1B"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640fd", public: "0xB4eFf29840cDcfF388890D0C923a7b76F66e24c6"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640fe", public: "0xf179d1bAAdF6B4f045aFf53F61783B6DC5012035"},
				{private: "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd03640ff", public: "0x02396f902B7C3aE09AE37155f7287a4a3F498f66"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotKeys := generateEthereumKeys(tt.args.pageNumber, tt.args.keysPerPage); !reflect.DeepEqual(gotKeys, tt.wantKeys) {
				t.Errorf("Expected:")
				for _, expectedKey := range tt.wantKeys {
					t.Errorf("%#v", expectedKey)
				}

				t.Errorf("Actual:")
				for _, actualKey := range gotKeys {
					t.Errorf("%#v", actualKey)
				}
			}
		})
	}
}

func Test_findEthPrivateKeyPage(t *testing.T) {
	type args struct {
		privateKey  string
		keysPerPage int
	}

	tests := []struct {
		name     string
		args     args
		wantPage string
	}{
		{
			"It can find the page that a random private key is on",
			args{"e44a4bdc91d35496190474dca11338059ffbab72d3a72f195c4a030632d49503", 128},
			"806707810447654934665982607721811039665835129933115297248770801612223785259",
		},
		{
			"It can find a key on the first page",
			args{"0000000000000000000000000000000000000000000000000000000000000001", 128},
			"1",
		},
		{
			"It can find the first key on the first page",
			args{"0000000000000000000000000000000000000000000000000000000000000000", 128},
			"1",
		},
		{
			"It can find a key on the last page",
			args{"fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364140", 128},
			"904625697166532776746648320380374280100293470930272690489102837043110636675",
		},
		{
			"It can find the last key on the last page",
			args{"fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd036415f", 128},
			"904625697166532776746648320380374280100293470930272690489102837043110636675",
		},
		{
			"It can find a key beyond the last page",
			args{"ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 128},
			"904625697166532776746648320380374280103671755200316906558262375061821325312",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPage := findEthPrivateKeyPage(tt.args.privateKey, tt.args.keysPerPage); !reflect.DeepEqual(gotPage, tt.wantPage) {
				t.Errorf("Expected: %v", tt.wantPage)
				t.Errorf("Actual:   %v", gotPage)
			}
		})
	}
}
