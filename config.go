package main

import (
	"fmt"
	"image/color"
	"strings"
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
	} `toml:"non-inferred-base_categories"`

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
	var cfg rawConfig
	_, err := toml.DecodeFile("assets/default-config.toml", &cfg)
	if err != nil {
		return fmt.Errorf("error decoding toml file: %w", err)
	}

	FontSize = cfg.FontSize
	CurrentLayoutName = cfg.CurrentLayoutName

	BaseWeightLowerLetters = cfg.BaseWeights.LowerLetters
	BaseWeightUpperLetters = cfg.BaseWeights.UpperLetters
	BaseWeightLowerSymbols = cfg.BaseWeights.LowerSymbols
	BaseWeightUpperSymbols = cfg.BaseWeights.UpperSymbols
	BaseWeightDigits = cfg.BaseWeights.Digits
	BaseWeightNonPrintableKeys1 = cfg.BaseWeights.NonPrintableKeys1
	BaseWeightNonPrintableKeys2 = cfg.BaseWeights.NonPrintableKeys2
	BaseWeightCustomChars = cfg.BaseWeights.CustomChars
	BaseWeightCustomKeys = cfg.BaseWeights.CustomKeys

	ModWeightNoModifiers = cfg.ModWeights.NoModifiers
	ModWeightControl = cfg.ModWeights.Control
	ModWeightAlt = cfg.ModWeights.Alt
	ModWeightShift = cfg.ModWeights.Shift
	ModWeightMeta = cfg.ModWeights.Meta
	ModWeightControlAlt = cfg.ModWeights.ControlAlt
	ModWeightControlShift = cfg.ModWeights.ControlShift
	ModWeightControlMeta = cfg.ModWeights.ControlMeta
	ModWeightAltShift = cfg.ModWeights.AltShift
	ModWeightAltMeta = cfg.ModWeights.AltMeta
	ModWeightShiftMeta = cfg.ModWeights.ShiftMeta
	ModWeightControlAltShift = cfg.ModWeights.ControlAltShift
	ModWeightControlAltMeta = cfg.ModWeights.ControlAltMeta
	ModWeightControlShiftMeta = cfg.ModWeights.ControlShiftMeta
	ModWeightAltShiftMeta = cfg.ModWeights.AltShiftMeta
	ModWeightControlAltShiftMeta = cfg.ModWeights.ControlAltShiftMeta

	setModCategories()
	setBaseCategories()

	for _, rawKey := range cfg.NonInferredBaseCategories.NonPrintableKeys1 {
		key, exist := stringToEbitenKeyMap[rawKey]
		if !exist {
			return fmt.Errorf("non_printable_keys_1 includes a key that does not exist: %s", rawKey)
		}
		NonPrintableKeys1 = append(NonPrintableKeys1, key)
	}
	for _, rawKey := range cfg.NonInferredBaseCategories.NonPrintableKeys2 {
		key, exist := stringToEbitenKeyMap[rawKey]
		if !exist {
			return fmt.Errorf("non_printable_keys_2 includes a key that does not exist: %s", rawKey)
		}
		NonPrintableKeys2 = append(NonPrintableKeys2, key)
	}
	for _, rawKey := range cfg.NonInferredBaseCategories.CustomKeys {
		key, exist := stringToEbitenKeyMap[rawKey]
		if !exist {
			return fmt.Errorf("custom_keys includes a key that does not exist: %s", rawKey)
		}
		CustomKeys = append(CustomKeys, key)
	}
	for _, charStr := range cfg.NonInferredBaseCategories.CustomChars {
		r, isValid := convertStringToRune(charStr)
		if !isValid {
			return fmt.Errorf("custom_chars includes a string that is not a valid rune: %s", charStr)
		}
		CustomChars = append(CustomChars, r)
	}

	BackgroundColor, err = getRGBAFromHex(cfg.Colors.Background)
	if err != nil {
		return err
	}
	UpcomingComboColor, err = getRGBAFromHex(cfg.Colors.UpcomingCombo)
	if err != nil {
		return err
	}
	CurrentComboColor, err = getRGBAFromHex(cfg.Colors.CurrentCombo)
	if err != nil {
		return err
	}
	CorrectPastComboColor, err = getRGBAFromHex(cfg.Colors.CorrectPastCombo)
	if err != nil {
		return err
	}
	IncorrectPastComboColor, err = getRGBAFromHex(cfg.Colors.IncorrectPastCombo)
	if err != nil {
		return err
	}

	return nil
}

func getRGBAFromHex(hexString string) (color.RGBA, error) {
	hexString = strings.TrimPrefix(hexString, "#")

	var r, g, b uint8
	n, err := fmt.Sscanf(hexString, "%02x%02x%02x", &r, &g, &b)

	if err != nil || n != 3 {
		return color.RGBA{}, fmt.Errorf("invalid hex color string: %s", hexString)
	}

	return color.RGBA{r, g, b, 255}, nil
}

