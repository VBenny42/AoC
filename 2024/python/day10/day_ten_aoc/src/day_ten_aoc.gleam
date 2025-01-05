import gleam/io
import gleam/erlang

pub fn read_input() -> Result(List(String), String) {
  case io.read_file("../sample-input.txt") {
    Ok(content) ->
      Ok(string.split(content, "\n")) |> List.map(fn line -> string.trim(line))
    Error(err) ->
      Error("Failed to read file: " <> io.error_to_string(err))
  }
}

pub fn main() {
  case read_input() {
    Ok(lines) -> erlang.print(lines)
    Error(err) -> erlang.print(err)
  }
}
