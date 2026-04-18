package main

import "github.com/hajimehoshi/ebiten/v2"

// Probability weights for modifier keys (keep int)
const NoModifiersWeight = 10
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
const LowercaseLettersWeight = 20
const UppercaseLettersWeight = 10
const SymbolsWeight = 10
const NumbersWeight = 10
const BaseLayerNonPrintableWeight = 10
const LowerLayerNonPrintableWeight = 5
const CustomWeight = 0

type BaseWeightIndex int

const EmptyBaseWeightIndex BaseWeightIndex = -1
const (
	LowercaseLettersWeightIndex BaseWeightIndex = iota
	UppercaseLettersWeightIndex
	SymbolsWeightIndex
	NumbersWeightIndex
	BaseLayerNonPrintableWeightIndex
	LowerLayerNonPrintableWeightIndex
	CustomWeightIndex
)

// Add chars to emphasize to the Custom slice, make CustomWeight non-zero for them to repeat more often.
var Custom = []rune{}

var USLower = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
var USUpper = []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
var USSymbols = []rune{
	'!', '@', '#', '$', '%', '^', '&', '*', '(', ')',
	'`', '~', '-', '_', '=', '+',
	'[', '{', ']', '}', '\\', '|',
	';', ':', '\'', '"',
	',', '<', '.', '>', '/', '?',
}

var Numbers = []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

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

var RULower = []rune{'а', 'б', 'в', 'г', 'д', 'е', 'ё', 'ж', 'з', 'и', 'й', 'к', 'л', 'м', 'н', 'о', 'п', 'р', 'с', 'т', 'у', 'ф', 'х', 'ц', 'ч', 'ш', 'щ', 'ъ', 'ы', 'ь', 'э', 'ю', 'я'}
var RUUpper = []rune{'А', 'Б', 'В', 'Г', 'Д', 'Е', 'Ё', 'Ж', 'З', 'И', 'Й', 'К', 'Л', 'М', 'Н', 'О', 'П', 'Р', 'С', 'Т', 'У', 'Ф', 'Х', 'Ц', 'Ч', 'Ш', 'Щ', 'Ъ', 'Ы', 'Ь', 'Э', 'Ю', 'Я'}
var RUSymbols = []rune{
	'!', '"', '№', ';', '%', ':', '?', '*', '(', ')',
	'-', '_', '=', '+', '\\', '/', '.', ',',
}

var UALower = []rune{'а', 'б', 'в', 'г', 'ґ', 'д', 'е', 'є', 'ж', 'з', 'и', 'і', 'ї', 'й', 'к', 'л', 'м', 'н', 'о', 'п', 'р', 'с', 'т', 'у', 'ф', 'х', 'ц', 'ч', 'ш', 'щ', 'ь', 'ю', 'я'}
var UAUpper = []rune{'А', 'Б', 'В', 'Г', 'Ґ', 'Д', 'Е', 'Є', 'Ж', 'З', 'И', 'І', 'Ї', 'Й', 'К', 'Л', 'М', 'Н', 'О', 'П', 'Р', 'С', 'Т', 'У', 'Ф', 'Х', 'Ц', 'Ч', 'Ш', 'Щ', 'Ь', 'Ю', 'Я'}
var UASymbols = []rune{
	'!', '"', '№', ';', '%', ':', '?', '*', '(', ')',
	'\'', 'ʼ', '-', '_', '=', '+', '.', ',',
}
