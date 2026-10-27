// web/tests/captures/discord-bot/discord-render.ts
// A faithful render of a Discord message payload (the JSON the Go bot sends) into a Discord-like
// HTML card: dark theme, embed with its colour bar, fields, thumbnail, footer, button rows and
// the string select. It draws what the payload says and nothing else; it is not the real client.

export interface Component {
  type: number;
  style?: number;
  label?: string;
  custom_id?: string;
  url?: string;
  placeholder?: string;
  disabled?: boolean;
  options?: { label: string; value: string; description?: string }[];
}
interface Field {
  name: string;
  value: string;
  inline?: boolean;
}
interface Embed {
  title?: string;
  description?: string;
  url?: string;
  color?: number;
  fields?: Field[];
  author?: { name: string };
  thumbnail?: { url: string };
  footer?: { text: string };
  timestamp?: string;
}
export interface Payload {
  content?: string;
  embeds?: Embed[];
  components?: { type: number; components: Component[] }[];
  flags?: number;
}
export interface Dumped {
  id: string;
  title: string;
  surface: 'channel' | 'ephemeral';
  message: Payload;
}

const EPHEMERAL_FLAG = 64;
const DEFAULT_BAR = '#202225';
const BUTTON_STYLE: Record<number, string> = { 2: 'secondary', 3: 'success', 5: 'link' };
const TIME_TAG = /<t:(\d+):([FfR])>/g;
const NEW_YORK = 'America/New_York';

const escapeHtml = (s: string): string =>
  s.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;').replace(/"/g, '&quot;');

function timeTag(unix: string, style: string): string {
  const at = new Date(Number(unix) * 1000);
  if (style === 'R') return 'in 5 hours';
  return new Intl.DateTimeFormat('en-US', {
    timeZone: NEW_YORK,
    weekday: 'long',
    month: 'long',
    day: 'numeric',
    year: 'numeric',
    hour: 'numeric',
    minute: '2-digit',
  }).format(at);
}

/** Discord's inline markdown, the part the bot uses: bold, strike-through, timestamps. */
function markdown(raw: string): string {
  return escapeHtml(raw)
    .replace(/&lt;t:(\d+):([FfR])&gt;/g, (_m, unix: string, style: string) => `<span class="ts">${timeTag(unix, style)}</span>`)
    .replace(/\*\*(.+?)\*\*/g, '<strong>$1</strong>')
    .replace(/~~(.+?)~~/g, '<s>$1</s>')
    .replace(/\n/g, '<br>');
}

function embedHtml(e: Embed): string {
  const bar = e.color === undefined || e.color === 0 ? DEFAULT_BAR : `#${e.color.toString(16).padStart(6, '0')}`;
  const title = e.title === undefined ? '' : e.url ? `<a class="etitle" href="#">${markdown(e.title)}</a>` : `<div class="etitle plain">${markdown(e.title)}</div>`;
  const fields = (e.fields ?? [])
    .map((f) => `<div class="field ${f.inline ? 'inline' : ''}"><div class="fname">${markdown(f.name)}</div><div>${markdown(f.value)}</div></div>`)
    .join('');
  const stamp = e.timestamp === undefined ? '' : ' • Today at 4:00 PM';
  const footer = e.footer === undefined ? '' : `<div class="footer">${markdown(e.footer.text)}${stamp}</div>`;
  const thumb = e.thumbnail === undefined ? '' : '<div class="thumb">icon</div>';
  return `<div class="embed" style="border-left-color:${bar}">
    <div class="ebody">${e.author ? `<div class="author">${markdown(e.author.name)}</div>` : ''}${title}
    ${e.description ? `<div class="edesc">${markdown(e.description)}</div>` : ''}<div class="fields">${fields}</div>${footer}</div>${thumb}</div>`;
}

