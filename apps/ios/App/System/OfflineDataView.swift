import DesignSystem
import SwiftUI

/// You › Offline data (spec §5.7 F15.4, board F15.4), pure over its inputs so the snapshot tests
/// draw every state: the saved data by kind, the two switches, Update now and Remove.
struct OfflineDataView: View {
    let report: OfflineReport
    let update: OfflineUpdateState
    let online: Bool
    let checkedAt: Date?
    let words: OfflineWords
    let now: Date
    @Binding var keepClassesReady: Bool
    @Binding var wifiOnly: Bool
    let updateNow: () -> Void
    let remove: () -> Void
    @State private var confirmingRemove = false

    /// A write that needs a connection reads as disabled offline (spec §5.7 Writes).
    private static let disabledOpacity = 0.4

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: FSSpacing.space4) {
                SimsTitle(text: Copy.offlineDataTitle).padding(.horizontal, -16)
                SectionLabel(text: Copy.offlineDataSection(size: OfflineWords.size(report.totalBytes)))
                if report.isEmpty { empty } else { rows }
                SectionLabel(text: Copy.offlineDataUpdates).padding(.top, FSSpacing.space2)
                switches
                actions
                Text(words.footer(checkedAt: checkedAt, now: now))
                    .font(ScaledType.meta).foregroundStyle(FSColor.muted)
                    .fixedSize(horizontal: false, vertical: true)
            }
            .padding(FSSpacing.space4)
        }
        .background(FSColor.bg)
        .alert(Copy.offlineDataRemoveConfirm, isPresented: $confirmingRemove) {
            Button(Copy.offlineDataRemoveAction, role: .destructive, action: remove)
            Button(Copy.commonCancel, role: .cancel) {}
        }
    }

    private var rows: some View {
        VStack(spacing: 0) {
            ForEach(Array(report.rows.enumerated()), id: \.element.id) { index, row in
                if index > 0 { Divider().overlay(FSColor.lineSoft) }
                OfflineDataRow(row: row, meta: words.meta(row, now: now))
            }
        }
        .quietCard()
    }

    private var empty: some View {
        VStack(alignment: .leading, spacing: FSSpacing.space1) {
            Text(Copy.offlineDataEmpty).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
            Text(Copy.offlineDataEmptySub).font(ScaledType.meta).foregroundStyle(FSColor.muted)
        }
        .fixedSize(horizontal: false, vertical: true)
        .frame(maxWidth: .infinity, alignment: .leading)
        .padding(FSSpacing.space4)
        .quietCard()
    }

    private var switches: some View {
        VStack(spacing: 0) {
            OfflineToggle(title: Copy.offlineDataKeep, note: Copy.offlineDataKeepNote, isOn: $keepClassesReady)
            Divider().overlay(FSColor.lineSoft)
            OfflineToggle(title: Copy.offlineDataWifi, note: Copy.offlineDataWifiNote, isOn: $wifiOnly)
        }
        .quietCard()
    }

    private var actions: some View {
        VStack(alignment: .leading, spacing: FSSpacing.space3) {
            Button(action: updateNow) {
                Label(update == .updating ? Copy.offlineDataUpdating : Copy.offlineDataUpdateNow, systemImage: "arrow.clockwise")
            }
            .buttonStyle(BrandButtonStyle(.primary, height: 48))
            .disabled(!online || update == .updating)
            .opacity(online ? 1 : Self.disabledOpacity)
            if let status { Text(status).font(ScaledType.meta).foregroundStyle(FSColor.muted) }
            Button { confirmingRemove = true } label: {
                Text(Copy.offlineDataRemove).foregroundStyle(FSColor.wipe)
            }
            .buttonStyle(BrandButtonStyle(.secondary, height: 48))
            .disabled(report.isEmpty)
        }
    }

    /// Under Update now: why it is disabled, or what it found.
    private var status: String? {
        if !online { return Copy.offlineNeedsConnection(thing: Copy.offlineDataUpdateNow) }
        switch update {
        case .upToDate: return Copy.offlineDataUpToDate
        case .failed: return Copy.offlineUpdateFailed
        default: return nil
        }
    }
}

/// One saved kind: a 40 pt tile beside the name and its line at every size (board F15.4).
private struct OfflineDataRow: View {
    let row: OfflineRow
    let meta: String

    private static let tile: CGFloat = 40

    var body: some View {
        HStack(alignment: .top, spacing: FSSpacing.space3) {
            Image(systemName: row.category.systemImage)
                .font(.system(size: 16))
                .foregroundStyle(FSColor.gold)
                .frame(width: Self.tile, height: Self.tile)
                .background(FSColor.gold.opacity(0.12), in: .rect(cornerRadius: FSRadius.pill))
                .accessibilityHidden(true)
            VStack(alignment: .leading, spacing: 2) {
                Text(row.category.title).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                MetaParts(meta: meta)
            }
            .fixedSize(horizontal: false, vertical: true)
            Spacer(minLength: 0)
        }
        .padding(.horizontal, FSSpacing.space4)
        .padding(.vertical, FSSpacing.space3)
        .accessibilityElement(children: .combine)
    }
}

/// A row line that wraps only between its ` · ` parts, each keeping its separator at the end, so
/// a wrapped line never starts with a bare dot.
private struct MetaParts: View {
    let meta: String
    @ScaledMetric(relativeTo: .footnote) private var gap: CGFloat = 3

    var body: some View {
        let parts = meta.components(separatedBy: " · ")
        WrapLayout(spacing: gap) {
            ForEach(Array(parts.enumerated()), id: \.offset) { index, part in
                Text(index < parts.count - 1 ? part + " ·" : part).font(ScaledType.meta).foregroundStyle(FSColor.muted)
            }
        }
        .accessibilityElement(children: .ignore)
        .accessibilityLabel(meta)
    }
}

/// A switch with its note under the title (board F15.4).
private struct OfflineToggle: View {
    let title: String
    let note: String
    @Binding var isOn: Bool

    var body: some View {
        Toggle(isOn: $isOn) {
            VStack(alignment: .leading, spacing: 2) {
                Text(title).font(ScaledType.bodyStrong).foregroundStyle(FSColor.text)
                Text(note).font(ScaledType.meta).foregroundStyle(FSColor.muted).fixedSize(horizontal: false, vertical: true)
            }
        }
        .tint(FSColor.gold)
        .padding(.horizontal, FSSpacing.space4)
        .padding(.vertical, FSSpacing.space3)
        .frame(minHeight: 56)
    }
}

/// The screen bound to its model.
struct OfflineDataScreen: View {
    @State var model: OfflineDataModel
    let now: () -> Date

    var body: some View {
        OfflineDataView(
            report: model.report, update: model.update, online: model.isOnline, checkedAt: model.checkedAt,
            words: OfflineWords(calendar: .current), now: now(),
            keepClassesReady: $model.keepClassesReady, wifiOnly: $model.wifiOnly,
            updateNow: { Task { await model.updateNow() } },
            remove: { Task { await model.removeAll() } }
        )
        .navigationTitle(Copy.offlineDataTitle)
        .navigationBarTitleDisplayMode(.inline)
        .toolbar(removing: .title)
        .task { await model.load() }
    }
}
