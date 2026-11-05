// web/src/lib/guild/professions.test.ts
import { describe, expect, it } from 'vitest';
import { professionName } from './professions';

describe('professionName', () => {
  it('reads the export slug as the client names it', () => {
    expect(professionName('first-aid')).toBe('First Aid');
    expect(professionName('leatherworking')).toBe('Leatherworking');
    expect(professionName('skinning')).toBe('Skinning');
  });

  it('capitalises each word of a slug it does not list', () => {
    expect(professionName('jewel-crafting')).toBe('Jewel Crafting');
  });
});
