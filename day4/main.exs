defmodule AoC do
  def read_file do
    {:ok, d} = File.read("in")
    d
    |> String.replace(~r/\n\n/, ":::")
    |> String.replace(~r/\n/, " ")
    |> String.split(":::", trim: true)
  end

  def basic_count([head | tail], all) do
    with 7 <- head |> String.split(" ") |> total_valid_attributes do
      basic_count(tail, all+1)
    else _ ->
      basic_count(tail, all)
    end
  end
  def basic_count([], all), do: all

  defp total_valid_attributes(list, valid \\ 0)
  defp total_valid_attributes([head | tail], valid) do
    with 1 <- valid(String.split(head, ~r/:/)) do
      total_valid_attributes(tail, valid+1)
    else _ ->
      total_valid_attributes(tail, valid)
    end
  end
  defp total_valid_attributes([], valid), do: valid

  defp valid(list)
  defp valid(["byr", _]), do: 1
  defp valid(["iyr", _]), do: 1
  defp valid(["eyr", _]), do: 1
  defp valid(["hgt", _]), do: 1
  defp valid(["hcl", _]), do: 1
  defp valid(["ecl", _]), do: 1
  defp valid(["pid", _]), do: 1
  defp valid(["cid", _]), do: 0
  defp valid([_, _]), do: 0
  defp valid([""]), do: 0

  def strict_count([head | tail], all) do
    with 7 <- head |> String.split(" ") |> proper_valid_attributes do
      strict_count(tail, all+1)
    else _ ->
      strict_count(tail, all)
    end
  end
  def strict_count([], all), do: all

  defp proper_valid_attributes(list, valid \\ 0)
  defp proper_valid_attributes([head | tail], valid) do
    with 1 <- proper_valid(String.split(head, ~r/:/)) do
      proper_valid_attributes(tail, valid+1)
    else _ ->
      proper_valid_attributes(tail, valid)
    end
  end
  defp proper_valid_attributes([], valid), do: valid

  defp proper_valid(list)
  defp proper_valid(["byr", value]) do
    {parsed, ""} = Integer.parse(value)
    with true <- parsed > 1919 and parsed < 2003 do
      1
    else _ ->
      0
    end
  end
  defp proper_valid(["iyr", value]) do
    {parsed, ""} = Integer.parse(value)
    with true <- parsed > 2009 and parsed < 2021 do
      1
    else _ ->
      0
    end
  end
  defp proper_valid(["eyr", value]) do
    {parsed, ""} = Integer.parse(value)
    with true <- parsed > 2019 and parsed < 2031 do
      1
    else _ ->
      0
    end
  end
  defp proper_valid(["hgt", value]) do
    match_height(Integer.parse(value))
  end
  defp proper_valid(["hcl", value]) do
    match_colour(String.split(value, "#"))
  end
  defp proper_valid(["ecl", value]) when value in ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"], do: 1
  defp proper_valid(["pid", value]) do
    with 9 <- String.length(value) do
      1
    else _ ->
      0
    end
  end
  defp proper_valid(_), do: 0

  defp match_height({value, "in"}) when value > 58 and value < 77, do: 1
  defp match_height({value, "cm"}) when value > 149 and value < 194, do: 1
  defp match_height(_), do: 0

  defp match_colour(["", value]) do
    with true <- Regex.match?(~r/[a-z0-9]{6}/, value) do
      1
    else _ ->
      0
    end
  end
  defp match_colour(_), do: 0
end

list = AoC.read_file

IO.puts "#1: " <> Integer.to_string(AoC.basic_count(list, 0))
IO.puts "#2: " <> Integer.to_string(AoC.strict_count(list, 0))
