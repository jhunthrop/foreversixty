"""Data-shaped markup shared by the round-3 pages: feed sentences, the paperdoll order, the BiS cells, writing a mock.

Style lives in style.py; this module knows nothing about how things look.
"""
from html import escape
from pathlib import Path

import data as D

OUT = Path(__file__).resolve().parent / 'html'

# The client's character-sheet order: left column, right column, weapons along the bottom (PaperDollFrame).
DOLL_LEFT = ['head', 'neck', 'shoulder', 'back', 'chest', 'wrist']
DOLL_RIGHT = ['hands', 'waist', 'legs', 'feet', 'finger1', 'finger2', 'trinket1', 'trinket2']
DOLL_BOTTOM = ['main_hand', 'off_hand', 'ranged']


def write(name: str, html: str) -> Path:
    OUT.mkdir(parents=True, exist_ok=True)
    path = OUT / f'{name}.html'
    path.write_text(html)
    return path


def sentence(row: dict) -> str:
    """A feed row's sentence: character names in their class colour, strong segments bold."""
    out = []
    for seg in row['segments']:
        if seg.get('style') == 'name':
            out.append(f'<b style="color:{D.CLASS_COLOR[seg["class"]]}">{escape(seg["text"])}</b>')
        elif seg.get('style') == 'strong':
            out.append(f'<b>{escape(seg["text"])}</b>')
        else:
            out.append(escape(seg['text']))
    return ''.join(out)


def doll_cells(facts: dict) -> dict[str, dict]:
    """The BiS facts' cells keyed by slot, whichever view (by slot or where to go) the facts were read in."""
    doll = facts['doll']
    cells = doll['left'] + doll['right'] + [c for g in doll['groups'] for c in g['cells']]
    return {c['slot']: c for c in cells}


def next_best(cell: dict) -> str:
    """The basis as a player says it: "Replaces <item>." or "Next best: <item>, 0.8 DPS less."."""
    b = cell.get('basis') or {}
    prefix, other, quality = b.get('prefix', ''), b.get('item', ''), b.get('quality', 'common')
    if not other:
        return ''
    named = f'<span style="color:var(--q-{quality},{D.QUALITY[quality]})">{escape(other)}</span>'  # the page's lifted colour where set
    if prefix.startswith('next:'):
        delta = prefix.removeprefix('next:').removesuffix('vs').strip().lstrip('−-')
        return f'Next best: {named}, {delta} DPS less.'
    return f'Replaces {named}.'

