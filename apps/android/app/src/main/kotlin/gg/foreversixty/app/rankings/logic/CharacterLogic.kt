package gg.foreversixty.app.rankings.logic

import gg.foreversixty.app.navigation.DeepLink
import android.content.res.Resources
import androidx.annotation.StringRes
import gg.foreversixty.app.R
import gg.foreversixty.app.bis.SpecCatalog
import gg.foreversixty.app.home.HomePresenter
import gg.foreversixty.app.model.CharacterFight
import gg.foreversixty.app.model.CharacterProfile
import gg.foreversixty.app.model.ScoreComponentCaps
import gg.foreversixty.app.model.ScoreComponents
import gg.foreversixty.app.model.ScorePoint
import java.text.NumberFormat
import java.time.Instant
import java.time.LocalDate
import java.time.ZoneId
import java.time.format.DateTimeFormatter
import java.util.Locale
import kotlin.math.roundToInt

/** A character page's address: the path of `GET /v1/characters/{region}/{ruleset}/{name}`. */
data class CharacterRoute(val region: String, val ruleset: String, val name: String) {
    /** The key's last segment (`voltaris`): the web path's name. */
    val slug: String get() = name.lowercase(Locale.ROOT).replace(' ', '-')

    companion object {
        /** From a `<region>/<ruleset>/<name-slug>` key, with the [name] to show while it loads. */
        fun ofKey(key: String, name: String? = null): CharacterRoute? {
            val parts = key.split('/')
            if (parts.size != KEY_PARTS) return null
            return CharacterRoute(parts[0], parts[1], name ?: parts[2])
        }

        private const val KEY_PARTS = 3
    }
}

/** The five Forever Score components, in the card's order. */
enum class ScoreComponent(@param:StringRes val label: Int) {
    Gear(R.string.char_component_gear),
    Experience(R.string.char_component_experience),
    Parses(R.string.char_component_parses),
    Attendance(R.string.char_component_attendance),
    Discipline(R.string.char_component_discipline),
}

/** One 8 dp bar: the component's value over its cap from the API. */
data class ScoreBar(val component: ScoreComponent, val value: Int, val cap: Int) {
    val fraction: Float get() = if (cap <= 0) 0f else (value.toFloat() / cap).coerceIn(0f, 1f)
}

/** The rules under a character page (spec F11.4 to F11.7). */
object CharacterCard {
    private const val SITE = "https://foreversixty.gg"
    const val HOW_URL = "$SITE/ratings"
    private const val HISTORY_ANCHOR = "#history"

    /** The five bars, or none with no score: a character with no rated fight shows no zero. */
    fun bars(components: ScoreComponents?, caps: ScoreComponentCaps?): List<ScoreBar> {
        if (components == null || caps == null) return emptyList()
        return listOf(
            ScoreBar(ScoreComponent.Gear, components.gear, caps.gear),
            ScoreBar(ScoreComponent.Experience, components.experience, caps.experience),
            ScoreBar(ScoreComponent.Parses, components.parses, caps.parses),
            ScoreBar(ScoreComponent.Attendance, components.attendance, caps.attendance),
            ScoreBar(ScoreComponent.Discipline, components.discipline, caps.discipline),
        )
    }

    /** The sparkline's points, oldest first, each 0 (lowest) to 1 (highest); flat at the middle when all equal. */
    fun sparkline(history: List<ScorePoint>): List<Float> {
        val scores = history.sortedBy { it.date }.map { it.score }
        val low = scores.minOrNull() ?: return emptyList()
        val high = scores.max()
        if (high == low) return scores.map { HALF }
        return scores.map { (it - low).toFloat() / (high - low) }
    }

    fun hasScore(profile: CharacterProfile): Boolean = profile.score != null

    fun isYours(profile: CharacterProfile, viewer: RankingsViewer): Boolean = profile.character.key in viewer.accountKeys

    /** `Parse history` shows only on your own character with `entitlements.history`. */
    fun showsParseHistory(profile: CharacterProfile, viewer: RankingsViewer): Boolean = viewer.history && isYours(profile, viewer)

    /** The spec the card is about: the gear's spec from the catalogue, else the newest best's spec. */
    fun specName(profile: CharacterProfile, specs: SpecCatalog): String? =
        profile.gearSpec?.let { specs.spec(it)?.name } ?: profile.best.firstNotNullOfOrNull { it.spec } ?: profile.history.firstNotNullOfOrNull { it.spec }

    /** The items file to read: the class slug. */
    fun classSlug(profile: CharacterProfile): String? = profile.character.className?.lowercase(Locale.ROOT)

    /**
     * Builds seen's `Open in planner` (lane F15, F13 follow-up): the planner on the character's class,
     * with no code and no race (rankings carry no build code). Null with no class or no builds seen.
     */
    fun plannerLink(profile: CharacterProfile): DeepLink.Planner? =
        classSlug(profile)?.takeIf { profile.buildsSeen.isNotEmpty() }?.let { DeepLink.Planner(it) }

