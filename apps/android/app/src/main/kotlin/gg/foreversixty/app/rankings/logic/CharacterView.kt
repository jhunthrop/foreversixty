package gg.foreversixty.app.rankings.logic

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.bis.SpecCatalog
import gg.foreversixty.app.home.HomePresenter
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.items.ItemNumbers
import gg.foreversixty.app.model.CharacterFight
import gg.foreversixty.app.model.CharacterHome
import gg.foreversixty.app.model.CharacterProfile
import gg.foreversixty.app.model.FeedRow
import gg.foreversixty.app.model.HomeGuildSims
import gg.foreversixty.app.model.HomeSpecWeek
import gg.foreversixty.app.model.RaidLine
import gg.foreversixty.app.paperdoll.PaperdollSlot
import gg.foreversixty.app.raids.logic.RaidClock
import java.text.NumberFormat
import java.time.Duration
import java.time.Instant
import java.time.ZoneId
import java.time.format.DateTimeFormatter
import java.util.Locale
import kotlin.math.abs
import kotlin.math.roundToInt

/** Where a link on the character page goes; the route turns each into a screen or a web page. */
sealed interface CharacterLink {
    data object Planner : CharacterLink
    data object Sim : CharacterLink
    data object Bis : CharacterLink
    data object HowScored : CharacterLink
    data object History : CharacterLink
    data object Plans : CharacterLink
    data object Coach : CharacterLink
    data object AllFights : CharacterLink

    /** A site path (`/guild/us/normal/olympus`, `/tiers`), opened on the web. */
    data class Web(val path: String) : CharacterLink
}

/** One piece of the head's fact line: words, a link, or the supporter tag. */
sealed interface HeadFact {
    data class Words(val text: String) : HeadFact
    data class Link(val text: String, val link: CharacterLink) : HeadFact
    data class Tag(val text: String) : HeadFact
}

/** The head's answer (C2): the score beside the upgrades the band's list offers. */
data class HeadAnswer(val score: String?, val total: String, val unit: String)

/** A sentence with its one link (`Next: ...  Open best in slot`). */
data class LinkedLine(val text: String, val action: String?, val link: CharacterLink?)

/** One paperdoll cell (C4): the worn item, its level, and for the owner the gain to the band's pick. */
data class GearCell(
    val slot: PaperdollSlot,
    val label: String,
    /** `Item level 62`; null for an empty slot or an item the data does not know. */
    val level: String?,
    val gain: CellGain?,
)

/** The gain column: `+18.4 HPS` over `vs <pick>`, or `Best in slot`, or a word in place of a figure. */
data class CellGain(val value: String, val unit: String?, val versus: ItemLook?, val muted: Boolean = false)

/** The Forever Score section (C5): figures, the one line, the parts. */
data class ScoreView(
    val score: String,
    val tier: String,
    val tierQuality: Int?,
    val mostToGain: String?,
    val parts: List<Pair<String, String>>,
)

/** One best fight (C6). */
data class BestRow(val fight: CharacterFight, val boss: String, val sub: String, val percentile: Int?, val value: String, val unit: String)

/** One build seen (C7). */
data class BuildRow(val split: String, val sub: String)

/** The owner's extras under the fights (C8): History with the entitlement or its lock, Coach history with it. */
data class OwnerSections(val history: Boolean, val coach: Boolean)

/** Everything the loaded page draws, built from the card and its extras so the screen only lays it out. */
data class CharacterPageView(
    val name: String,
    val className: String?,
    val facts: List<HeadFact>,
    val answer: HeadAnswer?,
    val next: LinkedLine?,
    val standing: List<LinkedLine>,
    val gearMeta: String?,
    val cells: List<GearCell>,
    val score: ScoreView?,
    val noScoreLine: String,
    val best: List<BestRow>,
    val allFights: Int,
    val builds: List<BuildRow>,
    val activity: List<FeedRow>,
    val owner: OwnerSections?,
)

/** Builds the character page (design/specs/2026-11-06-character.md C2 to C22) from data: data in, words out. */
class CharacterPageComposer(private val res: Resources, private val specs: SpecCatalog, private val zone: ZoneId = ZoneId.systemDefault()) {
    private val rankings = RankingsCopy(res, zone)

