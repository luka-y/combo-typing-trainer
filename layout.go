package main

import (
	"fmt"
	"unicode"

	"github.com/hajimehoshi/ebiten/v2"
)

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

const CurrentLayoutName = "US"

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

var InputUSKeyMap = map[KeyWithShift]rune{
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

var InputRUKeyMap = map[KeyWithShift]rune{
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

var InputUAKeyMap = map[KeyWithShift]rune{
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
