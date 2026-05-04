package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"image/color"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/BurntSushi/toml"
	"github.com/hajimehoshi/ebiten/v2"
)

type rawConfig struct {
	FontSize          int    `toml:"font_size"`
	CurrentLayoutName string `toml:"current_layout_name"`

	BaseWeights struct {
		LowerLetters      int `toml:"lower_letters"`
		UpperLetters      int `toml:"upper_letters"`
		LowerSymbols      int `toml:"lower_symbols"`
		UpperSymbols      int `toml:"upper_symbols"`
		Digits            int `toml:"digits"`
		NonPrintableKeys1 int `toml:"non_printable_keys_1"`
		NonPrintableKeys2 int `toml:"non_printable_keys_2"`
		CustomChars       int `toml:"custom_chars"`
		CustomKeys        int `toml:"custom_keys"`
	} `toml:"base_weights"`

	ModWeights struct {
		NoModifiers         int `toml:"no_modifiers"`
		Control             int `toml:"control"`
		Alt                 int `toml:"alt"`
		Shift               int `toml:"shift"`
		Meta                int `toml:"meta"`
		ControlAlt          int `toml:"control_alt"`
		ControlShift        int `toml:"control_shift"`
		ControlMeta         int `toml:"control_meta"`
		AltShift            int `toml:"alt_shift"`
		AltMeta             int `toml:"alt_meta"`
		ShiftMeta           int `toml:"shift_meta"`
		ControlAltShift     int `toml:"control_alt_shift"`
		ControlAltMeta      int `toml:"control_alt_meta"`
		ControlShiftMeta    int `toml:"control_shift_meta"`
		AltShiftMeta        int `toml:"alt_shift_meta"`
		ControlAltShiftMeta int `toml:"control_alt_shift_meta"`
	} `toml:"mod_weights"`

	NonInferredBaseCategories struct {
		NonPrintableKeys1 []string `toml:"non_printable_keys_1"`
		NonPrintableKeys2 []string `toml:"non_printable_keys_2"`
		CustomChars       []string `toml:"custom_chars"`
		CustomKeys        []string `toml:"custom_keys"`
	} `toml:"non_inferred_base_categories"`

	Colors struct {
		Background         string `toml:"background"`
		UpcomingCombo      string `toml:"upcoming_combo"`
		CurrentCombo       string `toml:"current_combo"`
		CorrectPastCombo   string `toml:"correct_past_combo"`
		IncorrectPastCombo string `toml:"incorrect_past_combo"`
	} `toml:"colors"`

	Layouts []struct {
		Name   string               `toml:"name"`
		KeyMap map[string][2]string `toml:"key_map"`
	} `toml:"layouts"`
}

