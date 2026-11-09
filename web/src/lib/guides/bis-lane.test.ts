// web/src/lib/guides/bis-lane.test.ts
// @vitest-environment jsdom
// The Best in slot lane follows the answer's character (guides spec 2026-11-08 G5, bug B2: a
// Battle.net character of this class counts, not only a pasted or addon one).
import { afterEach, beforeEach, describe, expect, it } from 'vitest';
import type { CurrentCharacter } from '../current-character';
import { writeCurrent } from '../current-character';
import type { GuideCharacter } from './guide-answer';
import { bisPanelKey, mountBisLane, ownBand, retargetFullList } from './bis-lane';
import { publishGuideCharacter, resetGuideCharacter } from './guide-character';

const CODE = (trees: string): string => `FS1:1.60.1.70532:mage:gnome:${trees}:`;

function mage(level: number | undefined, faction: 'alliance' | 'horde' | undefined): GuideCharacter {
  return {
    name: 'Frosty',
    classSlug: 'mage',
    specSlug: 'frost',
    className: 'Mage',
    specName: 'Frost',
    level,
    faction,
    levelLine: level === undefined ? null : `Level ${level}`,
    identity: 'Undead Frost Mage',
  };
}

describe('ownBand (a stored addon or code build)', () => {
  it('turns the talent points a stored build spends into the level range it is in', () => {
    // 11 points: level 20.
    expect(ownBand(CODE('5/5/1'), 'mage')).toBe(20);
    // 24 points: level 33.
    expect(ownBand(CODE('55/55/4'), 'mage')).toBe(30);
  });

  it('is 60 at a full build and never above it', () => {
    expect(ownBand(CODE('2030005/113023/25351013010030025'), 'mage')).toBe(60);
  });

  it('never goes below the first range', () => {
    expect(ownBand(CODE('1'), 'mage')).toBe(20);
  });

  it('is undefined for another class or something that is not a code', () => {
    expect(ownBand(CODE('2030005'), 'warrior')).toBeUndefined();
    expect(ownBand('Simfury@Realm', 'mage')).toBeUndefined();
  });
});

describe('bisPanelKey (bug B2)', () => {
  it('gives a character of this class, Battle.net or not, its own band and faction', () => {
    expect(bisPanelKey(mage(34, 'horde'), 'mage', 'alliance')).toBe('30-horde');
    expect(bisPanelKey(mage(60, 'alliance'), 'mage', 'horde')).toBe('60-alliance');
  });

  it('falls back to the page’s faction when the character’s is unknown, and 60 when its level is', () => {
    expect(bisPanelKey(mage(22, undefined), 'mage', 'alliance')).toBe('20-alliance');
    expect(bisPanelKey(mage(undefined, 'horde'), 'mage', 'alliance')).toBe('60-horde');
  });

  it('shows anyone else the guide’s level 60 picks in the default faction', () => {
    expect(bisPanelKey(mage(34, 'horde'), 'warrior', 'alliance')).toBe('60-alliance');
    expect(bisPanelKey(null, 'mage', 'horde')).toBe('60-horde');
  });
});

describe('mountBisLane', () => {
  let root: HTMLElement;

  function panel(key: string): string {
    return `<template data-bis-panel="${key}"><p data-testid="panel-${key}">${key}</p></template>`;
  }

  beforeEach(() => {
    localStorage.clear();
    root = document.createElement('div');
    root.dataset.guideClass = 'mage';
    root.dataset.defaultFaction = 'alliance';
    root.innerHTML = `<div data-bis-current data-panel="60-alliance"><p>60-alliance</p></div>${[
      '30-horde',
      '20-alliance',
      '60-horde',
    ]
      .map(panel)
      .join('')}`;
    document.body.append(root);
  });

  afterEach(() => {
    resetGuideCharacter();
    root.remove();
  });

  const current = (): HTMLElement => root.querySelector<HTMLElement>('[data-bis-current]')!;

  it('swaps in the answer character’s band and faction', () => {
    mountBisLane(root);
    publishGuideCharacter(mage(34, 'horde'));
    expect(current().dataset.panel).toBe('30-horde');
    expect(current().querySelector('[data-testid="panel-30-horde"]')).not.toBeNull();
  });

  it('opens a stored code build of this class on its band when the answer names nobody', () => {
    const pointer: CurrentCharacter = {
      source: 'code',
      ref: CODE('5/5/1'),
      label: 'Frosty · Frost Mage',
      classSlug: 'mage',
      savedAt: '2026-11-06T00:00:00Z',
    };
    writeCurrent(pointer);
    mountBisLane(root);
    publishGuideCharacter(null);
    expect(current().dataset.panel).toBe('20-alliance');
  });

  it('returns to the band the server drew when the character moves on to 60', () => {
    mountBisLane(root);
    publishGuideCharacter(mage(34, 'horde'));
    expect(current().dataset.panel).toBe('30-horde');
    publishGuideCharacter(mage(60, 'alliance'));
    expect(current().dataset.panel).toBe('60-alliance');
    expect(current().textContent).toBe('60-alliance');
  });

  it('keeps the panel it has when the key has no template', () => {
    mountBisLane(root);
    publishGuideCharacter(mage(44, 'alliance'));
    expect(current().dataset.panel).toBe('60-alliance');
    expect(current().textContent).toBe('60-alliance');
  });
});

describe('retargetFullList', () => {
  it('points the full list at the shown band and faction', () => {
    const root = document.createElement('div');
    root.dataset.specSlug = 'survival';
    root.innerHTML = '<a data-testid="guide-bis-full-list" href="/bis/hunter/survival">Full list</a>';
    retargetFullList(root, 'hunter', '30-horde');
    expect(root.querySelector('a')?.getAttribute('href')).toBe('/bis/hunter/survival/30-39?faction=horde');
  });
});
