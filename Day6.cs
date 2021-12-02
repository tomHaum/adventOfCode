using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace AdventOfCodeCSharp
{
    class Day6
    {
        public static void Run()
        {
            Console.WriteLine("Day 6");

            var groups = data.day6.Split("\r\n\r\n");

            var sum = 0;
            foreach(var g in groups)
            {
                //Console.WriteLine(new string(g.ToArray()));
                var hs = new HashSet<char>(g.AsEnumerable().Where(x => !char.IsWhiteSpace(x)));
                //Console.WriteLine($"{new string(hs.ToList().OrderBy(x => x).ToArray())}");
                sum += hs.Count;
            }

            Console.WriteLine("Part 1: " + sum);
            
            sum = 0;
            foreach(var g in groups)
            {
                //Console.WriteLine(new string(g.ToArray()));
                var people = g.Split("\r\n");
                var all = new HashSet<char>(people[0]);
                foreach (var person in people.Skip(1)) 
                {
                    all.IntersectWith(person);
                }

                //Console.WriteLine($"{new string(all.ToList().OrderBy(x => x).ToArray())}");
                sum += all.Count;
            }

            Console.WriteLine("Part 2: " + sum);
        }
    }
}
