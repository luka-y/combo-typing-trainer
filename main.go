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

const ScreenWidth = 1920
const ScreenHeight = 50

const GibberishOverInputLen = 100

var BaseCumulativeDistribution []float64
var ModifierCumulativeDistribution []float64

const FontSize = 16
const FontDPI = 72

var FontFace font.Face
var FontDrawer *font.Drawer

var BackgroundColor = color.RGBA{24, 22, 22, 255}
var UpcomingComboColor = color.RGBA{110, 110, 105, 255}
var CurrentComboColor = color.RGBA{196, 178, 138, 255}
var CorrectPastComboColor = color.RGBA{138, 154, 123, 255}
var IncorrectPastComboColor = color.RGBA{196, 116, 110, 255}

var GibberishYPos int
var InputYPos int

var layout Layout //Set on the start of the program, to reset restart the program.
type Layout int

const (
	USLayout Layout = iota
	UALayout
	RULayout
)

type KeyCombo struct {
	Key         ebiten.Key
	NormalRune  rune
	ShiftedRune rune

	Shift   bool
	Control bool
	Alt     bool
	Meta    bool

	StrToDraw string
}

func (kc *KeyCombo) setStringToDraw() {
	if kc.Shift && !kc.Control && !kc.Alt && !kc.Meta && kc.ShiftedRune != 0 {
		kc.StrToDraw = string(kc.ShiftedRune)
		return
	}

	res := ""
	if kc.Control {
		res += "Control+"
	}
	if kc.Alt {
		res += "Alt+"
	}
	if kc.Shift {
		res += "Shift+"
	}
	if kc.Meta {
		res += "Meta+"
	}
	if kc.NormalRune != 0 {
		res += string(kc.NormalRune)
	} else {
		res += kc.Key.String()
	}
	kc.StrToDraw = res
}

type Game struct {
	InputStream     []KeyCombo
	GibberishStream []KeyCombo

	TickCounter int
}

