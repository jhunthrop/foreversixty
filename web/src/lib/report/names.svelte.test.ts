import { describe, expect, it } from 'vitest';
import { homeRealmOf, playerLabel, realmLabel } from './names.svelte';

describe('homeRealmOf', () => {
  it('is the realm most players carry', () => {
    expect(
      homeRealmOf([
        'Ana-ClassicBetaPvP',
        'Bo-ClassicBetaPvP',
        'Cy-ClassicBetaPvP2',
        'Boss',
        'Di-ClassicBetaPvP',
      ]),
    ).toBe('ClassicBetaPvP');
  });
  it('is empty when no name carries one', () => {
    expect(homeRealmOf(['Ana', 'Bo'])).toBe('');
  });
});

describe('realmLabel', () => {
  it('spaces the normalised realm out', () => {
    expect(realmLabel('ClassicBetaPvP')).toBe('Classic Beta PvP');
    expect(realmLabel('ClassicBetaPvP2')).toBe('Classic Beta PvP 2');
    expect(realmLabel('Stormrage')).toBe('Stormrage');
  });
  it('prefers a name the characters data knows', () => {
    expect(realmLabel('ClassicBetaPvP2', new Map([['ClassicBetaPvP2', 'Classic Beta PvP 2 (EU)']]))).toBe(
      'Classic Beta PvP 2 (EU)',
    );
  });
});

describe('playerLabel', () => {
  it('shows the bare name for a player of the report’s own realm, with the full name on hover', () => {
    expect(playerLabel('Ana-ClassicBetaPvP', 'ClassicBetaPvP')).toEqual({
      name: 'Ana',
      realm: '',
      text: 'Ana',
      title: 'Ana-ClassicBetaPvP',
    });
  });
  it('adds the realm, spaced, for a cross-realm player', () => {
    expect(playerLabel('Cy-ClassicBetaPvP2', 'ClassicBetaPvP')).toEqual({
      name: 'Cy',
      realm: 'Classic Beta PvP 2',
      text: 'Cy · Classic Beta PvP 2',
      title: 'Cy-ClassicBetaPvP2',
    });
  });
  it('shows just the name when the report has no realm of its own or the log wrote none', () => {
    expect(playerLabel('Cy-ClassicBetaPvP2', '').text).toBe('Cy');
    expect(playerLabel('Cy', 'ClassicBetaPvP').text).toBe('Cy');
  });
  it('reads a two-word name out of the log’s First-Last-Realm- unit', () => {
    expect(playerLabel('Finnias-Falcon-ClassicBetaPvP-', 'ClassicBetaPvP')).toEqual({
      name: 'Finnias Falcon',
      realm: '',
      text: 'Finnias Falcon',
      title: 'Finnias-Falcon-ClassicBetaPvP-',
    });
    expect(playerLabel('На-Бутылку-ClassicBetaPvP2-', 'ClassicBetaPvP').text).toBe(
      'На Бутылку · Classic Beta PvP 2',
    );
  });
  it('counts the realm of a two-word unit, not its second word', () => {
    expect(homeRealmOf(['Bow-Jackzon-ClassicBetaPvP-', 'Finnias-Falcon-ClassicBetaPvP-', 'Boss'])).toBe(
      'ClassicBetaPvP',
    );
  });
});
