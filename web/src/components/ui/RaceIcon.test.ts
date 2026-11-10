import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import RaceIcon from './RaceIcon.svelte';

function html(props: { url: string; name: string; size?: 20 | 24 | 28 | 40 }): string {
  return render(RaceIcon, { props }).body.replace(/<!--[^>]*-->/g, '');
}

describe('RaceIcon', () => {
  const url = '/data/1.60.1.70790/icons/achievement_character_nightelf_male.webp';

  it('draws the portrait in the icon frame, titled with the race', () => {
    const body = html({ url, name: 'Night Elf' });
    expect(body).toContain('title="Night Elf"');
    expect(body).toContain(`src="${url}"`);
    expect(body).toContain('border-line-warm-strong');
    expect(body).toContain('width: 28px');
  });

  it('is decorative beside the name it stands for', () => {
    const body = html({ url, name: 'Undead', size: 20 });
    expect(body).toContain('aria-hidden="true"');
    expect(body).toContain('alt=""');
    expect(body).toContain('width="20"');
  });
});
