package gg.foreversixty.app.activity

import androidx.compose.runtime.Composable
import androidx.compose.ui.text.AnnotatedString
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.withStyle
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.flows.FlowRow
import gg.foreversixty.app.items.ItemQuality
import gg.foreversixty.app.model.ActivityRow
import java.util.Locale

/**
 * An Activity row's body as the feed styles it (`segments`): names in class colour, items in their
 * quality colour, strong figures; the plain body when the row has none (iOS ActivityBody).
 */
object ActivityBody {
    sealed interface Tone {
        data object Plain : Tone
        data object Strong : Tone
        data class ClassName(val slug: String?) : Tone
        data class Quality(val value: Int?) : Tone
    }

    data class Part(val text: String, val tone: Tone)

    private const val NAME_STYLE = "name"
    private const val ITEM_STYLE = "item"
    private const val STRONG_STYLE = "strong"

    /** The parts to draw; null when the row carries no segments. */
    fun parts(row: ActivityRow): List<Part>? = row.segments?.takeIf { it.isNotEmpty() }?.map { segment ->
        val tone = when (segment.style) {
            NAME_STYLE -> Tone.ClassName(segment.className)
            ITEM_STYLE -> Tone.Quality(segment.quality)
            STRONG_STYLE -> Tone.Strong
            else -> Tone.Plain
        }
        Part(segment.text, tone)
    }

    /** The row's figure (sim DPS, the guild ladder's), one decimal. */
    fun figure(row: ActivityRow): String? = row.figure?.let { String.format(Locale.US, "%.1f", it) }

    /** The styled line, then [suffix] after a dot (Home's age); null when the row has no segments (the caller draws its plain line). */
    fun annotated(row: ActivityRow, suffix: String? = null): AnnotatedString? = parts(row)?.let { parts ->
        buildAnnotatedString {
            parts.forEach { part -> withStyle(spanOf(part.tone)) { append(part.text) } }
            if (!suffix.isNullOrBlank()) append(" · $suffix")
        }
    }

    private fun spanOf(tone: Tone): SpanStyle = when (tone) {
        Tone.Plain -> SpanStyle()
        Tone.Strong -> SpanStyle(color = FsColor.strong, fontWeight = FontWeight.SemiBold)
        is Tone.ClassName -> SpanStyle(color = classColor(tone.slug))
        is Tone.Quality -> SpanStyle(color = ItemQuality(tone.value).text)
    }
}

/**
 * One Activity row as Activity and Home's New for you draw it: the crest or the game icon, the
 * title, the styled body (else [line]), the time, the unread dot, the figure at the right in
 * tabular figures.
 */
@Composable
fun ActivityRowView(
    row: ActivityRow,
    line: String,
    build: String?,
    last: Boolean,
    onClick: () -> Unit,
    meta: String? = null,
    lineSuffix: String? = null,
) {
    FlowRow(
        title = row.title,
        sub = line,
        styledSub = ActivityBody.annotated(row, lineSuffix),
        lead = activityLead(row, build),
        unread = !row.read,
        last = last,
        meta = meta,
        value = ActivityBody.figure(row),
        onClick = onClick,
    )
}
