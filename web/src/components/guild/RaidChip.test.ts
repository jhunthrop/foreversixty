// web/src/components/guild/RaidChip.test.ts
// The spec 3.6 chip table, rendered with svelte/server: text and recipe per status and per
// viewer answer.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import RaidChip from './RaidChip.svelte';

type ChipProps = Record<string, unknown>;

function chip(props: ChipProps): { text: string; tag: string } {
  const body = render(RaidChip, { props: props as never }).body;
  const tag = body.match(/<span[^>]*>/)?.[0] ?? '';
  const text = body
    .replace(/<!--[\s\S]*?-->/g, '')
    .replace(/<[^>]+>/g, '')
    .trim();
  return { text, tag };
}

const KILL = 'pill-community';

describe('RaidChip status', () => {
  it.each([
    ['planned', 'Planned', 'pill-site'],
    ['locked', 'Locked', 'pill-sample'],
    ['done', 'Done', KILL],
    ['ended', 'Ended', 'pill-sample'],
  ])('%s reads %s on %s', (status, text, recipe) => {
    const out = chip({ kind: 'status', status });
    expect(out.text).toBe(text);
    expect(out.tag).toContain('pill ');
    expect(out.tag).toContain(recipe);
  });

  it('draws cancelled in the failure colour', () => {
    const out = chip({ kind: 'status', status: 'cancelled' });
    expect(out.text).toBe('Cancelled');
    expect(out.tag).toContain('text-wipe');
    expect(out.tag).toContain('bg-wipe/14');
    expect(out.tag).toContain('border-wipe/35');
  });
});

describe('RaidChip viewer', () => {
  it('reads In and Late with the role, on the kill recipe', () => {
    expect(chip({ kind: 'viewer', state: 'in', role: 'dps' })).toMatchObject({ text: 'In · DPS' });
    expect(chip({ kind: 'viewer', state: 'late', role: 'healer' }).text).toBe('Late · Healer');
    expect(chip({ kind: 'viewer', state: 'in', role: 'tank' }).tag).toContain(KILL);
  });

  it('reads Bench and Out without a role, on the sample recipe', () => {
    const bench = chip({ kind: 'viewer', state: 'bench', role: 'dps' });
    expect(bench.text).toBe('Bench');
    expect(bench.tag).toContain('pill-sample');
    expect(chip({ kind: 'viewer', state: 'out', role: 'dps' }).text).toBe('Out');
  });

  it('reads Not signed up in gold while unsigned', () => {
    const out = chip({ kind: 'viewer', state: null });
    expect(out.text).toBe('Not signed up');
    expect(out.tag).toContain('pill-site');
  });
});
