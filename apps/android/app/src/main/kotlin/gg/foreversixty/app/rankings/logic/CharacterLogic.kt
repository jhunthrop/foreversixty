package gg.foreversixty.app.rankings.logic

import gg.foreversixty.app.navigation.DeepLink
import androidx.annotation.StringRes
import gg.foreversixty.app.R
import gg.foreversixty.app.bis.SpecCatalog
import gg.foreversixty.app.model.CharacterProfile
import gg.foreversixty.app.model.ScorePoint
import gg.foreversixty.app.model.ScoreComponentCaps
import gg.foreversixty.app.model.ScoreComponents
import java.util.Locale

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

/** One score part: the component's value and its cap from the API. */
data class ScoreBar(val component: ScoreComponent, val value: Int, val cap: Int)

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

    private const val HALF = 0.5f

    fun isYours(profile: CharacterProfile, viewer: RankingsViewer): Boolean = profile.character.key in viewer.accountKeys

    /** The spec the card is about: the gear's spec from the catalogue, else the newest best's spec. */
    fun specName(profile: CharacterProfile, specs: SpecCatalog): String? =
        profile.gearSpec?.let { specs.spec(it)?.name } ?: profile.best.firstNotNullOfOrNull { it.spec } ?: profile.history.firstNotNullOfOrNull { it.spec }

    /** The items file to read: the class slug. */
    fun classSlug(profile: CharacterProfile): String? = profile.character.className?.lowercase(Locale.ROOT)

    /**
     * The planner on the character's class (its talent split, `Talents in planner`, a build seen), with
     * no code and no race (rankings carry no build code). Null with no class.
     */
    fun plannerLink(profile: CharacterProfile): DeepLink.Planner? = classSlug(profile)?.let { DeepLink.Planner(it) }

    /** The web page a share and the share sheet open: `/character/{region}/{ruleset}/{name-slug}`. */
    fun shareUrl(route: CharacterRoute): String = "$SITE/character/${route.region}/${route.ruleset}/${route.slug}"

    fun shareUrl(key: String): String? = CharacterRoute.ofKey(key)?.let(::shareUrl)

    fun historyUrl(route: CharacterRoute): String = shareUrl(route) + HISTORY_ANCHOR

    /** A site path on the web (`/tiers`), or an address that already is one. */
    fun webUrl(path: String): String = if (path.startsWith("http")) path else SITE + path

    const val LOGIN_PATH = "/login"

    /** Where a page link opens on the web; the planner's in-app screen is the route's. */
    fun linkUrl(link: CharacterLink, route: CharacterRoute): String = when (link) {
        CharacterLink.Planner -> webUrl("/planner")
        CharacterLink.Sim -> webUrl("/sim")
        CharacterLink.Bis -> webUrl("/bis")
        CharacterLink.HowScored -> HOW_URL
        CharacterLink.History -> historyUrl(route)
        CharacterLink.Coach -> shareUrl(route) + COACH_ANCHOR
        CharacterLink.Plans -> webUrl("/premium")
        CharacterLink.AllFights -> webUrl("/logs")
        is CharacterLink.Web -> webUrl(link.path)
    }

    private const val COACH_ANCHOR = "#coach"

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
}
