package gg.foreversixty.app.premium

import androidx.compose.runtime.Composable
import androidx.compose.ui.res.stringResource
import gg.foreversixty.app.R
import gg.foreversixty.app.account.AccountCopy

/** [PlanSummary]'s words from the copy catalog: the site's `accountPlanCopy` lines. */
object PlanCopy {
    @Composable
    fun line(line: PlanLine): String = when (line) {
        PlanLine.Free -> stringResource(R.string.you_plan_free)
        PlanLine.Premium -> stringResource(R.string.plan_premium)
        is PlanLine.Renews -> stringResource(R.string.plan_renews, AccountCopy.longDate(line.date))
        is PlanLine.RenewsPrice -> stringResource(R.string.plan_line_renewsPrice, AccountCopy.longDate(line.date), price(line.price))
        is PlanLine.Ends -> stringResource(R.string.plan_ends, AccountCopy.longDate(line.date))
        is PlanLine.EndsThenSeat -> stringResource(R.string.plan_line_endsThenSeat, AccountCopy.longDate(line.date), line.guild)
        is PlanLine.GrantUntil -> stringResource(R.string.plan_line_grantUntil, AccountCopy.longDate(line.date))
        PlanLine.GrantNoEnd -> stringResource(R.string.plan_line_grantNoEnd)
        is PlanLine.Seat -> stringResource(R.string.plan_line_seat, line.guild)
    }

    @Composable
    fun given(given: PlanGiven): String = when (given) {
        PlanGiven.GrantUntil -> stringResource(R.string.plan_given_grantUntil)
        PlanGiven.GrantNoEnd -> stringResource(R.string.plan_given_grantNoEnd)
        is PlanGiven.Seat -> stringResource(R.string.plan_seatGivenBy, given.givenBy)
    }

    @Composable
    fun price(price: PlanPrice): String =
        stringResource(if (price.yearly) R.string.plan_price_yearly else R.string.plan_price_monthly, price.amount)

    @Composable
    fun feature(feature: PlanFeature): String = stringResource(
        when (feature) {
            PlanFeature.Coach -> R.string.plan_feature_coach
            PlanFeature.ServerSims -> R.string.plan_feature_serverSims
            PlanFeature.History -> R.string.plan_feature_history
            PlanFeature.Discord -> R.string.plan_feature_discord
            PlanFeature.Supporter -> R.string.plan_feature_supporter
        },
    )

    /** The Plan row's one TalkBack label (spec 5.8.1). */
    @Composable
    fun rowLabel(plan: PlanSummary, mayLink: Boolean): String = when {
        !plan.isPremium && mayLink -> stringResource(R.string.plan_a11y_seePlans)
        !plan.isPremium -> stringResource(R.string.plan_a11y_free)
        plan.founder -> stringResource(R.string.plan_a11y_rowFounder, line(plan.row))
        else -> stringResource(R.string.plan_a11y_row, line(plan.row))
    }
}
