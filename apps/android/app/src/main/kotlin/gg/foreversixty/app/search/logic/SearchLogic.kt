package gg.foreversixty.app.search.logic

import gg.foreversixty.app.model.CharacterSearchRow
import gg.foreversixty.app.model.SearchGuide
import gg.foreversixty.app.model.SearchItem
import gg.foreversixty.app.model.SearchTalent
import java.util.Locale
import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

/** The scope chips (spec F12.1), in the chips' order. */
enum class SearchScope { All, Items, Characters, Guides, Talents }

/** The result groups, in the order `All` draws them (spec F12.2), each with the scope that lifts its cap. */
enum class SearchGroup(val scope: SearchScope) {
    Guides(SearchScope.Guides),
    Talents(SearchScope.Talents),
    Items(SearchScope.Items),
    Characters(SearchScope.Characters),
}

/** How a name matches the query, best first (spec F12.2): exact, prefix, word prefix, contains. */
enum class MatchTier { Exact, Prefix, WordPrefix, Contains }

/** A name prepared once when the index loads, so a keystroke is a pure in-memory filter. */
class SearchName(raw: String) {
    private val lower = raw.lowercase(Locale.ROOT)
    private val wordStarts: IntArray = lower.indices
        .filter { it > 0 && lower[it].isLetterOrDigit() && !lower[it - 1].isLetterOrDigit() }
        .toIntArray()

    /** Null when [query] (already normalised) is not in the name at all. */
    fun tier(query: String): MatchTier? = when {
        lower == query -> MatchTier.Exact
        lower.startsWith(query) -> MatchTier.Prefix
        wordStarts.any { lower.startsWith(query, it) } -> MatchTier.WordPrefix
        lower.contains(query) -> MatchTier.Contains
        else -> null
    }
}

/** One searchable thing and the identity Recent keeps it by; also what Recent stores on the phone. */
@Serializable
sealed interface SearchTarget {
    val key: String

    @Serializable
    @SerialName("item")
    data class Item(val item: SearchItem) : SearchTarget {
        override val key: String get() = "item:${item.id}"
    }

    @Serializable
    @SerialName("character")
    data class Character(val character: CharacterSearchRow) : SearchTarget {
        override val key: String get() = "character:${character.key}"
    }

    @Serializable
    @SerialName("guide")
    data class Guide(val guide: SearchGuide) : SearchTarget {
        override val key: String get() = "guide:${guide.className}-${guide.spec}"
    }

    @Serializable
    @SerialName("talent")
    data class Talent(val talent: SearchTalent) : SearchTarget {
        override val key: String get() = "talent:${talent.id}"
    }
}

/**
 * `search.json` prepared for typing: names lowercased with their word starts, built once off the
 * main thread. [guideTitle] names a guide as its row does (`Elemental Shaman`), so a query matches
 * what the player reads.
 */
class PreparedIndex(
    val build: String,
    private val itemUsesShards: Int,
    items: List<SearchItem>,
    talents: List<SearchTalent>,
    guides: List<SearchGuide>,
    guideTitle: (SearchGuide) -> String,
) {
    internal val items: List<Pair<SearchItem, SearchName>> = items.map { it to SearchName(it.name) }
    internal val talents: List<Pair<SearchTalent, SearchName>> = talents.map { it to SearchName(it.name) }
    internal val guides: List<Pair<SearchGuide, SearchName>> = guides.map { it to SearchName(guideTitle(it)) }
    private val byId: Map<Int, SearchItem> = items.associateBy { it.id }

    fun item(id: Int): SearchItem? = byId[id]

    /** The item-uses file that holds [id]: `id % item_uses_shards` (app-v1.md). */
    fun shardOf(id: Int): Int = Math.floorMod(id, itemUsesShards.coerceAtLeast(1))
}

/**
 * What leaves the app for the website until its screen ships (UX S5): guides until F9, characters
 * and Rankings until F11. Drop an entry as its lane lands; the rows' external-link glyph follows.
 */
object WebHandoff {
    /**
     * [rankingsInApp]: F11 is on, so characters (F11.4) and Rankings (F11.1) open in the app;
     * [guidesInApp]: F9 is on, so a guide opens the native guide (lane F15).
     */
    fun opensWebsite(target: SearchTarget, rankingsInApp: Boolean = false, guidesInApp: Boolean = false): Boolean = when (target) {
        is SearchTarget.Guide -> !guidesInApp
        is SearchTarget.Character -> !rankingsInApp
        is SearchTarget.Item, is SearchTarget.Talent -> false
    }

