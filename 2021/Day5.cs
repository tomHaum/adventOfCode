    using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Text.RegularExpressions;
using System.Threading.Tasks;
using System.Drawing;
namespace AdventOfCode2021
{
    internal class Day5
    {
        internal class Line
        {
            private Regex regex = new Regex(@"(\d+),(\d+) -> (\d+),(\d+)");
            internal Line(string txt)
            {
                var match = regex.Match(txt);

                var values = match.Groups.Values.Skip(1).ToArray();

                Start = new Point
                {
                    X = int.Parse(values[0].Value),
                    Y = int.Parse(values[1].Value),
                };

                End = new Point
                {
                    X = int.Parse(values[2].Value),
                    Y = int.Parse(values[3].Value),
                };

            }
            internal Point Start;
            internal Point End;
            internal bool IsStraight
            {
                get
                {
                    return Start.X == End.X || Start.Y == End.Y;
                }
            }

            internal static void Print(List<Line> lines)
            {
                int xMax = 0, yMax = 0;
                foreach (var line in lines)
                {
                    if (line.Start.X > xMax)
                        xMax = line.Start.X;
                    if (line.End.X > xMax)
                        xMax = line.End.X;

                    if (line.Start.Y > yMax)
                        yMax = line.Start.Y;
                    if (line.End.Y > yMax)
                        yMax = line.End.Y;
                }
            }

            internal  IEnumerable<Point> WalkLine()
            {
                int i = 0;
                int xDir = this.Start.X == this.End.X ? 0 : this.Start.X > this.End.X ? -1 : 1;
                int yDir = this.Start.Y == this.End.Y ? 0 : this.Start.Y > this.End.Y ? -1 : 1;
                while (true)
                {
                    var point = new Point { X = this.Start.X + (i * xDir), Y = this.Start.Y + (i * yDir) };
                    yield return point;
                    i++;
                    if (this.End == point)
                        break;
                }
            }
        }

        public static void Run()
        {
            var txt = input.day5.Split("\r\n");
            int xMax = 0, yMax = 0;

            var lines = txt.Select(x => new Line(x)).ToList();

            Dictionary<Point, int> pointValues = new Dictionary<Point, int>();

            var twoOrMore = 0;

            foreach(var line in lines.Where(x => x.IsStraight))
            {
                foreach(var point in line.WalkLine())
                {
                    if(pointValues.TryGetValue(point, out var count))
                    {
                        pointValues[point] = count + 1;
                        if (count == 1)
                            twoOrMore++;
                    }
                    else
                    {
                        pointValues[point] = 1;
                    }
                }
            }

            Console.WriteLine($"Day 5 - Part 1: twoOrMore [{twoOrMore}]");
            pointValues = new Dictionary<Point, int>();
            twoOrMore = 0;

            foreach (var line in lines)
            {
                foreach (var point in line.WalkLine())
                {
                    if (pointValues.TryGetValue(point, out var count))
                    {
                        pointValues[point] = count + 1;
                        if (count == 1)
                            twoOrMore++;
                    }
                    else
                    {
                        pointValues[point] = 1;
                    }
                }
            }
            Console.WriteLine($"Day 5 - Part 2: twoOrMore [{twoOrMore}]");
        }
    }
}
