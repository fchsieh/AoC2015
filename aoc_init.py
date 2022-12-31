import argparse
import os
import subprocess
import sys

try:
    from aocd.models import Puzzle
except ImportError:
    print("Please install aocd: pip install advent-of-code-data")
    sys.exit(1)


def build_folder(day, code_type="py"):
    code_file = "{}.{}".format(day, code_type)
    # create folder
    folder = "day-{}".format(day)
    print("Creating folder: {}".format(folder))
    # create files
    files = ["test.txt", "input.txt", code_file]
    try:
        os.mkdir(folder)
    except FileExistsError:
        print("Folder {} already exists".format(folder))
    for file in files:
        try:
            open(os.path.join(folder, file), "w")
        except FileExistsError:
            print("File {} already exists".format(file))
    return folder


def download(day, year, path):
    print("Downloading input and example data...")
    # check if session key is set
    if "AOC_SESSION" not in os.environ:
        token = subprocess.check_output(["aocd-token".format(day, year)])
        token = token.decode("utf-8").split(" ")[0].strip()
        os.environ["AOC_SESSION"] = token

    puzzle = Puzzle(year=year, day=day)
    # download test
    test_location = os.path.join(path, "test.txt")
    with open(test_location, "w") as f:
        f.write(puzzle.example_data)

    # download input
    input_location = os.path.join(path, "input.txt")
    with open(input_location, "w") as f:
        f.write(puzzle.input_data)


def write_default_code(day, code="py"):
    if code == "py":
        """
        Write default code for python
        """
        with open("day-{}/{}.py".format(day, day), "w") as f:
            f.write("# Path: day-{}/{}.py\n".format(day, day))
            f.write("# Solution for day {} of Advent of Code\n".format(day))
            f.write("import os\n\nfrom aoc_tools import Test\n\n\n")
            f.write("def read_data(file):\n    pass\n\n\n")
            f.write("def solve(data):\n    pass\n\n\n")
            f.write('if __name__ == "__main__":\n')
            f.write("    os.chdir(os.path.dirname(os.path.abspath(__file__)))\n")
            f.write('    test, input = read_data("test.txt"), read_data("input.txt")\n')
            f.write("    # Test(test, solve, -1, -1)\n")
            f.write("    solve(test)\n")
    elif code == "go":
        """
        Write default code for go
        """
        with open("day-{}/{}.go".format(day, day), "w") as f:
            f.write("// Path: day-{}/{}.go\n".format(day, day))
            f.write("// Solution for day {} of Advent of Code\n".format(day))
            f.write("package main\n\nimport (\n")
            f.write('\t"fmt"\n\n\t"github.com/fchsieh/AoC2015/aocutils"\n')
            f.write(")\n\n")
            f.write(
                """func main() {
\tdata := aocutils.ReadPuzzle("input.txt")
\tfmt.Println("Part 1:", part1(data))
\tfmt.Println("Part 2:", part2(data))\n}"""
            )


def set_parser():
    parser = argparse.ArgumentParser(
        prog="Advent of Code Setup",
        description="Setup a new day of AoC",
        epilog="Happy coding!",
    )
    parser.add_argument("--year", "-y", type=int, help="Year of AoC", required=True)
    parser.add_argument("--day", "-d", type=int, help="Day of AoC", required=True)
    parser.add_argument(
        "--code",
        "-c",
        type=str,
        help="set extension of code file",
        required=False,
        default="py",
    )
    return parser


def main(args):
    parser = set_parser()
    args = parser.parse_args(args[1:])

    # setup a new day for aoc
    year = args.year
    day = args.day
    code = args.code
    # Check if folder exists
    if os.path.exists("day-{}".format(day)):
        print("Day {} already exists".format(day))
        sys.exit(1)
    print("Creating new day of AoC: {}-12-{}".format(year, day))

    # Create folder
    folder = build_folder(day, code)
    # Download input and test data
    download(day, year, folder)
    # write default code for day
    write_default_code(day, code)

    print("Done! Happy coding!")


if __name__ == "__main__":
    main(sys.argv)
