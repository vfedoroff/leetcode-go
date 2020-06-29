package solution

import (
	"reflect"
	"sort"
	"testing"
)

type Node struct {
	Сhildren  [27]*Node
	Sentences []string
}

func index(c byte) int {
	if c == ' ' {
		return 26
	}
	return int(c - 'a')
}

func (node *Node) Add(sentence string) {
	root := node
	for i := 0; i < len(sentence); i++ {
		letter := sentence[i]
		i := index(letter)
		child := root.Сhildren[i]
		if root.Сhildren[i] == nil {
			child = &Node{}
			root.Сhildren[i] = child
		}
		root = child
		for _, s := range root.Sentences {
			if s == sentence {
				return
			}
		}
		root.Sentences = append(root.Sentences, sentence)
	}
}

type AutocompleteSystem struct {
	Trie        Node
	Times       map[string]int
	SessionTrie *Node
	History     []byte
}

func Constructor(sentences []string, times []int) AutocompleteSystem {
	l := len(sentences)
	a := AutocompleteSystem{}
	a.Trie = Node{}
	a.Times = make(map[string]int, l)
	for i := 0; i < l; i++ {
		a.Trie.Add(sentences[i])
		a.Times[sentences[i]] = times[i]
	}
	a.SessionTrie = &a.Trie
	return a
}

func (a *AutocompleteSystem) Input(c byte) []string {
	ret := make([]string, 0)
	if c != byte('#') {
		a.History = append(a.History, c)
		var node *Node
		if a.SessionTrie != nil {
			if node = a.SessionTrie.Сhildren[index(c)]; node != nil {
				ret = node.Sentences
				sort.Slice(ret, func(i, j int) bool {
					x := a.Times[ret[i]]
					y := a.Times[ret[j]]
					if x == y {
						return ret[i] < ret[j]
					}
					return x > y
				})
			}
		}
		a.SessionTrie = node
	} else {
		history := string(a.History)
		a.Trie.Add(history)
		a.Times[history]++
		a.History = []byte{}
		a.SessionTrie = &a.Trie
		ret = []string{}
	}
	if len(ret) >= 3 {
		return ret[:3]
	}
	return ret
}

func TestAutoComplete(t *testing.T) {
	autoCompleteSystem := Constructor([]string{"i love you", "island", "iroman", "i love leetcode"}, []int{5, 3, 2, 2})
	type test struct {
		c    byte
		want []string
	}
	tests := []test{
		{
			c:    byte('i'),
			want: []string{"i love you", "island", "i love leetcode"},
		},
		{
			c:    byte(' '),
			want: []string{"i love you", "i love leetcode"},
		},
		{
			c:    byte('a'),
			want: []string{},
		},
		{
			c:    byte('#'),
			want: []string{},
		},
	}
	for _, tc := range tests {
		got := autoCompleteSystem.Input(tc.c)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("AutocompleteSystem.Input() = %#v, want %#v", got, tc.want)
		}
	}
}

