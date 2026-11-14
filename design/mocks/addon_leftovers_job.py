"""Review board for the Guild tab's "Your job" ink (2.0.57): the job's name moves from the client yellow to the
brand's white (Theme strong, base.strong), per design/specs/2026-11-14-addon-brand.md (gold only for the waiting
dot and the score).

    python3 design/mocks/addon_leftovers_job.py

Writes design/reviews/2026-11-15-addon-leftovers/window-guild-before.png and window-guild.png: the store
fixture's Guild window (design/listing/gen_captures.py, read only, its own captures untouched) as shipped, and
with surfaces/GuildHubTab.lua's one change applied to the board (the job span's class cg -> s).
"""
import sys
from pathlib import Path

HERE = Path(__file__).resolve().parent
ROOT = HERE.parents[1]
sys.path.insert(0, str(ROOT / 'design/listing'))
sys.path.insert(0, str(HERE))

import gen_captures as gc  # noqa: E402

OUT = ROOT / 'design/reviews/2026-11-15-addon-leftovers'
JOB_BEFORE = '<span class="cg">Phase 2 ranged, west side</span>'
JOB_AFTER = '<span class="s">Phase 2 ranged, west side</span>'


def main():
    build, selector, extra, most = gc.CAPTURES['window-guild']

    def after():
        html = build()
        if JOB_BEFORE not in html:
            raise SystemExit('the board no longer draws the job span this board recolours')
        return html.replace(JOB_BEFORE, JOB_AFTER)

    gc.OUT = OUT
    gc.CAPTURES = {'window-guild-before': (build, selector, extra, most),
                   'window-guild': (after, selector, extra, most)}
    gc.main(list(gc.CAPTURES))


if __name__ == '__main__':
    main()
