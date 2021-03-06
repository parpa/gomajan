package player

import (
	"fmt"
	"testing"

	"github.com/kyokomi/gomajan/agari"
	"github.com/kyokomi/gomajan/foo"
	"github.com/kyokomi/gomajan/pai"
	"github.com/kyokomi/gomajan/tehai"
)

func TestCheckAgari(t *testing.T) {
	type TestCase struct {
		inAgari  pai.MJP
		inTiles  []tehai.Tehai
		inFoos   []foo.Foo
		outAgari []agari.Agari
	}

	testCases := []TestCase{
		// シャボ待ち
		TestCase{
			inAgari: pai.HAK,
			inTiles: tehai.NewTehai(map[pai.MJP]int{
				pai.P1:  1,
				pai.P2:  1,
				pai.P3:  1,
				pai.S7:  1,
				pai.S8:  1,
				pai.S9:  1,
				pai.HAK: 3,
				pai.S1:  2,
			}),
			inFoos: nil,
			outAgari: []agari.Agari{
				agari.Agari{
					Pai:     pai.HAK,
					Syanten: [2]pai.MJP{pai.HAK, pai.HAK},
					Type:    agari.Shabo,
				},
				agari.Agari{
					Pai:     pai.S1,
					Syanten: [2]pai.MJP{pai.S1, pai.S1},
					Type:    agari.Shabo,
				},
			},
		},
		// 両面待ち
		TestCase{
			inAgari: pai.P1,
			inTiles: tehai.NewTehai(map[pai.MJP]int{
				pai.P1:  1,
				pai.P2:  1,
				pai.P3:  1,
				pai.S7:  1,
				pai.S8:  1,
				pai.S9:  1,
				pai.HAK: 3,
				pai.S1:  2,
			}),
			inFoos: nil,
			outAgari: []agari.Agari{
				agari.Agari{
					Pai:     pai.P1,
					Syanten: [2]pai.MJP{pai.P2, pai.P3},
					Type:    agari.Ryanmen,
				},
			},
		},
		// 辺張待ち（12=>3）
		TestCase{
			inAgari: pai.P3,
			inTiles: tehai.NewTehai(map[pai.MJP]int{
				pai.P1:  1,
				pai.P2:  1,
				pai.P3:  1,
				pai.S7:  1,
				pai.S8:  1,
				pai.S9:  1,
				pai.HAK: 3,
				pai.S1:  2,
			}),
			inFoos: nil,
			outAgari: []agari.Agari{
				agari.Agari{
					Pai:     pai.P3,
					Syanten: [2]pai.MJP{pai.P1, pai.P2},
					Type:    agari.Penchan,
				},
			},
		},
		// 辺張待ち（89=>7）
		TestCase{
			inAgari: pai.S7,
			inTiles: tehai.NewTehai(map[pai.MJP]int{
				pai.P1:  1,
				pai.P2:  1,
				pai.P3:  1,
				pai.S7:  1,
				pai.S8:  1,
				pai.S9:  1,
				pai.HAK: 3,
				pai.S1:  2,
			}),
			inFoos: nil,
			outAgari: []agari.Agari{
				agari.Agari{
					Pai:     pai.S7,
					Syanten: [2]pai.MJP{pai.S8, pai.S9},
					Type:    agari.Penchan,
				},
			},
		},
		// 嵌張待ち（13=>2）
		TestCase{
			inAgari: pai.P2,
			inTiles: tehai.NewTehai(map[pai.MJP]int{
				pai.P1:  1,
				pai.P2:  1,
				pai.P3:  1,
				pai.S7:  1,
				pai.S8:  1,
				pai.S9:  1,
				pai.HAK: 3,
				pai.S1:  2,
			}),
			inFoos: nil,
			outAgari: []agari.Agari{
				agari.Agari{
					Pai:     pai.P2,
					Syanten: [2]pai.MJP{pai.P1, pai.P3},
					Type:    agari.Kanchan,
				},
			},
		},
	}

	for _, testCase := range testCases {
		m := newMentuCheck(testCase.inAgari, testCase.inTiles, testCase.inFoos)

		var as []agari.Agari
		as = m.CheckAgari()

		// 件数
		if len(as) != len(testCase.outAgari) {
			t.Errorf("out len error %d != %d \n", len(as), len(testCase.outAgari))
		}

		for idx, a := range as {
			fmt.Println(a.String())
			if a.String() != testCase.outAgari[idx].String() {
				t.Errorf("out agari case error %s != %s \n", a.String(), testCase.outAgari[idx].String())
			}
		}
	}
}
