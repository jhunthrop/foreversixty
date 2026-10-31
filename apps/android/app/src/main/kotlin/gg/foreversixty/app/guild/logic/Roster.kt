package gg.foreversixty.app.guild.logic

import android.content.res.Resources
import androidx.annotation.StringRes
import gg.foreversixty.app.R
import gg.foreversixty.app.model.GuildHome
import gg.foreversixty.app.model.GuildMember
import gg.foreversixty.app.model.GuildRosterRow
import gg.foreversixty.app.raids.logic.GameClass
import gg.foreversixty.app.raids.logic.RaidCache
import gg.foreversixty.app.raids.logic.RaidLinks
import java.time.Instant

/** The role chips (spec F16.1), read from `members[].role` only. */
enum class RosterRole(val wire: String?, @param:StringRes val label: Int) {
    All(null, R.string.roster_chip_all),
    Tanks("tank", R.string.roster_chip_tanks),
    Healers("healer", R.string.roster_chip_healers),
    Dps("dps", R.string.roster_chip_dps),
}

/** The `Sort: {by}` menu, in the spec's order. */
enum class RosterSort(@param:StringRes val label: Int) {
    Rank(R.string.roster_sort_rank),
    ClassName(R.string.roster_sort_className),
    Name(R.string.roster_sort_name),
    Attendance(R.string.roster_sort_attendance),
    Level(R.string.roster_sort_level),
}

/** A guild rank as the API spells it, in rank order, with its word. */
enum class GuildRank(val wire: String, @param:StringRes val word: Int) {
    Leader("leader", R.string.roster_rank_leader),
    Officer("officer", R.string.roster_rank_officer),
    Member("member", R.string.roster_rank_member),
    ;

    companion object {
        fun of(wire: String): GuildRank = entries.firstOrNull { it.wire == wire } ?: Member
    }
}

/** What the right side of a row shows. */
sealed interface RosterStanding {
    data class Attended(val present: Int, val nights: Int) : RosterStanding
    data object NoNights : RosterStanding
    data object Waiting : RosterStanding
    data object NoAccount : RosterStanding
}

/** One roster row: a member joined to its roster row (attendance, supporter) by character key. */
data class RosterEntry(
    val name: String,
    val characterKey: String?,
    val className: String?,
    val level: Int?,
    val rank: GuildRank,
    val role: String?,
    val specName: String?,
    val specIcon: String?,
    val supporter: Boolean,
    val standing: RosterStanding,
) {
    /** A no-account name has no page, so no tap target. */
    val opensPage: Boolean get() = characterKey != null

    /** F11 is not built yet: the character's page on the site (`region/ruleset/name-slug`). */
    val webUrl: String? get() = characterKey?.let { "${RaidLinks.SITE}/character/$it" }
}

object RosterEntries {
    /** The window is the guild's last 8 logged nights; the bar never draws more. */
    const val MAX_NIGHTS = 8

    fun of(home: GuildHome): List<RosterEntry> {
        val rows = home.roster.associateBy { it.characterKey }
        return home.members.map { member -> entry(member, member.characterKey?.let(rows::get)) }
    }

    private fun entry(member: GuildMember, row: GuildRosterRow?) = RosterEntry(
        name = member.name,
        characterKey = member.characterKey,
        className = member.className ?: row?.className,
        level = member.level,
        rank = GuildRank.of(member.rank),
        role = member.role,
        specName = member.specName ?: row?.specName,
        specIcon = member.specIcon ?: row?.specIcon,
        supporter = row?.supporter == true,
        standing = standing(member, row),
    )

    private fun standing(member: GuildMember, row: GuildRosterRow?): RosterStanding {
        val nights = (row?.attendance?.nights ?: 0).coerceAtMost(MAX_NIGHTS)
        return when {
            member.state == "waiting" -> RosterStanding.Waiting
            member.characterKey == null -> RosterStanding.NoAccount
            nights == 0 -> RosterStanding.NoNights
            else -> RosterStanding.Attended((row?.attendance?.present ?: 0).coerceIn(0, nights), nights)
        }
    }
}

