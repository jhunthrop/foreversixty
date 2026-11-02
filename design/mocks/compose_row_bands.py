"""Before/after pairs for the row-band sweep (design/specs/2026-11-01-row-bands.md).

Reads the full-page captures web/tests/e2e/row-band-captures.spec.ts writes, and the focus boxes
(<prefix>-<shot>-<width>.focus.json) of the after run, and writes one side-by-side crop per
changed list and width: compare-<shot>-<width>-<list>.png, before on the left.

    python3 design/mocks/compose_row_bands.py design/reviews/2026-11-01-row-bands
"""

import glob
import json
import os
import sys

from PIL import Image, ImageDraw

MARGIN = 28
GAP = 24
MAX_HEIGHT = 1400
LABEL_HEIGHT = 28
GROUND = (5, 6, 9)
INK = (233, 228, 216)


def crop(path, box, width):
    page = Image.open(path).convert("RGB")
    left = max(0, int(box["x"]) - MARGIN)
    top = max(0, int(box["y"]) - MARGIN)
    right = min(page.width, int(box["x"] + box["width"]) + MARGIN, width)
    bottom = min(page.height, int(box["y"] + box["height"]) + MARGIN, top + MAX_HEIGHT)
    return page.crop((left, top, right, bottom))


def pair(before, after):
    width = before.width + GAP + after.width
    height = LABEL_HEIGHT + max(before.height, after.height)
    board = Image.new("RGB", (width, height), GROUND)
    board.paste(before, (0, LABEL_HEIGHT))
    board.paste(after, (before.width + GAP, LABEL_HEIGHT))
    draw = ImageDraw.Draw(board)
    draw.text((8, 8), "BEFORE", fill=INK)
    draw.text((before.width + GAP + 8, 8), "AFTER", fill=INK)
    return board


def main(folder):
    for focus in sorted(glob.glob(os.path.join(folder, "after-*.focus.json"))):
        stem = os.path.basename(focus)[len("after-") : -len(".focus.json")]
        width = int(stem.rsplit("-", 1)[1])
        before_page = os.path.join(folder, f"before-{stem}.png")
        after_page = os.path.join(folder, f"after-{stem}.png")
        for list_id, box in json.load(open(focus)).items():
            board = pair(crop(before_page, box, width), crop(after_page, box, width))
            out = os.path.join(folder, f"compare-{stem}-{list_id}.png")
            board.save(out)
            print(out)


if __name__ == "__main__":
    main(sys.argv[1])
