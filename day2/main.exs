defmodule AoC do
  def read_file do
    {:ok, d} = File.read("in")
    d
    |> String.split("\n", trim: true)
  end

  def first([head | tail], total) do
    with 1 <- passes?(head, ~r/(\d+)-(\d+) ([[:alpha:]]): ([[:alpha:]]+)$/) do
      first(tail, total + 1)
    else _ ->
      first(tail, total)
    end
  end
  def first([], total), do: total

  def second([head | tail], total) do
    with 1 <- positioned?(head, ~r/(\d+)-(\d+) ([[:alpha:]]): ([[:alpha:]]+)$/) do
      second(tail, total + 1)
    else _ ->
      second(tail, total)
    end
  end
  def second([], total), do: total

  defp passes?(str, pattern) do
    [low, high, letter, input] = Regex.run(pattern, str, capture: :all_but_first)
    {:ok, re} = Regex.compile("[^" <> letter <> "]*")
    with(
      {high_int, _} <- Integer.parse(high),
      {low_int, _}  <- Integer.parse(low),
      len           <- String.length(Regex.replace(re, input, "")),
      true          <- high_int >= len and len >= low_int
    ) do
      1
    else
      _ ->
        0
    end
  end

  defp positioned?(str, pattern) do
    [low, high, letter, input] = Regex.run(pattern, str, capture: :all_but_first)
    with(
      {high_int, _} <- Integer.parse(high),
      {low_int, _}  <- Integer.parse(low),
      first_find    <- String.at(input, low_int-1),
      second_find   <- String.at(input, high_int-1),
      true          <- (first_find == letter && second_find != letter) || (first_find != letter && second_find == letter)
    ) do
      1
    else
      _ ->
        0
    end
  end
end

list = AoC.read_file

IO.puts "#1: " <> Integer.to_string(AoC.first(list, 0))
IO.puts "#2: " <> Integer.to_string(AoC.second(list, 0))