func ParseConfig() error {
	_, err := os.Stat("config.toml")
	if errors.Is(err, os.ErrNotExist) {
		defaultConfigBytes, err := assets.ReadFile("assets/default-config.toml") //os.ReadFile("assets/")
		if err != nil {
			return fmt.Errorf("error reading default config: %w", err)
		}
		if err := os.WriteFile("config.toml", defaultConfigBytes, 0644); err != nil {
			return fmt.Errorf("error writing config.toml: %w", err)
		}

	} else if err != nil {
		return fmt.Errorf("error with os.Stat-ing config.toml: %w", err)
	}

	var cfg rawConfig
	metadata, err := toml.DecodeFile("config.toml", &cfg)
	if err != nil {
		return fmt.Errorf("error decoding toml file: %w", err)
	}
	for _, key := range metadata.Undecoded() {
		return fmt.Errorf("redundant key: %q", key.String())
	}

	if !metadata.IsDefined("font_size") || cfg.FontSize <= 0 {
		return fmt.Errorf("missed or non-positive int entry: font_size")
	}
	FontSize = cfg.FontSize

	if !metadata.IsDefined("current_layout_name") {
		return fmt.Errorf("missed string entry: current_layout_name")
	}
	CurrentLayoutName = cfg.CurrentLayoutName

	if !metadata.IsDefined("base_weights", "lower_letters") || cfg.BaseWeights.LowerLetters < 0 {
		return fmt.Errorf("missed or negative int entry: base_weights.lower_letters")
	}
	BaseWeightLowerLetters = cfg.BaseWeights.LowerLetters

	if !metadata.IsDefined("base_weights", "upper_letters") || cfg.BaseWeights.UpperLetters < 0 {
		return fmt.Errorf("missed or negative int entry: base_weights.upper_letters")
	}
	BaseWeightUpperLetters = cfg.BaseWeights.UpperLetters

	if !metadata.IsDefined("base_weights", "lower_symbols") || cfg.BaseWeights.LowerSymbols < 0 {
		return fmt.Errorf("missed or negative int entry: base_weights.lower_symbols")
	}
	BaseWeightLowerSymbols = cfg.BaseWeights.LowerSymbols

	if !metadata.IsDefined("base_weights", "upper_symbols") || cfg.BaseWeights.UpperSymbols < 0 {
		return fmt.Errorf("missed or negative int entry: base_weights.upper_symbols")
	}
	BaseWeightUpperSymbols = cfg.BaseWeights.UpperSymbols

	if !metadata.IsDefined("base_weights", "digits") || cfg.BaseWeights.Digits < 0 {
		return fmt.Errorf("missed or negative int entry: base_weights.digits")
	}
	BaseWeightDigits = cfg.BaseWeights.Digits

	if !metadata.IsDefined("base_weights", "non_printable_keys_1") || cfg.BaseWeights.NonPrintableKeys1 < 0 {
		return fmt.Errorf("missed or negative int entry: base_weights.non_printable_keys_1")
	}
	BaseWeightNonPrintableKeys1 = cfg.BaseWeights.NonPrintableKeys1

	if !metadata.IsDefined("base_weights", "non_printable_keys_2") || cfg.BaseWeights.NonPrintableKeys2 < 0 {
		return fmt.Errorf("missed or negative int entry: base_weights.non_printable_keys_2")
	}
	BaseWeightNonPrintableKeys2 = cfg.BaseWeights.NonPrintableKeys2

	if !metadata.IsDefined("base_weights", "custom_chars") || cfg.BaseWeights.CustomChars < 0 {
		return fmt.Errorf("missed or negative int entry: base_weights.custom_chars")
	}
	BaseWeightCustomChars = cfg.BaseWeights.CustomChars

	if !metadata.IsDefined("base_weights", "custom_keys") || cfg.BaseWeights.CustomKeys < 0 {
		return fmt.Errorf("missed or negative int entry: base_weights.custom_keys")
	}
	BaseWeightCustomKeys = cfg.BaseWeights.CustomKeys

	if !metadata.IsDefined("mod_weights", "no_modifiers") || cfg.ModWeights.NoModifiers < 0 {
		return fmt.Errorf("missed or negative int entry: mod_weights.no_modifiers")
	}
	ModWeightNoModifiers = cfg.ModWeights.NoModifiers

	if !metadata.IsDefined("mod_weights", "control") || cfg.ModWeights.Control < 0 {
		return fmt.Errorf("missed or negative int entry: mod_weights.control")
	}
	ModWeightControl = cfg.ModWeights.Control

	if !metadata.IsDefined("mod_weights", "alt") || cfg.ModWeights.Alt < 0 {
		return fmt.Errorf("missed or negative int entry: mod_weights.alt")
	}
	ModWeightAlt = cfg.ModWeights.Alt

	if !metadata.IsDefined("mod_weights", "shift") || cfg.ModWeights.Shift < 0 {
		return fmt.Errorf("missed or negative int entry: mod_weights.shift")
	}
	ModWeightShift = cfg.ModWeights.Shift

	if !metadata.IsDefined("mod_weights", "meta") || cfg.ModWeights.Meta < 0 {
		return fmt.Errorf("missed or negative int entry: mod_weights.meta")
	}
	ModWeightMeta = cfg.ModWeights.Meta

	if !metadata.IsDefined("mod_weights", "control_alt") || cfg.ModWeights.ControlAlt < 0 {
		return fmt.Errorf("missed or negative int entry: mod_weights.control_alt")
	}
	ModWeightControlAlt = cfg.ModWeights.ControlAlt

	if !metadata.IsDefined("mod_weights", "control_shift") || cfg.ModWeights.ControlShift < 0 {
		return fmt.Errorf("missed or negative int entry: mod_weights.control_shift")
	}
	ModWeightControlShift = cfg.ModWeights.ControlShift

	if !metadata.IsDefined("mod_weights", "control_meta") || cfg.ModWeights.ControlMeta < 0 {
		return fmt.Errorf("missed or negative int entry: mod_weights.control_meta")
	}
	ModWeightControlMeta = cfg.ModWeights.ControlMeta

	if !metadata.IsDefined("mod_weights", "alt_shift") || cfg.ModWeights.AltShift < 0 {
		return fmt.Errorf("missed or negative int entry: mod_weights.alt_shift")
	}
	ModWeightAltShift = cfg.ModWeights.AltShift

	if !metadata.IsDefined("mod_weights", "alt_meta") || cfg.ModWeights.AltMeta < 0 {
		return fmt.Errorf("missed or negative int entry: mod_weights.alt_meta")
	}
	ModWeightAltMeta = cfg.ModWeights.AltMeta

	if !metadata.IsDefined("mod_weights", "shift_meta") || cfg.ModWeights.ShiftMeta < 0 {
		return fmt.Errorf("missed or negative int entry: mod_weights.shift_meta")
	}
	ModWeightShiftMeta = cfg.ModWeights.ShiftMeta

	if !metadata.IsDefined("mod_weights", "control_alt_shift") || cfg.ModWeights.ControlAltShift < 0 {
		return fmt.Errorf("missed or negative int entry: mod_weights.control_alt_shift")
	}
	ModWeightControlAltShift = cfg.ModWeights.ControlAltShift

	if !metadata.IsDefined("mod_weights", "control_alt_meta") || cfg.ModWeights.ControlAltMeta < 0 {
		return fmt.Errorf("missed or negative int entry: mod_weights.control_alt_meta")
	}
	ModWeightControlAltMeta = cfg.ModWeights.ControlAltMeta

	if !metadata.IsDefined("mod_weights", "control_shift_meta") || cfg.ModWeights.ControlShiftMeta < 0 {
		return fmt.Errorf("missed or negative int entry: mod_weights.control_shift_meta")
	}
	ModWeightControlShiftMeta = cfg.ModWeights.ControlShiftMeta

	if !metadata.IsDefined("mod_weights", "alt_shift_meta") || cfg.ModWeights.AltShiftMeta < 0 {
		return fmt.Errorf("missed or negative int entry: mod_weights.alt_shift_meta")
	}
	ModWeightAltShiftMeta = cfg.ModWeights.AltShiftMeta

	if !metadata.IsDefined("mod_weights", "control_alt_shift_meta") || cfg.ModWeights.ControlAltShiftMeta < 0 {
		return fmt.Errorf("missed or negative int entry: mod_weights.control_alt_shift_meta")
	}
	ModWeightControlAltShiftMeta = cfg.ModWeights.ControlAltShiftMeta

	baseWeightsSum := BaseWeightLowerLetters + BaseWeightUpperLetters + BaseWeightLowerSymbols + BaseWeightUpperSymbols +
		BaseWeightDigits + BaseWeightNonPrintableKeys1 + BaseWeightNonPrintableKeys2 + BaseWeightCustomChars + BaseWeightCustomKeys
	if baseWeightsSum == 0 {
		return fmt.Errorf("sum of base_weights entries equals zero")
	}

	modWeightsSum := ModWeightNoModifiers + ModWeightControl + ModWeightAlt + ModWeightShift + ModWeightMeta +
		ModWeightControlAlt + ModWeightControlShift + ModWeightControlMeta + ModWeightAltShift + ModWeightAltMeta +
		ModWeightShiftMeta + ModWeightControlAltShift + ModWeightControlAltMeta + ModWeightControlShiftMeta +
		ModWeightAltShiftMeta + ModWeightControlAltShiftMeta
	if modWeightsSum == 0 {
		return fmt.Errorf("sum of mod_weights entries equals zero")
	}

	setModCategories()
	setBaseCategories()

	if !metadata.IsDefined("non_inferred_base_categories", "non_printable_keys_1") {
		return fmt.Errorf("missed []string entry: non_inferred_base_categories.non_printable_keys_1")
	}
	for _, rawKey := range cfg.NonInferredBaseCategories.NonPrintableKeys1 {
		key, exist := stringToEbitenKeyMap[rawKey]
		if !exist {
			return fmt.Errorf("non_printable_keys_1 includes a key that does not exist: %s", rawKey)
		}
		NonPrintableKeys1 = append(NonPrintableKeys1, key)
	}

	if !metadata.IsDefined("non_inferred_base_categories", "non_printable_keys_2") {
		return fmt.Errorf("missed []string entry: non_inferred_base_categories.non_printable_keys_2")
	}
	for _, rawKey := range cfg.NonInferredBaseCategories.NonPrintableKeys2 {
		key, exist := stringToEbitenKeyMap[rawKey]
		if !exist {
			return fmt.Errorf("non_printable_keys_2 includes a key that does not exist: %s", rawKey)
		}
		NonPrintableKeys2 = append(NonPrintableKeys2, key)
	}

	if !metadata.IsDefined("non_inferred_base_categories", "custom_keys") {
		return fmt.Errorf("missed []string entry: non_inferred_base_categories.custom_keys")
	}
	for _, rawKey := range cfg.NonInferredBaseCategories.CustomKeys {
		key, exist := stringToEbitenKeyMap[rawKey]
		if !exist {
			return fmt.Errorf("custom_keys includes a key that does not exist: %s", rawKey)
		}
		CustomKeys = append(CustomKeys, key)
	}

	if !metadata.IsDefined("non_inferred_base_categories", "custom_chars") {
		return fmt.Errorf("missed []string entry: non_inferred_base_categories.custom_chars")
	}
	for _, charStr := range cfg.NonInferredBaseCategories.CustomChars {
		r, isValid := convertStringToRune(charStr)
		if !isValid {
			return fmt.Errorf("custom_chars includes a string that is not a valid rune: %s", charStr)
		}
		CustomChars = append(CustomChars, r)
	}

	if !metadata.IsDefined("colors", "background") {
		return fmt.Errorf("missed string entry: colors.background")
	}
	BackgroundColor, err = getRGBAFromHex(cfg.Colors.Background)
	if err != nil {
		return err
	}

	if !metadata.IsDefined("colors", "upcoming_combo") {
		return fmt.Errorf("missed string entry: colors.upcoming_combo")
	}
	UpcomingComboColor, err = getRGBAFromHex(cfg.Colors.UpcomingCombo)
	if err != nil {
		return err
	}

	if !metadata.IsDefined("colors", "current_combo") {
		return fmt.Errorf("missed string entry: colors.current_combo")
	}
	CurrentComboColor, err = getRGBAFromHex(cfg.Colors.CurrentCombo)
	if err != nil {
		return err
	}

	if !metadata.IsDefined("colors", "correct_past_combo") {
		return fmt.Errorf("missed string entry: colors.correct_past_combo")
	}
	CorrectPastComboColor, err = getRGBAFromHex(cfg.Colors.CorrectPastCombo)
	if err != nil {
		return err
	}

	if !metadata.IsDefined("colors", "incorrect_past_combo") {
		return fmt.Errorf("missed string entry: colors.incorrect_past_combo")
	}
	IncorrectPastComboColor, err = getRGBAFromHex(cfg.Colors.IncorrectPastCombo)
	if err != nil {
		return err
	}

	seenLayoutNames := make(map[string]bool)
	for i, rawLayout := range cfg.Layouts {
		if rawLayout.Name == "" {
			return fmt.Errorf("missed string entry: layouts[%d].name", i)
		}
		if rawLayout.KeyMap == nil {
			return fmt.Errorf("missed map entry: layouts[%d].key_map", i)
		}
		if seenLayoutNames[rawLayout.Name] {
			return fmt.Errorf("duplicate layout name: %q", rawLayout.Name)
		}
		seenLayoutNames[rawLayout.Name] = true
		final := InputLayout{Name: rawLayout.Name}
		final.KeyMap = make(map[KeyWithShift]rune)
		for rawKeyName, doubleRune := range rawLayout.KeyMap {
			key, exist := stringToEbitenKeyMap[rawKeyName]
			if !exist {
				return fmt.Errorf("layout \"%s\" has a key that does not exist: %s", rawLayout.Name, rawKeyName)
			}
			lowerRune, isValid := convertStringToRune(doubleRune[0])
			if !isValid {
				return fmt.Errorf("layout \"%s\" has a key \"%s\" first value of which is not a valid rune: %s", rawLayout.Name, key.String(), doubleRune[0])
			}
			upperRune, isValid := convertStringToRune(doubleRune[1])
			if !isValid {
				return fmt.Errorf("layout \"%s\" has a key \"%s\" second value of which is not a valid rune: %s", rawLayout.Name, key.String(), doubleRune[1])
			}
			final.KeyMap[KeyWithShift{key, false}] = lowerRune
			final.KeyMap[KeyWithShift{key, true}] = upperRune
		}
		InputLayouts = append(InputLayouts, final)
	}

	for _, il := range InputLayouts {
		layout, err := getLayoutFromKeyMap(il.Name, il.KeyMap)
		if err != nil {
			return fmt.Errorf("err getting layout from the %s key map: %w", il.Name, err)
		}
		Layouts = append(Layouts, layout)
	}

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
			return fmt.Errorf("custom_chars contains a char '%c' that is absent in the current layout", r)
		}
	}

	for _, key := range CustomKeys {
		_, exist := CurrentLayout.KeyMap[KeyWithShift{Key: key, Shift: false}]
		if exist {
			return fmt.Errorf("custom_keys contain a key %q that is present in the current layout. Use custom_chars instead", key.String())
		}
	}

	for _, key := range NonPrintableKeys1 {
		_, exist := CurrentLayout.KeyMap[KeyWithShift{Key: key, Shift: false}]
		if exist {
			return fmt.Errorf("non_printable_keys_1 contain a key %q that is present in the current layout. Use custom_chars instead", key.String())
		}
	}

	for _, key := range NonPrintableKeys2 {
		_, exist := CurrentLayout.KeyMap[KeyWithShift{Key: key, Shift: false}]
		if exist {
			return fmt.Errorf("non_printable_keys_2 contain a key %q that is present in the current layout. Use custom_chars instead", key.String())
		}
	}

	for _, category := range BaseCategories {
		if category.Weight > 0 && !category.Validator() {
			return fmt.Errorf("base category \"%s\" weight is more than zero while slice it is pulling from is empty", category.Name)
		}
	}
	return nil
}

