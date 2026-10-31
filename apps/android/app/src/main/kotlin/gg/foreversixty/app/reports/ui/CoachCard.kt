package gg.foreversixty.app.reports.ui

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.LiveRegionMode
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.liveRegion
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.MicroLabel
import gg.foreversixty.app.designsystem.Tag
import gg.foreversixty.app.designsystem.quietCard
import gg.foreversixty.app.model.CoachReview
import gg.foreversixty.app.reports.logic.CoachPhase
import gg.foreversixty.app.reports.logic.CoachSection
import gg.foreversixty.app.reports.logic.ReportsCopy
import java.util.Locale

/** F10.4: Premium only, and only with a night line. The coach's words are printed as given. */
@Composable
fun CoachCard(phase: CoachPhase, copy: ReportsCopy, retry: () -> Unit) {
    if (phase == CoachPhase.Hidden) return
    Column(verticalArrangement = Arrangement.spacedBy(8.dp)) {
        MicroLabel(stringResource(R.string.report_coach))
        CoachBody(phase, copy, retry)
    }
}

@Composable
private fun CoachBody(phase: CoachPhase, copy: ReportsCopy, retry: () -> Unit) {
    Column(Modifier.fillMaxWidth().quietCard().padding(16.dp), verticalArrangement = Arrangement.spacedBy(14.dp)) {
        when (phase) {
            CoachPhase.Hidden -> Unit
            CoachPhase.Loading -> Column(verticalArrangement = Arrangement.spacedBy(10.dp)) { SkeletonBar(200.dp, 18.dp); SkeletonBar(260.dp, 13.dp) }
            CoachPhase.Writing -> Line(stringResource(R.string.coach_writing))
            CoachPhase.Preparing -> Line(stringResource(R.string.coach_preparing))
            CoachPhase.Failed -> ErrorRegion(stringResource(R.string.coach_failed), retry, inset = false)
            is CoachPhase.Ready -> Review(phase.review, copy)
        }
    }
}

@Composable
private fun Line(text: String) {
    Text(text, style = FsType.body, color = FsColor.text, modifier = Modifier.semantics { liveRegion = LiveRegionMode.Polite })
}

@Composable
private fun Review(review: CoachReview, copy: ReportsCopy) {
    Row(verticalAlignment = Alignment.CenterVertically, horizontalArrangement = Arrangement.spacedBy(10.dp)) {
        Text(
            stringResource(R.string.coach_title, review.spec).uppercase(Locale.ROOT), style = FsType.sec, color = FsColor.strong,
            modifier = Modifier.weight(1f).semantics { heading() },
        )
        Tag(stringResource(R.string.coach_premium), FsColor.gold)
    }
    CoachSection.of(review).forEach { section -> Section(section) }
    Text(copy.coachFooter(review.reviewedAt, review.basisLabels), style = FsType.meta, color = FsColor.muted)
}

@Composable
private fun Section(section: CoachSection) {
    Column(verticalArrangement = Arrangement.spacedBy(4.dp)) {
        when (section) {
            is CoachSection.FixFirst -> {
                MicroLabel(stringResource(R.string.coach_fixFirst))
                Text(section.title, style = FsType.bodyStrong, color = FsColor.strong)
                Text(section.why, style = FsType.body, color = FsColor.text)
            }
            is CoachSection.Practice -> {
                MicroLabel(stringResource(R.string.coach_practice))
                Text(section.text, style = FsType.body, color = FsColor.text)
            }
            is CoachSection.Then -> {
                MicroLabel(stringResource(R.string.coach_then))
                section.items.forEach { (title, why) ->
                    Text(title, style = FsType.bodyStrong, color = FsColor.strong)
                    Text(why, style = FsType.body, color = FsColor.text)
                }
            }
            is CoachSection.DoingWell -> {
                MicroLabel(stringResource(R.string.coach_doingWell))
                section.lines.forEach { Text(it, style = FsType.body, color = FsColor.text) }
            }
            is CoachSection.CouldntCheck -> {
                MicroLabel(stringResource(R.string.coach_couldntCheck))
                section.lines.forEach { Text(it, style = FsType.body, color = FsColor.muted) }
            }
        }
    }
}