    /** The web page a share and the share sheet open: `/character/{region}/{ruleset}/{name-slug}`. */
    fun shareUrl(route: CharacterRoute): String = "$SITE/character/${route.region}/${route.ruleset}/${route.slug}"

    fun shareUrl(key: String): String? = CharacterRoute.ofKey(key)?.let(::shareUrl)

    fun historyUrl(route: CharacterRoute): String = shareUrl(route) + HISTORY_ANCHOR

    /** The tier's item-quality number, so the tag is drawn in that quality's colour. */
    fun tierQuality(tier: String?): Int? = when (tier) {
        "poor" -> 0
        "common" -> 1
        "uncommon" -> 2
        "rare" -> 3
        "epic" -> 4
        "legendary" -> 5
        else -> null
    }

    private const val HALF = 0.5f
}

/** Every character-page sentence built from data; the words live in the copy catalog. */
class CharacterCopy(private val res: Resources, private val zone: ZoneId = ZoneId.systemDefault()) {
    private val rankings = RankingsCopy(res, zone)

    /** `Orc Elemental Shaman · Level 60`. */
    fun descriptor(profile: CharacterProfile, specName: String?): String {
        val c = profile.character
        val race = c.race?.let(HomePresenter::words).orEmpty()
        val cls = rankings.className(c.className?.lowercase(Locale.ROOT))
        val level = c.level
        val text = if (level != null) {
            res.getString(R.string.char_descriptor, race, specName.orEmpty(), cls, level)
        } else {
            res.getString(R.string.char_descriptorNoLevel, race, specName.orEmpty(), cls)
        }
        return RankingsCopy.collapse(text)
    }

    /** `Hordebreakers · US Normal`. */
    fun where(profile: CharacterProfile): String {
        val c = profile.character
        val region = c.region.uppercase(Locale.ROOT)
        val ruleset = rankings.rulesetLabel(c.ruleset)
        return c.guild?.name?.let { res.getString(R.string.char_where, it, region, ruleset) } ?: res.getString(R.string.char_whereNoGuild, region, ruleset)
    }

    /** `1,830`. */
    fun score(score: Int): String = NumberFormat.getIntegerInstance(Locale.US).format(score)

    fun tier(tier: String?): String = when (tier) {
        "common" -> res.getString(R.string.char_tier_common)
        "uncommon" -> res.getString(R.string.char_tier_uncommon)
        "rare" -> res.getString(R.string.char_tier_rare)
        "epic" -> res.getString(R.string.char_tier_epic)
        "legendary" -> res.getString(R.string.char_tier_legendary)
        else -> tier?.let(HomePresenter::words).orEmpty()
    }

    /** `As of 17 Dec`, from the API's day. */
    fun asOf(date: String?): String? = date?.let { d ->
        runCatching { LocalDate.parse(d) }.getOrNull()?.let { res.getString(R.string.char_asOf, DAY_MONTH.format(it)) }
    }

    /** `Forever Score 1,830, Epic. Gear 412 of 500, …` */
    fun scoreLabel(profile: CharacterProfile): String? {
        val score = profile.score ?: return null
        val bars = CharacterCard.bars(profile.components, profile.componentCaps)
        val parts = listOf(res.getString(R.string.char_a11y_score, score(score), tier(profile.tier))) +
            bars.map { res.getString(R.string.char_a11y_component, res.getString(it.component.label), it.value, it.cap) }
        return parts.joinToString(" ")
    }

    /** The unit word under a fight's value. */
    fun unit(metric: String): String = Measure.of(metric)?.let(rankings::unit) ?: metric.uppercase(Locale.ROOT)

    /** A fight's value as the board prints it. */
    fun value(fight: CharacterFight): String = when (Measure.of(fight.metric)) {
        Measure.Execution -> "${(fight.value * PERCENT).roundToInt()}%"
        else -> fight.value.roundToInt().toString()
    }

    /** `Normal · 3 d ago`. */
    fun bestMeta(profile: CharacterProfile, fight: CharacterFight, now: Instant): String =
        res.getString(R.string.char_bestMeta, rankings.rulesetLabel(profile.character.ruleset), rankings.age(fight.foughtAt, now))

    /** `Kill · 3 d ago` or `Wipe · 3 d ago`. */
    fun fightMeta(fight: CharacterFight, now: Instant): String =
        res.getString(R.string.char_fightMeta, res.getString(if (fight.kill == false) R.string.char_wipe else R.string.char_kill), rankings.age(fight.foughtAt, now))

    /** `Elemental · first seen 10 Dec`. */
    fun buildMeta(spec: String, firstSeen: String): String {
        val date = runCatching { Instant.parse(firstSeen) }.getOrNull()?.let(rankings::shortDate).orEmpty()
        return res.getString(R.string.char_buildMeta, spec, date)
    }

    companion object {
        private const val PERCENT = 100.0
        private val DAY_MONTH: DateTimeFormatter = DateTimeFormatter.ofPattern("d MMM", Locale.US)
    }
}
