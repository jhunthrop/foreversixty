import DesignSystem
import SwiftUI

/// You's Plan row (spec 5.8.1). Free in a link region: `See plans` and the whole row opens the
/// browser. Free elsewhere: `Plan · Free`, not a button. Premium: the plan line, pushing F8.4.
struct PlanRow: View {
    let plan: PlanSummary
    let mayLink: Bool
    @Environment(\.openURL) private var openURL

    var body: some View {
        switch plan.rowAction(mayLink: mayLink) {
        case .seePlans:
            Button { openURL(AccountLinks.premium) } label: { content }
                .accessibilityLabel(plan.rowLabel(mayLink: mayLink))
                .accessibilityAddTraits(.isLink)
        case .openPlan:
            NavigationLink(value: YouRoute.plan) { content }
                .accessibilityLabel(plan.rowLabel(mayLink: mayLink))
        case .none:
            content
                .accessibilityLabel(plan.rowLabel(mayLink: mayLink))
        }
    }

    private var content: some View {
        PlanRowContent(plan: plan, seePlans: plan.rowAction(mayLink: mayLink) == .seePlans)
    }
}

/// The row's drawing: one line at the default sizes, stacked at accessibility sizes so the plan
/// line, the tag and `See plans` wrap and never truncate.
private struct PlanRowContent: View {
    let plan: PlanSummary
    let seePlans: Bool
    @Environment(\.dynamicTypeSize) private var typeSize

    var body: some View {
        HStack(alignment: typeSize.isAccessibilitySize ? .top : .center, spacing: FSSpace.space3) {
            Image(systemName: "crown")
                .font(.callout)
                .dynamicTypeSize(...DynamicTypeSize.xxxLarge) // the glyph stays inside its 32 pt tile
                .foregroundStyle(FSColor.gold)
                .frame(width: 32, height: 32)
                .background(FSColor.gold.opacity(0.12), in: .rect(cornerRadius: FSRadius.legacyPill))
            if typeSize.isAccessibilitySize { stacked } else { inline }
            if plan.isPremium {
                Image(systemName: "chevron.right").font(.footnote.weight(.semibold)).foregroundStyle(FSColor.muted)
            }
        }
        .padding(.horizontal, FSSpace.space4)
        .padding(.vertical, FSSpace.space2)
        .frame(minHeight: 56)
        .contentShape(.rect)
        .accessibilityElement(children: .ignore)
    }

    private var inline: some View {
        HStack(spacing: FSSpace.space2) {
            VStack(alignment: .leading, spacing: 2) {
                Text(Copy.youPlan).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                if plan.isPremium { Text(plan.row).font(ScaledType.meta).foregroundStyle(FSColor.muted) }
            }
            Spacer(minLength: FSSpace.space2)
            if plan.founder { TagView(text: Copy.planFounder, color: FSColor.gold) }
            if !plan.isPremium { Text(Copy.youPlanFree).font(ScaledType.meta).foregroundStyle(FSColor.muted) }
            if seePlans { seePlansLabel }
        }
    }

    private var stacked: some View {
        VStack(alignment: .leading, spacing: FSSpace.space1) {
            Text(Copy.youPlan).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
            Text(plan.row).font(ScaledType.meta).foregroundStyle(FSColor.muted)
            if plan.founder { TagView(text: Copy.planFounder, color: FSColor.gold) }
            if seePlans { seePlansLabel }
        }
        .fixedSize(horizontal: false, vertical: true)
        .frame(maxWidth: .infinity, alignment: .leading)
    }

    private var seePlansLabel: some View {
        HStack(spacing: 4) {
            Text(Copy.youSeePlans).font(ScaledType.bodyStrong).foregroundStyle(FSColor.gold)
            Image(systemName: "arrow.up.right.square").font(.footnote.weight(.semibold)).foregroundStyle(FSColor.gold)
        }
    }
}

