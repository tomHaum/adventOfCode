using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

using System.Drawing;
namespace AdventOfCode2021
{
    internal class Day13
    {

        internal static void Run()
        {
            var pastBreak = false;
            var points = new List<Point>();
            int xMax = 0, yMax = 0;
            var folds = new List<(bool horizontal, int at)>();


            foreach (var line in input.day13.Split("\r\n"))
            {
                if (!pastBreak)
                {
                    if (string.IsNullOrWhiteSpace(line))
                    {

                        pastBreak = true;
                        continue;
                    }
                    else
                    {
                        var split = line.Split(',');
                        var p = new Point(int.Parse(split[0]), int.Parse(split[1]));

                        xMax = Math.Max(xMax, p.X);
                        yMax = Math.Max(yMax, p.Y);
                        points.Add(p);
                    }
                }
                else
                {
                    var split = line.Split('=');
                    var splitAlong = int.Parse(split[1]);

                    var splitHorizontal = split[0].Contains("y");

                    folds.Add((splitHorizontal, splitAlong));
                }
            }

            Print(points, xMax, yMax);

            foreach (var fold in folds)
            {
                for (var i = 0; i < points.Count; i++)
                {
                    var p = points[i];
                    if (fold.horizontal)
                    {
                        if (p.Y > fold.at)
                        {
                            p.Y = fold.at - (p.Y - fold.at);
                        }
                    }
                    else
                    {
                        if (p.X > fold.at)
                        {
                            p.X = fold.at - (p.X - fold.at);
                        }
                    }

                    points[i] = p;
                }
                if (fold.horizontal)
                {
                    yMax = fold.at;
                }
                else
                {
                    xMax = fold.at;
                }
                Console.WriteLine("Fold dot count: " + Print(points, xMax, yMax));
            }

            var accum = Print(points, xMax, yMax, true);

            Console.WriteLine($"Day 13 - Part 1: {1}");

            Console.WriteLine($"Day 13 - Part 2: {2}");
        }

        private static int Print(List<Point> points, int xMax, int yMax, bool print =false)
        {
            var accum = 0;
            for (int y = 0; y <= yMax; y++)
            {
                for (int x = 0; x <= xMax; x++)
                {
                    if (points.Any(p => p.X == x && p.Y == y))
                    {
                        if(print)Console.Write("#");
                        accum++;
                    }
                    else
                    {
                        if (print) Console.Write(".");
                    }
                }
                if (print) Console.WriteLine();
            }
            if (print) Console.WriteLine();
            return accum;
        }
    }
}
