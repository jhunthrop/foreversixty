"""Mock boards for the guild header art round (design/specs/2026-10-04-guild-page.md §12.2,
round 2). Scope: the header band only -- faction emblem as a crest-ring identity mark
(unchanged from round 1 of this header study) plus a composed faction banner or watermark as
the band's own art, over the existing night gradient. The roster/tabs/content below the band
are reused unchanged from gen_guild.py so each board still reads as a real page.

  python3 design/mocks/gen_guild_header.py
    -> guild-header-horde-banner.html     (1440, officer, option A: hanging cloth banner)
    -> guild-header-horde-watermark.html  (1440, officer, option B: large emblem watermark
                                            + diagonal vignette)
    -> guild-header-alliance.html         (1440, officer, option A -- the proposed pick)
    -> guild-header-2000.html             (2000 wide, member, option A -- confirms the banner
                                            stays anchored to the 1344px inner, not the
                                            viewport edge)
    -> guild-header-phone.html            (390, member, option A -- banner shrinks beside the
                                            title, 44px ringed crest unchanged)

Round-2 fix (owner review of round 1's boards): the tab strip sat flush against the viewport
edge at 2000px while the header text above it stayed centred at 1344px -- `tab_strip()`'s own
container carries a hardcoded `padding:0 48px`, which only happens to equal the 1344px
column's own left inset at exactly 1440px board width ((1440-1344)/2 = 48). Fixed here by
nesting the tab strip inside the SAME 1344px-max-width column as the eyebrow/h1/officer strip,
on every board this file builds -- the rule is now stated once in §12.2 ("the tab strip never
leaves the inner") so the build lane does not reproduce the live page's own correct nesting
incorrectly in a future mock.

Round-2 fix (owner review of round 1's art): the blurred-emblem "smear" is gone. Two options
instead, both built from nothing but the two real emblem files
(web/public/icons/hd/faction/{alliance,horde}.webp) plus CSS/SVG gradients in the faction's own
documented colour -- no Blizzard artwork beyond the emblem pixels themselves, no new raster
asset. See each builder's own docstring for exactly how.

The neutral/unknown-faction fallback is not re-rendered here either: design/mocks/renders/
guild-overview-officer.png (on file) remains its reference.
"""
from mocklib import GOLD, MUTED, TEXT, BODY, emblem, nav, nav_phone, footer, page, write
from gen_guild import (
    GUILD, RED, BLUE, UPDATED_LABEL, VIEWER_MEMBER, VIEWER_OFFICER,
    ROSTER_BY_NAME, header_facts, standing_line, officer_tools, tab_strip,
    overview_raids_summary, overview_progression_summary, mock_caption,
    QUIET_LINK_STYLE,
)

# Faction colour tokens (design/DESIGN-SYSTEM.md "Faction"): the ring/banner/glow use the bar
# variant (full-saturation swatch), matching ClassCrest.astro's own choice of classes.json's
# full-saturation `color` field for its ring, never the lightened text variant.
FACTION_BAR = {'alliance': '#2f6fd6', 'horde': '#c0392b'}
FACTION_TEXT = {'alliance': BLUE, 'horde': RED}
# Deep cloth gradient stops, option A (owner's own spec: "Horde: dark crimson to near-black;
# Alliance: deep royal blue to near-black").
BANNER_STOPS = {
    'alliance': ('#1d4d8f', '#0e2342', '#080e18'),
    'horde': ('#7a1012', '#300506', '#0d0302'),
}


def faction_crest(faction: str, size: int) -> str:
    """The identity mark (unchanged from round 1 of this study): ClassCrest.astro's own
    shipped ring recipe (circular crop, 2px box-shadow ring, --raised background fallback)
    applied to FactionMark's emblem asset instead of a class icon. A new mark, not a
    FactionMark.astro edit -- that component's own header comment states 'no ring, no
    background, unlike ClassCrest' for its nine existing callers."""
    ring = FACTION_BAR[faction]
    pad = round(size * 0.16)
    return (f'<img src="{emblem(faction)}" alt="" style="width:{size}px;height:{size}px;'
            f'border-radius:999px;object-fit:contain;padding:{pad}px;box-sizing:border-box;'
            f'box-shadow:0 0 0 2px {ring};background:var(--raised);flex-shrink:0">')


