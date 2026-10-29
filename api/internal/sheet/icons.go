package sheet

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"regexp"
	"sync"

	"golang.org/x/image/webp"
)

// assetFS is the site's own art the sheet draws, copied byte for byte from web/public by
// `make sheet-assets`: the nine 128 px HD crests, the 72 px faction marks and the 400 px seal.
//
//go:embed assets/crests/*.webp assets/faction/*.webp assets/seal-400.png
var assetFS embed.FS

// QuestionMark is the icon an item unknown to the build draws (spec §A.4, §B.3).
const QuestionMark = "inv_misc_questionmark"

// IconSource gives the sheet and the tiles their art: a game icon by its file name (no
// extension), and the class's ringed HD crest by class slug.
type IconSource interface {
	Icon(name string) (image.Image, error)
	Crest(classSlug string) (image.Image, error)
}

// iconName is the shape of an icon file name in the build's folder; anything else is refused
// before it reaches the file system.
var iconName = regexp.MustCompile(`^[a-z0-9_]+$`)

// ErrNoIcon is an icon the folder does not hold.
var ErrNoIcon = errors.New("sheet: no such icon")

// FileIcons reads a build's icon folder (data/builds/{build}/icons/*.webp, /data/{build}/icons
// in the image) and the embedded crests.
type FileIcons struct {
	Dir string
}

// NewFileIcons is a FileIcons over dir.
func NewFileIcons(dir string) *FileIcons { return &FileIcons{Dir: dir} }

// Icon decodes {Dir}/{name}.webp.
func (f *FileIcons) Icon(name string) (image.Image, error) {
	if !iconName.MatchString(name) {
		return nil, fmt.Errorf("%w: %q", ErrNoIcon, name)
	}
	raw, err := os.ReadFile(filepath.Join(f.Dir, name+".webp"))
	if errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("%w: %s", ErrNoIcon, name)
	}
	if err != nil {
		return nil, fmt.Errorf("sheet: icon %s: %w", name, err)
	}
	img, err := webp.Decode(bytes.NewReader(raw))
	if err != nil {
		return nil, fmt.Errorf("sheet: decode icon %s: %w", name, err)
	}
	return img, nil
}

// Crest is the embedded crest for classSlug.
func (f *FileIcons) Crest(classSlug string) (image.Image, error) { return EmbeddedCrest(classSlug) }

// EmbeddedCrest decodes the class's 128 px HD crest from the embedded copy.
func EmbeddedCrest(classSlug string) (image.Image, error) {
	if _, ok := classColors[classSlug]; !ok {
		return nil, fmt.Errorf("sheet: crest: %q is not a class", classSlug)
	}
	return embeddedImage("assets/crests/" + classSlug + ".webp")
}

func factionMark(faction string) (image.Image, error) {
	return embeddedImage("assets/faction/" + faction + ".webp")
}

func seal() (image.Image, error) { return embeddedImage("assets/seal-400.png") }

// embedded decodes each embedded image once; decoded images are only ever read.
var (
	embeddedMu    sync.Mutex
	embeddedCache = map[string]image.Image{}
)

func embeddedImage(name string) (image.Image, error) {
	embeddedMu.Lock()
	defer embeddedMu.Unlock()
	if img, ok := embeddedCache[name]; ok {
		return img, nil
	}
	raw, err := assetFS.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("sheet: asset %s: %w", name, err)
	}
	var img image.Image
	if filepath.Ext(name) == ".png" {
		img, err = png.Decode(bytes.NewReader(raw))
	} else {
		img, err = webp.Decode(bytes.NewReader(raw))
	}
	if err != nil {
		return nil, fmt.Errorf("sheet: decode %s: %w", name, err)
	}
	embeddedCache[name] = img
	return img, nil
}

// iconOrQuestionMark is the named icon, the question mark when the source lacks it.
func iconOrQuestionMark(src IconSource, name string) (image.Image, error) {
	if name != "" {
		img, err := src.Icon(name)
		if err == nil {
			return img, nil
		}
		if !errors.Is(err, ErrNoIcon) {
			return nil, err
		}
	}
	return src.Icon(QuestionMark)
}
