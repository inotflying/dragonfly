package chat

import (
	"fmt"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"golang.org/x/text/language"
)

// https://github.com/Mojang/bedrock-samples/blob/main/resource_pack/texts/en_GB.lang

var MessageJoin = Translate(str("%multiplayer.player.joined"), 1, `%v joined the game`).Enc("<yellow>%v</yellow>")
var MessageQuit = Translate(str("%multiplayer.player.left"), 1, `%v left the game`).Enc("<yellow>%v</yellow>")
var MessageServerDisconnect = Translate(str("%disconnect.disconnected"), 0, `Disconnected by Server`).Enc("<yellow>%v</yellow>")

var MessageCommandSyntax = Translate(str("%commands.generic.syntax"), 3, `Syntax error: unexpected value: at "%v>>%v<<%v"`)
var MessageCommandUsage = Translate(str("%commands.generic.usage"), 1, `Usage: %v`)
var MessageCommandUnknown = Translate(str("%commands.generic.unknown"), 1, `Unknown command: "%v": Please check that the command exists and that you have permission to use it.`)
var MessageCommandNoTargets = Translate(str("%commands.generic.noTargetMatch"), 0, `No targets matched selector`)

type str string

// Resolve returns the translation identifier as a string.
func (s str) Resolve(language.Tag) string { return string(s) }

// TranslationString is a value that can resolve a translated version of itself
// for a language.Tag passed.
type TranslationString interface {
	// Resolve finds a suitable translated version for a translation string for
	// a specific language.Tag.
	Resolve(l language.Tag) string
}

// Translate returns a Translation for a TranslationString. The required number
// of parameters specifies how many arguments may be passed to Translation.F.
// The fallback string should be a 'standard' translation of the string, which
// is used when translation.String is called on the translation that results
// from a call to Translation.F. This fallback string should have as many
// formatting identifiers (like in fmt.Sprintf) as the number of params.
func Translate(str TranslationString, params int, fallback string) Translation {
	return Translation{str: str, params: params, fallback: fallback, format: "%v"}
}

// Translation represents a TranslationString with additional formatting, that
// may be filled out by calling F on it with a list of arguments for the
// translation.
type Translation struct {
	str      TranslationString
	format   string
	params   int
	fallback string
}

// Zero returns false if a Translation was not created using Translate or
// Untranslated.
func (t Translation) Zero() bool {
	return t.format == ""
}

// Enc encapsulates the translation string into the format passed. This format
// should have exactly one formatting identifier, %v, to specify where the
// translation string should go, such as 'translation: %v'.
// Enc accepts colouring formats parsed by text.Colourf.
func (t Translation) Enc(format string) Translation {
	t.format = format
	return t
}

// Resolve passes 0 arguments to the translation and resolves the translation
// string for the language passed. It is equal to calling t.F().Resolve(l).
// Resolve panics if the Translation requires at least 1 argument.
func (t Translation) Resolve(l language.Tag) string {
	return t.F().Resolve(l)
}

// F takes arguments for a translation string passed and returns a filled out
// translation that may be sent to players. The number of arguments passed must
// be exactly equal to the number specified in Translate. If not, F will panic.
func (t Translation) F(a ...any) translation {
	if len(a) != t.params {
		panic(fmt.Sprintf("translation '%v' requires exactly %v parameters, got %v", t.format, t.params, len(a)))
	}
	params := make([]string, len(a))
	for i, arg := range a {
		params[i] = fmt.Sprint(arg)
	}
	return translation{t: t, params: params, fallbackParams: a}
}

// translation is a translation string with its arguments filled out. Resolve may
// be called to obtain the translated version of the translation string and
// Params may be called to obtain the parameters passed in Translation.F.
// translation implements the fmt.Stringer and error interfaces.
type translation struct {
	t              Translation
	params         []string
	fallbackParams []any
}

// Resolve translates the TranslationString of the translation to the language
// passed and returns it.
func (t translation) Resolve(l language.Tag) string {
	return text.Colourf(t.t.format, t.t.str.Resolve(l))
}

// Params returns a slice of values that are used to parameterise the
// translation returned by Resolve.
func (t translation) Params() []string {
	return t.params
}

// String formats and returns the fallback value of the translation.
func (t translation) String() string {
	return fmt.Sprintf(text.Colourf(t.t.format, t.t.fallback), t.fallbackParams...)
}

// Error formats and returns the fallback value of the translation.
func (t translation) Error() string {
	return t.String()
}
