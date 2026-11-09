<!-- web/src/components/ui/Checklist.svelte -->
<!-- Private checks (web layout spec 2026-11-05 §7.3): one open Row per check, a round mark for its
     state (done, to do, missing) and, on a check that is not done, a link to its fix. The member's
     pre-raid list, later "Before you raid" on the character page. A row with a fix is not itself a
     link: the fix is the row's one action (Row's href xor action rule). -->
<script lang="ts">
  import { LINK } from '../../lib/ui/primitives';
  import { MARK, type MarkTone } from '../../lib/ui/cells';
  import Row from './Row.svelte';

  let {
    items,
    testid,
  }: {
    items: readonly {
      key: string;
      tone: MarkTone;
      title: string;
      sub: string;
      fix: { label: string; href: string } | null;
    }[];
    testid?: string;
  } = $props();
</script>

<div data-testid={testid}>
  {#each items as item (item.key)}
    <Row
      title={item.title}
      sub={item.sub}
      open
      testid={testid === undefined ? undefined : `${testid}-${item.key}`}
    >
      {#snippet lead()}
        <span
          class="text-bg inline-flex size-[22px] flex-none items-center justify-center rounded-full text-[13px] leading-[22px] font-bold {MARK[
            item.tone
          ].ground}"
          role="img"
          aria-label={MARK[item.tone].label}>{MARK[item.tone].glyph}</span
        >
      {/snippet}
      {#snippet trail()}
        {#if item.fix}<a class={LINK} href={item.fix.href}>{item.fix.label} ›</a>{/if}
      {/snippet}
    </Row>
  {/each}
</div>
