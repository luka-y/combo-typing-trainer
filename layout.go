package main

import "github.com/hajimehoshi/ebiten/v2"

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
	Numbers      []rune
	LowerSymbols []rune
	UpperSymbols []rune
}

func init() {

}

var ActualUSKeyMap = map[KeyWithShift]rune{
	{Key: ebiten.KeyA, Shift: false}: 'a', {Key: ebiten.KeyA, Shift: true}: 'A',
	{Key: ebiten.KeyB, Shift: false}: 'b', {Key: ebiten.KeyB, Shift: true}: 'B',
	{Key: ebiten.KeyC, Shift: false}: 'c', {Key: ebiten.KeyC, Shift: true}: 'C',
	{Key: ebiten.KeyD, Shift: false}: 'd', {Key: ebiten.KeyD, Shift: true}: 'D',
	{Key: ebiten.KeyE, Shift: false}: 'e', {Key: ebiten.KeyE, Shift: true}: 'E',
	{Key: ebiten.KeyF, Shift: false}: 'f', {Key: ebiten.KeyF, Shift: true}: 'F',
	{Key: ebiten.KeyG, Shift: false}: 'g', {Key: ebiten.KeyG, Shift: true}: 'G',
	{Key: ebiten.KeyH, Shift: false}: 'h', {Key: ebiten.KeyH, Shift: true}: 'H',
	{Key: ebiten.KeyI, Shift: false}: 'i', {Key: ebiten.KeyI, Shift: true}: 'I',
	{Key: ebiten.KeyJ, Shift: false}: 'j', {Key: ebiten.KeyJ, Shift: true}: 'J',
	{Key: ebiten.KeyK, Shift: false}: 'k', {Key: ebiten.KeyK, Shift: true}: 'K',
	{Key: ebiten.KeyL, Shift: false}: 'l', {Key: ebiten.KeyL, Shift: true}: 'L',
	{Key: ebiten.KeyM, Shift: false}: 'm', {Key: ebiten.KeyM, Shift: true}: 'M',
	{Key: ebiten.KeyN, Shift: false}: 'n', {Key: ebiten.KeyN, Shift: true}: 'N',
	{Key: ebiten.KeyO, Shift: false}: 'o', {Key: ebiten.KeyO, Shift: true}: 'O',
	{Key: ebiten.KeyP, Shift: false}: 'p', {Key: ebiten.KeyP, Shift: true}: 'P',
	{Key: ebiten.KeyQ, Shift: false}: 'q', {Key: ebiten.KeyQ, Shift: true}: 'Q',
	{Key: ebiten.KeyR, Shift: false}: 'r', {Key: ebiten.KeyR, Shift: true}: 'R',
	{Key: ebiten.KeyS, Shift: false}: 's', {Key: ebiten.KeyS, Shift: true}: 'S',
	{Key: ebiten.KeyT, Shift: false}: 't', {Key: ebiten.KeyT, Shift: true}: 'T',
	{Key: ebiten.KeyU, Shift: false}: 'u', {Key: ebiten.KeyU, Shift: true}: 'U',
	{Key: ebiten.KeyV, Shift: false}: 'v', {Key: ebiten.KeyV, Shift: true}: 'V',
	{Key: ebiten.KeyW, Shift: false}: 'w', {Key: ebiten.KeyW, Shift: true}: 'W',
	{Key: ebiten.KeyX, Shift: false}: 'x', {Key: ebiten.KeyX, Shift: true}: 'X',
	{Key: ebiten.KeyY, Shift: false}: 'y', {Key: ebiten.KeyY, Shift: true}: 'Y',
	{Key: ebiten.KeyZ, Shift: false}: 'z', {Key: ebiten.KeyZ, Shift: true}: 'Z',
	{Key: ebiten.KeyBackquote, Shift: false}: '`', {Key: ebiten.KeyBackquote, Shift: true}: '~',
	{Key: ebiten.KeyBackslash, Shift: false}: '\\', {Key: ebiten.KeyBackslash, Shift: true}: '|',
	{Key: ebiten.KeyBracketLeft, Shift: false}: '[', {Key: ebiten.KeyBracketLeft, Shift: true}: '{',
	{Key: ebiten.KeyBracketRight, Shift: false}: ']', {Key: ebiten.KeyBracketRight, Shift: true}: '}',
	{Key: ebiten.KeyComma, Shift: false}: ',', {Key: ebiten.KeyComma, Shift: true}: '<',
	{Key: ebiten.KeyDigit0, Shift: false}: '0', {Key: ebiten.KeyDigit0, Shift: true}: ')',
	{Key: ebiten.KeyDigit1, Shift: false}: '1', {Key: ebiten.KeyDigit1, Shift: true}: '!',
	{Key: ebiten.KeyDigit2, Shift: false}: '2', {Key: ebiten.KeyDigit2, Shift: true}: '@',
	{Key: ebiten.KeyDigit3, Shift: false}: '3', {Key: ebiten.KeyDigit3, Shift: true}: '#',
	{Key: ebiten.KeyDigit4, Shift: false}: '4', {Key: ebiten.KeyDigit4, Shift: true}: '$',
	{Key: ebiten.KeyDigit5, Shift: false}: '5', {Key: ebiten.KeyDigit5, Shift: true}: '%',
	{Key: ebiten.KeyDigit6, Shift: false}: '6', {Key: ebiten.KeyDigit6, Shift: true}: '^',
	{Key: ebiten.KeyDigit7, Shift: false}: '7', {Key: ebiten.KeyDigit7, Shift: true}: '&',
	{Key: ebiten.KeyDigit8, Shift: false}: '8', {Key: ebiten.KeyDigit8, Shift: true}: '*',
	{Key: ebiten.KeyDigit9, Shift: false}: '9', {Key: ebiten.KeyDigit9, Shift: true}: '(',
	{Key: ebiten.KeyEqual, Shift: false}: '=', {Key: ebiten.KeyEqual, Shift: true}: '+',
	{Key: ebiten.KeyMinus, Shift: false}: '-', {Key: ebiten.KeyMinus, Shift: true}: '_',
	{Key: ebiten.KeyPeriod, Shift: false}: '.', {Key: ebiten.KeyPeriod, Shift: true}: '>',
	{Key: ebiten.KeyQuote, Shift: false}: '\'', {Key: ebiten.KeyQuote, Shift: true}: '"',
	{Key: ebiten.KeySemicolon, Shift: false}: ';', {Key: ebiten.KeySemicolon, Shift: true}: ':',
	{Key: ebiten.KeySlash, Shift: false}: '/', {Key: ebiten.KeySlash, Shift: true}: '?',
}

