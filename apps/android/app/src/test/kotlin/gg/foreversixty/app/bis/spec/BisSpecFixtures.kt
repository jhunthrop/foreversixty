package gg.foreversixty.app.bis.spec

import gg.foreversixty.app.items.ItemLook

/** The mocks' three lists (web-bis-dps-60, web-bis-low-20, web-bis-healer-60), word for word. */
object BisSpecFixtures {
    const val BUILD = "1.60.1.70790"
    private var nextId = 1

    private fun look(name: String, quality: Int) = ItemLook(nextId++, name, quality, null)

    private fun cell(
        slot: String, level: String?, look: ItemLook?, stats: String, source: String,
        gain: String?, unit: String?, versus: String?, versusItem: ItemLook?,
    ) = BisSpecCell(slot, level, look, stats, source, gain, unit, versus, versusItem)

    private fun options(labels: List<String>, on: String, icons: Boolean = false) =
        labels.map { BisSpecOption(it, it == on, icon = if (icons) "ability_hunter_focusfire" else null) }

    private val levels = listOf("20-29", "30-39", "40-49", "50-59", "60")
    private val factions = listOf("Alliance", "Horde")
    private val profiles = listOf("Raid-ready", "Bare")
    private val hunterSpecs = listOf("Beast Mastery", "Marksmanship", "Survival")
    private const val ADDON_LEGEND_TAIL = ", with DPS per point and the sim error."

    private fun controls(bar: String, specs: List<String>, spec: String, level: String, profile: String?, race: String, sources: String) = BisSpecControls(
        bar = bar,
        specs = options(specs, spec, icons = true),
        levels = options(levels, level),
        factions = options(factions, "Alliance"),
        profiles = if (profile == null) emptyList() else options(profiles, profile),
        race = race,
        sources = sources,
    )

