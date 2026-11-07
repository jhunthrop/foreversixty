<!-- web/src/components/ui/Toc.svelte -->
<!-- The contents rail of a reading page (guides spec 2026-11-08 §3A): a label, then one link per
     section, 14/18 600 in --color-nav on 32 px rows with a 2 px --color-line-soft bar; the current
     one gold with a gold bar (`aria-current="location"`, kept by the page's observer). -->
<script lang="ts">
  let {
    label,
    items,
    current,
    testid,
  }: {
    label: string;
    items: readonly { title: string; id: string }[];
    current?: string;
    testid?: string;
  } = $props();
</script>

<div class="ui-toc" data-testid={testid}>
  <span class="text-label tracking-label text-muted block pb-[10px] font-bold uppercase">{label}</span>
  <ol>
    {#each items as item (item.id)}
      <li>
        <a
          href={`#${item.id}`}
          data-toc-link={item.id}
          aria-current={item.id === current ? 'location' : undefined}>{item.title}</a
        >
      </li>
    {/each}
  </ol>
</div>

<style>
  .ui-toc ol {
    list-style: none;
    margin: 0;
    padding: 0;
    display: flex;
    flex-direction: column;
  }
  .ui-toc a {
    display: flex;
    align-items: center;
    min-height: 32px;
    padding-left: 12px;
    font-size: 14px;
    line-height: 18px;
    font-weight: 600;
    color: var(--color-nav);
    box-shadow: inset 2px 0 0 var(--color-line-soft);
  }
  .ui-toc a:hover {
    color: var(--color-strong);
  }
  .ui-toc a:focus-visible {
    outline: 2px solid var(--color-gold);
    outline-offset: 2px;
  }
  .ui-toc a[aria-current] {
    color: var(--color-gold);
    box-shadow: inset 2px 0 0 var(--color-gold);
  }
</style>
