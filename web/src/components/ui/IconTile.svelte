<!-- web/src/components/ui/IconTile.svelte -->
<!-- A square glyph tile (polish spec 4.2): 36 or 32 px, control radius, the tile ground and ink.
     Tones: gold (default), discord, muted. Decorative: the row or head beside it carries the name. -->
<script lang="ts" module>
  export type TileTone = 'gold' | 'discord' | 'muted';
</script>

<script lang="ts">
  import type { GlyphName } from '../../lib/ui/glyphs';
  import Glyph from './Glyph.svelte';

  const TONES: Readonly<Record<TileTone, string>> = {
    gold: 'bg-tile text-tile-ink',
    discord: 'bg-[rgba(88,101,242,.16)] text-[#9aa4ff]',
    muted: 'bg-white/[.05] text-muted',
  };

  let { glyph, size = 36, tone = 'gold' }: { glyph: GlyphName; size?: 36 | 32; tone?: TileTone } = $props();
</script>

<span
  class="rounded-control inline-flex shrink-0 items-center justify-center {TONES[tone]}"
  style:width="{size}px"
  style:height="{size}px"
  aria-hidden="true"
  data-testid="icon-tile"
  data-tone={tone}
>
  <Glyph name={glyph} size={size === 36 ? 18 : 16} />
</span>
