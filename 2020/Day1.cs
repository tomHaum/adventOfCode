using System;
using System.Resources;
using System.Reflection;
using System.Threading;
using System.Globalization;

using System.Collections;
using System.Linq;

namespace AdventOfCodeCSharp
{
    public class Day1
    {
        public static void Run()
        {
            Console.WriteLine("Day 1");
            var numbers = data.day1.Split(Environment.NewLine).Select(int.Parse);

            Console.Write("Part 1: ");

            var exit = false;
            foreach (var x in numbers)
            {
                foreach (var y in numbers)
                {
                    if (x + y == 2020)
                    {
                        Console.WriteLine(x*y);
                        exit = true;
                        break;
                    }
                }
                if (exit) break;
            }

            exit = false;
            Console.Write("Part 2: ");

            foreach (var x in numbers)
            {
                foreach (var y in numbers)
                {
                    foreach(var z in numbers)
                    {
                        if (x + y + z == 2020)
                        {
                            Console.WriteLine(x * y * z);
                            exit = true;
                            break;
                        }
                    }
                    if (exit) break;
                }
            }
        }
    }
}