    val sixty = BisSpecContent(
        title = "Marksmanship Hunter",
        classSlug = "hunter",
        meta = "Best in slot",
        answer = BisSpecAnswer(
            label = "This set, simmed 7 nov", value = "1083.6", unit = "DPS", facts = emptyList(),
            line = "Level 60 Night-elf · 51 talent points · level 63 unknown, 180 s · Raid-ready, Phase 1: 14 buffs, 8 debuffs and 6 consumables. · 3,000 runs ·",
            rankingSim = "View ranking sim ›", openSim = true,
        ),
        controls = controls("Marksmanship · 60 · Alliance · Raid-ready", hunterSpecs, "Marksmanship", "60", "Raid-ready", "Night Elf", "All sources"),
        summary = BisSpecSummary("16 of 17 slots have a pick at this level.", null, null),
        notes = listOf("Hit to cap: 1.1% for specials, 8.1% for spells", "Legionite Glaive is a two-hander: the off hand stays empty."),
        cells = listOf(
            cell("Head", "Level 56", look("Black Dragonscale Helm", 4), "+21 Agi · +38 Int", "Crafted · Leatherworking", "Same", null, "next: −1.8 vs", look("Lieutenant Commander's Chain Helm", 3)),
            cell("Neck", "Level 58", look("Sentinel's Medallion", 3), "+15 Agi · +10 Sta", "Silverwing Sentinels reputation · Honored", "Same", null, "next: −1.8 vs", look("Skibi's Pendant", 3)),
            cell("Shoulder", "Level 60", look("Theramore Pauldrons", 4), "+27 Agi · +12 Sta", "Theramore Expeditionary Force reputation · Exalted", "Same", null, "next: −7.9 vs", look("Field Marshal's Chain Spaulders", 4)),
            cell("Back", "Level 58", look("Cape of the Black Baron", 3), "+15 Agi · +20 AP", "Baron Rivendare · Stratholme (58-60)", "+4.0", "DPS", "vs", look("Shadow Prowler's Cloak", 3)),
            cell("Chest", "Level 56", look("Dawn Armor", 4), "+37 Agi · +14 Crit", "Crafted · Leatherworking", "Same", null, "next: −26.3 vs", look("Savage Gladiator Chain", 4)),
            cell("Wrist", "Level 60", look("Forest Stalker's Bracers", 4), "+19 Agi · +11 Str", "Silverwing Sentinels reputation · Exalted", "Same", null, "next: −1.8 vs", look("Knight-Captain's Chain Armguards", 3)),
            cell("Main hand", "Level 58", look("Legionite Glaive", 3), "+16 Agi · +54 AP", "Crafted · Blacksmithing", "Same", null, "next: −1.0 vs", look("Grand Marshal's Stave", 4)),
            cell("Off hand", null, null, "", "Legionite Glaive is a two-hander; the off hand is taken.", null, null, null, null),
            cell("Ranged", "Level 58", look("Eaglehorn Long Bow", 3), "+10 Agi · +4 Sta", "World drop (59-62) · BoE · on the AH", "Same", null, "next: −20.0 vs", look("Hyper Deluxe Sniper Rifle Mk XVII", 4)),
            cell("Hands", "Level 60", look("Theramore Gloves", 4), "+27 Agi · +11 Str", "Theramore Expeditionary Force reputation · Exalted", "Same", null, "next: −0.8 vs", look("Mourning Handguards", 4)),
            cell("Waist", "Level 60", look("Ranger's Belt", 4), "+24 Agi · +14 Int", "Vendor: Pix Xizzix", "+4.1", "DPS", "vs", look("Marshal's Chain Girdle", 4)),
            cell("Legs", "Level 60", look("Sentinel's Chain Leggings", 4), "+35 Agi · +26 Sta", "Vendor: Illiyana Moonblaze", "+7.2", "DPS", "vs", look("Marshal's Chain Legguards", 4)),
            cell("Feet", "Level 54", look("Beaststalker's Boots", 3), "+21 Agi · +9 Sta", "Nerub'enkan · Stratholme (58-60)", "Same", null, "No close alternative", null),
            cell("Ring", "Level 40", look("Masons Fraternity Ring", 3), "+14 Agi · +5 Sta", "Quest: Divino-matic Rod · requires level 40", "Same", null, "No close alternative", null),
            cell("Ring", "Level 55", look("Don Julio's Band", 4), "+10 Sta · +14 Crit", "Stormpike Guard reputation · Exalted", "Same", null, "next: −1.5 vs", look("Cutthroat's Signet", 3)),
            cell("Trinket", null, look("Devilsaur Eye", 3), "", "Quest: The Green Drake · requires level 50", "Same", null, "No close alternative", null),
            cell("Trinket", "Level 50", look("Frozen Heart of the Mountain", 3), "+7 Fire res · +7 Nature res", "Crafted · Enchanting", "Same", null, "next: −6.3 vs", look("Heart of Wyrmthalak", 3)),
        ),
        groups = emptyList(),
        weights = BisSpecWeights(
            sentence = "Stack agility, then crit rating, then intellect.",
            legend = "Normalized to Agility = 1.00. Per point of stat, normalized to Agility = 1.00$ADDON_LEGEND_TAIL Haste: 6.83 per 1%, not in the table because no item at this level range has it",
            stats = listOf(
                BisSpecWeight("Agility", "1.00", 1.00f, "0.891 DPS per point"),
                BisSpecWeight("Crit rating", "0.82", 0.82f, "0.728 DPS per point"),
                BisSpecWeight("Intellect", "0.36", 0.36f, "0.323 DPS per point"),
                BisSpecWeight("Ranged attack power", "0.30", 0.30f, "0.270 DPS per point"),
                BisSpecWeight("Hit rating", null, 0f, null),
                BisSpecWeight("Spell power", null, 0f, null),
            ),
        ),
        guide = "Marksmanship Hunter guide ›",
    )

