#!/usr/bin/env node
// The slop copy scan: runs tools/slop/patterns.txt over every copy source listed in
// tools/slop/sources.txt, minus comments, and reports each hit that allowlist.txt does
// not excuse. The web vitest runs it with --check; api/internal/slopguard mirrors it in Go.
import { readFileSync, readdirSync, statSync } from "node:fs";
import { dirname, join, extname } from "node:path";
import { fileURLToPath } from "node:url";

const HERE = dirname(fileURLToPath(import.meta.url));
const REPO = join(HERE, "..", "..");
const COPY_EXTENSIONS = new Set([".ts", ".js", ".mjs", ".svelte", ".astro", ".md", ".json", ".go", ".lua"]);
const GENERATED_HEADER = /generated|do not edit/i;
const GENERATED_HEADER_LINES = 5;

function dataLines(file) {
  return readFileSync(join(HERE, file), "utf8")
    .split("\n")
    .map((line) => line.replace(/\r$/, ""))
    .filter((line) => line.trim() !== "" && !line.startsWith("#"));
}

function splitTabs(line, count, file) {
  const parts = line.split("\t");
  if (parts.length !== count) {
    throw new Error(`${file}: want ${count} tab-separated fields, got ${parts.length}: ${line}`);
  }
  return parts;
}

export function loadPatterns() {
  return dataLines("patterns.txt").map((line) => {
    const [id, flags, source] = splitTabs(line, 3, "patterns.txt");
    return { id, regex: new RegExp(source, flags === "i" ? "i" : "") };
  });
}

export function loadAllowlist() {
  return dataLines("allowlist.txt").map((line) => {
    const [id, path, text] = splitTabs(line, 3, "allowlist.txt");
    return { id, path: new RegExp(path), text };
  });
}

export function loadSources() {
  const lines = dataLines("sources.txt");
  const pick = (mark) => lines.filter((l) => l.startsWith(mark)).map((l) => l.slice(2).trim());
  return { include: pick("+ "), exclude: pick("- ").map((r) => new RegExp(r)) };
}

function walk(path) {
  const stat = statSync(join(REPO, path), { throwIfNoEntry: false });
  if (!stat) throw new Error(`sources.txt: ${path} does not exist`);
  if (stat.isFile()) return [path];
  return readdirSync(join(REPO, path))
    .filter((name) => name !== "node_modules")
    .flatMap((name) => walk(join(path, name)));
}

function isGenerated(text) {
  return text.split("\n").slice(0, GENERATED_HEADER_LINES).some((line) => GENERATED_HEADER.test(line));
}

export function copyFiles(sources = loadSources()) {
  return sources.include
    .flatMap(walk)
    .filter((path) => COPY_EXTENSIONS.has(extname(path)))
    .filter((path) => !sources.exclude.some((re) => re.test(path)))
    .sort();
}

// Comment syntax per extension: block open/close pairs and the line-comment marker.
const COMMENTS = {
  c: { blocks: [["/*", "*/"], ["<!--", "-->"]], line: "//" },
  lua: { blocks: [["--[[", "]]"]], line: "--" },
  md: { blocks: [["<!--", "-->"]], line: null },
  json: { blocks: [], line: null },
};

function commentStyle(path) {
  const ext = extname(path);
  if (ext === ".lua") return COMMENTS.lua;
  if (ext === ".md") return COMMENTS.md;
  if (ext === ".json") return COMMENTS.json;
  return COMMENTS.c;
}

// stripLineComment drops a line comment that starts the line or follows whitespace,
// so "https://" and "a--b" survive.
function stripLineComment(text, marker) {
  if (!marker) return text;
  if (text.trimStart().startsWith(marker)) return "";
  const at = text.search(new RegExp(`\\s${marker.replace(/[-/]/g, "\\$&")}(\\s|$)`));
  return at === -1 ? text : text.slice(0, at);
}

