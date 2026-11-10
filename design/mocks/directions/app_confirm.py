"""Mock: the app sign-in confirm page (api/internal/auth/app_confirm.go) in round 5 (batch H, lane h-api-html).

The phone's browser opens it when the Forever Sixty app starts a sign-in. Its CSP is strict (default-src 'none';
style-src 'unsafe-inline'): no image, no font, no script is fetched, so the mock draws only what the page can draw.
- No painted art: the band is the flat band colour fading into the flat earth ground (both the art's fallbacks).
- No seal: the wordmark is type alone.
- No web fonts: Marcellus and Open Sans when the device has them installed, otherwise the nearest system faces
  (Optima, then Candara, then Georgia for the title; Segoe UI, Roboto, Helvetica Neue, Arial for the body).
- One framed panel holding the sentence, the teal primary (full width: one tap, the one thing to do) and the
  not-you line.
Copy is read from app_confirm.go (appConfirmCopy), never retyped. Colours are style.py's TOKENS, the same values
design/tokens/tokens.json gives the Go page through api/internal/brand/tokens_gen.go.

Run: python3 design/mocks/directions/app_confirm.py, then
     node design/mocks/directions/shoot.cjs web/node_modules app-confirm
"""
import re
from html import escape

import data as D
from common import write
from style import FRAME_LINE, R_LG, R_SM, TOKENS

SOURCE = (D.ROOT / 'api/internal/auth/app_confirm.go').read_text()
COPY = dict(re.findall(r'^\t(\w+):\s+"([^"]+)",$', SOURCE, re.M))

TITLE_FONT = "'Marcellus',Optima,Candara,Georgia,serif"
BODY_FONT = "'Open Sans','Segoe UI',Roboto,'Helvetica Neue',Arial,sans-serif"

CSS = f'''
:root{{color-scheme:dark}}*{{box-sizing:border-box}}
body{{margin:0;min-height:100vh;background:{TOKENS['earth']};color:{TOKENS['text']};font:16px/1.55 {BODY_FONT};-webkit-font-smoothing:antialiased}}
.band{{background:linear-gradient({TOKENS['band']} 0,{TOKENS['band']} calc(100% - 96px),{TOKENS['earth']} 100%);padding:0 18px 96px}}
.col{{max-width:560px;margin:0 auto}}
.mark{{display:block;padding:18px 0 0;font:400 22px/1 {TITLE_FONT};color:{TOKENS['cream']};letter-spacing:-.015em}}
h1{{margin:0;padding:22px 0 0;font:400 21px/25px {TITLE_FONT};letter-spacing:-.015em;color:{TOKENS['cream']};text-shadow:0 1px 2px rgba(0,0,0,.45)}}
main{{margin-top:-76px;padding:0 18px 40px}}
.panel{{position:relative;background:{TOKENS['card']};border:1px solid #000;box-shadow:0 0 0 1px rgba(130,112,91,.55);border-radius:{R_LG}px;padding:22px 24px}}
.panel::before{{content:"";position:absolute;inset:6px;border:1px solid {FRAME_LINE};border-radius:{R_LG - 4}px;pointer-events:none}}
.panel p{{margin:0}}
.go{{display:flex;align-items:center;justify-content:center;height:44px;margin-top:18px;padding:0 20px;border:1px solid {TOKENS['teal-edge']};
  border-radius:{R_SM}px;background:{TOKENS['teal']};color:#fff;font-weight:600;text-decoration:none}}
.note{{margin-top:16px!important;font-size:14px;line-height:20px;color:rgba(255,255,255,.5)}}
@media(min-width:768px){{.mark{{padding-top:24px}}h1{{padding-top:28px;font-size:26px;line-height:30px}}}}
'''


def app_confirm() -> str:
    body = (f'<div class="band"><div class="col"><span class="mark">Forever Sixty</span><h1>{escape(COPY["Heading"])}</h1></div></div>'
            f'<main><div class="col panel"><p>{escape(COPY["Body"])}</p>'
            f'<a class="go" href="#">{escape(COPY["Button"])}</a><p class="note">{escape(COPY["NotYou"])}</p></div></main>')
    return (f'<!doctype html><html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1">'
            f'<title>{escape(COPY["Title"])}</title><style>{CSS}</style></head><body>{body}</body></html>')


if __name__ == '__main__':
    print(write('app-confirm', app_confirm()))
