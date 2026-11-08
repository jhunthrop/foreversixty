package gg.foreversixty.app.home.flows

import android.content.res.Resources
import androidx.compose.ui.text.AnnotatedString
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.withStyle
import gg.foreversixty.app.R
import gg.foreversixty.app.auth.session.RelativeAge
import gg.foreversixty.app.bis.BisDataSource
import gg.foreversixty.app.bis.RaidsOpening
import gg.foreversixty.app.bis.SpecCatalog
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.flows.ClockPart
import gg.foreversixty.app.flows.FlowLead
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.items.ItemQuality
import gg.foreversixty.app.model.FeedRow
import gg.foreversixty.app.model.FeedSegment
import gg.foreversixty.app.model.KeyDate
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiEndpoints
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.phrase
import java.time.Duration
import java.time.Instant
import java.time.ZoneOffset
import java.time.format.DateTimeFormatter
import java.util.Locale

/** What signed-out Home reads (R3): the site feed, the build its icons come from, the key dates, the specs. */
interface HomeOutService {
    val specs: SpecCatalog

    /** `GET /v1/feed`, the first rows; null when it could not load. */
    suspend fun feed(): List<FeedRow>?
    suspend fun iconBuild(): String?
    suspend fun keyDates(): List<KeyDate>
}

class LiveHomeOutService(
    private val client: ApiClient,
    private val data: BisDataSource,
    override val specs: SpecCatalog,
) : HomeOutService {
    override suspend fun feed(): List<FeedRow>? =
        (client.send(ApiEndpoints.getFeed(filter = FEED_FILTER, limit = FEED_ROWS)) as? ApiResult.Success)?.data?.rows?.take(FEED_ROWS)

    override suspend fun iconBuild(): String? = data.activeBuild().takeIf { it.isNotEmpty() }

    override suspend fun keyDates(): List<KeyDate> = data.keyDates()

    companion object {
        /** The site's FEED_ROWS (web/src/lib/home/feed.ts). */
        const val FEED_ROWS = 7
        private const val FEED_FILTER = "all"
    }
}

/** A class to pick: its crest, name and specs (`Arms, Fury, Protection`). */
data class ClassPick(val slug: String, val name: String, val specs: String)

/** A first-tier raid under the countdown (web/src/lib/first-raids.ts). */
data class FirstRaid(val name: String, val icon: String)

/** The first raids' countdown (web RaidsCountdown): the eyebrow, the clock while ahead, the three raids. */
data class FirstRaidsCard(val eyebrow: String, val clock: List<ClockPart>, val clockLabel: String, val raids: List<FirstRaid>)

/** One line of `Forever right now`. */
data class FeedLine(val id: String, val lead: FlowLead, val sentence: AnnotatedString, val time: String, val kind: String, val href: String)

sealed interface HomeFeed {
    data object Loading : HomeFeed
    data object Failed : HomeFeed
    data object Empty : HomeFeed
    data class Rows(val lines: List<FeedLine>) : HomeFeed
}

/** Signed-out Home (mock android-home-out). */
data class HomeOutView(
    val facts: List<String>,
    val classes: List<ClassPick>,
    val raids: FirstRaidsCard?,
    val feed: HomeFeed,
    val build: String?,
)

/** Signed-out Home's words from what has loaded; [now] pins the clock and the feed's times. */
class HomeOutComposer(private val res: Resources, private val now: Instant) {
    fun view(specs: SpecCatalog, dates: List<KeyDate>, rows: List<FeedRow>?, loading: Boolean, build: String?): HomeOutView {
        val opening = RaidsOpening.from(dates)
        val facts = listOfNotNull(
            res.getString(R.string.shared_home_signedOutMeta),
            opening?.takeIf { it.isUpcoming(now.toEpochMilli()) }?.let { res.getString(R.string.shared_home_firstRaids, format(DAY_MONTH, it.startsMillis)) },
        )
        return HomeOutView(facts, classes(specs), opening?.let(::raids), feed(rows, loading, build), build)
    }

    fun classes(specs: SpecCatalog): List<ClassPick> = CLASS_ORDER.map { slug ->
        ClassPick(slug, slug.replaceFirstChar(Char::titlecase), specs.specsOf(slug).joinToString(", ") { it.name })
    }

