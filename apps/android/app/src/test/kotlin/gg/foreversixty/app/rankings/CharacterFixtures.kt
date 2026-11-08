package gg.foreversixty.app.rankings

import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.model.BuildSeen
import gg.foreversixty.app.model.BuildTalents
import gg.foreversixty.app.model.CharacterBuild
import gg.foreversixty.app.model.CharacterFight
import gg.foreversixty.app.model.CharacterGearSlot
import gg.foreversixty.app.model.CharacterGearSlotWorn
import gg.foreversixty.app.model.CharacterHome
import gg.foreversixty.app.model.CharacterProfile
import gg.foreversixty.app.model.CharacterProfileGuild
import gg.foreversixty.app.model.CharacterProfileHeader
import gg.foreversixty.app.model.FeedRow
import gg.foreversixty.app.model.FeedSegment
import gg.foreversixty.app.model.HomeAlreadyBest
import gg.foreversixty.app.model.HomeBlocks
import gg.foreversixty.app.model.HomeGuildSims
import gg.foreversixty.app.model.HomeItem
import gg.foreversixty.app.model.HomeLadderRow
import gg.foreversixty.app.model.HomePick
import gg.foreversixty.app.model.HomeSpecWeek
import gg.foreversixty.app.model.HomeTalents
import gg.foreversixty.app.model.HomeUpgrade
import gg.foreversixty.app.model.Item
import gg.foreversixty.app.model.RaidCounts
import gg.foreversixty.app.model.RaidLine
import gg.foreversixty.app.model.ScoreComponentCaps
import gg.foreversixty.app.model.ScoreComponents
import gg.foreversixty.app.rankings.logic.CharacterExtras
import gg.foreversixty.app.rankings.logic.RankingsViewer
import java.time.Duration
import java.time.Instant
import java.time.ZoneId

/**
 * The character mocks' data (design/mocks/web/web-character-*): Elyra Duskvale, a Holy Priest 60 of
 * Cx66 Olympus on US Hardcore, and Felsnap, a Beast Mastery Hunter 34 on US Normal. The clock is the
 * mocks' (Fri 6 Nov 2026, 18:00 UTC).
 */
object CharacterFixtures {
    val now: Instant = Instant.parse("2026-11-06T18:00:00Z")
    val zone: ZoneId = ZoneId.of("America/Chicago")
    const val BUILD = "1.60.1.70790"
    const val ELYRA_KEY = "us/hardcore/elyra-duskvale"
    const val FELSNAP_KEY = "us/normal/felsnap"

    fun owner(key: String) = RankingsViewer(signedIn = true, current = null, accountKeys = setOf(key), history = true)

    private fun ago(duration: Duration): String = now.minus(duration).toString()

    /** One worn item: its slot, name, quality and level; [pick] the band's pick with its gain, or `best`. */
    private data class Doll(
        val slot: String,
        val id: Int,
        val name: String,
        val quality: Int,
        val level: Int,
        val pick: String? = null,
        val pickQuality: Int = 3,
        val gain: Double? = null,
        val best: Boolean = false,
        val twoHand: Boolean = false,
    )

    private val elyraDoll = listOf(
        Doll("head", 101, "Mooncloth Circlet", 3, 62, "Devout Crown", 3, 18.4),
        Doll("neck", 102, "Lady Maye's Pendant", 4, 64, "Glowing Eye of Mordresh", 3, 2.9),
        Doll("shoulder", 103, "Earthenweave Mantle", 2, 57, "Mooncloth Shoulders", 3, 5.5),
        Doll("back", 104, "Caretaker's Cape", 3, 63, "Drape of Recovery", 3, 3.9),
        Doll("chest", 105, "Field Marshal's Satin Tunic", 4, 74, "Truefaith Vestments", 4, 0.8),
        Doll("wrist", 106, "Devout Bracers", 3, 57, best = true),
        Doll("main_hand", 107, "Resurgence Rod", 3, 53, "Glowing Brightwood Staff", 4, 21.1, twoHand = true),
        Doll("ranged", 109, "Brilliant Wand", 3, 60, "Torch of Light", 4, 5.3),
        Doll("hands", 110, "Hands of the Exalted Herald", 3, 59, "Devout Gloves", 3, 15.7),
        Doll("waist", 111, "Penitent's Cinch", 4, 60, "Wisdom of the Timbermaw", 3, 3.4),
        Doll("legs", 112, "Earthenweave Leggings", 2, 58, "Devout Skirt", 3, 25.7),
        Doll("feet", 113, "Faith Healer's Boots", 3, 63, "Devout Sandals", 3, 16.1),
        Doll("finger1", 114, "Rosewine Circle", 3, 60, "Blessed Band of Light", 3, 0.5),
        Doll("finger2", 114, "Rosewine Circle", 3, 60, "Band of Mending", 3, 1.4),
        Doll("trinket1", 116, "Weakness Analyzer", 3, 65, "Serenity Field", 3, 8.3),
        Doll("trinket2", 117, "Dormant Heart of the Mountain", 3, 55, best = true),
    )

