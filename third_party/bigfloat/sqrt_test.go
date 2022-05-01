package bigfloat_test

// Copied from https://github.com/ALTree/bigfloat since it has no release tags.

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"testing"

	"bitbucket.org/decimalteam/go-smart-node/third_party/bigfloat"
)

const maxPrec uint = 1100

// We can't guarantee that the result will have *prec* precision
// if we call Sqrt with an argument with *prec* precision, because
// the Newton's iteration will actually converge to a number that
// is not the square root of x with *prec+64* precision, but to a
// number that is the square root of x with *prec* precision.
// If we want Sqrt(x) with *prec* precision and correct rounding,
// we need to call Sqrt with an argument having precision greater
// than *prec*.
//
// This will happen for every test number which is not directly
// representable as a binary floating point value, so avoid those.

func TestSqrt(t *testing.T) {
	for _, test := range []struct {
		z    string
		want string
	}{
		// 350 decimal digits are enough to give us up to 1000 binary digits
		{"0.5", "0.70710678118654752440084436210484903928483593768847403658833986899536623923105351942519376716382078636750692311545614851246241802792536860632206074854996791570661133296375279637789997525057639103028573505477998580298513726729843100736425870932044459930477616461524215435716072541988130181399762570399484362669827316590441482031030762917619752737287514"},
		{"2.0", "1.4142135623730950488016887242096980785696718753769480731766797379907324784621070388503875343276415727350138462309122970249248360558507372126441214970999358314132226659275055927557999505011527820605714701095599716059702745345968620147285174186408891986095523292304843087143214508397626036279952514079896872533965463318088296406206152583523950547457503"},
		{"3.0", "1.7320508075688772935274463415058723669428052538103806280558069794519330169088000370811461867572485756756261414154067030299699450949989524788116555120943736485280932319023055820679748201010846749232650153123432669033228866506722546689218379712270471316603678615880190499865373798593894676503475065760507566183481296061009476021871903250831458295239598"},
		{"4.0", "2.0"},
		{"5", "2.2360679774997896964091736687312762354406183596115257242708972454105209256378048994144144083787822749695081761507737835042532677244470738635863601215334527088667781731918791658112766453226398565805357613504175337850034233924140644420864325390972525926272288762995174024406816117759089094984923713907297288984820886415426898940991316935770197486788844"},
		{"6", "2.4494897427831780981972840747058913919659474806566701284326925672509603774573150265398594331046402348185946012266141891248588654598377573416257839512372785528289127475276765712476301052709117702234813106789866908536324433525456040338088089393745855678465747243613041442702702161742018383000815898078380130897007286939936308371580944008004437386875492"},
		{"7", "2.6457513110645905905016157536392604257102591830824501803683344592010688232302836277603928864745436106150645783384974630957435298886272147844273905558801077227171507297283238922996895948650872607009780542037238280237159411003419391160015785255963059457410351523968027164073737990740415815199044034743194536713997305970050513996922375456160971190273782"},

		{"1p512", "1p256"},
		{"1p1024", "1p512"},
		{"1p2048", "1p1024"},
		{"1p4096", "1p2048"},
		{"2p512", "1.4142135623730950488016887242096980785696718753769480731766797379907324784621070388503875343276415727350138462309122970249248360558507372126441214970999358314132226659275055927557999505011527820605714701095599716059702745345968620147285174186408891986095523292304843087143214508397626036279952514079896872533965463318088296406206152583523950547457503p256"},
		{"2p1024", "1.4142135623730950488016887242096980785696718753769480731766797379907324784621070388503875343276415727350138462309122970249248360558507372126441214970999358314132226659275055927557999505011527820605714701095599716059702745345968620147285174186408891986095523292304843087143214508397626036279952514079896872533965463318088296406206152583523950547457503p512"},
		{"2p2048", "1.4142135623730950488016887242096980785696718753769480731766797379907324784621070388503875343276415727350138462309122970249248360558507372126441214970999358314132226659275055927557999505011527820605714701095599716059702745345968620147285174186408891986095523292304843087143214508397626036279952514079896872533965463318088296406206152583523950547457503p1024"},

		{"1p-512", "1p-256"},
		{"1p-1024", "1p-512"},
		{"1p-2048", "1p-1024"},
		{"1p-4096", "1p-2048"},
		{"2p-512", "1.4142135623730950488016887242096980785696718753769480731766797379907324784621070388503875343276415727350138462309122970249248360558507372126441214970999358314132226659275055927557999505011527820605714701095599716059702745345968620147285174186408891986095523292304843087143214508397626036279952514079896872533965463318088296406206152583523950547457503p-256"},
		{"2p-1024", "1.4142135623730950488016887242096980785696718753769480731766797379907324784621070388503875343276415727350138462309122970249248360558507372126441214970999358314132226659275055927557999505011527820605714701095599716059702745345968620147285174186408891986095523292304843087143214508397626036279952514079896872533965463318088296406206152583523950547457503p-512"},
		{"2p-2048", "1.4142135623730950488016887242096980785696718753769480731766797379907324784621070388503875343276415727350138462309122970249248360558507372126441214970999358314132226659275055927557999505011527820605714701095599716059702745345968620147285174186408891986095523292304843087143214508397626036279952514079896872533965463318088296406206152583523950547457503p-1024"},
	} {
		for _, prec := range []uint{24, 53, 64, 100, 200, 300, 400, 500, 600, 700, 800, 900, 1000} {
			want := new(big.Float).SetPrec(prec)
			want.Parse(test.want, 10)

			z := new(big.Float).SetPrec(prec)
			z.Parse(test.z, 10)

			x := bigfloat.Sqrt(z)

			if x.Cmp(want) != 0 {
				t.Errorf("prec = %d, Sqrt(%v) =\ngot  %g;\nwant %g", prec, test.z, x, want)
			}
		}
	}
}

