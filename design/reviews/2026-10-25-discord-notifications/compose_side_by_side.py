"""Item 13: the shipped rail (Plan + Discord panel) at 1440 beside the mock's matching board, same scale.

  python3 design/reviews/2026-10-25-discord-notifications/compose_side_by_side.py

Needs discord-saved-1440.png (capture.spec.ts) and the mock render (design/mocks/gen_discord_notifications.py,
git-ignored). Writes discord-side-by-side-1440.png: mock on the left, shipped on the right, both 1:1.
"""
from pathlib import Path

from PIL import Image, ImageDraw

HERE = Path(__file__).resolve().parent
MOCK = HERE.parents[1] / "mocks" / "renders" / "discord-notifications-states.png"
SHIPPED = HERE / "discord-saved-1440.png"
OUT = HERE / "discord-side-by-side-1440.png"

MOCK_BOX = (96, 178, 510, 836)  # the mock's first board: Plan panel and the saved Discord block
SHIPPED_BOX = (934, 660, 1358, 1342)  # the shipped rail: Plan panel and the Discord panel
GUTTER = 40
LABEL_H = 36
BACKGROUND = (5, 8, 15)
LABEL = (240, 200, 90)


def labelled(image: Image.Image, text: str, height: int) -> Image.Image:
    canvas = Image.new("RGB", (image.width, height + LABEL_H), BACKGROUND)
    ImageDraw.Draw(canvas).text((8, 10), text, fill=LABEL)
    canvas.paste(image, (0, LABEL_H))
    return canvas


def main() -> None:
    mock = Image.open(MOCK).convert("RGB").crop(MOCK_BOX)
    shipped = Image.open(SHIPPED).convert("RGB").crop(SHIPPED_BOX)
    height = max(mock.height, shipped.height)
    left = labelled(mock, "MOCK (design/mocks, 1440 rail)", height)
    right = labelled(shipped, "SHIPPED (fixture build, 1440)", height)
    sheet = Image.new("RGB", (left.width + GUTTER + right.width, height + LABEL_H), BACKGROUND)
    sheet.paste(left, (0, 0))
    sheet.paste(right, (left.width + GUTTER, 0))
    sheet.save(OUT)


if __name__ == "__main__":
    main()
