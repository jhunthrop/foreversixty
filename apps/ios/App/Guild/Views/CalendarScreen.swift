import DesignSystem
import SwiftUI
import UIKit
import UniformTypeIdentifiers

/// Everything the Calendar screen draws (spec F16.4, F16.5), so snapshot tests draw every state.
struct CalendarScreenState: Equatable {
    var sync: CalendarSyncPhase = .off
    var settings = CalendarSettings()
    var syncFailed = false
    var link: CalendarLinkPhase = .loading
    var linkOffline = false
    var linkMessage: String?
    var copied = false
    var resetting = false
}

struct CalendarActions {
    var setEnabled: (Bool) -> Void = { _ in }
    var setWhich: (CalendarWhich) -> Void = { _ in }
    var setAlert: (CalendarAlert) -> Void = { _ in }
    var openSettings: () -> Void = {}
    var copy: () -> Void = {}
    var askReset: () -> Void = {}
    var addToCalendar: () -> Void = {}
    var retryLink: () -> Void = {}
}

struct CalendarContentView: View {
    let state: CalendarScreenState
    let actions: CalendarActions

    var body: some View {
        GuildScreenScroll(offlineStrip: nil) {
            SectionLabel(text: Copy.calendarSection)
            if state.sync == .accessOff { accessOffCard }
            syncCard
            SectionLabel(text: Copy.calendarShare).padding(.top, 6)
            linkCard
            GuildFooter(text: Copy.calendarFooter)
        }
    }

    // MARK: Part 1, the calendar the app keeps

    private var accessOffCard: some View {
        VStack(alignment: .leading, spacing: 10) {
            Text(Copy.calendarOff).font(ScaledType.bodyStrong).foregroundStyle(FSColor.wipe)
            Text(Copy.calendarOffBody).font(ScaledType.meta).foregroundStyle(FSColor.muted)
            SecondaryButton(title: Copy.calendarOpenSettings, action: actions.openSettings)
        }
        .raidCardBox(16).quietCard()
    }

    private var syncCard: some View {
        let accessOff = state.sync == .accessOff
        let isOn = state.sync == .on
        return VStack(spacing: 0) {
            VStack(alignment: .leading, spacing: 6) {
                Toggle(isOn: Binding(get: { isOn }, set: { actions.setEnabled($0) })) {
                    VStack(alignment: .leading, spacing: 2) {
                        Text(Copy.calendarSwitch).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                        Text(Copy.calendarSwitchNote).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                            .fixedSize(horizontal: false, vertical: true)
                    }
                }
                .tint(FSColor.gold)
                .disabled(accessOff)
                .opacity(accessOff ? 0.5 : 1)
                if state.sync == .removed {
                    Text(Copy.calendarRemoved).font(ScaledType.meta).foregroundStyle(FSColor.gold)
                }
                if state.syncFailed {
                    Text(Copy.calendarSyncFailed).font(ScaledType.meta).foregroundStyle(FSColor.wipe)
                }
            }
            .padding(.horizontal, 16)
            .padding(.vertical, 12)
            Divider().overlay(FSColor.lineSoft)
            menuRow(Copy.calendarWhich, value: state.settings.which.label) {
                Picker(Copy.calendarWhich, selection: Binding(get: { state.settings.which }, set: { actions.setWhich($0) })) {
                    ForEach(CalendarWhich.allCases) { Text($0.label).tag($0) }
                }
            }
            Divider().overlay(FSColor.lineSoft)
            menuRow(Copy.calendarAlert, value: state.settings.alert.label) {
                Picker(Copy.calendarAlert, selection: Binding(get: { state.settings.alert }, set: { actions.setAlert($0) })) {
                    ForEach(CalendarAlert.allCases) { Text($0.label).tag($0) }
                }
            }
        }
        .quietCard()
    }

    private func menuRow(_ title: String, value: String, @ViewBuilder picker: () -> some View) -> some View {
        Menu { picker() } label: {
            ViewThatFits(in: .horizontal) {
                HStack(spacing: 8) {
                    Text(title).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                    Spacer(minLength: 8)
                    valueLabel(value)
                }
                VStack(alignment: .leading, spacing: 4) {
                    Text(title).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                    valueLabel(value)
                }
            }
            .padding(.horizontal, 16)
            .frame(minHeight: 52)
            .contentShape(.rect)
        }
        .accessibilityLabel(title)
        .accessibilityValue(value)
    }

    private func valueLabel(_ value: String) -> some View {
        HStack(spacing: 4) {
            Text(value).font(ScaledType.meta).foregroundStyle(FSColor.muted).multilineTextAlignment(.trailing)
            Image(systemName: "chevron.up.chevron.down").font(.system(size: 10, weight: .bold)).foregroundStyle(FSColor.muted)
                .accessibilityHidden(true)
        }
    }

    // MARK: Part 2, the calendar link