var stringToEbitenKeyMap = map[string]ebiten.Key{
	"A": ebiten.KeyA, "B": ebiten.KeyB, "C": ebiten.KeyC, "D": ebiten.KeyD, "E": ebiten.KeyE, "F": ebiten.KeyF,
	"G": ebiten.KeyG, "H": ebiten.KeyH, "I": ebiten.KeyI, "J": ebiten.KeyJ, "K": ebiten.KeyK, "L": ebiten.KeyL,
	"M": ebiten.KeyM, "N": ebiten.KeyN, "O": ebiten.KeyO, "P": ebiten.KeyP, "Q": ebiten.KeyQ, "R": ebiten.KeyR,
	"S": ebiten.KeyS, "T": ebiten.KeyT, "U": ebiten.KeyU, "V": ebiten.KeyV, "W": ebiten.KeyW, "X": ebiten.KeyX,
	"Y": ebiten.KeyY, "Z": ebiten.KeyZ, "Alt": ebiten.KeyAlt, "AltLeft": ebiten.KeyAltLeft, "AltRight": ebiten.KeyAltRight,
	"ArrowDown": ebiten.KeyArrowDown, "ArrowLeft": ebiten.KeyArrowLeft, "ArrowRight": ebiten.KeyArrowRight, "ArrowUp": ebiten.KeyArrowUp,
	"Backquote": ebiten.KeyBackquote, "Backslash": ebiten.KeyBackslash, "Backspace": ebiten.KeyBackspace,
	"BracketLeft": ebiten.KeyBracketLeft, "BracketRight": ebiten.KeyBracketRight, "CapsLock": ebiten.KeyCapsLock,
	"Comma": ebiten.KeyComma, "ContextMenu": ebiten.KeyContextMenu, "Control": ebiten.KeyControl,
	"ControlLeft": ebiten.KeyControlLeft, "ControlRight": ebiten.KeyControlRight, "Delete": ebiten.KeyDelete,
	"Digit0": ebiten.KeyDigit0, "Digit1": ebiten.KeyDigit1, "Digit2": ebiten.KeyDigit2, "Digit3": ebiten.KeyDigit3,
	"Digit4": ebiten.KeyDigit4, "Digit5": ebiten.KeyDigit5, "Digit6": ebiten.KeyDigit6, "Digit7": ebiten.KeyDigit7,
	"Digit8": ebiten.KeyDigit8, "Digit9": ebiten.KeyDigit9, "End": ebiten.KeyEnd, "Enter": ebiten.KeyEnter,
	"Equal": ebiten.KeyEqual, "Escape": ebiten.KeyEscape, "F1": ebiten.KeyF1, "F2": ebiten.KeyF2, "F3": ebiten.KeyF3,
	"F4": ebiten.KeyF4, "F5": ebiten.KeyF5, "F6": ebiten.KeyF6, "F7": ebiten.KeyF7, "F8": ebiten.KeyF8,
	"F9": ebiten.KeyF9, "F10": ebiten.KeyF10, "F11": ebiten.KeyF11, "F12": ebiten.KeyF12, "F13": ebiten.KeyF13,
	"F14": ebiten.KeyF14, "F15": ebiten.KeyF15, "F16": ebiten.KeyF16, "F17": ebiten.KeyF17, "F18": ebiten.KeyF18,
	"F19": ebiten.KeyF19, "F20": ebiten.KeyF20, "F21": ebiten.KeyF21, "F22": ebiten.KeyF22, "F23": ebiten.KeyF23,
	"F24": ebiten.KeyF24, "Home": ebiten.KeyHome, "Insert": ebiten.KeyInsert, "IntlBackslash": ebiten.KeyIntlBackslash,
	"Meta": ebiten.KeyMeta, "MetaLeft": ebiten.KeyMetaLeft, "MetaRight": ebiten.KeyMetaRight, "Minus": ebiten.KeyMinus,
	"NumLock": ebiten.KeyNumLock, "Numpad0": ebiten.KeyNumpad0, "Numpad1": ebiten.KeyNumpad1, "Numpad2": ebiten.KeyNumpad2,
	"Numpad3": ebiten.KeyNumpad3, "Numpad4": ebiten.KeyNumpad4, "Numpad5": ebiten.KeyNumpad5, "Numpad6": ebiten.KeyNumpad6,
	"Numpad7": ebiten.KeyNumpad7, "Numpad8": ebiten.KeyNumpad8, "Numpad9": ebiten.KeyNumpad9, "NumpadAdd": ebiten.KeyNumpadAdd,
	"NumpadDecimal": ebiten.KeyNumpadDecimal, "NumpadDivide": ebiten.KeyNumpadDivide, "NumpadEnter": ebiten.KeyNumpadEnter,
	"NumpadEqual": ebiten.KeyNumpadEqual, "NumpadMultiply": ebiten.KeyNumpadMultiply, "NumpadSubtract": ebiten.KeyNumpadSubtract,
	"PageDown": ebiten.KeyPageDown, "PageUp": ebiten.KeyPageUp, "Pause": ebiten.KeyPause, "Period": ebiten.KeyPeriod,
	"PrintScreen": ebiten.KeyPrintScreen, "Quote": ebiten.KeyQuote, "ScrollLock": ebiten.KeyScrollLock, "Semicolon": ebiten.KeySemicolon,
	"Shift": ebiten.KeyShift, "ShiftLeft": ebiten.KeyShiftLeft, "ShiftRight": ebiten.KeyShiftRight, "Slash": ebiten.KeySlash,
	"Space": ebiten.KeySpace, "Tab": ebiten.KeyTab,
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