    fun compose(profile: CharacterProfile, slots: List<PaperdollSlot>, extras: CharacterExtras, viewer: RankingsViewer, now: Instant): CharacterPageView {
        val owner = CharacterCard.isYours(profile, viewer)
        val home = extras.home?.takeIf { it.state == HOME_OK }
        val spec = pageSpec(profile)
        val total = home?.totalGain?.takeIf { it > 0.0 }
        return CharacterPageView(
            name = profile.character.name,
            className = profile.character.className?.lowercase(Locale.ROOT),
            facts = facts(profile, spec, if (owner) extras else null, now),
            answer = total?.let { HeadAnswer(profile.score?.let(::grouped), ItemNumbers.signedTenths(it), home.unitWord.orEmpty()) },
            next = if (owner && home != null) next(home, extras) else null,
            standing = if (owner) standing(extras, spec, profile, now) else emptyList(),
            gearMeta = extras.home?.band?.let(::bandLabel),
            cells = cells(slots, if (owner) home else null),
            score = score(profile),
            noScoreLine = res.getString(
                when {
                    profile.best.isNotEmpty() -> R.string.char_page_noScoreYet
                    owner -> R.string.char_page_noLogsOwner
                    else -> R.string.char_noLogs
                },
            ),
            best = profile.best.take(FIGHTS_SHOWN).map(::bestRow),
            allFights = profile.history.size,
            builds = profile.buildsSeen.sortedByDescending { it.firstSeen }.take(BUILDS_SHOWN).map { BuildRow(it.talentSplit, buildSub(it.spec, it.firstSeen)) },
            activity = extras.activity,
            owner = if (owner && profile.score != null) OwnerSections(history = viewer.history, coach = viewer.history) else null,
        )
    }

    /** The spec the page names: the gear's priced spec, else the newest fight's or build's. */
    private fun pageSpec(profile: CharacterProfile): String? = CharacterCard.specName(profile, specs) ?: profile.buildsSeen.firstOrNull()?.spec

    // The head (C3)

    private fun facts(profile: CharacterProfile, spec: String?, own: CharacterExtras?, now: Instant): List<HeadFact> {
        val c = profile.character
        val race = c.race?.let(HomePresenter::words)
        val cls = rankings.className(c.className?.lowercase(Locale.ROOT))
        val levelRace = listOfNotNull(c.level?.let { res.getString(R.string.char_page_level, it) }, race).joinToString(" ")
        val split = own?.own?.build?.talents?.points?.joinToString("/") ?: profile.buildsSeen.maxByOrNull { it.firstSeen }?.talentSplit
        val guild = c.guild?.name
        return listOfNotNull(
            levelRace.takeIf { it.isNotEmpty() }?.let(HeadFact::Words),
            listOfNotNull(spec, cls.takeIf { it.isNotEmpty() }).joinToString(" ").takeIf { it.isNotEmpty() }?.let(HeadFact::Words),
            split?.let { HeadFact.Link(it, CharacterLink.Planner) },
            guild?.let { HeadFact.Link(res.getString(R.string.char_page_guild, it), CharacterLink.Web(guildPath(c.region, c.ruleset, it))) },
            c.realm?.let(HeadFact::Words),
            HeadFact.Words(res.getString(R.string.char_page_chip, rankings.rulesetLabel(c.ruleset), c.region.uppercase(Locale.ROOT))),
            if (profile.supporter) HeadFact.Tag(res.getString(R.string.shared_premium_supporterChip)) else null,
            own?.home?.let { sync(it, now) }?.let(HeadFact::Words),
        )
    }

    /** `Addon · 1 day ago`: where the gear came from and when (owner only). */
    private fun sync(home: CharacterHome, now: Instant): String? {
        val at = home.syncedAt ?: return null
        val age = rankings.age(at, now).takeIf { it.isNotEmpty() } ?: return null
        return when (home.syncedFrom) {
            ADDON -> res.getString(R.string.char_page_syncAddon, age)
            BLIZZARD -> res.getString(R.string.char_page_syncBlizzard, age)
            else -> null
        }
    }

    // The summary (C11, C16 to C21)

    /** The one Next line: unspent talent points first, then the largest upgrade. */
    private fun next(home: CharacterHome, extras: CharacterExtras): LinkedLine {
        val top = home.upgrades.filter { it.modelled && (it.gain ?: 0.0) > 0.0 }.maxByOrNull { it.gain ?: 0.0 }
        val unspent = extras.own?.let { own -> unspentPoints(own.level, own.build?.talents?.points) } ?: 0
        if (unspent > 0) {
            val points = res.getQuantityString(R.plurals.char_page_talentPoints, unspent, unspent)
            val text = top?.pick?.name?.let { res.getString(R.string.char_page_nextTalentsThen, points, it) } ?: res.getString(R.string.char_page_nextTalents, points)
            return LinkedLine(text, res.getString(R.string.char_page_openPlannerNext), CharacterLink.Planner)
        }
        val pick = top?.pick ?: return LinkedLine(res.getString(R.string.char_page_nothingWaiting), null, null)
        return LinkedLine(res.getString(R.string.char_page_next, pick.name.orEmpty(), pick.sourceLine), res.getString(R.string.char_page_openBis), CharacterLink.Bis)
    }