var ActualRUKeyMap = map[KeyWithShift]rune{
	{Key: ebiten.KeyA, Shift: false}: 'ф', {Key: ebiten.KeyA, Shift: true}: 'Ф',
	{Key: ebiten.KeyB, Shift: false}: 'и', {Key: ebiten.KeyB, Shift: true}: 'И',
	{Key: ebiten.KeyC, Shift: false}: 'с', {Key: ebiten.KeyC, Shift: true}: 'С',
	{Key: ebiten.KeyD, Shift: false}: 'в', {Key: ebiten.KeyD, Shift: true}: 'В',
	{Key: ebiten.KeyE, Shift: false}: 'у', {Key: ebiten.KeyE, Shift: true}: 'У',
	{Key: ebiten.KeyF, Shift: false}: 'а', {Key: ebiten.KeyF, Shift: true}: 'А',
	{Key: ebiten.KeyG, Shift: false}: 'п', {Key: ebiten.KeyG, Shift: true}: 'П',
	{Key: ebiten.KeyH, Shift: false}: 'р', {Key: ebiten.KeyH, Shift: true}: 'Р',
	{Key: ebiten.KeyI, Shift: false}: 'ш', {Key: ebiten.KeyI, Shift: true}: 'Ш',
	{Key: ebiten.KeyJ, Shift: false}: 'о', {Key: ebiten.KeyJ, Shift: true}: 'О',
	{Key: ebiten.KeyK, Shift: false}: 'л', {Key: ebiten.KeyK, Shift: true}: 'Л',
	{Key: ebiten.KeyL, Shift: false}: 'д', {Key: ebiten.KeyL, Shift: true}: 'Д',
	{Key: ebiten.KeyM, Shift: false}: 'ь', {Key: ebiten.KeyM, Shift: true}: 'Ь',
	{Key: ebiten.KeyN, Shift: false}: 'т', {Key: ebiten.KeyN, Shift: true}: 'Т',
	{Key: ebiten.KeyO, Shift: false}: 'щ', {Key: ebiten.KeyO, Shift: true}: 'Щ',
	{Key: ebiten.KeyP, Shift: false}: 'з', {Key: ebiten.KeyP, Shift: true}: 'З',
	{Key: ebiten.KeyQ, Shift: false}: 'й', {Key: ebiten.KeyQ, Shift: true}: 'Й',
	{Key: ebiten.KeyR, Shift: false}: 'к', {Key: ebiten.KeyR, Shift: true}: 'К',
	{Key: ebiten.KeyS, Shift: false}: 'ы', {Key: ebiten.KeyS, Shift: true}: 'Ы',
	{Key: ebiten.KeyT, Shift: false}: 'е', {Key: ebiten.KeyT, Shift: true}: 'Е',
	{Key: ebiten.KeyU, Shift: false}: 'г', {Key: ebiten.KeyU, Shift: true}: 'Г',
	{Key: ebiten.KeyV, Shift: false}: 'м', {Key: ebiten.KeyV, Shift: true}: 'М',
	{Key: ebiten.KeyW, Shift: false}: 'ц', {Key: ebiten.KeyW, Shift: true}: 'Ц',
	{Key: ebiten.KeyX, Shift: false}: 'ч', {Key: ebiten.KeyX, Shift: true}: 'Ч',
	{Key: ebiten.KeyY, Shift: false}: 'н', {Key: ebiten.KeyY, Shift: true}: 'Н',
	{Key: ebiten.KeyZ, Shift: false}: 'я', {Key: ebiten.KeyZ, Shift: true}: 'Я',
	{Key: ebiten.KeyBackquote, Shift: false}: 'ё', {Key: ebiten.KeyBackquote, Shift: true}: 'Ё',
	{Key: ebiten.KeyBackslash, Shift: false}: '\\', {Key: ebiten.KeyBackslash, Shift: true}: '/',
	{Key: ebiten.KeyBracketLeft, Shift: false}: 'х', {Key: ebiten.KeyBracketLeft, Shift: true}: 'Х',
	{Key: ebiten.KeyBracketRight, Shift: false}: 'ъ', {Key: ebiten.KeyBracketRight, Shift: true}: 'Ъ',
	{Key: ebiten.KeyComma, Shift: false}: 'б', {Key: ebiten.KeyComma, Shift: true}: 'Б',
	{Key: ebiten.KeyDigit0, Shift: false}: '0', {Key: ebiten.KeyDigit0, Shift: true}: ')',
	{Key: ebiten.KeyDigit1, Shift: false}: '1', {Key: ebiten.KeyDigit1, Shift: true}: '!',
	{Key: ebiten.KeyDigit2, Shift: false}: '2', {Key: ebiten.KeyDigit2, Shift: true}: '"',
	{Key: ebiten.KeyDigit3, Shift: false}: '3', {Key: ebiten.KeyDigit3, Shift: true}: '№',
	{Key: ebiten.KeyDigit4, Shift: false}: '4', {Key: ebiten.KeyDigit4, Shift: true}: ';',
	{Key: ebiten.KeyDigit5, Shift: false}: '5', {Key: ebiten.KeyDigit5, Shift: true}: '%',
	{Key: ebiten.KeyDigit6, Shift: false}: '6', {Key: ebiten.KeyDigit6, Shift: true}: ':',
	{Key: ebiten.KeyDigit7, Shift: false}: '7', {Key: ebiten.KeyDigit7, Shift: true}: '?',
	{Key: ebiten.KeyDigit8, Shift: false}: '8', {Key: ebiten.KeyDigit8, Shift: true}: '*',
	{Key: ebiten.KeyDigit9, Shift: false}: '9', {Key: ebiten.KeyDigit9, Shift: true}: '(',
	{Key: ebiten.KeyEqual, Shift: false}: '=', {Key: ebiten.KeyEqual, Shift: true}: '+',
	{Key: ebiten.KeyMinus, Shift: false}: '-', {Key: ebiten.KeyMinus, Shift: true}: '_',
	{Key: ebiten.KeyPeriod, Shift: false}: 'ю', {Key: ebiten.KeyPeriod, Shift: true}: 'Ю',
	{Key: ebiten.KeyQuote, Shift: false}: 'э', {Key: ebiten.KeyQuote, Shift: true}: 'Э',
	{Key: ebiten.KeySemicolon, Shift: false}: 'ж', {Key: ebiten.KeySemicolon, Shift: true}: 'Ж',
	{Key: ebiten.KeySlash, Shift: false}: '.', {Key: ebiten.KeySlash, Shift: true}: ',',
}