    fun rankingsOpenWebsite(rankingsInApp: Boolean): Boolean = !rankingsInApp

    fun guidesOpenWebsite(guidesInApp: Boolean): Boolean = !guidesInApp
}

/** The phone-side matches for one query, each group in its final order and uncapped. */
data class LocalResults(
    val guides: List<SearchGuide> = emptyList(),
    val talents: List<SearchTalent> = emptyList(),
    val items: List<SearchItem> = emptyList(),
)

/** One drawn group: its rows (capped under `All`) and how many matched in all. */
data class SearchSection(val group: SearchGroup, val rows: List<SearchTarget>, val total: Int) {
    /** `Show all {n}` shows when the cap hid rows. */
    val showsAll: Boolean get() = total > rows.size
}

/** The search rules (spec F12.2), pure so the tests read them. iOS: SearchLogic.swift. */
object SearchLogic {
    /** `All` shows up to five rows per group. */
    const val GROUP_CAP = 5

    /** Characters come from the API, up to eight. */
    const val CHARACTER_LIMIT = 8

    /** Characters need at least two letters. */
    const val MIN_CHARACTER_QUERY = 2

    /** Characters are asked 250 ms after the last keystroke. */
    const val CHARACTER_DEBOUNCE_MILLIS = 250L

    /** Trimmed, lowercased, inner spaces collapsed. */
    fun normalize(raw: String): String = raw.trim().lowercase(Locale.ROOT).replace(WHITESPACE, " ")

    /** Items, guides and talents matching [query] (normalised); [classSlug] is the current character's class. */
    fun filter(index: PreparedIndex, query: String, classSlug: String?): LocalResults {
        if (query.isEmpty()) return LocalResults()
        val items = ranked(index.items, query, compareBy<SearchItem> { if (classSlug != null && classSlug in it.classes) 0 else 1 }
            .thenByDescending { it.quality }
            .thenByDescending { it.itemLevel }
            .thenBy { it.name })
        val talents = ranked(index.talents, query, compareBy<SearchTalent> { if (it.className == classSlug) 0 else 1 }.thenBy { it.name })
        val guides = ranked(index.guides, query, compareBy<SearchGuide> { if (it.className == classSlug) 0 else 1 }.thenBy { it.specName })
        return LocalResults(guides, talents, items)
    }

    private fun <T> ranked(entries: List<Pair<T, SearchName>>, query: String, within: Comparator<T>): List<T> =
        entries.mapNotNull { (value, name) -> name.tier(query)?.let { value to it } }
            .sortedWith(compareBy<Pair<T, MatchTier>> { it.second.ordinal }.thenComparator { a, b -> within.compare(a.first, b.first) })
            .map { it.first }

    /** Whether characters are asked for [query] (normalised) under [scope]. */
    fun asksCharacters(query: String, scope: SearchScope): Boolean =
        query.length >= MIN_CHARACTER_QUERY && (scope == SearchScope.All || scope == SearchScope.Characters)

    /** The API's rows with the current character's region and ruleset first, up to eight. */
    fun orderCharacters(rows: List<CharacterSearchRow>, region: String?, ruleset: String?): List<CharacterSearchRow> =
        rows.sortedBy { if (it.region == region && it.ruleset == ruleset) 0 else 1 }.take(CHARACTER_LIMIT)

    /**
     * The groups to draw: under `All` every group with a match, five rows each, in the order
     * Guides, Talents, Items, Characters; under a scope only its group, uncapped.
     */
    fun sections(local: LocalResults, characters: List<CharacterSearchRow>, scope: SearchScope): List<SearchSection> {
        val all = listOf(
            SearchGroup.Guides to local.guides.map { SearchTarget.Guide(it) },
            SearchGroup.Talents to local.talents.map { SearchTarget.Talent(it) },
            SearchGroup.Items to local.items.map { SearchTarget.Item(it) },
            SearchGroup.Characters to characters.map { SearchTarget.Character(it) },
        )
        return all.filter { (group, rows) -> rows.isNotEmpty() && (scope == SearchScope.All || group.scope == scope) }
            .map { (group, rows) -> SearchSection(group, if (scope == SearchScope.All) rows.take(GROUP_CAP) else rows, rows.size) }
    }

    private val WHITESPACE = Regex("\\s+")
}
