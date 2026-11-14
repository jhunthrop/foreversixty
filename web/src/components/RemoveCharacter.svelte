<!-- web/src/components/RemoveCharacter.svelte -->
<!-- /remove-my-character: the character directory's opt-out (Blizzard API terms 11(c); spec
     docs/superpowers/specs/2026-11-14-character-directory-discovery.md §6). One panel names the
     character (region and full name), with the lookup's answer under the name; two panels are the
     two proofs: a Battle.net sign-in that holds the character removes it at once, or a code in its
     public guild note that a guildmate's addon sends (the wait line under Check now). Removed is
     final: the directory deletes it and never stores it again. The page shows only the name the
     visitor typed, never a character's data. Page layout "form", like /connect. -->
<script lang="ts">
  import { onMount } from 'svelte';
  import { AccountError, battlenetStartUrl, fetchMeOnce, type Me } from '../lib/account/api';
  import { zonedDayMonthYear } from '../lib/dates';
  import {
    optOutStatus,
    removeOwnCharacter,
    startOptOutCode,
    validCharacterName,
    type OptOutCode,
    type OptOutRegion,
  } from '../lib/directory/opt-out-api';
  import { optOutCopy as copy } from '../lib/directory/opt-out-copy';
  import Button from './ui/Button.svelte';
  import Code from './ui/Code.svelte';
  import Field from './ui/Field.svelte';
  import Page from './ui/Page.svelte';
  import PageHeader from './ui/PageHeader.svelte';
  import Section from './ui/Section.svelte';
  import Segmented from './ui/Segmented.svelte';

  const PATH = '/remove-my-character';

  let region = $state<OptOutRegion>('us');
  let name = $state('');
  let nameError = $state('');
  let me = $state<Me | null>(null);
  let code = $state<OptOutCode | null>(null);
  let removed = $state(false);
  let message = $state('');
  let waiting = $state(false);
  let busy = $state(false);

  const shownName = $derived(name.trim());
  const nextPath = $derived(`${PATH}?region=${region}&name=${encodeURIComponent(shownName)}`);

  /** The form's name, or null after saying why it cannot be a character. */
  function who(): { region: OptOutRegion; name: string } | null {
    nameError = validCharacterName(name) ? '' : copy.nameInvalid;
    message = '';
    waiting = false;
    return nameError === '' ? { region, name: shownName } : null;
  }

  function sayError(error: unknown): void {
    if (error instanceof AccountError && error.status === 403) message = copy.notYours;
    else message = copy.failed;
  }

  async function run(work: () => Promise<void>): Promise<void> {
    busy = true;
    try {
      await work();
    } catch (error) {
      sayError(error);
    } finally {
      busy = false;
    }
  }

  const onRemoveOwn = (): Promise<void> =>
    run(async () => {
      const target = who();
      if (target === null) return;
      await removeOwnCharacter(target);
      removed = true;
    });

  const onGetCode = (): Promise<void> =>
    run(async () => {
      const target = who();
      if (target === null) return;
      code = await startOptOutCode(target);
    });

  const onCheck = (): Promise<void> =>
    run(async () => {
      const target = who();
      if (target === null || code === null) return;
      const state = await optOutStatus(target, code.code);
      removed = state === 'removed';
      waiting = !removed;
    });

  /** Back to an empty form after a removal. */
  function onAnother(): void {
    removed = false;
    code = null;
    name = '';
    message = '';
  }

  onMount(() => {
    const query = new URLSearchParams(window.location.search);
    if (query.get('region') === 'eu') region = 'eu';
    name = query.get('name') ?? '';
    void fetchMeOnce()
      .then((found) => (me = found))
      .catch(() => (me = null));
  });
</script>

{#snippet headMeta()}<span data-testid="remove-lead">{removed ? copy.removed(shownName) : copy.lead}</span
  >{/snippet}

{#snippet header()}<PageHeader title={removed ? copy.removedTitle : copy.title} meta={headMeta} />{/snippet}

<div data-testid="remove-character" data-state={removed ? 'removed' : code ? 'waiting' : 'form'}>
  <Page layout="form" {header}>
    {#if removed}
      <div>
        <Button onclick={onAnother} testid="remove-another">{copy.removeAnother}</Button>
      </div>
    {:else}
      <Section title={copy.findHeading} level={2} panel frame="framed">
        <div class="remove-body">
          <Segmented
            label={copy.regionLabel}
            options={copy.regions}
            value={region}
            onChange={(next) => (region = next)}
            testid="remove-region"
          />
          <Field
            id="remove-name"
            name="name"
            label={copy.nameLabel}
            placeholder={copy.namePlaceholder}
            autocomplete="off"
            maxlength={25}
            bind:value={name}
            error={nameError || undefined}
            testid="remove-name"
          />
          <p class="text-strong text-[14px] empty:hidden" role="status" data-testid="remove-message">
            {message}
          </p>
        </div>
      </Section>
      <Section title={copy.battlenetHeading} level={2} panel frame="framed">
        <div class="remove-body">
          <p>{copy.battlenetLine}</p>
          <div>
            {#if me === null}
              <Button tone="primary" href={battlenetStartUrl(nextPath)} testid="remove-battlenet"
                >{copy.battlenetSignIn}</Button
              >
            {:else}
              <Button tone="primary" disabled={busy} onclick={onRemoveOwn} testid="remove-own"
                >{copy.battlenetRemove(shownName || copy.nameLabel.toLowerCase())}</Button
              >
            {/if}
          </div>
        </div>
      </Section>
      <Section title={copy.noteHeading} level={2} panel frame="framed">
        <div class="remove-body">
          <p>{copy.noteLine}</p>
          {#if code}
            <p>{copy.codeLead(shownName)}</p>
            <Code text={code.code} big testid="remove-code" />
            <p class="text-muted">{copy.codeExpires(zonedDayMonthYear(code.expires_at))}</p>
            <div>
              <Button disabled={busy} onclick={onCheck} testid="remove-check"
                >{busy ? copy.checking : copy.check}</Button
              >
            </div>
            <p class="text-muted text-[14px] empty:hidden" role="status" data-testid="remove-waiting">
              {waiting ? copy.waiting : ''}
            </p>
          {:else}
            <div>
              <Button disabled={busy} onclick={onGetCode} testid="remove-get-code">{copy.getCode}</Button>
            </div>
          {/if}
        </div>
      </Section>
    {/if}
  </Page>
</div>

<style>
  .remove-body {
    display: flex;
    flex-direction: column;
    gap: 14px;
  }
  .remove-body > p {
    font-size: 15px;
    line-height: 24px;
    max-width: 72ch;
  }
</style>
