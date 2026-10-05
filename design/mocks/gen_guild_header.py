"""Mock boards for the guild header art round only (design/specs/2026-10-04-guild-page.md
§12.2). Scope: the header band alone -- faction emblem as a crest-ring identity mark, a
faction-coloured hero fade built from the same emblem, over the existing night gradient. The
roster/tabs/content below the band are reused unchanged from gen_guild.py so each board still
reads as a real page, not an isolated swatch.

  python3 design/mocks/gen_guild_header.py
    -> guild-header-alliance.html  (1440, officer, faction overridden to Alliance for this
                                     art study only -- the mock roster's own names/classes
                                     are unchanged, see each board's own caption)
    -> guild-header-horde.html     (1440, officer, the roster's native faction)
    -> guild-header-2000.html      (2000 wide, member, confirms the fade stays glued to the
                                     1344px content column, not the physical viewport edge)
    -> guild-header-phone.html     (390, member, 44px crest, fade scoped to the title row)

The neutral/unknown-faction fallback is not a fifth board here: it is pixel-identical to
today's shipped band (no crest, no fade, the plain night gradient alone) -- already on file
as design/mocks/renders/guild-overview-officer.png, cited rather than re-rendered.
"""
from mocklib import GOLD, MUTED, TEXT, BODY, emblem, nav, nav_phone, footer, page, write
from gen_guild import (
    GUILD, NIGHTS, RED, BLUE, UPDATED_LABEL, VIEWER_MEMBER, VIEWER_OFFICER,
    ROSTER_BY_NAME, header_facts, standing_line, officer_tools, tab_strip,
    overview_raids_summary, overview_progression_summary, mock_caption,
    QUIET_LINK_STYLE,
)

# Faction colour tokens (design/DESIGN-SYSTEM.md "Faction"): text vs bar variants are both
# documented there; the ring/fade use the bar variant (full-saturation swatch), matching
# ClassCrest.astro's own choice of classes.json's full-saturation `color` field for its ring,
# never the lightened text variant meant for small type on dark panels.
FACTION_BAR = {'alliance': '#2f6fd6', 'horde': '#c0392b'}
FACTION_TEXT = {'alliance': BLUE, 'horde': RED}


def faction_crest(faction: str, size: int) -> str:
    """§12.2 identity mark: ClassCrest.astro's own shipped ring recipe (circular crop, 2px
    box-shadow ring, --raised background fallback) applied to FactionMark's emblem asset
    instead of a class icon. A new mark, not a FactionMark.astro edit -- that component's own
    header comment states 'no ring, no background, unlike ClassCrest' for its nine existing
    callers (nav chip, row descriptors, the class header's faction toggle), and bolting a ring
    onto it here would change all nine. Padding keeps the shield/disc silhouette off the ring
    edge (the emblem's own art is not full-bleed square the way a class icon is)."""
    ring = FACTION_BAR[faction]
    pad = round(size * 0.16)
    return (f'<img src="{emblem(faction)}" alt="" style="width:{size}px;height:{size}px;'
            f'border-radius:999px;object-fit:contain;padding:{pad}px;box-sizing:border-box;'
            f'box-shadow:0 0 0 2px {ring};background:var(--raised);flex-shrink:0">')


