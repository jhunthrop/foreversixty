<!-- web/src/components/ui/Notice.svelte -->
<!-- One standing sentence under the page header (web layout spec 2026-11-05 §3A.2): a guild's
     standing line, a "before raids open" state. Text and one optional link, no box. -->
<script lang="ts">
  import { LINK } from '../../lib/ui/primitives';

  let {
    text,
    action,
    tone = 'gold',
    testid,
  }: {
    text: string;
    action?: { label: string; href: string };
    tone?: 'gold' | 'neutral';
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
  <span>{text}</span>
  {#if action}<a class={LINK} href={action.href}>{action.label} ›</a>{/if}
</p>
