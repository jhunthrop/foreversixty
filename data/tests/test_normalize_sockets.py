import pytest

from pipeline.normalize.sockets import SocketError, check_no_sockets, socketed_item_ids


def row(item_id: int, *sockets: int) -> dict[str, str]:
    padded = list(sockets) + [0] * (3 - len(sockets))
    return {"ID": str(item_id)} | {
        f"SocketType_{index}": str(value) for index, value in enumerate(padded)
    }


def test_a_client_with_no_socketed_item_passes():
    rows = [row(1), row(2), row(3)]
    assert socketed_item_ids(rows) == []
    check_no_sockets(rows, "1.60.1.69893")


def test_a_row_missing_the_socket_columns_entirely_is_not_socketed():
    """Classic Era's ItemSparse export carries the columns; a trimmed
    fixture or a future client that drops them must read as 'no sockets',
    not raise inside int()."""
    assert socketed_item_ids([{"ID": "7"}]) == []


def test_an_empty_socket_column_is_not_socketed():
    assert socketed_item_ids([{"ID": "7", "SocketType_0": ""}]) == []


def test_socketed_items_are_reported_in_id_order():
    assert socketed_item_ids([row(9, 0, 2), row(4), row(7, 1)]) == [7, 9]


def test_a_socketed_item_stops_the_build_and_names_the_ids():
    with pytest.raises(SocketError) as error:
        check_no_sockets([row(4), row(7, 1), row(9, 0, 0, 3)], "1.61.0.1")
    message = str(error.value)
    assert "1.61.0.1" in message
    assert "7, 9" in message
    assert "2 item(s)" in message


def test_a_long_list_is_truncated_but_counted():
    rows = [row(item_id, 1) for item_id in range(100, 130)]
    with pytest.raises(SocketError) as error:
        check_no_sockets(rows, "1.61.0.1")
    message = str(error.value)
    assert "30 item(s)" in message
    assert "and 10 more" in message