def hero_fade(faction: str, phone: bool = False) -> str:
    """§12.2 hero fade: the same emblem asset, large, blurred and darkened as a backdrop at
    low opacity -- the exact 'disc made from the icon itself' mechanic tenet 7 uses for
    crests, at atmosphere scale instead of identity-mark scale (no crisp foreground layer;
    this one is never meant to be legible, only atmosphere, same job the night gradient
    already does). A radial mask fades it to transparent on every edge so it blends into the
    flat band colour rather than ending in a hard circle -- 'fading to --bg' read literally.

    Measured round-2 fix: the raw emblem art is not faction-pure -- the Alliance lion shield
    samples mostly gold/tan (the lion and its rim), the Horde disc mostly near-black (sampled
    directly from the shipped web/public/icons/hd/faction/*.webp pixels) -- so a plain blur
    of either reads gold or a dark smudge, never recognisably blue or red. Fixed the same way
    the design system already lightens rare/epic item text for legibility (design/DESIGN-
    SYSTEM.md's rarity section: adjust the real colour for its use, never invent a new one):
    the blurred layer is desaturated first, then colourised with the faction's own documented
    bar colour (design/DESIGN-SYSTEM.md "Faction") via `mix-blend-mode:color`, so the glow's
    *shape* is still the real emblem's own blurred silhouette (never invented art) and its
    *hue* is the one faction colour the design system already names for this.

    Desktop: anchored to the band's own 1344px content column (not the raw viewport), so it
    stays glued to the column's right edge at 1920/2000px instead of drifting to the browser
    edge (guild-header-2000's own job). Phone: scoped to the identity row's own box only
    (§12.2's 'fade behind the title only'), far smaller, never reaching the facts/standing/
    tab regions below it."""
    if faction not in FACTION_BAR:
        return ''
    src = emblem(faction)
    bar = FACTION_BAR[faction]
    # Round-2 fix (this round's own mock review): an earlier draft sized this box 480px
    # square while the band's own real content height is ~290px, so the local clip box
    # (below) cut the glow off in a hard horizontal line top and bottom instead of letting
    # its own radial mask fade out first -- a glow reading as a flat-edged rectangle, not
    # atmosphere. Fixed: an explicitly wide-short box (ellipse, not circle) sized to fit
    # inside the band's own real content height with its soft edge intact.
    width = 170 if phone else 460
    height = 150 if phone else 300
    top = -14 if phone else -10
    right = -30 if phone else -40
    blur = 24 if phone else 44
    opacity = 0.6 if phone else 0.68
    # Round-2 fix, second pass: two overlapping elements with filter + mix-blend-mode +
    # mask-image each is fragile across engines (isolation/stacking-context edge cases) and
    # rendered as a near-solid rectangle with almost no visible falloff on this board's own
    # first Playwright capture. One element instead: two CSS background-image layers (the
    # flat faction colour on top, the emblem photo underneath) combined with
    # `background-blend-mode:color` inside a SINGLE box, so `filter:blur()` blurs the
    # already-colourised, already-composited result as one unit, and one `mask-image` on
    # that same box does the fade -- no cross-element blending, far more reliable.
    mask = 'radial-gradient(ellipse at 68% 42%,#000 0%,#000 22%,transparent 72%)'
    return (
        f'<div style="position:absolute;top:{top}px;right:{right}px;width:{width}px;height:{height}px;'
        f'pointer-events:none;overflow:hidden;mask-image:{mask};-webkit-mask-image:{mask}" aria-hidden="true">'
        f'<div style="position:absolute;inset:-20%;background-image:linear-gradient({bar},{bar}),url({src});'
        f'background-blend-mode:color;background-size:cover,cover;background-position:center;'
        f'filter:blur({blur}px) brightness(.68) saturate(1.3);opacity:{opacity}"></div>'
        f'</div>'
    )


def header_hero(role_slot: str, tabs_html: str, faction: str, phone: bool = False) -> str:
    """The shipped band (eyebrow/h1/Updated/facts/role line/tab strip, Guild.svelte's own
    order, unchanged) plus the identity row's crest and the band's own hero fade -- the two
    new elements this round adds, nothing else moved (§12.2's own instruction)."""
    pad = '22px 18px 0 18px' if phone else '28px 48px 0 48px'
    maxw = '' if phone else 'width:100%;max-width:1344px;margin:0 auto;'
    crest_size = 44 if phone else 64
    crest_html = faction_crest(faction, crest_size) if faction in FACTION_BAR else ''
    gap = 16 if phone else 18
    h1 = f'<h1 class="display" style="margin:0;font-size:22px;color:{TEXT}">{GUILD["name"]}</h1>'
    row = (f'<div style="display:flex;align-items:center;gap:{gap}px;position:relative;z-index:1">'
           f'{crest_html}{h1}</div>')

    if phone:
        # Fade clipped to the identity row's own box only -- never the facts/standing/tabs
        # below it, which this `header_hero` call also renders but outside this wrapper.
        identity = (f'<div style="position:relative;overflow:hidden;padding:4px 0 2px 0">'
                    f'{hero_fade(faction, phone=True)}{row}</div>')
        outer_fade = ''
    else:
        identity = row
        outer_fade = hero_fade(faction, phone=False)

    # The fade's own positioning context is the 1344px content column, not the full-bleed
    # outer wrapper -- so `right:` tracks the column's own right edge at any viewport width
    # (guild-header-2000's own job) instead of drifting toward the physical browser edge.
    content = f'''<div style="position:relative;overflow:hidden">
<div style="position:relative;{maxw}">
{outer_fade}
<div style="position:relative;z-index:1;display:flex;flex-direction:column;gap:14px;padding:{pad}">
<span class="label" style="color:{GOLD};display:flex;align-items:center;gap:10px"><i style="width:28px;height:1px;background:{GOLD};display:inline-block"></i>Guild</span>
{identity}
<span style="font-size:12px;color:{MUTED}">{UPDATED_LABEL}</span>
<p style="margin:0;font-size:14px;color:{BODY}">{header_facts()}</p>
{role_slot}
</div>
</div>
</div>'''
    return content + f'<div style="margin-top:18px">{tabs_html}</div>'


