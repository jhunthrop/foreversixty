<!-- web/src/components/report/ReportSkeleton.svelte -->
<!-- The report while it loads (logs spec 2026-11-07 6.3 "Loading"): the ready page's own frame in
     placeholder blocks, the same geometry so nothing shifts when the figures land. The head (the
     title, or a 520 x 40 block while the shell does not know it; a 320 x 16 meta; the answer's
     three lines), one disabled Summary tab, the bar's three selects, the Players Plate's six 56 px
     rows and the rail's This pull as five lines. Rendered twice from this one component: by the
     page shell before the island boots (Astro, no hydration) and by the island until the first
     summary arrives. The title's testid is the one the Worker fills (src/worker.ts). -->
<script lang="ts">
  import type { ControlGroup } from '../../lib/ui/primitives';
  import { reportCopy } from '../../lib/report/copy';
  import { REPORT_SHELL_TITLE_TESTID } from '../../lib/report/skeleton';
  import ControlBar from '../ui/ControlBar.svelte';
  import Page from '../ui/Page.svelte';
  import PageHeader from '../ui/PageHeader.svelte';
  import Plate from '../ui/Plate.svelte';
  import Section from '../ui/Section.svelte';
  import TabStripMore from '../ui/TabStripMore.svelte';

  let { title }: { title?: string } = $props();

  const placeholder = (param: string, label: string): ControlGroup => ({
    label,
    param,
    kind: 'select',
    value: '',
    options: [{ value: '', label: '…' }],
  });
  const GROUPS: readonly ControlGroup[] = [
    placeholder('fight', reportCopy.fight),
    placeholder('source', reportCopy.source),
    placeholder('window', reportCopy.window),
  ];
  const ROWS = [0, 1, 2, 3, 4, 5];
</script>

{#snippet block(width: string, height: string)}
  <span class="skeleton-block block {width} {height} max-w-full" aria-hidden="true"></span>
{/snippet}

<div class="report-skeleton" data-testid="report-skeleton" aria-busy="true">
  <Page layout="split">
    {#snippet header()}
      <PageHeader title={title ?? ' '} titleTestid={REPORT_SHELL_TITLE_TESTID}>
        {#snippet meta()}{@render block('w-[320px]', 'h-4')}{/snippet}
        {#snippet answer()}
          <span class="hidden min-[900px]:block min-[900px]:w-[360px]" aria-hidden="true"></span>
          {@render block('w-[200px]', 'h-4')}
          {@render block('w-[160px]', 'h-10')}
          {@render block('w-[220px]', 'h-4')}
        {/snippet}
        {#snippet tabs()}
          <TabStripMore
            tabs={[{ id: 'summary', label: 'Summary', href: '#', disabled: true }]}
            active=""
            md={1}
            xl={1}
            label={reportCopy.tabsLabel}
          />
        {/snippet}
      </PageHeader>
    {/snippet}
    {#snippet bar()}
      <ControlBar groups={GROUPS} label="Report choices" summary="…" labelsFrom={1920} staticField />
    {/snippet}
    {#snippet rail()}
      <Section title="This pull">
        <div class="-mt-3">
          {#each [0, 1, 2, 3, 4] as line (line)}<p>{@render block('w-[300px]', 'h-4')}</p>{/each}
        </div>
      </Section>
    {/snippet}
    <Section title="Players">
      <Plate label="Loading the report.">
        {#each ROWS as row (row)}
          <div class="flex h-14 items-center gap-3 px-4" aria-hidden="true">
            {@render block('w-7 rounded-full', 'h-7')}
            {@render block('w-[180px]', 'h-4')}
            <span class="ml-auto flex gap-6">
              {#each [0, 1, 2, 3] as cell (cell)}{@render block('w-12', 'h-4')}{/each}
            </span>
          </div>
        {/each}
      </Plate>
    </Section>
  </Page>
</div>
