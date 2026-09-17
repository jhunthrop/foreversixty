"""Each talent tree's background, taken from the game and made ours.

The client draws a tree's background as four textures under
`Interface\\TALENTFRAME\\`: a 256x256 top-left, a 64x256 top-right, a 256x128
bottom-left and a 64x128 bottom-right, which tile into one 320x384 panel. All
27 tabs of build 1.60.1.69893 have all four. They arrive the same way icons
do -- by file data id from CASC, cached as BLP -- and leave as one processed
WebP per tree at `builds/<build>/trees/<background>.webp`, which is what the
site ships. A raw texture is never published: Blizzard's art at full saturation
under this site's type would read as a screenshot of the game rather than as
this site, and the treatment below is the one place that decision lives.
"""

from __future__ import annotations

import io
import json
from collections.abc import Mapping
from dataclasses import dataclass
from pathlib import Path

import httpx
from PIL import Image, ImageEnhance

from pipeline.icons import CACHE_DIR, _atomic_write
from pipeline.wago import BASE_URL, USER_AGENT


class ArtDataError(ValueError):
    """The client's talent-frame textures are not the shape a panel needs."""


#: Quadrant name -> where its top-left corner goes in the composite.
QUADRANTS: dict[str, tuple[int, int]] = {
    "TopLeft": (0, 0),
    "TopRight": (256, 0),
    "BottomLeft": (0, 256),
    "BottomRight": (256, 256),
}
#: Quadrant name -> the size the client ships it at.
QUADRANT_SIZES: dict[str, tuple[int, int]] = {
    "TopLeft": (256, 256),
    "TopRight": (64, 256),
    "BottomLeft": (256, 128),
    "BottomRight": (64, 128),
}
BACKGROUND_SIZE = (320, 384)
_TALENT_FRAME_DIR = "interface\\talentframe\\"
_BLP_SUFFIX = ".blp"
#: Quadrant suffix, any case the client ships it in, -> the canonical QUADRANTS
#: key. 1.60.1.69893's own ManifestInterfaceData mixes cases within one build:
#: most tabs are "DruidBalance-BottomLeft.blp" but some (all of Paladin's) are
#: "PALADINCOMBAT-BOTTOMLEFT.BLP".
_QUADRANT_BY_CASE: dict[str, str] = {name.lower(): name for name in QUADRANTS}


@dataclass(frozen=True)
class BackgroundTreatment:
    """How a talent-frame texture becomes a Forever Sixty tree background.

    Applied in this order: pull the colour out, take the light down, then
    blend toward one palette token so every tree sits on the same ground and
    the gold borders and white type stay the brightest things on the panel.
    """

    #: 0 leaves a greyscale image, 1 leaves the texture's own colour.
    saturation: float
    #: Multiplies the remaining light. 1 leaves it alone.
    brightness: float
    #: The colour everything is blended toward: --color-raised in
    #: web/src/styles/tokens.css, the panel the tree grid sits on.
    tint: tuple[int, int, int]
    #: How far toward `tint`. 0 leaves the darkened texture, 1 leaves a flat fill.
    tint_strength: float


TREATMENT = BackgroundTreatment(
    saturation=0.30,
    brightness=0.60,
    tint=(0x0D, 0x11, 0x1A),
    tint_strength=0.35,
)


def process_pixel(
    rgb: tuple[int, int, int], treatment: BackgroundTreatment = TREATMENT
) -> tuple[int, int, int]:
    """What `process_background` does to one pixel. The treatment's definition."""
    red, green, blue = rgb
    # ITU-R 601-2 luma, the transform Image.convert("L") uses.
    grey = red * 299 / 1000 + green * 587 / 1000 + blue * 114 / 1000
    out = []
    for channel, tint in zip(rgb, treatment.tint, strict=True):
        desaturated = grey + (channel - grey) * treatment.saturation
        darkened = desaturated * treatment.brightness
        out.append(round(darkened + (tint - darkened) * treatment.tint_strength))
    return (out[0], out[1], out[2])