function componentHtml(c: Component): string {
  if (c.type === 3) {
    return `<div class="select"><span>${escapeHtml(c.placeholder ?? '')}</span><span>⌄</span></div>
      <div class="options">${(c.options ?? [])
        .map((o) => `<div class="option"><b>${escapeHtml(o.label)}</b><small>${escapeHtml(o.description ?? '')}</small></div>`)
        .join('')}</div>`;
  }
  const kind = BUTTON_STYLE[c.style ?? 2] ?? 'secondary';
  const external = kind === 'link' ? '<span class="ext">↗</span>' : '';
  return `<button class="btn ${kind}" ${c.disabled ? 'disabled' : ''}>${escapeHtml(c.label ?? '')}${external}</button>`;
}

/** One message as an HTML document body: the author line, the content, embeds and rows. */
export function renderMessage(entry: Dumped, width: number): string {
  const m = entry.message;
  const ephemeral = ((m.flags ?? 0) & EPHEMERAL_FLAG) !== 0;
  const rows = (m.components ?? [])
    .map((r) => `<div class="row">${r.components.map(componentHtml).join('')}</div>`)
    .join('');
  return `<!doctype html><meta charset="utf-8"><style>
  body{margin:0;background:#313338;color:#dbdee1;font:16px/1.375 "gg sans","Segoe UI",Helvetica,Arial,sans-serif}
  .wrap{box-sizing:border-box;width:${width}px;padding:16px;display:flex;gap:12px}
  .avatar{flex:none;width:40px;height:40px;border-radius:50%;background:#e5b955;color:#07090d;font-weight:700;display:grid;place-items:center}
  .main{min-width:0;flex:1}
  .head b{color:#f2f3f5}.app{background:#5865f2;color:#fff;font-size:10px;padding:1px 5px;border-radius:3px;margin:0 4px}
  .head small{color:#949ba4;font-size:12px}
  .content{margin:2px 0 4px}
  .embed{display:flex;gap:16px;max-width:432px;background:#2b2d31;border-left:4px solid #202225;border-radius:4px;padding:8px 16px 16px 12px;margin-top:4px}
  .ebody{min-width:0;flex:1}.author{font-size:14px;font-weight:600;margin-top:8px}
  .etitle{display:block;color:#00a8fc;font-weight:600;margin-top:8px;text-decoration:none}.etitle.plain{color:#f2f3f5}
  .edesc{font-size:14px;margin-top:8px}
  .fields{display:grid;grid-template-columns:repeat(3,1fr);gap:8px;margin-top:8px}
  .field{grid-column:1/-1;font-size:14px}.field.inline{grid-column:span 1}
  .fname{font-weight:600;color:#f2f3f5;margin-bottom:2px}
  .footer{font-size:12px;color:#b5bac1;margin-top:8px}
  .thumb{flex:none;width:64px;height:64px;background:#1e1f22;border-radius:4px;color:#6d6f78;font-size:11px;display:grid;place-items:center}
  .ts{background:#4e5058;border-radius:3px;padding:0 2px}
  .row{display:flex;flex-wrap:wrap;gap:8px;margin-top:8px}
  .btn{font:500 14px/1 inherit;font-family:inherit;border:0;border-radius:3px;padding:10px 16px;color:#fff;min-height:32px}
  .btn.secondary,.btn.link{background:#4e5058}.btn.success{background:#248046}.btn[disabled]{opacity:.5}.ext{margin-left:6px}
  .select{display:flex;justify-content:space-between;align-items:center;background:#1e1f22;border:1px solid #1e1f22;border-radius:4px;padding:10px 12px;width:100%;box-sizing:border-box;color:#b5bac1}
  .options{width:100%;margin-top:4px;background:#2b2d31;border-radius:4px;padding:4px}
  .option{padding:6px 8px;display:flex;flex-direction:column}.option small{color:#949ba4}
  .only{margin-top:6px;font-size:12px;color:#949ba4}.only a{color:#00a8fc}
  </style><div class="wrap" id="msg"><div class="avatar">FS</div><div class="main">
  <div class="head"><b>Forever Sixty</b><span class="app">APP</span><small>Today at 4:00 PM</small></div>
  ${m.content ? `<div class="content">${markdown(m.content)}</div>` : ''}
  ${(m.embeds ?? []).map(embedHtml).join('')}${rows}
  ${ephemeral ? '<div class="only">Only you can see this • <a href="#">Dismiss message</a></div>' : ''}
  </div></div>`;
}