    private fun standing(extras: CharacterExtras, spec: String?, profile: CharacterProfile, now: Instant): List<LinkedLine> {
        val guild = profile.character.guild?.name
        val guildLink = guild?.let { CharacterLink.Web(guildPath(profile.character.region, profile.character.ruleset, it)) }
        return listOfNotNull(
            extras.blocks?.guildSims?.let(::rankLine)?.let { LinkedLine(it, null, guildLink) },
            extras.blocks?.specWeek?.let { gradeLine(it, spec) }?.let { LinkedLine(it, null, CharacterLink.Web(TIERS_PATH)) },
            extras.nextRaid?.let { raidLine(it, now) }?.let { LinkedLine(it, null, guildLink) },
        )
    }

    /** C16: `#3 of 14 healers in Olympus XXVII, 60 raid`; null with no own row on the ladder. */
    private fun rankLine(sims: HomeGuildSims): String? {
        val own = sims.you ?: sims.rows.firstOrNull { it.you } ?: return null
        val role = res.getString(
            when (sims.role) {
                "healer" -> R.string.shared_standing_role_healer
                "tank" -> R.string.shared_standing_role_tank
                else -> R.string.shared_standing_role_dps
            },
        )
        val band = if (sims.band >= MAX_LEVEL) MAX_LEVEL.toString() else res.getString(R.string.char_page_bandRange, sims.band, sims.band + BAND_SPAN)
        val label = if (sims.preset == RAID) res.getString(R.string.char_page_bandRaid, band) else band
        return res.getString(R.string.shared_standing_rank, own.rank, sims.of, role, sims.guildName, label)
    }

    /** C19: `Holy is graded A this week`; a list with no letter reads its place. */
    private fun gradeLine(week: HomeSpecWeek, spec: String?): String? {
        spec ?: return null
        return week.grade?.let { res.getString(R.string.shared_standing_graded, spec, it) } ?: res.getString(R.string.shared_standing_place, spec, week.place, week.of)
    }

    /** C21: `Onyxia's Lair in 33 days 8 hours, 18 in`; null once it has started. */
    private fun raidLine(raid: RaidLine, now: Instant): String? {
        val start = runCatching { Instant.parse(raid.startsAt) }.getOrNull() ?: return null
        val left = Duration.between(now, start).takeIf { !it.isNegative } ?: return null
        val clock = RaidClock(res, zone).spanWords(left)
        val inCount = raid.counts.signedIn
        return if (inCount > 0) res.getString(R.string.shared_standing_raidSigned, raid.title, clock, inCount) else res.getString(R.string.shared_standing_raid, raid.title, clock)
    }

    // The gear (C4)

    private fun cells(slots: List<PaperdollSlot>, home: CharacterHome?): List<GearCell> {
        val bySlot = slots.associateBy { it.slot }
        return DOLL_ORDER.mapNotNull(bySlot::get).map { slot ->
            val look = slot.item
            GearCell(
                slot = slot,
                label = slot.slotName,
                level = look?.takeIf { it.name != null }?.itemLevel?.let { res.getString(R.string.char_page_itemLevel, it) },
                gain = home?.let { gain(slot.slot, it) },
            )
        }
    }

    private fun gain(slot: String, home: CharacterHome): CellGain? {
        if (home.alreadyBest.any { slotKey(it.slot) == slot }) return CellGain(res.getString(R.string.shared_paperdoll_best), null, null)
        val up = home.upgrades.firstOrNull { slotKey(it.slot) == slot } ?: return null
        val pick = ItemLook(up.pick.id, up.pick.name, up.pick.quality, up.pick.icon, up.pick.itemLevel)
        val value = up.gain
        return when {
            value == null || !up.modelled -> CellGain(res.getString(R.string.shared_paperdoll_notSimChecked), null, pick, muted = true)
            abs(value) < TIE -> CellGain(res.getString(R.string.shared_paperdoll_same), null, pick, muted = true)
            else -> CellGain(ItemNumbers.signedTenths(value), home.unitWord, pick)
        }
    }

    // The score (C5)

