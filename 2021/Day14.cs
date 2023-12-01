using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

using System.Drawing;
using System.Text.RegularExpressions;
namespace AdventOfCode2021
{
    internal class Day14
    {
        private static Regex ruleMatcher = new Regex(@"(\w{2}) \-> (\w)");
        internal static void Run()
        {
            var split = input.day14.Split("\r\n").ToArray();
            var polymerTemplateTemp = split[0];

            var polymerTemplate = new List<char>(2146435071);
            foreach (var c in polymerTemplateTemp) polymerTemplate.Add(c);

            var rules = new Dictionary<(char, char), char>();

            foreach(var line in split.Skip(2))
            {
                var x = line[0];
                var y = line[1];

                var z = line[6];

                rules.Add((x, y), z);
            }
            var count = 40;
            while (count > 0)
            {
                var char1 = polymerTemplate[0];
                for (int i = 1; i < polymerTemplate.Count; i++)
                {
                    var char2 = polymerTemplate[i];

                    if (rules.TryGetValue((char1, char2), out var insertChar))
                    {
                        polymerTemplate.Insert(i, insertChar);
                        i++;
                    }
                    char1 = char2;
                }
                count--;
                if(count == 30)
                {
                    var group = polymerTemplate.GroupBy(x => x).OrderByDescending(x => x.Count()).ToList();

                    var first = group.First();
                    var last = group.Last();

                    Console.WriteLine($"Day 14 - Part 1: {first.Key} [{first.Count()}] {last.Key} [{last.Count()}] | [{first.Count() - last.Count()}]");
                }
                Console.WriteLine("count" + count);
            }
            var group2 = polymerTemplate.GroupBy(x => x).OrderByDescending(x => x.Count()).ToList();

            var first2 = group2.First();
            var last2 = group2.Last();

            Console.WriteLine($"Day 14 - Part 2: {first2.Key} [{first2.Count()}] {last2.Key} [{last2.Count()}] | [{first2.Count() - last2.Count()}]");
        }

    }
}
