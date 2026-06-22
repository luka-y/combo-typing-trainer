package main

import "math/rand/v2"

func setBaseCategories() {
	BaseCategories = []BaseCategory{
		{
			Name:   "Lower Letters",
			Weight: BaseWeightLowerLetters,
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
			Weight: BaseWeightUpperLetters,
			Handler: func(kc *KeyCombo) error {
				kc.setKeyComboBasedOnRune(CurrentLayout.UpperLetters[rand.IntN(len(CurrentLayout.UpperLetters))])
				return nil
			},
			Validator: func() bool {
				return len(CurrentLayout.UpperLetters) > 0
			},
		},
		{
			Name:   "Digits",
			Weight: BaseWeightDigits,
			Handler: func(kc *KeyCombo) error {
				kc.setKeyComboBasedOnRune(CurrentLayout.Digits[rand.IntN(len(CurrentLayout.Digits))])
				return nil
			},
			Validator: func() bool {
				return len(CurrentLayout.Digits) > 0
			},
		},
		{
			Name:   "Upper Digit Symbols",
			Weight: BaseWeightUpperDigitSymbols,
			Handler: func(kc *KeyCombo) error {
				kc.setKeyComboBasedOnRune(CurrentLayout.UpperDigitSymbols[rand.IntN(len(CurrentLayout.UpperDigitSymbols))])
				return nil
			},
			Validator: func() bool {
				return len(CurrentLayout.UpperDigitSymbols) > 0
			},
		},
		{
			Name:   "Lower Symbols",
			Weight: BaseWeightLowerSymbols,
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
			Weight: BaseWeightUpperSymbols,
			Handler: func(kc *KeyCombo) error {
				kc.setKeyComboBasedOnRune(CurrentLayout.UpperSymbols[rand.IntN(len(CurrentLayout.UpperSymbols))])
				return nil
			},
			Validator: func() bool {
				return len(CurrentLayout.UpperSymbols) > 0
			},
		},
		{
			Name:   "Non-Printable Keys 1",
			Weight: BaseWeightNonPrintableKeys1,
			Handler: func(kc *KeyCombo) error {
				kc.Key = NonPrintableKeys1[rand.IntN(len(NonPrintableKeys1))]
				return nil
			},
			Validator: func() bool {
				return len(NonPrintableKeys1) > 0
			},
		},
		{
			Name:   "Non-Printable Keys 2",
			Weight: BaseWeightNonPrintableKeys2,
			Handler: func(kc *KeyCombo) error {
				kc.Key = NonPrintableKeys2[rand.IntN(len(NonPrintableKeys2))]
				return nil
			},
			Validator: func() bool {
				return len(NonPrintableKeys2) > 0
			},
		},
		{
			Name:   "Non-Printable Keys 3",
			Weight: BaseWeightNonPrintableKeys3,
			Handler: func(kc *KeyCombo) error {
				kc.Key = NonPrintableKeys3[rand.IntN(len(NonPrintableKeys3))]
				return nil
			},
			Validator: func() bool {
				return len(NonPrintableKeys3) > 0
			},
		},
		{
			Name:   "Non-Printable Keys 4",
			Weight: BaseWeightNonPrintableKeys4,
			Handler: func(kc *KeyCombo) error {
				kc.Key = NonPrintableKeys4[rand.IntN(len(NonPrintableKeys4))]
				return nil
			},
			Validator: func() bool {
				return len(NonPrintableKeys4) > 0
			},
		},
		{
			Name:   "Non-Printable Keys 5",
			Weight: BaseWeightNonPrintableKeys5,
			Handler: func(kc *KeyCombo) error {
				kc.Key = NonPrintableKeys5[rand.IntN(len(NonPrintableKeys5))]
				return nil
			},
			Validator: func() bool {
				return len(NonPrintableKeys5) > 0
			},
		},
		{
			Name:   "Custom Chars",
			Weight: BaseWeightCustomChars,
			Handler: func(kc *KeyCombo) error {
				kc.setKeyComboBasedOnRune(CustomChars[rand.IntN(len(CustomChars))])
				return nil
			},
			Validator: func() bool {
				return len(CustomChars) > 0
			},
		},
	}
}

