package main

import (
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
)

type ModCategory struct {
	Name    string
	Weight  int
	Handler func(*KeyCombo) error
}

var ModCategories = []ModCategory{
	{
		Name:   "No modifiers",
		Weight: 25,
		Handler: func(kc *KeyCombo) error {
			return nil
		},
	},
	{
		Name:   "Control",
		Weight: 2,
		Handler: func(kc *KeyCombo) error {
			kc.Control = true
			return nil
		},
	},
	{
		Name:   "Alt",
		Weight: 2,
		Handler: func(kc *KeyCombo) error {
			kc.Alt = true
			return nil
		},
	},
	{
		Name:   "Shift",
		Weight: 2,
		Handler: func(kc *KeyCombo) error {
			if kc.LowerRune == 0 && kc.UpperRune == 0 {
				kc.Shift = true
			}
			return nil
		},
	},
	{
		Name:   "Meta",
		Weight: 2,
		Handler: func(kc *KeyCombo) error {
			kc.Meta = true
			return nil
		},
	},
	{
		Name:   "Control+Alt",
		Weight: 1,
		Handler: func(kc *KeyCombo) error {
			kc.Control = true
			kc.Alt = true
			return nil
		},
	},
	{
		Name:   "Control+Shift",
		Weight: 1,
		Handler: func(kc *KeyCombo) error {
			kc.Control = true
			kc.Shift = true
			return nil
		},
	},
	{
		Name:   "Control+Meta",
		Weight: 0,
		Handler: func(kc *KeyCombo) error {
			kc.Control = true
			kc.Meta = true
			return nil
		},
	},
	{
		Name:   "Alt+Shift",
		Weight: 1,
		Handler: func(kc *KeyCombo) error {
			kc.Alt = true
			kc.Shift = true
			return nil
		},
	},
	{
		Name:   "Alt+Meta",
		Weight: 1,
		Handler: func(kc *KeyCombo) error {
			kc.Alt = true
			kc.Meta = true
			return nil
		},
	},
	{
		Name:   "Shift+Meta",
		Weight: 0,
		Handler: func(kc *KeyCombo) error {
			kc.Shift = true
			kc.Meta = true
			return nil
		},
	},
	{
		Name:   "Control+Alt+Shift",
		Weight: 1,
		Handler: func(kc *KeyCombo) error {
			kc.Control = true
			kc.Alt = true
			kc.Shift = true
			return nil
		},
	},
	{
		Name:   "Control+Alt+Meta",
		Weight: 0,
		Handler: func(kc *KeyCombo) error {
			kc.Control = true
			kc.Alt = true
			kc.Meta = true
			return nil
		},
	},
	{
		Name:   "Control+Shift+Meta",
		Weight: 0,
		Handler: func(kc *KeyCombo) error {
			kc.Control = true
			kc.Shift = true
			kc.Meta = true
			return nil
		},
	},
	{
		Name:   "Alt+Shift+Meta",
		Weight: 0,
		Handler: func(kc *KeyCombo) error {
			kc.Alt = true
			kc.Shift = true
			kc.Meta = true
			return nil
		},
	},
	{
		Name:   "Control+Alt+Shift+Meta",
		Weight: 0,
		Handler: func(kc *KeyCombo) error {
			kc.Control = true
			kc.Alt = true
			kc.Shift = true
			kc.Meta = true
			return nil
		},
	},
}

type BaseCategory struct {
	Name      string
	Weight    int
	Handler   func(*KeyCombo) error
	Validator func() bool
}

var BaseCategories = []BaseCategory{
	{
		Name:   "Lower Letters",
		Weight: 20,
		Handler: func(kc *KeyCombo) error {
			kc.setKeyComboBasedOnRune(CurrentLayout.LowerLetters[rand.IntN(len(CurrentLayout.LowerLetters))])
			return nil
		},
		Validator: func() bool {
			return len(CurrentLayout.LowerLetters) > 0
		},
	},
	{
		Name:   "Upper Letters",
		Weight: 10,
		Handler: func(kc *KeyCombo) error {
			kc.setKeyComboBasedOnRune(CurrentLayout.UpperLetters[rand.IntN(len(CurrentLayout.UpperLetters))])
			return nil
		},
		Validator: func() bool {
			return len(CurrentLayout.UpperLetters) > 0
		},
	},
	{
		Name:   "Lower Symbols",
		Weight: 10,
		Handler: func(kc *KeyCombo) error {
			kc.setKeyComboBasedOnRune(CurrentLayout.LowerSymbols[rand.IntN(len(CurrentLayout.LowerSymbols))])
			return nil
		},
		Validator: func() bool {
			return len(CurrentLayout.LowerSymbols) > 0
		},
	},
	{
		Name:   "Upper Symbols",
		Weight: 10,
		Handler: func(kc *KeyCombo) error {
			kc.setKeyComboBasedOnRune(CurrentLayout.UpperSymbols[rand.IntN(len(CurrentLayout.UpperSymbols))])
			return nil
		},
		Validator: func() bool {
			return len(CurrentLayout.UpperSymbols) > 0
		},
	},
	{
		Name:   "Digits",
		Weight: 10,
		Handler: func(kc *KeyCombo) error {
			kc.setKeyComboBasedOnRune(CurrentLayout.Digits[rand.IntN(len(CurrentLayout.Digits))])
			return nil
		},
		Validator: func() bool {
			return len(CurrentLayout.Digits) > 0
		},
	},
	{
		Name:   "Base Layer Non-Printable",
		Weight: 10,
		Handler: func(kc *KeyCombo) error {
			kc.Key = BaseLayerNonPrintableKeys[rand.IntN(len(BaseLayerNonPrintableKeys))]
			return nil
		},
		Validator: func() bool {
			return len(BaseLayerNonPrintableKeys) > 0
		},
	},
	{
		Name:   "Lower Layer Non-Printable",
		Weight: 5,
		Handler: func(kc *KeyCombo) error {
			kc.Key = LowerLayerNonPrintableKeys[rand.IntN(len(LowerLayerNonPrintableKeys))]
			return nil
		},
		Validator: func() bool {
			return len(LowerLayerNonPrintableKeys) > 0
		},
	},
	{
		Name:   "Custom Chars",
		Weight: 0,
		Handler: func(kc *KeyCombo) error {
			kc.setKeyComboBasedOnRune(CustomChars[rand.IntN(len(CustomChars))])
			return nil
		},
		Validator: func() bool {
			return len(CustomChars) > 0
		},
	},
	{
		Name:   "Custom Keys",
		Weight: 0,
		Handler: func(kc *KeyCombo) error {
			kc.Key = CustomKeys[rand.IntN(len(CustomKeys))]
			return nil
		},
		Validator: func() bool {
			return len(CustomKeys) > 0
		},
	},
}

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