    private val felsnapDoll = listOf(
        Doll("head", 201, "Brawler's Leather Hood", 3, 25, "Brawler's Leather Helm", 3, 0.6),
        Doll("neck", 202, "Archmagister's Faceted Pendant", 3, 35, "Sentinel's Medallion", 3, 1.1),
        Doll("shoulder", 203, "Mantle of Thieves", 3, 30, "Forest Tracker Epaulets", 3, 0.3),
        Doll("back", 204, "Hawkeye's Cloak", 2, 35, "Tigerstrike Mantle", 3, 0.3),
        Doll("chest", 205, "Tunic of Westfall", 3, 24, "Garb of Florid Feathers", 3, 0.0),
        Doll("wrist", 206, "Madwolf Bracers", 2, 29, "Hawkeye's Bracers", 3, 0.3),
        Doll("main_hand", 207, "Armor Piercer", 2, 29, "Felbough Arm", 3, 1.3, twoHand = true),
        Doll("ranged", 209, "Ironweaver", 3, 34, "Golemsight Long Gun", 3, 0.7),
        Doll("hands", 210, "Pilferer's Gloves", 2, 28, "Cartilage Shapers", 3, 0.8),
        Doll("waist", 211, "Highlander's Chain Girdle", 3, 33, "Stalker's Leather Belt", 3, 0.7),
        Doll("legs", 212, "Petrolspill Leggings", 3, 30, "Troll's Bane Leggings", 3, 0.2),
        Doll("feet", 213, "Blackened Defias Boots", 3, 20, "Insignia Boots", 3, 0.7),
        Doll("finger1", 214, "Monkey Ring", 2, 32, "Ring of Precision", 3, 0.3),
        Doll("finger2", 214, "Monkey Ring", 2, 32, "Protector's Band", 3, 0.3),
    )

    private fun item(d: Doll) = Item(
        id = d.id, name = d.name, icon = "inv_misc_questionmark", slot = d.slot, quality = d.quality, requiredLevel = d.level,
        requiredLevelSource = "", itemLevel = d.level, armor = 0, damageMin = 0.0, damageMax = 0.0, speed = 0.0, dps = 0.0,
        twoHand = d.twoHand, effectText = "", unique = false, clientUnconfirmed = false, iconSource = "",
    )

    val items: Map<Int, Item> = (elyraDoll + felsnapDoll).associate { it.id to item(it) }

    private fun gear(doll: List<Doll>) = doll.map { CharacterGearSlot(it.slot, CharacterGearSlotWorn(it.id, it.name, it.level)) }

    private fun home(key: String, band: Int, unit: String, total: Double, doll: List<Doll>) = CharacterHome(
        character = key, state = "ok", band = band, unitWord = unit, dataBuild = BUILD, upgradeCount = doll.count { it.gain != null },
        totalGain = total, notSimCheckedCount = 0, notModelledCount = 0, syncedAt = ago(Duration.ofDays(1)), syncedFrom = "addon",
        upgrades = doll.filter { it.pick != null }.mapIndexed { i, d ->
            HomeUpgrade(
                d.slot, d.slot, HomePick(PICK_IDS + i, d.pick, d.pickQuality, null, null, PICK_SOURCES[d.pick] ?: "Dungeon"),
                gain = d.gain, gainWords = "", modelled = true,
            )
        },
        alreadyBest = doll.filter { it.best }.map { HomeAlreadyBest(it.slot, it.slot, HomeItem(it.id, it.name, it.quality)) },
        talents = HomeTalents("optimized"),
    )

    private val PICK_SOURCES = mapOf("Devout Skirt" to "Stratholme: Baron Rivendare", "Felbough Arm" to "Wailing Caverns: Mutanus the Devourer")
    private const val PICK_IDS = 900

    private fun fight(boss: String, value: Double, percentile: Double, at: String, spec: String = "Holy") =
        CharacterFight(boss, boss.length, 3, "hps", value, percentile, spec, at, "report-$boss", 1, kill = true)

