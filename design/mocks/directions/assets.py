"""Finding the owner's assets (design/assets: the seals, the art and the nav icons, committed); a missing asset stops
the build."""
import base64
from functools import lru_cache
from pathlib import Path

ROOT = Path(__file__).resolve().parents[3]
ASSETS = ROOT / 'design/assets'


def owner_asset(rel: str) -> Path:
    path = ASSETS / rel
    if not path.exists():
        raise FileNotFoundError(f'design/assets/{rel} not found')
    return path


@lru_cache(maxsize=None)
def png_uri(rel: str) -> str:
    return 'data:image/png;base64,' + base64.b64encode(owner_asset(rel).read_bytes()).decode()


def logo_uri(px: int) -> str:
    """The painted seal, for large sizes."""
    return png_uri(f'logo/lx-seal-{px}.png')


def flat_logo_uri(px: int) -> str:
    """The flat seal, for the header (48 px and under) and favicons."""
    return png_uri(f'logo/lx-seal-flat-{px}.png')