func (g *Game) Update() error {
	if g.TickCounter == 0 {
		g.TickCounter++
		return g.FirstUpdateCall()
	}
	g.TickCounter++

	justPressedKeys := inpututil.AppendJustPressedKeys([]ebiten.Key{})

	shiftPressed := ebiten.IsKeyPressed(ebiten.KeyShift)
	controlPressed := ebiten.IsKeyPressed(ebiten.KeyControl)
	altPressed := ebiten.IsKeyPressed(ebiten.KeyAlt)
	metaPressed := ebiten.IsKeyPressed(ebiten.KeyMeta)

	var pressedIds []int
	for _, k := range justPressedKeys {
		if isKeyModifierToSkip(k) {
			continue
		}
		if k == ebiten.KeyBackspace {
			continue
		}
		normal, shifted := getRunesFromKey(k)
		inputCombo := KeyCombo{
			Key:         k,
			NormalRune:  normal,
			ShiftedRune: shifted,

			Shift:   shiftPressed,
			Control: controlPressed,
			Alt:     altPressed,
			Meta:    metaPressed,
		}
		inputCombo.setStringToDraw()
		g.InputStream = append(g.InputStream, inputCombo)

		pressedIds = append(pressedIds, len(g.InputStream)-1)
	}

	lenDiff := len(g.GibberishStream) - len(g.InputStream)
	if lenDiff < GibberishOverInputLen {
		for i := 0; i < GibberishOverInputLen-lenDiff; i++ {
			randomCombo, err := getRandomCombo(BaseCumulativeDistribution, ModifierCumulativeDistribution)
			if err != nil {
				return fmt.Errorf("err with getRandomCombo in Update: %w", err)
			}
			g.GibberishStream = append(g.GibberishStream, randomCombo)
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) && len(g.InputStream) > 0 {
		g.InputStream = g.InputStream[0 : len(g.InputStream)-1]
	}

	return nil
}

func measureStringWidth(s string) int {
	return int(FontDrawer.MeasureString(s) >> 6)
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
	switch layout {
	case USLayout:
		return USNormal[k], USShifted[k]
	case UALayout:
		return UANormal[k], UAShifted[k]
	case RULayout:
		return RUNormal[k], RUShifted[k]
	}
	log.Print("getRunesFromKey: unknown layout")
	return 0, 0
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(BackgroundColor)

	spaceBetweenCombos := "  "

	currentScreenXPosition := ScreenWidth / 2
	for i := 0; i < len(g.InputStream); i++ {
		inputWidth := measureStringWidth(g.InputStream[i].StrToDraw + spaceBetweenCombos)
		gibberishWidth := measureStringWidth(g.GibberishStream[i].StrToDraw + spaceBetweenCombos)
		if inputWidth < gibberishWidth {
			currentScreenXPosition -= gibberishWidth
		} else {
			currentScreenXPosition -= inputWidth
		}
	}

	for i := 0; i < len(g.GibberishStream); i++ {
		inputColor, gibberishColor := UpcomingComboColor, UpcomingComboColor
		if i == len(g.InputStream) {
			gibberishColor = CurrentComboColor
		}
		if i < len(g.InputStream) {
			if g.GibberishStream[i] == g.InputStream[i] {
				inputColor, gibberishColor = CorrectPastComboColor, CorrectPastComboColor
			} else {
				inputColor, gibberishColor = IncorrectPastComboColor, IncorrectPastComboColor
			}
		}

		if i < len(g.InputStream) {
			text.Draw(screen, g.InputStream[i].StrToDraw+spaceBetweenCombos, FontFace, currentScreenXPosition, InputYPos, inputColor)
		}
		text.Draw(screen, g.GibberishStream[i].StrToDraw+spaceBetweenCombos, FontFace, currentScreenXPosition, GibberishYPos, gibberishColor)

		inputWidth := -1
		if i < len(g.InputStream) {
			inputWidth = measureStringWidth(g.InputStream[i].StrToDraw + spaceBetweenCombos)
		}
		gibberishWidth := measureStringWidth(g.GibberishStream[i].StrToDraw + spaceBetweenCombos)
		if inputWidth < gibberishWidth {
			currentScreenXPosition += gibberishWidth
		} else {
			currentScreenXPosition += inputWidth
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *Game) FirstUpdateCall() error {
	switch ebiten.KeyName(ebiten.KeyS) {
	case "s":
		layout = USLayout
	case "і":
		layout = UALayout
	case "ы":
		layout = RULayout
	default:
		return fmt.Errorf("unknown layout")
	}

	err := setCumulativeDistributions()
	if err != nil {
		return err
	}
	gibberishStream := make([]KeyCombo, GibberishOverInputLen)
	for i := 0; i < GibberishOverInputLen; i++ {
		gibberishStream[i], err = getRandomCombo(BaseCumulativeDistribution, ModifierCumulativeDistribution)
		if err != nil {
			return fmt.Errorf("error making random combo: %w", err)
		}
	}
	g.GibberishStream = gibberishStream
	return nil
}

func main() {
	var err error
	FontFace, err = getFaceFromPath("assets/JetBrainsMono-Regular.ttf", FontSize, FontDPI)
	if err != nil {
		log.Fatal(err)
	}
	FontDrawer = &font.Drawer{Face: FontFace}

	InputYPos = FontSize + 2
	GibberishYPos = ScreenHeight - int(FontFace.Metrics().Descent.Ceil()) - 2

	if CustomKeysWeight > 0 && len(CustomKeys) == 0 {
		log.Fatal("config error: non-zero CustomKeysWeight while CustomKeys's len is 0")
	}
	if CustomCharsWeight > 0 && len(CustomChars) == 0 {
		log.Fatal("config error: non-zero CustomCharsWeight while CustomChars's len is 0")
	}

	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("combo-typing-trainer")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

func setCumulativeDistributions() error {
	var err error
	BaseCumulativeDistribution, err = convertWeightsToCumulativeDistribution([]int{
		LowercaseLettersWeight,
		UppercaseLettersWeight,
		SymbolsWeight,
		NumbersWeight,
		BaseLayerNonPrintableWeight,
		LowerLayerNonPrintableWeight,
		CustomCharsWeight,
		CustomKeysWeight,
	})
	if err != nil {
		return fmt.Errorf("error making cumulative distribution for base rune/key: %w", err)
	}
	ModifierCumulativeDistribution, err = convertWeightsToCumulativeDistribution([]int{
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
		return fmt.Errorf("error making cumulative distribution for modifiers: %w", err)
	}

	return nil
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

	var res = KeyCombo{}
	switch baseIndex {
	case LowercaseLettersWeightIndex:
		switch layout {
		case USLayout:
			_ = setKeyComboBasedOnRuneAndLayout(&res, USLower[rand.IntN(len(USLower))])
		case UALayout:
			_ = setKeyComboBasedOnRuneAndLayout(&res, UALower[rand.IntN(len(UALower))])
		case RULayout:
			_ = setKeyComboBasedOnRuneAndLayout(&res, RULower[rand.IntN(len(RULower))])
		}
	case UppercaseLettersWeightIndex:
		switch layout {
		case USLayout:
			_ = setKeyComboBasedOnRuneAndLayout(&res, USUpper[rand.IntN(len(USUpper))])
		case UALayout:
			_ = setKeyComboBasedOnRuneAndLayout(&res, UAUpper[rand.IntN(len(UAUpper))])
		case RULayout:
			_ = setKeyComboBasedOnRuneAndLayout(&res, RUUpper[rand.IntN(len(RUUpper))])
		}
	case SymbolsWeightIndex:
		switch layout {
		case USLayout:
			_ = setKeyComboBasedOnRuneAndLayout(&res, USSymbols[rand.IntN(len(USSymbols))])
		case UALayout:
			_ = setKeyComboBasedOnRuneAndLayout(&res, UASymbols[rand.IntN(len(UASymbols))])
		case RULayout:
			_ = setKeyComboBasedOnRuneAndLayout(&res, RUSymbols[rand.IntN(len(RUSymbols))])
		}
	case NumbersWeightIndex:
		_ = setKeyComboBasedOnRuneAndLayout(&res, Numbers[rand.IntN(len(Numbers))])
	case BaseLayerNonPrintableWeightIndex:
		res.Key = BaseLayerNonPrintableKeys[rand.IntN(len(BaseLayerNonPrintableKeys))]
	case LowerLayerNonPrintableWeightIndex:
		res.Key = LowerLayerNonPrintableKeys[rand.IntN(len(LowerLayerNonPrintableKeys))]
	case CustomCharsWeightIndex:
		if len(CustomChars) > 0 {
			err := setKeyComboBasedOnRuneAndLayout(&res, CustomChars[rand.IntN(len(CustomChars))])
			if err != nil {
				return KeyCombo{}, fmt.Errorf("error with Custom []rune: %w", err)
			}
		}
	case CustomKeysWeightIndex:
		if len(CustomKeys) > 0 {
			res.Key = CustomKeys[rand.IntN(len(CustomKeys))]
		}
	}

	switch modIndex {
	case NoModifiersWeightIndex:
	case ControlWeightIndex:
		res.Control = true
	case AltWeightIndex:
		res.Alt = true
	case ShiftWeightIndex:
		if baseIndex == BaseLayerNonPrintableWeightIndex || baseIndex == LowerLayerNonPrintableWeightIndex {
			res.Shift = true
		}
	case MetaWeightIndex:
		res.Meta = true
	case ControlAltWeightIndex:
		res.Control = true
		res.Alt = true
	case ControlShiftWeightIndex:
		res.Control = true
		res.Shift = true
	case ControlMetaWeightIndex:
		res.Control = true
		res.Meta = true
	case AltShiftWeightIndex:
		res.Alt = true
		res.Shift = true
	case AltMetaWeightIndex:
		res.Alt = true
		res.Meta = true
	case ShiftMetaWeightIndex:
		res.Shift = true
		res.Meta = true
	case ControlAltShiftWeightIndex:
		res.Control = true
		res.Alt = true
		res.Shift = true
	case ControlAltMetaWeightIndex:
		res.Control = true
		res.Alt = true
		res.Meta = true
	case ControlShiftMetaWeightIndex:
		res.Control = true
		res.Shift = true
		res.Meta = true
	case AltShiftMetaWeightIndex:
		res.Alt = true
		res.Shift = true
		res.Meta = true
	case ControlAltShiftMetaWeightIndex:
		res.Control = true
		res.Alt = true
		res.Shift = true
		res.Meta = true
	}

	res.setStringToDraw()
	return res, nil
}

func setKeyComboBasedOnRuneAndLayout(kc *KeyCombo, r rune) error {
	switch layout {
	case USLayout:
		keyWithShift, exist := USKeyMap[r]
		if !exist {
			return fmt.Errorf("no such rune in the USKeyMap: %c", r)
		}
		kc.Key = keyWithShift.Key
		kc.Shift = keyWithShift.Shift
		kc.NormalRune = USNormal[kc.Key]
		kc.ShiftedRune = USShifted[kc.Key]
	case UALayout:
		keyWithShift, exist := UAKeyMap[r]
		if !exist {
			return fmt.Errorf("no such rune in the UAKeyMap: %c", r)
		}
		kc.Key = keyWithShift.Key
		kc.Shift = keyWithShift.Shift
		kc.NormalRune = UANormal[kc.Key]
		kc.ShiftedRune = UAShifted[kc.Key]
	case RULayout:
		keyWithShift, exist := RUKeyMap[r]
		if !exist {
			return fmt.Errorf("no such rune in the RUKeyMap: %c", r)
		}
		kc.Key = keyWithShift.Key
		kc.Shift = keyWithShift.Shift
		kc.NormalRune = RUNormal[kc.Key]
		kc.ShiftedRune = RUShifted[kc.Key]
	}
	return nil
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
