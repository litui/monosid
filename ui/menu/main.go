package menu

import (
	"image/color"

	"github.com/litui/monosid/config"
	"github.com/litui/monosid/ui/rotaryencoder"
	"tinygo.org/x/drivers/ssd1306"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/notosans"
)

type Menu uint8

const (
	INCOMING_NOTE Menu = iota
	MIDI_CHANNEL
	WAVEFORM
	DETUNE_CENTS
	// PULSEWIDTH
	ATTACK
	DECAY
	SUSTAIN
	RELEASE
	DEBUG_LOG
	MENU_LENGTH
)

const (
	sixColWidth   = config.DISPLAY_WIDTH / 6
	threeColWidth = config.DISPLAY_WIDTH / 3
	twoColWidth   = config.DISPLAY_WIDTH / 2
	headerY       = (config.DISPLAY_HEIGHT / 2) - 2
	mainY         = config.DISPLAY_HEIGHT - 2
)

type EncoderRange struct {
	bottom        int
	top           int
	displayOffset int
}

var (
	BLACK     = color.RGBA{0, 0, 0, 255}
	WHITE     = color.RGBA{1, 1, 1, 255}
	largeFont = &notosans.Notosans12pt

	currentMenu Menu = DEBUG_LOG

	oldEncoderValue = [3]int{0, 0, 0}

	// bottom, top
	encoderRange = [MENU_LENGTH][3][2]int{
		{ // INCOMING_NOTE
			{0, 0},
			{0, 0},
			{0, 0},
		},
		{ // MIDI_CHANNEL
			{0, 15},
			{0, 15},
			{0, 15},
		},
		{ // WAVEFORM
			{1, 15},
			{1, 15},
			{1, 15},
		},
		{ // DETUNE CENTS
			{-99, 99},
			{-99, 99},
			{-99, 99},
		},
		{ // ATTACK RATE
			{0, 15},
			{0, 15},
			{0, 15},
		},
		{ // DECAY RATE
			{0, 15},
			{0, 15},
			{0, 15},
		},
		{ // SUSTAIN LEVEL
			{0, 15},
			{0, 15},
			{0, 15},
		},
		{ // RELEASE RATE
			{0, 15},
			{0, 15},
			{0, 15},
		},
		{ // DEBUG_LOG
			{0, 0},
			{0, 0},
			{0, 0},
		},
	}
)

func RenderMainMenu(display *ssd1306.Device, subEncoder []*rotaryencoder.Device) {
	display.ClearBuffer()

	switch currentMenu {
	case DEBUG_LOG:
		renderLogMenu(display)
	case INCOMING_NOTE:
		renderIncomingNoteMenu(display)
	case MIDI_CHANNEL:
		processChannelMenuEncoders(subEncoder)
		renderChannelMenu(display)
	case WAVEFORM:
		processWaveformMenuEncoders(subEncoder)
		renderWaveformMenu(display)
	case DETUNE_CENTS:
		processDetuneMenuEncoders(subEncoder)
		renderDetuneMenu(display)
	case ATTACK:
		processAttackMenuEncoders(subEncoder)
		renderAttackMenu(display)
	case DECAY:
		processDecayMenuEncoders(subEncoder)
		renderDecayMenu(display)
	case SUSTAIN:
		processSustainMenuEncoders(subEncoder)
		renderSustainMenu(display)
	case RELEASE:
		processReleaseMenuEncoders(subEncoder)
		renderReleaseMenu(display)
	}

	display.Display()
}

func ChangeMainMenu(encoderZero *rotaryencoder.Device) bool {
	newMenu := Menu(encoderZero.Value())
	if newMenu != currentMenu {
		currentMenu = newMenu
		return true
	}
	return false
}

func SetupEncoderMenuRanges(subEncoder []*rotaryencoder.Device) {
	for ei, e := range subEncoder {
		e.SetRange(encoderRange[currentMenu][ei][0], encoderRange[currentMenu][ei][1])
	}

	// Set initial values
	switch currentMenu {
	case DEBUG_LOG:
	case INCOMING_NOTE:
	case MIDI_CHANNEL:
		initChannelMenuValues(subEncoder)
	case WAVEFORM:
		initWaveformMenuValues(subEncoder)
	case DETUNE_CENTS:
		initDetuneMenuValues(subEncoder)
	case ATTACK:
		initAttackMenuValues(subEncoder)
	case DECAY:
		initDecayMenuValues(subEncoder)
	case SUSTAIN:
		initSustainMenuValues(subEncoder)
	case RELEASE:
		initReleaseMenuValues(subEncoder)
	}
}

func writeHeader(display *ssd1306.Device, header string) {
	tinyfont.WriteLineRotated(display, largeFont, 0, headerY, header, WHITE, tinyfont.NO_ROTATION)
}

func write3Box(display *ssd1306.Device, pos uint8, text string) {
	inner, _ := tinyfont.LineWidth(largeFont, text)
	exact := ((threeColWidth - int16(inner)) / 2) + int16(pos*threeColWidth)
	tinyfont.WriteLineRotated(display, largeFont, exact, mainY, text, WHITE, tinyfont.NO_ROTATION)
}