    private var linkCard: some View {
        VStack(alignment: .leading, spacing: 14) {
            HStack(alignment: .top, spacing: 12) {
                Image(systemName: "link")
                    .font(.callout)
                    .foregroundStyle(FSColor.gold)
                    .frame(width: 32, height: 32)
                    .background(FSColor.gold.opacity(0.12), in: .rect(cornerRadius: FSRadius.pill))
                    .accessibilityHidden(true)
                VStack(alignment: .leading, spacing: 2) {
                    Text(Copy.calendarLink).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                    Text(Copy.calendarLinkNote).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                        .fixedSize(horizontal: false, vertical: true)
                }
            }
            .accessibilityElement(children: .combine)
            linkBody
        }
        .raidCardBox(16).quietCard()
    }

    @ViewBuilder
    private var linkBody: some View {
        switch state.link {
        case .loading:
            ProgressView().tint(FSColor.gold).frame(maxWidth: .infinity, minHeight: 48)
        case .failed:
            Text(state.linkMessage ?? Copy.calendarLinkError).font(ScaledType.meta).foregroundStyle(FSColor.wipe)
            SecondaryButton(title: Copy.commonTryAgain, action: actions.retryLink)
        case .elsewhere:
            Text(Copy.calendarLinkElsewhere).font(ScaledType.meta).foregroundStyle(FSColor.text)
                .fixedSize(horizontal: false, vertical: true)
            resetButton.frame(maxWidth: .infinity)
            messageLine
        case .ready:
            ViewThatFits(in: .horizontal) {
                HStack(spacing: 10) { copyButton; resetButton }
                VStack(spacing: 10) { copyButton; resetButton }
            }
            linkButton(Copy.calendarAddToCalendar, systemImage: "calendar.badge.plus", action: actions.addToCalendar)
            if state.copied {
                Text(Copy.calendarCopied).font(ScaledType.meta).foregroundStyle(FSColor.kill)
            }
            messageLine
        }
    }

    @ViewBuilder
    private var messageLine: some View {
        if let message = state.linkMessage {
            Text(message).font(ScaledType.meta).foregroundStyle(FSColor.wipe)
        }
    }

    private var copyButton: some View {
        linkButton(Copy.calendarCopy, systemImage: "doc.on.doc", action: actions.copy)
    }

    private var resetButton: some View {
        Button(action: actions.askReset) {
            HStack(spacing: 8) {
                if state.resetting { ProgressView().controlSize(.small) }
                Text(Copy.calendarReset)
            }
        }
        .buttonStyle(BrandButtonStyle(.secondary, height: 48))
        .disabled(state.linkOffline || state.resetting)
        .opacity(state.linkOffline ? 0.5 : 1)
    }

    private func linkButton(_ title: String, systemImage: String, action: @escaping () -> Void) -> some View {
        Button(action: action) {
            HStack(spacing: 8) {
                Image(systemName: systemImage).accessibilityHidden(true)
                Text(title)
            }
        }
        .buttonStyle(BrandButtonStyle(.secondary, height: 48))
        .disabled(state.linkOffline)
        .opacity(state.linkOffline ? 0.5 : 1)
    }
}

/// You › Calendar: binds the sync controller and the link model to the content.
struct CalendarScreen: View {
    let sync: CalendarSyncController
    let link: CalendarLinkModel
    @State private var copied = false
    @State private var askingReset = false
    @Environment(\.openURL) private var openURL
    @Environment(\.scenePhase) private var scenePhase

    var body: some View {
        CalendarContentView(state: state, actions: actions)
            .navigationTitle(Copy.calendarTitle)
            .navigationBarTitleDisplayMode(.large)
            .alert(Copy.calendarResetAsk, isPresented: $askingReset) {
                Button(Copy.calendarResetConfirm, role: .destructive) {
                    copied = false
                    Task { await link.reset() }
                }
                Button(Copy.commonCancel, role: .cancel) {}
            }
            .task { await link.load() }
            .onAppear { sync.refresh() }
            .onChange(of: scenePhase) { _, phase in if phase == .active { sync.refresh() } }
    }

    private var state: CalendarScreenState {
        CalendarScreenState(
            sync: sync.phase, settings: sync.record.settings, syncFailed: sync.syncFailed,
            link: link.phase, linkOffline: link.offline, linkMessage: link.message, copied: copied, resetting: link.resetting
        )
    }

    private var actions: CalendarActions {
        CalendarActions(
            setEnabled: { on in Task { await sync.setEnabled(on) } },
            setWhich: { which in Task { await sync.setWhich(which) } },
            setAlert: { alert in Task { await sync.setAlert(alert) } },
            openSettings: { URL(string: UIApplication.openSettingsURLString).map { openURL($0) } },
            copy: copy,
            askReset: { askingReset = true },
            addToCalendar: { link.url.flatMap(CalendarLinkModel.webcal).map { openURL($0) } },
            retryLink: { Task { await link.load() } }
        )
    }

    /// The link goes to the clipboard marked local-only and expiring, never logged.
    private func copy() {
        guard let url = link.url else { return }
        UIPasteboard.general.setItems(
            [[UTType.url.identifier: url, UTType.plainText.identifier: url]],
            options: [.localOnly: true, .expirationDate: Date().addingTimeInterval(10 * 60)]
        )
        copied = true
    }
}
