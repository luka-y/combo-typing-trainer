package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"unicode"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var ScreenHeight = 50

const TrimDistanceBeforeCurrent = 100
const GenerateGibberishDistanceAfterCurrent = 100

var BaseCumulativeDistribution []float64
var ModifierCumulativeDistribution []float64

var FontFace font.Face
var FontDrawer *font.Drawer

var GibberishYPos int
var InputYPos int

type KeyCombo struct {
	Key       ebiten.Key
	LowerRune rune
	UpperRune rune

	Shift   bool
	Control bool
	Alt     bool
	Meta    bool

	StrToDraw string
}

func (kc *KeyCombo) setStringToDraw() {
	if kc.Shift && !kc.Control && !kc.Alt && !kc.Meta && kc.UpperRune != 0 {
		kc.StrToDraw = string(kc.UpperRune)
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
	if kc.LowerRune != 0 {
		res += string(kc.LowerRune)
	} else {
		res += kc.Key.String()
	}
	kc.StrToDraw = res
}

type Game struct {
	InputStream     []KeyCombo
	GibberishStream []KeyCombo

	ScreenImg *ebiten.Image

	TickCounter int
}

func (g *Game) Update() error {
	if g.TickCounter == 0 {
		g.TickCounter++
		return g.FirstUpdateCall()
	}
	g.TickCounter++

	changeThisFrame := false

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
		changeThisFrame = true
		normal, shifted := CurrentLayout.KeyMap[KeyWithShift{k, false}], CurrentLayout.KeyMap[KeyWithShift{k, true}]
		inputCombo := KeyCombo{
			Key:       k,
			LowerRune: normal,
			UpperRune: shifted,

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
	if lenDiff < GenerateGibberishDistanceAfterCurrent {
		changeThisFrame = true
		for i := 0; i < GenerateGibberishDistanceAfterCurrent-lenDiff; i++ {
			randomCombo, err := getRandomCombo(BaseCumulativeDistribution, ModifierCumulativeDistribution)
			if err != nil {
				return fmt.Errorf("err with getRandomCombo in Update: %w", err)
			}
			g.GibberishStream = append(g.GibberishStream, randomCombo)
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) && len(g.InputStream) > 0 {
		changeThisFrame = true
		g.InputStream = g.InputStream[0 : len(g.InputStream)-1]
	}

	if len(g.InputStream) > TrimDistanceBeforeCurrent {
		changeThisFrame = true
		diff := len(g.InputStream) - TrimDistanceBeforeCurrent
		g.InputStream = g.InputStream[diff:]
		g.GibberishStream = g.GibberishStream[diff:]
	}

	if changeThisFrame {
		g.UpdateScreenImg()
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

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.ScreenImg, &ebiten.DrawImageOptions{})
}

func (g *Game) UpdateScreenImg() {
	if g.ScreenImg == nil {
		g.ScreenImg = ebiten.NewImage(ScreenWidth, ScreenHeight)
	}
	g.ScreenImg.Fill(BackgroundColor)

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
			text.Draw(g.ScreenImg, g.InputStream[i].StrToDraw+spaceBetweenCombos, FontFace, currentScreenXPosition, InputYPos, inputColor)
		}
		text.Draw(g.ScreenImg, g.GibberishStream[i].StrToDraw+spaceBetweenCombos, FontFace, currentScreenXPosition, GibberishYPos, gibberishColor)

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
	err := setCumulativeDistributions()
	if err != nil {
		return err
	}
	gibberishStream := make([]KeyCombo, GenerateGibberishDistanceAfterCurrent)
	for i := 0; i < GenerateGibberishDistanceAfterCurrent; i++ {
		gibberishStream[i], err = getRandomCombo(BaseCumulativeDistribution, ModifierCumulativeDistribution)
		if err != nil {
			return fmt.Errorf("error making random combo: %w", err)
		}
	}
	g.GibberishStream = gibberishStream

	g.UpdateScreenImg()
	return nil
}

func main() {
	err := InitLayout()
	if err != nil {
		log.Fatal(err)
	}

	FontFace, err = getFaceFromPath("assets/JetBrainsMono-Regular.ttf", FontSize, 72)
	if err != nil {
		log.Fatal(err)
	}
	FontDrawer = &font.Drawer{Face: FontFace}

	InputYPos = FontSize + 2
	GibberishYPos = ScreenHeight - int(FontFace.Metrics().Descent.Ceil()) - 2

	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("combo-typing-trainer")

	if err = ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

func setCumulativeDistributions() error {
	var err error

	baseWeights := make([]int, len(BaseCategories))
	for i, bc := range BaseCategories {
		baseWeights[i] = bc.Weight
	}
	BaseCumulativeDistribution, err = convertWeightsToCumulativeDistribution(baseWeights)
	if err != nil {
		return fmt.Errorf("error making cumulative distribution for base rune/key: %w", err)
	}

	modWeights := make([]int, len(ModCategories))
	for i, mc := range ModCategories {
		modWeights[i] = mc.Weight
	}
	ModifierCumulativeDistribution, err = convertWeightsToCumulativeDistribution(modWeights)
	if err != nil {
		return fmt.Errorf("error making cumulative distribution for modifiers: %w", err)
	}

	return nil
}

func getRandomCombo(baseCumulativeDistribution, modifierCumulativeDistribution []float64) (KeyCombo, error) {
	var res = KeyCombo{}

	finalBaseCategory := BaseCategory{}
	randFloat := rand.Float64()
	for i, num := range baseCumulativeDistribution {
		if num >= randFloat {
			finalBaseCategory = BaseCategories[i]
			break
		}
	}
	err := finalBaseCategory.Handler(&res)
	if err != nil {
		return KeyCombo{}, fmt.Errorf("base category error: %w", err)
	}

	finalModCategory := ModCategory{}
	randFloat = rand.Float64()
	for i, num := range modifierCumulativeDistribution {
		if num >= randFloat {
			finalModCategory = ModCategories[i]
			break
		}
	}
	err = finalModCategory.Handler(&res)
	if err != nil {
		return KeyCombo{}, fmt.Errorf("mod category error: %w", err)
	}

	res.setStringToDraw()
	return res, nil
}

func (kc *KeyCombo) setKeyComboBasedOnRune(r rune) {
	keyWithShift := CurrentLayout.ReverseKeyMap[r]
	kc.Key = keyWithShift.Key
	kc.Shift = keyWithShift.Shift
	kc.LowerRune = CurrentLayout.KeyMap[KeyWithShift{kc.Key, false}]
	kc.UpperRune = CurrentLayout.KeyMap[KeyWithShift{kc.Key, true}]
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
	probabilities[len(probabilities)-1] = 1
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

type ModCategory struct {
	Name    string
	Weight  int
	Handler func(*KeyCombo) error
}

var ModCategories = []ModCategory{
	{
		Name:   "No modifiers",
		Weight: 25,
		Handler: func(kc *KeyCombo) error {
			return nil
		},
	},
	{
		Name:   "Control",
		Weight: 2,
		Handler: func(kc *KeyCombo) error {
			kc.Control = true
			return nil
		},
	},
	{
		Name:   "Alt",
		Weight: 2,
		Handler: func(kc *KeyCombo) error {
			kc.Alt = true
			return nil
		},
	},
	{
		Name:   "Shift",
		Weight: 2,
		Handler: func(kc *KeyCombo) error {
			if kc.LowerRune == 0 && kc.UpperRune == 0 {
				kc.Shift = true
			}
			return nil
		},
	},
	{
		Name:   "Meta",
		Weight: 2,
		Handler: func(kc *KeyCombo) error {
			kc.Meta = true
			return nil
		},
	},
	{
		Name:   "Control+Alt",
		Weight: 1,
		Handler: func(kc *KeyCombo) error {
			kc.Control = true
			kc.Alt = true
			return nil
		},
	},
	{
		Name:   "Control+Shift",
		Weight: 1,
		Handler: func(kc *KeyCombo) error {
			kc.Control = true
			kc.Shift = true
			return nil
		},
	},
	{
		Name:   "Control+Meta",
		Weight: 0,
		Handler: func(kc *KeyCombo) error {
			kc.Control = true
			kc.Meta = true
			return nil
		},
	},
	{
		Name:   "Alt+Shift",
		Weight: 1,
		Handler: func(kc *KeyCombo) error {
			kc.Alt = true
			kc.Shift = true
			return nil
		},
	},
	{
		Name:   "Alt+Meta",
		Weight: 1,
		Handler: func(kc *KeyCombo) error {
			kc.Alt = true
			kc.Meta = true
			return nil
		},
	},
	{
		Name:   "Shift+Meta",
		Weight: 0,
		Handler: func(kc *KeyCombo) error {
			kc.Shift = true
			kc.Meta = true
			return nil
		},
	},
	{
		Name:   "Control+Alt+Shift",
		Weight: 1,
		Handler: func(kc *KeyCombo) error {
			kc.Control = true
			kc.Alt = true
			kc.Shift = true
			return nil
		},
	},
	{
		Name:   "Control+Alt+Meta",
		Weight: 0,
		Handler: func(kc *KeyCombo) error {
			kc.Control = true
			kc.Alt = true
			kc.Meta = true
			return nil
		},
	},
	{
		Name:   "Control+Shift+Meta",
		Weight: 0,
		Handler: func(kc *KeyCombo) error {
			kc.Control = true
			kc.Shift = true
			kc.Meta = true
			return nil
		},
	},
	{
		Name:   "Alt+Shift+Meta",
		Weight: 0,
		Handler: func(kc *KeyCombo) error {
			kc.Alt = true
			kc.Shift = true
			kc.Meta = true
			return nil
		},
	},
	{
		Name:   "Control+Alt+Shift+Meta",
		Weight: 0,
		Handler: func(kc *KeyCombo) error {
			kc.Control = true
			kc.Alt = true
			kc.Shift = true
			kc.Meta = true
			return nil
		},
	},
}

type BaseCategory struct {
	Name      string
	Weight    int
	Handler   func(*KeyCombo) error
	Validator func() bool
}

var BaseCategories = []BaseCategory{
	{
		Name:   "Lower Letters",
		Weight: 20,
		Handler: func(kc *KeyCombo) error {
			kc.setKeyComboBasedOnRune(CurrentLayout.LowerLetters[rand.IntN(len(CurrentLayout.LowerLetters))])
			return nil
		},
		Validator: func() bool {
			return len(CurrentLayout.LowerLetters) > 0
		},
	},
	{
		Name:   "Upper Letters",
		Weight: 10,
		Handler: func(kc *KeyCombo) error {
			kc.setKeyComboBasedOnRune(CurrentLayout.UpperLetters[rand.IntN(len(CurrentLayout.UpperLetters))])
			return nil
		},
		Validator: func() bool {
			return len(CurrentLayout.UpperLetters) > 0
		},
	},
	{
		Name:   "Lower Symbols",
		Weight: 10,
		Handler: func(kc *KeyCombo) error {
			kc.setKeyComboBasedOnRune(CurrentLayout.LowerSymbols[rand.IntN(len(CurrentLayout.LowerSymbols))])
			return nil
		},
		Validator: func() bool {
			return len(CurrentLayout.LowerSymbols) > 0
		},
	},
	{
		Name:   "Upper Symbols",
		Weight: 10,
		Handler: func(kc *KeyCombo) error {
			kc.setKeyComboBasedOnRune(CurrentLayout.UpperSymbols[rand.IntN(len(CurrentLayout.UpperSymbols))])
			return nil
		},
		Validator: func() bool {
			return len(CurrentLayout.UpperSymbols) > 0
		},
	},
	{
		Name:   "Digits",
		Weight: 10,
		Handler: func(kc *KeyCombo) error {
			kc.setKeyComboBasedOnRune(CurrentLayout.Digits[rand.IntN(len(CurrentLayout.Digits))])
			return nil
		},
		Validator: func() bool {
			return len(CurrentLayout.Digits) > 0
		},
	},
	{
		Name:   "Non-Printable Keys 1",
		Weight: 10,
		Handler: func(kc *KeyCombo) error {
			kc.Key = NonPrintableKeys1[rand.IntN(len(NonPrintableKeys1))]
			return nil
		},
		Validator: func() bool {
			return len(NonPrintableKeys1) > 0
		},
	},
	{
		Name:   "Non-Printable Keys 2",
		Weight: 5,
		Handler: func(kc *KeyCombo) error {
			kc.Key = NonPrintableKeys2[rand.IntN(len(NonPrintableKeys2))]
			return nil
		},
		Validator: func() bool {
			return len(NonPrintableKeys2) > 0
		},
	},
	{
		Name:   "Custom Chars",
		Weight: 0,
		Handler: func(kc *KeyCombo) error {
			kc.setKeyComboBasedOnRune(CustomChars[rand.IntN(len(CustomChars))])
			return nil
		},
		Validator: func() bool {
			return len(CustomChars) > 0
		},
	},
	{
		Name:   "Custom Keys",
		Weight: 0,
		Handler: func(kc *KeyCombo) error {
			kc.Key = CustomKeys[rand.IntN(len(CustomKeys))]
			return nil
		},
		Validator: func() bool {
			return len(CustomKeys) > 0
		},
	},
}

type KeyWithShift struct {
	Key   ebiten.Key
	Shift bool
}

type Layout struct {
	Name   string
	KeyMap map[KeyWithShift]rune

	ReverseKeyMap map[rune]KeyWithShift

	LowerLetters []rune
	UpperLetters []rune
	Digits       []rune
	LowerSymbols []rune
	UpperSymbols []rune
}

var Layouts []Layout
var CurrentLayout Layout

func InitLayout() error {
	usLayout, err := GetLayoutFromKeyMap("US", InputUSKeyMap)
	if err != nil {
		return fmt.Errorf("err getting layout from the US key map: %w", err)
	}
	Layouts = append(Layouts, usLayout)
	ruLayout, err := GetLayoutFromKeyMap("RU", InputRUKeyMap)
	if err != nil {
		return fmt.Errorf("err getting layout from the RU key map: %w", err)
	}
	Layouts = append(Layouts, ruLayout)
	uaLayout, err := GetLayoutFromKeyMap("UA", InputUAKeyMap)
	if err != nil {
		return fmt.Errorf("err getting layout from the UA key map: %w", err)
	}
	Layouts = append(Layouts, uaLayout)

	for _, l := range Layouts {
		if l.Name == CurrentLayoutName {
			CurrentLayout = l
		}
	}
	if CurrentLayout.Name == "" {
		if CurrentLayoutName == "" {
			return fmt.Errorf("no layout is set")
		}
		return fmt.Errorf("no such layout: %s", CurrentLayoutName)
	}

	for _, r := range CustomChars {
		_, exist := CurrentLayout.ReverseKeyMap[r]
		if !exist {
			return fmt.Errorf("custom layout contains a char '%c' that is absent in the current layout", r)
		}
	}

	for _, category := range BaseCategories {
		if category.Weight > 0 && !category.Validator() {
			return fmt.Errorf("config error: base category \"%s\" weigh is more than zero while slice it is pulling from is empty", category.Name)
		}
	}
	return nil
}

func GetLayoutFromKeyMap(name string, keyMap map[KeyWithShift]rune) (Layout, error) {
	res := Layout{}
	res.Name = name
	res.KeyMap = keyMap

	reverseKeyMap := make(map[rune]KeyWithShift)
	var lowerLetters, upperLetters, digits, lowerSymbols, upperSymbols []rune

	for keyWithShift, r := range keyMap {
		if _, alreadyExist := reverseKeyMap[r]; alreadyExist {
			return Layout{}, fmt.Errorf("duplicate rune in the input layout map")
		}
		reverseKeyMap[r] = keyWithShift

		shifted := keyWithShift.Shift
		if unicode.Is(unicode.L, r) && !unicode.Is(unicode.Lm, r) && !shifted {
			lowerLetters = append(lowerLetters, r)
		} else if unicode.Is(unicode.L, r) && !unicode.Is(unicode.Lm, r) && shifted {
			upperLetters = append(upperLetters, r)
		} else if unicode.IsDigit(r) {
			digits = append(digits, r)
		} else {
			if !shifted {
				lowerSymbols = append(lowerSymbols, r)
			}
			if shifted {
				upperSymbols = append(upperSymbols, r)
			}
		}
	}

	res.ReverseKeyMap = reverseKeyMap
	res.LowerLetters = lowerLetters
	res.UpperLetters = upperLetters
	res.Digits = digits
	res.LowerSymbols = lowerSymbols
	res.UpperSymbols = upperSymbols

	return res, nil
}
