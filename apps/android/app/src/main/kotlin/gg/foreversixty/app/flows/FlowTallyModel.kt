package gg.foreversixty.app.flows

import android.content.res.Resources
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.text.AnnotatedString
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.withStyle
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor

/** How a role stands against the raid's target: met, close, short; [None] without a target. */
enum class TallyTone(val color: Color) {
    None(FsColor.strong),
    Met(FsColor.kill),
    Close(FsColor.gold),
    Short(FsColor.wipe),
}

/**
 * One role's tally, the one model both raid surfaces read (brand S6): `Healers 5 of 8 · 3 short` in its
 * tone with a target, `Healers 6` without one; compact on Home's line as `5/8`.
 */
data class TallyItem(val label: String, val have: Int, val want: Int?) {
    val short: Int get() = want?.let { (it - have).coerceAtLeast(0) } ?: 0

    val tone: TallyTone
        get() {
            val target = want ?: return TallyTone.None
            return when {
                have >= target -> TallyTone.Met
                have >= target * CLOSE_SHARE -> TallyTone.Close
                else -> TallyTone.Short
            }
        }

    /** The count part: `5 of 8 · 3 short`, `8 of 8`, or `6` without a target. */
    fun count(res: Resources): String {
        val target = want ?: return have.toString()
        val of = res.getString(R.string.raid_tab_ofWant, have, target)
        return if (short > 0) res.getString(R.string.raid_tab_tallyShort, of, short) else of
    }

    /** `2/2`, `5/8 · 3 short` (row 2: red alone does not say how many), or the count alone without a target. */
    fun compact(res: Resources): String {
        val target = want ?: return have.toString()
        val of = res.getString(R.string.raid_tab_tallyCompact, have, target)
        return if (short > 0) res.getString(R.string.raid_tab_tallyShort, of, short) else of
    }

    private companion object {
        /** Three quarters of the target reads close (gold), under it short (wipe). */
        const val CLOSE_SHARE = 0.75f
    }
}

/** `Tanks 2/2 · Healers 5/8 · Damage 11/30` in one line, each count in its tone (Home's line, S6). */
@Composable
fun FlowTallyLine(items: List<TallyItem>, modifier: Modifier = Modifier) {
    if (items.isEmpty()) return
    val res = LocalResources.current
    val text = buildAnnotatedString {
        items.forEachIndexed { index, item ->
            if (index > 0) append(SEPARATOR)
            append(item.label)
            append(" ")
            withStyle(SpanStyle(color = item.tone.color)) { append(item.compact(res)) }
        }
    }
    Text(text, style = SUB, color = FsColor.muted, modifier = modifier)
}

/** `Healers` then the count in its tone: a Who is coming row's title. */
@Composable
fun tallyTitle(item: TallyItem): AnnotatedString {
    val count = item.count(LocalResources.current)
    return buildAnnotatedString {
        append(item.label)
        append(" ")
        withStyle(SpanStyle(color = item.tone.color)) { append(count) }
    }
}

private const val SEPARATOR = " · "
