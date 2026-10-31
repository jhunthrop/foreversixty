package gg.foreversixty.app.premium

import androidx.compose.foundation.BorderStroke
import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.filled.KeyboardArrowRight
import androidx.compose.material.icons.filled.Check
import androidx.compose.material.icons.filled.Star
import androidx.compose.material3.ButtonDefaults
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Icon
import androidx.compose.material3.OutlinedButton
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.role
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.account.AccountPage
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsSpacing
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.GameColors
import gg.foreversixty.app.designsystem.MicroLabel
import gg.foreversixty.app.designsystem.Tag
import gg.foreversixty.app.designsystem.primaryCard
import java.util.Locale

/** From this font scale up the Plan row stacks, so nothing truncates (spec 5.8.1). */
private const val STACK_FONT_SCALE = 1.5f

/**
 * You's Plan row (spec 5.8.1). Free in a link region: `See plans` and the whole row opens the
 * browser. Free elsewhere: `Plan · Free`, not a button. Premium: the plan line, opening F8.4.
 */
@Composable
fun PlanRow(plan: PlanSummary, mayLink: Boolean, onSeePlans: () -> Unit, onPlan: () -> Unit) {
    val action = plan.rowAction(mayLink)
    val label = PlanCopy.rowLabel(plan, mayLink)
    val click: (() -> Unit)? = when (action) {
        PlanRowAction.SeePlans -> onSeePlans
        PlanRowAction.OpenPlan -> onPlan
        PlanRowAction.None -> null
    }
    val stacked = LocalDensity.current.fontScale >= STACK_FONT_SCALE
    Row(
        Modifier
            .fillMaxWidth()
            .heightIn(min = 56.dp)
            .let { if (click != null) it.clickable(onClick = click) else it }
            .clearAndSetSemantics {
                contentDescription = label
                if (click != null) role = Role.Button
            }
            .padding(horizontal = FsSpacing.space4, vertical = FsSpacing.space2),
        verticalAlignment = if (stacked) Alignment.Top else Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(FsSpacing.space3),
    ) {
        Box(
            Modifier.size(32.dp).background(FsColor.gold.copy(alpha = 0.12f), RoundedCornerShape(FsRadius.pill)),
            contentAlignment = Alignment.Center,
        ) { Icon(Icons.Filled.Star, contentDescription = null, tint = FsColor.gold, modifier = Modifier.size(20.dp)) }
        if (stacked) Stacked(plan, action == PlanRowAction.SeePlans, Modifier.weight(1f)) else Inline(plan, action == PlanRowAction.SeePlans, Modifier.weight(1f))
        if (plan.isPremium) Icon(Icons.AutoMirrored.Filled.KeyboardArrowRight, contentDescription = null, tint = FsColor.muted)
    }
}

@Composable
private fun Inline(plan: PlanSummary, seePlans: Boolean, modifier: Modifier) {
    Row(modifier, verticalAlignment = Alignment.CenterVertically, horizontalArrangement = Arrangement.spacedBy(FsSpacing.space2)) {
        Column(Modifier.weight(1f)) {
            Text(stringResource(R.string.you_plan), style = FsType.bodyStrong, color = FsColor.strong)
            if (plan.isPremium) Text(PlanCopy.line(plan.row), style = FsType.meta, color = FsColor.muted)
        }
        if (plan.founder) Tag(stringResource(R.string.plan_founder), FsColor.gold)
        if (!plan.isPremium) Text(stringResource(R.string.you_plan_free), style = FsType.meta, color = FsColor.muted)
        if (seePlans) SeePlans()
    }
}

@Composable
private fun Stacked(plan: PlanSummary, seePlans: Boolean, modifier: Modifier) {
    Column(modifier, verticalArrangement = Arrangement.spacedBy(FsSpacing.space1)) {
        Text(stringResource(R.string.you_plan), style = FsType.bodyStrong, color = FsColor.strong)
        Text(PlanCopy.line(plan.row), style = FsType.meta, color = FsColor.muted)
        if (plan.founder) Tag(stringResource(R.string.plan_founder), FsColor.gold)
        if (seePlans) SeePlans()
    }
}

