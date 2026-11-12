import { describe, expect, it } from 'vitest';
import { SEAL } from './seal';
import { FRAME_BOTTOM, FRAME_LEFT, FRAME_RIGHT, FRAME_SLOTS } from '../gear/slots';
import {
  CENTRE_BOX,
  FRAME_SIZE,
  INNER_BOX,
  MIN_FRAME_PX,
  MIN_SLOT_PX,
  boxStyle,
  frameFits,
  medallionSrc,
  slotBox,
} from './paperdoll-frame';

const STUB_AVATAR = 'https://example.test/avatar.jpg';

describe('the paperdoll frame', () => {
  it('lists the client order: left, right, then the three weapon sockets', () => {
    expect(FRAME_LEFT[0]).toBe('head');
    expect(FRAME_LEFT.at(-1)).toBe('wrist');
    expect(FRAME_RIGHT[0]).toBe('hands');
    expect(FRAME_RIGHT.slice(-2)).toEqual(['trinket1', 'trinket2']);
    expect(FRAME_BOTTOM).toEqual(['main_hand', 'off_hand', 'ranged']);
    expect(FRAME_SLOTS).toHaveLength(19);
  });

  it('places the left column left of the centre, the right column right of it, the weapons below', () => {
    for (const slot of FRAME_LEFT) expect(slotBox(slot).x + slotBox(slot).w).toBeLessThan(CENTRE_BOX.x);
    for (const slot of FRAME_RIGHT) expect(slotBox(slot).x).toBeGreaterThan(CENTRE_BOX.x + CENTRE_BOX.w);
    for (const slot of FRAME_BOTTOM) expect(slotBox(slot).y).toBeGreaterThan(CENTRE_BOX.y + CENTRE_BOX.h - 5);
    const ys = FRAME_LEFT.map((slot) => slotBox(slot).y);
    expect(ys).toEqual(ys.toSorted((a, b) => a - b));
  });

  it('centres the centre box inside the frame’s inner edges: equal insets within 1 px of the 2x frame', () => {
    const toPx = (pct: number): number => (pct * FRAME_SIZE.width) / 100;
    const toPy = (pct: number): number => (pct * FRAME_SIZE.height) / 100;
    const left = toPx(CENTRE_BOX.x - INNER_BOX.x);
    const right = toPx(INNER_BOX.x + INNER_BOX.w - CENTRE_BOX.x - CENTRE_BOX.w);
    const top = toPy(CENTRE_BOX.y - INNER_BOX.y);
    const bottom = toPy(INNER_BOX.y + INNER_BOX.h - CENTRE_BOX.y - CENTRE_BOX.h);
    expect(Math.abs(left - right)).toBeLessThanOrEqual(1);
    expect(Math.abs(top - bottom)).toBeLessThanOrEqual(1);
    expect(Math.min(left, right, top, bottom)).toBeGreaterThanOrEqual(0);
  });

  it('fails loudly on a slot the frame does not have', () => {
    expect(() => slotBox('feet2')).toThrow(/no slot "feet2"/);
  });

  it('writes a box as percent positions', () => {
    expect(boxStyle({ x: 1, y: 2, w: 3, h: 4 })).toBe('left:1%;top:2%;width:3%;height:4%');
  });

  it('stays a frame while every slot is at least 36 px wide: 423 px', () => {
    expect(MIN_SLOT_PX).toBe(36);
    expect(MIN_FRAME_PX).toBe(423);
    expect(frameFits(424)).toBe(true);
    expect(frameFits(680)).toBe(true);
    expect(frameFits(422)).toBe(false);
    expect(frameFits(390)).toBe(false);
  });

  it('shows the Battle.net head render, else the painted seal, and the seal when the render fails', () => {
    expect(medallionSrc(STUB_AVATAR)).toBe(STUB_AVATAR);
    expect(medallionSrc(null)).toBe(SEAL.disc128);
    expect(medallionSrc('')).toBe(SEAL.disc128);
    expect(medallionSrc(STUB_AVATAR, true)).toBe(SEAL.disc128);
  });
});
