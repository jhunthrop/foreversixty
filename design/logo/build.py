"""Forever Sixty logo: the LX seal (a level badge with sixty as the Roman numeral) as
vector mark, mono marks, wordmark and lockups, plus the favicon and the PNG exports.
Glyphs are the real Cinzel outlines (800 for LX, 700 for the wordmark) from the web
package's fontsource files. Run from the repo root: python3 design/logo/build.py"""
from fontTools.ttLib import TTFont
from fontTools.pens.svgPathPen import SVGPathPen
from fontTools.pens.transformPen import TransformPen
import math, shutil, subprocess
from pathlib import Path
ROOT=Path(__file__).resolve().parents[2]
W=str(ROOT/'web/node_modules/@fontsource/cinzel/files')+'/'
OUT=Path(__file__).resolve().parent; EXPORT=OUT/'export'; EXPORT.mkdir(exist_ok=True)
def text_path(text, weight, size, x, y, tracking=0.0):
    f=TTFont(W+f'cinzel-latin-{weight}-normal.woff2'); gs=f.getGlyphSet(); cmap=f.getBestCmap(); upm=f['head'].unitsPerEm
    s=size/upm; parts=[]; adv=0; kern=None
    for ch in text:
        if ch==' ': adv+=upm*0.28; continue
        g=cmap[ord(ch)]; pen=SVGPathPen(gs); tp=TransformPen(pen,(s,0,0,-s,x+adv*s,y)); gs[g].draw(tp)
        parts.append(pen.getCommands()); adv+=gs[g].width+tracking*upm
    return ' '.join(parts), adv*s
GRAD='<linearGradient id="g" x1="0" y1="0" x2="0" y2="1"><stop offset="0" stop-color="#fbe7a1"/><stop offset=".45" stop-color="#e5b955"/><stop offset="1" stop-color="#a8762a"/></linearGradient>'
def seal(size=256, ink=None, dark=True, ground=None):
    c=size/2; r=c*0.88; sw=size*0.05
    d,adv=text_path('LX',800,size*0.44,0,0,tracking=-0.02)
    tx=c-adv/2; ty=c+size*0.44*0.36
    fill = ink or 'url(#g)'
    dots=[]; n=36; rr=r-size*0.10
    for i in range(n):
        a=2*math.pi*i/n; dots.append(f'<circle cx="{c+rr*math.cos(a):.2f}" cy="{c+rr*math.sin(a):.2f}" r="{size*0.008:.2f}" fill="{ink or '#a8762a'}"/>')
    disc = f'<circle cx="{c}" cy="{c}" r="{r}" fill="{ground}"/>' if ground else (f'<circle cx="{c}" cy="{c}" r="{r}" fill="#0d111a"/>' if dark else '')
    defs = '' if ink else f'<defs>{GRAD}</defs>'
    return (f'<svg xmlns="http://www.w3.org/2000/svg" width="{size}" height="{size}" viewBox="0 0 {size} {size}">{defs}'
            f'{disc}<circle cx="{c}" cy="{c}" r="{r}" fill="none" stroke="{fill}" stroke-width="{sw}"/>{"".join(dots)}'
            f'<path transform="translate({tx:.2f},{ty:.2f})" d="{d}" fill="{fill}"/></svg>')
def wordmark(ink='#f2eee4', size=64):
    d,adv=text_path('FOREVER SIXTY',700,size,0,size*0.78,tracking=0.10)
    w=adv+size*0.1
    return f'<svg xmlns="http://www.w3.org/2000/svg" width="{w:.0f}" height="{size}" viewBox="0 0 {w:.0f} {size}"><path d="{d}" fill="{ink}"/></svg>', w
def lockup(horizontal=True, ink_word='#f2eee4', mono=None):
    m=seal(64, ink=mono, dark=mono is None)
    inner=m.split('>',1)[1].rsplit('</svg>',1)[0]
    if mono is None: inner=inner.replace('<defs>','').replace('</defs>','')
    wm,w=wordmark(mono or ink_word, 40)
    winner=wm.split('>',1)[1].rsplit('</svg>',1)[0]
    if horizontal:
        W_=64+18+w; H=64
        return f'<svg xmlns="http://www.w3.org/2000/svg" width="{W_:.0f}" height="{H}" viewBox="0 0 {W_:.0f} {H}"><defs>{GRAD}</defs>{inner}<g transform="translate(82,12)">{winner}</g></svg>'
    W_=max(96,w); H=96+16+40
    return f'<svg xmlns="http://www.w3.org/2000/svg" width="{W_:.0f}" height="{H}" viewBox="0 0 {W_:.0f} {H}"><defs>{GRAD}</defs><g transform="translate({(W_-64)/2:.1f},0)">{inner}</g><g transform="translate({(W_-w)/2:.1f},80)">{winner}</g></svg>'
(OUT/'foreversixty-mark.svg').write_text(seal(256))
(OUT/'foreversixty-mark-mono.svg').write_text(seal(256, ink='#07090d', dark=False))
(OUT/'foreversixty-mark-mono-on-dark.svg').write_text(seal(256, ink='#f2eee4', dark=False))
(OUT/'foreversixty-wordmark-mono.svg').write_text(wordmark()[0])
(OUT/'foreversixty-lockup-horizontal.svg').write_text(lockup(True))
(OUT/'foreversixty-lockup-stacked.svg').write_text(lockup(False))
(OUT/'favicon.svg').write_text(seal(32))
(OUT/'foreversixty-wordmark.svg').write_text(wordmark('#e5b955')[0])
rsvg=shutil.which('rsvg-convert')
if rsvg:
    for name,svg,w in (('foreversixty-icon-1024.png','foreversixty-mark.svg',1024),('foreversixty-icon-400.png','foreversixty-mark.svg',400),('foreversixty-lockup-1200.png','foreversixty-lockup-horizontal.svg',1200),('foreversixty-icon-64.png','foreversixty-mark.svg',64)):
        subprocess.run([rsvg,'-w',str(w),str(OUT/svg),'-o',str(EXPORT/name)],check=True)
print('logo built')
