package main

import "github.com/hajimehoshi/ebiten/v2"

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
	ebiten.KeyDelete, ebiten.KeyEnter,
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
