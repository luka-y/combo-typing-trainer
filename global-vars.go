package main

import (
	"embed"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

//go:embed assets/*
var assets embed.FS

var (
	ScreenWidth  int
	ScreenHeight int

	FontFace   font.Face
	FontDrawer *font.Drawer

	GibberishYPos int
	InputYPos     int

	ConfigErr error

	BaseCumulativeDistribution     []float64
	ModifierCumulativeDistribution []float64

	BaseCategories []BaseCategory
	ModCategories  []ModCategory

	Layouts       []Layout
	CurrentLayout Layout

	TrimDistanceBeforeCurrent             = 100
	GenerateGibberishDistanceAfterCurrent = 100
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
