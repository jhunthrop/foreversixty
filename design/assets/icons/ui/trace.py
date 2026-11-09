"""Trace the generated UI icon sheets into single-path, currentColor SVGs.

Run: python3 -I design/assets/icons/ui/trace.py   (from the repo root)

Sources live in design/assets/icons/ui/source/: the 5x5 sheet (row 3 is a stray glyph and is
skipped; its tier-list and paste-export cells are replaced by the single-icon sources).
Each icon is thresholded on luminance (cream on dark), its contours traced with holes kept
(even-odd), simplified, and written centred in a square viewBox with a fixed margin.
"""

from pathlib import Path

import cv2
import numpy as np

HERE = Path(__file__).resolve().parent
SOURCE = HERE / "source"
THRESHOLD = 200          # cream glyph vs dark ground (and the dim ghosts in tier-list.jpeg)
MERGE_PX = 41            # dilation that joins one icon's pieces into one cluster
MIN_AREA = 400           # specks below this are noise
EPSILON = 1.2            # contour simplification, in source pixels
MARGIN = 0.06            # empty border around each glyph, as a share of its box

SHEET_ORDER = [
    "home", "bis", "guides", None, None,
    "planner", "logs", None, None, "search",
    None,                                   # the stray mountain in row 3
    "account", "premium", "setup", "character", None,
    "sign-out", "manage", "upload", "link", "download",
]
SINGLES = {
    "tier-list": "tier-list.jpeg",
    "paste-export": "paste-export.jpeg",
    "sim": "sim.jpeg",
    "guild": "guild.jpeg",
    "rankings": "rankings.jpeg",
}


def glyph_mask(path: Path) -> np.ndarray:
    gray = cv2.cvtColor(cv2.imread(str(path)), cv2.COLOR_BGR2GRAY)
    return (gray > THRESHOLD).astype(np.uint8) * 255


def clusters(mask: np.ndarray) -> list[tuple[int, int, int, int]]:
    kernel = np.ones((MERGE_PX, MERGE_PX), np.uint8)
    merged = cv2.dilate(mask, kernel)
    found, _ = cv2.findContours(merged, cv2.RETR_EXTERNAL, cv2.CHAIN_APPROX_SIMPLE)
    boxes = [cv2.boundingRect(c) for c in found if cv2.contourArea(c) > MIN_AREA * 10]
    rows: list[list[tuple[int, int, int, int]]] = []
    for box in sorted(boxes, key=lambda b: b[1] + b[3] / 2):
        centre = box[1] + box[3] / 2
        if rows and abs(rows[-1][0][1] + rows[-1][0][3] / 2 - centre) < 120:
            rows[-1].append(box)
        else:
            rows.append([box])
    return [box for row in rows for box in sorted(row, key=lambda b: b[0])]


def to_svg(mask: np.ndarray, box: tuple[int, int, int, int]) -> str:
    x, y, w, h = box
    crop = mask[y:y + h, x:x + w]
    contours, _ = cv2.findContours(crop, cv2.RETR_CCOMP, cv2.CHAIN_APPROX_NONE)
    side = max(w, h)
    pad = round(side * MARGIN)
    view = side + 2 * pad
    dx, dy = pad + (side - w) / 2, pad + (side - h) / 2
    parts = []
    for contour in contours:
        if cv2.contourArea(contour) < MIN_AREA / 4:
            continue
        points = cv2.approxPolyDP(contour, EPSILON, True).reshape(-1, 2)
        coords = " L".join(f"{px + dx:.0f},{py + dy:.0f}" for px, py in points)
        parts.append(f"M{coords}Z")
    path = " ".join(parts)
    return (f'<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 {view} {view}">'
            f'<path fill="currentColor" fill-rule="evenodd" d="{path}"/></svg>\n')


def main() -> None:
    sheet = glyph_mask(SOURCE / "sheet.jpeg")
    boxes = clusters(sheet)
    if len(boxes) != len(SHEET_ORDER):
        raise SystemExit(f"expected {len(SHEET_ORDER)} icons on the sheet, found {len(boxes)}")
    written = []
    for name, box in zip(SHEET_ORDER, boxes):
        if name:
            (HERE / f"{name}.svg").write_text(to_svg(sheet, box))
            written.append(name)
    for name, file in SINGLES.items():
        mask = glyph_mask(SOURCE / file)
        boxes = clusters(mask)
        if not boxes:
            raise SystemExit(f"{file}: no icon found")
        # A single-icon source is one glyph however many pieces it has: take their union.
        x0 = min(b[0] for b in boxes)
        y0 = min(b[1] for b in boxes)
        x1 = max(b[0] + b[2] for b in boxes)
        y1 = max(b[1] + b[3] for b in boxes)
        (HERE / f"{name}.svg").write_text(to_svg(mask, (x0, y0, x1 - x0, y1 - y0)))
        written.append(name)
    print(f"wrote {len(written)} icons: {', '.join(sorted(written))}")


if __name__ == "__main__":
    main()
