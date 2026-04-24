package main

import "github.com/hajimehoshi/ebiten/v2"

// Probability weights for modifier keys (keep int)
const NoModifiersWeight = 25
const ControlWeight = 2
const AltWeight = 2
const ShiftWeight = 2 //non-printable only
const MetaWeight = 2
const ControlAltWeight = 1
const ControlShiftWeight = 1
const ControlMetaWeight = 0
const AltShiftWeight = 1
const AltMetaWeight = 1
const ShiftMetaWeight = 0
const ControlAltShiftWeight = 1
const ControlAltMetaWeight = 0
const ControlShiftMetaWeight = 0
const AltShiftMetaWeight = 0
const ControlAltShiftMetaWeight = 0

type ModWeightIndex int

const EmptyModWeightIndex ModWeightIndex = -1
const (
	NoModifiersWeightIndex ModWeightIndex = iota
	ControlWeightIndex
	AltWeightIndex
	ShiftWeightIndex
	MetaWeightIndex
	ControlAltWeightIndex
	ControlShiftWeightIndex
	ControlMetaWeightIndex
	AltShiftWeightIndex
	AltMetaWeightIndex
	ShiftMetaWeightIndex
	ControlAltShiftWeightIndex
	ControlAltMetaWeightIndex
	ControlShiftMetaWeightIndex
	AltShiftMetaWeightIndex
	ControlAltShiftMetaWeightIndex
)

// Probability weights for base rune/key (keep int)
const LowerLettersWeight = 20
const UpperLettersWeight = 10
const LowerSymbolsWeight = 10
const UpperSymbolsWeight = 10
const DigitsWeight = 10
const BaseLayerNonPrintableWeight = 10
const LowerLayerNonPrintableWeight = 5
const CustomCharsWeight = 0
const CustomKeysWeight = 0

type BaseWeightIndex int

const EmptyBaseWeightIndex BaseWeightIndex = -1
const (
	LowerLettersWeightIndex BaseWeightIndex = iota
	UpperLettersWeightIndex
	LowerSymbolsWeightIndex
	UpperSymbolsWeightIndex
	DigitsWeightIndex
	BaseLayerNonPrintableWeightIndex
	LowerLayerNonPrintableWeightIndex
	CustomCharsWeightIndex
	CustomKeysWeightIndex
)

// Add chars to emphasize to the CustomChars slice, make CustomCharsWeight non-zero for them to repeat more often.
var CustomChars = []rune{}

// Same, but for keys. Add only keys from BaseLayerNonPrintableKeys or LowerLayerNonPrintableKeys. You would not be able to type Shift, Control, Alt, Meta, or Backspace. If key correspond to a character, add a char to CustomChars instead.
var CustomKeys = []ebiten.Key{}

var BaseLayerNonPrintableKeys = []ebiten.Key{
	ebiten.KeyEscape, ebiten.KeyTab, ebiten.KeyCapsLock,
	ebiten.KeyDelete, ebiten.KeyEnter, ebiten.KeySpace,
	ebiten.KeyInsert, ebiten.KeyHome, ebiten.KeyEnd,
	ebiten.KeyPageUp, ebiten.KeyPageDown,
	ebiten.KeyArrowUp, ebiten.KeyArrowDown, ebiten.KeyArrowLeft, ebiten.KeyArrowRight,
}

var LowerLayerNonPrintableKeys = []ebiten.Key{
	ebiten.KeyF1, ebiten.KeyF2, ebiten.KeyF3, ebiten.KeyF4,
	ebiten.KeyF5, ebiten.KeyF6, ebiten.KeyF7, ebiten.KeyF8,
	ebiten.KeyF9, ebiten.KeyF10, ebiten.KeyF11, ebiten.KeyF12,
	ebiten.KeyF13, ebiten.KeyF14, ebiten.KeyF15, ebiten.KeyF16,
	ebiten.KeyPrintScreen, ebiten.KeyScrollLock, ebiten.KeyPause,
}