/** The chips, the search and the sort applied to the entries. */
data class RosterQuery(val role: RosterRole = RosterRole.All, val sort: RosterSort = RosterSort.Rank, val search: String = "")

object RosterList {
    fun apply(entries: List<RosterEntry>, query: RosterQuery): List<RosterEntry> {
        val needle = query.search.trim()
        return entries
            .filter { query.role.wire == null || it.role == query.role.wire }
            .filter { needle.isEmpty() || it.name.contains(needle, ignoreCase = true) }
            .sortedWith(comparator(query.sort))
    }

    private val byName: Comparator<RosterEntry> = compareBy(String.CASE_INSENSITIVE_ORDER) { it.name }

    /** Higher attendance first; a row without attendance (waiting, no account, no nights) last. */
    private val byAttendance: Comparator<RosterEntry> = compareByDescending<RosterEntry> { attendanceShare(it) }
        .thenByDescending { (it.standing as? RosterStanding.Attended)?.present ?: -1 }

    private fun attendanceShare(entry: RosterEntry): Double =
        (entry.standing as? RosterStanding.Attended)?.let { it.present.toDouble() / it.nights } ?: -1.0

    private fun comparator(sort: RosterSort): Comparator<RosterEntry> = when (sort) {
        RosterSort.Rank -> compareBy<RosterEntry> { it.rank.ordinal }.then(byAttendance).then(byName)
        RosterSort.ClassName -> compareBy<RosterEntry> { GameClass.rank(it.className.orEmpty()) }.then(byName)
        RosterSort.Name -> byName
        RosterSort.Attendance -> byAttendance.then(byName)
        RosterSort.Level -> compareByDescending<RosterEntry> { it.level ?: 0 }.then(byName)
    }
}

/** Every roster sentence built from data; the words live in the copy catalog. */
class RosterCopy(private val res: Resources) {
    /** The full spec name, else the class name. */
    fun spec(entry: RosterEntry): String =
        entry.specName ?: GameClass.of(entry.className.orEmpty())?.let { res.getString(it.label) }.orEmpty()

    /** `Protection Warrior · 60 · Guild Master`. */
    fun meta(entry: RosterEntry): String {
        val rank = res.getString(entry.rank.word)
        return entry.level?.let { res.getString(R.string.roster_meta, spec(entry), it, rank) }
            ?: res.getString(R.string.roster_metaNoLevel, spec(entry), rank)
    }

    /** `8 of 8`. */
    fun attendance(attended: RosterStanding.Attended): String =
        res.getString(R.string.roster_attendance, attended.present, attended.nights)

    /** Spec 6: `Kraggor, Protection Warrior, level 60, Guild Master, attended 8 of 8 recent raid nights`. */
    fun a11y(entry: RosterEntry): String = listOfNotNull(
        entry.name,
        spec(entry).takeIf { it.isNotEmpty() },
        entry.level?.let { res.getString(R.string.roster_a11y_level, it) },
        res.getString(entry.rank.word),
        when (val standing = entry.standing) {
            is RosterStanding.Attended -> res.getString(R.string.roster_a11y_attended, standing.present, standing.nights)
            RosterStanding.NoNights -> res.getString(R.string.roster_noNights)
            RosterStanding.Waiting -> res.getString(R.string.roster_a11y_waiting)
            RosterStanding.NoAccount -> res.getString(R.string.roster_noAccountTag)
        },
    ).joinToString(", ")
}

/** `GET /v1/guilds/{id}/home` for the roster and the Raids tab's `Roster` row, cache first. */
class RosterStore(val guild: Long, service: GuildService, cache: RaidCache, now: () -> Instant) {
    private val loader = CacheFirst("guild-$guild-home", GuildHome.serializer(), cache, now) { service.home(guild) }

    val state: Loadable<GuildHome> get() = loader.state

    /** `counts.on_roster`, once the home has answered (or is cached). */
    val onRoster: Int? get() = (state as? Loadable.Loaded)?.value?.counts?.onRoster

    suspend fun load() = loader.load()
}
