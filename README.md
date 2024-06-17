## MONOSID

This is 3rd-party, work-in-progress firmware by Litui (Aria Burrell) for the MIDISID (original hardware by Shiela Dixon and available on [Tindie](https://www.tindie.com/products/shieladixon/midisid-midi-in-2-x-sid-clones/)). This is not yet a feature-complete synthesizer firmware! If you flash it and need to recover, you will need to obtain the latest version of the original firmware from the hardware creator.

This started as part of a successful attempt to reverse engineer the Raspberry Pi Pico-based MIDISID hardware and write my own custom-tailored firmware. The original firmware is very good but tailored more toward General MIDI and the like. I wanted the SID to act more like a mono synth with more immediate controls for live techno/chiptune performance.

While I started writing this custom firmware in PlatformIO/Arduino, I later changed gears and started rewriting what I'd completed in TinyGo after realizing its RP2040 support was mature enough to make this work and provide me a good opportunity to learn Golang.

I haven't yet finished rewriting my custom firmware much less completed the full feature set, but it's coming along nicely.

### Working
* General settings
    * MIDI Channel
    * Selected patch
* Patch settings
    * Waveform
    * Detune (cents)
    * Attack rate
    * Decay rate
    * Sustain level
    * Release rate
* Log view
* Save/Load menu (press first encoder)
    * Saving and loading works but hasn't been extensively tested for bugs

### To-do

* Configuration screens for all basic parameters
    * General volume setting
    * Pulse width
    * Filter cutoff and resonance
    * Ring modulation
 * Extras