/// F8.4 Plan: the card with the site's plan line, what it includes, and billing where allowed.
struct PlanView: View {
    let plan: PlanSummary
    let mayLink: Bool
    @Environment(\.openURL) private var openURL

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 0) {
                BandHead(title: Copy.youPlan)
                card.padding(FSSpace.space4)
            }
        }
        .screenGround()
        .pushedBandBar(Copy.youPlan)
    }

    private var card: some View {
        VStack(alignment: .leading, spacing: FSSpace.space3) {
            // The Founder tag beside the heading, or under it when the heading needs the room.
            ViewThatFits(in: .horizontal) {
                HStack(alignment: .firstTextBaseline) {
                    heading
                    Spacer(minLength: FSSpace.space2)
                    founderTag
                }
                VStack(alignment: .leading, spacing: FSSpace.space2) {
                    heading
                    founderTag
                }
            }
            VStack(alignment: .leading, spacing: FSSpace.space1) {
                if let price = plan.priceLine {
                    priceFigure(price)
                    Text(price.renews).font(ScaledType.body).foregroundStyle(FSColor.text)
                } else {
                    Text(plan.line).font(ScaledType.body).foregroundStyle(FSColor.text)
                }
                if let line2 = plan.line2 { muted(line2) }
            }
            if !plan.includes.isEmpty { includes }
            billing
        }
        .padding(FSSpace.space4)
        .frame(maxWidth: .infinity, alignment: .leading)
        .primaryCard()
    }

    /// The price as a Figure (spec 3.4): the amount in gold tabular figures, its period beside it.
    private func priceFigure(_ price: PlanSummary.PriceLine) -> some View {
        HStack(alignment: .firstTextBaseline, spacing: FSSpace.space2) {
            Text(price.amount).type(FSType.figure).foregroundStyle(FSColor.gold)
            Text(price.period).font(ScaledType.bodyStrong).foregroundStyle(FSColor.muted)
        }
        .accessibilityElement(children: .combine)
    }

    private var heading: some View {
        Text(Copy.planPremium).font(ScaledType.h2).foregroundStyle(FSColor.strong)
            .accessibilityAddTraits(.isHeader)
    }

    @ViewBuilder
    private var founderTag: some View {
        if plan.founder { TagView(text: Copy.planFounder, color: FSColor.gold) }
    }

    private var includes: some View {
        VStack(alignment: .leading, spacing: FSSpace.space2) {
            Divider().overlay(FSColor.lineSoft)
            Label12(Copy.planIncludes)
            ForEach(plan.includes, id: \.self) { line in
                HStack(alignment: .firstTextBaseline, spacing: FSSpace.space2) {
                    Image(systemName: "checkmark").font(.footnote.weight(.bold)).foregroundStyle(FSColor.kill)
                        .accessibilityHidden(true)
                    Text(line).font(ScaledType.body).foregroundStyle(FSColor.text)
                        .fixedSize(horizontal: false, vertical: true)
                }
            }
        }
    }

    @ViewBuilder
    private var billing: some View {
        switch plan.billingLine(mayLink: mayLink) {
        case .manageButton:
            // Spec 5.8.1: the quiet outline, the least loud control on the card.
            Button { openURL(AccountLinks.billing) } label: {
                HStack(spacing: FSSpace.space2) {
                    Text(Copy.planManage)
                    Image(systemName: "arrow.up.right.square")
                }
                .textCase(nil) // Spec 5.8.1: sentence case, never upper case.
            }
            .buttonStyle(BrandButtonStyle(.quiet, height: 44))
            .accessibilityElement(children: .ignore)
            .accessibilityLabel(Copy.planA11yManage)
            .accessibilityAddTraits(.isLink)
        case .managedOnWeb:
            muted(Copy.planManagedOnWeb)
        case .nothing:
            EmptyView()
        }
    }

    private func muted(_ text: String) -> some View {
        Text(text).font(ScaledType.meta).foregroundStyle(FSColor.muted).fixedSize(horizontal: false, vertical: true)
    }
}

/// The storefront the app reads: StoreKit, or in debug builds the `-FSStorefront <code>` launch
/// argument (`USA`, `GBR`, or `none` for an unreadable storefront), so captures show each region.
enum StorefrontLaunch {
    static var reader: any StorefrontReading {
        #if DEBUG
            if let code = UserDefaults.standard.string(forKey: "FSStorefront") {
                return FixedStorefront(code: code == "none" ? nil : code)
            }
        #endif
        return StoreKitStorefront()
    }
}
