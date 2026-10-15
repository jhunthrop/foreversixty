// design/listing/md2html.mjs
// Renders a listing file's "## Description" section to the HTML CurseForge's description
// editor takes (paste into its HTML source view). Covers only what the two listings use:
// headings, paragraphs, bold, inline code, links, bullet lists and pipe tables.
//   node design/listing/md2html.mjs design/listing/curseforge-description.md > /tmp/desc.html
import { readFile } from 'node:fs/promises';

const escape = (text) =>
  text.replaceAll('&', '&amp;').replaceAll('<', '&lt;').replaceAll('>', '&gt;');

function inline(text) {
  return escape(text)
    .replace(/\*\*(.+?)\*\*/g, '<strong>$1</strong>')
    .replace(/`([^`]+)`/g, '<code>$1</code>')
    .replace(/\[([^\]]+)\]\(([^)]+)\)/g, '<a href="$2">$1</a>');
}

function descriptionOf(markdown) {
  const start = markdown.indexOf('\n## Description\n');
  if (start < 0) throw new Error('no "## Description" section');
  return markdown.slice(start + '\n## Description\n'.length).trim();
}

function render(markdown) {
  const out = [];
  const lines = markdown.split('\n');
  let i = 0;
  while (i < lines.length) {
    const line = lines[i];
    if (line.trim() === '') {
      i += 1;
    } else if (line.startsWith('### ')) {
      out.push(`<h3>${inline(line.slice(4))}</h3>`);
      i += 1;
    } else if (line.startsWith('## ')) {
      out.push(`<h2>${inline(line.slice(3))}</h2>`);
      i += 1;
    } else if (line.startsWith('- ')) {
      const items = [];
      while (i < lines.length && lines[i].startsWith('- ')) {
        items.push(`<li>${inline(lines[i].slice(2))}</li>`);
        i += 1;
      }
      out.push(`<ul>\n${items.join('\n')}\n</ul>`);
    } else if (line.startsWith('|')) {
      const rows = [];
      while (i < lines.length && lines[i].startsWith('|')) {
        rows.push(lines[i]);
        i += 1;
      }
      const cells = (row) => row.split('|').slice(1, -1).map((cell) => cell.trim());
      const [head, , ...body] = rows;
      const th = cells(head).map((cell) => `<th>${inline(cell)}</th>`).join('');
      const trs = body.map((row) => `<tr>${cells(row).map((cell) => `<td>${inline(cell)}</td>`).join('')}</tr>`);
      out.push(`<table>\n<thead><tr>${th}</tr></thead>\n<tbody>\n${trs.join('\n')}\n</tbody>\n</table>`);
    } else {
      const paragraph = [];
      while (i < lines.length && lines[i].trim() !== '' && !/^(#|- |\|)/.test(lines[i])) {
        paragraph.push(lines[i].trim());
        i += 1;
      }
      out.push(`<p>${inline(paragraph.join(' '))}</p>`);
    }
  }
  return out.join('\n');
}

const file = process.argv[2];
if (file === undefined) {
  process.stderr.write('usage: node md2html.mjs <listing.md>\n');
  process.exit(2);
}
process.stdout.write(`${render(descriptionOf(await readFile(file, 'utf8')))}\n`);
