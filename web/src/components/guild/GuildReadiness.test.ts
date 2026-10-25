// web/src/components/guild/GuildReadiness.test.ts
// The Readiness tab rendered with svelte/server: officer-read-only (officer tools off) shows the
// plan sentence once and the Read-only label, and keeps every Nudge on screen but disabled; a
// member sees none of it.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { buildMockReadiness } from '../../fixtures/guild/mock-guild';
import { guildGateCopy } from '../../lib/guild/copy';
import GuildReadiness from './GuildReadiness.svelte';

const SETTINGS_HREF = '/guild/us/hardcore/the-last-watch/#settings';

function html(officer: boolean, officerTools: boolean): string {
  return render(GuildReadiness, {
    props: {
      rows: buildMockReadiness().rows,
      officer,
      officerTools,
      myCharacterKey: null,
      settingsHref: SETTINGS_HREF,
    },
  }).body;
}

function text(markup: string): string {
  return markup
    .replace(/<!--[\s\S]*?-->/g, '')
    .replace(/<[^>]+>/g, ' ')
    .replace(/&nbsp;/g, ' ')
    .replace(/\s+/g, ' ')
    .trim();
}

function nudges(markup: string): string[] {
  return [...markup.matchAll(/<button[^>]*>/g)]
    .map((m) => m[0])
    .filter((tag) => tag.includes('data-testid="guild-readiness-nudge-'));
}

describe('GuildReadiness', () => {
  it('says the plan sentence once, under the h2 and above the header row, in officer-read-only', () => {
    const markup = html(true, false);
    expect(markup.split('data-testid="guild-readiness-plan-gate"').length - 1).toBe(1);
    expect(markup).toContain('id="officer-tools-gate-readiness"');
    const gateAt = markup.indexOf('data-testid="guild-readiness-plan-gate"');
    expect(gateAt).toBeGreaterThan(markup.indexOf('<h2'));
    expect(gateAt).toBeLessThan(markup.indexOf('data-testid="guild-readiness-header"'));
    const gate = markup.slice(gateAt, markup.indexOf('</p>', gateAt));
    expect(text(gate)).toContain(`${guildGateCopy.sentence.trim()} ${guildGateCopy.link}`);
    expect(gate).toContain(`href="${SETTINGS_HREF}"`);
  });

  it('shows the Read-only label beside the h2 in officer-read-only only', () => {
    expect(text(html(true, false))).toContain(`Readiness ${guildGateCopy.readOnly}`);
    expect(text(html(true, true))).not.toContain(guildGateCopy.readOnly);
    expect(text(html(false, false))).not.toContain(guildGateCopy.readOnly);
  });

  it('disables every Nudge, described by the sentence, and keeps the footnote', () => {
    const markup = html(true, false);
    const tags = nudges(markup);
    expect(tags.length).toBeGreaterThan(0);
    for (const tag of tags) {
      expect(tag).toMatch(/\sdisabled(=""|\s|>)/);
      expect(tag).toContain('aria-describedby="officer-tools-gate-readiness"');
      expect(tag).toContain('disabled:opacity-50');
      expect(tag).toContain('disabled:cursor-default');
    }
    expect(text(markup)).toContain('Nudge copies a message you can paste to them.');
  });

  it('keeps live Nudge with no sentence while officer tools are on', () => {
    const markup = html(true, true);
    expect(markup).not.toContain('guild-readiness-plan-gate');
    for (const tag of nudges(markup)) {
      expect(tag).not.toMatch(/\sdisabled(=""|\s|>)/);
      expect(tag).not.toContain('aria-describedby');
    }
  });

  it('leaves a member exactly as before: no sentence, no Nudge', () => {
    const markup = html(false, false);
    expect(markup).not.toContain('guild-readiness-plan-gate');
    expect(nudges(markup)).toHaveLength(0);
  });
});
