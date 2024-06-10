package shared

type SidChip uint8
type VoiceIndex uint8
type AttackRate uint8
type ReleaseRate uint8
type DecayRate uint8

const (
	ChipLeft SidChip = iota
	ChipRight
)

const (
	Voice1 VoiceIndex = iota
	Voice2
	Voice3
	VoiceEx // used only for filter setting
)

const (
	Attack2MS AttackRate = iota
	Attack8MS
	Attack16MS
	Attack24MS
	Attack38MS
	Attack56MS
	Attack68MS
	Attack80MS
	Attack100MS
	Attack250MS
	Attack500MS
	Attack800MS
	Attack1000MS
	Attack3000MS
	Attack5000MS
	Attack6000MS
)

var (
	AttackText = []string{
		"2ms", "8ms", "16ms", "24ms",
		"38ms", "56ms", "68ms", "80ms",
		".1s", ".25s", ".5s", ".8s",
		"1s", "3s", "5s", "6s",
	}
)

const (
	Release6MS ReleaseRate = iota
	Release24MS
	Release48MS
	Release72MS
	Release114MS
	Release168MS
	Release204MS
	Release240MS
	Release300MS
	Release750MS
	Release1500MS
	Release2400MS
	Release3000MS
	Release9000MS
	Release15000MS
	Release24000MS
)

var (
	ReleaseText = []string{
		"6ms", "24ms", "48ms", "72ms",
		".114s", ".168s", ".204s", ".24s",
		".3s", ".75s", "1.5s", "2.4s",
		"3s", "9s", "15s", "24s",
	}
)

const (
	Decay6MS DecayRate = iota
	Decay24MS
	Decay48MS
	Decay72MS
	Decay114MS
	Decay168MS
	Decay204MS
	Decay240MS
	Decay300MS
	Decay750MS
	Decay1500MS
	Decay2400MS
	Decay3000MS
	Decay9000MS
	Decay15000MS
	Decay24000MS
)

var (
	DecayText = []string{
		"6ms", "24ms", "48ms", "72ms",
		".114s", ".168s", ".204s", ".24s",
		".3s", ".75s", "1.5s", "2.4s",
		"3s", "9s", "15s", "24s",
	}
)

func BToI(trueOrFalse bool) uint8 {
	if trueOrFalse {
		return 1
	}
	return 0
}

func IToB(zeroOrOne uint8) bool {
	if zeroOrOne == 1 {
		return true
	}
	return false
}