def compose_background(quadrants: Mapping[str, bytes]) -> Image.Image:
    """The four textures tiled into one 320x384 RGB image.

    Alpha is dropped rather than composited: three of the four quadrants carry
    an alpha channel, the panel behind them is opaque in game, and pasting
    with a mask would leave the site's page colour showing through the edges
    of every tree.
    """
    panel = Image.new("RGB", BACKGROUND_SIZE, (0, 0, 0))
    for name, offset in QUADRANTS.items():
        data = quadrants.get(name)
        if data is None:
            raise ArtDataError(f"no {name} texture; a tree background needs all four")
        with Image.open(io.BytesIO(data)) as tile:
            if tile.size != QUADRANT_SIZES[name]:
                raise ArtDataError(
                    f"{name} is {tile.size}, not the {QUADRANT_SIZES[name]} it tiles at"
                )
            panel.paste(tile.convert("RGB"), offset)
    return panel


def process_background(
    image: Image.Image, treatment: BackgroundTreatment = TREATMENT
) -> Image.Image:
    """`process_pixel` over a whole image, through Pillow's own C loops."""
    rgb = image.convert("RGB")
    desaturated = ImageEnhance.Color(rgb).enhance(treatment.saturation)
    darkened = ImageEnhance.Brightness(desaturated).enhance(treatment.brightness)
    tint = Image.new("RGB", darkened.size, treatment.tint)
    return Image.blend(darkened, tint, treatment.tint_strength)


def background_webp(
    quadrants: Mapping[str, bytes], treatment: BackgroundTreatment = TREATMENT
) -> bytes:
    processed = process_background(compose_background(quadrants), treatment)
    buffer = io.BytesIO()
    processed.save(buffer, "WEBP", quality=90, method=6)
    return buffer.getvalue()


def talent_frame_ids(manifest_rows: list[dict[str, str]]) -> dict[str, dict[str, int]]:
    """Background name (lowercased) -> quadrant name -> file data id."""
    out: dict[str, dict[str, int]] = {}
    for row in manifest_rows:
        if row["FilePath"].lower() != _TALENT_FRAME_DIR:
            continue
        name = row["FileName"]
        if not name.lower().endswith(_BLP_SUFFIX):
            continue
        stem = name[: -len(_BLP_SUFFIX)]
        background, _, quadrant = stem.rpartition("-")
        canonical = _QUADRANT_BY_CASE.get(quadrant.lower())
        if canonical is None:
            continue
        out.setdefault(background.lower(), {})[canonical] = int(row["ID"])
    return out


def _wanted_backgrounds(build_dir: Path) -> set[str]:
    """Every `background` the build's emitted trees name."""
    names: set[str] = set()
    for path in sorted((build_dir / "talents").glob("*.json")):
        payload = json.loads(path.read_text(encoding="utf-8"))
        for tree in payload["trees"]:
            names.add(tree["background"])
    return names


def backgrounds_for_build(
    build: str,
    root: Path = Path("builds"),
    cache_dir: Path = CACHE_DIR,
    client: httpx.Client | None = None,
) -> int:
    """Write one processed background per tree. Returns the number written."""
    from pipeline.csvio import read_csv

    build_dir = root / build
    raw = build_dir / "raw"
    if not raw.exists():
        raise SystemExit(f"no raw data at {raw}; run `python -m pipeline fetch` first")
    frames = talent_frame_ids(read_csv(raw / "ManifestInterfaceData.csv"))
    wanted = sorted(_wanted_backgrounds(build_dir))
    out_dir = build_dir / "trees"
    out_dir.mkdir(parents=True, exist_ok=True)
    cache_dir.mkdir(parents=True, exist_ok=True)
    own = client is None
    if client is None:
        client = httpx.Client(base_url=BASE_URL, headers={"User-Agent": USER_AGENT})
    written = 0
    try:
        for background in wanted:
            target = out_dir / f"{background}.webp"
            if target.exists():
                continue
            ids = frames.get(background)
            if not ids or set(ids) != set(QUADRANTS):
                raise ArtDataError(
                    f"{background} has {sorted(ids or ())} in ManifestInterfaceData, "
                    f"not the four quadrants {sorted(QUADRANTS)}"
                )
            quadrants: dict[str, bytes] = {}
            for quadrant, file_id in ids.items():
                cached = cache_dir / f"{file_id}.blp"
                if not cached.exists():
                    response = client.get(f"/api/casc/{file_id}", timeout=60)
                    response.raise_for_status()
                    _atomic_write(cached, response.content)
                quadrants[quadrant] = cached.read_bytes()
            _atomic_write(target, background_webp(quadrants))
            written += 1
    finally:
        if own:
            client.close()
    print(f"{written} tree backgrounds written, {len(wanted)} referenced")
    return written
