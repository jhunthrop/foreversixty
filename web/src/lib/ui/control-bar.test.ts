// @vitest-environment jsdom
import { afterEach, describe, expect, it, vi } from 'vitest';
import { CONTROL_BAR_EVENT, bootControlBar, paramValue, radioKeyTarget, withParam } from './control-bar';

describe('withParam', () => {
  it('sets, replaces and drops a parameter, keeping the others', () => {
    expect(withParam('', 'band', '60')).toBe('?band=60');
    expect(withParam('?band=20&faction=horde', 'band', '60')).toBe('?band=60&faction=horde');
    expect(withParam('?band=20&faction=horde', 'band', '20', '20')).toBe('?faction=horde');
    expect(withParam('?band=20', 'band', '')).toBe('');
  });
});

describe('paramValue', () => {
  it('reads an allowed value and falls back otherwise', () => {
    expect(paramValue('?band=60', 'band', ['20', '60'], '20')).toBe('60');
    expect(paramValue('?band=70', 'band', ['20', '60'], '20')).toBe('20');
    expect(paramValue('', 'band', ['20'], '20')).toBe('20');
  });
});

describe('radioKeyTarget', () => {
  it('moves with the arrows, wraps, and jumps with Home and End', () => {
    expect(radioKeyTarget('ArrowRight', 0, 3)).toBe(1);
    expect(radioKeyTarget('ArrowRight', 2, 3)).toBe(0);
    expect(radioKeyTarget('ArrowLeft', 0, 3)).toBe(2);
    expect(radioKeyTarget('Home', 2, 3)).toBe(0);
    expect(radioKeyTarget('End', 0, 3)).toBe(2);
    expect(radioKeyTarget('Tab', 0, 3)).toBeNull();
  });
});

describe('bootControlBar', () => {
  afterEach(() => {
    document.body.innerHTML = '';
    window.history.replaceState(null, '', '/');
  });

  function bar(): HTMLElement {
    document.body.innerHTML = `
      <div data-control-bar>
        <button data-control-bar-toggle aria-expanded="false"></button>
        <div role="radiogroup" data-param="band" data-default="20">
          <a role="radio" href="?band=20" data-value="20" aria-checked="true" tabindex="0">20-29</a>
          <a role="radio" href="?band=30" data-value="30" aria-checked="false" tabindex="-1">30-39</a>
          <a role="radio" href="?band=40" data-value="40" aria-checked="false" tabindex="-1" aria-disabled="true">40-49</a>
          <a role="radio" href="?band=60" data-value="60" aria-checked="false" tabindex="-1">60</a>
        </div>
        <a role="radio" href="/bis/hunter/survival" data-navigate="true" data-value="sv">Survival</a>
      </div>`;
    const element = document.querySelector<HTMLElement>('[data-control-bar]')!;
    bootControlBar(element);
    return element;
  }

  it('writes the choice to the URL, checks it, and announces it', () => {
    const element = bar();
    const heard = vi.fn();
    document.addEventListener(CONTROL_BAR_EVENT, (event) => heard((event as CustomEvent).detail));
    document.querySelector<HTMLElement>('[data-value="60"]')!.click();
    expect(window.location.search).toBe('?band=60');
    expect(heard).toHaveBeenCalledWith({ param: 'band', value: '60' });
    expect(element.querySelector('[data-value="60"]')!.getAttribute('aria-checked')).toBe('true');
    expect(element.querySelector('[data-value="20"]')!.getAttribute('tabindex')).toBe('-1');
  });

  it('drops the parameter for the default choice', () => {
    bar();
    document.querySelector<HTMLElement>('[data-value="60"]')!.click();
    document.querySelector<HTMLElement>('[data-value="20"]')!.click();
    expect(window.location.search).toBe('');
  });

  it('ignores a disabled choice and skips it with the arrow keys', () => {
    bar();
    document.querySelector<HTMLElement>('[data-value="40"]')!.click();
    expect(window.location.search).toBe('');
    const thirty = document.querySelector<HTMLElement>('[data-value="30"]')!;
    thirty.focus();
    thirty.dispatchEvent(new KeyboardEvent('keydown', { key: 'ArrowRight', bubbles: true }));
    expect(window.location.search).toBe('?band=60');
    expect(document.activeElement?.getAttribute('data-value')).toBe('60');
  });

  it('opens and closes the phone controls, and collapses on a choice', () => {
    const element = bar();
    const toggle = element.querySelector<HTMLElement>('[data-control-bar-toggle]')!;
    toggle.click();
    expect(element.hasAttribute('data-open')).toBe(true);
    expect(toggle.getAttribute('aria-expanded')).toBe('true');
    document.querySelector<HTMLElement>('[data-value="30"]')!.click();
    expect(element.hasAttribute('data-open')).toBe(false);
  });

  it('lets a path choice navigate as a link', () => {
    bar();
    const link = document.querySelector<HTMLElement>('[data-navigate="true"]')!;
    const event = new MouseEvent('click', { bubbles: true, cancelable: true });
    link.dispatchEvent(event);
    expect(event.defaultPrevented).toBe(false);
  });
});
