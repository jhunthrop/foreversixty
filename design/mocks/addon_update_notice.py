"""Review board for the 2.0.56 update notice: the one chat line a session when a guildmate's addon is newer.

    python3 design/mocks/addon_update_notice.py

Writes design/reviews/2026-11-15-addon-followups/update-notice.png at device scale 2: the default chat frame's
General tab with guild chat around the line, the line in print()'s white with the addon's gold
"Forever Sixty:" prefix (Theme.chatLine), its words Locale.lua's commsUpdateAvailable. Every name is made up.
"""
import os
import re

from addon_hub import esc
from addon_ux import ADDON, ROOT

OUT = ROOT / 'design/reviews/2026-11-15-addon-followups'
LOCALE = (ROOT / 'addon/ForeverSixty/Locale.lua').read_text()
LINE = re.search(r'commsUpdateAvailable = "([^"]+)"', LOCALE).group(1).replace('%s', '2.0.57', 1).replace('%s', '2.0.56', 1)
GUILD = '#40ff40'  # the client's guild chat green
SYSTEM = '#ffffff'  # print() to the default chat frame: white

CSS = '''
body { margin: 0; background: #2a2f25; }
#board { width: 620px; padding: 14px; font-family: "Arial Narrow", Arial, sans-serif; font-size: 15px;
  line-height: 19px; text-shadow: 1px 1px 0 #000; background: linear-gradient(#3b4a35, #23291f); }
.chat { background: rgba(0, 0, 0, 0.42); padding: 8px 10px; }
.tab { display: inline-block; padding: 3px 10px; margin-bottom: 4px; color: #ffd100; font-size: 13px;
  background: rgba(0, 0, 0, 0.5); }
'''


def board():
    rows = [
        (GUILD, '[Guild] [Brannoc]: summons up at the Onyxia entrance in 10'),
        (GUILD, '[Guild] [Mirelle]: omw'),
        (SYSTEM, f'<span style="color:{ADDON["gold"]}">Forever Sixty</span>: {esc(LINE)}'),
        (GUILD, '[Guild] [Kethra]: anyone got a spare fire prot potion?'),
    ]
    lines = ''.join(f'<div style="color:{c}">{t}</div>' for c, t in rows)
    return (f'<!doctype html><html><head><meta charset="utf-8"><style>{CSS}</style></head><body>'
            f'<div id="board"><div class="tab">General</div><div class="chat">{lines}</div></div></body></html>')


def main():
    from playwright.sync_api import sync_playwright
    OUT.mkdir(parents=True, exist_ok=True)
    with sync_playwright() as pw:
        exe = os.environ.get('PLAYWRIGHT_CHROMIUM')
        browser = pw.chromium.launch(executable_path=exe) if exe else pw.chromium.launch()
        page = browser.new_context(device_scale_factor=2).new_page()
        page.set_content(board())
        page.locator('#board').screenshot(path=str(OUT / 'update-notice.png'))
        browser.close()
    print(OUT / 'update-notice.png')


if __name__ == '__main__':
    main()
