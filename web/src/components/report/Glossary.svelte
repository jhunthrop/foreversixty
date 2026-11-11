<!-- web/src/components/report/Glossary.svelte -->
<!-- The report's help (logs spec 2026-11-07, coordinator ruling: the glossary stays): a quiet
     "What these mean ›" line at the foot of the main column that opens the words the tables use
     in place, the shared Disclosure (a native details element) that reads as what it is. No box. The terms are their own chunk,
     fetched the first time it opens, so a reader who never asks pays nothing. A link to an entry
     (the threat chart's "What it leaves out") opens it and scrolls to the entry. -->
<script lang="ts">
  import type { GlossaryTerm } from '../../lib/report/glossary';
  import { REGION } from '../../lib/ui/regions';
  import Disclosure from '../ui/Disclosure.svelte';

  let open = $state(false);
  let terms = $state<readonly GlossaryTerm[]>([]);
  const loadTerms = async (): Promise<readonly GlossaryTerm[]> => {
    if (terms.length === 0) terms = (await import('../../lib/report/glossary')).GLOSSARY;
    return terms;
  };
  $effect(() => {
    if (open) void loadTerms();
  });
  $effect(() => {
    const follow = (): void => {
      const id = window.location.hash.slice(1);
      if (!id.startsWith('glossary-')) return;
      void loadTerms().then((list) => {
        if (!list.some((entry) => entry.id === id)) return;
        open = true;
        requestAnimationFrame(() => document.getElementById(id)?.scrollIntoView({ block: 'center' }));
      });
    };
    follow();
    window.addEventListener('hashchange', follow);
    return () => window.removeEventListener('hashchange', follow);
  });
</script>

<div data-region={REGION.footerNote}>
  <Disclosure bind:open summary={open ? 'Hide what these mean' : 'What these mean'} link testid="glossary">
    <dl class="grid grid-cols-1 gap-x-6 gap-y-3 pb-2 text-[14px] leading-5 md:grid-cols-2">
      {#each terms as entry (entry.term)}
        <div class="flex scroll-mt-4 flex-col gap-0.5" id={entry.id}>
          <dt class="text-strong font-semibold">{entry.term}</dt>
          <dd class="text-muted m-0">{entry.meaning}</dd>
        </div>
      {/each}
    </dl>
  </Disclosure>
</div>