    val twenty = BisSpecContent(
        title = "Beast Mastery Hunter",
        classSlug = "hunter",
        meta = "Best in slot",
        answer = BisSpecAnswer(
            label = "This set, simmed 7 nov", value = "91.0", unit = "DPS", facts = emptyList(),
            line = "Level 29 Night-elf · 11 talent points · level 23 unknown, 180 s · Bare character: no raid buffs, debuffs or consumables. · 3,000 runs",
            rankingSim = null, openSim = true,
        ),
        controls = controls("Beast Mastery · 20-29 · Alliance", hunterSpecs, "Beast Mastery", "20-29", null, "Night Elf", "Sources (7)"),
        summary = BisSpecSummary("14 of 17 slots have a pick at this level.", "Changes from 20 to 29: see 30 to 39 ›", null),
        notes = listOf(
            "Reputation above Friendly is off below level 60, so those slots are re-picked without it. Turn it on in Sources to include it.",
            "Hit to cap: 9% for specials, 16% for spells",
            "Impaling Harpoon is a two-hander: the off hand stays empty.",
        ),
        cells = emptyList(),
        groups = listOf(
            BisSpecGroup("Wailing Caverns", "4 pieces", listOf(
                cell("Shoulder", "Level 18", look("Serpent's Shoulders", 3), "+7 Spi · +5 Agi", "Lady Anacondra · Wailing Caverns (15-24) · 40%", "+0.1", "DPS", "vs", look("Magician's Mantle", 3)),
                cell("Back", "Level 18", look("Glowing Lizardscale Cloak", 3), "+6 Agi · +2 Nature res", "Skum · Wailing Caverns (15-24) · 50%", "+0.3", "DPS", "vs", look("Cape of the Brotherhood", 3)),
                cell("Hands", "Level 14", look("Gloves of the Fang", 3), "+6 Agi · +4 Str", "Druid of the Fang · Wailing Caverns (15-24) · 2%", "Same", null, "next: −0.1 vs", look("Trapper's Leather Gloves", 3)),
                cell("Legs", "Level 17", look("Leggings of the Fang", 3), "+5 Str · +4 Sta", "Lord Cobrahn · Wailing Caverns (15-24) · 20%", "Same", null, "next: −0.3 vs", look("Duty Bound Leggings", 2)),
            )),
            BisSpecGroup("The Deadmines", "3 pieces", listOf(
                cell("Waist", "Level 17", look("Blackened Defias Belt", 3), "+18 AP", "Captain Greenskin · The Deadmines (17-26) · 30%", "+0.1", "DPS", "vs", look("Trapper's Leather Belt", 3)),
                cell("Ring", "Level 17", look("Lavishly Jeweled Ring", 3), "+2 Agi · +6 Int", "Gilnid · The Deadmines (17-26) · 45%", "+1.1", "DPS", "vs", look("Volcanic Rock Ring", 2)),
                cell("Main hand", "Level 17", look("Impaling Harpoon", 3), "+9 Agi · +5 Int", "Captain Greenskin · The Deadmines (17-26) · 30%", "+0.4", "DPS", "vs", look("Bronze Dory", 2)),
            )),
            BisSpecGroup("Crafted · Leatherworking", "1 piece", listOf(
                cell("Head", "Level 20", look("Brawler's Leather Hood", 3), "+10 Sta · +8 Agi", "Crafted · Leatherworking", "+1.0", "DPS", "vs", look("Trapper's Leather Hood", 3)),
            )),
            BisSpecGroup("Silverwing Sentinels reputation", "2 pieces", listOf(
                cell("Neck", "Level 18", look("Sentinel's Medallion", 3), "+6 Agi · +2 Sta", "Silverwing Sentinels reputation · Honored", "+1.2", "DPS", "vs", look("Erudite's Amulet", 2)),
                cell("Ring", "Level 18", look("Protector's Band", 3), "+4 Str · +4 Agi", "Silverwing Sentinels reputation · Honored", "+0.8", "DPS", "vs", look("Volcanic Rock Ring", 2)),
            )),
            BisSpecGroup("Quest", "1 piece", listOf(
                cell("Chest", "Level 14", look("Tunic of Westfall", 3), "+5 Sta · +11 Agi", "Quest: The Defias Brotherhood · requires level 14", "+0.3", "DPS", "vs", look("Trapper's Leather Armor", 3)),
            )),
            BisSpecGroup("World / BoE", "3 pieces", listOf(
                cell("Wrist", "Level 19", look("Forest Leather Bracers", 2), "+5 Agi", "World drop (20-26) · BoE · on the AH", "+0.3", "DPS", "vs", look("Wolf Bracers", 2)),
                cell("Feet", "Level 19", look("Feet of the Lynx", 3), "+8 Agi · +3 Str", "World drop (20-27) · BoE · on the AH", "+0.3", "DPS", "vs", look("Trapper's Leather Boots", 3)),
                cell("Ranged", "Level 20", look("Ranger Bow", 3), "+1 Agi", "World drop (21-28) · BoE · on the AH", "+0.3", "DPS", "vs", look("Venomstrike", 3)),
            )),
        ),
        weights = BisSpecWeights(
            sentence = "Stack agility, then intellect, then ranged attack power.",
            legend = "Normalized to Agility = 1.00. Per point of stat, normalized to Agility = 1.00$ADDON_LEGEND_TAIL Haste: 1.40 per 1%, not in the table because no item at this level range has it",
            stats = listOf(
                BisSpecWeight("Agility", "1.00", 1.00f, "0.316 DPS per point"),
                BisSpecWeight("Intellect", "0.51", 0.51f, "0.162 DPS per point"),
                BisSpecWeight("Ranged attack power", "0.32", 0.32f, "0.100 DPS per point"),
                BisSpecWeight("Hit rating", "0.31", 0.31f, "0.097 DPS per point"),
                BisSpecWeight("Crit rating", "0.21", 0.21f, "0.066 DPS per point"),
                BisSpecWeight("Spell power", null, 0f, null),
            ),
        ),
        guide = "Beast Mastery Hunter guide ›",
    )