FACTION_CAPTION = {
    'alliance': ('Header art study &mdash; faction overridden to Alliance for this round only. '
                 'The mock roster&rsquo;s own names and classes are unchanged (round 2&rsquo;s '
                 'roster is written Horde-flavoured); a real Alliance guild&rsquo;s page reads '
                 'identically, with its own roster.'),
    'horde': ('Header art study &mdash; the mock roster&rsquo;s native faction (Horde). '
              'Body content below the band is reused unchanged from the round-2 boards.'),
}


def build_header_board(name: str, faction: str, role: str, width_tag: str = '', phone: bool = False) -> str:
    officer = role == 'officer'
    viewer_name = VIEWER_OFFICER if officer else VIEWER_MEMBER
    viewer_row = ROSTER_BY_NAME[viewer_name]
    role_slot = standing_line(viewer_name, with_fails=(role == 'member'))
    if officer:
        role_slot += officer_tools()
    tabs_html = tab_strip('Overview', role, phone=phone)
    band = header_hero(role_slot, tabs_html, faction, phone=phone)

    persona = {'class': 'warrior' if officer else 'hunter', 'faction': faction if faction in FACTION_BAR else 'horde',
               'battletag': viewer_name}
    # Bug found on this round's own first phone capture: the desktop `nav()` (seven
    # unwrapped links + chip + Discord button in one flex row) was used here regardless of
    # `phone`, forcing a ~1270px-wide nav row inside a 390px viewport and pushing the whole
    # page's scrollWidth out to 986px -- the screenshot came back 986x1279 instead of
    # 390-wide. `nav_phone()` is gen_guild.py's own existing phone nav (logo + crest + menu
    # glyph only); using it here is the fix, matching gen_guild.py's own `nav_fn` switch.
    nav_html = nav_phone(True, persona) if phone else nav('', True, persona)

    body_pad = '18px' if phone else '48px'
    body = overview_raids_summary() + overview_progression_summary()
    caption = FACTION_CAPTION.get(faction, '')
    content = (f'<div style="display:flex;flex-direction:column;gap:22px;padding:24px {body_pad} 32px {body_pad};'
               + ('' if phone else 'width:100%;max-width:1344px;margin:0 auto;') + '">'
               + body + mock_caption()
               + (f'<p style="margin:0;font-size:12px;color:{MUTED};font-style:italic">{caption}</p>' if caption else '')
               + '</div>')

    band_wrap = f'<div class="band" style="display:flex;flex-direction:column;flex-shrink:0">{nav_html}{band}</div>'
    out_html = QUIET_LINK_STYLE + band_wrap + content + footer(phone=phone)
    width = 390 if phone else (2000 if width_tag == '2000' else 1440)
    return write(f'{name}.html', page(out_html, width)).as_posix()


print(build_header_board('guild-header-alliance', 'alliance', 'officer'))
print(build_header_board('guild-header-horde', 'horde', 'officer'))
print(build_header_board('guild-header-2000', 'horde', 'member', width_tag='2000'))
print(build_header_board('guild-header-phone', 'horde', 'member', phone=True))
