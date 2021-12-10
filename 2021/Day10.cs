using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

using System.Drawing;
namespace AdventOfCode2021
{
    internal class Day10
    {
        private static Dictionary<char, char> closers = new Dictionary<char, char>();
        private static Dictionary<char, int> part1Scores = new Dictionary<char, int>();
        private static Dictionary<char, int> part2Scores = new Dictionary<char, int>();

        static Day10()
        {
            closers.Add(')', '(');
            closers.Add(']', '[');
            closers.Add('}', '{');
            closers.Add('>', '<');

            part1Scores.Add(')', 3);
            part1Scores.Add(']', 57);
            part1Scores.Add('}', 1197);
            part1Scores.Add('>', 25137);

            part2Scores.Add('(', 1);
            part2Scores.Add('[', 2);
            part2Scores.Add('{', 3);
            part2Scores.Add('<', 4);
        }
        
           
        internal static void Run()
        {
            var lines = input.day10.Split("\r\n").ToArray();

            var part1Score = 0;
            var part2Results = new List<long>();
            foreach(var line in lines)
            {
                Dictionary<char, int> openers = new Dictionary<char, int>();
                Stack<char> stack = new Stack<char>();
                var failed = false;
                foreach (var c in line)
                {
                    if (closers.ContainsKey(c))
                    {
                        if(stack.Peek() == closers[c]) {
                            stack.Pop();
                        }
                        else
                        {
                            part1Score += part1Scores[c];
                            failed = true;
                            break;
                        }
                    }
                    else
                    {
                        stack.Push(c);
                    }
                }
                if (failed)
                {
                    continue;
                }
                long score = 0;
                while(stack.Count != 0)
                {
                    var c = stack.Pop();
                    score = score * 5;
                    score += part2Scores[c];
                }
                part2Results.Add(score);
            }

            Console.WriteLine($"Day 10 - Part 1: {part1Score}");

            var part2 = part2Results.OrderBy(x => x).Skip(part2Results.Count / 2).First();
            Console.WriteLine($"Day 10 - Part 2: {part2}");
        }

    }
}
