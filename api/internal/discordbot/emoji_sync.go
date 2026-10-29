package discordbot

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"time"

	xdraw "golang.org/x/image/draw"
	"golang.org/x/image/webp"

	"github.com/jhunthrop/foreversixty/api/internal/raidspec"
)

// The spec emoji images (spec 2026-10-28 §P.9.1): 128 px PNGs, Discord's emoji size, scaled from
// the game's 64 px talent-tab icons; each stays far under Discord's 256 KiB.
const (
	emojiSide     = 128
	emojiMaxBytes = 256 << 10
)

// EmojiImage is one spec's emoji, ready to upload.
type EmojiImage struct {
	Name string
	PNG  []byte
}

// EmojiImages builds every spec's emoji from iconsDir (web/public/data/{build}/icons), one per
// specs.json row, named fs_{class}_{spec}.
func EmojiImages(iconsDir string) ([]EmojiImage, error) {
	specs := raidspec.All()
	out := make([]EmojiImage, 0, len(specs))
	for _, sp := range specs {
		raw, err := os.ReadFile(filepath.Join(iconsDir, sp.Icon+".webp"))
		if err != nil {
			return nil, fmt.Errorf("discordbot: emoji %s: %w", sp.Emoji(), err)
		}
		img, err := emojiPNG(raw)
		if err != nil {
			return nil, fmt.Errorf("discordbot: emoji %s: %w", sp.Emoji(), err)
		}
		out = append(out, EmojiImage{Name: sp.Emoji(), PNG: img})
	}
	return out, nil
}

// emojiPNG decodes a webp icon and encodes it as a 128 px PNG with its alpha kept.
func emojiPNG(webpBytes []byte) ([]byte, error) {
	src, err := webp.Decode(bytes.NewReader(webpBytes))
	if err != nil {
		return nil, fmt.Errorf("decode icon: %w", err)
	}
	dst := image.NewNRGBA(image.Rect(0, 0, emojiSide, emojiSide))
	xdraw.CatmullRom.Scale(dst, dst.Bounds(), src, src.Bounds(), xdraw.Src, nil)
	var buf bytes.Buffer
	if err := png.Encode(&buf, dst); err != nil {
		return nil, fmt.Errorf("encode png: %w", err)
	}
	if buf.Len() > emojiMaxBytes {
		return nil, fmt.Errorf("png is %d bytes, over Discord's 256 KiB", buf.Len())
	}
	return buf.Bytes(), nil
}

// SyncReport counts what one sync did.
type SyncReport struct{ Uploaded, Kept, Deleted int }

// SyncEmoji uploads every spec emoji the application lacks; with replace it first deletes the
// spec emoji it holds and uploads them all again. It is safe to re-run: a second run uploads
// nothing. It prints names and counts only, never a credential.
func SyncEmoji(ctx context.Context, api EmojiAPI, images []EmojiImage, replace bool, out io.Writer) (SyncReport, error) {
	var held []AppEmoji
	err := retryRateLimited(ctx, func() (err error) {
		held, err = api.ListAppEmoji(ctx)
		return err
	})
	if err != nil {
		return SyncReport{}, fmt.Errorf("discordbot: list emoji: %w", err)
	}
	have := map[string]string{}
	for _, e := range held {
		have[e.Name] = e.ID
	}
	var rep SyncReport
	for _, img := range images {
		id, exists := have[img.Name]
		if exists && replace {
			if err := retryRateLimited(ctx, func() error { return api.DeleteAppEmoji(ctx, id) }); err != nil {
				return rep, fmt.Errorf("discordbot: delete emoji %s: %w", img.Name, err)
			}
			rep.Deleted++
			exists = false
		}
		if exists {
			rep.Kept++
			fmt.Fprintf(out, "kept     %s\n", img.Name)
			continue
		}
		data := "data:image/png;base64," + base64.StdEncoding.EncodeToString(img.PNG)
		if err := retryRateLimited(ctx, func() error {
			_, err := api.CreateAppEmoji(ctx, img.Name, data)
			return err
		}); err != nil {
			return rep, fmt.Errorf("discordbot: upload emoji %s: %w", img.Name, err)
		}
		rep.Uploaded++
		fmt.Fprintf(out, "uploaded %s\n", img.Name)
	}
	fmt.Fprintf(out, "%d uploaded, %d kept, %d deleted, %d specs\n", rep.Uploaded, rep.Kept, rep.Deleted, len(images))
	return rep, nil
}

// retryRateLimited runs call, and once more after Discord's retry_after when it answers 429.
func retryRateLimited(ctx context.Context, call func() error) error {
	err := call()
	var api *APIError
	if !errors.As(err, &api) || !api.RateLimited() {
		return err
	}
	wait := time.NewTimer(api.RetryAfter)
	defer wait.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-wait.C:
	}
	return call()
}
