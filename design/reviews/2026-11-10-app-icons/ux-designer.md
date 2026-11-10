# ux-designer: app icon variants (2026-11-10)

**Verdict: SHIP.** No must-fix items.

Reviewed: `preview.png`, the drawables, `ic_launcher.xml`, the AppIcon set and the generators, at commit 78e72f574.

## Should
1. **LiveNotifier keeps gold.**
   - Live raid and sim notifications use `setColor(FsColor.gold)`, while every other push uses turquoise.
   - Gold should stay only if the progress bar needs it. Otherwise move LiveNotifier to `PushNotifier.builder`.
   - Gold header text on a light shade is about 2.3:1.
   - **Status:** left for the owner, because this lane stays out of LiveNotifier while the Android restyle lands.
2. **iOS Dark filled the tile edge to edge.** It needed a 6 to 8% inset.
   - **Fixed:** the flat seal is now drawn at 85% of the canvas, which is a 7.5% margin.
3. **Tinted could merge with the system background.** Check on a device.
   - **Adjusted:** a levels curve now puts the disc at mid grey (about 120), apart from both the white LX and the black outline.
   - Still needs a device check. The simulator is unavailable on this Mac.
4. **mdpi serifs are about 1 px.** The icon reads as LX but only just. Check once on an mdpi emulator. This is polish.

## Checks passed
- **Notification icon:** `ic_notification_seal` is white on transparent, 22 dp inside the 24 dp grid. FCM's default icon and colour are set in the manifest.
- **Push builders:** both share `PushNotifier.builder`.
- **Themed icon:** the `<monochrome>` layer is 64 dp, inside the 66 dp safe zone. It reads in all eight tints.
- **iOS asset catalog:** `Contents.json` is valid, with the dark and tinted luminosity appearances. Any is unchanged.
- **Palette:** Dark uses only the three brand colours.
- **Build output:** the generated `mipmap-anydpi-v26/ic_launcher.xml` in the build carries the `<monochrome>` line.
