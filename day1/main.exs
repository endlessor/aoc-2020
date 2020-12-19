defmodule AoC do
  def read_file do
    {:ok, d} = File.read("in")
    d
    |> String.split
    |> Enum.map(&String.to_integer/1)
  end

  def first([head | tail]) do
    with {:false} <- multi_pair(tail, head) do
      first(tail)
    end
  end

  def second([head | tail]) when head > 900, do: second(tail)
  def second([head | [mul | tail]]) when mul > 900, do: second([head | tail])
  def second([head | [mul | tail]]) do
    with {:false} <- multi_group(tail, head, mul) do
      second([mul | tail])
    end
    with {:false} <- multi_group(tail, mul, head) do
      second([head | tail])
    end
  end
  def second(list) when length(list) <= 1, do: {:false}

  defp multi_group([head | _], root, mul) when head + root + mul == 2020 do
    IO.puts "#2: " <> Integer.to_string(head) <> "*" <> Integer.to_string(root) <> "*" <> Integer.to_string(mul) <> "=" <> Integer.to_string(head * root * mul)
  end
  defp multi_group([head | tail], root, mul) when head > 900 do
    multi_group(tail, root, mul)
  end
  defp multi_group([_ | tail], root, mul) do
    multi_group(tail, root, mul)
  end
  defp multi_group([], _, _), do: {:false}

  defp multi_pair([head | _], mul) when head + mul == 2020 do
    IO.puts "#1: " <> Integer.to_string(head) <> "*" <> Integer.to_string(mul) <> "=" <> Integer.to_string(head * mul)
  end
  defp multi_pair([_ | tail], mul) do
    multi_pair(tail, mul)
  end
  defp multi_pair([], _), do: {:false}
end

list = AoC.read_file

AoC.first(list)
AoC.second(list)
