// Both apps' strings are generated from apps/shared/copy/en.json (web/scripts/gen-mobile-copy.mjs).
// This fails when a checked-in catalog differs from a fresh generation: after editing en.json run
// `node web/scripts/gen-mobile-copy.mjs` and commit the result.
import { readFileSync } from 'node:fs';
import { describe, expect, it } from 'vitest';
import {
  buildCatalog,
  generate,
  goDeclaration,
  parseEntry,
  renderAndroid,
  renderGo,
  renderXcstrings,
  swiftAccessor,
} from '../../../scripts/gen-mobile-copy.mjs';

const HINT = 'run `node web/scripts/gen-mobile-copy.mjs` and commit the result';

describe('mobile copy codegen', () => {
  const generated = generate() as Record<string, string>;

  it.each(Object.keys(generated))('%s matches the generator', (file) => {
    expect(readFileSync(file, 'utf8'), HINT).toBe(generated[file]);
  });

  it('numbers placeholders by first appearance, ints as integers', () => {
    const entry = parseEntry('raid.count', {
      text: '{n} of {size} signed up',
      args: { n: 'int', size: 'int' },
    });
    expect(renderAndroid([entry], '// x')).toContain(
      '<string name="raid_count">%1$d of %2$d signed up</string>',
    );
    expect(renderXcstrings([entry])).toContain('"value": "%1$lld of %2$lld signed up"');
  });

  it('writes plurals for both platforms', () => {
    const entry = parseEntry('talents.unspent', {
      count: 'n',
      one: '1 point unspent.',
      other: '{n} points unspent.',
    });
    expect(renderAndroid([entry], '// x')).toContain('<item quantity="other">%1$d points unspent.</item>');
    const ios = JSON.parse(renderXcstrings([entry])).strings['talents.unspent'].localizations.en;
    expect(ios.variations.plural.other.stringUnit.value).toBe('%1$lld points unspent.');
  });

  it('escapes a placeholder named after a Swift keyword', () => {
    const swift = swiftAccessor(
      parseEntry('home.raid.counts', {
        text: '{in} in, {late} late',
        notPlural: true,
        args: { in: 'int', late: 'int' },
      }),
    );
    expect(swift).toContain('static func homeRaidCounts(`in`: Int, late: Int) -> String');
    expect(swift).toContain('format("home.raid.counts", `in`, late)');
    expect(swiftAccessor(parseEntry('a.b', 'Hi {class}'))).toContain('(className: String)');
  });

  it('keeps edge spaces and escapes apostrophes for aapt', () => {
    const android = renderAndroid([parseEntry('a.b', " · it's {x}")], '// x');
    expect(android).toContain(`<string name="a_b">" · it\\'s %1$s"</string>`);
  });

  it('keeps shared.* keys in shared.json and only there', () => {
    const app = (key: string) => JSON.stringify({ strings: { [key]: 'x' } });
    expect(buildCatalog(app('a.b'), app('shared.a.b')).map((e: { key: string }) => e.key)).toEqual([
      'a.b',
      'shared.a.b',
    ]);
    expect(() => buildCatalog(app('shared.a.b'))).toThrow(/live in apps\/shared\/copy\/shared.json/);
    expect(() => buildCatalog(app('a.b'), app('a.c'))).toThrow(/keys start shared\./);
  });

  it('refuses mismatched placeholders, stray plural args and banned words', () => {
    expect(() => parseEntry('a.b', { ios: 'Hi {name}', android: 'Hi' })).toThrow(/same placeholders/);
    expect(() => parseEntry('a.b', { count: 'n', one: '1 {x}', other: '{n} {x}' })).toThrow(/names only/);
    expect(() => parseEntry('a.b', 'Go Premium')).toThrow(/banned/);
    expect(() => parseEntry('a.b', 'Done!')).toThrow(/banned/);
    expect(() => parseEntry('A.b', 'x')).toThrow(/lowerCamel/);
  });

  it('keeps server.json keys to the server namespaces, out of both apps and the site', () => {
    const one = (key: string) => JSON.stringify({ strings: { [key]: 'x' } });
    const entries = buildCatalog(one('a.b'), one('shared.a.b'), one('discord.a.b'));
    expect(entries.map((e: { key: string }) => e.key)).toEqual(['a.b', 'shared.a.b', 'discord.a.b']);
    expect(() => buildCatalog(one('a.b'), undefined, one('a.c'))).toThrow(/keys start discord\./);
    expect(() => buildCatalog(one('mail.a'))).toThrow(/live in apps\/shared\/copy\/server.json/);
    const outputs = generate() as Record<string, string>;
    const client = Object.entries(outputs).filter(([file]) => !file.endsWith('.go'));
    for (const [file, text] of client) expect(text, file).not.toContain('discord.common.unbound');
  });

  it('lets a server string carry a bare percent but never a printf verb', () => {
    expect(parseEntry('discord.a', '{chance}% from {from}', { server: true }).key).toBe('discord.a');
    expect(() => parseEntry('discord.a', 'Your %s sim', { server: true })).toThrow(/% is not allowed/);
    expect(() => parseEntry('a.b', '{chance}% from {from}')).toThrow(/% is not allowed/);
  });

  it('writes one Go constant per key and a Plural per count', () => {
    expect(goDeclaration(parseEntry('discord.press.locked', 'Sign-ups are "locked".'))).toBe(
      '// DiscordPressLocked is discord.press.locked: Sign-ups are "locked".\nconst DiscordPressLocked = "Sign-ups are \\"locked\\"."',
    );
    expect(
      goDeclaration(parseEntry('push.event.kills', { count: 'n', one: '1 kill', other: '{n} kills' })),
    ).toContain('var PushEventKills = Plural{Count: "n", One: "1 kill", Other: "{n} kills"}');
    expect(goDeclaration(parseEntry('a.b', { ios: 'Tap', android: 'Press' }))).toBeNull();
    const go = renderGo(
      buildCatalog('{"strings":{}}', undefined, JSON.stringify({ strings: { 'mail.a': 'x' } })),
    );
    const server = Object.entries(go).find(([file]) => file.endsWith('server_gen.go'))?.[1] ?? '';
    expect(server.startsWith('// Code generated by web/scripts/gen-mobile-copy.mjs')).toBe(true);
    expect(server).toContain('DO NOT EDIT.\n\npackage copygen\n');
    expect(server).toContain('const MailA = "x"');
  });

  it('refuses two keys that make one Go name', () => {
    const app = JSON.stringify({ strings: { 'a.bC': 'x', 'a.b.c': 'y' } });
    expect(() => buildCatalog(app)).toThrow(/share the Go name ABC/);
  });
});
