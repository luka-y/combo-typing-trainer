package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	FontSize          int
	CurrentLayoutName string

	BaseWeightLowerLetters      int
	BaseWeightUpperLetters      int
	BaseWeightLowerSymbols      int
	BaseWeightUpperSymbols      int
	BaseWeightDigits            int
	BaseWeightNonPrintableKeys1 int
	BaseWeightNonPrintableKeys2 int
	BaseWeightCustomChars       int
	BaseWeightCustomKeys        int

	ModWeightNoModifiers         int
	ModWeightControl             int
	ModWeightAlt                 int
	ModWeightShift               int
	ModWeightMeta                int
	ModWeightControlAlt          int
	ModWeightControlShift        int
	ModWeightControlMeta         int
	ModWeightAltShift            int
	ModWeightAltMeta             int
	ModWeightShiftMeta           int
	ModWeightControlAltShift     int
	ModWeightControlAltMeta      int
	ModWeightControlShiftMeta    int
	ModWeightAltShiftMeta        int
	ModWeightControlAltShiftMeta int

	NonPrintableKeys1 []ebiten.Key
	NonPrintableKeys2 []ebiten.Key
	CustomChars       []rune
	CustomKeys        []ebiten.Key

	UpcomingComboColor      color.RGBA
	CurrentComboColor       color.RGBA
	CorrectPastComboColor   color.RGBA
	IncorrectPastComboColor color.RGBA
	BackgroundColor         color.RGBA

	InputLayouts []InputLayout
)