var ActualUAKeyMap = map[KeyWithShift]rune{
	{Key: ebiten.KeyA, Shift: false}: 'ф', {Key: ebiten.KeyA, Shift: true}: 'Ф',
	{Key: ebiten.KeyB, Shift: false}: 'и', {Key: ebiten.KeyB, Shift: true}: 'И',
	{Key: ebiten.KeyC, Shift: false}: 'с', {Key: ebiten.KeyC, Shift: true}: 'С',
	{Key: ebiten.KeyD, Shift: false}: 'в', {Key: ebiten.KeyD, Shift: true}: 'В',
	{Key: ebiten.KeyE, Shift: false}: 'у', {Key: ebiten.KeyE, Shift: true}: 'У',
	{Key: ebiten.KeyF, Shift: false}: 'а', {Key: ebiten.KeyF, Shift: true}: 'А',
	{Key: ebiten.KeyG, Shift: false}: 'п', {Key: ebiten.KeyG, Shift: true}: 'П',
	{Key: ebiten.KeyH, Shift: false}: 'р', {Key: ebiten.KeyH, Shift: true}: 'Р',
	{Key: ebiten.KeyI, Shift: false}: 'ш', {Key: ebiten.KeyI, Shift: true}: 'Ш',
	{Key: ebiten.KeyJ, Shift: false}: 'о', {Key: ebiten.KeyJ, Shift: true}: 'О',
	{Key: ebiten.KeyK, Shift: false}: 'л', {Key: ebiten.KeyK, Shift: true}: 'Л',
	{Key: ebiten.KeyL, Shift: false}: 'д', {Key: ebiten.KeyL, Shift: true}: 'Д',
	{Key: ebiten.KeyM, Shift: false}: 'ь', {Key: ebiten.KeyM, Shift: true}: 'Ь',
	{Key: ebiten.KeyN, Shift: false}: 'т', {Key: ebiten.KeyN, Shift: true}: 'Т',
	{Key: ebiten.KeyO, Shift: false}: 'щ', {Key: ebiten.KeyO, Shift: true}: 'Щ',
	{Key: ebiten.KeyP, Shift: false}: 'з', {Key: ebiten.KeyP, Shift: true}: 'З',
	{Key: ebiten.KeyQ, Shift: false}: 'й', {Key: ebiten.KeyQ, Shift: true}: 'Й',
	{Key: ebiten.KeyR, Shift: false}: 'к', {Key: ebiten.KeyR, Shift: true}: 'К',
	{Key: ebiten.KeyS, Shift: false}: 'і', {Key: ebiten.KeyS, Shift: true}: 'І',
	{Key: ebiten.KeyT, Shift: false}: 'е', {Key: ebiten.KeyT, Shift: true}: 'Е',
	{Key: ebiten.KeyU, Shift: false}: 'г', {Key: ebiten.KeyU, Shift: true}: 'Г',
	{Key: ebiten.KeyV, Shift: false}: 'м', {Key: ebiten.KeyV, Shift: true}: 'М',
	{Key: ebiten.KeyW, Shift: false}: 'ц', {Key: ebiten.KeyW, Shift: true}: 'Ц',
	{Key: ebiten.KeyX, Shift: false}: 'ч', {Key: ebiten.KeyX, Shift: true}: 'Ч',
	{Key: ebiten.KeyY, Shift: false}: 'н', {Key: ebiten.KeyY, Shift: true}: 'Н',
	{Key: ebiten.KeyZ, Shift: false}: 'я', {Key: ebiten.KeyZ, Shift: true}: 'Я',
	{Key: ebiten.KeyBackquote, Shift: false}: '\'', {Key: ebiten.KeyBackquote, Shift: true}: 'ʼ',
	{Key: ebiten.KeyBackslash, Shift: false}: 'ґ', {Key: ebiten.KeyBackslash, Shift: true}: 'Ґ',
	{Key: ebiten.KeyBracketLeft, Shift: false}: 'х', {Key: ebiten.KeyBracketLeft, Shift: true}: 'Х',
	{Key: ebiten.KeyBracketRight, Shift: false}: 'ї', {Key: ebiten.KeyBracketRight, Shift: true}: 'Ї',
	{Key: ebiten.KeyComma, Shift: false}: 'б', {Key: ebiten.KeyComma, Shift: true}: 'Б',
	{Key: ebiten.KeyDigit0, Shift: false}: '0', {Key: ebiten.KeyDigit0, Shift: true}: ')',
	{Key: ebiten.KeyDigit1, Shift: false}: '1', {Key: ebiten.KeyDigit1, Shift: true}: '!',
	{Key: ebiten.KeyDigit2, Shift: false}: '2', {Key: ebiten.KeyDigit2, Shift: true}: '"',
	{Key: ebiten.KeyDigit3, Shift: false}: '3', {Key: ebiten.KeyDigit3, Shift: true}: '№',
	{Key: ebiten.KeyDigit4, Shift: false}: '4', {Key: ebiten.KeyDigit4, Shift: true}: ';',
	{Key: ebiten.KeyDigit5, Shift: false}: '5', {Key: ebiten.KeyDigit5, Shift: true}: '%',
	{Key: ebiten.KeyDigit6, Shift: false}: '6', {Key: ebiten.KeyDigit6, Shift: true}: ':',
	{Key: ebiten.KeyDigit7, Shift: false}: '7', {Key: ebiten.KeyDigit7, Shift: true}: '?',
	{Key: ebiten.KeyDigit8, Shift: false}: '8', {Key: ebiten.KeyDigit8, Shift: true}: '*',
	{Key: ebiten.KeyDigit9, Shift: false}: '9', {Key: ebiten.KeyDigit9, Shift: true}: '(',
	{Key: ebiten.KeyEqual, Shift: false}: '=', {Key: ebiten.KeyEqual, Shift: true}: '+',
	{Key: ebiten.KeyMinus, Shift: false}: '-', {Key: ebiten.KeyMinus, Shift: true}: '_',
	{Key: ebiten.KeyPeriod, Shift: false}: 'ю', {Key: ebiten.KeyPeriod, Shift: true}: 'Ю',
	{Key: ebiten.KeyQuote, Shift: false}: 'є', {Key: ebiten.KeyQuote, Shift: true}: 'Є',
	{Key: ebiten.KeySemicolon, Shift: false}: 'ж', {Key: ebiten.KeySemicolon, Shift: true}: 'Ж',
	{Key: ebiten.KeySlash, Shift: false}: '.', {Key: ebiten.KeySlash, Shift: true}: ',',
}

