import struct
from pathlib import Path

from PIL import Image

from pipeline.csvio import read_csv
from pipeline.icons import blp_to_webp, icon_names

HERE = Path(__file__).parent


def make_blp2(
    width: int = 8,
    height: int = 8,
    colour: tuple[int, int, int, int] = (13, 34, 56, 255),
) -> bytes:
    """A minimal palettised BLP2, the format wago.tools serves from CASC."""
    red, green, blue, alpha = colour
    palette = struct.pack("<4B", blue, green, red, 255) + bytes(255 * 4)
    pixels = bytes(width * height)  # every pixel is palette index 0
    alpha_channel = bytes([alpha]) * (width * height)
    body = pixels + alpha_channel
    header = b"BLP2" + struct.pack("<i", 1) + struct.pack("<4b", 1, 8, 0, 1)
    header += struct.pack("<II", width, height)
    header += struct.pack("<16I", 148 + 1024, *([0] * 15))
    header += struct.pack("<16I", len(body), *([0] * 15))
    return header + palette + body


def test_icon_names_lowercase_and_drop_the_extension():
    names = icon_names(read_csv(HERE / "fixtures/ManifestInterfaceData.csv"))
    assert names[132154] == "ability_golemthunderclap"
    assert names[132759] == "inv_chest_samurai"


def test_icon_names_skip_non_blp_entries():
    names = icon_names([{"ID": "1", "FilePath": "Interface\\FrameXML\\", "FileName": "Fonts.xml"}])
    assert names == {}


def test_blp_to_webp_produces_a_64_square(tmp_path: Path):
    data = blp_to_webp(make_blp2())
    out = tmp_path / "icon.webp"
    out.write_bytes(data)
    with Image.open(out) as img:
        assert img.format == "WEBP"
        assert img.size == (64, 64)
        assert img.convert("RGBA").getpixel((32, 32)) == (13, 34, 56, 255)


def test_blp_to_webp_is_deterministic():
    blp = make_blp2()
    assert blp_to_webp(blp) == blp_to_webp(blp)
