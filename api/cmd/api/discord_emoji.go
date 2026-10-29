// api/cmd/api/discord_emoji.go
//
// discord-emoji uploads the raid roster's spec marks (spec 2026-10-28 §P.9.1) as application
// emoji: one 128 px PNG per specs.json row, drawn from the active build's icon folder; and the
// six item-quality dots the gear panel leads each name with (spec 2026-10-29 §A.4.1,
// fs_q_poor to fs_q_legendary), drawn in Go in the site's quality colours. It uploads
// only the names the application lacks, so it is safe to re-run; --replace deletes and uploads
// them all again after an icon changes. It prints names and counts, never the token. Run it from
// a checkout (the icons are generated web data, not in the image), once per application.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/discordbot"
	"github.com/jhunthrop/foreversixty/api/internal/sheet"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// defaultDataRoot is web/public/data seen from the api directory, where the command is run.
const defaultDataRoot = "../web/public/data"

// defaultActiveBuildFile is web/src/data/active-build.json seen from the api directory: the
// file tools/activate-build.sh writes, which names the build the site serves. A build fetched
// into web/public/data but not activated is never the icon source.
const defaultActiveBuildFile = "../web/src/data/" + trees.ActiveBuildFile

// emojiAPIFactory builds the application-emoji client; tests swap it for a fake.
type emojiAPIFactory func(applicationID, botToken string) discordbot.EmojiAPI

func productionEmojiAPI(applicationID, botToken string) discordbot.EmojiAPI {
	return discordbot.NewHTTPREST(applicationID, botToken)
}

// discordEmojiCommand is the setup command's name, a row of extraJobs.
const discordEmojiCommand = "discord-emoji"

// runDiscordEmojiCommand is runDiscordEmoji over the process's environment and Discord.
func runDiscordEmojiCommand(ctx context.Context, _ *slog.Logger, args []string) error {
	return runDiscordEmoji(ctx, args, os.Getenv, productionEmojiAPI, os.Stdout)
}

func runDiscordEmoji(ctx context.Context, args []string, getenv func(string) string, newAPI emojiAPIFactory, out io.Writer) error {
	fs := flag.NewFlagSet("discord-emoji", flag.ContinueOnError)
	fs.SetOutput(out)
	icons := fs.String("icons", "", "the icon folder (default: the active build under "+defaultDataRoot+")")
	replace := fs.Bool("replace", false, "delete and upload every spec emoji again")
	if err := fs.Parse(args); err != nil {
		return err
	}
	appID, token := strings.TrimSpace(getenv("DISCORD_APPLICATION_ID")), strings.TrimSpace(getenv("DISCORD_BOT_TOKEN"))
	if appID == "" || token == "" {
		return errors.New("discord-emoji: DISCORD_APPLICATION_ID and DISCORD_BOT_TOKEN are required")
	}
	dir := *icons
	if dir == "" {
		build, err := trees.ReadActiveBuild(defaultActiveBuildFile)
		if err != nil {
			return err
		}
		dir = filepath.Join(defaultDataRoot, build, "icons")
	}
	fmt.Fprintf(out, "icons from %s\n", dir)
	images, err := emojiImages(dir)
	if err != nil {
		return err
	}
	_, err = discordbot.SyncEmoji(ctx, newAPI(appID, token), images, *replace, out)
	return err
}

// emojiImages is every emoji the application holds: the spec marks, then the quality dots.
func emojiImages(iconsDir string) ([]discordbot.EmojiImage, error) {
	images, err := discordbot.EmojiImages(iconsDir)
	if err != nil {
		return nil, err
	}
	dots, err := sheet.QualityDots()
	if err != nil {
		return nil, err
	}
	for _, d := range dots {
		images = append(images, discordbot.EmojiImage{Name: d.Name, PNG: d.PNG})
	}
	return images, nil
}
