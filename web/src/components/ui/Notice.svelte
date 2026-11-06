<!-- web/src/components/ui/Notice.svelte -->
<!-- One standing sentence under the page header (web layout spec 2026-11-05 §3A.2): a guild's
     standing line, a "before raids open" state. Text and one optional link, no box. `lead` puts a
     small mark before the sentence (the sim's Worth knowing lines lead with the spell's icon). -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { LINK } from '../../lib/ui/primitives';

  let {
    text,
    action,
    tone = 'gold',
    tail,
    lead,
    testid,
  }: {
    text: string;
    action?: { label: string; href: string };
    tone?: 'gold' | 'neutral';
    /** Words that finish the sentence after the text (a signed-out "Sign in to ..." clause). */
    tail?: Snippet;
    /** A mark before the text (a 20 px spell icon). */
    lead?: Snippet;
    testid?: string;
  } = $props();
</script>

<p
  class={[
    'ui-notice flex flex-wrap items-baseline gap-x-[14px] gap-y-[6px] pt-[14px] pb-1 text-[15px] leading-6',
    tone === 'gold' ? 'text-gold' : 'text-text',
  ]}
  data-testid={testid}
>
  {#if lead}<span class="inline-flex shrink-0 self-center">{@render lead()}</span>{/if}
  <span
    >{#if tail}{`${text} `}{@render tail()}{:else}{text}{/if}</span
  >
  {#if action}<a class={LINK} href={action.href}>{action.label} ›</a>{/if}
</p>
