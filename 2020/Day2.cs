using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace AdventOfCodeCSharp
{
    class Day2
    {
        private static string[] testData = new string[] { "1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc" };

        public static void Run()
        {
            Console.WriteLine("Day 2");
            var lines = data.day2.Split(Environment.NewLine);
            var test = testData.Count(x => isCorrect1(x));
            var ret = lines.Select(isCorrect1).Count(x => x == true);
            Console.WriteLine("Part 1: " + ret);
            test = testData.Count(x => isCorrect2(x));
            ret = lines.Count(x => isCorrect2(x));
            Console.WriteLine("Part 2: " + ret);
        }

        private static bool isCorrect1(string s)
        {
            var dash = s.IndexOf("-");
            var space = s.IndexOf(" ");
            var colon = s.IndexOf(":");

            var lower = int.Parse(s.Substring(0, s.IndexOf("-")));
            var upper = int.Parse(s.Substring(dash + 1, space - dash - 1));
            var lookFor = s[space + 1];
            var pass = s.Substring(colon + 2);

            var count = pass.Count(x => x == lookFor);
            var ret = count >= lower && count <= upper;
            return ret;
        }

        private static bool isCorrect2(string s)
        {
            var dash = s.IndexOf("-");
            var space = s.IndexOf(" ");
            var colon = s.IndexOf(":");

            var lowerIndex = int.Parse(s.Substring(0, s.IndexOf("-")));
            var upperIndex = int.Parse(s.Substring(dash + 1, space - dash - 1));
            var lookFor = s[space + 1];
            var pass = s.Substring(colon + 2);

            var ret = pass[lowerIndex - 1] == lookFor ^ pass[upperIndex - 1] == lookFor;
            return ret;
        }
    }
}
