package main

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/BurntSushi/toml"
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

	//set all:
	//BaseWeightLowerLetters      int
	//BaseWeightUpperLetters      int
	//BaseWeightLowerSymbols      int
	//BaseWeightUpperSymbols      int
	//BaseWeightDigits            int
	//BaseWeightNonPrintableKeys1 int
	//BaseWeightNonPrintableKeys2 int
	//BaseWeightCustomChars       int
	//BaseWeightCustomKeys        int
	//
	//ModWeightNoModifiers         int
	//ModWeightControl             int
	//ModWeightAlt                 int
	//ModWeightShift               int
	//ModWeightMeta                int
	//ModWeightControlAlt          int
	//ModWeightControlShift        int
	//ModWeightControlMeta         int
	//ModWeightAltShift            int
	//ModWeightAltMeta             int
	//ModWeightShiftMeta           int
	//ModWeightControlAltShift     int
	//ModWeightControlAltMeta      int
	//ModWeightControlShiftMeta    int
	//ModWeightAltShiftMeta        int
	//ModWeightControlAltShiftMeta int

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
