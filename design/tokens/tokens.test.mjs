// design/tokens/tokens.test.mjs: the token outputs (web, iOS, Android, the companion, the addon)
// must be what design/tokens/build.mjs writes from tokens.json today, and each must open with the
// GENERATED line (a hand edit that removed it fails too). Run: node --test
// design/tokens/tokens.test.mjs (the web unit job and mobile-codegen.yml).
import assert from 'node:assert/strict';
import { existsSync, readdirSync, readFileSync } from 'node:fs';
import { join, relative } from 'node:path';
import { describe, it } from 'node:test';
import {
  ADDON_FONT_DIR,
  COMMAND,
  HEADER,
  OUTPUTS,
  SHARED_FONT_DIR,
  flatten,
  generate,
  gradientStops,
  parseColor,
} from './build.mjs';

const generated = generate();
const repoPath = (file) => relative(new URL('../..', import.meta.url).pathname, file);

describe('design tokens', () => {
  for (const [key, file] of Object.entries(OUTPUTS)) {
    const onDisk = readFileSync(file, 'utf8');

    it(`${repoPath(file)} opens with the GENERATED line`, () => {
      const first = onDisk.split('\n', 1)[0];
      assert.ok(
        first.includes(HEADER),
        `${repoPath(file)} lost its first line "${HEADER}"; run \`${COMMAND}\``,
      );
    });

    it(`${repoPath(file)} matches tokens.json`, () => {
      assert.equal(
        onDisk,
        generated[key],
        `${repoPath(file)} is stale or hand-edited; edit design/tokens/tokens.json, run \`${COMMAND}\` and commit the result`,
      );
    });
  }

  it('names every group the spec lists', () => {
    const doc = JSON.parse(readFileSync(new URL('./tokens.json', import.meta.url), 'utf8'));
    const { tokens } = flatten(doc);
    const paths = new Set(tokens.map((t) => t.path.join('.')));
    for (const path of [
      'color.class.warrior',
      'color.rarity.epicText',
      'color.parse.greyText',
      'color.faction.horde',
      'color.tier.S',
      'color.school.frost',
      'font.title',
      'type.label',
      'space.4',
      'size.rowH',
      'radius.plate',
      'stroke.hairline',
      'elevation.overlay',
      'motion.base',
    ]) {
      assert.ok(paths.has(path), path);
    }
  });

  it('parses hex and rgba colours', () => {
    assert.deepEqual(parseColor('#abc'), { rgb: 0xaabbcc, alpha: 1 });
    assert.deepEqual(parseColor('rgba(229, 185, 85, 0.17)'), {
      rgb: 0xe5b955,
      alpha: 0.17,
    });
    assert.throws(() => parseColor('var(--x)'));
  });
});

describe('the band fade', () => {
  const doc = JSON.parse(readFileSync(new URL('./tokens.json', import.meta.url), 'utf8'));
  const byPath = new Map(flatten(doc).tokens.map((t) => [t.path.join('.'), t]));
  const bandCss = readFileSync(new URL('../../web/src/components/ui/band.css', import.meta.url), 'utf8');

  it('holds the web mask ramp to gradient.bandFade', () => {
    const mask = /mask-image:\s*linear-gradient\(\s*to bottom,([\s\S]*?)\);/.exec(bandCss);
    assert.ok(mask, 'web/src/components/ui/band.css lost its fade mask-image');
    const webStops = mask[1].split(/,\s*\n/).map((stop) => {
      const s = stop.trim();
      if (s.startsWith('black')) return { alpha: 1, position: 0 };
      if (s === 'transparent') return { alpha: 0, position: 1 };
      const m = /^rgb\(0 0 0 \/ ([\d.]+)\) calc\(100% - var\(--band-fade\) \* ([\d.]+)\)$/.exec(s);
      assert.ok(m, `unexpected stop in band.css: ${s}`);
      return { alpha: Number(m[1]), position: 1 - Number(m[2]) };
    });
    const tokenStops = gradientStops(byPath.get('gradient.bandFade')).map(({ alpha, position }) => ({
      alpha,
      position,
    }));
    assert.deepEqual(webStops, tokenStops);
  });

  it('holds the web fade lengths to size.bandFade', () => {
    const t = byPath.get('size.bandFade');
    const [desktop, phone] = [...bandCss.matchAll(/--band-fade:\s*(\d+px);/g)].map((m) => m[1]);
    assert.equal(desktop, t.value);
    assert.equal(phone, t.ext.phone);
  });
});

describe('the addon fonts', () => {
  const doc = JSON.parse(readFileSync(new URL('./tokens.json', import.meta.url), 'utf8'));
  const shipped = flatten(doc).tokens.filter((t) => t.top === 'font' && t.ext.lua);

  it('ships a face for Open Sans and Marcellus', () => {
    assert.deepEqual(shipped.map((t) => t.name).sort(), ['body', 'title']);
  });

  for (const t of shipped) {
    for (const postScript of Object.values(t.ext.faces)) {
      it(`ships ${postScript}.ttf unaltered, with its licence, and names a client fallback`, () => {
        const file = join(ADDON_FONT_DIR, `${postScript}.ttf`);
        assert.ok(existsSync(file), `copy apps/shared/fonts/${postScript}.ttf into addon/ForeverSixty/media/fonts`);
        assert.ok(
          readFileSync(file).equals(readFileSync(join(SHARED_FONT_DIR, `${postScript}.ttf`))),
          `${postScript}.ttf differs from apps/shared/fonts: ship the upstream TTF, never a subset`,
        );
        assert.match(t.ext.lua.fallback, /^Fonts\\[A-Z_]+\.TTF$/);
      });
    }
  }

  it('ships each family\'s OFL licence beside its faces', () => {
    for (const name of ['OFL-Marcellus.txt', 'OFL-OpenSans.txt']) {
      assert.ok(existsSync(join(ADDON_FONT_DIR, name)), name);
    }
  });
});

describe('the Android XML colours', () => {
  const res = new URL('../../apps/android/app/src/main/res/', import.meta.url).pathname;
  // Every colour a values file or a shape drawable sets comes from colors_gen.xml (@color/fs_*). The
  // ic_* vectors are white glyph masks the code tints, so they are not colours of their own.
  const files = ['values', 'drawable'].flatMap((dir) =>
    readdirSync(join(res, dir))
      .filter((name) => name.endsWith('.xml') && !name.startsWith('ic_'))
      .map((name) => join(res, dir, name)),
  );

  it('writes no colour value outside the generated file', () => {
    const found = files.flatMap((file) =>
      [...readFileSync(file, 'utf8').matchAll(/#[0-9A-Fa-f]{6,8}\b/g)].map((m) => `${repoPath(file)}: ${m[0]}`),
    );
    assert.deepEqual(found, []);
  });

  it('opens the window, so the splash, on the token ground', () => {
    const themes = readFileSync(join(res, 'values/themes.xml'), 'utf8');
    assert.match(themes, /android:windowBackground">@color\/fs_bg</);
  });
});
