defmodule AoC do
  def read_file do
    {:ok, d} = File.read("in")
    d
    |> String.split("\n", trim: true)
  end

  def get_highest(list, max \\ 0, seats \\ [])
  def get_highest([head | tail], max, seats) do
    {row, seat} = tuple_space(head)
    {int_row, ""} = Integer.parse(row_map(row), 2)
    {int_seat, ""} = Integer.parse(seat_map(seat), 2)
    new_max = Kernel.max(max, (int_row * 8) + int_seat)

    get_highest(tail, new_max, [((int_row * 8) + int_seat) | seats])
  end
  def get_highest([], max, seats) do
    IO.puts "#2: " <> Integer.to_string(missing_seat(Enum.sort(seats), 0))

    max
  end

  defp tuple_space(alloc), do: {String.slice(alloc, 0..6), String.slice(alloc, -3..-1)}
  defp row_map(row), do: String.replace(String.replace(row, "B", "1"), "F", "0")
  defp seat_map(seat), do: String.replace(String.replace(seat, "R", "1"), "L", "0")

  defp missing_seat([head | [second | tail]], missing) do
    with 1 <- second - head do
      missing_seat([second | tail], missing)
    else _ ->
      missing_seat([], head+1)
    end
  end
  defp missing_seat(_, missing), do: missing
end

lines = AoC.read_file

IO.puts "#1: " <> Integer.to_string(AoC.get_highest(lines))
