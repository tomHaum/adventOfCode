using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

using System.Drawing;
namespace AdventOfCode2021
{
    internal class Day9
    {


        internal static void Run()
        {

            int[][] map = input.day9.Split("\r\n").Select(x => x.Select(y => int.Parse(y.ToString())).ToArray()).ToArray();
            var mins = new List<int>();

            for (int x = 0; x < map.Length; x++)
            {
                for (int y = 0; y < map[0].Length; y++)
                {
                    var current = map[x][y];
                    //checkLeft
                    if (x != 0 && current >= map[x - 1][y])
                        continue;
                    //checkRight
                    if (x != map.Length - 1 && current >= map[x + 1][y])
                        continue;
                    //checkTop
                    if (y != 0 && current >= map[x][y - 1])
                        continue;
                    //checkBot
                    if (y != map[0].Length - 1 && current >= map[x][y + 1])
                        continue;
                    mins.Add(current);
                }
            }

            Console.WriteLine($"Day 9 - Part 1: {mins.Sum() + mins.Count()}");

            mins = new List<int>();
            var basinOrgins = new List<Point>();
            for (int x = 0; x < map.Length; x++)
            {
                for (int y = 0; y < map[0].Length; y++)
                {
                    var current = map[x][y];
                    //checkLeft
                    if (x != 0 && current >= map[x - 1][y])
                        continue;
                    //checkRight
                    if (x != map.Length - 1 && current >= map[x + 1][y])
                        continue;
                    //checkTop
                    if (y != 0 && current >= map[x][y - 1])
                        continue;
                    //checkBot
                    if (y != map[0].Length - 1 && current >= map[x][y + 1])
                        continue;
                    mins.Add(current);
                    basinOrgins.Add(new Point() { X = x, Y = y });
                }
            }

            var basins = new List<HashSet<Point>>();
            foreach (var basin in basinOrgins)
            {
                var visited = new HashSet<Point>();
                var active = new Queue<Point>();
                active.Enqueue(basin);

                while (active.Any())
                {
                    var currentTile = active.Dequeue();
                    visited.Add(currentTile);
                    var x = currentTile.X;
                    var y = currentTile.Y;

                    //checkLeft
                    if (x != 0 && 9 != map[x - 1][y])
                    {
                        var p = new Point(x - 1, y);
                        if (!visited.Contains(p))
                        {
                            active.Enqueue(p);
                        }
                    }
                    //checkRight
                    if (x != map.Length - 1 && 9 != map[x + 1][y])
                    {
                        var p = new Point(x + 1, y);
                        if (!visited.Contains(p))
                        {
                            active.Enqueue(p);
                        }
                    }
                    //checkTop
                    if (y != 0 && 9 != map[x][y - 1])
                    {
                        var p = new Point(x, y - 1);
                        if (!visited.Contains(p))
                        {
                            active.Enqueue(p);
                        }
                    }
                    //checkBot
                    if (y != map[0].Length - 1 && 9 != map[x][y + 1])
                    {
                        var p = new Point(x, y + 1);
                        if (!visited.Contains(p))
                        {
                            active.Enqueue(p);
                        }
                    }
                }
                basins.Add(visited);
            }
            var product = basins.OrderByDescending(x => x.Count).Take(3).Aggregate(1, (accum, x) => accum * x.Count);
            Console.WriteLine($"Day 9 - Part 2: {product}");
        }

    }
}
