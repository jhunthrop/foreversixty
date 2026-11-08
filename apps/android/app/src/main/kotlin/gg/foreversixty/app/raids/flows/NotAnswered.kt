package gg.foreversixty.app.raids.flows

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.home.flows.RaidIcons
import gg.foreversixty.app.model.RaidView
import gg.foreversixty.app.raids.logic.GameClass
import gg.foreversixty.app.raids.logic.RaidClock
import gg.foreversixty.app.raids.logic.ShortRole
import gg.foreversixty.app.raids.logic.SignupState
import java.time.Instant
import java.time.ZoneId

/** The raid page's filters (spec F3): Not answered preselected, then the answers, then everyone. */
enum class NotAnsweredFilter(val state: SignupState?) {
    NotAnswered(null),
    In(SignupState.In),
    Late(SignupState.Late),
    Bench(SignupState.Bench),
    Out(SignupState.Out),
    All(null),
}

data class NotAnsweredChip(val filter: NotAnsweredFilter, val title: String)

/** One raider: the name in the class colour over the class crest, the spec and last week. [marked] is a mark held on the row. */
data class NotAnsweredRow(
    val characterKey: String,
    val name: String,
    val className: String,
    val sub: String,
    val role: String?,
    val marked: SignupState? = null,
)

/** The mention: the short role's button (null when none of it is unanswered), Mention all, and the message it posts. */
data class MentionContent(val role: String?, val all: String, val preview: String)

/** What Not answered draws (mock raid-unanswered). */
data class NotAnsweredContent(
    val title: String,
    val icon: String,
    val facts: List<String>,
    val chips: List<NotAnsweredChip>,
    val filter: NotAnsweredFilter,
    val mention: MentionContent?,
    val listTitle: String,
    val listMeta: String?,
    val rows: List<NotAnsweredRow>,
    /** Only Not answered's rows carry Mark In, Bench and Discord DM. */
    val actionable: Boolean,
    val benchTitle: String?,
    val bench: List<NotAnsweredRow>,
    val benchNote: String?,
)

/**
 * Keep the raid full (spec F3): the unanswered with the short role first, the mention that names it,
 * rows marked on the page, and the short role's bench when the unanswered cannot close the gap.
 * [marks] are the marks held on their rows (Mark In, Bench) before they are sent.
 */