// copyLines returns [lineNumber, text] for every line with its comments removed.
export function copyLines(path, text) {
  const style = commentStyle(path);
  let closing = null;
  return text.split("\n").map((raw, index) => {
    let line = raw;
    let kept = "";
    while (line !== "") {
      if (closing) {
        const end = line.indexOf(closing);
        if (end === -1) { line = ""; break; }
        line = line.slice(end + closing.length);
        closing = null;
        continue;
      }
      const opens = style.blocks
        .map(([open, close]) => ({ at: line.indexOf(open), open, close }))
        .filter((o) => o.at !== -1)
        .sort((a, b) => a.at - b.at);
      if (opens.length === 0) { kept += line; break; }
      kept += line.slice(0, opens[0].at);
      line = line.slice(opens[0].at + opens[0].open.length);
      closing = opens[0].close;
    }
    return [index + 1, stripLineComment(kept, style.line)];
  });
}

export function scan({ patterns = loadPatterns(), allowlist = loadAllowlist(), files = copyFiles() } = {}) {
  const hits = [];
  for (const path of files) {
    const text = readFileSync(join(REPO, path), "utf8");
    if (isGenerated(text)) continue;
    for (const [line, copy] of copyLines(path, text)) {
      for (const { id, regex } of patterns) {
        if (!regex.test(copy)) continue;
        const allowed = allowlist.some((a) => a.id === id && a.path.test(path) && copy.includes(a.text));
        hits.push({ id, path, line, text: copy.trim(), allowed });
      }
    }
  }
  return hits;
}

function countsById(hits, patterns) {
  return patterns.map(({ id }) => [id, hits.filter((h) => h.id === id && !h.allowed).length]);
}

function markdown(hits, patterns) {
  const rows = countsById(hits, patterns).map(([id, n]) => `| ${id} | ${n} |`);
  const body = patterns.flatMap(({ id }) => {
    const mine = hits.filter((h) => h.id === id && !h.allowed);
    if (mine.length === 0) return [];
    return [`### ${id} (${mine.length})`, "", ...mine.map((h) => `- \`${h.path}:${h.line}\` ${h.text.replace(/\|/g, "\\|").slice(0, 220)}`), ""];
  });
  const allowed = hits.filter((h) => h.allowed).map((h) => `- ${h.id} \`${h.path}:${h.line}\``);
  return ["| pattern | hits |", "|---|---|", ...rows, "", ...body, "### allowlisted", "", ...allowed, ""].join("\n");
}

// --probe <text> prints the id of every pattern the text trips, one per line (the guards'
// self-check that each tell is still caught).
function probe(patterns, text) {
  patterns.filter(({ regex }) => regex.test(text)).forEach(({ id }) => console.log(id));
}

// --skip <prefix> leaves out a tree another guard owns (the web vitest skips api/, which
// api/internal/slopguard covers).
function skippedPrefixes(argv) {
  return argv.flatMap((arg, index) => (arg === "--skip" && argv[index + 1] ? [argv[index + 1]] : []));
}

function main(argv) {
  const patterns = loadPatterns();
  const probeAt = argv.indexOf("--probe");
  if (probeAt !== -1) return probe(patterns, argv[probeAt + 1] ?? "");
  const skip = skippedPrefixes(argv);
  const files = copyFiles().filter((path) => !skip.some((prefix) => path.startsWith(prefix)));
  const hits = scan({ patterns, files });
  const open = hits.filter((h) => !h.allowed);
  if (argv.includes("--md")) process.stdout.write(markdown(hits, patterns));
  else if (argv.includes("--counts")) countsById(hits, patterns).forEach(([id, n]) => console.log(`${id}\t${n}`));
  else open.forEach((h) => console.log(`${h.path}:${h.line}: [${h.id}] ${h.text}`));
  if (argv.includes("--check") && open.length > 0) process.exitCode = 1;
}

if (process.argv[1] && fileURLToPath(import.meta.url) === process.argv[1]) main(process.argv.slice(2));
