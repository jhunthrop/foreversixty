<!-- web/src/components/logs/LogsPanel.svelte -->
<!-- One titled panel of the logs landing: the shared Card (site alignment spec 2026-10-29 §5.12)
     with its tile, its title and the muted aside as the head's meta, over the panel's own content.
     The card owns the padding; the content brings none. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import type { GlyphName } from '../../lib/ui/glyphs';
  import Card from '../ui/Card.svelte';

  let {
    title,
    glyph,
    aside = '',
    id,
    children,
  }: { title: string; glyph: GlyphName; aside?: string; id?: string; children: Snippet } = $props();
</script>

<Card {glyph} {title} meta={aside === '' ? undefined : aside} {id} class="scroll-mt-6">
  {@render children()}
</Card>
