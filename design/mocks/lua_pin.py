"""Answer the lane generators' reads of the addon's Lua from the store captures' commit.

    import lua_pin  # before any addon_hub_* module

The lane generators read the addon's Lua (Data.lua's packed picks, MeFrame.lua's layout, Locale.lua) through
Path.read_text. That Lua has moved on since the store fixture was set (design/listing/gen_captures.py), so
both the store captures and the brand skin's boards (addon_brand.py, addon_brand_build.py) read it as of
CAPTURE_COMMIT: every board is the same evening of the same made-up character. A reader that wants the
build's own Lua (Theme.SIZES, Hub.LAYOUT) or its textures opens the file directly, which is not pinned.
"""
import pathlib
import subprocess

ROOT = pathlib.Path(__file__).resolve().parents[2]
CAPTURE_COMMIT = 'aa2c8bbb7'
MARKER = 'addon/ForeverSixty/'


def _install():
    original = pathlib.Path.read_text
    if getattr(original, 'lua_pin', False):
        return
    pinned = {}

    def read_text(self, *args, **kwargs):
        path = str(self)
        if MARKER in path and path.endswith('.lua'):
            rel = path[path.index(MARKER):]
            if rel not in pinned:
                pinned[rel] = subprocess.run(['git', '-C', str(ROOT), 'show', f'{CAPTURE_COMMIT}:{rel}'],
                                             check=True, capture_output=True, text=True).stdout
            return pinned[rel]
        return original(self, *args, **kwargs)
    read_text.lua_pin = True
    pathlib.Path.read_text = read_text


_install()
