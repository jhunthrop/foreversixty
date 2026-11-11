"""The addon's palette and faces for the mock boards, read from design/tokens/tokens.json exactly as
design/tokens/build.mjs writes addon/ForeverSixty/Tokens.lua: the web value of each colour
(platformValue.web where the site leads), references resolved. So a board draws the colours the
addon draws, from the same one source, and never a hex of its own.

    from addon_tokens import color, FONT_FILES
    color('base.bg')        # '#100a06'
"""
import json
from pathlib import Path

ROOT = Path(__file__).resolve().parents[2]
_DOC = json.loads((ROOT / 'design/tokens/tokens.json').read_text())
EXT = 'gg.foreversixty'


def _node(path: str) -> dict:
    node = _DOC
    for key in path.split('.'):
        node = node[key]
    return node


def _value(path: str):
    node = _node(path)
    web = node.get('$extensions', {}).get(EXT, {}).get('platformValue', {}).get('web')
    value = web if web is not None and not isinstance(web, dict) else node['$value']
    if isinstance(value, str) and value.startswith('{') and value.endswith('}'):
        return _value(value[1:-1])
    return value


def color(path: str) -> str:
    """A colour token as CSS: '#rrggbb', or rgba() where the token carries an alpha."""
    path = path if path.startswith('color.') else f'color.{path}'
    return _value(path)


def _faces(font: str) -> dict:
    faces = _node(f'font.{font}')['$extensions'][EXT]['faces']
    return {int(w): ROOT / 'addon/ForeverSixty/media/fonts' / f'{ps}.ttf' for w, ps in faces.items()}


#: The faces the addon ships (media/fonts), by family and weight.
FONT_FILES = {'body': _faces('body'), 'title': _faces('title')}


def role(name: str) -> dict:
    """An addon type role (type.addon.*): font family key, weight, size and line in px."""
    v = _node(f'type.addon.{name}')['$value']
    return {'font': v['fontFamily'][len('{font.'):-1], 'weight': v['fontWeight'],
            'size': int(v['fontSize'][:-2]), 'line': int(v['lineHeight'][:-2])}