class NotAnsweredPresenter(
    private val res: Resources,
    private val clock: RaidClock,
    private val raid: RaidView,
    private val guild: String,
    private val marks: Map<String, SignupState>,
) {
    fun content(filter: NotAnsweredFilter): NotAnsweredContent {
        val short = shortRole()
        val unanswered = unansweredRows(short)
        val open = unanswered.filter { it.marked == null }
        val notAnswered = filter == NotAnsweredFilter.NotAnswered
        val bench = if (notAnswered) benchRows(short, open) else emptyList()
        val facts = listOfNotNull(guild, whenText(), short?.let { res.getString(R.string.home_officer_short, res.getString(it.lower), it.missing) })
        return NotAnsweredContent(
            title = raid.zoneName,
            icon = RaidIcons.of(raid),
            facts = facts,
            chips = chips(open.size),
            filter = filter,
            mention = if (notAnswered && open.isNotEmpty()) mention(open, short) else null,
            listTitle = label(filter),
            listMeta = if (notAnswered) listMeta(open.size, short) else null,
            rows = if (notAnswered) unanswered else signedRows(filter.state),
            actionable = notAnswered,
            benchTitle = if (bench.isEmpty()) null else short?.let { res.getString(R.string.raid_na_bench, res.getString(it.group)) },
            bench = bench,
            benchNote = if (bench.isEmpty()) null else benchNote(short, open, bench),
        )
    }

    private val total get() = raid.signups.size + raid.notSigned.size

    private fun label(filter: NotAnsweredFilter): String = when (filter) {
        NotAnsweredFilter.NotAnswered -> res.getString(R.string.raid_na_title)
        NotAnsweredFilter.All -> res.getString(R.string.raid_na_all)
        else -> res.getString(requireNotNull(filter.state).label)
    }

    private fun count(state: SignupState): Int =
        raid.signups.count { it.state == state.wire && it.characterKey !in marks } + marks.values.count { it == state }

    private fun chips(unanswered: Int): List<NotAnsweredChip> = NotAnsweredFilter.entries.map { filter ->
        val n = when (filter) {
            NotAnsweredFilter.NotAnswered -> unanswered
            NotAnsweredFilter.All -> total
            else -> count(requireNotNull(filter.state))
        }
        NotAnsweredChip(filter, res.getString(R.string.raid_na_filter, label(filter), n))
    }

    /** The role furthest short, less the raiders of it marked In here. */
    private fun shortRole(): ShortRole? {
        val short = ShortRole.of(raid) ?: return null
        val marked = marks.count { (key, state) -> state == SignupState.In && roleOf(key) == short.wire }
        return short.takeIf { it.missing > marked }?.copy(missing = short.missing - marked)
    }

    private fun roleOf(key: String): String? =
        raid.notSignedRaiders?.firstOrNull { it.characterKey == key }?.role ?: raid.signups.firstOrNull { it.characterKey == key }?.role

    private fun whenText(): String {
        val start = Instant.parse(raid.startsAt)
        val zone = runCatching { ZoneId.of(raid.guildTimezone) }.getOrDefault(clock.zone)
        return res.getString(R.string.home_raid_when, clock.shortDate(start), clock.guildClock(start, zone))
    }

    private fun kind(specName: String?, className: String): String {
        val cls = GameClass.of(className)?.let { res.getString(it.label) }.orEmpty()
        return listOfNotNull(specName?.takeIf { it.isNotBlank() }, cls.takeIf { it.isNotEmpty() }).joinToString(" ")
    }

    /** The unanswered, the short role first, then tanks, healers, damage, in the server's order. */
    private fun unansweredRows(short: ShortRole?): List<NotAnsweredRow> {
        fun rank(role: String?): Int = if (short != null && role == short.wire) -1 else ROLE_ORDER.indexOf(role ?: ShortRole.DPS).let { if (it < 0) ROLE_ORDER.size else it }
        return raid.notSignedRaiders.orEmpty().withIndex()
            .sortedWith(compareBy({ rank(it.value.role) }, { it.index }))
            .map { (_, raider) ->
                val spec = raid.notSigned.firstOrNull { it.characterKey == raider.characterKey }?.specName
                NotAnsweredRow(
                    raider.characterKey, raider.name, raider.className,
                    res.getString(R.string.raid_na_sub, kind(spec, raider.className), lastWeek(raider.lastWeek)),
                    raider.role, marks[raider.characterKey],
                )
            }
    }

    private fun lastWeek(state: String?): String {
        if (state == null) return res.getString(R.string.raid_na_new)
        val answer = SignupState.of(state) ?: return res.getString(R.string.raid_na_missed)
        return res.getString(R.string.raid_na_lastWeek, res.getString(answer.label))
    }

    /** The sign-ups in a state (All: everyone, the unanswered after them). */
    private fun signedRows(state: SignupState?): List<NotAnsweredRow> {
        val signed = raid.signups.filter { state == null || it.state == state.wire }.map {
            NotAnsweredRow(it.characterKey, it.name, it.className, kind(it.specName, it.className), it.role)
        }
        return if (state == null) signed + unansweredRows(shortRole()) else signed
    }

    /** The short role's bench, when its unanswered cannot close the gap. */
    private fun benchRows(short: ShortRole?, open: List<NotAnsweredRow>): List<NotAnsweredRow> {
        if (short == null || open.count { it.role == short.wire } >= short.missing) return emptyList()
        return raid.signups.filter { it.state == SignupState.Bench.wire && it.role == short.wire }.map {
            NotAnsweredRow(
                it.characterKey, it.name, it.className,
                res.getString(R.string.raid_na_sub, kind(it.specName, it.className), res.getString(R.string.raid_na_onBench)),
                it.role, marks[it.characterKey],
            )
        }
    }

    private fun benchNote(short: ShortRole?, open: List<NotAnsweredRow>, bench: List<NotAnsweredRow>): String? {
        short ?: return null
        val names = open.filter { it.role == short.wire }.map { it.name } + bench.filter { it.marked == null }.map { it.name }
        if (names.size < short.missing) return null
        return res.getString(R.string.raid_na_fills, list(names), res.getString(short.lower))
    }

    private fun listMeta(open: Int, short: ShortRole?): String =
        short?.let { res.getString(R.string.raid_na_meta, open, total, res.getString(it.lower)) } ?: res.getString(R.string.raid_na_metaPlain, open, total)

    private fun mention(open: List<NotAnsweredRow>, short: ShortRole?): MentionContent {
        val ofRole = short?.let { s -> open.filter { it.role == s.wire } }.orEmpty()
        val named = ofRole.ifEmpty { open }
        val names = named.joinToString(" ") { "@${it.name}" }
        val preview = if (short != null && ofRole.isNotEmpty()) {
            res.getString(R.string.raid_na_preview, names, raid.zoneName, whenText(), res.getString(short.lower))
        } else {
            res.getString(R.string.raid_na_previewAll, names, raid.zoneName, whenText())
        }
        return MentionContent(
            role = short?.takeIf { ofRole.isNotEmpty() }?.let { res.getString(R.string.raid_na_mentionRole, ShortRole.people(res, it.wire, ofRole.size)) },
            all = res.getString(R.string.raid_na_mentionAll, open.size),
            preview = preview,
        )
    }

    /** `Mirethil, Selka and Aurelith`. */
    private fun list(names: List<String>): String = when (names.size) {
        0 -> ""
        1 -> names.first()
        else -> res.getString(R.string.raid_list_and, names.dropLast(1).joinToString(", "), names.last())
    }

    private companion object {
        val ROLE_ORDER = listOf(ShortRole.TANK, ShortRole.HEALER, ShortRole.DPS)
    }
}
