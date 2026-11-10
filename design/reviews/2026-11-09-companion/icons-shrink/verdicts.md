# The icon size cut: re-check on the icon sheet only

The owner asked for a smaller icon set. The cut takes it from 2.6 MB to 214 KB:

- every PNG is quantized and then run through oxipng;
- the .icns keeps one entry per pixel size, from 16 to 512;
- the binary embeds only the icons its own platform uses.

Two images were reviewed: `icons.png` (the icons as each OS shows them) and `before-after.png` (the full-colour render against each shipped file, at 100 % and 200 %).

## ux-designer

The seal looks the same at every size each OS shows. Dock and Explorer keep the rim, the rune ring, the letterforms and the marbling. At 200 % there is no banding. The tray tones are untouched. The PSNR runs from 41 to 46 dB.

Polish: the 16 px tray badge is about 4 px wide, as it was before the cut. Check it on a real taskbar at 100 % scale.

VERDICT: SHIP — the quantized icons show no visible change from the full-colour render at any size, and the 214 KB total is a fair trade for the 2.6 MB original.

## ai-slop-reviewer

There is no banding, dither speckle, halo or fringe on the dark or light grounds. The softness at 16 and 24 px is the size, not the compression.

Standing note: if the seal art was generated, the owner should confirm that is acceptable. This change did not cause the note.

VERDICT: PASS — The quantized icons show no visible banding, dither noise, halos or fringes against the full-colour render on dark or light grounds at the sizes I could see.
