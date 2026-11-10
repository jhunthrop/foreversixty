# ai-slop-reviewer: app icon variants (2026-11-10)

**Verdict: PASS (with fixes, applied).**

No hard tells were found:
- no glow, no decorative gradient, no glass;
- the art is the owner's flat seal, not redrawn.

The reviewer called these authentic:
- one mark across Any, Dark, Tinted, the notification icon and the themed icon;
- the knocked-out disc that stays legible at mdpi 24 px;
- the notification copy, written in a player's voice.

## Findings and fixes
1. **Captions on the sheet over-explained** (tell 22, sheet chrome only). **Fixed:** the captions are now "Any", "Dark", "Tinted (system tint simulated in amber)", "xxhdpi" and "Seal: 64 dp."
2. **Labels were set in system sans** (tell 14, review-aid chrome, not shipped). No change.
3. **Tinted contrast:** the disc (about #B8B8B8) sat too close to the white LX. **Fixed:** a levels curve takes the disc to about #787878. The LX now separates clearly from the disc, and the disc from the outline. The amber in the sheet simulates the system tint; the shipped file is greyscale.
4. **Dark alpha:** the reviewer asked to confirm that the background is transparent, not white. **Verified:** the file is RGBA, and its corners have alpha 0.

Still open: check the Tinted contrast on a device once the iOS simulator or a device is available.
