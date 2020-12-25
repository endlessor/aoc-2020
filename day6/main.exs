defmodule AoC do
  def read_file do
    {:ok, d} = File.read("in")
    d
    |> String.split(~r/\n\n/, trim: true)
    |> Enum.map(&(String.split(&1, ~r/\n/)))
  end

  def count_anyone(list, count \\ 0)
  def count_anyone([head | tail], count), do: count_anyone(tail, count + count_letters(head))
  def count_anyone([], count), do: count

  def count_everyone(list, count \\ 0)
  def count_everyone([head | tail], count), do: count_everyone(tail, count + count_matching_letters(head))
  def count_everyone([], count), do: count

  defp count_letters(list, letters \\ %{}, count \\ 0)
  defp count_letters([head | tail], letters, count), do: count_letters(tail, Map.merge(letters, Enum.frequencies(String.split(head, "", trim: true)), fn _, v1, v2 -> v1 + v2 end), count)
  defp count_letters([], letters, _), do: Enum.count(letters, &(&1))

  defp count_matching_letters(list, letters \\ %{}, count \\ 0)
  defp count_matching_letters([head | tail], letters, count), do: count_matching_letters(tail, Map.merge(letters, Enum.frequencies(String.split(head, "", trim: true)), fn _, v1, v2 -> v1 + v2 end), count + 1)
  defp count_matching_letters([], letters, count), do: Enum.count(Enum.reduce(letters, %{}, fn {letter, found}, acc -> validate(letter, found, count, acc) end), &(&1))

  defp validate(letter, found, count, acc) when found == count, do: Map.merge(acc, %{letter => count})
  defp validate(_, _, _, acc), do: acc
end

lines = AoC.read_file

IO.puts "#1: " <> Integer.to_string(AoC.count_anyone(lines))
IO.puts "#2: " <> Integer.to_string(AoC.count_everyone(lines))
