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
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/font/opentype"
)

type Game struct {
	InputStream     []KeyCombo
	GibberishStream []KeyCombo

	ScreenImg *ebiten.Image

	TickCounter int
}

func (g *Game) Update() error {
	if g.TickCounter == 0 {
		g.TickCounter++
		if ConfigErr != nil {
			return g.FirstUpdateCallConfigErr()
		}
		return g.FirstUpdateCall()
	}
	g.TickCounter++

	if ConfigErr != nil {
		return ConfigErrUpdate()
	}

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

func ConfigErrUpdate() error {
	if len(inpututil.AppendJustPressedKeys([]ebiten.Key{})) > 0 {
		return ConfigErr
	}
	return nil
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

func (g *Game) Draw(screen *ebiten.Image) {
	if g.ScreenImg != nil {
		screen.DrawImage(g.ScreenImg, &ebiten.DrawImageOptions{})
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
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

func (g *Game) FirstUpdateCallConfigErr() error {
	g.ScreenImg = ebiten.NewImage(ScreenWidth, ScreenHeight)

	text.Draw(g.ScreenImg, "Config error: "+ConfigErr.Error(), FontFace, 0, InputYPos, color.RGBA{255, 100, 100, 255})
	text.Draw(g.ScreenImg, "Press any key to exit", FontFace, 0, GibberishYPos, color.RGBA{255, 255, 255, 255})
	return nil
}

func main() {
	err := ParseConfig()
	ConfigErr = err

	if FontSize <= 0 {
		FontFace = basicfont.Face7x13
		FontSize = 7
	} else {
		FontFace, err = getFaceFromPath("assets/JetBrainsMono-Regular.ttf", float64(FontSize), 72)
		if err != nil {
			log.Fatal(err)
		}
	}
	FontDrawer = &font.Drawer{Face: FontFace}

	ScreenWidth, _ = ebiten.Monitor().Size()
	ScreenHeight = int(float64(FontFace.Metrics().Ascent.Ceil()+FontFace.Metrics().Descent.Ceil()) * 2.15)

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
