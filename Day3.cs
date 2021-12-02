using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace AdventOfCodeCSharp
{
    class Day3
       
    {
        public static void Run()
        {
            Console.WriteLine("Day 3");
            var trees = data.day3.Split(Environment.NewLine).Select(x => x.ToCharArray().Select(y => y == '#').ToArray()).ToArray();

            Console.WriteLine("Part 1: " + collisions(trees, 3, 1));

            Console.WriteLine("Part 2: " + (
                collisions(trees, 1, 1) * collisions(trees, 3, 1) * collisions(trees, 5, 1) * collisions(trees, 7, 1) * collisions(trees, 1, 2)
            ));
        }
        private static long collisions(bool[][] trees, int right, int down)
        {
            int x = 0, y = 0, count = 0;
            do { 
                if (trees[y][x])
                    count++;
                x = (x + right) % trees[0].Length;
                y = y + down;
            } while (y < trees.Length);

            return count;
        }

    }
}