func testSqrtFloat64(scale float64, nTests int, t *testing.T) {
	for i := 0; i < nTests; i++ {
		r := rand.Float64() * scale

		z := big.NewFloat(r)
		x64, acc := bigfloat.Sqrt(z).Float64()

		want := math.Sqrt(r)

		if x64 != want || acc != big.Exact {
			t.Errorf("Sqrt(%g) =\n got %g (%s);\nwant %g (Exact)", z, x64, acc, want)
		}
	}
}

func TestSqrtFloat64Small(t *testing.T) {
	testSqrtFloat64(1e-100, 1e5, t)
	testSqrtFloat64(1e-10, 1e5, t)
}

func TestSqrtFloa64Medium(t *testing.T) {
	testSqrtFloat64(1, 1e5, t)
	testSqrtFloat64(100, 1e5, t)
}

func TestSqrtFloat64Big(t *testing.T) {
	testSqrtFloat64(1e10, 1e5, t)
	testSqrtFloat64(1e100, 1e5, t)
}

func TestSqrtSpecialValues(t *testing.T) {
	for _, f := range []float64{
		+0.0,
		-0.0,
		math.Inf(+1),
	} {
		z := big.NewFloat(f)
		x64, acc := bigfloat.Sqrt(z).Float64()
		want := math.Sqrt(f)
		if x64 != want || acc != big.Exact {
			t.Errorf("Sqrt(%g) =\n got %g (%s);\nwant %g (Exact)", z, x64, acc, want)
		}
	}
}

// ---------- Benchmarks ----------

func BenchmarkSqrt(b *testing.B) {
	for _, prec := range []uint{1e2, 1e3, 1e4, 1e5} {
		z := big.NewFloat(2).SetPrec(prec)
		b.Run(fmt.Sprintf("%v", prec), func(b *testing.B) {
			b.ReportAllocs()
			for n := 0; n < b.N; n++ {
				bigfloat.Sqrt(z)
			}
		})
	}
}
