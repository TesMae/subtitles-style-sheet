package main

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"
)

func TestTextDecoration(t *testing.T) {
	input0 := ""
	input01 := "bold italic underline line-through something-else"

	input1 := "none"

	input2 := "bold italic underline line-through"

	input21 := "italic underline line-through"
	input22 := "bold underline line-through"
	input23 := "bold italic line-through"
	input24 := "bold italic underline"

	input31 := "bold italic"
	input32 := "bold underline"
	input33 := "bold line-through"
	input34 := "italic underline"
	input35 := "italic line-through"
	input36 := "underline line-through"

	input41 := "bold"
	input42 := "italic"
	input43 := "underline"
	input44 := "line-through"

	expectedOutput1 := TextDecoration{
		Bold:      0,
		Italic:    0,
		Underline: 0,
		StrikeOut: 0,
	}
	expectedOutput2 := TextDecoration{
		Bold:      1,
		Italic:    1,
		Underline: 1,
		StrikeOut: 1,
	}
	expectedOutput21 := TextDecoration{
		Bold:      0,
		Italic:    1,
		Underline: 1,
		StrikeOut: 1,
	}
	expectedOutput22 := TextDecoration{
		Bold:      1,
		Italic:    0,
		Underline: 1,
		StrikeOut: 1,
	}
	expectedOutput23 := TextDecoration{
		Bold:      1,
		Italic:    1,
		Underline: 0,
		StrikeOut: 1,
	}
	expectedOutput24 := TextDecoration{
		Bold:      1,
		Italic:    1,
		Underline: 1,
		StrikeOut: 0,
	}

	expectedOutput31 := TextDecoration{
		Bold:      1,
		Italic:    1,
		Underline: 0,
		StrikeOut: 0,
	}
	expectedOutput32 := TextDecoration{
		Bold:      1,
		Italic:    0,
		Underline: 1,
		StrikeOut: 0,
	}
	expectedOutput33 := TextDecoration{
		Bold:      1,
		Italic:    0,
		Underline: 0,
		StrikeOut: 1,
	}
	expectedOutput34 := TextDecoration{
		Bold:      0,
		Italic:    1,
		Underline: 1,
		StrikeOut: 0,
	}
	expectedOutput35 := TextDecoration{
		Bold:      0,
		Italic:    1,
		Underline: 0,
		StrikeOut: 1,
	}
	expectedOutput36 := TextDecoration{
		Bold:      0,
		Italic:    0,
		Underline: 1,
		StrikeOut: 1,
	}

	expectedOutput41 := TextDecoration{
		Bold:      1,
		Italic:    0,
		Underline: 0,
		StrikeOut: 0,
	}
	expectedOutput42 := TextDecoration{
		Bold:      0,
		Italic:    1,
		Underline: 0,
		StrikeOut: 0,
	}
	expectedOutput43 := TextDecoration{
		Bold:      0,
		Italic:    0,
		Underline: 1,
		StrikeOut: 0,
	}
	expectedOutput44 := TextDecoration{
		Bold:      0,
		Italic:    0,
		Underline: 0,
		StrikeOut: 1,
	}

	if mapTextDecoration(input0) != expectedOutput1 {
		t.Error("Text Decoration mapping 0 Failed !")
	}
	if mapTextDecoration(input01) != expectedOutput1 {
		t.Error("Text Decoration mapping 01 Failed !")
	}
	if mapTextDecoration(input1) != expectedOutput1 {
		t.Error("Text Decoration mapping 1 Failed !")
	}
	if mapTextDecoration(input2) != expectedOutput2 {
		t.Error("Text Decoration mapping 2 Failed !")
	}

	if mapTextDecoration(input21) != expectedOutput21 {
		t.Error("Text Decoration mapping 21 Failed !")
	}
	if mapTextDecoration(input22) != expectedOutput22 {
		t.Error("Text Decoration mapping 22 Failed !")
	}
	if mapTextDecoration(input23) != expectedOutput23 {
		t.Error("Text Decoration mapping 23 Failed !")
	}
	if mapTextDecoration(input24) != expectedOutput24 {
		t.Error("Text Decoration mapping 24 Failed !")
	}

	if mapTextDecoration(input31) != expectedOutput31 {
		t.Error("Text Decoration mapping 31 Failed !")
	}
	if mapTextDecoration(input32) != expectedOutput32 {
		t.Error("Text Decoration mapping 32 Failed !")
	}
	if mapTextDecoration(input33) != expectedOutput33 {
		t.Error("Text Decoration mapping 33 Failed !")
	}
	if mapTextDecoration(input34) != expectedOutput34 {
		t.Error("Text Decoration mapping 34 Failed !")
	}
	if mapTextDecoration(input35) != expectedOutput35 {
		t.Error("Text Decoration mapping 35 Failed !")
	}
	if mapTextDecoration(input36) != expectedOutput36 {
		t.Error("Text Decoration mapping 36 Failed !")
	}

	if mapTextDecoration(input41) != expectedOutput41 {
		t.Error("Text Decoration mapping 41 Failed !")
	}
	if mapTextDecoration(input42) != expectedOutput42 {
		t.Error("Text Decoration mapping 42 Failed !")
	}
	if mapTextDecoration(input43) != expectedOutput43 {
		t.Error("Text Decoration mapping 43 Failed !")
	}
	if mapTextDecoration(input44) != expectedOutput44 {
		t.Error("Text Decoration mapping 44 Failed !")
	}

}

