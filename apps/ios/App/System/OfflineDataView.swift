import DesignSystem
import SwiftUI

/// Offline data (mocks ios-offline-data, -empty, -strip), pure over its inputs so the snapshot
/// tests draw every state: On this phone by kind, Updates with the two switches, Update now,
/// Remove downloaded data and the footer; offline, the Notice under the title.
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
            VStack(alignment: .leading, spacing: 0) {
                PageHeader(title: Copy.offlineDataTitle)
                if !online {
                    Notice(Copy.offlineStrip(time: OfflineStamp.format(checkedAt ?? now, now: now)))
                }
                FSSection(report.isEmpty ? Copy.offlineDataSectionEmpty : Copy.offlineDataSection(size: OfflineWords.size(report.totalBytes))) {
                    if report.isEmpty { empty } else { rows }
                }
                .padding(.top, DevicesView.sectionGap)
                FSSection(Copy.offlineDataUpdates) { switches }
                    .padding(.top, DevicesView.sectionGap)
                actions.padding(.top, DevicesView.stackGap)
                SmallText(words.footer(checkedAt: checkedAt, now: now)).padding(.top, FSSpace.space4)
            }
            .padding(.top, FSSpace.space2)
            .padding(.horizontal, FSSpace.gutterPhone)
            .padding(.bottom, FSSpace.space6)
            .parityContent()
        }
        .background(FSColor.bg)
        .alert(Copy.offlineDataRemoveConfirm, isPresented: $confirmingRemove) {
            Button(Copy.offlineDataRemoveAction, role: .destructive, action: remove)
            Button(Copy.commonCancel, role: .cancel) {}
        }
    }

    private var rows: some View {
        VStack(spacing: 0) {
            ForEach(report.rows) { row in
                FSRow(title: row.category.title, sub: words.meta(row, now: now), inset: 0) { EmptyView() }
            }
        }
    }

    private var empty: some View {
        SmallText("\(Copy.offlineDataEmpty) \(Copy.offlineDataEmptySub)").padding(.top, FSSpace.space3)
    }

    private var switches: some View {
        VStack(spacing: 0) {
            toggleRow(Copy.offlineDataKeep, note: Copy.offlineDataKeepNote, isOn: $keepClassesReady)
            toggleRow(Copy.offlineDataWifi, note: Copy.offlineDataWifiNote, isOn: $wifiOnly)
        }
    }

    private func toggleRow(_ title: String, note: String, isOn: Binding<Bool>) -> some View {
        FSRow(title: title, sub: note, inset: 0, trailPlacement: .below, lead: { EmptyView() }) {
            LabelledSwitch(title, isOn: isOn)
        }
    }

    private var actions: some View {
        VStack(alignment: .leading, spacing: FSSpace.space3) {
            FSButton(update == .updating ? Copy.offlineDataUpdating : Copy.offlineDataUpdateNow, action: updateNow)
                .disabled(!online || update == .updating)
                .opacity(online ? 1 : Self.disabledOpacity)
            if let status { SmallText(status) }
            FSButton(Copy.offlineDataRemove, kind: .danger) { confirmingRemove = true }
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
