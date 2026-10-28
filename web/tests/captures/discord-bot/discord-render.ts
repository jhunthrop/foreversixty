// web/tests/captures/discord-bot/discord-render.ts
// A faithful render of a Discord message payload (the JSON the Go bot sends) into a Discord-like
// HTML card: dark theme, embed with its colour bar, fields, thumbnail, footer, button rows and
// the string select. It draws what the payload says and nothing else; it is not the real client.
// Times are drawn the way Discord draws its timestamp markup for a reader in the guild's raid
// zone, relative to the entry's clock (the render's "now"), so a post that is over reads "ago".

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
  /** The spec §9 Discord row the frame belongs to (D1 to D10). */
  row: string;
  title: string;
  surface: 'channel' | 'ephemeral';
  /** The render's "now", RFC 3339. */
  clock: string;
  message: Payload;
}

export interface RenderOptions {
  width: number;
  /** Draw a string select opened, its options listed. */
  selectOpen: boolean;
  /** The line Discord draws above a follow-up of a command (`Zulmara used /bis`). */
  usedCommand?: string;
}

const EPHEMERAL_FLAG = 64;
const DEFAULT_BAR = '#202225';
const BUTTON_STYLE: Record<number, string> = { 2: 'secondary', 3: 'success', 5: 'link' };
/** The guild's raid zone: the example guild raids at 8:00 pm CST. */
const READER_ZONE = 'America/Chicago';
const MINUTE = 60;
const HOUR = 60 * MINUTE;
const DAY = 24 * HOUR;
const MONTH = 30 * DAY;

const escapeHtml = (s: string): string =>
  s.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;').replace(/"/g, '&quot;');

function format(at: Date, options: Intl.DateTimeFormatOptions): string {
  return new Intl.DateTimeFormat('en-US', { timeZone: READER_ZONE, ...options }).format(at);
}

/** Discord's relative stamp: `in 4 hours`, `11 hours ago`, `a day ago`, `2 days ago`. */
export function relative(at: Date, now: Date): string {
  const seconds = Math.round((at.getTime() - now.getTime()) / 1000);
  const size = Math.abs(seconds);
  const unit = (n: number, word: string): string =>
    n === 1 ? `${word === 'hour' ? 'an' : 'a'} ${word}` : `${n} ${word}s`;
  let text: string;
  if (size < MINUTE) text = 'a few seconds';
  else if (size < HOUR) text = unit(Math.round(size / MINUTE), 'minute');
  else if (size < DAY) text = unit(Math.round(size / HOUR), 'hour');
  else if (size < MONTH) text = unit(Math.round(size / DAY), 'day');
  else text = unit(Math.round(size / MONTH), 'month');
  return seconds >= 0 ? `in ${text}` : `${text} ago`;
}

function sameDay(a: Date, b: Date): boolean {
  const day = (d: Date): string => format(d, { year: 'numeric', month: 'numeric', day: 'numeric' });
  return day(a) === day(b);
}

/** The message header's and an embed footer's time: `Today at 4:00 PM`, else the date. */
function clockText(at: Date, now: Date): string {
  const time = format(at, { hour: 'numeric', minute: '2-digit' });
  if (sameDay(at, now)) return `Today at ${time}`;
  if (sameDay(at, new Date(now.getTime() - DAY * 1000))) return `Yesterday at ${time}`;
  return `${format(at, { month: '2-digit', day: '2-digit', year: 'numeric' })} ${time}`;
}

function timeTag(unix: string, style: string, now: Date): string {
  const at = new Date(Number(unix) * 1000);
  if (style === 'R') return relative(at, now);
  if (style === 'f')
    return format(at, { month: 'long', day: 'numeric', year: 'numeric', hour: 'numeric', minute: '2-digit' });
  return format(at, {
    weekday: 'long',
    month: 'long',
    day: 'numeric',
    year: 'numeric',
    hour: 'numeric',
    minute: '2-digit',
  });
}

