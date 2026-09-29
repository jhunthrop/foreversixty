package main

// engineImplementedSetIDs is every data/builds/<build>/sets.json set id
// whose NAME matches a core.NewItemSet(core.ItemSet{Name: "..."})
// registration somewhere in wowsims-forever's own source -- generated
// once (2026-09-28 weights-effects lane) the same way
// setclass_generated.go's own table is: grep every
// item_sets_pve.go/item_sets_pvp.go (rogue: items_sets_pve.go/
// items_sets_pvp.go) plus sim/common/item_sets/*.go for a NewItemSet
// call's own Name field, across every class, then match each name
// against sets.json's own name column. 232 of this build's 536 sets
// matched; the other 304 are either a set this engine build has never
// registered at all, or one whose registered name does not match this
// data pipeline's own name for it exactly (a data/engine drift this
// map cannot distinguish from "truly unimplemented" -- see the lane
// report's coverage table for the caveat).
//
// A set id this map calls "implemented" only means the engine runs
// SOME bonus function for it; some of those functions are themselves
// "// Not implemented in sim" stubs for a specific piece count (see
// band.go's own crossClassSetItem doc for a related, class-scoped
// engine gap) -- this map cannot see that finer distinction, so a
// verified set-completion run (rank.go's trySetCompletion) is what
// actually proves a specific piece count did something, not this map
// by itself.
//
// Regenerate by re-running that same two-pass script against
// sim/enginever.Version's commit next.
var engineImplementedSetIDs = map[int]bool{
	41: true, 65: true, 81: true, 121: true, 122: true, 123: true, 124: true, 141: true, 142: true, 143: true,
	144: true, 181: true, 182: true, 183: true, 184: true, 185: true, 186: true, 187: true, 188: true, 189: true,
	201: true, 202: true, 203: true, 204: true, 205: true, 206: true, 207: true, 208: true, 209: true, 210: true,
	211: true, 212: true, 213: true, 214: true, 215: true, 216: true, 217: true, 218: true, 261: true, 301: true,
	321: true, 342: true, 344: true, 361: true, 362: true, 383: true, 384: true, 386: true, 387: true, 388: true,
	389: true, 390: true, 391: true, 392: true, 393: true, 394: true, 395: true, 396: true, 397: true, 398: true,
	402: true, 421: true, 442: true, 443: true, 444: true, 461: true, 462: true, 463: true, 464: true, 465: true,
	466: true, 467: true, 468: true, 469: true, 470: true, 471: true, 472: true, 473: true, 474: true, 475: true,
	477: true, 478: true, 480: true, 481: true, 482: true, 483: true, 484: true, 485: true, 486: true, 487: true,
	488: true, 489: true, 490: true, 491: true, 494: true, 495: true, 496: true, 497: true, 498: true, 499: true,
	501: true, 502: true, 503: true, 505: true, 506: true, 507: true, 509: true, 511: true, 512: true, 513: true,
	515: true, 517: true, 518: true, 519: true, 520: true, 523: true, 524: true, 525: true, 526: true, 527: true,
	528: true, 529: true, 530: true, 533: true, 534: true, 535: true, 536: true, 537: true, 538: true, 539: true,
	540: true, 541: true, 542: true, 543: true, 544: true, 545: true, 547: true, 548: true, 549: true, 550: true,
	551: true, 1666: true, 1667: true, 1668: true, 1669: true, 1670: true, 1671: true, 1672: true, 1674: true, 1676: true,
	1677: true, 1678: true, 1679: true, 1680: true, 1681: true, 1682: true, 1720: true, 1721: true, 1722: true, 1725: true,
	1727: true, 1728: true, 1730: true, 1731: true, 1734: true, 1737: true, 1739: true, 1740: true, 1741: true, 1743: true,
	1746: true, 1747: true, 1749: true, 1752: true, 1754: true, 1755: true, 1759: true, 1763: true, 1766: true, 1768: true,
	1769: true, 1778: true, 1780: true, 1781: true, 1782: true, 1783: true, 1784: true, 1785: true, 1786: true, 1787: true,
	1788: true, 1789: true, 1790: true, 1791: true, 1792: true, 1793: true, 1795: true, 1796: true, 1797: true, 1798: true,
	1799: true, 1800: true, 1822: true, 1823: true, 1825: true, 1826: true, 1827: true, 1828: true, 1829: true, 1831: true,
	1832: true, 1856: true, 1857: true, 1858: true, 1861: true, 1863: true, 1864: true, 1866: true, 1882: true, 1885: true,
	1887: true, 1891: true, 1892: true, 1895: true, 1898: true, 1899: true, 1905: true, 1906: true, 1907: true, 1908: true,
	2072: true, 2075: true,
}

func setEffectImplemented(setID int) bool {
	return engineImplementedSetIDs[setID]
}
