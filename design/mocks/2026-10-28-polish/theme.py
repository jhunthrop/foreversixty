"""Theme tokens for the 2026-10-28 polish boards: the two candidate type and surface
directions (A Frost, cooler and sharper; B Forge, warmer and heavier) as CSS variable sets,
and the shared component CSS both directions draw with. Run gen_polish.py, not this file."""

DIRECTIONS = {
    "A": {
        "name": "Frost",
        "fonts": "family=Rajdhani:wght@600;700&family=Inter:wght@400;500;600;700&family=JetBrains+Mono:wght@500;700&family=Cinzel:wght@700",
        "vars": {
            "--fh": "\"Rajdhani\", \"Inter\", sans-serif", "--fhw": "700", "--fht": ".03em",
            "--fb": "\"Inter\", \"Helvetica Neue\", Arial, sans-serif", "--fbs": "14px",
            "--card": "linear-gradient(180deg, rgba(130,175,255,.065), rgba(130,175,255,.012) 55%), #0a0f18",
            "--card-b": "rgba(150,185,235,.15)", "--hl": "rgba(124,196,255,.30)",
            "--shadow": "0 14px 34px rgba(0,0,0,.5)",
            "--rp": "6px", "--rc": "4px", "--rpill": "3px",
            "--tile": "rgba(124,196,255,.09)", "--tile-c": "#9fd2ff",
            "--btn-case": "none", "--btn-size": "14px", "--btn-w": "600", "--btn-ls": ".01em",
            "--band": "radial-gradient(1200px 380px at 70% -10%, rgba(60,110,190,.35), transparent 70%), linear-gradient(180deg, #0b1320, #07090d)",
            "--row-hover": "rgba(124,196,255,.06)", "--glow": "0 0 18px rgba(229,185,85,.22)",
            "--h1": "44px", "--h1-phone": "34px", "--h2": "22px",
        },
    },
    "B": {
        "name": "Forge",
        "fonts": "family=Barlow+Condensed:wght@600;700;800&family=Barlow:wght@400;500;600;700&family=JetBrains+Mono:wght@500;700&family=Cinzel:wght@700",
        "vars": {
            "--fh": "\"Barlow Condensed\", \"Barlow\", sans-serif", "--fhw": "800", "--fht": ".01em",
            "--fb": "\"Barlow\", \"Helvetica Neue\", Arial, sans-serif", "--fbs": "15px",
            "--card": "linear-gradient(180deg, rgba(229,185,85,.075), rgba(229,185,85,.012) 45%), #0d0f15",
            "--card-b": "rgba(229,185,85,.17)", "--hl": "rgba(245,210,122,.45)",
            "--shadow": "0 16px 38px rgba(0,0,0,.55)",
            "--rp": "12px", "--rc": "8px", "--rpill": "6px",
            "--tile": "rgba(229,185,85,.12)", "--tile-c": "#f5d27a",
            "--btn-case": "uppercase", "--btn-size": "13px", "--btn-w": "700", "--btn-ls": ".07em",
            "--band": "radial-gradient(1100px 380px at 72% -10%, rgba(214,110,40,.28), transparent 70%), linear-gradient(180deg, #14100b, #07090d)",
            "--row-hover": "rgba(229,185,85,.06)", "--glow": "0 0 26px rgba(229,185,85,.32)",
            "--h1": "52px", "--h1-phone": "40px", "--h2": "26px",
        },
    },
}