    val healer = BisSpecContent(
        title = "Holy Priest",
        classSlug = "priest",
        meta = "Best in slot",
        answer = BisSpecAnswer(
            label = "This set, simmed 7 nov", value = "867.3", unit = "HPS",
            facts = listOf(BisSpecFact("Overheal", "6%"), BisSpecFact("Mana lasts", "2:00"), BisSpecFact("Healing per mana", "6.6")),
            line = "Level 60 Human · 51 talent points · level 63 unknown, 180 s · Raid-ready, Phase 1: 14 buffs, 8 debuffs and 10 consumables. · 3,003 runs",
            rankingSim = null, openSim = false,
        ),
        controls = controls("Holy · 60 · Alliance · Raid-ready", listOf("Discipline", "Holy", "Shadow"), "Holy", "60", "Raid-ready", "Human", "All sources"),
        summary = BisSpecSummary("16 of 17 slots have a pick at this level.", null, null),
        notes = listOf("Glowing Brightwood Staff is a two-hander: the off hand stays empty."),
        cells = listOf(
            cell("Head", "Level 57", look("Devout Crown", 3), "+24 Int · +15 Spi", "Darkmaster Gandling · Scholomance (58-60)", "Same", null, "next: −18.4 vs", look("Mooncloth Circlet", 3)),
            cell("Neck", "Level 36", look("Glowing Eye of Mordresh", 3), "+20 Healing · +8 Int", "Mordresh Fire Eye · Razorfen Downs (37-46)", "Same", null, "next: −2.9 vs", look("Lady Maye's Pendant", 4)),
            cell("Shoulder", "Level 56", look("Mooncloth Shoulders", 3), "+37 Healing · +13 SP", "Crafted · Tailoring", "Same", null, "next: −5.5 vs", look("Earthenweave Mantle", 2)),
            cell("Back", "Level 60", look("Drape of Recovery", 3), "+31 Healing · +10 SP", "Vendor: Pix Xizzix", "+3.9", "HPS", "vs", look("Caretaker's Cape", 3)),
            cell("Chest", "Level 57", look("Truefaith Vestments", 4), "+14 Sta · +73 Healing", "Crafted · Tailoring", "+0.8", "HPS", "vs", look("Field Marshal's Satin Tunic", 4)),
            cell("Wrist", "Level 52", look("Devout Bracers", 3), "+8 Int · +10 Spi", "Thuzadin Shadowcaster · Stratholme (58-60) · 1%", "Same", null, "No close alternative", null),
            cell("Main hand", "Level 49", look("Glowing Brightwood Staff", 4), "+29 Int · +15 Sta", "World drop (50-57) · BoE · on the AH", "+21.1", "HPS", "vs", look("Resurgence Rod", 3)),
            cell("Off hand", null, null, "", "Glowing Brightwood Staff is a two-hander; the off hand is taken.", null, null, null, null),
            cell("Ranged", "Level 60", look("Torch of Light", 4), "+17 Healing · +6 SP", "Crafted · Enchanting", "+5.3", "HPS", "vs", look("Brilliant Wand", 3)),
            cell("Hands", "Level 54", look("Devout Gloves", 3), "+17 Spi · +10 Int", "Archivist Galford · Stratholme (58-60)", "Same", null, "next: −15.7 vs", look("Hands of the Exalted Herald", 3)),
            cell("Waist", "Level 53", look("Wisdom of the Timbermaw", 3), "+14 Int · +31 Healing", "Crafted · Tailoring", "+3.4", "HPS", "vs", look("Penitent's Cinch", 4)),
            cell("Legs", "Level 56", look("Devout Skirt", 3), "+23 Spi · +15 Int", "Baron Rivendare · Stratholme (58-60)", "Same", null, "next: −25.7 vs", look("Earthenweave Leggings", 2)),
            cell("Feet", "Level 54", look("Devout Sandals", 3), "+17 Spi · +10 Int", "Maleki the Pallid · Stratholme (58-60)", "Same", null, "next: −16.1 vs", look("Faith Healer's Boots", 3)),
            cell("Ring", "Level 60", look("Blessed Band of Light", 3), "+31 Healing · +10 SP", "Vendor: Pix Xizzix", "Same", null, "next: −0.5 vs", look("Rosewine Circle", 3)),
            cell("Ring", "Level 57", look("Band of Mending", 3), "+7 Sta · +7 Int", "Balnazzar · Stratholme (58-60)", "Same", null, "next: −1.4 vs", look("Rosewine Circle", 3)),
            cell("Trinket", "Level 60", look("Serenity Field", 3), "+42 Healing · +14 SP", "Vendor: Pix Xizzix", "+8.3", "HPS", "vs", look("Weakness Analyzer", 3)),
            cell("Trinket", "Level 50", look("Dormant Heart of the Mountain", 3), "+7 Fire res · +7 Nature res", "Crafted · Enchanting", "Same", null, "No close alternative", null),
        ),
        groups = emptyList(),
        weights = BisSpecWeights(
            sentence = "Stack mp5, then healing power, then spirit.",
            legend = "Normalized to MP5 = 1.00. Per point of stat, normalized to MP5 = 1.00, with HPS per point and the sim error. Haste: not significant",
            stats = listOf(
                BisSpecWeight("MP5", "1.00", 1.00f, "0.569 HPS per point"),
                BisSpecWeight("Healing power", "0.73", 0.73f, "0.417 HPS per point"),
                BisSpecWeight("Spirit", "0.55", 0.55f, "0.312 HPS per point"),
                BisSpecWeight("Intellect", "0.38", 0.38f, "0.214 HPS per point"),
                BisSpecWeight("Crit rating", "0.30", 0.30f, "0.171 HPS per point"),
            ),
        ),
        guide = "Holy Priest guide ›",
    )
}
