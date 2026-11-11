// @vitest-environment jsdom
import { describe, expect, it } from 'vitest';
import { ICON_FAIL_SCRIPT, ICON_HIDE_ON_FAIL } from './icon-fail';

/** Runs the inline listener the way the page does, once for the file. */
new Function(ICON_FAIL_SCRIPT)();

function frame(hideOnFail: boolean): { frame: HTMLSpanElement; image: HTMLImageElement } {
  const span = document.createElement('span');
  if (hideOnFail) span.setAttribute(ICON_HIDE_ON_FAIL, '');
  const image = document.createElement('img');
  span.append(image);
  document.body.append(span);
  return { frame: span, image };
}

describe('the icon fail listener', () => {
  it('hides a frame marked to go with its failed image, keeping its box', () => {
    const { frame: span, image } = frame(true);
    image.dispatchEvent(new Event('error'));
    expect(span.style.visibility).toBe('hidden');
    expect(span.style.display).toBe('');
  });

  it('leaves a socket frame (one holding a column) for socket.css', () => {
    const { frame: span, image } = frame(false);
    image.dispatchEvent(new Event('error'));
    expect(span.style.visibility).toBe('');
  });
});
