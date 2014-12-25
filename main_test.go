package main

import (
	"testing"

	"fmt"
	"github.com/kyokomi/gomajan/mjp"
	"github.com/kyokomi/gomajan/mjp/pai"
)

type TestCase struct {
	in     []mjp.Tehai
	inFoos []mjp.Foo
	out    []mjp.Yaku
}

func TestYakuCheck(t *testing.T) {
	testCases := []TestCase{
		TestCase{
			in: mjp.NewTehai(map[pai.MJP]int{
				pai.S1:  2,
				pai.S9:  1,
				pai.P1:  1,
				pai.P9:  1,
				pai.M1:  1,
				pai.M9:  1,
				pai.TON: 1,
				pai.NAN: 1,
				pai.SHA: 1,
				pai.PEI: 1,
				pai.HAK: 1,
				pai.HAT: 1,
				pai.CHN: 1,
			}),
			inFoos: nil,
			out:    []mjp.Yaku{mjp.M国士無双},
		},
		TestCase{
			in: mjp.NewTehai(map[pai.MJP]int{
				pai.S1:  2,
				pai.S2:  2,
				pai.M3:  2,
				pai.M4:  2,
				pai.P5:  2,
				pai.P6:  2,
				pai.TON: 2,
			}),
			inFoos: nil,
			out:    []mjp.Yaku{mjp.M七対子},
		},
		// ソーズ清一色
		TestCase{
			in: mjp.NewTehai(map[pai.MJP]int{
				pai.S1: 2,
				pai.S2: 2,
				pai.S3: 2,
				pai.S4: 1,
				pai.S5: 1,
				pai.S6: 1,
				pai.S7: 3,
				pai.S9: 2,
			}),
			inFoos: nil,
			out:    []mjp.Yaku{mjp.M清一色},
		},
		// マンズ清一色
		TestCase{
			in: mjp.NewTehai(map[pai.MJP]int{
				pai.M1: 2,
				pai.M2: 2,
				pai.M3: 2,
				pai.M4: 1,
				pai.M5: 1,
				pai.M6: 1,
				pai.M7: 3,
				pai.M9: 2,
			}),
			inFoos: nil,
			out:    []mjp.Yaku{mjp.M清一色},
		},
		// ピンズ清一色
		TestCase{
			in: mjp.NewTehai(map[pai.MJP]int{
				pai.P1: 2,
				pai.P2: 2,
				pai.P3: 2,
				pai.P4: 1,
				pai.P5: 1,
				pai.P6: 1,
				pai.P7: 3,
				pai.P9: 2,
			}),
			inFoos: nil,
			out:    []mjp.Yaku{mjp.M清一色},
		},
		// ソーズ混一色
		TestCase{
			in: mjp.NewTehai(map[pai.MJP]int{
				pai.S1:  2,
				pai.S2:  2,
				pai.S3:  2,
				pai.S4:  1,
				pai.S5:  1,
				pai.S6:  1,
				pai.S7:  3,
				pai.HAK: 2,
			}),
			inFoos: nil,
			out:    []mjp.Yaku{mjp.M混一色},
		},
		// マンズ混一色
		TestCase{
			in: mjp.NewTehai(map[pai.MJP]int{
				pai.M1:  2,
				pai.M2:  2,
				pai.M3:  2,
				pai.M4:  1,
				pai.M5:  1,
				pai.M6:  1,
				pai.M7:  3,
				pai.TON: 2,
			}),
			inFoos: nil,
			out:    []mjp.Yaku{mjp.M混一色},
		},
		// ピンズ混一色
		TestCase{
			in: mjp.NewTehai(map[pai.MJP]int{
				pai.P1:  2,
				pai.P2:  2,
				pai.P3:  2,
				pai.P4:  1,
				pai.P5:  1,
				pai.P6:  1,
				pai.P7:  3,
				pai.CHN: 2,
			}),
			inFoos: nil,
			out:    []mjp.Yaku{mjp.M混一色},
		},
		// 断么九
		TestCase{
			in: mjp.NewTehai(map[pai.MJP]int{
				pai.P2: 1,
				pai.P3: 1,
				pai.P4: 1,
				pai.S4: 1,
				pai.S5: 1,
				pai.S6: 1,
				pai.M7: 3,
				pai.S7: 3,
				pai.P8: 2,
			}),
			inFoos: nil,
			out:    []mjp.Yaku{mjp.M断么九},
		},
		// 緑一色
		TestCase{
			in: mjp.NewTehai(map[pai.MJP]int{
				pai.S2:  2,
				pai.S3:  2,
				pai.S4:  2,
				pai.S6:  3,
				pai.HAT: 3,
				pai.S8:  2,
			}),
			inFoos: nil,
			out:    []mjp.Yaku{mjp.M緑一色},
		},
		// 大三元
		TestCase{
			in: mjp.NewTehai(map[pai.MJP]int{
				pai.S2:  1,
				pai.S3:  1,
				pai.S4:  1,
				pai.P5:  2,
				pai.HAT: 3,
				pai.CHN: 3,
			}),
			inFoos: []mjp.Foo{mjp.NewFooPon(mjp.Toimen, pai.HAK)},
			out:    []mjp.Yaku{mjp.M大三元},
		},
		// 字一色
		TestCase{
			in: mjp.NewTehai(map[pai.MJP]int{
				pai.TON: 3,
				pai.NAN: 3,
				pai.HAT: 3,
				pai.CHN: 2,
			}),
			inFoos: []mjp.Foo{mjp.NewFooPon(mjp.Toimen, pai.HAK)},
			out:    []mjp.Yaku{mjp.M字一色},
		},
		// 大四喜
		TestCase{
			in: mjp.NewTehai(map[pai.MJP]int{
				pai.TON: 3,
				pai.NAN: 3,
				pai.PEI: 3,
				pai.P1:  2,
			}),
			inFoos: []mjp.Foo{mjp.NewFooPon(mjp.Toimen, pai.SHA)},
			out:    []mjp.Yaku{mjp.M大四喜},
		},
		// 小四喜
		TestCase{
			in: mjp.NewTehai(map[pai.MJP]int{
				pai.TON: 3,
				pai.NAN: 3,
				pai.PEI: 2,
				pai.P1:  1,
				pai.P2:  1,
				pai.P3:  1,
			}),
			inFoos: []mjp.Foo{mjp.NewFooPon(mjp.Toimen, pai.SHA)},
			out:    []mjp.Yaku{mjp.M小四喜},
		},
		// 四暗刻
		TestCase{
			in: mjp.NewTehai(map[pai.MJP]int{
				pai.TON: 3,
				pai.S1:  3,
				pai.S5:  3,
				pai.P9:  3,
				pai.CHN: 2,
			}),
			inFoos: nil,
			out:    []mjp.Yaku{mjp.M四暗刻},
		},
		// 三暗刻
		TestCase{
			in: mjp.NewTehai(map[pai.MJP]int{
				pai.TON: 3,
				pai.S1:  3,
				pai.S5:  3,
				pai.P7:  1,
				pai.P8:  1,
				pai.P9:  1,
				pai.CHN: 2,
			}),
			inFoos: nil,
			out:    []mjp.Yaku{mjp.M三暗刻},
		},
		// 清老頭
		TestCase{
			in: mjp.NewTehai(map[pai.MJP]int{
				pai.S1: 2,
				pai.P1: 3,
				pai.P9: 3,
				pai.S9: 3,
			}),
			inFoos: []mjp.Foo{mjp.NewFooPon(mjp.Toimen, pai.M1)},
			out:    []mjp.Yaku{mjp.M清老頭},
		},
		// 小三元
		TestCase{
			in: mjp.NewTehai(map[pai.MJP]int{
				pai.S2:  1,
				pai.S3:  1,
				pai.S4:  1,
				pai.P5:  1,
				pai.P6:  1,
				pai.P7:  1,
				pai.HAK: 3,
				pai.HAT: 3,
				pai.CHN: 2,
			}),
			inFoos: nil,
			out:    []mjp.Yaku{mjp.M小三元},
		},
		// 四槓子
		TestCase{
			in: mjp.NewTehai(map[pai.MJP]int{
				pai.CHN: 2,
			}),
			inFoos: []mjp.Foo{
				mjp.NewFooMinKan(mjp.Toimen, pai.M1),
				mjp.NewFooAnnKan(pai.M2),
				mjp.NewFooAnnKan(pai.M3),
				mjp.NewFooAnnKan(pai.M4),
			},
			out:    []mjp.Yaku{mjp.M四槓子},
		},
	}

	for _, testCase := range testCases {
		p := mjp.NewPlayer(testCase.in, testCase.inFoos)

		yakuCheck := p.NewYakuCheck()

		// 役の数
		if len(yakuCheck.Yakus()) != len(testCase.out) {
			t.Error(testCase.out, "!= ", yakuCheck.Yakus(), " error 手牌 => ", p)
			continue
		}

		// 役の内容（順番もチェックしてる）
		ok := true
		for idx, y := range yakuCheck.Yakus() {
			if !y.Equal(testCase.out[idx]) {
				t.Error(testCase.out, "!= ", yakuCheck.Yakus(), " error 手牌 => ", p)
				ok = false
				break
			}
		}
		if ok {
			fmt.Println(" => ", yakuCheck.String())
		}
	}
}
