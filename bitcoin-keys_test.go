package main

import (
	"math/big"
	"reflect"
	"testing"
)

func Test_generateBitcoinKeys(t *testing.T) {
	type args struct {
		firstSeed *big.Int
		amount    int
	}

	tests := []struct {
		name     string
		args     args
		wantKeys []key
	}{
		{
			"It can generate keys starting from the first seed",
			args{makeBigInt("1"), 10},
			[]key{
				{"5HpHagT65TZzG1PH3CSu63k8DbpvD8s5ip4nEB3kEsreAnchuDf", "1BgGZ9tcN4rm9KBzDn7KprQz87SZ26SAMH", "1EHNa6Q4Jz2uvNExL497mE43ikXhwF6kZm"},
				{"5HpHagT65TZzG1PH3CSu63k8DbpvD8s5ip4nEB3kEsreAvUcVfH", "1cMh228HTCiwS8ZsaakH8A8wze1JR5ZsP", "1LagHJk2FyCV2VzrNHVqg3gYG4TSYwDV4m"},
				{"5HpHagT65TZzG1PH3CSu63k8DbpvD8s5ip4nEB3kEsreB1FQ8BZ", "1CUNEBjYrCn2y1SdiUMohaKUi4wpP326Lb", "1NZUP3JAc9JkmbvmoTv7nVgZGtyJjirKV1"},
				{"5HpHagT65TZzG1PH3CSu63k8DbpvD8s5ip4nEB3kEsreB4AD8Yi", "1JtK9CQw1syfWj1WtFMWomrYdV3W2tWBF9", "1MnyqgrXCmcWJHBYEsAWf7oMyqJAS81eC"},
				{"5HpHagT65TZzG1PH3CSu63k8DbpvD8s5ip4nEB3kEsreBF8or94", "17Vu7st1U1KwymUKU4jJheHHGRVNqrcfLD", "1E1NUNmYw1G5c3FKNPd435QmDvuNG3auYk"},
				{"5HpHagT65TZzG1PH3CSu63k8DbpvD8s5ip4nEB3kEsreBKdE2NK", "1Cf2hs39Woi61YNkYGUAcohL2K2q4pawBq", "1UCZSVufT1PNimutbPdJUiEyCYSiZAD6n"},
				{"5HpHagT65TZzG1PH3CSu63k8DbpvD8s5ip4nEB3kEsreBR6zCMU", "19ZewH8Kk1PDbSNdJ97FP4EiCjTRaZMZQA", "1BYbgHpSKQCtMrQfwN6b6n5S718EJkEJ41"},
				{"5HpHagT65TZzG1PH3CSu63k8DbpvD8s5ip4nEB3kEsreBbMaQX1", "1EhqbyUMvvs7BfL8goY6qcPbD6YKfPqb7e", "1JMcEcKXQ7xA7JLAMPsBmHz68bzugYtdrv"},
				{"5HpHagT65TZzG1PH3CSu63k8DbpvD8s5ip4nEB3kEsreBd7uGcN", "1HSxWThjiwbC4dJbXHMpBfwRenB12UguG5", "1CijKR7rDvJJBJfSPyUYrWC8kAsQLy2B2e"},
				{"5HpHagT65TZzG1PH3CSu63k8DbpvD8s5ip4nEB3kEsreBoNWTw6", "13DaZ9nfmJLfzU6oBnD2sdCiDmf3M5fmLx", "1GDWJm5dPj6JTxF68WEVhicAS4gS3pvjo7"},
			},
		},
		{
			"It can generate keys until the last seed",
			args{makeBigInt("115792089237316195423570985008687907852837564279074904382605163141518161494327"), 10},
			[]key{
				{"5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqesmCC6YY", "1Aea8LKoEEWpPqTqaSwRYfksmUScVqV1F6", "18PUeum1Su423DmV2jEGdSd3ewiPfsZZ7z"},
				{"5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqestnHCQU", "1Nk4wGvaSinFVdrnMfEexLDnBZvWPY393C", "1GLiZZVt326aA8JHG2dEJHC591DXDQNKTs"},
				{"5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqet3sujS6", "1A81LWBrirUNAKpUVFS37xWT4GAMYU5qgD", "1MFyofP8SVtsEYDHQbZg7XJgfDeSP4ysPm"},
				{"5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqet8uM8zj", "1BJYFk5827oeYipArjTvLL7JdR4ivCGFYj", "1XunvtCGpmb7uw9qxWwaZFfHNFdUmuMVG"},
				{"5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqetEoeLmv", "1Lvxa3uJyPyRLbrNpGx761aSDWrJ77aTNm", "1J2zofmGpMUSaNGdTZEhMRYXdWsBQFMpS"},
				{"5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqetNQLySX", "1JSVicNeasrtuiDpb6r4J5fWxjfdU7ZyWT", "1LWBSfTeaLRNS1vyGSKy2BVW2nd6W9sk8Q"},
				{"5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqetVTGEAr", "1FjMR9gvnmZ3JYMxBbyc3aZK717b5txJoC", "1F3zbGb5JLBnmCAAYjCCv35zkggrXfi8LR"},
				{"5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqetbh69Dr", "1HjFHBmhUQkKntPPeWmiLiNGewRAMQWNYs", "15K4QVHD5T1KvW4it56qNuGJoTGMpUaFMj"},
				{"5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqetd9ZKJ4", "1NjSB7UL4MtdjmPbTUfaHne9R5C2YGxUSA", "1Knh2eFMtzMEtmvGHW14ELG8F9Ny6jV4s3"},
				{"5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqetqj84qw", "1GrLCmVQXoyJXaPJQdqssNqwxvha1eUo2E", "1JPbzbsAx1HyaDQoLMapWGoqf9pD5uha5m"},
			},
		},
		{
			"It stops when reaching the max seed",
			args{makeBigInt("115792089237316195423570985008687907852837564279074904382605163141518161494335"), 9999},
			[]key{
				{"5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqetd9ZKJ4", "1NjSB7UL4MtdjmPbTUfaHne9R5C2YGxUSA", "1Knh2eFMtzMEtmvGHW14ELG8F9Ny6jV4s3"},
				{"5Km2kuu7vtFDPpxywn4u3NLpbr5jKpTB3jsuDU2KYEqetqj84qw", "1GrLCmVQXoyJXaPJQdqssNqwxvha1eUo2E", "1JPbzbsAx1HyaDQoLMapWGoqf9pD5uha5m"},
			},
		},
		{
			"It generates nothing when out of range",
			args{makeBigInt("115792089237316195423570985008687907852837564279074904382605163141518161499999"), 9999},
			[]key{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotKeys := generateBitcoinKeys(tt.args.firstSeed, tt.args.amount); !reflect.DeepEqual(gotKeys, tt.wantKeys) {
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