def art_banner(faction: str, phone: bool = False) -> str:
    """Option A: a composed hanging banner, drawn by us, never Blizzard art beyond the real
    emblem's own pixels. Built from three layers:
      1. A soft, blurred faction-colour ambient glow behind the cloth (a plain CSS radial-
         gradient in the faction's bar colour -- no image involved at all).
      2. The cloth itself: one inline SVG <path>, a pointed-bottom banner silhouette, filled
         with a 3-stop linear gradient (deep faction colour -> near-black, the owner's own
         stops) and outlined with a 2px gold stroke (`--gold`, `#e5b955`) -- SVG stroke
         follows the path's own diagonal edges correctly, which a CSS border on a
         clip-path'd div cannot do. `filter:drop-shadow(...)` on the SVG casts a shadow that
         follows the cloth's own silhouette, not a rectangle, so it reads as hanging in
         front of the band's existing night gradient.
      3. The real HD emblem, crisp, centred on the cloth -- the one piece of real game art in
         this layer, large (160px desktop / 70px phone) and unblurred, exactly the owner's
         own ask ("the HD emblem crisp and large").
    The cloth's own bottom 25% fades to transparent (a mask-image on the SVG) so its point
    dissolves into the band's base rather than hard-colliding with the officer strip/tab
    strip beneath it -- verified clear on every board this file renders.
    """
    top_stop, mid_stop, bottom_stop = BANNER_STOPS[faction]
    gid = f'bannerFill-{faction}-{"p" if phone else "d"}'
    w, h = (64, 100) if phone else (150, 230)
    point = h  # pointed-bottom apex sits at the shape's own full height
    shoulder = round(h * 0.82)
    path = f'M0,0 L{w},0 L{w},{shoulder} L{round(w/2)},{point} L0,{shoulder} Z'
    emblem_size = 70 if phone else 160
    emblem_top = round(h * 0.16)
    glow_size = round(w * 3.6)
    mask = 'linear-gradient(180deg,#000 0%,#000 70%,transparent 100%)'
    svg = (
        f'<svg width="{w}" height="{h}" viewBox="0 0 {w} {h}" '
        f'style="position:absolute;top:0;left:0;filter:drop-shadow(0 10px 18px rgba(0,0,0,.55));'
        f'-webkit-mask-image:{mask};mask-image:{mask}">'
        f'<defs><linearGradient id="{gid}" x1="0" y1="0" x2="0" y2="1">'
        f'<stop offset="0%" stop-color="{top_stop}"/>'
        f'<stop offset="55%" stop-color="{mid_stop}"/>'
        f'<stop offset="100%" stop-color="{bottom_stop}"/></linearGradient></defs>'
        f'<path d="{path}" fill="url(#{gid})" stroke="{GOLD}" stroke-width="2" stroke-linejoin="round"/>'
        f'</svg>'
    )
    glow = (f'<div style="position:absolute;top:{-round(glow_size*0.22)}px;left:{round(w/2-glow_size/2)}px;'
            f'width:{glow_size}px;height:{glow_size}px;border-radius:50%;'
            f'background:radial-gradient(circle,{FACTION_BAR[faction]}3d 0%,transparent 68%);'
            f'filter:blur({14 if phone else 30}px);pointer-events:none"></div>')
    emblem_img = (f'<img src="{emblem(faction)}" alt="" style="position:absolute;top:{emblem_top}px;'
                  f'left:{round(w/2 - emblem_size/2)}px;width:{emblem_size}px;height:{emblem_size}px;'
                  f'object-fit:contain;filter:drop-shadow(0 2px 6px rgba(0,0,0,.6))">')
    return (f'<div style="position:absolute;top:0;right:{18 if phone else 70}px;width:{w}px;height:{h}px;'
            f'pointer-events:none" aria-hidden="true">{glow}{svg}{emblem_img}</div>')


def art_watermark(faction: str, phone: bool = False) -> str:
    """Option B: a crisp large emblem watermark, no blur, over a sharp diagonal faction-
    colour vignette -- the identity mark beside the h1 is unchanged (still the 64px/44px
    ringed FactionCrest, rendered separately by the caller). Two layers, both built from
    nothing but the real emblem and the faction's own documented colour:
      1. The vignette: a single `linear-gradient(135deg, transparent 55%, <bar-colour-at-low-
         alpha> 100%)` wash across the band's own right portion -- a sharp diagonal line of
         colour, never a soft blurred cloud.
      2. The emblem itself at ~320px (desktop), 10-14% opacity, `object-fit:contain`, no
         filter at all -- crisp edges, positioned so its own right edge is cropped by the
         band's own right boundary (`right:-90px`, inside the overflow:hidden band), the
         owner's own "cropped by the band's right edge."
    """
    bar = FACTION_BAR[faction]
    size = 170 if phone else 320
    opacity = 0.14 if phone else 0.12
    vignette = (f'<div style="position:absolute;inset:0;pointer-events:none;'
                f'background:linear-gradient(135deg,transparent 52%,color-mix(in srgb,{bar} 38%,transparent) 100%)">'
                f'</div>')
    wm = (f'<img src="{emblem(faction)}" alt="" style="position:absolute;top:{-round(size*0.12)}px;'
          f'right:{-round(size*0.28)}px;width:{size}px;height:{size}px;object-fit:contain;'
          f'opacity:{opacity};pointer-events:none" aria-hidden="true">')
    return vignette + wm