    val elyra = CharacterProfile(
        character = CharacterProfileHeader(
            ELYRA_KEY, "us", "hardcore", "Elyra Duskvale", CharacterProfileGuild("Cx66 Olympus", "Officer"), 60, "Night Elf", "Doomhowl", "priest",
        ),
        best = listOf(
            fight("Lucifron", 1012.0, 98.0, "2026-11-03T21:00:00Z"),
            fight("Magmadar", 948.0, 91.0, "2026-11-03T21:20:00Z"),
            fight("Gehennas", 902.0, 76.0, "2026-11-03T21:40:00Z"),
            fight("Garr", 866.0, 64.0, "2026-10-27T21:00:00Z"),
            fight("Baron Geddon", 790.0, 42.0, "2026-10-27T21:30:00Z"),
        ),
        history = List(HISTORY) { fight("Lucifron", 1000.0, 90.0, "2026-11-03T21:00:00Z") },
        buildsSeen = listOf(
            BuildSeen("15/31/5", "Holy", "2026-11-03T21:00:00Z"),
            BuildSeen("21/30/0", "Holy", "2026-10-27T21:00:00Z"),
            BuildSeen("31/20/0", "Discipline", "2026-10-20T21:00:00Z"),
        ),
        gear = gear(elyraDoll),
        gearSpec = "priest-holy",
        supporter = true,
        score = 1830,
        tier = "epic",
        asOf = "2026-11-05",
        components = ScoreComponents(410, 520, 470, 330, 100),
        componentCaps = ScoreComponentCaps(500, 700, 600, 200, 100),
    )

    val felsnap = CharacterProfile(
        character = CharacterProfileHeader(FELSNAP_KEY, "us", "normal", "Felsnap", null, 34, "Dwarf", "Skyborne", "hunter"),
        buildsSeen = emptyList(),
        gear = gear(felsnapDoll),
        gearSpec = "hunter-beast-mastery",
        supporter = false,
    )

    private fun row(id: String, tag: String, at: Duration, vararg segments: FeedSegment) =
        FeedRow(id, "kind", ago(at), tag, crest = "priest", href = "/character/us/hardcore/elyra-duskvale", dateOnly = false, segments = segments.toList())

    private val name = FeedSegment("Elyra Duskvale", "name", "priest")

    private val elyraActivity = listOf(
        row("a1", "your_guild", Duration.ofMinutes(30), name, FeedSegment(" simmed 912 HPS, #3 of 14 healers in Cx66 Olympus (Holy, 60 raid)")),
        row("a2", "logs", Duration.ofHours(1), name, FeedSegment(" uploaded Molten Core, 10 bosses in 92 min")),
        row("a3", "rankings", Duration.ofHours(21), FeedSegment("Holy, 60:", "spec"), FeedSegment(" "), name, FeedSegment(" climbed 6 places to "), FeedSegment("#4", "spec")),
        row("a4", "logs", Duration.ofDays(2), name, FeedSegment(" downed Lucifron in Molten Core, the first kill on Forever")),
    )

    private fun own(key: String, name: String, level: Int, points: List<Int>) = AccountCharacter(
        key = key, region = "us", ruleset = key.split('/')[1], name = name, source = "blizzard", level = level,
        build = CharacterBuild("addon", ago(Duration.ofDays(1)), talents = BuildTalents(points = points)),
    )

    private val raid = RaidLine(
        id = 1, zone = "onyxia", zoneName = "Onyxia's Lair", title = "Onyxia's Lair", startsAt = now.plus(Duration.ofDays(33).plusHours(8)).toString(),
        status = "open", size = 40, guildTimezone = "America/Chicago",
        counts = RaidCounts(signedIn = 18, late = 3, tanks = 2, healers = 5, dps = 11, melee = 6, ranged = 5),
    )

    private val sims = HomeGuildSims(
        guildId = 66, guildName = "Cx66 Olympus", role = "healer", band = 60, preset = "raid", unit = "hps", of = 14,
        you = HomeLadderRow(3, "Elyra Duskvale", "priest", ELYRA_KEY, 912.0, simmedAt = ago(Duration.ofMinutes(30)), simId = "s1", you = true),
    )

    private val elyraHome = home(ELYRA_KEY, 60, "HPS", 128.9, elyraDoll)
    private val felsnapHome = home(FELSNAP_KEY, 30, "DPS", 7.7, felsnapDoll)

    val elyraOwnerExtras = CharacterExtras(
        home = elyraHome,
        blocks = HomeBlocks(specWeek = HomeSpecWeek("priest-holy", 60, "alliance", "healer", "A", null, 2, 9), guildSims = sims),
        nextRaid = raid,
        activity = elyraActivity,
        own = own(ELYRA_KEY, "Elyra Duskvale", 60, listOf(15, 31, 5)),
    )

    val elyraVisitorExtras = CharacterExtras(home = elyraHome, activity = elyraActivity.drop(1))

    val felsnapOwnerExtras = CharacterExtras(
        home = felsnapHome,
        blocks = HomeBlocks(specWeek = HomeSpecWeek("hunter-beast-mastery", 30, "alliance", "dps", "B", null, 6, 9)),
        own = own(FELSNAP_KEY, "Felsnap", 34, listOf(23, 0, 0)),
    )

    val felsnapVisitorExtras = CharacterExtras(home = felsnapHome)

    private const val HISTORY = 7
}
