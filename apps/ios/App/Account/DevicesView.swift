import Auth
import DesignSystem
import Networking
import SwiftUI

/// What F1.6 shows: the rows, or the load failure.
enum DevicesState: Equatable {
    case loading
    case loaded([DeviceRow])
    case failed
}

/// Devices (mock ios-devices): Signed in, a row per device with Revoke (this one tagged), Sign out
/// everywhere else, and the footer. Pure over its state; `DevicesScreen` loads and acts.
struct DevicesView: View {
    let state: DevicesState
    /// The row whose revoke failed: `Could not revoke that device.` shows under it.
    let failedRevoke: String?
    var othersFailed = false
    let onRevoke: (DeviceRow) -> Void
    var onRevokeOthers: () -> Void = {}
    let onRetry: () -> Void
    @State private var confirming: DeviceRow?
    @State private var confirmingOthers = false

    /// `.ap section { margin-top: 28px }`; `.stack { margin-top: 20px }`.
    static let sectionGap: CGFloat = 28
    static let stackGap: CGFloat = FSSpace.space5

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 0) {
                PageHeader(title: Copy.devicesTitle)
                FSSection(Copy.devicesSignedInTitle) { content }
                    .padding(.top, Self.sectionGap)
                if case let .loaded(rows) = state, rows.contains(where: \.canRevoke) {
                    FSButton(Copy.devicesSignOutOthers, kind: .danger) { confirmingOthers = true }
                        .padding(.top, Self.stackGap)
                    if othersFailed { SmallText(Copy.devicesSignOutOthersFailed, color: FSColor.wipe).padding(.top, FSSpace.space2) }
                }
                SmallText(Copy.devicesFooter).padding(.top, FSSpace.space4)
            }
            .padding(.top, FSSpace.space2)
            .padding(.horizontal, FSSpace.gutterPhone)
            .padding(.bottom, FSSpace.space6)
            .parityContent()
        }
        .background(FSColor.bg)
        .navigationTitle(Copy.devicesTitle)
        .navigationBarTitleDisplayMode(.inline)
        .alert(
            confirming.map { Copy.devicesRevokeTitle(device: $0.name) } ?? "",
            isPresented: Binding(get: { confirming != nil }, set: { if !$0 { confirming = nil } }),
            presenting: confirming
        ) { row in
            Button(Copy.commonCancel, role: .cancel) {}
            Button(Copy.devicesRevoke, role: .destructive) { onRevoke(row) }
        } message: { _ in
            Text(Copy.devicesRevokeBody)
        }
        .alert(Copy.devicesSignOutOthersTitle, isPresented: $confirmingOthers) {
            Button(Copy.commonCancel, role: .cancel) {}
            Button(Copy.devicesSignOutOthers, role: .destructive, action: onRevokeOthers)
        } message: {
            Text(Copy.devicesSignOutOthersBody)
        }
    }

    @ViewBuilder
    private var content: some View {
        switch state {
        case .loading:
            VStack(spacing: 0) {
                ForEach(0..<3, id: \.self) { _ in
                    FSRow(title: "Placeholder device", sub: "Last used", inset: 0) { EmptyView() }
                }
            }
            .redacted(reason: .placeholder)
        case let .loaded(rows):
            VStack(spacing: 0) {
                ForEach(rows) { row in deviceRow(row) }
            }
        case .failed:
            LoadError(Copy.accountDevicesLoadFailed, retryTitle: Copy.commonTryAgain, retry: onRetry)
        }
    }

    private func deviceRow(_ row: DeviceRow) -> some View {
        FSRow(title: row.name, sub: AccountCopy.deviceMeta(row), inset: 0, trailPlacement: .below, lead: { EmptyView() }) {
            VStack(alignment: .leading, spacing: FSSpace.space2) {
                if row.canRevoke {
                    Button(Copy.devicesRevoke) { confirming = row }
                        .buttonStyle(FSButtonStyle(.danger, fullWidth: false, small: true))
                } else {
                    FSTag(Copy.devicesThisTag)
                }
                if failedRevoke == row.id {
                    SmallText(Copy.devicesRevokeFailed, color: FSColor.wipe)
                }
            }
        }
    }
}

/// Loads `GET /v1/devices`, revokes with `DELETE /v1/devices/{id}`.
struct DevicesScreen: View {
    let session: SessionStore
    let onCountChange: (Int) -> Void
    @Environment(\.openURL) private var openURL
    @State private var state: DevicesState = .loading
    @State private var failedRevoke: String?
    @State private var othersFailed = false

    var body: some View {
        DevicesView(
            state: state, failedRevoke: failedRevoke, othersFailed: othersFailed, onRevoke: revoke, onRevokeOthers: revokeOthers,
            onRetry: { Task { await load() } }
        )
            .task { await load() }
    }

    private func load() async {
        do throws(APIError) {
            let rows = try await session.devices()
            withAnimation { state = .loaded(rows) }
            onCountChange(rows.count)
        } catch {
            state = .failed
        }
    }

    private func revokeOthers() {
        Task {
            switch await session.revokeOthers() {
            case .done:
                othersFailed = false
                await load()
            case .failed, .refused, .confirmMismatch:
                othersFailed = true
            case .needsBrowser:
                openURL(AccountLinks.account)
            }
        }
    }

    private func revoke(_ row: DeviceRow) {
        Task {
            switch await session.revoke(row.id) {
            case .done:
                failedRevoke = nil
                await load()
            case .failed, .refused, .confirmMismatch:
                failedRevoke = row.id
            case .needsBrowser:
                openURL(AccountLinks.account)
            }
        }
    }
}
