package main

import "github.com/hajimehoshi/ebiten/v2"

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

func (kc *KeyCombo) setKeyComboBasedOnRune(r rune) {
	keyWithShift := CurrentLayout.ReverseKeyMap[r]
	kc.Key = keyWithShift.Key
	kc.Shift = keyWithShift.Shift
	kc.LowerRune = CurrentLayout.KeyMap[KeyWithShift{kc.Key, false}]
	kc.UpperRune = CurrentLayout.KeyMap[KeyWithShift{kc.Key, true}]
}

type BaseCategory struct {
	Name      string
	Weight    int
	Handler   func(*KeyCombo) error
	Validator func() bool
}

type ModCategory struct {
	Name    string
	Weight  int
	Handler func(*KeyCombo) error
}

type KeyWithShift struct {
	Key   ebiten.Key
	Shift bool
}

type InputLayout struct {
	Name   string
	KeyMap map[KeyWithShift]rune
}

type Layout struct {
	Name   string
	KeyMap map[KeyWithShift]rune

	ReverseKeyMap map[rune]KeyWithShift

	LowerLetters      []rune
	UpperLetters      []rune
	Digits            []rune
	UpperDigitSymbols []rune
	LowerSymbols      []rune
	UpperSymbols      []rune
}
