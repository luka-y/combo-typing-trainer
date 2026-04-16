package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand/v2"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	screenWidth  = 1600
	screenHeight = 900
)

const FontSize = 16
const FontDPI = 72

var FontFace font.Face

const MaxLen = 10

type KeyCombo struct {
	Key         ebiten.Key
	NormalRune  rune
	ShiftedRune rune

	Shift   bool
	Control bool
	Alt     bool
	Meta    bool
}
type Game struct {
	InputStream     []KeyCombo
	GibberishStream []KeyCombo
}

func (g *Game) Update() error {
	justPressedKeys := inpututil.AppendJustPressedKeys([]ebiten.Key{})

	shiftPressed := ebiten.IsKeyPressed(ebiten.KeyShift)
	controlPressed := ebiten.IsKeyPressed(ebiten.KeyControl)
	altPressed := ebiten.IsKeyPressed(ebiten.KeyAlt)
	metaPressed := ebiten.IsKeyPressed(ebiten.KeyMeta)

	for _, k := range justPressedKeys {
		if isKeyModifierToSkip(k) {
			continue
		}
		normal, shifted := getRunesFromKey(k)
		g.InputStream = append(g.InputStream, KeyCombo{
			Key:         k,
			NormalRune:  normal,
			ShiftedRune: shifted,

			Shift:   shiftPressed,
			Control: controlPressed,
			Alt:     altPressed,
			Meta:    metaPressed,
		})
	}

	if len(g.InputStream) > MaxLen {
		g.InputStream = g.InputStream[len(g.InputStream)-MaxLen:]
	}

	return nil
}

func isKeyModifierToSkip(key ebiten.Key) bool {
	switch key {
	case ebiten.KeyShift, ebiten.KeyShiftLeft, ebiten.KeyShiftRight:
		return true
	case ebiten.KeyControl, ebiten.KeyControlLeft, ebiten.KeyControlRight:
		return true
	case ebiten.KeyAlt, ebiten.KeyAltLeft, ebiten.KeyAltRight:
		return true
	case ebiten.KeyMeta, ebiten.KeyMetaLeft, ebiten.KeyMetaRight:
		return true
	}
	return false
}

func getRunesFromKey(k ebiten.Key) (rune, rune) {
	switch ebiten.KeyName(ebiten.KeyS) {
	case "s":
		return USNormal[k], USShifted[k]
	case "і":
		return UANormal[k], UAShifted[k]
	case "ы":
		return RUNormal[k], RUShifted[k]
	}
	log.Print("getRunesFromKey: unknown layout")
	return 0, 0
}

func (g *Game) Draw(screen *ebiten.Image) {
	inputStreamString := ""
	for _, combo := range g.InputStream {
		inputStreamString += getStringFromCombo(combo) + ", "
	}
	text.Draw(screen, inputStreamString, FontFace, 0, FontSize, color.White)
	gibberishStreamString := ""
	for _, combo := range g.GibberishStream {
		gibberishStreamString += getStringFromCombo(combo) + ", "
	}
	text.Draw(screen, gibberishStreamString, FontFace, 0, FontSize*2, color.White)
}

func getStringFromCombo(combo KeyCombo) string {
	if combo.Shift && !combo.Control && !combo.Alt && !combo.Meta && combo.ShiftedRune != 0 {
		return string(combo.ShiftedRune)
	}

	res := ""
	if combo.Shift {
		res += "Shift + "
	}
	if combo.Control {
		res += "Control + "
	}
	if combo.Alt {
		res += "Alt + "
	}
	if combo.Meta {
		res += "Meta + "
	}
	if combo.NormalRune != 0 {
		res += string(combo.NormalRune)
	} else {
		res += combo.Key.String()
	}
	return res
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	var err error
	FontFace, err = getFaceFromPath("assets/JetBrainsMono-Regular.ttf", FontSize, FontDPI)
	if err != nil {
		log.Fatal(err)
	}

	gibberishStream, err := generateGibberishStream(100)
	if err != nil {
		log.Fatal(err)
	}
	game := &Game{GibberishStream: gibberishStream}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("combo-typing-trainer")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func generateGibberishStream(size int) ([]KeyCombo, error) {
	baseCumulativeDistribution, err := convertWeightsToCumulativeDistribution([]int{
		LowercaseLettersWeight,
		UppercaseLettersWeight,
		SymbolsWeight,
		NumbersWeight,
		BaseLayerNonPrintableWeight,
		LowerLayerNonPrintableWeight,
		CustomWeight,
	})
	if err != nil {
		return nil, fmt.Errorf("error making cumulative distribution for base rune/key: %w", err)
	}
	modifierCumulativeDistribution, err := convertWeightsToCumulativeDistribution([]int{
		NoModifiersWeight,
		ControlWeight,
		AltWeight,
		ShiftWeight,
		MetaWeight,
		ControlAltWeight,
		ControlShiftWeight,
		ControlMetaWeight,
		AltShiftWeight,
		AltMetaWeight,
		ShiftMetaWeight,
		ControlAltShiftWeight,
		ControlAltMetaWeight,
		ControlShiftMetaWeight,
		AltShiftMetaWeight,
		ControlAltShiftMetaWeight,
	})
	if err != nil {
		return nil, fmt.Errorf("error making cumulative distribution for modifiers: %w", err)
	}

	res := make([]KeyCombo, size)
	for i := 0; i < size; i++ {
		res[i], err = getRandomCombo(baseCumulativeDistribution, modifierCumulativeDistribution)
		if err != nil {
			return nil, fmt.Errorf("error making random combo: %w", err)
		}
	}
	return res, nil
}

func getRandomCombo(baseCumulativeDistribution, modifierCumulativeDistribution []float64) (KeyCombo, error) {
	baseIndex := EmptyBaseWeightIndex
	randFloat := rand.Float64()
	for i, num := range baseCumulativeDistribution {
		if num >= randFloat {
			baseIndex = BaseWeightIndex(i)
			break
		}
	}
	if baseIndex == EmptyBaseWeightIndex {
		return KeyCombo{}, fmt.Errorf("impossible error baseIndex")
	}
	modIndex := EmptyModWeightIndex
	randFloat = rand.Float64()
	for i, num := range modifierCumulativeDistribution {
		if num >= randFloat {
			modIndex = ModWeightIndex(i)
			break
		}
	}
	if modIndex == EmptyModWeightIndex {
		return KeyCombo{}, fmt.Errorf("impossible error modIndex")
	}

	//TODO finish

	return KeyCombo{}, nil
}

func convertWeightsToCumulativeDistribution(inputSlice []int) ([]float64, error) {
	if len(inputSlice) == 0 {
		return nil, fmt.Errorf("empty input slice")
	}
	weightsSum := 0
	for _, weight := range inputSlice {
		if weight < 0 {
			return nil, fmt.Errorf("input slice included a negative weight")
		}
		weightsSum += weight
	}
	if weightsSum == 0 {
		return nil, fmt.Errorf("sum of weight equals zero")
	}
	probabilities := make([]float64, len(inputSlice))
	for i := 0; i < len(inputSlice); i++ {
		probabilities[i] = float64(inputSlice[i]) / float64(weightsSum)
		if i > 0 {
			probabilities[i] += probabilities[i-1]
		}
	}
	return probabilities, nil
}

func getFaceFromPath(path string, size, dpi float64) (font.Face, error) {
	fontBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	f, err := opentype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}
	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    size,
		DPI:     dpi,
		Hinting: font.HintingNone,
	})
	if err != nil {
		return nil, err
	}
	return face, nil
}
