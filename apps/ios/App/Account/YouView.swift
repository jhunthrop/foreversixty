import Auth
import DesignSystem
import Models
import SwiftUI

/// What the You tab's account section shows, already worded (spec F8.1's account head and
/// rows that F1 owns: Plan, Devices, Delete account, Sign out). F7 and F8 add their rows.
struct YouSummary: Equatable {
    let battletag: String
    let plan: String
    let founder: Bool
    /// `This iPhone and 2 more`, nil until the device list answers.
    let devices: String?

    init(battletag: String, plan: String, founder: Bool, devices: String?) {
        self.battletag = battletag
        self.plan = plan
        self.founder = founder
        self.devices = devices
    }

    init(me: Me, deviceCount: Int?) {
        let billing = me.entitlements.billing
        self.init(
            battletag: me.user.battletag ?? "",
            plan: billing.map(Self.planLine) ?? Copy.youPlanFree,
            founder: billing?.founder == true,
            devices: deviceCount.map { Copy.youDevicesSub(device: Copy.accountDeviceWord, n: max(0, $0 - 1)) }
        )
    }

    /// The site's `accountPlanCopy` lines: `Premium · renews 27 Oct 2027`, or `ends` once cancelled.
    static func planLine(_ billing: Billing) -> String {
        let end = billing.currentPeriodEnd.flatMap { try? Date($0, strategy: .iso8601) }
        guard let end else { return Copy.youPlanFree }
        let date = AccountCopy.longDate(end)
        return billing.cancelAtPeriodEnd ? Copy.planEnds(date: date) : Copy.planRenews(date: date)
    }
}

enum YouRoute: Hashable {
    case devices
    case deleteAccount
}

/// The You tab's account section. Pure: the snapshot tests render it with a summary.
struct YouView: View {
    let summary: YouSummary
    let onSignOut: () -> Void
    @State private var confirmingSignOut = false

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: FSSpacing.space5) {
                accountHead
                VStack(spacing: 0) {
                    AccountRow(icon: "crown", title: Copy.youPlan, trailing: summary.plan, tag: summary.founder ? Copy.planFounder : nil)
                    Divider().overlay(FSColor.lineSoft)
                    NavigationLink(value: YouRoute.devices) {
                        AccountRow(icon: "iphone", title: Copy.youDevices, subtitle: summary.devices, chevron: true)
                    }
                }
                .accountCard()
                NavigationLink(value: YouRoute.deleteAccount) {
                    AccountRow(icon: "trash", title: Copy.accountDeleteRow, chevron: true, tint: FSColor.wipe)
                }
                .accountCard()
                Button(Copy.youSignOut) { confirmingSignOut = true }
                    .buttonStyle(BrandButtonStyle(.quiet, height: 48))
            }
            .buttonStyle(.plain)
            .padding(FSSpacing.space4)
        }
        .background(FSColor.bg)
        .navigationTitle(Copy.youTitle)
        .confirmationDialog(
            Copy.youSignOutConfirm(device: Copy.accountDeviceWord),
            isPresented: $confirmingSignOut,
            titleVisibility: .visible
        ) {
            Button(Copy.accountSignOutAction, role: .destructive, action: onSignOut)
            Button(Copy.commonCancel, role: .cancel) {}
        }
    }

    private var accountHead: some View {
        HStack(spacing: FSSpacing.space3) {
            Image(systemName: "person")
                .font(.title3)
                .foregroundStyle(FSColor.gold)
                .frame(width: 44, height: 44)
                .background(FSColor.gold.opacity(0.12), in: .rect(cornerRadius: FSRadius.control))
                .accessibilityHidden(true)
            VStack(alignment: .leading, spacing: 2) {
                Text(summary.battletag).font(ScaledType.name).foregroundStyle(FSColor.strong)
                Text(Copy.youSignedInWith).font(ScaledType.meta).foregroundStyle(FSColor.muted)
            }
        }
        .accessibilityElement(children: .combine)
    }
}

/// One list row: a gold tile, the title and its sub line, a trailing value, tag or chevron.
struct AccountRow: View {
    let icon: String
    let title: String
    var subtitle: String?
    var trailing: String?
    var tag: String?
    var chevron = false
    var tint: Color = FSColor.strong

    var body: some View {
        HStack(spacing: FSSpacing.space3) {
            Image(systemName: icon)
                .font(.callout)
                .foregroundStyle(tint == FSColor.strong ? FSColor.gold : tint)
                .frame(width: 32, height: 32)
                .background(FSColor.gold.opacity(0.12), in: .rect(cornerRadius: FSRadius.pill))
                .accessibilityHidden(true)
            VStack(alignment: .leading, spacing: 2) {
                Text(title).font(ScaledType.bodyStrong).foregroundStyle(tint)
                if let subtitle {
                    Text(subtitle).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                }
            }
            Spacer(minLength: FSSpacing.space2)
            if let tag {
                Text(tag)
                    .font(ScaledType.micro)
                    .textCase(.uppercase)
                    .foregroundStyle(FSColor.gold)
                    .padding(.horizontal, FSSpacing.space2)
                    .frame(minHeight: 22)
                    .overlay(Capsule().stroke(FSColor.gold.opacity(0.4)))
            }
            if let trailing {
                Text(trailing).font(ScaledType.meta).foregroundStyle(FSColor.muted).multilineTextAlignment(.trailing)
            }
            if chevron {
                Image(systemName: "chevron.right").font(.footnote.weight(.semibold)).foregroundStyle(FSColor.muted)
                    .accessibilityHidden(true)
            }
        }
        .padding(.horizontal, FSSpacing.space4)
        .frame(minHeight: 56)
        .contentShape(.rect)
        .accessibilityElement(children: .combine)
    }
}

extension View {
    /// The quiet list card: raised, hairline border, panel radius.
    func accountCard() -> some View {
        background(FSColor.raised, in: .rect(cornerRadius: FSRadius.panel))
            .overlay(RoundedRectangle(cornerRadius: FSRadius.panel).stroke(FSColor.lineSoft))
    }
}