func TestParseNamedColor(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		ok       bool
	}{
		{"yellow", "&H0000FFFF", true},
		{"red", "&H000000FF", true},
		{"lime", "&H0000FF00", true},
		{"green", "&H00008000", true},
		{"cyan", "&H00FFFF00", true},
		{"magenta", "&H00FF00FF", true},
		{"YELLOW", "&H0000FFFF", true},
		{"unknown", "", false},
		{"", "", false},
	}
	for _, tt := range tests {
		hex, ok := parseNamedColor(tt.input)
		if hex != tt.expected || ok != tt.ok {
			t.Errorf("parseNamedColor(%q) = (%q, %v), want (%q, %v)", tt.input, hex, ok, tt.expected, tt.ok)
		}
	}
}

func TestParseCSSColor(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		err      bool
	}{
		{"rgba(255, 255, 255, 1)", "&H00FFFFFF", false},
		{"rgba(255, 0, 0, 1)", "&H000000FF", false},
		{"rgba(0, 0, 0, 1)", "&H00000000", false},
		{"rgba(0, 255, 0, 1)", "&H0000FF00", false},
		{"rgba(255, 255, 0, 1)", "&H0000FFFF", false},
		{"rgba(128, 0, 128, 1)", "&H00800080", false},
		{"rgba(255, 255, 255, 0)", "&HFFFFFFFF", false},
		{"invalid", "", true},
		{"rgba(1, 2)", "", true},
	}
	for _, tt := range tests {
		hex, err := parseCSSColor(tt.input)
		if tt.err {
			if err == nil {
				t.Errorf("parseCSSColor(%q) expected error", tt.input)
			}
		} else {
			if err != nil {
				t.Errorf("parseCSSColor(%q) unexpected error: %v", tt.input, err)
			}
			if hex != tt.expected {
				t.Errorf("parseCSSColor(%q) = %q, want %q", tt.input, hex, tt.expected)
			}
		}
	}
}