    private fun score(profile: CharacterProfile): ScoreView? {
        val score = profile.score ?: return null
        val bars = CharacterCard.bars(profile.components, profile.componentCaps)
        val most = bars.filter { it.cap > it.value }.maxByOrNull { it.cap - it.value }
        return ScoreView(
            score = grouped(score),
            tier = tierLabel(profile.tier),
            tierQuality = CharacterCard.tierQuality(profile.tier),
            mostToGain = most?.let { res.getString(R.string.char_page_mostToGain, res.getString(it.component.label), it.cap - it.value) },
            parts = bars.map { res.getString(it.component.label) to res.getString(R.string.char_page_partSub, it.value, it.cap) },
        )
    }

    private fun tierLabel(tier: String?): String = when (tier) {
        "common" -> res.getString(R.string.char_tier_common)
        "uncommon" -> res.getString(R.string.char_tier_uncommon)
        "rare" -> res.getString(R.string.char_tier_rare)
        "epic" -> res.getString(R.string.char_tier_epic)
        "legendary" -> res.getString(R.string.char_tier_legendary)
        else -> tier?.let(HomePresenter::words).orEmpty()
    }

    // The fights and builds (C6, C7)

    private fun bestRow(fight: CharacterFight): BestRow = BestRow(
        fight = fight,
        boss = fight.encounter,
        sub = listOfNotNull(dayMonth(fight.foughtAt, year = false), fight.spec).joinToString(" · "),
        percentile = fight.percentile?.roundToInt(),
        value = grouped(fight.value.roundToInt()),
        unit = if (fight.metric == HPS) res.getString(R.string.rank_unit_hps) else res.getString(R.string.rank_unit_dps),
    )

    private fun buildSub(spec: String, firstSeen: String): String {
        val seen = res.getString(R.string.char_page_firstSeen, dayMonth(firstSeen, year = true).orEmpty())
        return if (spec.isEmpty()) seen else "$spec · $seen"
    }

    /** `3 Nov`, or `3 Nov 2026` with the year. */
    private fun dayMonth(iso: String, year: Boolean): String? =
        runCatching { Instant.parse(iso) }.getOrNull()?.let { (if (year) DAY_MONTH_YEAR else DAY_MONTH).format(it.atZone(UTC)) }

    private fun bandLabel(band: Int): String =
        if (band >= MAX_LEVEL) res.getString(R.string.char_page_band60) else res.getString(R.string.char_page_band, band, band + BAND_SPAN)

    companion object {
        /** C6: the fights shown before the link to all of them; C7: the builds shown. */
        const val FIGHTS_SHOWN = 5
        const val BUILDS_SHOWN = 3
        private const val MAX_LEVEL = 60
        private const val BAND_SPAN = 9
        private const val TALENT_BASE_LEVEL = 9
        private const val HOME_OK = "ok"
        private const val ADDON = "addon"
        private const val BLIZZARD = "blizzard"
        private const val RAID = "raid"
        private const val HPS = "hps"

        /** A gain under a twentieth of a point reads as the same item (the web's `isTie`). */
        private const val TIE = 0.05
        private const val TIERS_PATH = "/tiers"
        private val UTC: ZoneId = ZoneId.of("UTC")
        private val DAY_MONTH: DateTimeFormatter = DateTimeFormatter.ofPattern("d MMM", Locale.UK)
        private val DAY_MONTH_YEAR: DateTimeFormatter = DateTimeFormatter.ofPattern("d MMM yyyy", Locale.UK)

        /** The doll's order: the left column, the weapons, then the right column (the web's phone paperdoll). */
        val DOLL_ORDER = listOf(
            "head", "neck", "shoulder", "back", "chest", "wrist", "main_hand", "off_hand", "ranged",
            "hands", "waist", "legs", "feet", "finger1", "finger2", "trinket1", "trinket2",
        )

        /** The upgrades' slot keys (`finger_1`) as the card's (`finger1`). */
        fun slotKey(slot: String): String = slot.replace("_1", "1").replace("_2", "2")

        /** Points a level earns (one a level from 10) less the points spent; 0 when unknown. */
        fun unspentPoints(level: Int?, points: List<Int>?): Int {
            if (level == null || points == null) return 0
            return (level - TALENT_BASE_LEVEL).coerceAtLeast(0) - points.sum()
        }

        fun guildPath(region: String, ruleset: String, guild: String): String =
            "/guild/$region/$ruleset/${guild.lowercase(Locale.ROOT).trim().replace(Regex("\\s+"), "-")}"

        private fun grouped(n: Int): String = NumberFormat.getIntegerInstance(Locale.US).format(n)
    }
}
