<!-- web/src/components/sim/SimLanding.svelte -->
<!-- /sim with no character (spec 2026-11-06-sim §4.2, mock web-sim): the paste box on the page's
     framed panel, three other ways in as rows on a panel, and the "Sim a best-in-slot set" panel (ruling S6:
     the published BiS sets, no figures). sim.astro server-renders it as the page's first paint and
     the island mounts the same component in its place, so nothing moves on hydration. Without
     handlers (the static render) every form still submits to /sim the plain way. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { classTone } from '../../lib/ui/primitives';
  import { validateColdPaste } from '../../lib/sim/cold-paste';
  import { simPageCopy as copy } from '../../lib/sim/page-copy';
  import type { BisSetClass } from '../../lib/sim/bis-set-rail';
  import Crest from '../ui/Crest.svelte';
  import Button from '../ui/Button.svelte';
  import Field from '../ui/Field.svelte';
  import FormRow from '../ui/FormRow.svelte';
  import Page from '../ui/Page.svelte';
  import Row from '../ui/Row.svelte';
  import RowGrid from '../ui/RowGrid.svelte';
  import Section from '../ui/Section.svelte';

  let {
    bisSets,
    busy = false,
    pasteError = '',
    message = null,
    onpaste,
    onbuild,
    onfight,
    onsignin,
    railTail,
  }: {
    bisSets: readonly BisSetClass[];
    /** A character is loading: Run reads busy. */
    busy?: boolean;
    /** The paste did not decode. */
    pasteError?: string;
    /** A source's refusal (a build or fight that would not load). */
    message?: string | null;
    onpaste?: (code: string) => void;
    onbuild?: (ref: string) => void;
    onfight?: (ref: string) => void;
    onsignin?: () => void;
    /** The rail's tail after the best-in-slot sets (Your sims, signed in). */
    railTail?: Snippet;
  } = $props();

  let paste = $state('');
  let buildRef = $state('');
  let fightRef = $state('');

  // The paste decodes locally (§5 "Paste plate"): Run waits for a paste that decodes, and one that
  // does not says why under the field. The static render (no handler) leaves the plain form be.
  const pasteCheck = $derived(validateColdPaste(paste));
  const pasteMessage = $derived(
    pasteError !== '' ? pasteError : onpaste !== undefined && !pasteCheck.ok ? pasteCheck.message : null,
  );
  const pasteBlocked = $derived(onpaste !== undefined && !pasteCheck.ok);

  function submitPaste(event: SubmitEvent): void {
    if (onpaste === undefined) return;
    event.preventDefault();
    if (pasteCheck.ok) onpaste(pasteCheck.code);
  }

  function handle(event: SubmitEvent, value: string, handler: ((value: string) => void) | undefined): void {
    if (handler === undefined) return;
    event.preventDefault();
    if (value.trim() !== '') handler(value.trim());
  }
</script>

<Page panels testid="sim-landing">
  <Section title={copy.sections.simYourCharacter} panel testid="sim-cold-paste">
    <div class="flex flex-col gap-3">
      <FormRow label={copy.paste.label} action="/sim" method="get" inline onsubmit={submitPaste}>
        <Field
          id="sim-cold-paste-input"
          name="code"
          label={copy.paste.label}
          placeholder={copy.paste.placeholder}
          labelHidden
          grow
          bind:value={paste}
          error={pasteMessage ?? undefined}
          testid="sim-cold-paste-input"
        />
        <!-- A Run pressed before the island mounts submits the plain way; run=1 makes that
             arrival run the pasted character too, as the island's Run does. -->
        <input type="hidden" name="run" value="1" />
        {#snippet submit()}
          <Button tone="primary" type="submit" disabled={busy || pasteBlocked} testid="sim-cold-paste-run"
            >{copy.paste.run}</Button
          >
        {/snippet}
      </FormRow>
      <p class="text-small text-muted m-0">
        {copy.paste.hint}
        <a class="text-gold font-semibold whitespace-nowrap" href={copy.paste.setUpHref}>{copy.paste.setUp}</a
        >
      </p>
    </div>
  </Section>
  <Section title={copy.sections.otherWaysIn} panel testid="sim-sources">
    <div>
      <Row title={copy.ways.characters} open fillTrail>
        {#snippet trail()}
          <Button href={onsignin ? undefined : '/account'} onclick={onsignin} testid="sim-sign-in"
            >{copy.ways.signIn}</Button
          >
        {/snippet}
      </Row>
      <Row title={copy.ways.build} open fillTrail>
        {#snippet trail()}
          <FormRow
            label={copy.ways.build}
            action="/sim"
            method="get"
            inline
            onsubmit={(event) => handle(event, buildRef, onbuild)}
          >
            <Field
              id="sim-build-input"
              name="ref"
              label={copy.ways.build}
              placeholder={copy.ways.buildPlaceholder}
              labelHidden
              grow
              bind:value={buildRef}
              testid="sim-build-input"
            />
            {#snippet submit()}
              <Button type="submit" disabled={busy} testid="sim-build-load">{copy.ways.load}</Button>
            {/snippet}
          </FormRow>
        {/snippet}
      </Row>
      <Row title={copy.ways.fight} open fillTrail>
        {#snippet trail()}
          <FormRow
            label={copy.ways.fight}
            action="/sim"
            method="get"
            inline
            onsubmit={(event) => handle(event, fightRef, onfight)}
          >
            <Field
              id="sim-fight-input"
              name="ref"
              label={copy.ways.fight}
              placeholder={copy.ways.fightPlaceholder}
              labelHidden
              grow
              bind:value={fightRef}
              testid="sim-fight-input"
            />
            {#snippet submit()}
              <Button type="submit" disabled={busy} testid="sim-fight-load">{copy.ways.load}</Button>
            {/snippet}
          </FormRow>
        {/snippet}
      </Row>
    </div>
    {#if message}
      <p class="text-small text-wipe mt-2" role="alert" data-testid="sim-landing-message">{message}</p>
    {/if}
  </Section>
  <Section title={copy.sections.bisSet} panel testid="sim-bis-sets">
    <RowGrid cols={3}>
      {#each bisSets as klass (klass.slug)}
        <Row title={klass.name} titleTone={classTone(klass.slug)} open>
          {#snippet lead()}<Crest cls={klass.slug} size={32} />{/snippet}
          {#snippet subContent()}
            <!-- On a phone each spec link takes its own line (no orphans, no collisions); one line from 768. -->
            <span class="flex flex-col pt-[2px] md:flex-row md:flex-wrap md:gap-x-4 md:gap-y-1">
              {#each klass.specs as spec (spec.href)}
                <!-- A 44 px hit area on a touch screen (spec §5, Touch); the line stays 20 px on desktop. -->
                <a
                  class="text-gold inline-flex min-h-11 items-center text-[14px] font-semibold whitespace-nowrap md:min-h-0"
                  href={spec.href}>{spec.name} ›</a
                >
              {/each}
            </span>
          {/snippet}
        </Row>
      {/each}
    </RowGrid>
  </Section>
  {#if railTail}{@render railTail()}{/if}
</Page>