func TestParseColor(t *testing.T) {
	if got := parseColor("red"); got != "&H000000FF" {
		t.Errorf("parseColor(\"red\") = %q, want %q", got, "&H000000FF")
	}
	if got := parseColor("YELLOW"); got != "&H0000FFFF" {
		t.Errorf("parseColor(\"YELLOW\") = %q, want %q", got, "&H0000FFFF")
	}
	if got := parseColor("rgba(0, 255, 0, 1)"); got != "&H0000FF00" {
		t.Errorf("parseColor(\"rgba(0,255,0,1)\") = %q, want %q", got, "&H0000FF00")
	}
	if got := parseColor("rgba(255, 255, 255, 1)"); got != "&H00FFFFFF" {
		t.Errorf("parseColor(\"rgba(255,255,255,1)\") = %q, want %q", got, "&H00FFFFFF")
	}
}

func TestMapAlignment(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"bottom-left", 1},
		{"bottom-center", 2},
		{"bottom-right", 3},
		{"middle-left", 4},
		{"middle-center", 5},
		{"middle-right", 6},
		{"top-left", 7},
		{"top-center", 8},
		{"top-right", 9},
		{"invalid", 2},
	}
	for _, tt := range tests {
		if got := mapAlignment(tt.input); got != tt.expected {
			t.Errorf("mapAlignment(%q) = %d, want %d", tt.input, got, tt.expected)
		}
	}
}

func TestMapBorderStyle(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"outline-shadow", 1},
		{"opaque-box", 3},
		{"invalid", 1},
	}
	for _, tt := range tests {
		if got := mapBorderStyle(tt.input); got != tt.expected {
			t.Errorf("mapBorderStyle(%q) = %d, want %d", tt.input, got, tt.expected)
		}
	}
}

func TestMapEncoding(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"ansi", 0},
		{"utf-8", 1},
		{"UTF-8", 1},
		{"0", 0},
		{"1", 1},
		{"invalid", 0},
	}
	for _, tt := range tests {
		if got := mapEncoding(tt.input); got != tt.expected {
			t.Errorf("mapEncoding(%q) = %d, want %d", tt.input, got, tt.expected)
		}
	}
}

func TestConvertGlobalStyle(t *testing.T) {
	props := IR1StyleProps{
		FontFamily:            "Arial",
		FontSize:              20,
		Color:                 "rgba(255, 255, 255, 1)",
		KaraokeSecondaryColor: "rgba(255, 0, 0, 1)",
		TextStrokeColor:       "rgba(0, 0, 0, 1)",
		BoxBackgroundColor:    "rgba(0, 0, 0, 1)",
		TextDecoration:        "none",
		Transform:             IR1Transform{ScaleX: 1, ScaleY: 1, Rotate: 0},
		LetterSpacing:         0,
		BorderStyle:           "outline-shadow",
		TextStrokeWidth:       2,
		TextShadowDepth:       0,
		Alignment:             "middle-center",
		MarginLeft:            10,
		MarginRight:           10,
		MarginVertical:        540,
		Encoding:              "ansi",
	}
	got := convertGlobalStyle("Default", props)

	if got.Name != "Default" {
		t.Errorf("Name = %q, want %q", got.Name, "Default")
	}
	if got.Fontname != "Arial" {
		t.Errorf("Fontname = %q, want %q", got.Fontname, "Arial")
	}
	if got.Fontsize != 20 {
		t.Errorf("Fontsize = %d, want %d", got.Fontsize, 20)
	}
	if got.PrimaryColour != "&H00FFFFFF" {
		t.Errorf("PrimaryColour = %q, want %q", got.PrimaryColour, "&H00FFFFFF")
	}
	if got.SecondaryColour != "&H000000FF" {
		t.Errorf("SecondaryColour = %q, want %q", got.SecondaryColour, "&H000000FF")
	}
	if got.OutlineColour != "&H00000000" {
		t.Errorf("OutlineColour = %q, want %q", got.OutlineColour, "&H00000000")
	}
	if got.BackColour != "&H00000000" {
		t.Errorf("BackColour = %q, want %q", got.BackColour, "&H00000000")
	}
	if got.Bold != 0 {
		t.Errorf("Bold = %d, want %d", got.Bold, 0)
	}
	if got.Italic != 0 {
		t.Errorf("Italic = %d, want %d", got.Italic, 0)
	}
	if got.Underline != 0 {
		t.Errorf("Underline = %d, want %d", got.Underline, 0)
	}
	if got.StrikeOut != 0 {
		t.Errorf("StrikeOut = %d, want %d", got.StrikeOut, 0)
	}
	if got.ScaleX != 100 {
		t.Errorf("ScaleX = %d, want %d", got.ScaleX, 100)
	}
	if got.ScaleY != 100 {
		t.Errorf("ScaleY = %d, want %d", got.ScaleY, 100)
	}
	if got.Spacing != 0 {
		t.Errorf("Spacing = %d, want %d", got.Spacing, 0)
	}
	if got.Angle != 0 {
		t.Errorf("Angle = %d, want %d", got.Angle, 0)
	}
	if got.BorderStyle != 1 {
		t.Errorf("BorderStyle = %d, want %d", got.BorderStyle, 1)
	}
	if got.Outline != 2 {
		t.Errorf("Outline = %d, want %d", got.Outline, 2)
	}
	if got.Shadow != 0 {
		t.Errorf("Shadow = %d, want %d", got.Shadow, 0)
	}
	if got.Alignment != 5 {
		t.Errorf("Alignment = %d, want %d", got.Alignment, 5)
	}
	if got.MarginL != 10 {
		t.Errorf("MarginL = %d, want %d", got.MarginL, 10)
	}
	if got.MarginR != 10 {
		t.Errorf("MarginR = %d, want %d", got.MarginR, 10)
	}
	if got.MarginV != 540 {
		t.Errorf("MarginV = %d, want %d", got.MarginV, 540)
	}
	if got.Encoding != 0 {
		t.Errorf("Encoding = %d, want %d", got.Encoding, 0)
	}
}

