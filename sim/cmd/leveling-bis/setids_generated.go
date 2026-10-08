package main

// engineImplementedSetIDs is every data/builds/<build>/sets.json set id
// the engine registers a set for: by the same two keys the engine counts
// worn pieces with (core.Character.GetActiveSetBonuses), the ItemSet id
// first, then the set's name (or its alternative name).
//
// A set id this map calls "implemented" only means the engine runs SOME
// bonus function for it; some of those functions are themselves "Not
// implemented in sim" stubs for a specific piece count -- this map cannot
// see that finer distinction, so a verified set-completion run
// (sets.go's trySetCompletion) is what actually proves a specific piece
// count did something, not this map by itself.
//
// The map is rewritten, not edited: TestEngineImplementedSetIDs
// (setids_test.go) rebuilds it from the engine's registrations and fails
// when this file differs; FOREVER_UPDATE_SETIDS=1 go test ./cmd/leveling-bis
// -run TestEngineImplementedSetIDs rewrites it.
var engineImplementedSetIDs = map[int]bool{
	1: true, 41: true, 65: true, 81: true, 121: true, 122: true, 123: true, 124: true, 141: true, 142: true,
	143: true, 144: true, 161: true, 162: true, 163: true, 181: true, 182: true, 183: true, 184: true, 185: true,
	186: true, 187: true, 188: true, 189: true, 201: true, 202: true, 203: true, 204: true, 205: true, 206: true,
	207: true, 208: true, 209: true, 210: true, 211: true, 212: true, 213: true, 214: true, 215: true, 216: true,
	217: true, 218: true, 261: true, 281: true, 282: true, 301: true, 321: true, 341: true, 342: true, 343: true,
	344: true, 345: true, 346: true, 347: true, 348: true, 361: true, 362: true, 381: true, 382: true, 383: true,
	384: true, 386: true, 387: true, 388: true, 389: true, 390: true, 391: true, 392: true, 393: true, 394: true,
	395: true, 396: true, 397: true, 398: true, 421: true, 442: true, 443: true, 444: true, 461: true, 462: true,
	463: true, 464: true, 465: true, 466: true, 467: true, 468: true, 469: true, 470: true, 471: true, 472: true,
	473: true, 474: true, 475: true, 476: true, 477: true, 478: true, 479: true, 480: true, 481: true, 482: true,
	483: true, 484: true, 485: true, 486: true, 487: true, 488: true, 489: true, 490: true, 491: true, 492: true,
	494: true, 496: true, 497: true, 498: true, 499: true, 501: true, 502: true, 503: true, 505: true, 506: true,
	507: true, 509: true, 511: true, 512: true, 513: true, 514: true, 515: true, 516: true, 517: true, 518: true,
	519: true, 520: true, 522: true, 523: true, 524: true, 525: true, 526: true, 527: true, 528: true, 529: true,
	530: true, 533: true, 534: true, 535: true, 536: true, 537: true, 538: true, 539: true, 540: true, 541: true,
	542: true, 543: true, 544: true, 545: true, 546: true, 547: true, 548: true, 549: true, 550: true, 551: true,
	1618: true, 1619: true, 1620: true, 1621: true, 1622: true, 1623: true, 1624: true, 1625: true, 1626: true, 1627: true,
	1628: true, 1629: true, 1630: true, 1631: true, 1632: true, 1633: true, 1634: true, 1635: true, 1636: true, 1665: true,
	1666: true, 1667: true, 1668: true, 1669: true, 1670: true, 1671: true, 1672: true, 1673: true, 1674: true, 1675: true,
	1676: true, 1677: true, 1678: true, 1679: true, 1680: true, 1681: true, 1682: true, 1721: true, 1722: true, 1723: true,
	1724: true, 1725: true, 1726: true, 1727: true, 1728: true, 1729: true, 1730: true, 1731: true, 1732: true, 1733: true,
	1734: true, 1735: true, 1736: true, 1737: true, 1738: true, 1739: true, 1740: true, 1741: true, 1742: true, 1743: true,
	1744: true, 1745: true, 1746: true, 1747: true, 1748: true, 1749: true, 1750: true, 1751: true, 1752: true, 1753: true,
	1754: true, 1755: true, 1756: true, 1757: true, 1758: true, 1759: true, 1760: true, 1761: true, 1762: true, 1763: true,
	1764: true, 1765: true, 1766: true, 1767: true, 1768: true, 1769: true, 1770: true, 1774: true, 1775: true, 1776: true,
	1777: true, 1778: true, 1780: true, 1781: true, 1782: true, 1783: true, 1784: true, 1785: true, 1786: true, 1787: true,
	1788: true, 1789: true, 1790: true, 1791: true, 1792: true, 1793: true, 1795: true, 1796: true, 1797: true, 1798: true,
	1799: true, 1800: true, 1824: true, 1825: true, 1826: true, 1827: true, 1828: true, 1829: true, 1830: true, 1831: true,
	1832: true, 1858: true, 1861: true, 1863: true, 1864: true, 1882: true, 1885: true, 1887: true, 1891: true, 1892: true,
	1895: true, 1898: true, 1899: true, 1905: true, 1906: true, 1907: true, 1908: true, 1968: true, 1972: true, 2071: true,
	2072: true, 2073: true, 2074: true, 2075: true, 2076: true, 2077: true, 2078: true, 2079: true, 2080: true, 2081: true,
	2086: true, 2087: true, 2090: true, 2091: true, 2094: true, 2095: true, 2098: true, 2099: true, 2100: true, 2101: true,
	2102: true, 2103: true, 2104: true, 2105: true, 2106: true, 2107: true, 2108: true, 2109: true, 2110: true, 2111: true,
	2112: true, 2113: true, 2114: true, 2115: true, 2132: true,
}

func setEffectImplemented(setID int) bool {
	return engineImplementedSetIDs[setID]
}