def header_hero(role_slot: str, tabs_html: str, faction: str, phone: bool = False, style: str = 'banner') -> str:
    """The shipped band (eyebrow/h1/Updated/facts/role line/tab strip, Guild.svelte's own
    order, unchanged) plus the identity row's crest and the band's own art layer. Round-2 fix:
    the tab strip is now a child of the SAME max-width:1344px column as the header text --
    previously a sibling of that column, which only happened to align at exactly 1440px board
    width (see this file's own module docstring)."""
    pad = '22px 18px 0 18px' if phone else '28px 48px 0 48px'
    maxw = '' if phone else 'width:100%;max-width:1344px;margin:0 auto;'
    crest_size = 44 if phone else 64
    crest_html = faction_crest(faction, crest_size) if faction in FACTION_BAR else ''
    gap = 16 if phone else 18
    h1 = f'<h1 class="display" style="margin:0;font-size:22px;color:{TEXT}">{GUILD["name"]}</h1>'
    row = (f'<div style="display:flex;align-items:center;gap:{gap}px;position:relative;z-index:1">'
           f'{crest_html}{h1}</div>')

    art_fn = art_banner if style == 'banner' else art_watermark

    if phone:
        # Art layer clipped to the identity row's own box only -- never the facts/standing/
        # tabs below it (owner's own "fade behind the title only" instruction, round 1).
        identity = (f'<div style="position:relative;overflow:hidden;padding:4px 0 2px 0">'
                    f'{art_fn(faction, phone=True) if faction in FACTION_BAR else ""}{row}</div>')
        band_art = ''
    else:
        identity = row
        band_art = art_fn(faction, phone=False) if faction in FACTION_BAR else ''

    tabs_row = f'<div style="margin-top:18px;position:relative;z-index:1">{tabs_html}</div>'

    content = f'''<div style="position:relative;overflow:hidden">
<div style="position:relative;{maxw}">
{band_art}
<div style="position:relative;z-index:1;display:flex;flex-direction:column;gap:14px;padding:{pad}">
<span class="label" style="color:{GOLD};display:flex;align-items:center;gap:10px"><i style="width:28px;height:1px;background:{GOLD};display:inline-block"></i>Guild</span>
{identity}
<span style="font-size:12px;color:{MUTED}">{UPDATED_LABEL}</span>
<p style="margin:0;font-size:14px;color:{BODY}">{header_facts()}</p>
{role_slot}
</div>
{tabs_row}
</div>
</div>'''
    return content


FACTION_CAPTION = {
    'alliance': ('Header art study, option A &mdash; faction overridden to Alliance for this '
                 'round only. The mock roster&rsquo;s own names and classes are unchanged '
                 '(round 2&rsquo;s roster is written Horde-flavoured).'),
    'horde-banner': 'Header art study, option A: the hanging banner.',
    'horde-watermark': 'Header art study, option B: the emblem watermark.',
}


def build_header_board(name: str, faction: str, role: str, style: str, width_tag: str = '', phone: bool = False) -> str:
    officer = role == 'officer'
    viewer_name = VIEWER_OFFICER if officer else VIEWER_MEMBER
    role_slot = standing_line(viewer_name, with_fails=(role == 'member'))
    if officer:
        role_slot += officer_tools()
    tabs_html = tab_strip('Overview', role, phone=phone)
    band = header_hero(role_slot, tabs_html, faction, phone=phone, style=style)

    persona = {'class': 'warrior' if officer else 'hunter',
               'faction': faction if faction in FACTION_BAR else 'horde', 'battletag': viewer_name}
    nav_html = nav_phone(True, persona) if phone else nav('', True, persona)

    body_pad = '18px' if phone else '48px'
    body = overview_raids_summary() + overview_progression_summary()
    caption = FACTION_CAPTION.get(name.replace('guild-header-', ''), FACTION_CAPTION.get(faction, ''))
    content = (f'<div style="display:flex;flex-direction:column;gap:22px;padding:24px {body_pad} 32px {body_pad};'
               + ('' if phone else 'width:100%;max-width:1344px;margin:0 auto;') + '">'
               + body + mock_caption()
               + (f'<p style="margin:0;font-size:12px;color:{MUTED};font-style:italic">{caption}</p>' if caption else '')
               + '</div>')

    band_wrap = f'<div class="band" style="display:flex;flex-direction:column;flex-shrink:0">{nav_html}{band}</div>'
    out_html = QUIET_LINK_STYLE + band_wrap + content + footer(phone=phone)
    width = 390 if phone else (2000 if width_tag == '2000' else 1440)
    return write(f'{name}.html', page(out_html, width)).as_posix()


print(build_header_board('guild-header-horde-banner', 'horde', 'officer', style='banner'))
print(build_header_board('guild-header-horde-watermark', 'horde', 'officer', style='watermark'))
print(build_header_board('guild-header-alliance', 'alliance', 'officer', style='banner'))
print(build_header_board('guild-header-2000', 'horde', 'member', style='banner', width_tag='2000'))
print(build_header_board('guild-header-phone', 'horde', 'member', style='banner', phone=True))
