package main

import (
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
	return err
}