/** A masked link, `[text](https://…)`, as Discord draws it in embed descriptions and field values. */
const MASKED_LINK = /\[([^\]]+)\]\((https?:\/\/[^)\s]+)\)/g;

/** Discord's inline markdown, the part the bot uses: masked links, bold, strike-through, timestamps. */
function markdown(raw: string, now: Date): string {
  return escapeHtml(raw)
    .replace(MASKED_LINK, (_m, text: string, url: string) => `<a class="mlink" href="${url}">${text}</a>`)
    .replace(
      /&lt;t:(\d+):([FfR])&gt;/g,
      (_m, unix: string, style: string) => `<span class="ts">${timeTag(unix, style, now)}</span>`,
    )
    .replace(/\*\*(.+?)\*\*/g, '<strong>$1</strong>')
    .replace(/~~(.+?)~~/g, '<s>$1</s>')
    .replace(/\n/g, '<br>');
}

function embedHtml(e: Embed, now: Date): string {
  const md = (s: string): string => markdown(s, now);
  const bar =
    e.color === undefined || e.color === 0 ? DEFAULT_BAR : `#${e.color.toString(16).padStart(6, '0')}`;
  const title =
    e.title === undefined
      ? ''
      : e.url
        ? `<a class="etitle" href="#">${md(e.title)}</a>`
        : `<div class="etitle plain">${md(e.title)}</div>`;
  const fields = (e.fields ?? [])
    .map(
      (f) =>
        `<div class="field ${f.inline ? 'inline' : ''}"><div class="fname">${md(f.name)}</div><div>${md(f.value)}</div></div>`,
    )
    .join('');
  const stamp = e.timestamp === undefined ? '' : ` • ${clockText(new Date(e.timestamp), now)}`;
  const footer = e.footer === undefined ? '' : `<div class="footer">${md(e.footer.text)}${stamp}</div>`;
  const thumb =
    e.thumbnail === undefined ? '' : `<img class="thumb" src="${escapeHtml(e.thumbnail.url)}" alt="">`;
  return `<div class="embed" style="border-left-color:${bar}">
    <div class="ebody">${e.author ? `<div class="author">${md(e.author.name)}</div>` : ''}${title}
    ${e.description ? `<div class="edesc">${md(e.description)}</div>` : ''}<div class="fields">${fields}</div>${footer}</div>${thumb}</div>`;
}

function componentHtml(c: Component, selectOpen: boolean): string {
  if (c.type === 3) {
    const options = (c.options ?? [])
      .map(
        (o) =>
          `<div class="option"><b>${escapeHtml(o.label)}</b><small>${escapeHtml(o.description ?? '')}</small></div>`,
      )
      .join('');
    return `<div class="select ${selectOpen ? 'open' : ''}"><span>${escapeHtml(c.placeholder ?? '')}</span><span>${selectOpen ? '⌃' : '⌄'}</span></div>${
      selectOpen ? `<div class="options">${options}</div>` : ''
    }`;
  }
  const kind = BUTTON_STYLE[c.style ?? 2] ?? 'secondary';
  const external = kind === 'link' ? '<span class="ext">↗</span>' : '';
  return `<button class="btn ${kind}" ${c.disabled ? 'disabled' : ''}>${escapeHtml(c.label ?? '')}${external}</button>`;
}