func TestJsonGeneration(t *testing.T) {
	JsonGeneration("ir_1.json", "ir_example_test.json")
	defer os.Remove("ir_example_test.json")

	data, err := os.ReadFile("ir_example_test.json")
	if err != nil {
		t.Fatal("failed to read generated file:", err)
	}

	var ir IR
	if err := json.Unmarshal(data, &ir); err != nil {
		t.Fatal("invalid JSON output:", err)
	}

	if len(ir.GlobalStyle) != 1 {
		t.Fatalf("expected 1 global style, got %d", len(ir.GlobalStyle))
	}

	s := ir.GlobalStyle[0]
	if s.Name != "Default" {
		t.Errorf("Name = %q, want %q", s.Name, "Default")
	}
	if s.Fontname != "Arial" {
		t.Errorf("Fontname = %q, want %q", s.Fontname, "Arial")
	}
	if s.PrimaryColour != "&H00FFFFFF" {
		t.Errorf("PrimaryColour = %q, want %q", s.PrimaryColour, "&H00FFFFFF")
	}
	if s.BorderStyle != 1 {
		t.Errorf("BorderStyle = %d, want %d", s.BorderStyle, 1)
	}
	if s.Alignment != 5 {
		t.Errorf("Alignment = %d, want %d", s.Alignment, 5)
	}
	if s.ScaleX != 100 {
		t.Errorf("ScaleX = %d, want %d", s.ScaleX, 100)
	}
	if s.Encoding != 0 {
		t.Errorf("Encoding = %d, want %d", s.Encoding, 0)
	}
}

func TestAssContent(t *testing.T) {
	outputExample, err := os.ReadFile("output_example.ass")
	if err != nil {
		t.Fatal("Error occured when trying to read the output_example.ass file")
	}

	output, err := os.ReadFile("output.ass")
	if err != nil {
		t.Fatal("Error occured when trying to read the output.ass file")
	}

	if !bytes.Equal(outputExample, output) {
		t.Error("ass outputs are different")
	}
}
