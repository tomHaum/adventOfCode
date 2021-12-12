using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

using System.Drawing;
namespace AdventOfCode2021
{
    internal class Day12
    {
        internal class Cave
        {
            internal string CaveName;
            internal bool IsBig;
            Cave(string caveName)
            {
                this.CaveName = caveName;
                this.IsBig = char.IsUpper(caveName[0]);
            }

            internal HashSet<Cave> Peers = new HashSet<Cave>();

            internal static Cave GetOrMakeCave(Dictionary<string, Cave> caves, string caveName)
            {
                if (caves.ContainsKey(caveName))
                {
                    return caves[caveName];
                }

                return caves[caveName] = new Cave(caveName);
            }

            internal void Connect(Cave cave)
            {
                this.Peers.Add(cave);
                cave.Peers.Add(this);
            }

            internal static IEnumerable<string> GetPaths(Cave start, Cave end, bool visitTwice)
            {
                var queue = new Queue<(Cave c, string path, bool twice)>();
                queue.Enqueue((start, start.CaveName, false));

                while (queue.Any())
                {
                    var c = queue.Dequeue();
                    if (c.c == end)
                        yield return c.path;

                    foreach (var p in c.c.Peers)
                    {
                        if (!visitTwice)
                        {
                            //part 1
                            if (p.IsBig || !c.path.Contains(p.CaveName))
                                queue.Enqueue((p, c.path + "-" + p.CaveName, false));
                        }
                        else
                        {
                            //part 2
                            if (p.IsBig)
                            {
                                queue.Enqueue((p, c.path + "-" + p.CaveName, c.twice));
                            }
                            else if (!c.twice)
                            {
                                if (!c.path.Contains(p.CaveName))
                                    queue.Enqueue((p, c.path + "-" + p.CaveName, false));
                                else if(p != start && p != end)
                                    queue.Enqueue((p, c.path + "-" + p.CaveName, true));
                            }
                            else
                            {
                                if (!c.path.Contains(p.CaveName))
                                    queue.Enqueue((p, c.path + "-" + p.CaveName, true));
                            }
                        }
                    }
                }
            }

            public override string ToString()
            {
                return this.CaveName;
            }
        }

        internal static void Run()
        {
            Dictionary<string, Cave> caves = new Dictionary<string, Cave>();

            foreach (var line in input.day12.Split("\r\n"))
            {
                var split = line.Split("-");
                var str1 = split[0];
                var str2 = split[1];

                var cave1 = Cave.GetOrMakeCave(caves, str1);
                var cave2 = Cave.GetOrMakeCave(caves, str2);

                cave1.Connect(cave2);
            }

            var start = caves["start"];
            var end = caves["end"];
            var paths = Cave.GetPaths(start, end, false).ToList();
            
            Console.WriteLine($"Day 12 - Part 1: {paths.Count}");

            var accum = 0;
            foreach (var p in Cave.GetPaths(start, end, true))
            {
                accum++;
            }
            Console.WriteLine($"Day 12 - Part 2: {accum}");

        }
    }
}