/** One message as an HTML document body: the author line, the content, embeds and rows. */
export function renderMessage(entry: Dumped, options: RenderOptions): string {
  const m = entry.message;
  const now = new Date(entry.clock);
  const ephemeral = ((m.flags ?? 0) & EPHEMERAL_FLAG) !== 0;
  const rows = (m.components ?? [])
    .map(
      (r) =>
        `<div class="row">${r.components.map((c) => componentHtml(c, options.selectOpen)).join('')}</div>`,
    )
    .join('');
  const used = options.usedCommand
    ? `<div class="used">↱ <b>${escapeHtml(options.usedCommand)}</b></div>`
    : '';
  return `<!doctype html><meta charset="utf-8"><style>
  body{margin:0;background:#313338;color:#dbdee1;font:16px/1.375 "gg sans","Segoe UI",Helvetica,Arial,sans-serif}
  .wrap{box-sizing:border-box;width:${options.width}px;padding:16px}
  .used{font-size:13px;color:#949ba4;margin:0 0 4px 52px}.used b{color:#c4c9ce;font-weight:500}
  .msg{display:flex;gap:12px}
  .avatar{flex:none;width:40px;height:40px;border-radius:50%;background:#e5b955;color:#07090d;font-weight:700;display:grid;place-items:center}
  .main{min-width:0;flex:1}
  .head b{color:#f2f3f5}.app{background:#5865f2;color:#fff;font-size:10px;padding:1px 5px;border-radius:3px;margin:0 4px}
  .head small{color:#949ba4;font-size:12px}
  .content{margin:2px 0 4px}
  .embed{display:flex;gap:16px;max-width:520px;box-sizing:border-box;background:#2b2d31;border-left:4px solid #202225;border-radius:4px;padding:8px 16px 16px 12px;margin-top:4px}
  .ebody{min-width:0;flex:1}.author{font-size:14px;font-weight:600;margin-top:8px;color:#f2f3f5}
  .etitle{display:block;color:#00a8fc;font-weight:600;margin-top:8px;text-decoration:none}.etitle.plain{color:#f2f3f5}
  .edesc{font-size:14px;margin-top:8px}
  .fields{display:grid;grid-template-columns:repeat(3,1fr);gap:8px;margin-top:8px}
  .field{grid-column:1/-1;font-size:14px}.field.inline{grid-column:span 1}
  @media (max-width:500px){.field.inline{grid-column:1/-1}}
  .fname{font-weight:600;color:#f2f3f5;margin-bottom:2px}
  .footer{font-size:12px;color:#b5bac1;margin-top:8px}
  .thumb{flex:none;width:80px;height:80px;margin-top:8px;border-radius:4px;object-fit:contain}
  .ts{background:#4e5058;border-radius:3px;padding:0 2px}
  .mlink{color:#00a8fc;text-decoration:none}
  .row{display:flex;flex-wrap:wrap;gap:8px;margin-top:8px}
  .btn{font:500 14px/1 inherit;font-family:inherit;border:0;border-radius:3px;padding:10px 16px;color:#fff;min-height:32px}
  .btn.secondary,.btn.link{background:#4e5058}.btn.success{background:#248046}.btn[disabled]{opacity:.5}.ext{margin-left:6px}
  .select{display:flex;justify-content:space-between;align-items:center;background:#1e1f22;border:1px solid #1e1f22;border-radius:4px;padding:10px 12px;width:100%;max-width:400px;box-sizing:border-box;color:#b5bac1}
  .select.open{border-color:#00a8fc}
  .options{width:100%;max-width:400px;margin-top:4px;background:#2b2d31;border:1px solid #1e1f22;border-radius:4px;padding:4px;box-sizing:border-box}
  .option{padding:6px 8px;display:flex;flex-direction:column}.option b{color:#dbdee1;font-weight:500}.option small{color:#949ba4}
  .only{margin-top:6px;font-size:12px;color:#949ba4}.only a{color:#00a8fc}
  </style><div class="wrap" id="msg">${used}<div class="msg"><div class="avatar">FS</div><div class="main">
  <div class="head"><b>Forever Sixty</b><span class="app">APP</span><small>${clockText(now, now)}</small></div>
  ${m.content ? `<div class="content">${markdown(m.content, now)}</div>` : ''}
  ${(m.embeds ?? []).map((e) => embedHtml(e, now)).join('')}${rows}
  ${ephemeral ? '<div class="only">Only you can see this • <a href="#">Dismiss message</a></div>' : ''}
  </div></div></div>`;
}