func getRGBAFromHex(hexString string) (color.RGBA, error) {
	original := hexString
	hexString = strings.TrimSpace(hexString)
	hexString = strings.TrimPrefix(hexString, "#")

	if len(hexString) != 6 {
		return color.RGBA{}, fmt.Errorf("invalid hex color string: %q (expected 6 hex digits)", original)
	}

	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		return color.RGBA{}, fmt.Errorf("invalid hex color string: %q: %w", original, err)
	}

	return color.RGBA{bytes[0], bytes[1], bytes[2], 255}, nil
}

func convertStringToRune(s string) (r rune, isValid bool) {
	if len([]rune(s)) != 1 {
		return 0, false
	}
	r, _ = utf8.DecodeRuneInString(s)
	if r == utf8.RuneError {
		return 0, false
	}
	return r, true
}

func getLayoutFromKeyMap(name string, keyMap map[KeyWithShift]rune) (Layout, error) {
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

var stringToEbitenKeyMap = map[string]ebiten.Key{
	"A": ebiten.KeyA, "B": ebiten.KeyB, "C": ebiten.KeyC, "D": ebiten.KeyD, "E": ebiten.KeyE, "F": ebiten.KeyF,
	"G": ebiten.KeyG, "H": ebiten.KeyH, "I": ebiten.KeyI, "J": ebiten.KeyJ, "K": ebiten.KeyK, "L": ebiten.KeyL,
	"M": ebiten.KeyM, "N": ebiten.KeyN, "O": ebiten.KeyO, "P": ebiten.KeyP, "Q": ebiten.KeyQ, "R": ebiten.KeyR,
	"S": ebiten.KeyS, "T": ebiten.KeyT, "U": ebiten.KeyU, "V": ebiten.KeyV, "W": ebiten.KeyW, "X": ebiten.KeyX,
	"Y": ebiten.KeyY, "Z": ebiten.KeyZ,
	"ArrowDown": ebiten.KeyArrowDown, "ArrowLeft": ebiten.KeyArrowLeft, "ArrowRight": ebiten.KeyArrowRight, "ArrowUp": ebiten.KeyArrowUp,
	"Backquote": ebiten.KeyBackquote, "Backslash": ebiten.KeyBackslash,
	"BracketLeft": ebiten.KeyBracketLeft, "BracketRight": ebiten.KeyBracketRight, "CapsLock": ebiten.KeyCapsLock,
	"Comma": ebiten.KeyComma, "ContextMenu": ebiten.KeyContextMenu,
	"Delete": ebiten.KeyDelete,
	"Digit0": ebiten.KeyDigit0, "Digit1": ebiten.KeyDigit1, "Digit2": ebiten.KeyDigit2, "Digit3": ebiten.KeyDigit3,
	"Digit4": ebiten.KeyDigit4, "Digit5": ebiten.KeyDigit5, "Digit6": ebiten.KeyDigit6, "Digit7": ebiten.KeyDigit7,
	"Digit8": ebiten.KeyDigit8, "Digit9": ebiten.KeyDigit9, "End": ebiten.KeyEnd, "Enter": ebiten.KeyEnter,
	"Equal": ebiten.KeyEqual, "Escape": ebiten.KeyEscape, "F1": ebiten.KeyF1, "F2": ebiten.KeyF2, "F3": ebiten.KeyF3,
	"F4": ebiten.KeyF4, "F5": ebiten.KeyF5, "F6": ebiten.KeyF6, "F7": ebiten.KeyF7, "F8": ebiten.KeyF8,
	"F9": ebiten.KeyF9, "F10": ebiten.KeyF10, "F11": ebiten.KeyF11, "F12": ebiten.KeyF12, "F13": ebiten.KeyF13,
	"F14": ebiten.KeyF14, "F15": ebiten.KeyF15, "F16": ebiten.KeyF16, "F17": ebiten.KeyF17, "F18": ebiten.KeyF18,
	"F19": ebiten.KeyF19, "F20": ebiten.KeyF20, "F21": ebiten.KeyF21, "F22": ebiten.KeyF22, "F23": ebiten.KeyF23,
	"F24": ebiten.KeyF24, "Home": ebiten.KeyHome, "Insert": ebiten.KeyInsert, "IntlBackslash": ebiten.KeyIntlBackslash,
	"Minus":   ebiten.KeyMinus,
	"NumLock": ebiten.KeyNumLock, "Numpad0": ebiten.KeyNumpad0, "Numpad1": ebiten.KeyNumpad1, "Numpad2": ebiten.KeyNumpad2,
	"Numpad3": ebiten.KeyNumpad3, "Numpad4": ebiten.KeyNumpad4, "Numpad5": ebiten.KeyNumpad5, "Numpad6": ebiten.KeyNumpad6,
	"Numpad7": ebiten.KeyNumpad7, "Numpad8": ebiten.KeyNumpad8, "Numpad9": ebiten.KeyNumpad9, "NumpadAdd": ebiten.KeyNumpadAdd,
	"NumpadDecimal": ebiten.KeyNumpadDecimal, "NumpadDivide": ebiten.KeyNumpadDivide, "NumpadEnter": ebiten.KeyNumpadEnter,
	"NumpadEqual": ebiten.KeyNumpadEqual, "NumpadMultiply": ebiten.KeyNumpadMultiply, "NumpadSubtract": ebiten.KeyNumpadSubtract,
	"PageDown": ebiten.KeyPageDown, "PageUp": ebiten.KeyPageUp, "Pause": ebiten.KeyPause, "Period": ebiten.KeyPeriod,
	"PrintScreen": ebiten.KeyPrintScreen, "Quote": ebiten.KeyQuote, "ScrollLock": ebiten.KeyScrollLock, "Semicolon": ebiten.KeySemicolon,
	"Slash": ebiten.KeySlash,
	"Space": ebiten.KeySpace, "Tab": ebiten.KeyTab,
}