BASE_CSS = """
*{box-sizing:border-box}html,body{margin:0;background:#07090d;color:#e9e4d8;font-family:var(--fb);font-size:var(--fbs);line-height:1.45;-webkit-font-smoothing:antialiased}
a{color:inherit;text-decoration:none}h1,h2,h3,p{margin:0}
:root{--bg:#07090d;--line:#262e40;--line-soft:#1c2230;--text:#e9e4d8;--strong:#f2eee4;--muted:#9a9484;--nav:#b9b3a4;--gold:#e5b955;--gold-hover:#f5d27a;--kill:#7fd48a;--wipe:#f0736b;--discord:#5865f2}
.mono{font-family:"JetBrains Mono",Menlo,monospace;font-variant-numeric:tabular-nums}
.wordmark{display:flex;align-items:center;gap:10px;font-family:"Cinzel",serif;font-weight:700;letter-spacing:.10em;color:var(--strong);text-transform:uppercase;font-size:17px;white-space:nowrap}
.micro{font-size:11px;font-weight:700;letter-spacing:.14em;text-transform:uppercase;color:var(--muted)}
.micro.gold{color:var(--gold)}
.h1{font-family:var(--fh);font-weight:var(--fhw);letter-spacing:var(--fht);text-transform:uppercase;font-size:var(--h1);line-height:1;color:var(--strong)}
.h1 .grad{background:linear-gradient(180deg,#fbe7a1,#e5b955 55%,#a8762a);-webkit-background-clip:text;background-clip:text;color:transparent}
.h2{font-family:var(--fh);font-weight:var(--fhw);letter-spacing:var(--fht);text-transform:uppercase;font-size:var(--h2);line-height:1.1;color:var(--strong)}
.h3{font-family:var(--fh);font-weight:var(--fhw);letter-spacing:var(--fht);text-transform:uppercase;font-size:18px;line-height:1.15;color:var(--strong)}
.card{background:var(--card);border:1px solid var(--card-b);border-radius:var(--rp);box-shadow:inset 0 1px 0 var(--hl),var(--shadow)}
.card.quiet{box-shadow:var(--shadow)}
.card-head{display:flex;align-items:center;gap:12px;padding:16px 20px}
.card-head .grow{flex:1;min-width:0}
.itile{flex:none;width:36px;height:36px;border-radius:var(--rc);display:grid;place-items:center;background:var(--tile);color:var(--tile-c)}
.itile.sm{width:32px;height:32px}
.itile.discord{background:rgba(88,101,242,.16);color:#9aa4ff}
.row{display:flex;align-items:center;gap:14px;min-height:56px;padding:10px 20px;border-top:1px solid var(--line-soft)}
.row.link:hover,.row.hover{background:var(--row-hover)}
.row .t{flex:1;min-width:0}.row .lab{font-size:15px;font-weight:600;color:var(--strong);line-height:1.3}.row .sub{font-size:13px;color:var(--muted);line-height:1.35;margin-top:2px}
.row .val{font-size:14px;color:var(--text);text-align:right;white-space:nowrap}.row .val.muted{color:var(--muted)}
.row.nested{padding-left:70px;min-height:52px}
.chev{flex:none;color:var(--muted)}
.switch{flex:none;position:relative;width:44px;height:24px;border-radius:999px;background:#232a3a;border:1px solid #303a50}
.switch::after{content:"";position:absolute;top:2px;left:2px;width:18px;height:18px;border-radius:999px;background:#8a8576;transition:left .12s}
.switch.on{background:linear-gradient(180deg,#f0cc6c,#c99a3a);border-color:#a8762a;box-shadow:var(--glow)}
.switch.on::after{left:22px;background:#0b0e14}
.switch.dis{opacity:.45}
.switch.focus{outline:2px solid var(--gold);outline-offset:3px}
.btn{display:inline-flex;align-items:center;justify-content:center;gap:8px;height:44px;padding:0 18px;border-radius:var(--rc);font-family:var(--fb);font-size:var(--btn-size);font-weight:var(--btn-w);letter-spacing:var(--btn-ls);text-transform:var(--btn-case);white-space:nowrap;border:1px solid transparent}
.btn-primary{background:linear-gradient(180deg,#f0cc6c,#c99a3a);color:#0b0e14;box-shadow:inset 0 0 0 1px #a8762a,0 6px 20px rgba(229,185,85,.22)}
.btn-secondary{background:rgba(255,255,255,.045);border-color:var(--card-b);color:var(--strong)}
.btn-ghost{color:var(--gold);padding:0 10px}
.btn-danger{color:var(--wipe);background:rgba(240,115,107,.06);border-color:rgba(240,115,107,.32)}
.btn.sm{height:36px;padding:0 14px}
.btn.full{width:100%}
.tag{display:inline-flex;align-items:center;gap:6px;height:22px;padding:0 8px;border-radius:var(--rpill);font-size:11px;font-weight:700;letter-spacing:.08em;text-transform:uppercase;white-space:nowrap;border:1px solid}
.tag-gold{color:var(--gold);background:rgba(229,185,85,.12);border-color:rgba(229,185,85,.38)}
.tag-green{color:var(--kill);background:rgba(127,212,138,.10);border-color:rgba(127,212,138,.32)}
.tag-muted{color:var(--muted);background:rgba(154,148,132,.10);border-color:rgba(154,148,132,.30)}
.dot{width:7px;height:7px;border-radius:999px;background:currentColor;box-shadow:0 0 8px currentColor}
.alink{display:inline-flex;align-items:center;gap:6px;min-height:44px;color:var(--gold);font-size:14px;font-weight:600}
.ilink{color:var(--gold);text-decoration:underline;text-decoration-color:rgba(229,185,85,.4);text-underline-offset:3px;text-decoration-thickness:1px}
.crest{flex:none;border-radius:999px;object-fit:cover;box-shadow:0 0 0 2px color-mix(in srgb,var(--c) 60%,transparent),0 0 14px color-mix(in srgb,var(--c) 25%,transparent);background:#0d111a}
.stat{display:flex;flex-direction:column;gap:2px;padding:12px 14px;border-radius:var(--rc);background:rgba(255,255,255,.03);border:1px solid var(--line-soft)}
.stat .n{font-family:var(--fh);font-weight:var(--fhw);font-size:28px;line-height:1;color:var(--strong)}
.icon{flex:none;border-radius:var(--rpill);border:1px solid var(--q,#262e40);object-fit:cover}
.tick{color:var(--kill);flex:none;margin-top:3px}
.seg{display:inline-flex;padding:3px;border-radius:calc(var(--rc) + 3px);background:rgba(255,255,255,.04);border:1px solid var(--card-b)}
.seg span{display:inline-flex;align-items:center;height:38px;padding:0 18px;border-radius:var(--rc);font-size:14px;font-weight:600;color:var(--nav)}
.seg span.on{background:rgba(229,185,85,.16);color:var(--gold);box-shadow:inset 0 0 0 1px rgba(229,185,85,.55)}
.field{display:flex;align-items:center;gap:10px;height:52px;padding:0 14px;border-radius:var(--rc);background:rgba(0,0,0,.35);border:1px solid var(--card-b);color:var(--muted)}
.select{display:flex;align-items:center;justify-content:space-between;gap:10px;height:52px;padding:0 14px;border-radius:var(--rc);background:rgba(255,255,255,.04);border:1px solid var(--card-b);color:var(--strong);font-weight:600}
.kbd{font-family:"JetBrains Mono",monospace;font-size:11px;color:var(--muted);border:1px solid var(--line);border-radius:4px;padding:1px 6px}
.page-col{width:100%;max-width:1344px;margin:0 auto;padding:0 48px}
.phone .page-col{padding:0 18px}
.note{font-size:12px;color:var(--muted);font-family:"JetBrains Mono",monospace}
"""


def css_vars(key: str) -> str:
    return ":root{" + ";".join(f"{k}:{v}" for k, v in DIRECTIONS[key]["vars"].items()) + "}"


def head(key: str, extra: str = "") -> str:
    fonts = DIRECTIONS[key]["fonts"]
    return ("<!doctype html><html lang=\"en\"><head><meta charset=\"utf-8\"><title>Forever Sixty polish mock</title>"
            f"<link rel=\"stylesheet\" href=\"https://fonts.googleapis.com/css2?{fonts}&display=swap\">"
            f"<style>{css_vars(key)}{BASE_CSS}{extra}</style></head><body>")
