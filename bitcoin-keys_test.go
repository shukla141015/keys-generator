package main

import (
	"reflect"
	"testing"
)

func Test_generateBitcoinKeys(t *testing.T) {
	type args struct {
		pageNumber  string
		keysPerpage int
	}

	tests := []struct {
		name     string
		args     args
		wantKeys []key
	}{
		{
			"It can generate keys starting from the first seed",
			args{"1", 10},
			[]key{
				{private: "5HpHagT65TZzG1PH3CSu63k8DbpvD8s5ip4nEB3kEsreAnchuDf", compressed: "1BgGZ9tcN4rm9KBzDn7KprQz87SZ26SAMH", uncompressed: "1EHNa6Q4Jz2uvNExL497mE43ikXhwF6kZm"},
				{private: "5HpHagT65TZzG1PH3CSu63k8DbpvD8s5ip4nEB3kEsreAvUcVfH", compressed: "1cMh228HTCiwS8ZsaakH8A8wze1JR5ZsP", uncompressed: "1LagHJk2FyCV2VzrNHVqg3gYG4TSYwDV4m"},
				{private: "5HpHagT65TZzG1PH3CSu63k8DbpvD8s5ip4nEB3kEsreB1FQ8BZ", compressed: "1CUNEBjYrCn2y1SdiUMohaKUi4wpP326Lb", uncompressed: "1NZUP3JAc9JkmbvmoTv7nVgZGtyJjirKV1"},
				{private: "5HpHagT65TZzG1PH3CSu63k8DbpvD8s5ip4nEB3kEsreB4AD8Yi", compressed: "1JtK9CQw1syfWj1WtFMWomrYdV3W2tWBF9", uncompressed: "1MnyqgrXCmcWJHBYEsAWf7oMyqJAS81eC"},
				{private: "5HpHagT65TZzG1PH3CSu63k8DbpvD8s5ip4nEB3kEsreBF8or94", compressed: "17Vu7st1U1KwymUKU4jJheHHGRVNqrcfLD", uncompressed: "1E1NUNmYw1G5c3FKNPd435QmDvuNG3auYk"},
				{private: "5HpHagT65TZzG1PH3CSu63k8DbpvD8s5ip4nEB3kEsreBKdE2NK", compressed: "1Cf2hs39Woi61YNkYGUAcohL2K2q4pawBq", uncompressed: "1UCZSVufT1PNimutbPdJUiEyCYSiZAD6n"},
				{private: "5HpHagT65TZzG1PH3CSu63k8DbpvD8s5ip4nEB3kEsreBR6zCMU", compressed: "19ZewH8Kk1PDbSNdJ97FP4EiCjTRaZMZQA", uncompressed: "1BYbgHpSKQCtMrQfwN6b6n5S718EJkEJ41"},
				{private: "5HpHagT65TZzG1PH3CSu63k8DbpvD8s5ip4nEB3kEsreBbMaQX1", compressed: "1EhqbyUMvvs7BfL8goY6qcPbD6YKfPqb7e", uncompressed: "1JMcEcKXQ7xA7JLAMPsBmHz68bzugYtdrv"},
				{private: "5HpHagT65TZzG1PH3CSu63k8DbpvD8s5ip4nEB3kEsreBd7uGcN", compressed: "1HSxWThjiwbC4dJbXHMpBfwRenB12UguG5", uncompressed: "1CijKR7rDvJJBJfSPyUYrWC8kAsQLy2B2e"},
				{private: "5HpHagT65TZzG1PH3CSu63k8DbpvD8s5ip4nEB3kEsreBoNWTw6", compressed: "13DaZ9nfmJLfzU6oBnD2sdCiDmf3M5fmLx", uncompressed: "1GDWJm5dPj6JTxF68WEVhicAS4gS3pvjo7"},
			},
		},
		{
			"It can generate keys for the last page",
			args{"904625697166532776746648320380374280100293470930272690489102837043110636675", 128},
			[]key{
				// 64 keys
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqemizF9vA", compressed: "12d8ggXP5MSJoEuqtRJqyZpLxqUAztmrpH", uncompressed: "1PDSZN2qgFcuay1vVRxYo1yp9gfXeSKJgt"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqemmsbvAo", compressed: "1Et3i5Bjbn5cLbqwngT3HeSxQG3sXyvC7L", uncompressed: "1EsDryguZoanBraPCYCk9bUoynfY6PoNvj"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqemtri4qS", compressed: "1MJu8dVQVRx6AeSLuHA9avGQALqxGFB4iw", uncompressed: "1FKs7XQkQS5MHqEmFeKmx9vhpBRNYUxBn5"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqemyyKQGD", compressed: "16qmCy9t35haJZfbnq4PkXfeKMjNQvx5h", uncompressed: "1GkQuui5ofmtJMnQvrMzVs3Rw2qnBj8Hms"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqenBBbpLQ", compressed: "1DHQUMNsRoZiCpcd7PhmHgrQvDUPGwGptK", uncompressed: "13bFiKHMPA6ydmC4jctqqdvRNPHq8JLhQc"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqenEc4n5A", compressed: "1MiUJRU3fSSgvSeF56BjMZaFHjcWmZS8w", uncompressed: "16gK4BTErckvm22uqTAcbztsEzdRq4JwT4"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqenQFmCR7", compressed: "1HWyLvUVJvkwmFgF2SvPkhHA5ttRhjGR1h", uncompressed: "1QeXZe66ay57kpkjxT6ydcpRe5J1TA997"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqenVqqaRt", compressed: "1DPinkkKGeh4B5Qynr1aHwfGPz38BBCd67", uncompressed: "18FeyYSiZBLvsSuKVtwDugRCvvtVU4t4LE"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqencAZErF", compressed: "1BPdPPj9jgtHT3usnF8AizRfnbXVVFPVDT", uncompressed: "126uVWnkbykXpUzNEuk7erFyuMYaePSWoV"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqenfnqBWw", compressed: "1CvKupTzRqsDi5Zf4QdbVYhmaQUkF667hM", uncompressed: "1Cud5ZFu44376mtdGytFVQxoXZFsAf396W"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqenripixH", compressed: "1J5kaAUPpLZor6UVkTeJYtBojgXrtWsknv", uncompressed: "129Zk4KrdjCtTkPDKDA9yKEoyQMKg7nnY4"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqenwBMVJd", compressed: "1MYwjHMGQZjWFYnXmWMMnaueyn6fHCYL6L", uncompressed: "1GahK7oUFETxTRp1tpcHt6EXchCerop1sj"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqeo3RQvkv", compressed: "18HrrgAZ3csXpqJevSembrbCv43UR2LoTo", uncompressed: "1BTcZcviXTJSoHRxaQZwvPWeUCwJMqj9id"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqeo6TidXi", compressed: "19kvXX4hHGF9cTJmomrN6CBePfEhpKDRWP", uncompressed: "1MadvbXBmgUo18XiwiS7w3nx4gPyHbGiqL"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqeoCq3htf", compressed: "15av1HesW2XF4hs8XP9aNGjezNnuJa3pjW", uncompressed: "1DYHVPZKncADbTRUyqQ6vLzAzotJBBdNVZ"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqeoQJAair", compressed: "15pTbF1pm6oEDHEMW4K1TUv3xdMUodTfWu", uncompressed: "1CaZUpjd7VmsyWDFrk9WG9nTYMLcLLvvCw"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqeoUG67kV", compressed: "158eqSFXqk53iyMnMZoENAE2o965Fe4dHy", uncompressed: "1DdXcnmYs4zWryEvfXJJWqu86T4DbQD2a8"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqeodFUkKT", compressed: "1CqJeCZBiLkB3bSgGiRoEURSj6LGqVqqRg", uncompressed: "1DZSj1cyJbhCzgz1UgTvPZHRZVvoGyDUAX"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqeog8LP2s", compressed: "1Ksi2xmc9vi5wnxNdxKkcs3pmsLQakoBBF", uncompressed: "14jo3BJqdNzVJcz4YrF4EMYvHSGgwdYKYY"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqeoqicdaf", compressed: "1MydgvXarZNjDs8Nzh5SkR4LsJbSszAEEU", uncompressed: "1DBXK2tjeJXdy128r6yhBqET55wPqSGSvc"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqeovMf19o", compressed: "1968U6xwiis6ipAaE4uP7H985Sg2xPtPiL", uncompressed: "1237sbJWPKg2MdZzuSqRaqEaaaKGXA1Cou"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqep2apkpy", compressed: "14jbc7bNhF94oiWX5p8dSHP6UkyPhysYW4", uncompressed: "1KnvaEg8NFdeRY3GjcUZm1NoVq8PcdNLcV"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqep9EVtUJ", compressed: "14csyS7vQKBLUn9Am1HHEu1ZfaLd3L6VgQ", uncompressed: "1JcTeDgX1dVMwiW9DN61Gt1x4U7rbLrQAs"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqepEKAcje", compressed: "1LoNTxsB9bGXRqRBLqbKwNz9yzF39amkiC", uncompressed: "14gXLdfh2sqCnuMvDEPMpAzgAMRt9iDPT9"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqepQhCqTK", compressed: "19z7VNJxr5bfEiKLWot8B2rnMe2uMazX2F", uncompressed: "1NPNXYwZXjHHUmNZ5yjGCRXycWbJSduFfz"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqepToGvwg", compressed: "1CBotfPmWKCTP7qEB63nB4D6cSbsBv8qXn", uncompressed: "1FPQXEjTh5RAfnJeNAPyv5xfEwTNVbGTHn"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqepaaXPow", compressed: "16B4ucETeu2MwKxm5WxxzbTLfBBrj1NRGZ", uncompressed: "1GHKXCXPYhyJpPizgewyt22e67gcBben5Y"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqepktCsqr", compressed: "19RQEGMBaKNGGQNnftS1VeaHEQSo7iv9NC", uncompressed: "1HskyAaSKozKoQ3YzretZsBzhcFGusoMiU"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqeptR3dcN", compressed: "1NsUwAZiojCb9ufDiLoTiijBFy2UvAj4tp", uncompressed: "1NtcfxvCYX76JR1cqp61ToAmtXzwYxatQb"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqeq1Uaa93", compressed: "17tBvAGvVofr253SMc1H2Y4MALpvK8nqdV", uncompressed: "1GtuZvcbKw1TsDkCQhDSvyvnfqnBiZHZqk"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqeq35A1TA", compressed: "1LFEfrvJb2tsaDm5h94AxDK77X1dqfdnEu", uncompressed: "1Gqw7w79Mb69tUFFRGpHHwtPkY6nEKint3"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqeqBRruvU", compressed: "1DDNFwqNX3m5kVWRg1ePAqoJngAavxBmnM", uncompressed: "1DeRhsAca8qjL6Qk2E2d9L3wopD6MLWcDy"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqeqGkUfLX", compressed: "12eqsJftjhW21MkA4AMqFntXUfMVkTzDvz", uncompressed: "1yqGdk4DoEd1xiN3bhFdzf6ZGPxiymDvX"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqeqPropqU", compressed: "1bD9vUMnriNiAE9yRRuxPs8cZ6FnF5mTz", uncompressed: "1E4NhgXkqpZvnZPBZwynzNmTAcgKLiPUre"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqeqVd2ZnE", compressed: "146WXvKwWuLH8xrFB7esbMHrN37qYYrWpN", uncompressed: "1LJ5utuGegyKa6YbVTtxzZndFnSdHNzC5C"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqeqaKfzBW", compressed: "1HrRcmLhkirJeYcyKoYq5uJc2Zeb94E1vz", uncompressed: "1KVPFr2XEwessL8J4zmqi5yFqD2BYVJ2Dk"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqeqhYRAUZ", compressed: "125wpFbHQrdRFuLvesEY96RavHM9T1yTFF", uncompressed: "1NMiUoStJMYxfWw6APLRoMs24Fqsr9tmg7"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqequMbUK2", compressed: "1PdyW6CsrYbfcQsLNvp2BprLjvPTDYvBo3", uncompressed: "12GQjWsXZ7rfRYCR4E5bHMg8AkoSxmPBox"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqeqy1cgUY", compressed: "1HYXgyq17sNtGPVsrakdE2bfW1Hu1qpCq", uncompressed: "14AJuXrdKFD8RzVtsF89FYVN4DSmb9xEPf"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqer5piUXL", compressed: "1LnSLVs7CGEQBbY6w2J7kMAU8P48rA8zKQ", uncompressed: "1N1sRyurQe7YouraPgh4rxV8JfARdv7zAH"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqerBWAT6z", compressed: "1LihsAYyVCsfuXXQaQfexNtHDCSWVzu2QQ", uncompressed: "12gG3cNVexUjXCY3KqHi891Kiafsb8AaBy"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqerKGMy47", compressed: "1BMNiGzCvpAMQGTn7NSFPTUtjwSNeB27nP", uncompressed: "16oS9HkwfDmrCSGkaFe7KDQgkMFy5GXFoc"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqerQMuaXp", compressed: "1qapNkhu4ARLB2VvhjiRzoQUQdBedWx69", uncompressed: "1E7rN6ZJ7g6mHYEZ643bJSFXSkwLw6Zzam"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqerai1D7C", compressed: "1WZ1qft3wFmk8QP4dfUSqpyC4JEUiV1FR", uncompressed: "18yhGBghaycjg3UhR2fiquffntYQpUGDE7"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqerf2U3AK", compressed: "12yHuvGnsJbAEgvqajjPdCve91Aa294AHt", uncompressed: "17QPbFArTP6M6QRg2ZE18D3fvzZYxnRUSb"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqeromqrGX", compressed: "1EmghU6CBBfw1wyJqguXeWtjUhW3kmzwbU", uncompressed: "1Fp1zhPoKnKfm8MLYkQZ33GZRbJpE9inpB"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqertvRkec", compressed: "19CEpYsRwMirXiFFSM7daVxtwALrERMaWf", uncompressed: "12TqhXBmGoaaJoudt1MdysYb2JqGWUGoL1"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqes1UC6J5", compressed: "13wyRkVE4XGmNW3g2xgA2SGpKysDtjy1Ka", uncompressed: "14X7DSjXSQBqvFVshZuNwVW6GZyNp79AjF"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqes6XLxKo", compressed: "1AaoXdKGqj5bHoFAUSLwfv5C2CkAi5RjFE", uncompressed: "1ADGZZSKRqz3ydkn714Qzw1FJSbUZZGEr1"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqesAbJPSv", compressed: "1EsZ8f9hGrd9cH35gWLuKbP3J793rArBSt", uncompressed: "1E1oVu22jUEvmQTFDy9bTgabSfmns6fQFY"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqesHpSgWj", compressed: "1Lj2EgsaunRNwsyEK32ebjofbu1tPxxtEy", uncompressed: "1KWhn5gquQvXyXp9BMgJ6HYfNwpHZDmJ5c"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqesTFJEbX", compressed: "12sQJfPVt5YuAbmDCWnym5tNDfshtpBXhB", uncompressed: "122Vo9PeKd4j8zSGBeQHdmks6GnkpycXNz"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqesYcYp9K", compressed: "1FyqVysQjVwyQatuoop3ByZYPecUhj6bnr", uncompressed: "1PMB9Etp3xaDKxpmofy1MmjJF1kvCtH8UA"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqesi68P9B", compressed: "1DNv5wVZKZvFp5gktKtN83ZwEfcQt8oKac", uncompressed: "1zrbUnLczbHkA6pzXuZDD6jNsoKMqGBcy"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqesmCC6YY", compressed: "1Aea8LKoEEWpPqTqaSwRYfksmUScVqV1F6", uncompressed: "18PUeum1Su423DmV2jEGdSd3ewiPfsZZ7z"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqestnHCQU", compressed: "1Nk4wGvaSinFVdrnMfEexLDnBZvWPY393C", uncompressed: "1GLiZZVt326aA8JHG2dEJHC591DXDQNKTs"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqet3sujS6", compressed: "1A81LWBrirUNAKpUVFS37xWT4GAMYU5qgD", uncompressed: "1MFyofP8SVtsEYDHQbZg7XJgfDeSP4ysPm"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqet8uM8zj", compressed: "1BJYFk5827oeYipArjTvLL7JdR4ivCGFYj", uncompressed: "1XunvtCGpmb7uw9qxWwaZFfHNFdUmuMVG"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqetEoeLmv", compressed: "1Lvxa3uJyPyRLbrNpGx761aSDWrJ77aTNm", uncompressed: "1J2zofmGpMUSaNGdTZEhMRYXdWsBQFMpS"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqetNQLySX", compressed: "1JSVicNeasrtuiDpb6r4J5fWxjfdU7ZyWT", uncompressed: "1LWBSfTeaLRNS1vyGSKy2BVW2nd6W9sk8Q"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqetVTGEAr", compressed: "1FjMR9gvnmZ3JYMxBbyc3aZK717b5txJoC", uncompressed: "1F3zbGb5JLBnmCAAYjCCv35zkggrXfi8LR"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqetbh69Dr", compressed: "1HjFHBmhUQkKntPPeWmiLiNGewRAMQWNYs", uncompressed: "15K4QVHD5T1KvW4it56qNuGJoTGMpUaFMj"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqetd9ZKJ4", compressed: "1NjSB7UL4MtdjmPbTUfaHne9R5C2YGxUSA", uncompressed: "1Knh2eFMtzMEtmvGHW14ELG8F9Ny6jV4s3"},
				{private: "5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqetqj84qw", compressed: "1GrLCmVQXoyJXaPJQdqssNqwxvha1eUo2E", uncompressed: "1JPbzbsAx1HyaDQoLMapWGoqf9pD5uha5m"},
			},
		},
		{
			"It generates nothing when out of range",
			args{"904625697166532776746648320380374280100293470930272690489102837043110636999", 128},
			[]key{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotKeys := generateBitcoinKeys(tt.args.pageNumber, tt.args.keysPerpage); !reflect.DeepEqual(gotKeys, tt.wantKeys) {
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

func Test_findBtcWifPage(t *testing.T) {
	type args struct {
		wifString   string
		keysPerPage int
	}

	tests := []struct {
		name     string
		args     args
		wantPage string
	}{
		{
			"It can find the page that a random WIF is on",
			args{"5KQkycVaH2urSTz9CQ4fGdWz3a5n9TFKLDwxzREv8tBtcXYW9Ua", 128},
			"741968862012117112677494014490987968047399326671284349197372731288562495168",
		},
		{
			"It can find a WIF on the first page",
			args{"5HpHagT65TZzG1PH3CSu63k8DbpvD8s5ip4nEB3kEsreAnchuDf", 128},
			"1",
		},
		{
			"It can find a WIF on the last page",
			args{"5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqetqj84qw", 128},
			"904625697166532776746648320380374280100293470930272690489102837043110636675",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPage := findBtcWifPage(tt.args.wifString, tt.args.keysPerPage); !reflect.DeepEqual(gotPage, tt.wantPage) {
				t.Errorf("Expected: %v", tt.wantPage)
				t.Errorf("Actual: %v", gotPage)
			}
		})
	}
}