//everything below needs to be outdated. Delete after all usages are removed. Until anything exist, the things isn't done.

var layout LayoutEnum

type LayoutEnum int

const (
	USLayout LayoutEnum = iota
	UALayout
	RULayout
)

var Numbers = []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

var USLower = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
var USUpper = []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
var USSymbols = []rune{
	'!', '@', '#', '$', '%', '^', '&', '*', '(', ')',
	'`', '~', '-', '_', '=', '+',
	'[', '{', ']', '}', '\\', '|',
	';', ':', '\'', '"',
	',', '<', '.', '>', '/', '?',
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

var USNormal = map[ebiten.Key]rune{
	ebiten.KeyA: 'a', ebiten.KeyB: 'b', ebiten.KeyC: 'c', ebiten.KeyD: 'd',
	ebiten.KeyE: 'e', ebiten.KeyF: 'f', ebiten.KeyG: 'g', ebiten.KeyH: 'h',
	ebiten.KeyI: 'i', ebiten.KeyJ: 'j', ebiten.KeyK: 'k', ebiten.KeyL: 'l',
	ebiten.KeyM: 'm', ebiten.KeyN: 'n', ebiten.KeyO: 'o', ebiten.KeyP: 'p',
	ebiten.KeyQ: 'q', ebiten.KeyR: 'r', ebiten.KeyS: 's', ebiten.KeyT: 't',
	ebiten.KeyU: 'u', ebiten.KeyV: 'v', ebiten.KeyW: 'w', ebiten.KeyX: 'x',
	ebiten.KeyY: 'y', ebiten.KeyZ: 'z',
	ebiten.KeyDigit1: '1', ebiten.KeyDigit2: '2', ebiten.KeyDigit3: '3',
	ebiten.KeyDigit4: '4', ebiten.KeyDigit5: '5', ebiten.KeyDigit6: '6',
	ebiten.KeyDigit7: '7', ebiten.KeyDigit8: '8', ebiten.KeyDigit9: '9',
	ebiten.KeyDigit0: '0',
	ebiten.KeyMinus:  '-', ebiten.KeyEqual: '=',
	ebiten.KeyBracketLeft: '[', ebiten.KeyBracketRight: ']',
	ebiten.KeyBackslash: '\\', ebiten.KeySemicolon: ';',
	ebiten.KeyQuote: '\'', ebiten.KeyBackquote: '`',
	ebiten.KeyComma: ',', ebiten.KeyPeriod: '.', ebiten.KeySlash: '/',
}

var USShifted = map[ebiten.Key]rune{
	ebiten.KeyA: 'A', ebiten.KeyB: 'B', ebiten.KeyC: 'C', ebiten.KeyD: 'D',
	ebiten.KeyE: 'E', ebiten.KeyF: 'F', ebiten.KeyG: 'G', ebiten.KeyH: 'H',
	ebiten.KeyI: 'I', ebiten.KeyJ: 'J', ebiten.KeyK: 'K', ebiten.KeyL: 'L',
	ebiten.KeyM: 'M', ebiten.KeyN: 'N', ebiten.KeyO: 'O', ebiten.KeyP: 'P',
	ebiten.KeyQ: 'Q', ebiten.KeyR: 'R', ebiten.KeyS: 'S', ebiten.KeyT: 'T',
	ebiten.KeyU: 'U', ebiten.KeyV: 'V', ebiten.KeyW: 'W', ebiten.KeyX: 'X',
	ebiten.KeyY: 'Y', ebiten.KeyZ: 'Z',
	ebiten.KeyDigit1: '!', ebiten.KeyDigit2: '@', ebiten.KeyDigit3: '#',
	ebiten.KeyDigit4: '$', ebiten.KeyDigit5: '%', ebiten.KeyDigit6: '^',
	ebiten.KeyDigit7: '&', ebiten.KeyDigit8: '*', ebiten.KeyDigit9: '(',
	ebiten.KeyDigit0: ')',
	ebiten.KeyMinus:  '_', ebiten.KeyEqual: '+',
	ebiten.KeyBracketLeft: '{', ebiten.KeyBracketRight: '}',
	ebiten.KeyBackslash: '|', ebiten.KeySemicolon: ':',
	ebiten.KeyQuote: '"', ebiten.KeyBackquote: '~',
	ebiten.KeyComma: '<', ebiten.KeyPeriod: '>', ebiten.KeySlash: '?',
}

var RUNormal = map[ebiten.Key]rune{
	ebiten.KeyA: 'ф', ebiten.KeyB: 'и', ebiten.KeyC: 'с', ebiten.KeyD: 'в',
	ebiten.KeyE: 'у', ebiten.KeyF: 'а', ebiten.KeyG: 'п', ebiten.KeyH: 'р',
	ebiten.KeyI: 'ш', ebiten.KeyJ: 'о', ebiten.KeyK: 'л', ebiten.KeyL: 'д',
	ebiten.KeyM: 'ь', ebiten.KeyN: 'т', ebiten.KeyO: 'щ', ebiten.KeyP: 'з',
	ebiten.KeyQ: 'й', ebiten.KeyR: 'к', ebiten.KeyS: 'ы', ebiten.KeyT: 'е',
	ebiten.KeyU: 'г', ebiten.KeyV: 'м', ebiten.KeyW: 'ц', ebiten.KeyX: 'ч',
	ebiten.KeyY: 'н', ebiten.KeyZ: 'я',
	ebiten.KeyDigit1: '1', ebiten.KeyDigit2: '2', ebiten.KeyDigit3: '3',
	ebiten.KeyDigit4: '4', ebiten.KeyDigit5: '5', ebiten.KeyDigit6: '6',
	ebiten.KeyDigit7: '7', ebiten.KeyDigit8: '8', ebiten.KeyDigit9: '9',
	ebiten.KeyDigit0: '0',
	ebiten.KeyMinus:  '-', ebiten.KeyEqual: '=',
	ebiten.KeyBracketLeft: 'х', ebiten.KeyBracketRight: 'ъ',
	ebiten.KeyBackslash: '\\', ebiten.KeySemicolon: 'ж',
	ebiten.KeyQuote: 'э', ebiten.KeyBackquote: 'ё',
	ebiten.KeyComma: 'б', ebiten.KeyPeriod: 'ю', ebiten.KeySlash: '.',
}

var RUShifted = map[ebiten.Key]rune{
	ebiten.KeyA: 'Ф', ebiten.KeyB: 'И', ebiten.KeyC: 'С', ebiten.KeyD: 'В',
	ebiten.KeyE: 'У', ebiten.KeyF: 'А', ebiten.KeyG: 'П', ebiten.KeyH: 'Р',
	ebiten.KeyI: 'Ш', ebiten.KeyJ: 'О', ebiten.KeyK: 'Л', ebiten.KeyL: 'Д',
	ebiten.KeyM: 'Ь', ebiten.KeyN: 'Т', ebiten.KeyO: 'Щ', ebiten.KeyP: 'З',
	ebiten.KeyQ: 'Й', ebiten.KeyR: 'К', ebiten.KeyS: 'Ы', ebiten.KeyT: 'Е',
	ebiten.KeyU: 'Г', ebiten.KeyV: 'М', ebiten.KeyW: 'Ц', ebiten.KeyX: 'Ч',
	ebiten.KeyY: 'Н', ebiten.KeyZ: 'Я',
	ebiten.KeyDigit1: '!', ebiten.KeyDigit2: '"', ebiten.KeyDigit3: '№',
	ebiten.KeyDigit4: ';', ebiten.KeyDigit5: '%', ebiten.KeyDigit6: ':',
	ebiten.KeyDigit7: '?', ebiten.KeyDigit8: '*', ebiten.KeyDigit9: '(',
	ebiten.KeyDigit0: ')',
	ebiten.KeyMinus:  '_', ebiten.KeyEqual: '+',
	ebiten.KeyBracketLeft: 'Х', ebiten.KeyBracketRight: 'Ъ',
	ebiten.KeyBackslash: '/', ebiten.KeySemicolon: 'Ж',
	ebiten.KeyQuote: 'Э', ebiten.KeyBackquote: 'Ё',
	ebiten.KeyComma: 'Б', ebiten.KeyPeriod: 'Ю', ebiten.KeySlash: ',',
}

var UANormal = map[ebiten.Key]rune{
	ebiten.KeyA: 'ф', ebiten.KeyB: 'и', ebiten.KeyC: 'с', ebiten.KeyD: 'в',
	ebiten.KeyE: 'у', ebiten.KeyF: 'а', ebiten.KeyG: 'п', ebiten.KeyH: 'р',
	ebiten.KeyI: 'ш', ebiten.KeyJ: 'о', ebiten.KeyK: 'л', ebiten.KeyL: 'д',
	ebiten.KeyM: 'ь', ebiten.KeyN: 'т', ebiten.KeyO: 'щ', ebiten.KeyP: 'з',
	ebiten.KeyQ: 'й', ebiten.KeyR: 'к', ebiten.KeyS: 'і', ebiten.KeyT: 'е',
	ebiten.KeyU: 'г', ebiten.KeyV: 'м', ebiten.KeyW: 'ц', ebiten.KeyX: 'ч',
	ebiten.KeyY: 'н', ebiten.KeyZ: 'я',
	ebiten.KeyDigit1: '1', ebiten.KeyDigit2: '2', ebiten.KeyDigit3: '3',
	ebiten.KeyDigit4: '4', ebiten.KeyDigit5: '5', ebiten.KeyDigit6: '6',
	ebiten.KeyDigit7: '7', ebiten.KeyDigit8: '8', ebiten.KeyDigit9: '9',
	ebiten.KeyDigit0: '0',
	ebiten.KeyMinus:  '-', ebiten.KeyEqual: '=',
	ebiten.KeyBracketLeft: 'х', ebiten.KeyBracketRight: 'ї',
	ebiten.KeyBackslash: 'ґ', ebiten.KeySemicolon: 'ж',
	ebiten.KeyQuote: 'є', ebiten.KeyBackquote: '\'',
	ebiten.KeyComma: 'б', ebiten.KeyPeriod: 'ю', ebiten.KeySlash: '.',
}

var UAShifted = map[ebiten.Key]rune{
	ebiten.KeyA: 'Ф', ebiten.KeyB: 'И', ebiten.KeyC: 'С', ebiten.KeyD: 'В',
	ebiten.KeyE: 'У', ebiten.KeyF: 'А', ebiten.KeyG: 'П', ebiten.KeyH: 'Р',
	ebiten.KeyI: 'Ш', ebiten.KeyJ: 'О', ebiten.KeyK: 'Л', ebiten.KeyL: 'Д',
	ebiten.KeyM: 'Ь', ebiten.KeyN: 'Т', ebiten.KeyO: 'Щ', ebiten.KeyP: 'З',
	ebiten.KeyQ: 'Й', ebiten.KeyR: 'К', ebiten.KeyS: 'І', ebiten.KeyT: 'Е',
	ebiten.KeyU: 'Г', ebiten.KeyV: 'М', ebiten.KeyW: 'Ц', ebiten.KeyX: 'Ч',
	ebiten.KeyY: 'Н', ebiten.KeyZ: 'Я',
	ebiten.KeyDigit1: '!', ebiten.KeyDigit2: '"', ebiten.KeyDigit3: '№',
	ebiten.KeyDigit4: ';', ebiten.KeyDigit5: '%', ebiten.KeyDigit6: ':',
	ebiten.KeyDigit7: '?', ebiten.KeyDigit8: '*', ebiten.KeyDigit9: '(',
	ebiten.KeyDigit0: ')',
	ebiten.KeyMinus:  '_', ebiten.KeyEqual: '+',
	ebiten.KeyBracketLeft: 'Х', ebiten.KeyBracketRight: 'Ї',
	ebiten.KeyBackslash: 'Ґ', ebiten.KeySemicolon: 'Ж',
	ebiten.KeyQuote: 'Є', ebiten.KeyBackquote: 'ʼ',
	ebiten.KeyComma: 'Б', ebiten.KeyPeriod: 'Ю', ebiten.KeySlash: ',',
}

var USKeyMap = map[rune]KeyWithShift{
	'a': {Key: ebiten.KeyA, Shift: false}, 'b': {Key: ebiten.KeyB, Shift: false}, 'c': {Key: ebiten.KeyC, Shift: false}, 'd': {Key: ebiten.KeyD, Shift: false},
	'e': {Key: ebiten.KeyE, Shift: false}, 'f': {Key: ebiten.KeyF, Shift: false}, 'g': {Key: ebiten.KeyG, Shift: false}, 'h': {Key: ebiten.KeyH, Shift: false},
	'i': {Key: ebiten.KeyI, Shift: false}, 'j': {Key: ebiten.KeyJ, Shift: false}, 'k': {Key: ebiten.KeyK, Shift: false}, 'l': {Key: ebiten.KeyL, Shift: false},
	'm': {Key: ebiten.KeyM, Shift: false}, 'n': {Key: ebiten.KeyN, Shift: false}, 'o': {Key: ebiten.KeyO, Shift: false}, 'p': {Key: ebiten.KeyP, Shift: false},
	'q': {Key: ebiten.KeyQ, Shift: false}, 'r': {Key: ebiten.KeyR, Shift: false}, 's': {Key: ebiten.KeyS, Shift: false}, 't': {Key: ebiten.KeyT, Shift: false},
	'u': {Key: ebiten.KeyU, Shift: false}, 'v': {Key: ebiten.KeyV, Shift: false}, 'w': {Key: ebiten.KeyW, Shift: false}, 'x': {Key: ebiten.KeyX, Shift: false},
	'y': {Key: ebiten.KeyY, Shift: false}, 'z': {Key: ebiten.KeyZ, Shift: false},
	'A': {Key: ebiten.KeyA, Shift: true}, 'B': {Key: ebiten.KeyB, Shift: true}, 'C': {Key: ebiten.KeyC, Shift: true}, 'D': {Key: ebiten.KeyD, Shift: true},
	'E': {Key: ebiten.KeyE, Shift: true}, 'F': {Key: ebiten.KeyF, Shift: true}, 'G': {Key: ebiten.KeyG, Shift: true}, 'H': {Key: ebiten.KeyH, Shift: true},
	'I': {Key: ebiten.KeyI, Shift: true}, 'J': {Key: ebiten.KeyJ, Shift: true}, 'K': {Key: ebiten.KeyK, Shift: true}, 'L': {Key: ebiten.KeyL, Shift: true},
	'M': {Key: ebiten.KeyM, Shift: true}, 'N': {Key: ebiten.KeyN, Shift: true}, 'O': {Key: ebiten.KeyO, Shift: true}, 'P': {Key: ebiten.KeyP, Shift: true},
	'Q': {Key: ebiten.KeyQ, Shift: true}, 'R': {Key: ebiten.KeyR, Shift: true}, 'S': {Key: ebiten.KeyS, Shift: true}, 'T': {Key: ebiten.KeyT, Shift: true},
	'U': {Key: ebiten.KeyU, Shift: true}, 'V': {Key: ebiten.KeyV, Shift: true}, 'W': {Key: ebiten.KeyW, Shift: true}, 'X': {Key: ebiten.KeyX, Shift: true},
	'Y': {Key: ebiten.KeyY, Shift: true}, 'Z': {Key: ebiten.KeyZ, Shift: true},
	'1': {Key: ebiten.KeyDigit1, Shift: false}, '2': {Key: ebiten.KeyDigit2, Shift: false}, '3': {Key: ebiten.KeyDigit3, Shift: false},
	'4': {Key: ebiten.KeyDigit4, Shift: false}, '5': {Key: ebiten.KeyDigit5, Shift: false}, '6': {Key: ebiten.KeyDigit6, Shift: false},
	'7': {Key: ebiten.KeyDigit7, Shift: false}, '8': {Key: ebiten.KeyDigit8, Shift: false}, '9': {Key: ebiten.KeyDigit9, Shift: false},
	'0': {Key: ebiten.KeyDigit0, Shift: false},
	'!': {Key: ebiten.KeyDigit1, Shift: true}, '@': {Key: ebiten.KeyDigit2, Shift: true}, '#': {Key: ebiten.KeyDigit3, Shift: true},
	'$': {Key: ebiten.KeyDigit4, Shift: true}, '%': {Key: ebiten.KeyDigit5, Shift: true}, '^': {Key: ebiten.KeyDigit6, Shift: true},
	'&': {Key: ebiten.KeyDigit7, Shift: true}, '*': {Key: ebiten.KeyDigit8, Shift: true}, '(': {Key: ebiten.KeyDigit9, Shift: true},
	')': {Key: ebiten.KeyDigit0, Shift: true},
	'-': {Key: ebiten.KeyMinus, Shift: false}, '=': {Key: ebiten.KeyEqual, Shift: false},
	'_': {Key: ebiten.KeyMinus, Shift: true}, '+': {Key: ebiten.KeyEqual, Shift: true},
	'[': {Key: ebiten.KeyBracketLeft, Shift: false}, ']': {Key: ebiten.KeyBracketRight, Shift: false},
	'{': {Key: ebiten.KeyBracketLeft, Shift: true}, '}': {Key: ebiten.KeyBracketRight, Shift: true},
	'\\': {Key: ebiten.KeyBackslash, Shift: false}, '|': {Key: ebiten.KeyBackslash, Shift: true},
	';': {Key: ebiten.KeySemicolon, Shift: false}, ':': {Key: ebiten.KeySemicolon, Shift: true},
	'\'': {Key: ebiten.KeyQuote, Shift: false}, '"': {Key: ebiten.KeyQuote, Shift: true},
	'`': {Key: ebiten.KeyBackquote, Shift: false}, '~': {Key: ebiten.KeyBackquote, Shift: true},
	',': {Key: ebiten.KeyComma, Shift: false}, '<': {Key: ebiten.KeyComma, Shift: true},
	'.': {Key: ebiten.KeyPeriod, Shift: false}, '>': {Key: ebiten.KeyPeriod, Shift: true},
	'/': {Key: ebiten.KeySlash, Shift: false}, '?': {Key: ebiten.KeySlash, Shift: true},
}

var RUKeyMap = map[rune]KeyWithShift{
	'ф': {Key: ebiten.KeyA, Shift: false}, 'и': {Key: ebiten.KeyB, Shift: false}, 'с': {Key: ebiten.KeyC, Shift: false}, 'в': {Key: ebiten.KeyD, Shift: false},
	'у': {Key: ebiten.KeyE, Shift: false}, 'а': {Key: ebiten.KeyF, Shift: false}, 'п': {Key: ebiten.KeyG, Shift: false}, 'р': {Key: ebiten.KeyH, Shift: false},
	'ш': {Key: ebiten.KeyI, Shift: false}, 'о': {Key: ebiten.KeyJ, Shift: false}, 'л': {Key: ebiten.KeyK, Shift: false}, 'д': {Key: ebiten.KeyL, Shift: false},
	'ь': {Key: ebiten.KeyM, Shift: false}, 'т': {Key: ebiten.KeyN, Shift: false}, 'щ': {Key: ebiten.KeyO, Shift: false}, 'з': {Key: ebiten.KeyP, Shift: false},
	'й': {Key: ebiten.KeyQ, Shift: false}, 'к': {Key: ebiten.KeyR, Shift: false}, 'ы': {Key: ebiten.KeyS, Shift: false}, 'е': {Key: ebiten.KeyT, Shift: false},
	'г': {Key: ebiten.KeyU, Shift: false}, 'м': {Key: ebiten.KeyV, Shift: false}, 'ц': {Key: ebiten.KeyW, Shift: false}, 'ч': {Key: ebiten.KeyX, Shift: false},
	'н': {Key: ebiten.KeyY, Shift: false}, 'я': {Key: ebiten.KeyZ, Shift: false},
	'Ф': {Key: ebiten.KeyA, Shift: true}, 'И': {Key: ebiten.KeyB, Shift: true}, 'С': {Key: ebiten.KeyC, Shift: true}, 'В': {Key: ebiten.KeyD, Shift: true},
	'У': {Key: ebiten.KeyE, Shift: true}, 'А': {Key: ebiten.KeyF, Shift: true}, 'П': {Key: ebiten.KeyG, Shift: true}, 'Р': {Key: ebiten.KeyH, Shift: true},
	'Ш': {Key: ebiten.KeyI, Shift: true}, 'О': {Key: ebiten.KeyJ, Shift: true}, 'Л': {Key: ebiten.KeyK, Shift: true}, 'Д': {Key: ebiten.KeyL, Shift: true},
	'Ь': {Key: ebiten.KeyM, Shift: true}, 'Т': {Key: ebiten.KeyN, Shift: true}, 'Щ': {Key: ebiten.KeyO, Shift: true}, 'З': {Key: ebiten.KeyP, Shift: true},
	'Й': {Key: ebiten.KeyQ, Shift: true}, 'К': {Key: ebiten.KeyR, Shift: true}, 'Ы': {Key: ebiten.KeyS, Shift: true}, 'Е': {Key: ebiten.KeyT, Shift: true},
	'Г': {Key: ebiten.KeyU, Shift: true}, 'М': {Key: ebiten.KeyV, Shift: true}, 'Ц': {Key: ebiten.KeyW, Shift: true}, 'Ч': {Key: ebiten.KeyX, Shift: true},
	'Н': {Key: ebiten.KeyY, Shift: true}, 'Я': {Key: ebiten.KeyZ, Shift: true},
	'1': {Key: ebiten.KeyDigit1, Shift: false}, '2': {Key: ebiten.KeyDigit2, Shift: false}, '3': {Key: ebiten.KeyDigit3, Shift: false},
	'4': {Key: ebiten.KeyDigit4, Shift: false}, '5': {Key: ebiten.KeyDigit5, Shift: false}, '6': {Key: ebiten.KeyDigit6, Shift: false},
	'7': {Key: ebiten.KeyDigit7, Shift: false}, '8': {Key: ebiten.KeyDigit8, Shift: false}, '9': {Key: ebiten.KeyDigit9, Shift: false},
	'0': {Key: ebiten.KeyDigit0, Shift: false},
	'!': {Key: ebiten.KeyDigit1, Shift: true}, '"': {Key: ebiten.KeyDigit2, Shift: true}, '№': {Key: ebiten.KeyDigit3, Shift: true},
	';': {Key: ebiten.KeyDigit4, Shift: true}, '%': {Key: ebiten.KeyDigit5, Shift: true}, ':': {Key: ebiten.KeyDigit6, Shift: true},
	'?': {Key: ebiten.KeyDigit7, Shift: true}, '*': {Key: ebiten.KeyDigit8, Shift: true}, '(': {Key: ebiten.KeyDigit9, Shift: true},
	')': {Key: ebiten.KeyDigit0, Shift: true},
	'-': {Key: ebiten.KeyMinus, Shift: false}, '=': {Key: ebiten.KeyEqual, Shift: false},
	'_': {Key: ebiten.KeyMinus, Shift: true}, '+': {Key: ebiten.KeyEqual, Shift: true},
	'х': {Key: ebiten.KeyBracketLeft, Shift: false}, 'ъ': {Key: ebiten.KeyBracketRight, Shift: false},
	'Х': {Key: ebiten.KeyBracketLeft, Shift: true}, 'Ъ': {Key: ebiten.KeyBracketRight, Shift: true},
	'\\': {Key: ebiten.KeyBackslash, Shift: false}, '/': {Key: ebiten.KeyBackslash, Shift: true},
	'ж': {Key: ebiten.KeySemicolon, Shift: false}, 'Ж': {Key: ebiten.KeySemicolon, Shift: true},
	'э': {Key: ebiten.KeyQuote, Shift: false}, 'Э': {Key: ebiten.KeyQuote, Shift: true},
	'ё': {Key: ebiten.KeyBackquote, Shift: false}, 'Ё': {Key: ebiten.KeyBackquote, Shift: true},
	'б': {Key: ebiten.KeyComma, Shift: false}, 'Б': {Key: ebiten.KeyComma, Shift: true},
	'ю': {Key: ebiten.KeyPeriod, Shift: false}, 'Ю': {Key: ebiten.KeyPeriod, Shift: true},
	'.': {Key: ebiten.KeySlash, Shift: false}, ',': {Key: ebiten.KeySlash, Shift: true},
}

var UAKeyMap = map[rune]KeyWithShift{
	'ф': {Key: ebiten.KeyA, Shift: false}, 'и': {Key: ebiten.KeyB, Shift: false}, 'с': {Key: ebiten.KeyC, Shift: false}, 'в': {Key: ebiten.KeyD, Shift: false},
	'у': {Key: ebiten.KeyE, Shift: false}, 'а': {Key: ebiten.KeyF, Shift: false}, 'п': {Key: ebiten.KeyG, Shift: false}, 'р': {Key: ebiten.KeyH, Shift: false},
	'ш': {Key: ebiten.KeyI, Shift: false}, 'о': {Key: ebiten.KeyJ, Shift: false}, 'л': {Key: ebiten.KeyK, Shift: false}, 'д': {Key: ebiten.KeyL, Shift: false},
	'ь': {Key: ebiten.KeyM, Shift: false}, 'т': {Key: ebiten.KeyN, Shift: false}, 'щ': {Key: ebiten.KeyO, Shift: false}, 'з': {Key: ebiten.KeyP, Shift: false},
	'й': {Key: ebiten.KeyQ, Shift: false}, 'к': {Key: ebiten.KeyR, Shift: false}, 'і': {Key: ebiten.KeyS, Shift: false}, 'е': {Key: ebiten.KeyT, Shift: false},
	'г': {Key: ebiten.KeyU, Shift: false}, 'м': {Key: ebiten.KeyV, Shift: false}, 'ц': {Key: ebiten.KeyW, Shift: false}, 'ч': {Key: ebiten.KeyX, Shift: false},
	'н': {Key: ebiten.KeyY, Shift: false}, 'я': {Key: ebiten.KeyZ, Shift: false},
	'Ф': {Key: ebiten.KeyA, Shift: true}, 'И': {Key: ebiten.KeyB, Shift: true}, 'С': {Key: ebiten.KeyC, Shift: true}, 'В': {Key: ebiten.KeyD, Shift: true},
	'У': {Key: ebiten.KeyE, Shift: true}, 'А': {Key: ebiten.KeyF, Shift: true}, 'П': {Key: ebiten.KeyG, Shift: true}, 'Р': {Key: ebiten.KeyH, Shift: true},
	'Ш': {Key: ebiten.KeyI, Shift: true}, 'О': {Key: ebiten.KeyJ, Shift: true}, 'Л': {Key: ebiten.KeyK, Shift: true}, 'Д': {Key: ebiten.KeyL, Shift: true},
	'Ь': {Key: ebiten.KeyM, Shift: true}, 'Т': {Key: ebiten.KeyN, Shift: true}, 'Щ': {Key: ebiten.KeyO, Shift: true}, 'З': {Key: ebiten.KeyP, Shift: true},
	'Й': {Key: ebiten.KeyQ, Shift: true}, 'К': {Key: ebiten.KeyR, Shift: true}, 'І': {Key: ebiten.KeyS, Shift: true}, 'Е': {Key: ebiten.KeyT, Shift: true},
	'Г': {Key: ebiten.KeyU, Shift: true}, 'М': {Key: ebiten.KeyV, Shift: true}, 'Ц': {Key: ebiten.KeyW, Shift: true}, 'Ч': {Key: ebiten.KeyX, Shift: true},
	'Н': {Key: ebiten.KeyY, Shift: true}, 'Я': {Key: ebiten.KeyZ, Shift: true},
	'1': {Key: ebiten.KeyDigit1, Shift: false}, '2': {Key: ebiten.KeyDigit2, Shift: false}, '3': {Key: ebiten.KeyDigit3, Shift: false},
	'4': {Key: ebiten.KeyDigit4, Shift: false}, '5': {Key: ebiten.KeyDigit5, Shift: false}, '6': {Key: ebiten.KeyDigit6, Shift: false},
	'7': {Key: ebiten.KeyDigit7, Shift: false}, '8': {Key: ebiten.KeyDigit8, Shift: false}, '9': {Key: ebiten.KeyDigit9, Shift: false},
	'0': {Key: ebiten.KeyDigit0, Shift: false},
	'!': {Key: ebiten.KeyDigit1, Shift: true}, '"': {Key: ebiten.KeyDigit2, Shift: true}, '№': {Key: ebiten.KeyDigit3, Shift: true},
	';': {Key: ebiten.KeyDigit4, Shift: true}, '%': {Key: ebiten.KeyDigit5, Shift: true}, ':': {Key: ebiten.KeyDigit6, Shift: true},
	'?': {Key: ebiten.KeyDigit7, Shift: true}, '*': {Key: ebiten.KeyDigit8, Shift: true}, '(': {Key: ebiten.KeyDigit9, Shift: true},
	')': {Key: ebiten.KeyDigit0, Shift: true},
	'-': {Key: ebiten.KeyMinus, Shift: false}, '=': {Key: ebiten.KeyEqual, Shift: false},
	'_': {Key: ebiten.KeyMinus, Shift: true}, '+': {Key: ebiten.KeyEqual, Shift: true},
	'х': {Key: ebiten.KeyBracketLeft, Shift: false}, 'ї': {Key: ebiten.KeyBracketRight, Shift: false},
	'Х': {Key: ebiten.KeyBracketLeft, Shift: true}, 'Ї': {Key: ebiten.KeyBracketRight, Shift: true},
	'ґ': {Key: ebiten.KeyBackslash, Shift: false}, 'Ґ': {Key: ebiten.KeyBackslash, Shift: true},
	'ж': {Key: ebiten.KeySemicolon, Shift: false}, 'Ж': {Key: ebiten.KeySemicolon, Shift: true},
	'є': {Key: ebiten.KeyQuote, Shift: false}, 'Є': {Key: ebiten.KeyQuote, Shift: true},
	'\'': {Key: ebiten.KeyBackquote, Shift: false}, 'ʼ': {Key: ebiten.KeyBackquote, Shift: true},
	'б': {Key: ebiten.KeyComma, Shift: false}, 'Б': {Key: ebiten.KeyComma, Shift: true},
	'ю': {Key: ebiten.KeyPeriod, Shift: false}, 'Ю': {Key: ebiten.KeyPeriod, Shift: true},
	'.': {Key: ebiten.KeySlash, Shift: false}, ',': {Key: ebiten.KeySlash, Shift: true},
}
