package main

import (
	"image/color"
	"log"
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
	PressedCombos []KeyCombo
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
		g.PressedCombos = append(g.PressedCombos, KeyCombo{
			Key:         k,
			NormalRune:  normal,
			ShiftedRune: shifted,

			Shift:   shiftPressed,
			Control: controlPressed,
			Alt:     altPressed,
			Meta:    metaPressed,
		})
	}

	if len(g.PressedCombos) > MaxLen {
		g.PressedCombos = g.PressedCombos[len(g.PressedCombos)-MaxLen:]
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
	keysString := ""
	for _, combo := range g.PressedCombos {
		keysString += getStringFromCombo(combo) + ", "
	}
	text.Draw(screen, keysString, FontFace, 0, FontSize, color.White)

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

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("combo-typing-trainer")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
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
