package discordbot

import (
	"encoding/json"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/jhunthrop/foreversixty/api/internal/notify"
)

// componentsV2Flag is IS_COMPONENTS_V2 (spec 2026-10-29 §0): the message is drawn from its
// components alone, so it carries no content and no embeds. Once set it can never be unset.
const componentsV2Flag = 1 << 15

// The Components V2 layout kinds (spec §0).
const (
	KindSection      ComponentKind = 9
	KindTextDisplay  ComponentKind = 10
	KindThumbnail    ComponentKind = 11
	KindMediaGallery ComponentKind = 12
	KindSeparator    ComponentKind = 14
	KindContainer    ComponentKind = 17
)

// separatorSmall is a Separator's small spacing.
const separatorSmall = 1

// Node is one top-level or nested Components V2 layout component; Kind decides which fields
// count:
//   - Container: Accent (nil for none) and Children.
//   - Section: Children (1 to 3 Text Displays) and Accessory (a Thumbnail).
//   - Text Display: Content, Discord markdown.
//   - Thumbnail: Media (an https or attachment:// URL) and Description, its alt text.
//   - Media Gallery: Items.
//   - Separator: always a divider, Spacing 1 or 2.
//
// Action rows are not Nodes: a V2 message keeps them in Message.Components, after its Layout.
type Node struct {
	Kind        ComponentKind
	Accent      *int
	Children    []Node
	Accessory   *Node
	Content     string
	Media       string
	Description string
	Items       []GalleryItem
	Spacing     int
}

// GalleryItem is one Media Gallery image with its alt text.
type GalleryItem struct {
	Media       string
	Description string
}

type mediaJSON struct {
	URL string `json:"url"`
}

type galleryItemJSON struct {
	Media       mediaJSON `json:"media"`
	Description string    `json:"description,omitempty"`
}

// MarshalJSON writes the fields of the node's kind in the reference payloads' order
// (design/mocks/2026-10-29-discord-gear-panel/payloads); a kind the bot does not draw fails here.
func (n Node) MarshalJSON() ([]byte, error) {
	switch n.Kind {
	case KindContainer:
		return json.Marshal(struct {
			Type     ComponentKind `json:"type"`
			Accent   *int          `json:"accent_color,omitempty"`
			Children []Node        `json:"components"`
		}{n.Kind, n.Accent, n.Children})
	case KindSection:
		if n.Accessory == nil {
			return nil, fmt.Errorf("discordbot: section without an accessory")
		}
		return json.Marshal(struct {
			Type      ComponentKind `json:"type"`
			Children  []Node        `json:"components"`
			Accessory *Node         `json:"accessory"`
		}{n.Kind, n.Children, n.Accessory})
	case KindTextDisplay:
		return json.Marshal(struct {
			Type    ComponentKind `json:"type"`
			Content string        `json:"content"`
		}{n.Kind, n.Content})
	case KindThumbnail:
		return json.Marshal(struct {
			Type        ComponentKind `json:"type"`
			Media       mediaJSON     `json:"media"`
			Description string        `json:"description,omitempty"`
		}{n.Kind, mediaJSON{n.Media}, n.Description})
	case KindMediaGallery:
		items := make([]galleryItemJSON, len(n.Items))
		for i, it := range n.Items {
			items[i] = galleryItemJSON{mediaJSON{it.Media}, it.Description}
		}
		return json.Marshal(struct {
			Type  ComponentKind     `json:"type"`
			Items []galleryItemJSON `json:"items"`
		}{n.Kind, items})
	case KindSeparator:
		return json.Marshal(struct {
			Type    ComponentKind `json:"type"`
			Divider bool          `json:"divider"`
			Spacing int           `json:"spacing"`
		}{n.Kind, true, max(n.Spacing, separatorSmall)})
	}
	return nil, fmt.Errorf("discordbot: layout kind %d is not drawn", n.Kind)
}

// textNode is a Text Display.
func textNode(content string) Node { return Node{Kind: KindTextDisplay, Content: content} }

// separatorNode is a small divider.
func separatorNode() Node { return Node{Kind: KindSeparator, Spacing: separatorSmall} }

// thumbnailNode is a Thumbnail with its alt text.
func thumbnailNode(url, alt string) *Node {
	return &Node{Kind: KindThumbnail, Media: url, Description: alt}
}

// sectionNode is one Text Display with an accessory; with no accessory it is the Text Display
// alone, since a Section needs one.
func sectionNode(content string, accessory *Node) Node {
	if accessory == nil {
		return textNode(content)
	}
	return Node{Kind: KindSection, Children: []Node{textNode(content)}, Accessory: accessory}
}

// containerNode is a Container with an accent colour.
func containerNode(accent int, children ...Node) Node {
	return Node{Kind: KindContainer, Accent: &accent, Children: children}
}

// galleryNode is a Media Gallery of one image.
func galleryNode(url, alt string) Node {
	return Node{Kind: KindMediaGallery, Items: []GalleryItem{{Media: url, Description: alt}}}
}

// countNodes is every component a layout holds, nested ones and accessories included.
func countNodes(nodes []Node) int {
	n := 0
	for _, node := range nodes {
		n++
		n += countNodes(node.Children)
		if node.Accessory != nil {
			n++
		}
	}
	return n
}

// textLength is the characters across every Text Display of a layout.
func textLength(nodes []Node) int {
	n := 0
	for _, node := range nodes {
		if node.Kind == KindTextDisplay {
			n += utf8.RuneCountInString(node.Content)
		}
		n += textLength(node.Children)
	}
	return n
}

// componentCount is every component a message holds: its layout, its rows and their parts
// (Discord's 40 counts them all; the bot holds 38, spec §A.6).
func (m Message) componentCount() int {
	n := countNodes(m.Layout)
	for _, row := range m.Components {
		n += 1 + len(row)
	}
	return n
}

// asComponentsV2 is m in the Components V2 encoding (spec §A.7): its content line one
// top-level Text Display, each embed a Container in its colour holding the author as subtext,
// the title as a heading (linked when it has a URL) and the description, beside the crest
// thumbnail when the embed has one. Footers, timestamps and fields are dropped, as V2 has
// none. crestAlt is the thumbnail's alt text. A message already V2 is returned as it is; one
// with rows alone becomes a V2 message of those rows.
func asComponentsV2(m Message, crestAlt string) Message {
	if m.componentsV2() {
		return m
	}
	var layout []Node
	if m.Content != "" {
		layout = append(layout, textNode(m.Content))
	}
	for _, e := range m.Embeds {
		layout = append(layout, embedContainer(e, crestAlt))
	}
	m.Content, m.Embeds, m.Layout, m.ComponentsV2 = "", nil, layout, true
	return m
}

func embedContainer(e notify.Embed, crestAlt string) Node {
	var lines []string
	if e.Author != "" {
		lines = append(lines, subtext(e.Author))
	}
	if title := e.Title; title != "" {
		if e.URL != "" {
			title = "[" + title + "](" + e.URL + ")"
		}
		lines = append(lines, heading(title))
	}
	if e.Description != "" {
		lines = append(lines, e.Description)
	}
	var crest *Node
	if e.Thumbnail != "" {
		crest = thumbnailNode(e.Thumbnail, crestAlt)
	}
	body := sectionNode(strings.Join(lines, "\n"), crest)
	if e.Color == 0 {
		return Node{Kind: KindContainer, Children: []Node{body}}
	}
	return containerNode(e.Color, body)
}

// heading is a Text Display's heading line; subtext its small muted line.
func heading(s string) string { return "### " + s }
func subtext(s string) string { return "-# " + s }
