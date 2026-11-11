import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import SpellIcon from './SpellIcon.svelte';
import { ICON_HIDE_ON_FAIL } from '../../lib/ui/icon-fail';

type Props = Parameters<typeof SpellIcon>[1];

function html(props: Props): string {
  return render(SpellIcon, { props }).body.replace(/<!--[^>]*-->/g, '');
}

const URL = '/data/b/icons/spell_holy_heal.webp';

describe('SpellIcon', () => {
  it('draws the art in the warm frame at its size, decorative beside a printed name', () => {
    const body = html({ url: URL, size: 24 });
    expect(body).toContain(`src="${URL}"`);
    expect(body).toContain('border-rule-warm');
    expect(body).toContain('width: 24px');
    expect(body).toContain('aria-hidden="true"');
    expect(body).toContain(ICON_HIDE_ON_FAIL);
  });

  it('draws nothing for no icon, unless it holds a column', () => {
    expect(html({ url: null })).toBe('');
    const socket = html({ url: null, socket: true });
    expect(socket).toContain('data-icon-socket');
    expect(socket).not.toContain('<img');
    expect(socket).not.toContain(ICON_HIDE_ON_FAIL);
  });

  it('is an unframed glyph at 16 px and below', () => {
    expect(html({ url: URL, size: 16 })).not.toContain('border-rule-warm');
    expect(html({ url: URL, size: 18 })).toContain('border-rule-warm');
  });

  it("opens the spell's or the item's tooltip and names an icon that stands alone", () => {
    expect(html({ url: URL, tip: 2050 })).toContain('data-game-tip="spell:2050"');
    expect(html({ url: URL, item: 19019 })).toContain('data-bis-item="19019"');
    const named = html({ url: URL, label: 'Tank' });
    expect(named).toContain('role="img"');
    expect(named).toContain('aria-label="Tank"');
  });

  it('holds the skeleton while its table loads', () => {
    expect(html({ url: null, loading: true })).toContain('skeleton-block');
  });
});