func TestAutoComplete2(t *testing.T) {
	autoCompleteSystem := Constructor([]string{"uqpewwnxyqxxlhiptuzevjxbwedbaozz", "ewftoujyxdgjtazppyztom", "pvyqceqrdrxottnukgbdfcr", "qtdkgdbcyozhllfycfjhdsdnuhycqcofaojknuqqnozltrjcabyxrdqwrxvqrztkcxpenbbtnnnkfhmebj", "jwfbusbwahyugiaiazysqbxkwgcawpniptbtmhqyrlxdwxxwhtumglihrgizrczv", "cfptjitfzdcrhw", "aitqgitjgrcbacgnaasvbouqsqcwbyskkpsnigtfeecmlkcjbgduban", "utsqkmiqqgglufourfdpgdmrkbippffacwvtkpflzrvdlkdxykfpkoqcb", "ethtbdopotpamvrwuomlpahtveyw", "jiaqkaxovsqtkpdjfbkajpvpyetuoqwnrnpjdhoojbsdvneecsdvgqpyurmsvcy", "j", "btbnuplyeuccjbernsfbnveillrwdbqledwvpmvdbcugkurrkabtpykhlcogeszclyfuquafouv", "hndjzblegevtfkgbjttektox", "gtvnlninpvenapyfgmsjdisfnmiktitrutctawosjflvzfkbegnprixzqwzcyhoovsivuwmofsveqkyosowuyamuvy", "sawrirvrfrbfagreahrioaombukmdwztbpggnxd", "mgdcwptvbvhzyvvumvbjjn", "otjvvkegwleyyxtghwgfmlsqlhrlibdvqfinyyebotjpwoaejhtornfgikmifdmwswbqgwhcbzuhrpajxuqicegcptszct", "zlondsttyvnnnnxjtoqnlktitwzurissczzbyfsbgpoawodwjpsmavaugnhqtsbeixwl", "yehvdehbtmwqkmcjmvpivfzqvevkotwzvjoyfvp", "bjximtpayjdcxbrnksbtfnpynzaygygdflowewprqngdadzdhxcpgapjejojrkzrutgcsfpfvpluagniqimfqddldxqiw", "bysyrxfykivyauysytgxfhqcrxliulahuizjvozpywrokxujhzpauxwufcxiitukljiiclatfrspqcljjoxpxziumstnhqr", "uxtvutlgqapyfltiulwrplesmtowzoyhhjhzihatpuvmutxqgxfawpwypedbz", "jzgsdjdawrqfladolduldhpdpagmvllvzamypuqlrpbmhxxadqaqrqavtxeghcyysjynovkiyjtvdluttodtmtocajgttmv", "mbijfkmepalhdiubposdksdmmttxblkodcdrxbnxaqebnwliatnxpwaohbwkidia", "ljggggbyxwrwanhjonoramexdmgjigrtpz", "cqfvkutpipxjepfgsufonvjtotwfxyn", "kvseesjazssavispavchdpzvdhibptowhyrrshyntpwkez", "nveuzbaosuayteiozmnelxlwkrrrjlwvhejxhupvchfwmvnqukphgoacnazuoimcliubvhv", "uwrpwhfdrxfnarxqpkhrylkwiuhzubjfk", "bniyggdcloefwy", "ihranmhbsahqjxesbtmdkjfsupzdzjvdfovgbtwhqfjdddwhdvrnlyscvqlnqpzegnvvzyymrajvso", "lscreasfuxpdxsiinymuzybjexkpfjiplevqcjxlm", "uwgnfozopsygnmptdtmmuumahoungpkodwxrcvfymqpeymaqruayvqqgoddgbnhemtsjifhxwiehncswxzrghf", "nyfbxgcpfrzyqwfjzgmhuohjhrjizsyjqgmertmooeiaadcmiuyyylpcosnweoyydeauazhzbeaqn", "tpylrxbwnkrfxckfdlvrbytaezuzmyscpvruthuvbxjenkeolvqsrjqzizyclzmqtjvnamdansmzyspcfghfprorqprua", "nhldlmxpuckxeekipkrzugatjiivtazjbjyxokksyueyjbgmrovbckbxqcqefaiavzsarbbypgmpxe", "sylraqsd", "xr", "xkzpxkhrucyatpatkigvntohihibyisyqtkjdhatdvyvxbjttz", "nvnz", "blzddwxphkgqfsfzfclwytstpvpzgcdeggdwzukzirscfzcteeuqbmmrfxcnokbbyxkqrxtjfarcefiynwfmy", "inuxmuhtdwpuvyludwtokhtalxbuccepsayrjycbcwbtnfholjvkmypodv", "awwillrm", "xznodxngrstjrwqzmlmigpw", "khlxjdtictufdfbkgfusdtaaeuspbbfmtjodflgqofzlqnulkdztflm", "nlngmckslyqzjiyiexbropbxnynjcstziluewypboqhqndqsxhtnosrgrameajovsclrgwqjgnztvxrkhwnxkfrf", "yroadxhxyacaexrwju", "ujxlbpcbxdqrvubifnpzjmmkolyljzjhdegaiowaal", "tnfnjgtxbckbpyplucprxcqzhrfjimylmlhdglntfydepltjvklyxesndzuubienhvuaqc", "ouedhtkpkg", "ygchsrrubucqffewifsxaefwocfaiiupqbomktvrcddggqfgnaycstpccwtbheyaqwhosxajqeqqxzyjsfng", "jqqgpjvfkgjh", "csowoazaiyejgyixszqmtctpzlkccccqregyhtvxccvrpkupwcyhqatxscevzdfbdqnuyadiyfnhysddfyxpgqtjiogmxsmzbbkr", "dlzxdpchkdaztkqtrjmuujgoiae", "plcjkwukkyqluxjbhxsyeaqvviinfuujsafwsquidvmutsrukxwrv", "yopqbtpoqhpcktjangauzcvvpephhprpaaclbbkgchlqkrwdsaupeizlwxzcpkchoagmrrkwdkthosmrjefgbumnrjsb"}, []int{12, 9, 4, 4, 1, 5, 3, 4, 7, 9, 2, 4, 2, 3, 11, 13, 1, 3, 4, 10, 7, 1, 9, 5, 10, 14, 5, 3, 2, 11, 5, 14, 4, 13, 11, 5, 15, 8, 1, 12, 2, 11, 4, 2, 11, 14, 9, 12, 1, 7, 13, 11, 7, 2, 6, 10})
	type test struct {
		c    byte
		want []string
	}
	tests := []test{
		{
			c:    byte('w'),
			want: []string{},
		},
		{
			c:    byte('o'),
			want: []string{},
		},
	}
	for _, tc := range tests {
		got := autoCompleteSystem.Input(tc.c)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("AutocompleteSystem.Input() = %#v, want %#v", got, tc.want)
		}
	}
}

func BenchmarkAutoComplete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		autoCompleteSystem := Constructor([]string{"i love you", "island", "iroman", "i love leetcode"}, []int{5, 3, 2, 2})
		type test struct {
			c byte
		}
		tests := []test{
			{
				c: byte('i'),
			},
			{
				c: byte(' '),
			},
			{
				c: byte('a'),
			},
			{
				c: byte('#'),
			},
		}
		for _, tc := range tests {
			autoCompleteSystem.Input(tc.c)
		}
	}

}
