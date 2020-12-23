defmodule AoC do
  def read_file do
    {:ok, d} = File.read("in")
    d
    |> String.split("\n", trim: true)
  end

  def traverse(list, offset, jump, to_the_right \\ 0, total_trees \\ 0)
  def traverse([head | tail], offset, 1, to_the_right, total_trees) do
    with "#" <- String.at(head, rem(to_the_right, String.length(head))) do
      traverse(tail, offset, 1, to_the_right + offset, total_trees + 1)
    else _ ->
      traverse(tail, offset, 1, to_the_right + offset, total_trees)
    end
  end
  def traverse([head | [_ | tail]], offset, 2, to_the_right, total_trees) do
    with "#" <- String.at(head, rem(to_the_right, String.length(head))) do
      traverse(tail, offset, 2, to_the_right + offset, total_trees + 1)
    else _ ->
      traverse(tail, offset, 2, to_the_right + offset, total_trees)
    end
  end
  def traverse([head], offset, 2, to_the_right, total_trees) do
    with "#" <- String.at(head, rem(to_the_right, String.length(head))) do
      traverse([], offset, 2, to_the_right + offset, total_trees + 1)
    else _ ->
      traverse([], offset, 2, to_the_right + offset, total_trees)
    end
  end
  def traverse([], _, _, _, total_trees), do: total_trees
end

list = AoC.read_file

IO.puts "#1: " <> Integer.to_string(AoC.traverse(list, 3, 1, 0))
IO.puts "#2: " <> (
  Integer.to_string(
    AoC.traverse(list, 1, 1) *
    AoC.traverse(list, 3, 1) *
    AoC.traverse(list, 5, 1) *
    AoC.traverse(list, 7, 1) *
    AoC.traverse(list, 1, 2)
  )
)