func setModCategories() {
	ModCategories = []ModCategory{
		{
			Name:   "No modifiers",
			Weight: ModWeightNoModifiers,
			Handler: func(kc *KeyCombo) error {
				return nil
			},
		},
		{
			Name:   "Control",
			Weight: ModWeightControl,
			Handler: func(kc *KeyCombo) error {
				if kc.Shift == true {
					return nil
				}
				kc.Control = true
				return nil
			},
		},
		{
			Name:   "Alt",
			Weight: ModWeightAlt,
			Handler: func(kc *KeyCombo) error {
				if kc.Shift == true {
					return nil
				}
				kc.Alt = true
				return nil
			},
		},
		{
			Name:   "Shift",
			Weight: ModWeightShift,
			Handler: func(kc *KeyCombo) error {
				if kc.Shift == true {
					return nil
				}
				if kc.LowerRune == 0 && kc.UpperRune == 0 {
					kc.Shift = true
				}
				return nil
			},
		},
		{
			Name:   "Meta",
			Weight: ModWeightMeta,
			Handler: func(kc *KeyCombo) error {
				if kc.Shift == true {
					return nil
				}
				kc.Meta = true
				return nil
			},
		},
		{
			Name:   "Control+Alt",
			Weight: ModWeightControlAlt,
			Handler: func(kc *KeyCombo) error {
				if kc.Shift == true {
					return nil
				}
				kc.Control = true
				kc.Alt = true
				return nil
			},
		},
		{
			Name:   "Control+Shift",
			Weight: ModWeightControlShift,
			Handler: func(kc *KeyCombo) error {
				if kc.Shift == true {
					return nil
				}
				kc.Control = true
				kc.Shift = true
				return nil
			},
		},
		{
			Name:   "Control+Meta",
			Weight: ModWeightControlMeta,
			Handler: func(kc *KeyCombo) error {
				if kc.Shift == true {
					return nil
				}
				kc.Control = true
				kc.Meta = true
				return nil
			},
		},
		{
			Name:   "Alt+Shift",
			Weight: ModWeightAltShift,
			Handler: func(kc *KeyCombo) error {
				if kc.Shift == true {
					return nil
				}
				kc.Alt = true
				kc.Shift = true
				return nil
			},
		},
		{
			Name:   "Alt+Meta",
			Weight: ModWeightAltMeta,
			Handler: func(kc *KeyCombo) error {
				if kc.Shift == true {
					return nil
				}
				kc.Alt = true
				kc.Meta = true
				return nil
			},
		},
		{
			Name:   "Shift+Meta",
			Weight: ModWeightShiftMeta,
			Handler: func(kc *KeyCombo) error {
				if kc.Shift == true {
					return nil
				}
				kc.Shift = true
				kc.Meta = true
				return nil
			},
		},
		{
			Name:   "Control+Alt+Shift",
			Weight: ModWeightControlAltShift,
			Handler: func(kc *KeyCombo) error {
				if kc.Shift == true {
					return nil
				}
				kc.Control = true
				kc.Alt = true
				kc.Shift = true
				return nil
			},
		},
		{
			Name:   "Control+Alt+Meta",
			Weight: ModWeightControlAltMeta,
			Handler: func(kc *KeyCombo) error {
				if kc.Shift == true {
					return nil
				}
				kc.Control = true
				kc.Alt = true
				kc.Meta = true
				return nil
			},
		},
		{
			Name:   "Control+Shift+Meta",
			Weight: ModWeightControlShiftMeta,
			Handler: func(kc *KeyCombo) error {
				if kc.Shift == true {
					return nil
				}
				kc.Control = true
				kc.Shift = true
				kc.Meta = true
				return nil
			},
		},
		{
			Name:   "Alt+Shift+Meta",
			Weight: ModWeightAltShiftMeta,
			Handler: func(kc *KeyCombo) error {
				if kc.Shift == true {
					return nil
				}
				kc.Alt = true
				kc.Shift = true
				kc.Meta = true
				return nil
			},
		},
		{
			Name:   "Control+Alt+Shift+Meta",
			Weight: ModWeightControlAltShiftMeta,
			Handler: func(kc *KeyCombo) error {
				if kc.Shift == true {
					return nil
				}
				kc.Control = true
				kc.Alt = true
				kc.Shift = true
				kc.Meta = true
				return nil
			},
		},
	}
}