@Composable
private fun SeePlans() {
    Row(verticalAlignment = Alignment.CenterVertically, horizontalArrangement = Arrangement.spacedBy(4.dp)) {
        Text(stringResource(R.string.you_seePlans), style = FsType.bodyStrong, color = FsColor.gold)
        Icon(painterResource(R.drawable.ic_open_external), contentDescription = null, tint = FsColor.gold, modifier = Modifier.size(18.dp))
    }
}

/** F8.4 Plan: the card with the site's plan line, what it includes, and billing where allowed. */
@Composable
fun PlanScreen(plan: PlanSummary, mayLink: Boolean, onManageBilling: () -> Unit) {
    AccountPage(stringResource(R.string.you_plan)) {
        Column(
            Modifier.fillMaxWidth().primaryCard().padding(FsSpacing.space4),
            verticalArrangement = Arrangement.spacedBy(FsSpacing.space3),
        ) {
            Row(verticalAlignment = Alignment.CenterVertically) {
                Text(
                    stringResource(R.string.plan_premium).uppercase(Locale.ROOT),
                    style = FsType.h2,
                    color = FsColor.strong,
                    modifier = Modifier.weight(1f).semantics { heading() },
                )
                if (plan.founder) Tag(stringResource(R.string.plan_founder), FsColor.gold)
            }
            Column(verticalArrangement = Arrangement.spacedBy(FsSpacing.space1)) {
                Text(PlanCopy.line(plan.line), style = FsType.body, color = FsColor.text)
                plan.line2?.let { Text(PlanCopy.given(it), style = FsType.meta, color = FsColor.muted) }
                if (plan.founder) Text(stringResource(R.string.plan_founderNote), style = FsType.meta, color = FsColor.muted)
            }
            if (plan.includes.isNotEmpty()) Includes(plan.includes)
            when (plan.billingLine(mayLink)) {
                PlanBillingLine.ManageButton -> ManageBilling(onManageBilling)
                PlanBillingLine.ManagedOnWeb -> Text(stringResource(R.string.plan_managedOnWeb), style = FsType.meta, color = FsColor.muted)
                PlanBillingLine.Nothing -> Unit
            }
        }
        Text(stringResource(R.string.plan_freeForAll), style = FsType.meta, color = FsColor.muted)
    }
}

@Composable
private fun Includes(features: List<PlanFeature>) {
    Column(verticalArrangement = Arrangement.spacedBy(FsSpacing.space2)) {
        HorizontalDivider(color = FsColor.lineSoft)
        MicroLabel(stringResource(R.string.plan_includes))
        features.forEach { feature ->
            Row(horizontalArrangement = Arrangement.spacedBy(FsSpacing.space2)) {
                Icon(Icons.Filled.Check, contentDescription = null, tint = FsColor.kill, modifier = Modifier.padding(top = 2.dp).size(16.dp))
                Text(PlanCopy.feature(feature), style = FsType.body, color = FsColor.text)
            }
        }
    }
}

/** Spec 5.8.1: the quiet outline in sentence case, the least loud control on the card. */
@Composable
private fun ManageBilling(onClick: () -> Unit) {
    val label = stringResource(R.string.plan_a11y_manage)
    OutlinedButton(
        onClick = onClick,
        shape = RoundedCornerShape(FsRadius.control),
        border = BorderStroke(1.dp, FsColor.lineCard),
        colors = ButtonDefaults.outlinedButtonColors(containerColor = GameColors.controlFill, contentColor = FsColor.nav),
        modifier = Modifier.fillMaxWidth().heightIn(min = 48.dp).semantics { contentDescription = label },
    ) {
        Text(stringResource(R.string.plan_manage), style = FsType.button)
        Icon(painterResource(R.drawable.ic_open_external), contentDescription = null, modifier = Modifier.padding(start = FsSpacing.space2).size(16.dp))
    }
}