    private fun raids(opening: RaidsOpening): FirstRaidsCard {
        if (!opening.isUpcoming(now.toEpochMilli())) return FirstRaidsCard(res.getString(R.string.shared_home_raidsOpen), emptyList(), "", FIRST_RAIDS)
        val until = Duration.between(now, Instant.ofEpochMilli(opening.startsMillis))
        val days = until.toDays().toInt()
        val hours = (until.toHours() % HOURS_PER_DAY).toInt()
        val minutes = (until.toMinutes() % MINUTES_PER_HOUR).toInt()
        val clock = listOf(
            ClockPart(twoDigits(days), res.getString(R.string.home_out_days)),
            ClockPart(twoDigits(hours), res.getString(R.string.home_raid_hours)),
            ClockPart(twoDigits(minutes), res.getString(R.string.home_raid_min)),
        )
        return FirstRaidsCard(
            res.getString(R.string.shared_home_raidsEyebrow, format(WEEKDAY_DAY_MONTH, opening.startsMillis)),
            clock,
            res.getString(R.string.home_out_clock, days, hours, minutes),
            FIRST_RAIDS,
        )
    }

    private fun feed(rows: List<FeedRow>?, loading: Boolean, build: String?): HomeFeed = when {
        rows != null && rows.isEmpty() -> HomeFeed.Empty
        rows != null -> HomeFeed.Rows(rows.map { line(it, build) })
        loading -> HomeFeed.Loading
        else -> HomeFeed.Failed
    }

    fun line(row: FeedRow, build: String?): FeedLine {
        val time = if (row.dateOnly) {
            parse(row.at)?.let { format(WEEKDAY_DAY_MONTH, it.toEpochMilli()) }.orEmpty()
        } else {
            parse(row.at)?.let { RelativeAge.of(it, now).phrase(res) }.orEmpty()
        }
        return FeedLine(row.id, lead(row, build), sentence(row.segments), time, res.getString(kind(row.tag)), row.href)
    }

    private fun lead(row: FeedRow, build: String?): FlowLead = when {
        row.crest != null -> FlowLead.Crest(row.crest.orEmpty())
        build != null -> FlowLead.Item(ItemLook(0, null, row.quality, row.icon), build)
        else -> FlowLead.Icon(row.icon, null)
    }

    private fun format(pattern: DateTimeFormatter, millis: Long): String = pattern.format(Instant.ofEpochMilli(millis).atZone(ZoneOffset.UTC))

    private fun parse(at: String): Instant? = runCatching { Instant.parse(at) }.getOrNull()

    companion object {
        /** The nine classes in the game's order. */
        val CLASS_ORDER = listOf("warrior", "paladin", "hunter", "rogue", "priest", "shaman", "mage", "warlock", "druid")

        val FIRST_RAIDS = listOf(
            FirstRaid("Barrow Deeps", "spell_shadow_raisedead"),
            FirstRaid("Hyjal Summit", "spell_nature_timestop"),
            FirstRaid("Onyxia's Lair", "inv_misc_head_dragon_black"),
        )

        // The site reads these dates in UTC (web/src/lib/home/feed.ts dayText).
        private val DAY_MONTH = DateTimeFormatter.ofPattern("d MMM", Locale.UK)
        private val WEEKDAY_DAY_MONTH = DateTimeFormatter.ofPattern("EEE d MMM", Locale.UK)
        private const val HOURS_PER_DAY = 24
        private const val MINUTES_PER_HOUR = 60

        private fun twoDigits(n: Int) = n.toString().padStart(2, '0')

        private fun kind(tag: String): Int = when (tag) {
            "your_spec" -> R.string.shared_feed_tag_yourSpec
            "logs" -> R.string.shared_feed_tag_logs
            "rankings" -> R.string.shared_feed_tag_rankings
            "best_in_slot" -> R.string.shared_feed_tag_bestInSlot
            "your_guild" -> R.string.shared_feed_tag_yourGuild
            "patch" -> R.string.shared_feed_tag_patch
            else -> R.string.shared_feed_tag_guides
        }

        /** The site's FeedSentence: names in their class colour, items in their quality colour, strong words bold. */
        fun sentence(segments: List<FeedSegment>): AnnotatedString = buildAnnotatedString {
            segments.forEach { segment ->
                val color = when (segment.style) {
                    "name" -> segment.className?.let(::classColor) ?: FsColor.strong
                    "item" -> ItemQuality(segment.quality ?: 1).text
                    "strong" -> FsColor.strong
                    else -> null
                }
                if (color == null) append(segment.text) else withStyle(SpanStyle(color = color, fontWeight = FontWeight.Bold)) { append(segment.text) }
            }
        }
    }
}
