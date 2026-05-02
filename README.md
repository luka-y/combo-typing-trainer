Incredibly simple typing trainer: just a stream of gibberish that besides regular characters can
contain any key on the keyboard, as well as combinations with modifiers (Ctrl, Alt, Shift, Meta).

*TODO: put gif of it here*

Inspired by [Programmer's Typing Practice by climech](https://climech.github.io/typing-practice/).

# Installation

Download the binary from [the latest release](/releases).

### OR build from source

Clone the repo and `go build .` (you would need go and
[ebiten dependencies](https://ebitengine.org/en/documents/install.html)
installed).

# Configuration

Every time you run the binary it looks for `config.toml` in the same directory with it.
If absent, it will create a new, [default one](assets/default-config.toml).

In the config, you can configure:

### Probability weights

The primary thing you would configure. It responsible for chance with which certain categories of chars and keys appear.
For example:

Just the letters:
*TODO: gif of a steam*

Just the symbols: *TODO: gif of a steam*

Just a selected set of keys: *TODO: gif of a steam*

Modifier heavy with just letters: *TODO: gif of a steam*

Or any combination of just the parts you want with any probabilities you might want.

### Layouts and custom categories

You can add layouts, for example for different languages, or things like Dvorak or Colemak.

Most categories are inferred from the layout (letters, symbols, etc.), but some you can set yourself.
Category is just a set of chars or keys that you can set own probability weights on.

### Appearance

You can also change all colors and font size.

*TODO: add one or few gifs with different colors and font sizes*

# Note for agents

The project is tiny (less than 2k loc), load the entire thing into your context, no need to pick selectively.